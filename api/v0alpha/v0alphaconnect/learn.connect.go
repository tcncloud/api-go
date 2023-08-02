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

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v0alpha/learn.proto

package v0alphaconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v0alpha "github.com/tcncloud/api-go/api/v0alpha"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// LearnName is the fully-qualified name of the Learn service.
	LearnName = "api.v0alpha.Learn"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// LearnExistProcedure is the fully-qualified name of the Learn's Exist RPC.
	LearnExistProcedure = "/api.v0alpha.Learn/Exist"
	// LearnContentProcedure is the fully-qualified name of the Learn's Content RPC.
	LearnContentProcedure = "/api.v0alpha.Learn/Content"
	// LearnExportManyProcedure is the fully-qualified name of the Learn's ExportMany RPC.
	LearnExportManyProcedure = "/api.v0alpha.Learn/ExportMany"
	// LearnSearchContentProcedure is the fully-qualified name of the Learn's SearchContent RPC.
	LearnSearchContentProcedure = "/api.v0alpha.Learn/SearchContent"
	// LearnListSearchResultsProcedure is the fully-qualified name of the Learn's ListSearchResults RPC.
	LearnListSearchResultsProcedure = "/api.v0alpha.Learn/ListSearchResults"
	// LearnStandaloneProcedure is the fully-qualified name of the Learn's Standalone RPC.
	LearnStandaloneProcedure = "/api.v0alpha.Learn/Standalone"
	// LearnContentEditorDataProcedure is the fully-qualified name of the Learn's ContentEditorData RPC.
	LearnContentEditorDataProcedure = "/api.v0alpha.Learn/ContentEditorData"
	// LearnUpdateProcedure is the fully-qualified name of the Learn's Update RPC.
	LearnUpdateProcedure = "/api.v0alpha.Learn/Update"
	// LearnStoreStaticImageProcedure is the fully-qualified name of the Learn's StoreStaticImage RPC.
	LearnStoreStaticImageProcedure = "/api.v0alpha.Learn/StoreStaticImage"
	// LearnUploadDynamicScreenshotProcedure is the fully-qualified name of the Learn's
	// UploadDynamicScreenshot RPC.
	LearnUploadDynamicScreenshotProcedure = "/api.v0alpha.Learn/UploadDynamicScreenshot"
	// LearnDeleteStandaloneProcedure is the fully-qualified name of the Learn's DeleteStandalone RPC.
	LearnDeleteStandaloneProcedure = "/api.v0alpha.Learn/DeleteStandalone"
	// LearnSnippetProcedure is the fully-qualified name of the Learn's Snippet RPC.
	LearnSnippetProcedure = "/api.v0alpha.Learn/Snippet"
	// LearnDeleteLearnPagesProcedure is the fully-qualified name of the Learn's DeleteLearnPages RPC.
	LearnDeleteLearnPagesProcedure = "/api.v0alpha.Learn/DeleteLearnPages"
)

// LearnClient is a client for the api.v0alpha.Learn service.
type LearnClient interface {
	// check if learning page already exists
	Exist(context.Context, *connect_go.Request[v0alpha.ExistReq]) (*connect_go.Response[v0alpha.ExistRes], error)
	// retreive content from learning pages
	Content(context.Context, *connect_go.Request[v0alpha.ContentReq]) (*connect_go.Response[v0alpha.ContentRes], error)
	// exports multiple pages of the learning center markdown as PDF
	ExportMany(context.Context, *connect_go.Request[v0alpha.ExportManyReq]) (*connect_go.Response[v0alpha.ExportRes], error)
	// search content in learning pages
	// we allow all the logged in agents/admins to view search content
	SearchContent(context.Context, *connect_go.Request[v0alpha.SearchContentReq]) (*connect_go.Response[v0alpha.SearchRes], error)
	// stream search content results in learning pages
	// we allow all the logged in agents/admins to view search content
	ListSearchResults(context.Context, *connect_go.Request[v0alpha.SearchContentReq]) (*connect_go.ServerStreamForClient[v0alpha.SearchRes], error)
	// get standalone articles from learning pages
	// we allow all the logged in agents/admins to view standalone articles
	Standalone(context.Context, *connect_go.Request[v0alpha.StandaloneReq]) (*connect_go.Response[v0alpha.StandaloneRes], error)
	// retrieve user who edited the content last
	ContentEditorData(context.Context, *connect_go.Request[v0alpha.ContentEditorDataReq]) (*connect_go.Response[v0alpha.ContentEditorDataRes], error)
	// update contents for learning pages
	Update(context.Context, *connect_go.Request[v0alpha.UpdateReq]) (*connect_go.Response[v0alpha.UpdateRes], error)
	// upload url for static images
	StoreStaticImage(context.Context, *connect_go.Request[v0alpha.StoreStaticImageReq]) (*connect_go.Response[v0alpha.StoreStaticImageRes], error)
	// upload dynamic learning image screenshot
	UploadDynamicScreenshot(context.Context, *connect_go.Request[v0alpha.UploadDynamicScreenshotReq]) (*connect_go.Response[v0alpha.UploadDynamicScreenshotRes], error)
	// delete standalone articles from learning pages
	DeleteStandalone(context.Context, *connect_go.Request[v0alpha.DeleteStandaloneReq]) (*connect_go.Response[v0alpha.DeleteStandaloneRes], error)
	// get snippet content from learning pages
	// we allow all the logged in agents/admins to view snippet content
	Snippet(context.Context, *connect_go.Request[v0alpha.SnippetReq]) (*connect_go.Response[v0alpha.SnippetRes], error)
	// delete learning pages
	DeleteLearnPages(context.Context, *connect_go.Request[v0alpha.DeleteLearnPagesReq]) (*connect_go.Response[v0alpha.DeleteLearnPagesRes], error)
}

// NewLearnClient constructs a client for the api.v0alpha.Learn service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewLearnClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) LearnClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &learnClient{
		exist: connect_go.NewClient[v0alpha.ExistReq, v0alpha.ExistRes](
			httpClient,
			baseURL+LearnExistProcedure,
			opts...,
		),
		content: connect_go.NewClient[v0alpha.ContentReq, v0alpha.ContentRes](
			httpClient,
			baseURL+LearnContentProcedure,
			opts...,
		),
		exportMany: connect_go.NewClient[v0alpha.ExportManyReq, v0alpha.ExportRes](
			httpClient,
			baseURL+LearnExportManyProcedure,
			opts...,
		),
		searchContent: connect_go.NewClient[v0alpha.SearchContentReq, v0alpha.SearchRes](
			httpClient,
			baseURL+LearnSearchContentProcedure,
			opts...,
		),
		listSearchResults: connect_go.NewClient[v0alpha.SearchContentReq, v0alpha.SearchRes](
			httpClient,
			baseURL+LearnListSearchResultsProcedure,
			opts...,
		),
		standalone: connect_go.NewClient[v0alpha.StandaloneReq, v0alpha.StandaloneRes](
			httpClient,
			baseURL+LearnStandaloneProcedure,
			opts...,
		),
		contentEditorData: connect_go.NewClient[v0alpha.ContentEditorDataReq, v0alpha.ContentEditorDataRes](
			httpClient,
			baseURL+LearnContentEditorDataProcedure,
			opts...,
		),
		update: connect_go.NewClient[v0alpha.UpdateReq, v0alpha.UpdateRes](
			httpClient,
			baseURL+LearnUpdateProcedure,
			opts...,
		),
		storeStaticImage: connect_go.NewClient[v0alpha.StoreStaticImageReq, v0alpha.StoreStaticImageRes](
			httpClient,
			baseURL+LearnStoreStaticImageProcedure,
			opts...,
		),
		uploadDynamicScreenshot: connect_go.NewClient[v0alpha.UploadDynamicScreenshotReq, v0alpha.UploadDynamicScreenshotRes](
			httpClient,
			baseURL+LearnUploadDynamicScreenshotProcedure,
			opts...,
		),
		deleteStandalone: connect_go.NewClient[v0alpha.DeleteStandaloneReq, v0alpha.DeleteStandaloneRes](
			httpClient,
			baseURL+LearnDeleteStandaloneProcedure,
			opts...,
		),
		snippet: connect_go.NewClient[v0alpha.SnippetReq, v0alpha.SnippetRes](
			httpClient,
			baseURL+LearnSnippetProcedure,
			opts...,
		),
		deleteLearnPages: connect_go.NewClient[v0alpha.DeleteLearnPagesReq, v0alpha.DeleteLearnPagesRes](
			httpClient,
			baseURL+LearnDeleteLearnPagesProcedure,
			opts...,
		),
	}
}

// learnClient implements LearnClient.
type learnClient struct {
	exist                   *connect_go.Client[v0alpha.ExistReq, v0alpha.ExistRes]
	content                 *connect_go.Client[v0alpha.ContentReq, v0alpha.ContentRes]
	exportMany              *connect_go.Client[v0alpha.ExportManyReq, v0alpha.ExportRes]
	searchContent           *connect_go.Client[v0alpha.SearchContentReq, v0alpha.SearchRes]
	listSearchResults       *connect_go.Client[v0alpha.SearchContentReq, v0alpha.SearchRes]
	standalone              *connect_go.Client[v0alpha.StandaloneReq, v0alpha.StandaloneRes]
	contentEditorData       *connect_go.Client[v0alpha.ContentEditorDataReq, v0alpha.ContentEditorDataRes]
	update                  *connect_go.Client[v0alpha.UpdateReq, v0alpha.UpdateRes]
	storeStaticImage        *connect_go.Client[v0alpha.StoreStaticImageReq, v0alpha.StoreStaticImageRes]
	uploadDynamicScreenshot *connect_go.Client[v0alpha.UploadDynamicScreenshotReq, v0alpha.UploadDynamicScreenshotRes]
	deleteStandalone        *connect_go.Client[v0alpha.DeleteStandaloneReq, v0alpha.DeleteStandaloneRes]
	snippet                 *connect_go.Client[v0alpha.SnippetReq, v0alpha.SnippetRes]
	deleteLearnPages        *connect_go.Client[v0alpha.DeleteLearnPagesReq, v0alpha.DeleteLearnPagesRes]
}

// Exist calls api.v0alpha.Learn.Exist.
func (c *learnClient) Exist(ctx context.Context, req *connect_go.Request[v0alpha.ExistReq]) (*connect_go.Response[v0alpha.ExistRes], error) {
	return c.exist.CallUnary(ctx, req)
}

// Content calls api.v0alpha.Learn.Content.
func (c *learnClient) Content(ctx context.Context, req *connect_go.Request[v0alpha.ContentReq]) (*connect_go.Response[v0alpha.ContentRes], error) {
	return c.content.CallUnary(ctx, req)
}

// ExportMany calls api.v0alpha.Learn.ExportMany.
func (c *learnClient) ExportMany(ctx context.Context, req *connect_go.Request[v0alpha.ExportManyReq]) (*connect_go.Response[v0alpha.ExportRes], error) {
	return c.exportMany.CallUnary(ctx, req)
}

// SearchContent calls api.v0alpha.Learn.SearchContent.
func (c *learnClient) SearchContent(ctx context.Context, req *connect_go.Request[v0alpha.SearchContentReq]) (*connect_go.Response[v0alpha.SearchRes], error) {
	return c.searchContent.CallUnary(ctx, req)
}

// ListSearchResults calls api.v0alpha.Learn.ListSearchResults.
func (c *learnClient) ListSearchResults(ctx context.Context, req *connect_go.Request[v0alpha.SearchContentReq]) (*connect_go.ServerStreamForClient[v0alpha.SearchRes], error) {
	return c.listSearchResults.CallServerStream(ctx, req)
}

// Standalone calls api.v0alpha.Learn.Standalone.
func (c *learnClient) Standalone(ctx context.Context, req *connect_go.Request[v0alpha.StandaloneReq]) (*connect_go.Response[v0alpha.StandaloneRes], error) {
	return c.standalone.CallUnary(ctx, req)
}

// ContentEditorData calls api.v0alpha.Learn.ContentEditorData.
func (c *learnClient) ContentEditorData(ctx context.Context, req *connect_go.Request[v0alpha.ContentEditorDataReq]) (*connect_go.Response[v0alpha.ContentEditorDataRes], error) {
	return c.contentEditorData.CallUnary(ctx, req)
}

// Update calls api.v0alpha.Learn.Update.
func (c *learnClient) Update(ctx context.Context, req *connect_go.Request[v0alpha.UpdateReq]) (*connect_go.Response[v0alpha.UpdateRes], error) {
	return c.update.CallUnary(ctx, req)
}

// StoreStaticImage calls api.v0alpha.Learn.StoreStaticImage.
func (c *learnClient) StoreStaticImage(ctx context.Context, req *connect_go.Request[v0alpha.StoreStaticImageReq]) (*connect_go.Response[v0alpha.StoreStaticImageRes], error) {
	return c.storeStaticImage.CallUnary(ctx, req)
}

// UploadDynamicScreenshot calls api.v0alpha.Learn.UploadDynamicScreenshot.
func (c *learnClient) UploadDynamicScreenshot(ctx context.Context, req *connect_go.Request[v0alpha.UploadDynamicScreenshotReq]) (*connect_go.Response[v0alpha.UploadDynamicScreenshotRes], error) {
	return c.uploadDynamicScreenshot.CallUnary(ctx, req)
}

// DeleteStandalone calls api.v0alpha.Learn.DeleteStandalone.
func (c *learnClient) DeleteStandalone(ctx context.Context, req *connect_go.Request[v0alpha.DeleteStandaloneReq]) (*connect_go.Response[v0alpha.DeleteStandaloneRes], error) {
	return c.deleteStandalone.CallUnary(ctx, req)
}

// Snippet calls api.v0alpha.Learn.Snippet.
func (c *learnClient) Snippet(ctx context.Context, req *connect_go.Request[v0alpha.SnippetReq]) (*connect_go.Response[v0alpha.SnippetRes], error) {
	return c.snippet.CallUnary(ctx, req)
}

// DeleteLearnPages calls api.v0alpha.Learn.DeleteLearnPages.
func (c *learnClient) DeleteLearnPages(ctx context.Context, req *connect_go.Request[v0alpha.DeleteLearnPagesReq]) (*connect_go.Response[v0alpha.DeleteLearnPagesRes], error) {
	return c.deleteLearnPages.CallUnary(ctx, req)
}

// LearnHandler is an implementation of the api.v0alpha.Learn service.
type LearnHandler interface {
	// check if learning page already exists
	Exist(context.Context, *connect_go.Request[v0alpha.ExistReq]) (*connect_go.Response[v0alpha.ExistRes], error)
	// retreive content from learning pages
	Content(context.Context, *connect_go.Request[v0alpha.ContentReq]) (*connect_go.Response[v0alpha.ContentRes], error)
	// exports multiple pages of the learning center markdown as PDF
	ExportMany(context.Context, *connect_go.Request[v0alpha.ExportManyReq]) (*connect_go.Response[v0alpha.ExportRes], error)
	// search content in learning pages
	// we allow all the logged in agents/admins to view search content
	SearchContent(context.Context, *connect_go.Request[v0alpha.SearchContentReq]) (*connect_go.Response[v0alpha.SearchRes], error)
	// stream search content results in learning pages
	// we allow all the logged in agents/admins to view search content
	ListSearchResults(context.Context, *connect_go.Request[v0alpha.SearchContentReq], *connect_go.ServerStream[v0alpha.SearchRes]) error
	// get standalone articles from learning pages
	// we allow all the logged in agents/admins to view standalone articles
	Standalone(context.Context, *connect_go.Request[v0alpha.StandaloneReq]) (*connect_go.Response[v0alpha.StandaloneRes], error)
	// retrieve user who edited the content last
	ContentEditorData(context.Context, *connect_go.Request[v0alpha.ContentEditorDataReq]) (*connect_go.Response[v0alpha.ContentEditorDataRes], error)
	// update contents for learning pages
	Update(context.Context, *connect_go.Request[v0alpha.UpdateReq]) (*connect_go.Response[v0alpha.UpdateRes], error)
	// upload url for static images
	StoreStaticImage(context.Context, *connect_go.Request[v0alpha.StoreStaticImageReq]) (*connect_go.Response[v0alpha.StoreStaticImageRes], error)
	// upload dynamic learning image screenshot
	UploadDynamicScreenshot(context.Context, *connect_go.Request[v0alpha.UploadDynamicScreenshotReq]) (*connect_go.Response[v0alpha.UploadDynamicScreenshotRes], error)
	// delete standalone articles from learning pages
	DeleteStandalone(context.Context, *connect_go.Request[v0alpha.DeleteStandaloneReq]) (*connect_go.Response[v0alpha.DeleteStandaloneRes], error)
	// get snippet content from learning pages
	// we allow all the logged in agents/admins to view snippet content
	Snippet(context.Context, *connect_go.Request[v0alpha.SnippetReq]) (*connect_go.Response[v0alpha.SnippetRes], error)
	// delete learning pages
	DeleteLearnPages(context.Context, *connect_go.Request[v0alpha.DeleteLearnPagesReq]) (*connect_go.Response[v0alpha.DeleteLearnPagesRes], error)
}

// NewLearnHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewLearnHandler(svc LearnHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	learnExistHandler := connect_go.NewUnaryHandler(
		LearnExistProcedure,
		svc.Exist,
		opts...,
	)
	learnContentHandler := connect_go.NewUnaryHandler(
		LearnContentProcedure,
		svc.Content,
		opts...,
	)
	learnExportManyHandler := connect_go.NewUnaryHandler(
		LearnExportManyProcedure,
		svc.ExportMany,
		opts...,
	)
	learnSearchContentHandler := connect_go.NewUnaryHandler(
		LearnSearchContentProcedure,
		svc.SearchContent,
		opts...,
	)
	learnListSearchResultsHandler := connect_go.NewServerStreamHandler(
		LearnListSearchResultsProcedure,
		svc.ListSearchResults,
		opts...,
	)
	learnStandaloneHandler := connect_go.NewUnaryHandler(
		LearnStandaloneProcedure,
		svc.Standalone,
		opts...,
	)
	learnContentEditorDataHandler := connect_go.NewUnaryHandler(
		LearnContentEditorDataProcedure,
		svc.ContentEditorData,
		opts...,
	)
	learnUpdateHandler := connect_go.NewUnaryHandler(
		LearnUpdateProcedure,
		svc.Update,
		opts...,
	)
	learnStoreStaticImageHandler := connect_go.NewUnaryHandler(
		LearnStoreStaticImageProcedure,
		svc.StoreStaticImage,
		opts...,
	)
	learnUploadDynamicScreenshotHandler := connect_go.NewUnaryHandler(
		LearnUploadDynamicScreenshotProcedure,
		svc.UploadDynamicScreenshot,
		opts...,
	)
	learnDeleteStandaloneHandler := connect_go.NewUnaryHandler(
		LearnDeleteStandaloneProcedure,
		svc.DeleteStandalone,
		opts...,
	)
	learnSnippetHandler := connect_go.NewUnaryHandler(
		LearnSnippetProcedure,
		svc.Snippet,
		opts...,
	)
	learnDeleteLearnPagesHandler := connect_go.NewUnaryHandler(
		LearnDeleteLearnPagesProcedure,
		svc.DeleteLearnPages,
		opts...,
	)
	return "/api.v0alpha.Learn/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case LearnExistProcedure:
			learnExistHandler.ServeHTTP(w, r)
		case LearnContentProcedure:
			learnContentHandler.ServeHTTP(w, r)
		case LearnExportManyProcedure:
			learnExportManyHandler.ServeHTTP(w, r)
		case LearnSearchContentProcedure:
			learnSearchContentHandler.ServeHTTP(w, r)
		case LearnListSearchResultsProcedure:
			learnListSearchResultsHandler.ServeHTTP(w, r)
		case LearnStandaloneProcedure:
			learnStandaloneHandler.ServeHTTP(w, r)
		case LearnContentEditorDataProcedure:
			learnContentEditorDataHandler.ServeHTTP(w, r)
		case LearnUpdateProcedure:
			learnUpdateHandler.ServeHTTP(w, r)
		case LearnStoreStaticImageProcedure:
			learnStoreStaticImageHandler.ServeHTTP(w, r)
		case LearnUploadDynamicScreenshotProcedure:
			learnUploadDynamicScreenshotHandler.ServeHTTP(w, r)
		case LearnDeleteStandaloneProcedure:
			learnDeleteStandaloneHandler.ServeHTTP(w, r)
		case LearnSnippetProcedure:
			learnSnippetHandler.ServeHTTP(w, r)
		case LearnDeleteLearnPagesProcedure:
			learnDeleteLearnPagesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedLearnHandler returns CodeUnimplemented from all methods.
type UnimplementedLearnHandler struct{}

func (UnimplementedLearnHandler) Exist(context.Context, *connect_go.Request[v0alpha.ExistReq]) (*connect_go.Response[v0alpha.ExistRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Learn.Exist is not implemented"))
}

func (UnimplementedLearnHandler) Content(context.Context, *connect_go.Request[v0alpha.ContentReq]) (*connect_go.Response[v0alpha.ContentRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Learn.Content is not implemented"))
}

func (UnimplementedLearnHandler) ExportMany(context.Context, *connect_go.Request[v0alpha.ExportManyReq]) (*connect_go.Response[v0alpha.ExportRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Learn.ExportMany is not implemented"))
}

func (UnimplementedLearnHandler) SearchContent(context.Context, *connect_go.Request[v0alpha.SearchContentReq]) (*connect_go.Response[v0alpha.SearchRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Learn.SearchContent is not implemented"))
}

func (UnimplementedLearnHandler) ListSearchResults(context.Context, *connect_go.Request[v0alpha.SearchContentReq], *connect_go.ServerStream[v0alpha.SearchRes]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Learn.ListSearchResults is not implemented"))
}

func (UnimplementedLearnHandler) Standalone(context.Context, *connect_go.Request[v0alpha.StandaloneReq]) (*connect_go.Response[v0alpha.StandaloneRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Learn.Standalone is not implemented"))
}

func (UnimplementedLearnHandler) ContentEditorData(context.Context, *connect_go.Request[v0alpha.ContentEditorDataReq]) (*connect_go.Response[v0alpha.ContentEditorDataRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Learn.ContentEditorData is not implemented"))
}

func (UnimplementedLearnHandler) Update(context.Context, *connect_go.Request[v0alpha.UpdateReq]) (*connect_go.Response[v0alpha.UpdateRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Learn.Update is not implemented"))
}

func (UnimplementedLearnHandler) StoreStaticImage(context.Context, *connect_go.Request[v0alpha.StoreStaticImageReq]) (*connect_go.Response[v0alpha.StoreStaticImageRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Learn.StoreStaticImage is not implemented"))
}

func (UnimplementedLearnHandler) UploadDynamicScreenshot(context.Context, *connect_go.Request[v0alpha.UploadDynamicScreenshotReq]) (*connect_go.Response[v0alpha.UploadDynamicScreenshotRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Learn.UploadDynamicScreenshot is not implemented"))
}

func (UnimplementedLearnHandler) DeleteStandalone(context.Context, *connect_go.Request[v0alpha.DeleteStandaloneReq]) (*connect_go.Response[v0alpha.DeleteStandaloneRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Learn.DeleteStandalone is not implemented"))
}

func (UnimplementedLearnHandler) Snippet(context.Context, *connect_go.Request[v0alpha.SnippetReq]) (*connect_go.Response[v0alpha.SnippetRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Learn.Snippet is not implemented"))
}

func (UnimplementedLearnHandler) DeleteLearnPages(context.Context, *connect_go.Request[v0alpha.DeleteLearnPagesReq]) (*connect_go.Response[v0alpha.DeleteLearnPagesRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Learn.DeleteLearnPages is not implemented"))
}
