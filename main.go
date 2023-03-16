package main

import (
	"github.com/pepabo/protoc-gen-go-client/gen"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(genp *protogen.Plugin) error {
		g := gen.New(genp)
		return g.Generate()
	})
}
