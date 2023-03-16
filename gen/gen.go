package gen

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const (
	optPackage     = "package"
	optSamePackage = "same_package"
)

type Generator struct {
	genp        *protogen.Plugin
	samePackage bool
	packageName string
}

func New(genp *protogen.Plugin) *Generator {
	genp.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
	return &Generator{
		genp: genp,
	}
}

func (gen *Generator) Generate() error {
	if err := gen.parseOpts(); err != nil {
		return err
	}
	var tmppf *protogen.File
	for _, pf := range gen.genp.Files {
		if !pf.Generate {
			continue
		}
		tmppf = pf
	}
	filename := filepath.Join(filepath.Dir(tmppf.GeneratedFilenamePrefix), "client.go")
	g := gen.genp.NewGeneratedFile(filename, tmppf.GoImportPath)
	if gen.packageName != "" {
		g.P("package ", gen.packageName)
	} else {
		g.P("package ", tmppf.GoPackageName)
	}
	g.P("")
	g.P(`import (`)
	if !gen.samePackage {
		g.P(fmt.Sprintf("%s %s", tmppf.GoPackageName, tmppf.GoImportPath.String()))
	}
	g.P(`"google.golang.org/grpc"`)
	g.P(`)`)

	g.P(`type Client struct {`)
	for _, pf := range gen.genp.Files {
		if !pf.Generate {
			continue
		}
		for _, s := range pf.Services {
			if gen.samePackage {
				g.P(fmt.Sprintf("%s %sClient", s.Desc.FullName().Name(), s.Desc.FullName().Name()))
			} else {
				g.P(fmt.Sprintf("%s %s.%sClient", s.Desc.FullName().Name(), tmppf.GoPackageName, s.Desc.FullName().Name()))
			}
		}
	}
	g.P(`}`)

	g.P(`func New(cc grpc.ClientConnInterface) *Client {`)
	g.P(`return &Client{`)
	for _, pf := range gen.genp.Files {
		if !pf.Generate {
			continue
		}
		for _, s := range pf.Services {
			if gen.samePackage {
				g.P(fmt.Sprintf("%s: New%sClient(cc),", s.Desc.FullName().Name(), s.Desc.FullName().Name()))
			} else {
				g.P(fmt.Sprintf("%s: %s.New%sClient(cc),", s.Desc.FullName().Name(), tmppf.GoPackageName, s.Desc.FullName().Name()))
			}
		}
	}
	g.P(`}`)
	g.P(`}`)
	g.P("")
	return nil
}

func (gen *Generator) parseOpts() error {
	opts := strings.Split(gen.genp.Request.GetParameter(), ",")
	for _, o := range opts {
		o := strings.TrimSpace(o)
		switch {
		case o == optSamePackage:
			gen.samePackage = true
		case strings.HasPrefix(o, fmt.Sprintf("%s=", optPackage)):
			gen.packageName = strings.TrimPrefix(o, fmt.Sprintf("%s=", optPackage))
		}
	}
	if gen.samePackage && gen.packageName != "" {
		return errors.New("package name cannot be specified if it is the same package as the Go package")
	}
	return nil
}
