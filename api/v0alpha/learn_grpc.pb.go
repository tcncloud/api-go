// Copyright (c) 2019, TCN Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/v0alpha/learn.proto

package v0alpha

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
	Learn_Exist_FullMethodName                      = "/api.v0alpha.Learn/Exist"
	Learn_Content_FullMethodName                    = "/api.v0alpha.Learn/Content"
	Learn_ExportMany_FullMethodName                 = "/api.v0alpha.Learn/ExportMany"
	Learn_SearchContent_FullMethodName              = "/api.v0alpha.Learn/SearchContent"
	Learn_ListSearchResults_FullMethodName          = "/api.v0alpha.Learn/ListSearchResults"
	Learn_Standalone_FullMethodName                 = "/api.v0alpha.Learn/Standalone"
	Learn_ContentEditorData_FullMethodName          = "/api.v0alpha.Learn/ContentEditorData"
	Learn_Update_FullMethodName                     = "/api.v0alpha.Learn/Update"
	Learn_StoreStaticImage_FullMethodName           = "/api.v0alpha.Learn/StoreStaticImage"
	Learn_UploadDynamicScreenshot_FullMethodName    = "/api.v0alpha.Learn/UploadDynamicScreenshot"
	Learn_DeleteStandalone_FullMethodName           = "/api.v0alpha.Learn/DeleteStandalone"
	Learn_Snippet_FullMethodName                    = "/api.v0alpha.Learn/Snippet"
	Learn_DeleteLearnPages_FullMethodName           = "/api.v0alpha.Learn/DeleteLearnPages"
	Learn_CreateEditVersion_FullMethodName          = "/api.v0alpha.Learn/CreateEditVersion"
	Learn_PublishVersion_FullMethodName             = "/api.v0alpha.Learn/PublishVersion"
	Learn_ContentByVersion_FullMethodName           = "/api.v0alpha.Learn/ContentByVersion"
	Learn_UpdateByVersion_FullMethodName            = "/api.v0alpha.Learn/UpdateByVersion"
	Learn_ListSearchResultsByVersion_FullMethodName = "/api.v0alpha.Learn/ListSearchResultsByVersion"
	Learn_ReviewFileVersions_FullMethodName         = "/api.v0alpha.Learn/ReviewFileVersions"
	Learn_ReviewVersion_FullMethodName              = "/api.v0alpha.Learn/ReviewVersion"
	Learn_ExportManyStream_FullMethodName           = "/api.v0alpha.Learn/ExportManyStream"
)

// LearnClient is the client API for Learn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LearnClient interface {
	// check if learning page already exists
	Exist(ctx context.Context, in *ExistReq, opts ...grpc.CallOption) (*ExistRes, error)
	// retreive content from learning pages
	Content(ctx context.Context, in *ContentReq, opts ...grpc.CallOption) (*ContentRes, error)
	// exports multiple pages of the learning center markdown as PDF
	ExportMany(ctx context.Context, in *ExportManyReq, opts ...grpc.CallOption) (*ExportRes, error)
	// search content in learning pages
	// we allow all the logged in agents/admins to view search content
	SearchContent(ctx context.Context, in *SearchContentReq, opts ...grpc.CallOption) (*SearchRes, error)
	// stream search content results in learning pages
	// we allow all the logged in agents/admins to view search content
	ListSearchResults(ctx context.Context, in *SearchContentReq, opts ...grpc.CallOption) (Learn_ListSearchResultsClient, error)
	// get standalone articles from learning pages
	// we allow all the logged in agents/admins to view standalone articles
	Standalone(ctx context.Context, in *StandaloneReq, opts ...grpc.CallOption) (*StandaloneRes, error)
	// retrieve user who edited the content last
	ContentEditorData(ctx context.Context, in *ContentEditorDataReq, opts ...grpc.CallOption) (*ContentEditorDataRes, error)
	// update contents for learning pages
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateRes, error)
	// upload url for static images
	StoreStaticImage(ctx context.Context, in *StoreStaticImageReq, opts ...grpc.CallOption) (*StoreStaticImageRes, error)
	// upload dynamic learning image screenshot
	UploadDynamicScreenshot(ctx context.Context, in *UploadDynamicScreenshotReq, opts ...grpc.CallOption) (*UploadDynamicScreenshotRes, error)
	// delete standalone articles from learning pages
	DeleteStandalone(ctx context.Context, in *DeleteStandaloneReq, opts ...grpc.CallOption) (*DeleteStandaloneRes, error)
	// get snippet content from learning pages
	// we allow all the logged in agents/admins to view snippet content
	Snippet(ctx context.Context, in *SnippetReq, opts ...grpc.CallOption) (*SnippetRes, error)
	// delete learning pages
	DeleteLearnPages(ctx context.Context, in *DeleteLearnPagesReq, opts ...grpc.CallOption) (*DeleteLearnPagesRes, error)
	// create edit version
	CreateEditVersion(ctx context.Context, in *CreateEditVersionReq, opts ...grpc.CallOption) (*CreateEditVersionRes, error)
	// publish version
	PublishVersion(ctx context.Context, in *PublishVersionReq, opts ...grpc.CallOption) (*PublishVersionRes, error)
	// retrieve content from learning pages based on version
	ContentByVersion(ctx context.Context, in *ContentByVersionReq, opts ...grpc.CallOption) (*ContentRes, error)
	// update contents for learning pages by version
	UpdateByVersion(ctx context.Context, in *UpdateByVersionReq, opts ...grpc.CallOption) (*UpdateRes, error)
	// stream search content results in learning pages by version
	// we allow all the logged in agents/admins to view search content
	ListSearchResultsByVersion(ctx context.Context, in *SearchContentByVersionReq, opts ...grpc.CallOption) (Learn_ListSearchResultsByVersionClient, error)
	// return diff by comparing file contens from any version
	ReviewFileVersions(ctx context.Context, in *ReviewFileVersionsReq, opts ...grpc.CallOption) (*ReviewFileVersionsRes, error)
	// returns list of file details after comparing different versions
	ReviewVersion(ctx context.Context, in *ReviewVersionReq, opts ...grpc.CallOption) (*ReviewVersionRes, error)
	// exports multiple pages of the learning center markdown as PDF file stream
	ExportManyStream(ctx context.Context, in *ExportManyReq, opts ...grpc.CallOption) (Learn_ExportManyStreamClient, error)
}

type learnClient struct {
	cc grpc.ClientConnInterface
}

func NewLearnClient(cc grpc.ClientConnInterface) LearnClient {
	return &learnClient{cc}
}

func (c *learnClient) Exist(ctx context.Context, in *ExistReq, opts ...grpc.CallOption) (*ExistRes, error) {
	out := new(ExistRes)
	err := c.cc.Invoke(ctx, Learn_Exist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnClient) Content(ctx context.Context, in *ContentReq, opts ...grpc.CallOption) (*ContentRes, error) {
	out := new(ContentRes)
	err := c.cc.Invoke(ctx, Learn_Content_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnClient) ExportMany(ctx context.Context, in *ExportManyReq, opts ...grpc.CallOption) (*ExportRes, error) {
	out := new(ExportRes)
	err := c.cc.Invoke(ctx, Learn_ExportMany_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnClient) SearchContent(ctx context.Context, in *SearchContentReq, opts ...grpc.CallOption) (*SearchRes, error) {
	out := new(SearchRes)
	err := c.cc.Invoke(ctx, Learn_SearchContent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnClient) ListSearchResults(ctx context.Context, in *SearchContentReq, opts ...grpc.CallOption) (Learn_ListSearchResultsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Learn_ServiceDesc.Streams[0], Learn_ListSearchResults_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &learnListSearchResultsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Learn_ListSearchResultsClient interface {
	Recv() (*SearchRes, error)
	grpc.ClientStream
}

type learnListSearchResultsClient struct {
	grpc.ClientStream
}

func (x *learnListSearchResultsClient) Recv() (*SearchRes, error) {
	m := new(SearchRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *learnClient) Standalone(ctx context.Context, in *StandaloneReq, opts ...grpc.CallOption) (*StandaloneRes, error) {
	out := new(StandaloneRes)
	err := c.cc.Invoke(ctx, Learn_Standalone_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnClient) ContentEditorData(ctx context.Context, in *ContentEditorDataReq, opts ...grpc.CallOption) (*ContentEditorDataRes, error) {
	out := new(ContentEditorDataRes)
	err := c.cc.Invoke(ctx, Learn_ContentEditorData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateRes, error) {
	out := new(UpdateRes)
	err := c.cc.Invoke(ctx, Learn_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnClient) StoreStaticImage(ctx context.Context, in *StoreStaticImageReq, opts ...grpc.CallOption) (*StoreStaticImageRes, error) {
	out := new(StoreStaticImageRes)
	err := c.cc.Invoke(ctx, Learn_StoreStaticImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnClient) UploadDynamicScreenshot(ctx context.Context, in *UploadDynamicScreenshotReq, opts ...grpc.CallOption) (*UploadDynamicScreenshotRes, error) {
	out := new(UploadDynamicScreenshotRes)
	err := c.cc.Invoke(ctx, Learn_UploadDynamicScreenshot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnClient) DeleteStandalone(ctx context.Context, in *DeleteStandaloneReq, opts ...grpc.CallOption) (*DeleteStandaloneRes, error) {
	out := new(DeleteStandaloneRes)
	err := c.cc.Invoke(ctx, Learn_DeleteStandalone_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnClient) Snippet(ctx context.Context, in *SnippetReq, opts ...grpc.CallOption) (*SnippetRes, error) {
	out := new(SnippetRes)
	err := c.cc.Invoke(ctx, Learn_Snippet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnClient) DeleteLearnPages(ctx context.Context, in *DeleteLearnPagesReq, opts ...grpc.CallOption) (*DeleteLearnPagesRes, error) {
	out := new(DeleteLearnPagesRes)
	err := c.cc.Invoke(ctx, Learn_DeleteLearnPages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnClient) CreateEditVersion(ctx context.Context, in *CreateEditVersionReq, opts ...grpc.CallOption) (*CreateEditVersionRes, error) {
	out := new(CreateEditVersionRes)
	err := c.cc.Invoke(ctx, Learn_CreateEditVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnClient) PublishVersion(ctx context.Context, in *PublishVersionReq, opts ...grpc.CallOption) (*PublishVersionRes, error) {
	out := new(PublishVersionRes)
	err := c.cc.Invoke(ctx, Learn_PublishVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnClient) ContentByVersion(ctx context.Context, in *ContentByVersionReq, opts ...grpc.CallOption) (*ContentRes, error) {
	out := new(ContentRes)
	err := c.cc.Invoke(ctx, Learn_ContentByVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnClient) UpdateByVersion(ctx context.Context, in *UpdateByVersionReq, opts ...grpc.CallOption) (*UpdateRes, error) {
	out := new(UpdateRes)
	err := c.cc.Invoke(ctx, Learn_UpdateByVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnClient) ListSearchResultsByVersion(ctx context.Context, in *SearchContentByVersionReq, opts ...grpc.CallOption) (Learn_ListSearchResultsByVersionClient, error) {
	stream, err := c.cc.NewStream(ctx, &Learn_ServiceDesc.Streams[1], Learn_ListSearchResultsByVersion_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &learnListSearchResultsByVersionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Learn_ListSearchResultsByVersionClient interface {
	Recv() (*SearchRes, error)
	grpc.ClientStream
}

type learnListSearchResultsByVersionClient struct {
	grpc.ClientStream
}

func (x *learnListSearchResultsByVersionClient) Recv() (*SearchRes, error) {
	m := new(SearchRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *learnClient) ReviewFileVersions(ctx context.Context, in *ReviewFileVersionsReq, opts ...grpc.CallOption) (*ReviewFileVersionsRes, error) {
	out := new(ReviewFileVersionsRes)
	err := c.cc.Invoke(ctx, Learn_ReviewFileVersions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnClient) ReviewVersion(ctx context.Context, in *ReviewVersionReq, opts ...grpc.CallOption) (*ReviewVersionRes, error) {
	out := new(ReviewVersionRes)
	err := c.cc.Invoke(ctx, Learn_ReviewVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learnClient) ExportManyStream(ctx context.Context, in *ExportManyReq, opts ...grpc.CallOption) (Learn_ExportManyStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Learn_ServiceDesc.Streams[2], Learn_ExportManyStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &learnExportManyStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Learn_ExportManyStreamClient interface {
	Recv() (*ExportRes, error)
	grpc.ClientStream
}

type learnExportManyStreamClient struct {
	grpc.ClientStream
}

func (x *learnExportManyStreamClient) Recv() (*ExportRes, error) {
	m := new(ExportRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LearnServer is the server API for Learn service.
// All implementations must embed UnimplementedLearnServer
// for forward compatibility
type LearnServer interface {
	// check if learning page already exists
	Exist(context.Context, *ExistReq) (*ExistRes, error)
	// retreive content from learning pages
	Content(context.Context, *ContentReq) (*ContentRes, error)
	// exports multiple pages of the learning center markdown as PDF
	ExportMany(context.Context, *ExportManyReq) (*ExportRes, error)
	// search content in learning pages
	// we allow all the logged in agents/admins to view search content
	SearchContent(context.Context, *SearchContentReq) (*SearchRes, error)
	// stream search content results in learning pages
	// we allow all the logged in agents/admins to view search content
	ListSearchResults(*SearchContentReq, Learn_ListSearchResultsServer) error
	// get standalone articles from learning pages
	// we allow all the logged in agents/admins to view standalone articles
	Standalone(context.Context, *StandaloneReq) (*StandaloneRes, error)
	// retrieve user who edited the content last
	ContentEditorData(context.Context, *ContentEditorDataReq) (*ContentEditorDataRes, error)
	// update contents for learning pages
	Update(context.Context, *UpdateReq) (*UpdateRes, error)
	// upload url for static images
	StoreStaticImage(context.Context, *StoreStaticImageReq) (*StoreStaticImageRes, error)
	// upload dynamic learning image screenshot
	UploadDynamicScreenshot(context.Context, *UploadDynamicScreenshotReq) (*UploadDynamicScreenshotRes, error)
	// delete standalone articles from learning pages
	DeleteStandalone(context.Context, *DeleteStandaloneReq) (*DeleteStandaloneRes, error)
	// get snippet content from learning pages
	// we allow all the logged in agents/admins to view snippet content
	Snippet(context.Context, *SnippetReq) (*SnippetRes, error)
	// delete learning pages
	DeleteLearnPages(context.Context, *DeleteLearnPagesReq) (*DeleteLearnPagesRes, error)
	// create edit version
	CreateEditVersion(context.Context, *CreateEditVersionReq) (*CreateEditVersionRes, error)
	// publish version
	PublishVersion(context.Context, *PublishVersionReq) (*PublishVersionRes, error)
	// retrieve content from learning pages based on version
	ContentByVersion(context.Context, *ContentByVersionReq) (*ContentRes, error)
	// update contents for learning pages by version
	UpdateByVersion(context.Context, *UpdateByVersionReq) (*UpdateRes, error)
	// stream search content results in learning pages by version
	// we allow all the logged in agents/admins to view search content
	ListSearchResultsByVersion(*SearchContentByVersionReq, Learn_ListSearchResultsByVersionServer) error
	// return diff by comparing file contens from any version
	ReviewFileVersions(context.Context, *ReviewFileVersionsReq) (*ReviewFileVersionsRes, error)
	// returns list of file details after comparing different versions
	ReviewVersion(context.Context, *ReviewVersionReq) (*ReviewVersionRes, error)
	// exports multiple pages of the learning center markdown as PDF file stream
	ExportManyStream(*ExportManyReq, Learn_ExportManyStreamServer) error
	mustEmbedUnimplementedLearnServer()
}

// UnimplementedLearnServer must be embedded to have forward compatible implementations.
type UnimplementedLearnServer struct {
}

func (UnimplementedLearnServer) Exist(context.Context, *ExistReq) (*ExistRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exist not implemented")
}
func (UnimplementedLearnServer) Content(context.Context, *ContentReq) (*ContentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Content not implemented")
}
func (UnimplementedLearnServer) ExportMany(context.Context, *ExportManyReq) (*ExportRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportMany not implemented")
}
func (UnimplementedLearnServer) SearchContent(context.Context, *SearchContentReq) (*SearchRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchContent not implemented")
}
func (UnimplementedLearnServer) ListSearchResults(*SearchContentReq, Learn_ListSearchResultsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListSearchResults not implemented")
}
func (UnimplementedLearnServer) Standalone(context.Context, *StandaloneReq) (*StandaloneRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Standalone not implemented")
}
func (UnimplementedLearnServer) ContentEditorData(context.Context, *ContentEditorDataReq) (*ContentEditorDataRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContentEditorData not implemented")
}
func (UnimplementedLearnServer) Update(context.Context, *UpdateReq) (*UpdateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedLearnServer) StoreStaticImage(context.Context, *StoreStaticImageReq) (*StoreStaticImageRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreStaticImage not implemented")
}
func (UnimplementedLearnServer) UploadDynamicScreenshot(context.Context, *UploadDynamicScreenshotReq) (*UploadDynamicScreenshotRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadDynamicScreenshot not implemented")
}
func (UnimplementedLearnServer) DeleteStandalone(context.Context, *DeleteStandaloneReq) (*DeleteStandaloneRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStandalone not implemented")
}
func (UnimplementedLearnServer) Snippet(context.Context, *SnippetReq) (*SnippetRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Snippet not implemented")
}
func (UnimplementedLearnServer) DeleteLearnPages(context.Context, *DeleteLearnPagesReq) (*DeleteLearnPagesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLearnPages not implemented")
}
func (UnimplementedLearnServer) CreateEditVersion(context.Context, *CreateEditVersionReq) (*CreateEditVersionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEditVersion not implemented")
}
func (UnimplementedLearnServer) PublishVersion(context.Context, *PublishVersionReq) (*PublishVersionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishVersion not implemented")
}
func (UnimplementedLearnServer) ContentByVersion(context.Context, *ContentByVersionReq) (*ContentRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ContentByVersion not implemented")
}
func (UnimplementedLearnServer) UpdateByVersion(context.Context, *UpdateByVersionReq) (*UpdateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateByVersion not implemented")
}
func (UnimplementedLearnServer) ListSearchResultsByVersion(*SearchContentByVersionReq, Learn_ListSearchResultsByVersionServer) error {
	return status.Errorf(codes.Unimplemented, "method ListSearchResultsByVersion not implemented")
}
func (UnimplementedLearnServer) ReviewFileVersions(context.Context, *ReviewFileVersionsReq) (*ReviewFileVersionsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReviewFileVersions not implemented")
}
func (UnimplementedLearnServer) ReviewVersion(context.Context, *ReviewVersionReq) (*ReviewVersionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReviewVersion not implemented")
}
func (UnimplementedLearnServer) ExportManyStream(*ExportManyReq, Learn_ExportManyStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ExportManyStream not implemented")
}
func (UnimplementedLearnServer) mustEmbedUnimplementedLearnServer() {}

// UnsafeLearnServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LearnServer will
// result in compilation errors.
type UnsafeLearnServer interface {
	mustEmbedUnimplementedLearnServer()
}

func RegisterLearnServer(s grpc.ServiceRegistrar, srv LearnServer) {
	s.RegisterService(&Learn_ServiceDesc, srv)
}

func _Learn_Exist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExistReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnServer).Exist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Learn_Exist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnServer).Exist(ctx, req.(*ExistReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Learn_Content_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnServer).Content(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Learn_Content_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnServer).Content(ctx, req.(*ContentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Learn_ExportMany_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportManyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnServer).ExportMany(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Learn_ExportMany_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnServer).ExportMany(ctx, req.(*ExportManyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Learn_SearchContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchContentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnServer).SearchContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Learn_SearchContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnServer).SearchContent(ctx, req.(*SearchContentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Learn_ListSearchResults_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchContentReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LearnServer).ListSearchResults(m, &learnListSearchResultsServer{stream})
}

type Learn_ListSearchResultsServer interface {
	Send(*SearchRes) error
	grpc.ServerStream
}

type learnListSearchResultsServer struct {
	grpc.ServerStream
}

func (x *learnListSearchResultsServer) Send(m *SearchRes) error {
	return x.ServerStream.SendMsg(m)
}

func _Learn_Standalone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StandaloneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnServer).Standalone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Learn_Standalone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnServer).Standalone(ctx, req.(*StandaloneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Learn_ContentEditorData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentEditorDataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnServer).ContentEditorData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Learn_ContentEditorData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnServer).ContentEditorData(ctx, req.(*ContentEditorDataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Learn_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Learn_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Learn_StoreStaticImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreStaticImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnServer).StoreStaticImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Learn_StoreStaticImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnServer).StoreStaticImage(ctx, req.(*StoreStaticImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Learn_UploadDynamicScreenshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadDynamicScreenshotReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnServer).UploadDynamicScreenshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Learn_UploadDynamicScreenshot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnServer).UploadDynamicScreenshot(ctx, req.(*UploadDynamicScreenshotReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Learn_DeleteStandalone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStandaloneReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnServer).DeleteStandalone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Learn_DeleteStandalone_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnServer).DeleteStandalone(ctx, req.(*DeleteStandaloneReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Learn_Snippet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SnippetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnServer).Snippet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Learn_Snippet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnServer).Snippet(ctx, req.(*SnippetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Learn_DeleteLearnPages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLearnPagesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnServer).DeleteLearnPages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Learn_DeleteLearnPages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnServer).DeleteLearnPages(ctx, req.(*DeleteLearnPagesReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Learn_CreateEditVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEditVersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnServer).CreateEditVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Learn_CreateEditVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnServer).CreateEditVersion(ctx, req.(*CreateEditVersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Learn_PublishVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishVersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnServer).PublishVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Learn_PublishVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnServer).PublishVersion(ctx, req.(*PublishVersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Learn_ContentByVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContentByVersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnServer).ContentByVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Learn_ContentByVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnServer).ContentByVersion(ctx, req.(*ContentByVersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Learn_UpdateByVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateByVersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnServer).UpdateByVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Learn_UpdateByVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnServer).UpdateByVersion(ctx, req.(*UpdateByVersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Learn_ListSearchResultsByVersion_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchContentByVersionReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LearnServer).ListSearchResultsByVersion(m, &learnListSearchResultsByVersionServer{stream})
}

type Learn_ListSearchResultsByVersionServer interface {
	Send(*SearchRes) error
	grpc.ServerStream
}

type learnListSearchResultsByVersionServer struct {
	grpc.ServerStream
}

func (x *learnListSearchResultsByVersionServer) Send(m *SearchRes) error {
	return x.ServerStream.SendMsg(m)
}

func _Learn_ReviewFileVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewFileVersionsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnServer).ReviewFileVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Learn_ReviewFileVersions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnServer).ReviewFileVersions(ctx, req.(*ReviewFileVersionsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Learn_ReviewVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReviewVersionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearnServer).ReviewVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Learn_ReviewVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearnServer).ReviewVersion(ctx, req.(*ReviewVersionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Learn_ExportManyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExportManyReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LearnServer).ExportManyStream(m, &learnExportManyStreamServer{stream})
}

type Learn_ExportManyStreamServer interface {
	Send(*ExportRes) error
	grpc.ServerStream
}

type learnExportManyStreamServer struct {
	grpc.ServerStream
}

func (x *learnExportManyStreamServer) Send(m *ExportRes) error {
	return x.ServerStream.SendMsg(m)
}

// Learn_ServiceDesc is the grpc.ServiceDesc for Learn service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Learn_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v0alpha.Learn",
	HandlerType: (*LearnServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Exist",
			Handler:    _Learn_Exist_Handler,
		},
		{
			MethodName: "Content",
			Handler:    _Learn_Content_Handler,
		},
		{
			MethodName: "ExportMany",
			Handler:    _Learn_ExportMany_Handler,
		},
		{
			MethodName: "SearchContent",
			Handler:    _Learn_SearchContent_Handler,
		},
		{
			MethodName: "Standalone",
			Handler:    _Learn_Standalone_Handler,
		},
		{
			MethodName: "ContentEditorData",
			Handler:    _Learn_ContentEditorData_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Learn_Update_Handler,
		},
		{
			MethodName: "StoreStaticImage",
			Handler:    _Learn_StoreStaticImage_Handler,
		},
		{
			MethodName: "UploadDynamicScreenshot",
			Handler:    _Learn_UploadDynamicScreenshot_Handler,
		},
		{
			MethodName: "DeleteStandalone",
			Handler:    _Learn_DeleteStandalone_Handler,
		},
		{
			MethodName: "Snippet",
			Handler:    _Learn_Snippet_Handler,
		},
		{
			MethodName: "DeleteLearnPages",
			Handler:    _Learn_DeleteLearnPages_Handler,
		},
		{
			MethodName: "CreateEditVersion",
			Handler:    _Learn_CreateEditVersion_Handler,
		},
		{
			MethodName: "PublishVersion",
			Handler:    _Learn_PublishVersion_Handler,
		},
		{
			MethodName: "ContentByVersion",
			Handler:    _Learn_ContentByVersion_Handler,
		},
		{
			MethodName: "UpdateByVersion",
			Handler:    _Learn_UpdateByVersion_Handler,
		},
		{
			MethodName: "ReviewFileVersions",
			Handler:    _Learn_ReviewFileVersions_Handler,
		},
		{
			MethodName: "ReviewVersion",
			Handler:    _Learn_ReviewVersion_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListSearchResults",
			Handler:       _Learn_ListSearchResults_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListSearchResultsByVersion",
			Handler:       _Learn_ListSearchResultsByVersion_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ExportManyStream",
			Handler:       _Learn_ExportManyStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/v0alpha/learn.proto",
}
