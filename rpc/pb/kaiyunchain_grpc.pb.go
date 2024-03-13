// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: kaiyunchain.proto

package pb

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
	Rpc_GetLatestBlockHeight_FullMethodName = "/Rpc/GetLatestBlockHeight"
	Rpc_GetAllBlock_FullMethodName          = "/Rpc/GetAllBlock"
	Rpc_GetBlock_FullMethodName             = "/Rpc/GetBlock"
	Rpc_GetLatestTxNum_FullMethodName       = "/Rpc/GetLatestTxNum"
	Rpc_GetAllTx_FullMethodName             = "/Rpc/GetAllTx"
	Rpc_GetTx_FullMethodName                = "/Rpc/GetTx"
	Rpc_NewAccount_FullMethodName           = "/Rpc/NewAccount"
)

// RpcClient is the client API for Rpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcClient interface {
	GetLatestBlockHeight(ctx context.Context, in *GetLatestBlockHeightReq, opts ...grpc.CallOption) (*GetLatestBlockHeightResp, error)
	GetAllBlock(ctx context.Context, in *GetAllBlockReq, opts ...grpc.CallOption) (*GetAllBlockResp, error)
	GetBlock(ctx context.Context, in *GetBlockReq, opts ...grpc.CallOption) (*GetBlockResp, error)
	GetLatestTxNum(ctx context.Context, in *GetLatestTxNumReq, opts ...grpc.CallOption) (*GetLatestTxNumResp, error)
	GetAllTx(ctx context.Context, in *GetAllTxReq, opts ...grpc.CallOption) (*GetAllTxResp, error)
	GetTx(ctx context.Context, in *GetTxReq, opts ...grpc.CallOption) (*GetTxResp, error)
	NewAccount(ctx context.Context, in *NewAccountReq, opts ...grpc.CallOption) (*NewAccountResp, error)
}

type rpcClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcClient(cc grpc.ClientConnInterface) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) GetLatestBlockHeight(ctx context.Context, in *GetLatestBlockHeightReq, opts ...grpc.CallOption) (*GetLatestBlockHeightResp, error) {
	out := new(GetLatestBlockHeightResp)
	err := c.cc.Invoke(ctx, Rpc_GetLatestBlockHeight_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) GetAllBlock(ctx context.Context, in *GetAllBlockReq, opts ...grpc.CallOption) (*GetAllBlockResp, error) {
	out := new(GetAllBlockResp)
	err := c.cc.Invoke(ctx, Rpc_GetAllBlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) GetBlock(ctx context.Context, in *GetBlockReq, opts ...grpc.CallOption) (*GetBlockResp, error) {
	out := new(GetBlockResp)
	err := c.cc.Invoke(ctx, Rpc_GetBlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) GetLatestTxNum(ctx context.Context, in *GetLatestTxNumReq, opts ...grpc.CallOption) (*GetLatestTxNumResp, error) {
	out := new(GetLatestTxNumResp)
	err := c.cc.Invoke(ctx, Rpc_GetLatestTxNum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) GetAllTx(ctx context.Context, in *GetAllTxReq, opts ...grpc.CallOption) (*GetAllTxResp, error) {
	out := new(GetAllTxResp)
	err := c.cc.Invoke(ctx, Rpc_GetAllTx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) GetTx(ctx context.Context, in *GetTxReq, opts ...grpc.CallOption) (*GetTxResp, error) {
	out := new(GetTxResp)
	err := c.cc.Invoke(ctx, Rpc_GetTx_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) NewAccount(ctx context.Context, in *NewAccountReq, opts ...grpc.CallOption) (*NewAccountResp, error) {
	out := new(NewAccountResp)
	err := c.cc.Invoke(ctx, Rpc_NewAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcServer is the server API for Rpc service.
// All implementations must embed UnimplementedRpcServer
// for forward compatibility
type RpcServer interface {
	GetLatestBlockHeight(context.Context, *GetLatestBlockHeightReq) (*GetLatestBlockHeightResp, error)
	GetAllBlock(context.Context, *GetAllBlockReq) (*GetAllBlockResp, error)
	GetBlock(context.Context, *GetBlockReq) (*GetBlockResp, error)
	GetLatestTxNum(context.Context, *GetLatestTxNumReq) (*GetLatestTxNumResp, error)
	GetAllTx(context.Context, *GetAllTxReq) (*GetAllTxResp, error)
	GetTx(context.Context, *GetTxReq) (*GetTxResp, error)
	NewAccount(context.Context, *NewAccountReq) (*NewAccountResp, error)
	mustEmbedUnimplementedRpcServer()
}

// UnimplementedRpcServer must be embedded to have forward compatible implementations.
type UnimplementedRpcServer struct {
}

func (UnimplementedRpcServer) GetLatestBlockHeight(context.Context, *GetLatestBlockHeightReq) (*GetLatestBlockHeightResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestBlockHeight not implemented")
}
func (UnimplementedRpcServer) GetAllBlock(context.Context, *GetAllBlockReq) (*GetAllBlockResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllBlock not implemented")
}
func (UnimplementedRpcServer) GetBlock(context.Context, *GetBlockReq) (*GetBlockResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (UnimplementedRpcServer) GetLatestTxNum(context.Context, *GetLatestTxNumReq) (*GetLatestTxNumResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestTxNum not implemented")
}
func (UnimplementedRpcServer) GetAllTx(context.Context, *GetAllTxReq) (*GetAllTxResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTx not implemented")
}
func (UnimplementedRpcServer) GetTx(context.Context, *GetTxReq) (*GetTxResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTx not implemented")
}
func (UnimplementedRpcServer) NewAccount(context.Context, *NewAccountReq) (*NewAccountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAccount not implemented")
}
func (UnimplementedRpcServer) mustEmbedUnimplementedRpcServer() {}

// UnsafeRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcServer will
// result in compilation errors.
type UnsafeRpcServer interface {
	mustEmbedUnimplementedRpcServer()
}

func RegisterRpcServer(s grpc.ServiceRegistrar, srv RpcServer) {
	s.RegisterService(&Rpc_ServiceDesc, srv)
}

func _Rpc_GetLatestBlockHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestBlockHeightReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).GetLatestBlockHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_GetLatestBlockHeight_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).GetLatestBlockHeight(ctx, req.(*GetLatestBlockHeightReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_GetAllBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllBlockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).GetAllBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_GetAllBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).GetAllBlock(ctx, req.(*GetAllBlockReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_GetBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).GetBlock(ctx, req.(*GetBlockReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_GetLatestTxNum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestTxNumReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).GetLatestTxNum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_GetLatestTxNum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).GetLatestTxNum(ctx, req.(*GetLatestTxNumReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_GetAllTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllTxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).GetAllTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_GetAllTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).GetAllTx(ctx, req.(*GetAllTxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_GetTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTxReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).GetTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_GetTx_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).GetTx(ctx, req.(*GetTxReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_NewAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAccountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).NewAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rpc_NewAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).NewAccount(ctx, req.(*NewAccountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Rpc_ServiceDesc is the grpc.ServiceDesc for Rpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLatestBlockHeight",
			Handler:    _Rpc_GetLatestBlockHeight_Handler,
		},
		{
			MethodName: "GetAllBlock",
			Handler:    _Rpc_GetAllBlock_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _Rpc_GetBlock_Handler,
		},
		{
			MethodName: "GetLatestTxNum",
			Handler:    _Rpc_GetLatestTxNum_Handler,
		},
		{
			MethodName: "GetAllTx",
			Handler:    _Rpc_GetAllTx_Handler,
		},
		{
			MethodName: "GetTx",
			Handler:    _Rpc_GetTx_Handler,
		},
		{
			MethodName: "NewAccount",
			Handler:    _Rpc_NewAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kaiyunchain.proto",
}
