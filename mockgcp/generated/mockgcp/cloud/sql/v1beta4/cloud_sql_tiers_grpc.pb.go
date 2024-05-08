// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/sql/v1beta4/cloud_sql_tiers.proto

package sqlpb

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

// SqlTiersServiceClient is the client API for SqlTiersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SqlTiersServiceClient interface {
	// Lists all available machine types (tiers) for Cloud SQL, for example,
	// `db-custom-1-3840`. For related information, see [Pricing](/sql/pricing).
	List(ctx context.Context, in *SqlTiersListRequest, opts ...grpc.CallOption) (*TiersListResponse, error)
}

type sqlTiersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSqlTiersServiceClient(cc grpc.ClientConnInterface) SqlTiersServiceClient {
	return &sqlTiersServiceClient{cc}
}

func (c *sqlTiersServiceClient) List(ctx context.Context, in *SqlTiersListRequest, opts ...grpc.CallOption) (*TiersListResponse, error) {
	out := new(TiersListResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.sql.v1beta4.SqlTiersService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SqlTiersServiceServer is the server API for SqlTiersService service.
// All implementations must embed UnimplementedSqlTiersServiceServer
// for forward compatibility
type SqlTiersServiceServer interface {
	// Lists all available machine types (tiers) for Cloud SQL, for example,
	// `db-custom-1-3840`. For related information, see [Pricing](/sql/pricing).
	List(context.Context, *SqlTiersListRequest) (*TiersListResponse, error)
	mustEmbedUnimplementedSqlTiersServiceServer()
}

// UnimplementedSqlTiersServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSqlTiersServiceServer struct {
}

func (UnimplementedSqlTiersServiceServer) List(context.Context, *SqlTiersListRequest) (*TiersListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSqlTiersServiceServer) mustEmbedUnimplementedSqlTiersServiceServer() {}

// UnsafeSqlTiersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SqlTiersServiceServer will
// result in compilation errors.
type UnsafeSqlTiersServiceServer interface {
	mustEmbedUnimplementedSqlTiersServiceServer()
}

func RegisterSqlTiersServiceServer(s grpc.ServiceRegistrar, srv SqlTiersServiceServer) {
	s.RegisterService(&SqlTiersService_ServiceDesc, srv)
}

func _SqlTiersService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqlTiersListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SqlTiersServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.sql.v1beta4.SqlTiersService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SqlTiersServiceServer).List(ctx, req.(*SqlTiersListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SqlTiersService_ServiceDesc is the grpc.ServiceDesc for SqlTiersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SqlTiersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.sql.v1beta4.SqlTiersService",
	HandlerType: (*SqlTiersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _SqlTiersService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/cloud/sql/v1beta4/cloud_sql_tiers.proto",
}