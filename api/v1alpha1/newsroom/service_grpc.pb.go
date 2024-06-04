// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: api/v1alpha1/newsroom/service.proto

package newsroom

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	NewsroomAPI_CreateNewsArticle_FullMethodName        = "/api.v1alpha1.newsroom.NewsroomAPI/CreateNewsArticle"
	NewsroomAPI_ListNewsArticles_FullMethodName         = "/api.v1alpha1.newsroom.NewsroomAPI/ListNewsArticles"
	NewsroomAPI_GetNewsArticleById_FullMethodName       = "/api.v1alpha1.newsroom.NewsroomAPI/GetNewsArticleById"
	NewsroomAPI_UpdateNewsArticle_FullMethodName        = "/api.v1alpha1.newsroom.NewsroomAPI/UpdateNewsArticle"
	NewsroomAPI_CreatePublishedArticle_FullMethodName   = "/api.v1alpha1.newsroom.NewsroomAPI/CreatePublishedArticle"
	NewsroomAPI_ListPublishedArticles_FullMethodName    = "/api.v1alpha1.newsroom.NewsroomAPI/ListPublishedArticles"
	NewsroomAPI_GetPublishedArticleById_FullMethodName  = "/api.v1alpha1.newsroom.NewsroomAPI/GetPublishedArticleById"
	NewsroomAPI_UserActivity_FullMethodName             = "/api.v1alpha1.newsroom.NewsroomAPI/UserActivity"
	NewsroomAPI_GetNewsForUser_FullMethodName           = "/api.v1alpha1.newsroom.NewsroomAPI/GetNewsForUser"
	NewsroomAPI_StoreNewsArticleImage_FullMethodName    = "/api.v1alpha1.newsroom.NewsroomAPI/StoreNewsArticleImage"
	NewsroomAPI_ListImagesForNewsArticle_FullMethodName = "/api.v1alpha1.newsroom.NewsroomAPI/ListImagesForNewsArticle"
)

// NewsroomAPIClient is the client API for NewsroomAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service for interacting with TCN's Newsroom API system.
// Accessing all of the methods Requestuire an authenticated user with the correct
// permissions.
type NewsroomAPIClient interface {
	// create news article
	CreateNewsArticle(ctx context.Context, in *CreateNewsArticleRequest, opts ...grpc.CallOption) (*CreateNewsArticleResponse, error)
	// list news articles
	ListNewsArticles(ctx context.Context, in *ListNewsArticlesRequest, opts ...grpc.CallOption) (*ListNewsArticlesResponse, error)
	// get news article details by the id
	GetNewsArticleById(ctx context.Context, in *GetNewsArticleByIdRequest, opts ...grpc.CallOption) (*GetNewsArticleByIdResponse, error)
	// update news article
	UpdateNewsArticle(ctx context.Context, in *UpdateNewsArticleRequest, opts ...grpc.CallOption) (*UpdateNewsArticleResponse, error)
	// create published article
	CreatePublishedArticle(ctx context.Context, in *CreatePublishedArticleRequest, opts ...grpc.CallOption) (*CreatePublishedArticleResponse, error)
	// list published articles
	ListPublishedArticles(ctx context.Context, in *ListPublishedArticlesRequest, opts ...grpc.CallOption) (*ListPublishedArticlesResponse, error)
	// get published article details by the id
	GetPublishedArticleById(ctx context.Context, in *GetPublishedArticleByIdRequest, opts ...grpc.CallOption) (*GetPublishedArticleByIdResponse, error)
	// user activity updates
	UserActivity(ctx context.Context, in *UserActivityRequest, opts ...grpc.CallOption) (*UserActivityResponse, error)
	// fetch the unseen articles for the user
	GetNewsForUser(ctx context.Context, in *GetNewsForUserRequest, opts ...grpc.CallOption) (*GetNewsForUserResponse, error)
	// upload newsroom image for the news article
	StoreNewsArticleImage(ctx context.Context, in *StoreNewsArticleImageRequest, opts ...grpc.CallOption) (*StoreNewsArticleImageResponse, error)
	// list newsroom images
	ListImagesForNewsArticle(ctx context.Context, in *ListImagesForNewsArticleRequest, opts ...grpc.CallOption) (*ListImagesForNewsArticleResponse, error)
}

type newsroomAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsroomAPIClient(cc grpc.ClientConnInterface) NewsroomAPIClient {
	return &newsroomAPIClient{cc}
}

func (c *newsroomAPIClient) CreateNewsArticle(ctx context.Context, in *CreateNewsArticleRequest, opts ...grpc.CallOption) (*CreateNewsArticleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateNewsArticleResponse)
	err := c.cc.Invoke(ctx, NewsroomAPI_CreateNewsArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsroomAPIClient) ListNewsArticles(ctx context.Context, in *ListNewsArticlesRequest, opts ...grpc.CallOption) (*ListNewsArticlesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNewsArticlesResponse)
	err := c.cc.Invoke(ctx, NewsroomAPI_ListNewsArticles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsroomAPIClient) GetNewsArticleById(ctx context.Context, in *GetNewsArticleByIdRequest, opts ...grpc.CallOption) (*GetNewsArticleByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNewsArticleByIdResponse)
	err := c.cc.Invoke(ctx, NewsroomAPI_GetNewsArticleById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsroomAPIClient) UpdateNewsArticle(ctx context.Context, in *UpdateNewsArticleRequest, opts ...grpc.CallOption) (*UpdateNewsArticleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateNewsArticleResponse)
	err := c.cc.Invoke(ctx, NewsroomAPI_UpdateNewsArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsroomAPIClient) CreatePublishedArticle(ctx context.Context, in *CreatePublishedArticleRequest, opts ...grpc.CallOption) (*CreatePublishedArticleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePublishedArticleResponse)
	err := c.cc.Invoke(ctx, NewsroomAPI_CreatePublishedArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsroomAPIClient) ListPublishedArticles(ctx context.Context, in *ListPublishedArticlesRequest, opts ...grpc.CallOption) (*ListPublishedArticlesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPublishedArticlesResponse)
	err := c.cc.Invoke(ctx, NewsroomAPI_ListPublishedArticles_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsroomAPIClient) GetPublishedArticleById(ctx context.Context, in *GetPublishedArticleByIdRequest, opts ...grpc.CallOption) (*GetPublishedArticleByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPublishedArticleByIdResponse)
	err := c.cc.Invoke(ctx, NewsroomAPI_GetPublishedArticleById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsroomAPIClient) UserActivity(ctx context.Context, in *UserActivityRequest, opts ...grpc.CallOption) (*UserActivityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserActivityResponse)
	err := c.cc.Invoke(ctx, NewsroomAPI_UserActivity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsroomAPIClient) GetNewsForUser(ctx context.Context, in *GetNewsForUserRequest, opts ...grpc.CallOption) (*GetNewsForUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNewsForUserResponse)
	err := c.cc.Invoke(ctx, NewsroomAPI_GetNewsForUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsroomAPIClient) StoreNewsArticleImage(ctx context.Context, in *StoreNewsArticleImageRequest, opts ...grpc.CallOption) (*StoreNewsArticleImageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StoreNewsArticleImageResponse)
	err := c.cc.Invoke(ctx, NewsroomAPI_StoreNewsArticleImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsroomAPIClient) ListImagesForNewsArticle(ctx context.Context, in *ListImagesForNewsArticleRequest, opts ...grpc.CallOption) (*ListImagesForNewsArticleResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListImagesForNewsArticleResponse)
	err := c.cc.Invoke(ctx, NewsroomAPI_ListImagesForNewsArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewsroomAPIServer is the server API for NewsroomAPI service.
// All implementations must embed UnimplementedNewsroomAPIServer
// for forward compatibility
//
// Service for interacting with TCN's Newsroom API system.
// Accessing all of the methods Requestuire an authenticated user with the correct
// permissions.
type NewsroomAPIServer interface {
	// create news article
	CreateNewsArticle(context.Context, *CreateNewsArticleRequest) (*CreateNewsArticleResponse, error)
	// list news articles
	ListNewsArticles(context.Context, *ListNewsArticlesRequest) (*ListNewsArticlesResponse, error)
	// get news article details by the id
	GetNewsArticleById(context.Context, *GetNewsArticleByIdRequest) (*GetNewsArticleByIdResponse, error)
	// update news article
	UpdateNewsArticle(context.Context, *UpdateNewsArticleRequest) (*UpdateNewsArticleResponse, error)
	// create published article
	CreatePublishedArticle(context.Context, *CreatePublishedArticleRequest) (*CreatePublishedArticleResponse, error)
	// list published articles
	ListPublishedArticles(context.Context, *ListPublishedArticlesRequest) (*ListPublishedArticlesResponse, error)
	// get published article details by the id
	GetPublishedArticleById(context.Context, *GetPublishedArticleByIdRequest) (*GetPublishedArticleByIdResponse, error)
	// user activity updates
	UserActivity(context.Context, *UserActivityRequest) (*UserActivityResponse, error)
	// fetch the unseen articles for the user
	GetNewsForUser(context.Context, *GetNewsForUserRequest) (*GetNewsForUserResponse, error)
	// upload newsroom image for the news article
	StoreNewsArticleImage(context.Context, *StoreNewsArticleImageRequest) (*StoreNewsArticleImageResponse, error)
	// list newsroom images
	ListImagesForNewsArticle(context.Context, *ListImagesForNewsArticleRequest) (*ListImagesForNewsArticleResponse, error)
	mustEmbedUnimplementedNewsroomAPIServer()
}

// UnimplementedNewsroomAPIServer must be embedded to have forward compatible implementations.
type UnimplementedNewsroomAPIServer struct {
}

func (UnimplementedNewsroomAPIServer) CreateNewsArticle(context.Context, *CreateNewsArticleRequest) (*CreateNewsArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewsArticle not implemented")
}
func (UnimplementedNewsroomAPIServer) ListNewsArticles(context.Context, *ListNewsArticlesRequest) (*ListNewsArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNewsArticles not implemented")
}
func (UnimplementedNewsroomAPIServer) GetNewsArticleById(context.Context, *GetNewsArticleByIdRequest) (*GetNewsArticleByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewsArticleById not implemented")
}
func (UnimplementedNewsroomAPIServer) UpdateNewsArticle(context.Context, *UpdateNewsArticleRequest) (*UpdateNewsArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNewsArticle not implemented")
}
func (UnimplementedNewsroomAPIServer) CreatePublishedArticle(context.Context, *CreatePublishedArticleRequest) (*CreatePublishedArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePublishedArticle not implemented")
}
func (UnimplementedNewsroomAPIServer) ListPublishedArticles(context.Context, *ListPublishedArticlesRequest) (*ListPublishedArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPublishedArticles not implemented")
}
func (UnimplementedNewsroomAPIServer) GetPublishedArticleById(context.Context, *GetPublishedArticleByIdRequest) (*GetPublishedArticleByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublishedArticleById not implemented")
}
func (UnimplementedNewsroomAPIServer) UserActivity(context.Context, *UserActivityRequest) (*UserActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserActivity not implemented")
}
func (UnimplementedNewsroomAPIServer) GetNewsForUser(context.Context, *GetNewsForUserRequest) (*GetNewsForUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewsForUser not implemented")
}
func (UnimplementedNewsroomAPIServer) StoreNewsArticleImage(context.Context, *StoreNewsArticleImageRequest) (*StoreNewsArticleImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreNewsArticleImage not implemented")
}
func (UnimplementedNewsroomAPIServer) ListImagesForNewsArticle(context.Context, *ListImagesForNewsArticleRequest) (*ListImagesForNewsArticleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListImagesForNewsArticle not implemented")
}
func (UnimplementedNewsroomAPIServer) mustEmbedUnimplementedNewsroomAPIServer() {}

// UnsafeNewsroomAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewsroomAPIServer will
// result in compilation errors.
type UnsafeNewsroomAPIServer interface {
	mustEmbedUnimplementedNewsroomAPIServer()
}

func RegisterNewsroomAPIServer(s grpc.ServiceRegistrar, srv NewsroomAPIServer) {
	s.RegisterService(&NewsroomAPI_ServiceDesc, srv)
}

func _NewsroomAPI_CreateNewsArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewsArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsroomAPIServer).CreateNewsArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsroomAPI_CreateNewsArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsroomAPIServer).CreateNewsArticle(ctx, req.(*CreateNewsArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsroomAPI_ListNewsArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNewsArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsroomAPIServer).ListNewsArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsroomAPI_ListNewsArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsroomAPIServer).ListNewsArticles(ctx, req.(*ListNewsArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsroomAPI_GetNewsArticleById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewsArticleByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsroomAPIServer).GetNewsArticleById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsroomAPI_GetNewsArticleById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsroomAPIServer).GetNewsArticleById(ctx, req.(*GetNewsArticleByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsroomAPI_UpdateNewsArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNewsArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsroomAPIServer).UpdateNewsArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsroomAPI_UpdateNewsArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsroomAPIServer).UpdateNewsArticle(ctx, req.(*UpdateNewsArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsroomAPI_CreatePublishedArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePublishedArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsroomAPIServer).CreatePublishedArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsroomAPI_CreatePublishedArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsroomAPIServer).CreatePublishedArticle(ctx, req.(*CreatePublishedArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsroomAPI_ListPublishedArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPublishedArticlesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsroomAPIServer).ListPublishedArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsroomAPI_ListPublishedArticles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsroomAPIServer).ListPublishedArticles(ctx, req.(*ListPublishedArticlesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsroomAPI_GetPublishedArticleById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublishedArticleByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsroomAPIServer).GetPublishedArticleById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsroomAPI_GetPublishedArticleById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsroomAPIServer).GetPublishedArticleById(ctx, req.(*GetPublishedArticleByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsroomAPI_UserActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsroomAPIServer).UserActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsroomAPI_UserActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsroomAPIServer).UserActivity(ctx, req.(*UserActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsroomAPI_GetNewsForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewsForUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsroomAPIServer).GetNewsForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsroomAPI_GetNewsForUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsroomAPIServer).GetNewsForUser(ctx, req.(*GetNewsForUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsroomAPI_StoreNewsArticleImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreNewsArticleImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsroomAPIServer).StoreNewsArticleImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsroomAPI_StoreNewsArticleImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsroomAPIServer).StoreNewsArticleImage(ctx, req.(*StoreNewsArticleImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsroomAPI_ListImagesForNewsArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImagesForNewsArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsroomAPIServer).ListImagesForNewsArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsroomAPI_ListImagesForNewsArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsroomAPIServer).ListImagesForNewsArticle(ctx, req.(*ListImagesForNewsArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NewsroomAPI_ServiceDesc is the grpc.ServiceDesc for NewsroomAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NewsroomAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1alpha1.newsroom.NewsroomAPI",
	HandlerType: (*NewsroomAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewsArticle",
			Handler:    _NewsroomAPI_CreateNewsArticle_Handler,
		},
		{
			MethodName: "ListNewsArticles",
			Handler:    _NewsroomAPI_ListNewsArticles_Handler,
		},
		{
			MethodName: "GetNewsArticleById",
			Handler:    _NewsroomAPI_GetNewsArticleById_Handler,
		},
		{
			MethodName: "UpdateNewsArticle",
			Handler:    _NewsroomAPI_UpdateNewsArticle_Handler,
		},
		{
			MethodName: "CreatePublishedArticle",
			Handler:    _NewsroomAPI_CreatePublishedArticle_Handler,
		},
		{
			MethodName: "ListPublishedArticles",
			Handler:    _NewsroomAPI_ListPublishedArticles_Handler,
		},
		{
			MethodName: "GetPublishedArticleById",
			Handler:    _NewsroomAPI_GetPublishedArticleById_Handler,
		},
		{
			MethodName: "UserActivity",
			Handler:    _NewsroomAPI_UserActivity_Handler,
		},
		{
			MethodName: "GetNewsForUser",
			Handler:    _NewsroomAPI_GetNewsForUser_Handler,
		},
		{
			MethodName: "StoreNewsArticleImage",
			Handler:    _NewsroomAPI_StoreNewsArticleImage_Handler,
		},
		{
			MethodName: "ListImagesForNewsArticle",
			Handler:    _NewsroomAPI_ListImagesForNewsArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1alpha1/newsroom/service.proto",
}
