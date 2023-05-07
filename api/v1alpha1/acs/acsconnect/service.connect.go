// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1alpha1/acs/service.proto

package acsconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	acs "github.com/tcncloud/api-go/v1alpha1/acs"
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
	// AcsName is the fully-qualified name of the Acs service.
	AcsName = "api.v1alpha1.acs.Acs"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AcsCreateScriptProcedure is the fully-qualified name of the Acs's CreateScript RPC.
	AcsCreateScriptProcedure = "/api.v1alpha1.acs.Acs/CreateScript"
	// AcsUpdateScriptProcedure is the fully-qualified name of the Acs's UpdateScript RPC.
	AcsUpdateScriptProcedure = "/api.v1alpha1.acs.Acs/UpdateScript"
	// AcsGetScriptByIdProcedure is the fully-qualified name of the Acs's GetScriptById RPC.
	AcsGetScriptByIdProcedure = "/api.v1alpha1.acs.Acs/GetScriptById"
	// AcsListScriptsProcedure is the fully-qualified name of the Acs's ListScripts RPC.
	AcsListScriptsProcedure = "/api.v1alpha1.acs.Acs/ListScripts"
	// AcsSaveResponsesProcedure is the fully-qualified name of the Acs's SaveResponses RPC.
	AcsSaveResponsesProcedure = "/api.v1alpha1.acs.Acs/SaveResponses"
	// AcsGetResponsesProcedure is the fully-qualified name of the Acs's GetResponses RPC.
	AcsGetResponsesProcedure = "/api.v1alpha1.acs.Acs/GetResponses"
	// AcsCreateRuleProcedure is the fully-qualified name of the Acs's CreateRule RPC.
	AcsCreateRuleProcedure = "/api.v1alpha1.acs.Acs/CreateRule"
	// AcsUpdateRulesProcedure is the fully-qualified name of the Acs's UpdateRules RPC.
	AcsUpdateRulesProcedure = "/api.v1alpha1.acs.Acs/UpdateRules"
	// AcsListRulesProcedure is the fully-qualified name of the Acs's ListRules RPC.
	AcsListRulesProcedure = "/api.v1alpha1.acs.Acs/ListRules"
	// AcsCreateRegexEvaluatorProcedure is the fully-qualified name of the Acs's CreateRegexEvaluator
	// RPC.
	AcsCreateRegexEvaluatorProcedure = "/api.v1alpha1.acs.Acs/CreateRegexEvaluator"
	// AcsUpdateRegexEvaluatorProcedure is the fully-qualified name of the Acs's UpdateRegexEvaluator
	// RPC.
	AcsUpdateRegexEvaluatorProcedure = "/api.v1alpha1.acs.Acs/UpdateRegexEvaluator"
	// AcsListRegexEvaluatorsProcedure is the fully-qualified name of the Acs's ListRegexEvaluators RPC.
	AcsListRegexEvaluatorsProcedure = "/api.v1alpha1.acs.Acs/ListRegexEvaluators"
	// AcsDeleteRuleProcedure is the fully-qualified name of the Acs's DeleteRule RPC.
	AcsDeleteRuleProcedure = "/api.v1alpha1.acs.Acs/DeleteRule"
	// AcsGetScriptProcedure is the fully-qualified name of the Acs's GetScript RPC.
	AcsGetScriptProcedure = "/api.v1alpha1.acs.Acs/GetScript"
)

// AcsClient is a client for the api.v1alpha1.acs.Acs service.
type AcsClient interface {
	CreateScript(context.Context, *connect_go.Request[acs.CreateScriptRequest]) (*connect_go.Response[acs.CreateScriptResponse], error)
	UpdateScript(context.Context, *connect_go.Request[acs.UpdateScriptRequest]) (*connect_go.Response[acs.UpdateScriptResponse], error)
	GetScriptById(context.Context, *connect_go.Request[acs.GetScriptByIdRequest]) (*connect_go.Response[acs.GetScriptByIdResponse], error)
	ListScripts(context.Context, *connect_go.Request[acs.ListScriptsRequest]) (*connect_go.ServerStreamForClient[acs.ScriptDetails], error)
	SaveResponses(context.Context, *connect_go.Request[acs.SaveResponsesRequest]) (*connect_go.Response[acs.SaveResponsesResponse], error)
	GetResponses(context.Context, *connect_go.Request[acs.GetResponsesRequest]) (*connect_go.Response[acs.GetResponsesResponse], error)
	CreateRule(context.Context, *connect_go.Request[acs.CreateRuleRequest]) (*connect_go.Response[acs.CreateRuleResponse], error)
	UpdateRules(context.Context, *connect_go.Request[acs.UpdateRulesRequest]) (*connect_go.Response[acs.UpdateRulesResponse], error)
	ListRules(context.Context, *connect_go.Request[acs.ListRulesRequest]) (*connect_go.Response[acs.ListRulesResponse], error)
	CreateRegexEvaluator(context.Context, *connect_go.Request[acs.CreateRegexEvaluatorRequest]) (*connect_go.Response[acs.CreateRegexEvaluatorResponse], error)
	UpdateRegexEvaluator(context.Context, *connect_go.Request[acs.UpdateRegexEvaluatorRequest]) (*connect_go.Response[acs.UpdateRegexEvaluatorResponse], error)
	ListRegexEvaluators(context.Context, *connect_go.Request[acs.ListRegexEvaluatorsRequest]) (*connect_go.Response[acs.ListRegexEvaluatorsResponse], error)
	DeleteRule(context.Context, *connect_go.Request[acs.DeleteRuleRequest]) (*connect_go.Response[acs.DeleteRuleResponse], error)
	GetScript(context.Context, *connect_go.Request[acs.GetScriptRequest]) (*connect_go.Response[acs.GetScriptResponse], error)
}

// NewAcsClient constructs a client for the api.v1alpha1.acs.Acs service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAcsClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AcsClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &acsClient{
		createScript: connect_go.NewClient[acs.CreateScriptRequest, acs.CreateScriptResponse](
			httpClient,
			baseURL+AcsCreateScriptProcedure,
			opts...,
		),
		updateScript: connect_go.NewClient[acs.UpdateScriptRequest, acs.UpdateScriptResponse](
			httpClient,
			baseURL+AcsUpdateScriptProcedure,
			opts...,
		),
		getScriptById: connect_go.NewClient[acs.GetScriptByIdRequest, acs.GetScriptByIdResponse](
			httpClient,
			baseURL+AcsGetScriptByIdProcedure,
			opts...,
		),
		listScripts: connect_go.NewClient[acs.ListScriptsRequest, acs.ScriptDetails](
			httpClient,
			baseURL+AcsListScriptsProcedure,
			opts...,
		),
		saveResponses: connect_go.NewClient[acs.SaveResponsesRequest, acs.SaveResponsesResponse](
			httpClient,
			baseURL+AcsSaveResponsesProcedure,
			opts...,
		),
		getResponses: connect_go.NewClient[acs.GetResponsesRequest, acs.GetResponsesResponse](
			httpClient,
			baseURL+AcsGetResponsesProcedure,
			opts...,
		),
		createRule: connect_go.NewClient[acs.CreateRuleRequest, acs.CreateRuleResponse](
			httpClient,
			baseURL+AcsCreateRuleProcedure,
			opts...,
		),
		updateRules: connect_go.NewClient[acs.UpdateRulesRequest, acs.UpdateRulesResponse](
			httpClient,
			baseURL+AcsUpdateRulesProcedure,
			opts...,
		),
		listRules: connect_go.NewClient[acs.ListRulesRequest, acs.ListRulesResponse](
			httpClient,
			baseURL+AcsListRulesProcedure,
			opts...,
		),
		createRegexEvaluator: connect_go.NewClient[acs.CreateRegexEvaluatorRequest, acs.CreateRegexEvaluatorResponse](
			httpClient,
			baseURL+AcsCreateRegexEvaluatorProcedure,
			opts...,
		),
		updateRegexEvaluator: connect_go.NewClient[acs.UpdateRegexEvaluatorRequest, acs.UpdateRegexEvaluatorResponse](
			httpClient,
			baseURL+AcsUpdateRegexEvaluatorProcedure,
			opts...,
		),
		listRegexEvaluators: connect_go.NewClient[acs.ListRegexEvaluatorsRequest, acs.ListRegexEvaluatorsResponse](
			httpClient,
			baseURL+AcsListRegexEvaluatorsProcedure,
			opts...,
		),
		deleteRule: connect_go.NewClient[acs.DeleteRuleRequest, acs.DeleteRuleResponse](
			httpClient,
			baseURL+AcsDeleteRuleProcedure,
			opts...,
		),
		getScript: connect_go.NewClient[acs.GetScriptRequest, acs.GetScriptResponse](
			httpClient,
			baseURL+AcsGetScriptProcedure,
			opts...,
		),
	}
}

// acsClient implements AcsClient.
type acsClient struct {
	createScript         *connect_go.Client[acs.CreateScriptRequest, acs.CreateScriptResponse]
	updateScript         *connect_go.Client[acs.UpdateScriptRequest, acs.UpdateScriptResponse]
	getScriptById        *connect_go.Client[acs.GetScriptByIdRequest, acs.GetScriptByIdResponse]
	listScripts          *connect_go.Client[acs.ListScriptsRequest, acs.ScriptDetails]
	saveResponses        *connect_go.Client[acs.SaveResponsesRequest, acs.SaveResponsesResponse]
	getResponses         *connect_go.Client[acs.GetResponsesRequest, acs.GetResponsesResponse]
	createRule           *connect_go.Client[acs.CreateRuleRequest, acs.CreateRuleResponse]
	updateRules          *connect_go.Client[acs.UpdateRulesRequest, acs.UpdateRulesResponse]
	listRules            *connect_go.Client[acs.ListRulesRequest, acs.ListRulesResponse]
	createRegexEvaluator *connect_go.Client[acs.CreateRegexEvaluatorRequest, acs.CreateRegexEvaluatorResponse]
	updateRegexEvaluator *connect_go.Client[acs.UpdateRegexEvaluatorRequest, acs.UpdateRegexEvaluatorResponse]
	listRegexEvaluators  *connect_go.Client[acs.ListRegexEvaluatorsRequest, acs.ListRegexEvaluatorsResponse]
	deleteRule           *connect_go.Client[acs.DeleteRuleRequest, acs.DeleteRuleResponse]
	getScript            *connect_go.Client[acs.GetScriptRequest, acs.GetScriptResponse]
}

// CreateScript calls api.v1alpha1.acs.Acs.CreateScript.
func (c *acsClient) CreateScript(ctx context.Context, req *connect_go.Request[acs.CreateScriptRequest]) (*connect_go.Response[acs.CreateScriptResponse], error) {
	return c.createScript.CallUnary(ctx, req)
}

// UpdateScript calls api.v1alpha1.acs.Acs.UpdateScript.
func (c *acsClient) UpdateScript(ctx context.Context, req *connect_go.Request[acs.UpdateScriptRequest]) (*connect_go.Response[acs.UpdateScriptResponse], error) {
	return c.updateScript.CallUnary(ctx, req)
}

// GetScriptById calls api.v1alpha1.acs.Acs.GetScriptById.
func (c *acsClient) GetScriptById(ctx context.Context, req *connect_go.Request[acs.GetScriptByIdRequest]) (*connect_go.Response[acs.GetScriptByIdResponse], error) {
	return c.getScriptById.CallUnary(ctx, req)
}

// ListScripts calls api.v1alpha1.acs.Acs.ListScripts.
func (c *acsClient) ListScripts(ctx context.Context, req *connect_go.Request[acs.ListScriptsRequest]) (*connect_go.ServerStreamForClient[acs.ScriptDetails], error) {
	return c.listScripts.CallServerStream(ctx, req)
}

// SaveResponses calls api.v1alpha1.acs.Acs.SaveResponses.
func (c *acsClient) SaveResponses(ctx context.Context, req *connect_go.Request[acs.SaveResponsesRequest]) (*connect_go.Response[acs.SaveResponsesResponse], error) {
	return c.saveResponses.CallUnary(ctx, req)
}

// GetResponses calls api.v1alpha1.acs.Acs.GetResponses.
func (c *acsClient) GetResponses(ctx context.Context, req *connect_go.Request[acs.GetResponsesRequest]) (*connect_go.Response[acs.GetResponsesResponse], error) {
	return c.getResponses.CallUnary(ctx, req)
}

// CreateRule calls api.v1alpha1.acs.Acs.CreateRule.
func (c *acsClient) CreateRule(ctx context.Context, req *connect_go.Request[acs.CreateRuleRequest]) (*connect_go.Response[acs.CreateRuleResponse], error) {
	return c.createRule.CallUnary(ctx, req)
}

// UpdateRules calls api.v1alpha1.acs.Acs.UpdateRules.
func (c *acsClient) UpdateRules(ctx context.Context, req *connect_go.Request[acs.UpdateRulesRequest]) (*connect_go.Response[acs.UpdateRulesResponse], error) {
	return c.updateRules.CallUnary(ctx, req)
}

// ListRules calls api.v1alpha1.acs.Acs.ListRules.
func (c *acsClient) ListRules(ctx context.Context, req *connect_go.Request[acs.ListRulesRequest]) (*connect_go.Response[acs.ListRulesResponse], error) {
	return c.listRules.CallUnary(ctx, req)
}

// CreateRegexEvaluator calls api.v1alpha1.acs.Acs.CreateRegexEvaluator.
func (c *acsClient) CreateRegexEvaluator(ctx context.Context, req *connect_go.Request[acs.CreateRegexEvaluatorRequest]) (*connect_go.Response[acs.CreateRegexEvaluatorResponse], error) {
	return c.createRegexEvaluator.CallUnary(ctx, req)
}

// UpdateRegexEvaluator calls api.v1alpha1.acs.Acs.UpdateRegexEvaluator.
func (c *acsClient) UpdateRegexEvaluator(ctx context.Context, req *connect_go.Request[acs.UpdateRegexEvaluatorRequest]) (*connect_go.Response[acs.UpdateRegexEvaluatorResponse], error) {
	return c.updateRegexEvaluator.CallUnary(ctx, req)
}

// ListRegexEvaluators calls api.v1alpha1.acs.Acs.ListRegexEvaluators.
func (c *acsClient) ListRegexEvaluators(ctx context.Context, req *connect_go.Request[acs.ListRegexEvaluatorsRequest]) (*connect_go.Response[acs.ListRegexEvaluatorsResponse], error) {
	return c.listRegexEvaluators.CallUnary(ctx, req)
}

// DeleteRule calls api.v1alpha1.acs.Acs.DeleteRule.
func (c *acsClient) DeleteRule(ctx context.Context, req *connect_go.Request[acs.DeleteRuleRequest]) (*connect_go.Response[acs.DeleteRuleResponse], error) {
	return c.deleteRule.CallUnary(ctx, req)
}

// GetScript calls api.v1alpha1.acs.Acs.GetScript.
func (c *acsClient) GetScript(ctx context.Context, req *connect_go.Request[acs.GetScriptRequest]) (*connect_go.Response[acs.GetScriptResponse], error) {
	return c.getScript.CallUnary(ctx, req)
}

// AcsHandler is an implementation of the api.v1alpha1.acs.Acs service.
type AcsHandler interface {
	CreateScript(context.Context, *connect_go.Request[acs.CreateScriptRequest]) (*connect_go.Response[acs.CreateScriptResponse], error)
	UpdateScript(context.Context, *connect_go.Request[acs.UpdateScriptRequest]) (*connect_go.Response[acs.UpdateScriptResponse], error)
	GetScriptById(context.Context, *connect_go.Request[acs.GetScriptByIdRequest]) (*connect_go.Response[acs.GetScriptByIdResponse], error)
	ListScripts(context.Context, *connect_go.Request[acs.ListScriptsRequest], *connect_go.ServerStream[acs.ScriptDetails]) error
	SaveResponses(context.Context, *connect_go.Request[acs.SaveResponsesRequest]) (*connect_go.Response[acs.SaveResponsesResponse], error)
	GetResponses(context.Context, *connect_go.Request[acs.GetResponsesRequest]) (*connect_go.Response[acs.GetResponsesResponse], error)
	CreateRule(context.Context, *connect_go.Request[acs.CreateRuleRequest]) (*connect_go.Response[acs.CreateRuleResponse], error)
	UpdateRules(context.Context, *connect_go.Request[acs.UpdateRulesRequest]) (*connect_go.Response[acs.UpdateRulesResponse], error)
	ListRules(context.Context, *connect_go.Request[acs.ListRulesRequest]) (*connect_go.Response[acs.ListRulesResponse], error)
	CreateRegexEvaluator(context.Context, *connect_go.Request[acs.CreateRegexEvaluatorRequest]) (*connect_go.Response[acs.CreateRegexEvaluatorResponse], error)
	UpdateRegexEvaluator(context.Context, *connect_go.Request[acs.UpdateRegexEvaluatorRequest]) (*connect_go.Response[acs.UpdateRegexEvaluatorResponse], error)
	ListRegexEvaluators(context.Context, *connect_go.Request[acs.ListRegexEvaluatorsRequest]) (*connect_go.Response[acs.ListRegexEvaluatorsResponse], error)
	DeleteRule(context.Context, *connect_go.Request[acs.DeleteRuleRequest]) (*connect_go.Response[acs.DeleteRuleResponse], error)
	GetScript(context.Context, *connect_go.Request[acs.GetScriptRequest]) (*connect_go.Response[acs.GetScriptResponse], error)
}

// NewAcsHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAcsHandler(svc AcsHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(AcsCreateScriptProcedure, connect_go.NewUnaryHandler(
		AcsCreateScriptProcedure,
		svc.CreateScript,
		opts...,
	))
	mux.Handle(AcsUpdateScriptProcedure, connect_go.NewUnaryHandler(
		AcsUpdateScriptProcedure,
		svc.UpdateScript,
		opts...,
	))
	mux.Handle(AcsGetScriptByIdProcedure, connect_go.NewUnaryHandler(
		AcsGetScriptByIdProcedure,
		svc.GetScriptById,
		opts...,
	))
	mux.Handle(AcsListScriptsProcedure, connect_go.NewServerStreamHandler(
		AcsListScriptsProcedure,
		svc.ListScripts,
		opts...,
	))
	mux.Handle(AcsSaveResponsesProcedure, connect_go.NewUnaryHandler(
		AcsSaveResponsesProcedure,
		svc.SaveResponses,
		opts...,
	))
	mux.Handle(AcsGetResponsesProcedure, connect_go.NewUnaryHandler(
		AcsGetResponsesProcedure,
		svc.GetResponses,
		opts...,
	))
	mux.Handle(AcsCreateRuleProcedure, connect_go.NewUnaryHandler(
		AcsCreateRuleProcedure,
		svc.CreateRule,
		opts...,
	))
	mux.Handle(AcsUpdateRulesProcedure, connect_go.NewUnaryHandler(
		AcsUpdateRulesProcedure,
		svc.UpdateRules,
		opts...,
	))
	mux.Handle(AcsListRulesProcedure, connect_go.NewUnaryHandler(
		AcsListRulesProcedure,
		svc.ListRules,
		opts...,
	))
	mux.Handle(AcsCreateRegexEvaluatorProcedure, connect_go.NewUnaryHandler(
		AcsCreateRegexEvaluatorProcedure,
		svc.CreateRegexEvaluator,
		opts...,
	))
	mux.Handle(AcsUpdateRegexEvaluatorProcedure, connect_go.NewUnaryHandler(
		AcsUpdateRegexEvaluatorProcedure,
		svc.UpdateRegexEvaluator,
		opts...,
	))
	mux.Handle(AcsListRegexEvaluatorsProcedure, connect_go.NewUnaryHandler(
		AcsListRegexEvaluatorsProcedure,
		svc.ListRegexEvaluators,
		opts...,
	))
	mux.Handle(AcsDeleteRuleProcedure, connect_go.NewUnaryHandler(
		AcsDeleteRuleProcedure,
		svc.DeleteRule,
		opts...,
	))
	mux.Handle(AcsGetScriptProcedure, connect_go.NewUnaryHandler(
		AcsGetScriptProcedure,
		svc.GetScript,
		opts...,
	))
	return "/api.v1alpha1.acs.Acs/", mux
}

// UnimplementedAcsHandler returns CodeUnimplemented from all methods.
type UnimplementedAcsHandler struct{}

func (UnimplementedAcsHandler) CreateScript(context.Context, *connect_go.Request[acs.CreateScriptRequest]) (*connect_go.Response[acs.CreateScriptResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.acs.Acs.CreateScript is not implemented"))
}

func (UnimplementedAcsHandler) UpdateScript(context.Context, *connect_go.Request[acs.UpdateScriptRequest]) (*connect_go.Response[acs.UpdateScriptResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.acs.Acs.UpdateScript is not implemented"))
}

func (UnimplementedAcsHandler) GetScriptById(context.Context, *connect_go.Request[acs.GetScriptByIdRequest]) (*connect_go.Response[acs.GetScriptByIdResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.acs.Acs.GetScriptById is not implemented"))
}

func (UnimplementedAcsHandler) ListScripts(context.Context, *connect_go.Request[acs.ListScriptsRequest], *connect_go.ServerStream[acs.ScriptDetails]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.acs.Acs.ListScripts is not implemented"))
}

func (UnimplementedAcsHandler) SaveResponses(context.Context, *connect_go.Request[acs.SaveResponsesRequest]) (*connect_go.Response[acs.SaveResponsesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.acs.Acs.SaveResponses is not implemented"))
}

func (UnimplementedAcsHandler) GetResponses(context.Context, *connect_go.Request[acs.GetResponsesRequest]) (*connect_go.Response[acs.GetResponsesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.acs.Acs.GetResponses is not implemented"))
}

func (UnimplementedAcsHandler) CreateRule(context.Context, *connect_go.Request[acs.CreateRuleRequest]) (*connect_go.Response[acs.CreateRuleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.acs.Acs.CreateRule is not implemented"))
}

func (UnimplementedAcsHandler) UpdateRules(context.Context, *connect_go.Request[acs.UpdateRulesRequest]) (*connect_go.Response[acs.UpdateRulesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.acs.Acs.UpdateRules is not implemented"))
}

func (UnimplementedAcsHandler) ListRules(context.Context, *connect_go.Request[acs.ListRulesRequest]) (*connect_go.Response[acs.ListRulesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.acs.Acs.ListRules is not implemented"))
}

func (UnimplementedAcsHandler) CreateRegexEvaluator(context.Context, *connect_go.Request[acs.CreateRegexEvaluatorRequest]) (*connect_go.Response[acs.CreateRegexEvaluatorResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.acs.Acs.CreateRegexEvaluator is not implemented"))
}

func (UnimplementedAcsHandler) UpdateRegexEvaluator(context.Context, *connect_go.Request[acs.UpdateRegexEvaluatorRequest]) (*connect_go.Response[acs.UpdateRegexEvaluatorResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.acs.Acs.UpdateRegexEvaluator is not implemented"))
}

func (UnimplementedAcsHandler) ListRegexEvaluators(context.Context, *connect_go.Request[acs.ListRegexEvaluatorsRequest]) (*connect_go.Response[acs.ListRegexEvaluatorsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.acs.Acs.ListRegexEvaluators is not implemented"))
}

func (UnimplementedAcsHandler) DeleteRule(context.Context, *connect_go.Request[acs.DeleteRuleRequest]) (*connect_go.Response[acs.DeleteRuleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.acs.Acs.DeleteRule is not implemented"))
}

func (UnimplementedAcsHandler) GetScript(context.Context, *connect_go.Request[acs.GetScriptRequest]) (*connect_go.Response[acs.GetScriptResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.acs.Acs.GetScript is not implemented"))
}
