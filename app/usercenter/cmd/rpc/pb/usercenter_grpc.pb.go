// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: usercenter.proto

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
	Usercenter_Login_FullMethodName                = "/pb.usercenter/login"
	Usercenter_Register_FullMethodName             = "/pb.usercenter/register"
	Usercenter_GetUserInfo_FullMethodName          = "/pb.usercenter/getUserInfo"
	Usercenter_GetUserAuthByAuthKey_FullMethodName = "/pb.usercenter/getUserAuthByAuthKey"
	Usercenter_GetUserAuthByUserId_FullMethodName  = "/pb.usercenter/getUserAuthByUserId"
	Usercenter_GenerateToken_FullMethodName        = "/pb.usercenter/generateToken"
	Usercenter_SetAdmin_FullMethodName             = "/pb.usercenter/setAdmin"
	Usercenter_CheckIsAdmin_FullMethodName         = "/pb.usercenter/CheckIsAdmin"
	Usercenter_AddUserContact_FullMethodName       = "/pb.usercenter/AddUserContact"
	Usercenter_UpdateUserContact_FullMethodName    = "/pb.usercenter/UpdateUserContact"
	Usercenter_DelUserContact_FullMethodName       = "/pb.usercenter/DelUserContact"
	Usercenter_GetUserContactById_FullMethodName   = "/pb.usercenter/GetUserContactById"
	Usercenter_SearchUserContact_FullMethodName    = "/pb.usercenter/SearchUserContact"
	Usercenter_AddUserAddress_FullMethodName       = "/pb.usercenter/AddUserAddress"
	Usercenter_UpdateUserAddress_FullMethodName    = "/pb.usercenter/UpdateUserAddress"
	Usercenter_DelUserAddress_FullMethodName       = "/pb.usercenter/DelUserAddress"
	Usercenter_GetUserAddressById_FullMethodName   = "/pb.usercenter/GetUserAddressById"
	Usercenter_SearchUserAddress_FullMethodName    = "/pb.usercenter/SearchUserAddress"
)

// UsercenterClient is the client API for Usercenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UsercenterClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
	GetUserAuthByAuthKey(ctx context.Context, in *GetUserAuthByAuthKeyReq, opts ...grpc.CallOption) (*GetUserAuthByAuthKeyResp, error)
	GetUserAuthByUserId(ctx context.Context, in *GetUserAuthByUserIdReq, opts ...grpc.CallOption) (*GetUserAuthyUserIdResp, error)
	GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
	UpdateUserBaseInfo(ctx context.Context, in *UpdateUserBaseInfoReq, opts ...grpc.CallOption) (*UpdateUserBaseInfoResp, error)
	// -----------------------用户联系方式----------------------
	SetAdmin(ctx context.Context, in *SetAdminReq, opts ...grpc.CallOption) (*SetAdminResp, error)
	CheckIsAdmin(ctx context.Context, in *CheckIsAdminReq, opts ...grpc.CallOption) (*CheckIsAdminResp, error)
	// -----------------------用户联系方式----------------------
	AddUserContact(ctx context.Context, in *AddUserContactReq, opts ...grpc.CallOption) (*AddUserContactResp, error)
	UpdateUserContact(ctx context.Context, in *UpdateUserContactReq, opts ...grpc.CallOption) (*UpdateUserContactResp, error)
	DelUserContact(ctx context.Context, in *DelUserContactReq, opts ...grpc.CallOption) (*DelUserContactResp, error)
	GetUserContactById(ctx context.Context, in *GetUserContactByIdReq, opts ...grpc.CallOption) (*GetUserContactByIdResp, error)
	SearchUserContact(ctx context.Context, in *SearchUserContactReq, opts ...grpc.CallOption) (*SearchUserContactResp, error)
	// -----------------------用户收货地址表-----------------------
	AddUserAddress(ctx context.Context, in *AddUserAddressReq, opts ...grpc.CallOption) (*AddUserAddressResp, error)
	UpdateUserAddress(ctx context.Context, in *UpdateUserAddressReq, opts ...grpc.CallOption) (*UpdateUserAddressResp, error)
	DelUserAddress(ctx context.Context, in *DelUserAddressReq, opts ...grpc.CallOption) (*DelUserAddressResp, error)
	GetUserAddressById(ctx context.Context, in *GetUserAddressByIdReq, opts ...grpc.CallOption) (*GetUserAddressByIdResp, error)
	SearchUserAddress(ctx context.Context, in *SearchUserAddressReq, opts ...grpc.CallOption) (*SearchUserAddressResp, error)
}

type usercenterClient struct {
	cc grpc.ClientConnInterface
}

func NewUsercenterClient(cc grpc.ClientConnInterface) UsercenterClient {
	return &usercenterClient{cc}
}

func (c *usercenterClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	out := new(GetUserInfoResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/getUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) GetUserAuthByAuthKey(ctx context.Context, in *GetUserAuthByAuthKeyReq, opts ...grpc.CallOption) (*GetUserAuthByAuthKeyResp, error) {
	out := new(GetUserAuthByAuthKeyResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/getUserAuthByAuthKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) GetUserAuthByUserId(ctx context.Context, in *GetUserAuthByUserIdReq, opts ...grpc.CallOption) (*GetUserAuthyUserIdResp, error) {
	out := new(GetUserAuthyUserIdResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/getUserAuthByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	out := new(GenerateTokenResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/generateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) UpdateUserBaseInfo(ctx context.Context, in *UpdateUserBaseInfoReq, opts ...grpc.CallOption) (*UpdateUserBaseInfoResp, error) {
	out := new(UpdateUserBaseInfoResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/updateUserBaseInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) SetAdmin(ctx context.Context, in *SetAdminReq, opts ...grpc.CallOption) (*SetAdminResp, error) {
	out := new(SetAdminResp)
	err := c.cc.Invoke(ctx, Usercenter_SetAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) CheckIsAdmin(ctx context.Context, in *CheckIsAdminReq, opts ...grpc.CallOption) (*CheckIsAdminResp, error) {
	out := new(CheckIsAdminResp)
	err := c.cc.Invoke(ctx, Usercenter_CheckIsAdmin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) AddUserContact(ctx context.Context, in *AddUserContactReq, opts ...grpc.CallOption) (*AddUserContactResp, error) {
	out := new(AddUserContactResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/AddUserContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) UpdateUserContact(ctx context.Context, in *UpdateUserContactReq, opts ...grpc.CallOption) (*UpdateUserContactResp, error) {
	out := new(UpdateUserContactResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/UpdateUserContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) DelUserContact(ctx context.Context, in *DelUserContactReq, opts ...grpc.CallOption) (*DelUserContactResp, error) {
	out := new(DelUserContactResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/DelUserContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) GetUserContactById(ctx context.Context, in *GetUserContactByIdReq, opts ...grpc.CallOption) (*GetUserContactByIdResp, error) {
	out := new(GetUserContactByIdResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/GetUserContactById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) SearchUserContact(ctx context.Context, in *SearchUserContactReq, opts ...grpc.CallOption) (*SearchUserContactResp, error) {
	out := new(SearchUserContactResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/SearchUserContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) AddUserAddress(ctx context.Context, in *AddUserAddressReq, opts ...grpc.CallOption) (*AddUserAddressResp, error) {
	out := new(AddUserAddressResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/AddUserAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) UpdateUserAddress(ctx context.Context, in *UpdateUserAddressReq, opts ...grpc.CallOption) (*UpdateUserAddressResp, error) {
	out := new(UpdateUserAddressResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/UpdateUserAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) DelUserAddress(ctx context.Context, in *DelUserAddressReq, opts ...grpc.CallOption) (*DelUserAddressResp, error) {
	out := new(DelUserAddressResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/DelUserAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) GetUserAddressById(ctx context.Context, in *GetUserAddressByIdReq, opts ...grpc.CallOption) (*GetUserAddressByIdResp, error) {
	out := new(GetUserAddressByIdResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/GetUserAddressById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usercenterClient) SearchUserAddress(ctx context.Context, in *SearchUserAddressReq, opts ...grpc.CallOption) (*SearchUserAddressResp, error) {
	out := new(SearchUserAddressResp)
	err := c.cc.Invoke(ctx, "/pb.usercenter/SearchUserAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsercenterServer is the server API for Usercenter service.
// All implementations must embed UnimplementedUsercenterServer
// for forward compatibility
type UsercenterServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error)
	GetUserAuthByAuthKey(context.Context, *GetUserAuthByAuthKeyReq) (*GetUserAuthByAuthKeyResp, error)
	GetUserAuthByUserId(context.Context, *GetUserAuthByUserIdReq) (*GetUserAuthyUserIdResp, error)
	GenerateToken(context.Context, *GenerateTokenReq) (*GenerateTokenResp, error)
	SetAdmin(context.Context, *SetAdminReq) (*SetAdminResp, error)
	CheckIsAdmin(context.Context, *CheckIsAdminReq) (*CheckIsAdminResp, error)
	// -----------------------用户联系方式----------------------
	UpdateUserBaseInfo(context.Context, *UpdateUserBaseInfoReq) (*UpdateUserBaseInfoResp, error)
	// -----------------------用户联系方式----------------------
	AddUserContact(context.Context, *AddUserContactReq) (*AddUserContactResp, error)
	UpdateUserContact(context.Context, *UpdateUserContactReq) (*UpdateUserContactResp, error)
	DelUserContact(context.Context, *DelUserContactReq) (*DelUserContactResp, error)
	GetUserContactById(context.Context, *GetUserContactByIdReq) (*GetUserContactByIdResp, error)
	SearchUserContact(context.Context, *SearchUserContactReq) (*SearchUserContactResp, error)
	// -----------------------用户收货地址表-----------------------
	AddUserAddress(context.Context, *AddUserAddressReq) (*AddUserAddressResp, error)
	UpdateUserAddress(context.Context, *UpdateUserAddressReq) (*UpdateUserAddressResp, error)
	DelUserAddress(context.Context, *DelUserAddressReq) (*DelUserAddressResp, error)
	GetUserAddressById(context.Context, *GetUserAddressByIdReq) (*GetUserAddressByIdResp, error)
	SearchUserAddress(context.Context, *SearchUserAddressReq) (*SearchUserAddressResp, error)
	mustEmbedUnimplementedUsercenterServer()
}

// UnimplementedUsercenterServer must be embedded to have forward compatible implementations.
type UnimplementedUsercenterServer struct {
}

func (UnimplementedUsercenterServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUsercenterServer) Register(context.Context, *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUsercenterServer) GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUsercenterServer) GetUserAuthByAuthKey(context.Context, *GetUserAuthByAuthKeyReq) (*GetUserAuthByAuthKeyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAuthByAuthKey not implemented")
}
func (UnimplementedUsercenterServer) GetUserAuthByUserId(context.Context, *GetUserAuthByUserIdReq) (*GetUserAuthyUserIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAuthByUserId not implemented")
}
func (UnimplementedUsercenterServer) GenerateToken(context.Context, *GenerateTokenReq) (*GenerateTokenResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateToken not implemented")
}
func (UnimplementedUsercenterServer) UpdateUserBaseInfo(context.Context, *UpdateUserBaseInfoReq) (*UpdateUserBaseInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserBaseInfo not implemented")
}
func (UnimplementedUsercenterServer) SetAdmin(context.Context, *SetAdminReq) (*SetAdminResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAdmin not implemented")
}
func (UnimplementedUsercenterServer) CheckIsAdmin(context.Context, *CheckIsAdminReq) (*CheckIsAdminResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIsAdmin not implemented")
}
func (UnimplementedUsercenterServer) AddUserContact(context.Context, *AddUserContactReq) (*AddUserContactResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserContact not implemented")
}
func (UnimplementedUsercenterServer) UpdateUserContact(context.Context, *UpdateUserContactReq) (*UpdateUserContactResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserContact not implemented")
}
func (UnimplementedUsercenterServer) DelUserContact(context.Context, *DelUserContactReq) (*DelUserContactResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelUserContact not implemented")
}
func (UnimplementedUsercenterServer) GetUserContactById(context.Context, *GetUserContactByIdReq) (*GetUserContactByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserContactById not implemented")
}
func (UnimplementedUsercenterServer) SearchUserContact(context.Context, *SearchUserContactReq) (*SearchUserContactResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUserContact not implemented")
}
func (UnimplementedUsercenterServer) AddUserAddress(context.Context, *AddUserAddressReq) (*AddUserAddressResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserAddress not implemented")
}
func (UnimplementedUsercenterServer) UpdateUserAddress(context.Context, *UpdateUserAddressReq) (*UpdateUserAddressResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserAddress not implemented")
}
func (UnimplementedUsercenterServer) DelUserAddress(context.Context, *DelUserAddressReq) (*DelUserAddressResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelUserAddress not implemented")
}
func (UnimplementedUsercenterServer) GetUserAddressById(context.Context, *GetUserAddressByIdReq) (*GetUserAddressByIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAddressById not implemented")
}
func (UnimplementedUsercenterServer) SearchUserAddress(context.Context, *SearchUserAddressReq) (*SearchUserAddressResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUserAddress not implemented")
}
func (UnimplementedUsercenterServer) mustEmbedUnimplementedUsercenterServer() {}

// UnsafeUsercenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsercenterServer will
// result in compilation errors.
type UnsafeUsercenterServer interface {
	mustEmbedUnimplementedUsercenterServer()
}

func RegisterUsercenterServer(s grpc.ServiceRegistrar, srv UsercenterServer) {
	s.RegisterService(&Usercenter_ServiceDesc, srv)
}

func _Usercenter_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/getUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GetUserInfo(ctx, req.(*GetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_GetUserAuthByAuthKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAuthByAuthKeyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GetUserAuthByAuthKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/getUserAuthByAuthKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GetUserAuthByAuthKey(ctx, req.(*GetUserAuthByAuthKeyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_GetUserAuthByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAuthByUserIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GetUserAuthByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/getUserAuthByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GetUserAuthByUserId(ctx, req.(*GetUserAuthByUserIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_GenerateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GenerateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/generateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GenerateToken(ctx, req.(*GenerateTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_SetAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetAdminReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).SetAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_SetAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).SetAdmin(ctx, req.(*SetAdminReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_CheckIsAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckIsAdminReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).CheckIsAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Usercenter_CheckIsAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).CheckIsAdmin(ctx, req.(*CheckIsAdminReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_UpdateUserBaseInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserBaseInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).UpdateUserBaseInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/updateUserBaseInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).UpdateUserBaseInfo(ctx, req.(*UpdateUserBaseInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_AddUserContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserContactReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).AddUserContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/AddUserContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).AddUserContact(ctx, req.(*AddUserContactReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_UpdateUserContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserContactReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).UpdateUserContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/UpdateUserContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).UpdateUserContact(ctx, req.(*UpdateUserContactReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_DelUserContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelUserContactReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).DelUserContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/DelUserContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).DelUserContact(ctx, req.(*DelUserContactReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_GetUserContactById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserContactByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GetUserContactById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/GetUserContactById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GetUserContactById(ctx, req.(*GetUserContactByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_SearchUserContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserContactReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).SearchUserContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/SearchUserContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).SearchUserContact(ctx, req.(*SearchUserContactReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_AddUserAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).AddUserAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/AddUserAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).AddUserAddress(ctx, req.(*AddUserAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_UpdateUserAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).UpdateUserAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/UpdateUserAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).UpdateUserAddress(ctx, req.(*UpdateUserAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_DelUserAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelUserAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).DelUserAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/DelUserAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).DelUserAddress(ctx, req.(*DelUserAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_GetUserAddressById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAddressByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).GetUserAddressById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/GetUserAddressById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).GetUserAddressById(ctx, req.(*GetUserAddressByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Usercenter_SearchUserAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUserAddressReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsercenterServer).SearchUserAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.usercenter/SearchUserAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsercenterServer).SearchUserAddress(ctx, req.(*SearchUserAddressReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Usercenter_ServiceDesc is the grpc.ServiceDesc for Usercenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Usercenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.usercenter",
	HandlerType: (*UsercenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "login",
			Handler:    _Usercenter_Login_Handler,
		},
		{
			MethodName: "register",
			Handler:    _Usercenter_Register_Handler,
		},
		{
			MethodName: "getUserInfo",
			Handler:    _Usercenter_GetUserInfo_Handler,
		},
		{
			MethodName: "getUserAuthByAuthKey",
			Handler:    _Usercenter_GetUserAuthByAuthKey_Handler,
		},
		{
			MethodName: "getUserAuthByUserId",
			Handler:    _Usercenter_GetUserAuthByUserId_Handler,
		},
		{
			MethodName: "generateToken",
			Handler:    _Usercenter_GenerateToken_Handler,
		},
		{
			MethodName: "updateUserBaseInfo",
			Handler:    _Usercenter_UpdateUserBaseInfo_Handler,
		},
		{
			MethodName: "setAdmin",
			Handler:    _Usercenter_SetAdmin_Handler,
		},
		{
			MethodName: "CheckIsAdmin",
			Handler:    _Usercenter_CheckIsAdmin_Handler,
		},
		{
			MethodName: "AddUserContact",
			Handler:    _Usercenter_AddUserContact_Handler,
		},
		{
			MethodName: "UpdateUserContact",
			Handler:    _Usercenter_UpdateUserContact_Handler,
		},
		{
			MethodName: "DelUserContact",
			Handler:    _Usercenter_DelUserContact_Handler,
		},
		{
			MethodName: "GetUserContactById",
			Handler:    _Usercenter_GetUserContactById_Handler,
		},
		{
			MethodName: "SearchUserContact",
			Handler:    _Usercenter_SearchUserContact_Handler,
		},
		{
			MethodName: "AddUserAddress",
			Handler:    _Usercenter_AddUserAddress_Handler,
		},
		{
			MethodName: "UpdateUserAddress",
			Handler:    _Usercenter_UpdateUserAddress_Handler,
		},
		{
			MethodName: "DelUserAddress",
			Handler:    _Usercenter_DelUserAddress_Handler,
		},
		{
			MethodName: "GetUserAddressById",
			Handler:    _Usercenter_GetUserAddressById_Handler,
		},
		{
			MethodName: "SearchUserAddress",
			Handler:    _Usercenter_SearchUserAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "usercenter.proto",
}
