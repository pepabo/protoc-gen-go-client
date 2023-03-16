// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: example/project.proto

package example

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ProjectService_ListProjects_FullMethodName  = "/example.ProjectService/ListProjects"
	ProjectService_CreateProject_FullMethodName = "/example.ProjectService/CreateProject"
)

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectServiceClient interface {
	ListProjects(ctx context.Context, in *DummyProjectRequest, opts ...grpc.CallOption) (ProjectService_ListProjectsClient, error)
	CreateProject(ctx context.Context, in *DummyProjectRequest, opts ...grpc.CallOption) (*DummyProjectResponse, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) ListProjects(ctx context.Context, in *DummyProjectRequest, opts ...grpc.CallOption) (ProjectService_ListProjectsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProjectService_ServiceDesc.Streams[0], ProjectService_ListProjects_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &projectServiceListProjectsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProjectService_ListProjectsClient interface {
	Recv() (*DummyProjectResponse, error)
	grpc.ClientStream
}

type projectServiceListProjectsClient struct {
	grpc.ClientStream
}

func (x *projectServiceListProjectsClient) Recv() (*DummyProjectResponse, error) {
	m := new(DummyProjectResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *projectServiceClient) CreateProject(ctx context.Context, in *DummyProjectRequest, opts ...grpc.CallOption) (*DummyProjectResponse, error) {
	out := new(DummyProjectResponse)
	err := c.cc.Invoke(ctx, ProjectService_CreateProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations must embed UnimplementedProjectServiceServer
// for forward compatibility
type ProjectServiceServer interface {
	ListProjects(*DummyProjectRequest, ProjectService_ListProjectsServer) error
	CreateProject(context.Context, *DummyProjectRequest) (*DummyProjectResponse, error)
	mustEmbedUnimplementedProjectServiceServer()
}

// UnimplementedProjectServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProjectServiceServer struct {
}

func (UnimplementedProjectServiceServer) ListProjects(*DummyProjectRequest, ProjectService_ListProjectsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListProjects not implemented")
}
func (UnimplementedProjectServiceServer) CreateProject(context.Context, *DummyProjectRequest) (*DummyProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedProjectServiceServer) mustEmbedUnimplementedProjectServiceServer() {}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_ListProjects_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DummyProjectRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProjectServiceServer).ListProjects(m, &projectServiceListProjectsServer{stream})
}

type ProjectService_ListProjectsServer interface {
	Send(*DummyProjectResponse) error
	grpc.ServerStream
}

type projectServiceListProjectsServer struct {
	grpc.ServerStream
}

func (x *projectServiceListProjectsServer) Send(m *DummyProjectResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProjectService_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DummyProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_CreateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).CreateProject(ctx, req.(*DummyProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "example.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProject",
			Handler:    _ProjectService_CreateProject_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListProjects",
			Handler:       _ProjectService_ListProjects_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "example/project.proto",
}
