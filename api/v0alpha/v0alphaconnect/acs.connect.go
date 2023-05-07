// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v0alpha/acs.proto

package v0alphaconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v0alpha "github.com/tcncloud/api-go/v0alpha"
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
	AcsName = "api.v0alpha.Acs"
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
	AcsCreateScriptProcedure = "/api.v0alpha.Acs/CreateScript"
	// AcsUpdateScriptProcedure is the fully-qualified name of the Acs's UpdateScript RPC.
	AcsUpdateScriptProcedure = "/api.v0alpha.Acs/UpdateScript"
	// AcsGetScriptByIdProcedure is the fully-qualified name of the Acs's GetScriptById RPC.
	AcsGetScriptByIdProcedure = "/api.v0alpha.Acs/GetScriptById"
	// AcsListScriptsProcedure is the fully-qualified name of the Acs's ListScripts RPC.
	AcsListScriptsProcedure = "/api.v0alpha.Acs/ListScripts"
	// AcsSaveResponsesProcedure is the fully-qualified name of the Acs's SaveResponses RPC.
	AcsSaveResponsesProcedure = "/api.v0alpha.Acs/SaveResponses"
	// AcsGetResponsesProcedure is the fully-qualified name of the Acs's GetResponses RPC.
	AcsGetResponsesProcedure = "/api.v0alpha.Acs/GetResponses"
	// AcsCreateRuleProcedure is the fully-qualified name of the Acs's CreateRule RPC.
	AcsCreateRuleProcedure = "/api.v0alpha.Acs/CreateRule"
	// AcsUpdateRulesProcedure is the fully-qualified name of the Acs's UpdateRules RPC.
	AcsUpdateRulesProcedure = "/api.v0alpha.Acs/UpdateRules"
	// AcsListRulesProcedure is the fully-qualified name of the Acs's ListRules RPC.
	AcsListRulesProcedure = "/api.v0alpha.Acs/ListRules"
	// AcsCreateRegexEvaluatorProcedure is the fully-qualified name of the Acs's CreateRegexEvaluator
	// RPC.
	AcsCreateRegexEvaluatorProcedure = "/api.v0alpha.Acs/CreateRegexEvaluator"
	// AcsUpdateRegexEvaluatorProcedure is the fully-qualified name of the Acs's UpdateRegexEvaluator
	// RPC.
	AcsUpdateRegexEvaluatorProcedure = "/api.v0alpha.Acs/UpdateRegexEvaluator"
	// AcsListRegexEvaluatorsProcedure is the fully-qualified name of the Acs's ListRegexEvaluators RPC.
	AcsListRegexEvaluatorsProcedure = "/api.v0alpha.Acs/ListRegexEvaluators"
	// AcsDeleteRuleProcedure is the fully-qualified name of the Acs's DeleteRule RPC.
	AcsDeleteRuleProcedure = "/api.v0alpha.Acs/DeleteRule"
	// AcsGetScriptProcedure is the fully-qualified name of the Acs's GetScript RPC.
	AcsGetScriptProcedure = "/api.v0alpha.Acs/GetScript"
)

// AcsClient is a client for the api.v0alpha.Acs service.
type AcsClient interface {
	CreateScript(context.Context, *connect_go.Request[v0alpha.CreateScriptRequest]) (*connect_go.Response[v0alpha.CreateScriptResponse], error)
	UpdateScript(context.Context, *connect_go.Request[v0alpha.UpdateScriptRequest]) (*connect_go.Response[v0alpha.UpdateScriptResponse], error)
	GetScriptById(context.Context, *connect_go.Request[v0alpha.GetScriptByIdRequest]) (*connect_go.Response[v0alpha.GetScriptByIdResponse], error)
	ListScripts(context.Context, *connect_go.Request[v0alpha.ListScriptsRequest]) (*connect_go.ServerStreamForClient[v0alpha.ScriptDetails], error)
	SaveResponses(context.Context, *connect_go.Request[v0alpha.SaveResponsesRequest]) (*connect_go.Response[v0alpha.SaveResponsesResponse], error)
	GetResponses(context.Context, *connect_go.Request[v0alpha.GetResponsesRequest]) (*connect_go.Response[v0alpha.GetResponsesResponse], error)
	CreateRule(context.Context, *connect_go.Request[v0alpha.CreateRuleRequest]) (*connect_go.Response[v0alpha.CreateRuleResponse], error)
	UpdateRules(context.Context, *connect_go.Request[v0alpha.UpdateRulesRequest]) (*connect_go.Response[v0alpha.UpdateRulesResponse], error)
	ListRules(context.Context, *connect_go.Request[v0alpha.ListRulesRequest]) (*connect_go.Response[v0alpha.ListRulesResponse], error)
	CreateRegexEvaluator(context.Context, *connect_go.Request[v0alpha.CreateRegexEvaluatorRequest]) (*connect_go.Response[v0alpha.CreateRegexEvaluatorResponse], error)
	UpdateRegexEvaluator(context.Context, *connect_go.Request[v0alpha.UpdateRegexEvaluatorRequest]) (*connect_go.Response[v0alpha.UpdateRegexEvaluatorResponse], error)
	ListRegexEvaluators(context.Context, *connect_go.Request[v0alpha.ListRegexEvaluatorsRequest]) (*connect_go.Response[v0alpha.ListRegexEvaluatorsResponse], error)
	DeleteRule(context.Context, *connect_go.Request[v0alpha.DeleteRuleRequest]) (*connect_go.Response[v0alpha.DeleteRuleResponse], error)
	GetScript(context.Context, *connect_go.Request[v0alpha.GetScriptRequest]) (*connect_go.Response[v0alpha.GetScriptResponse], error)
}

// NewAcsClient constructs a client for the api.v0alpha.Acs service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAcsClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AcsClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &acsClient{
		createScript: connect_go.NewClient[v0alpha.CreateScriptRequest, v0alpha.CreateScriptResponse](
			httpClient,
			baseURL+AcsCreateScriptProcedure,
			opts...,
		),
		updateScript: connect_go.NewClient[v0alpha.UpdateScriptRequest, v0alpha.UpdateScriptResponse](
			httpClient,
			baseURL+AcsUpdateScriptProcedure,
			opts...,
		),
		getScriptById: connect_go.NewClient[v0alpha.GetScriptByIdRequest, v0alpha.GetScriptByIdResponse](
			httpClient,
			baseURL+AcsGetScriptByIdProcedure,
			opts...,
		),
		listScripts: connect_go.NewClient[v0alpha.ListScriptsRequest, v0alpha.ScriptDetails](
			httpClient,
			baseURL+AcsListScriptsProcedure,
			opts...,
		),
		saveResponses: connect_go.NewClient[v0alpha.SaveResponsesRequest, v0alpha.SaveResponsesResponse](
			httpClient,
			baseURL+AcsSaveResponsesProcedure,
			opts...,
		),
		getResponses: connect_go.NewClient[v0alpha.GetResponsesRequest, v0alpha.GetResponsesResponse](
			httpClient,
			baseURL+AcsGetResponsesProcedure,
			opts...,
		),
		createRule: connect_go.NewClient[v0alpha.CreateRuleRequest, v0alpha.CreateRuleResponse](
			httpClient,
			baseURL+AcsCreateRuleProcedure,
			opts...,
		),
		updateRules: connect_go.NewClient[v0alpha.UpdateRulesRequest, v0alpha.UpdateRulesResponse](
			httpClient,
			baseURL+AcsUpdateRulesProcedure,
			opts...,
		),
		listRules: connect_go.NewClient[v0alpha.ListRulesRequest, v0alpha.ListRulesResponse](
			httpClient,
			baseURL+AcsListRulesProcedure,
			opts...,
		),
		createRegexEvaluator: connect_go.NewClient[v0alpha.CreateRegexEvaluatorRequest, v0alpha.CreateRegexEvaluatorResponse](
			httpClient,
			baseURL+AcsCreateRegexEvaluatorProcedure,
			opts...,
		),
		updateRegexEvaluator: connect_go.NewClient[v0alpha.UpdateRegexEvaluatorRequest, v0alpha.UpdateRegexEvaluatorResponse](
			httpClient,
			baseURL+AcsUpdateRegexEvaluatorProcedure,
			opts...,
		),
		listRegexEvaluators: connect_go.NewClient[v0alpha.ListRegexEvaluatorsRequest, v0alpha.ListRegexEvaluatorsResponse](
			httpClient,
			baseURL+AcsListRegexEvaluatorsProcedure,
			opts...,
		),
		deleteRule: connect_go.NewClient[v0alpha.DeleteRuleRequest, v0alpha.DeleteRuleResponse](
			httpClient,
			baseURL+AcsDeleteRuleProcedure,
			opts...,
		),
		getScript: connect_go.NewClient[v0alpha.GetScriptRequest, v0alpha.GetScriptResponse](
			httpClient,
			baseURL+AcsGetScriptProcedure,
			opts...,
		),
	}
}

// acsClient implements AcsClient.
type acsClient struct {
	createScript         *connect_go.Client[v0alpha.CreateScriptRequest, v0alpha.CreateScriptResponse]
	updateScript         *connect_go.Client[v0alpha.UpdateScriptRequest, v0alpha.UpdateScriptResponse]
	getScriptById        *connect_go.Client[v0alpha.GetScriptByIdRequest, v0alpha.GetScriptByIdResponse]
	listScripts          *connect_go.Client[v0alpha.ListScriptsRequest, v0alpha.ScriptDetails]
	saveResponses        *connect_go.Client[v0alpha.SaveResponsesRequest, v0alpha.SaveResponsesResponse]
	getResponses         *connect_go.Client[v0alpha.GetResponsesRequest, v0alpha.GetResponsesResponse]
	createRule           *connect_go.Client[v0alpha.CreateRuleRequest, v0alpha.CreateRuleResponse]
	updateRules          *connect_go.Client[v0alpha.UpdateRulesRequest, v0alpha.UpdateRulesResponse]
	listRules            *connect_go.Client[v0alpha.ListRulesRequest, v0alpha.ListRulesResponse]
	createRegexEvaluator *connect_go.Client[v0alpha.CreateRegexEvaluatorRequest, v0alpha.CreateRegexEvaluatorResponse]
	updateRegexEvaluator *connect_go.Client[v0alpha.UpdateRegexEvaluatorRequest, v0alpha.UpdateRegexEvaluatorResponse]
	listRegexEvaluators  *connect_go.Client[v0alpha.ListRegexEvaluatorsRequest, v0alpha.ListRegexEvaluatorsResponse]
	deleteRule           *connect_go.Client[v0alpha.DeleteRuleRequest, v0alpha.DeleteRuleResponse]
	getScript            *connect_go.Client[v0alpha.GetScriptRequest, v0alpha.GetScriptResponse]
}

// CreateScript calls api.v0alpha.Acs.CreateScript.
func (c *acsClient) CreateScript(ctx context.Context, req *connect_go.Request[v0alpha.CreateScriptRequest]) (*connect_go.Response[v0alpha.CreateScriptResponse], error) {
	return c.createScript.CallUnary(ctx, req)
}

// UpdateScript calls api.v0alpha.Acs.UpdateScript.
func (c *acsClient) UpdateScript(ctx context.Context, req *connect_go.Request[v0alpha.UpdateScriptRequest]) (*connect_go.Response[v0alpha.UpdateScriptResponse], error) {
	return c.updateScript.CallUnary(ctx, req)
}

// GetScriptById calls api.v0alpha.Acs.GetScriptById.
func (c *acsClient) GetScriptById(ctx context.Context, req *connect_go.Request[v0alpha.GetScriptByIdRequest]) (*connect_go.Response[v0alpha.GetScriptByIdResponse], error) {
	return c.getScriptById.CallUnary(ctx, req)
}

// ListScripts calls api.v0alpha.Acs.ListScripts.
func (c *acsClient) ListScripts(ctx context.Context, req *connect_go.Request[v0alpha.ListScriptsRequest]) (*connect_go.ServerStreamForClient[v0alpha.ScriptDetails], error) {
	return c.listScripts.CallServerStream(ctx, req)
}

// SaveResponses calls api.v0alpha.Acs.SaveResponses.
func (c *acsClient) SaveResponses(ctx context.Context, req *connect_go.Request[v0alpha.SaveResponsesRequest]) (*connect_go.Response[v0alpha.SaveResponsesResponse], error) {
	return c.saveResponses.CallUnary(ctx, req)
}

// GetResponses calls api.v0alpha.Acs.GetResponses.
func (c *acsClient) GetResponses(ctx context.Context, req *connect_go.Request[v0alpha.GetResponsesRequest]) (*connect_go.Response[v0alpha.GetResponsesResponse], error) {
	return c.getResponses.CallUnary(ctx, req)
}

// CreateRule calls api.v0alpha.Acs.CreateRule.
func (c *acsClient) CreateRule(ctx context.Context, req *connect_go.Request[v0alpha.CreateRuleRequest]) (*connect_go.Response[v0alpha.CreateRuleResponse], error) {
	return c.createRule.CallUnary(ctx, req)
}

// UpdateRules calls api.v0alpha.Acs.UpdateRules.
func (c *acsClient) UpdateRules(ctx context.Context, req *connect_go.Request[v0alpha.UpdateRulesRequest]) (*connect_go.Response[v0alpha.UpdateRulesResponse], error) {
	return c.updateRules.CallUnary(ctx, req)
}

// ListRules calls api.v0alpha.Acs.ListRules.
func (c *acsClient) ListRules(ctx context.Context, req *connect_go.Request[v0alpha.ListRulesRequest]) (*connect_go.Response[v0alpha.ListRulesResponse], error) {
	return c.listRules.CallUnary(ctx, req)
}

// CreateRegexEvaluator calls api.v0alpha.Acs.CreateRegexEvaluator.
func (c *acsClient) CreateRegexEvaluator(ctx context.Context, req *connect_go.Request[v0alpha.CreateRegexEvaluatorRequest]) (*connect_go.Response[v0alpha.CreateRegexEvaluatorResponse], error) {
	return c.createRegexEvaluator.CallUnary(ctx, req)
}

// UpdateRegexEvaluator calls api.v0alpha.Acs.UpdateRegexEvaluator.
func (c *acsClient) UpdateRegexEvaluator(ctx context.Context, req *connect_go.Request[v0alpha.UpdateRegexEvaluatorRequest]) (*connect_go.Response[v0alpha.UpdateRegexEvaluatorResponse], error) {
	return c.updateRegexEvaluator.CallUnary(ctx, req)
}

// ListRegexEvaluators calls api.v0alpha.Acs.ListRegexEvaluators.
func (c *acsClient) ListRegexEvaluators(ctx context.Context, req *connect_go.Request[v0alpha.ListRegexEvaluatorsRequest]) (*connect_go.Response[v0alpha.ListRegexEvaluatorsResponse], error) {
	return c.listRegexEvaluators.CallUnary(ctx, req)
}

// DeleteRule calls api.v0alpha.Acs.DeleteRule.
func (c *acsClient) DeleteRule(ctx context.Context, req *connect_go.Request[v0alpha.DeleteRuleRequest]) (*connect_go.Response[v0alpha.DeleteRuleResponse], error) {
	return c.deleteRule.CallUnary(ctx, req)
}

// GetScript calls api.v0alpha.Acs.GetScript.
func (c *acsClient) GetScript(ctx context.Context, req *connect_go.Request[v0alpha.GetScriptRequest]) (*connect_go.Response[v0alpha.GetScriptResponse], error) {
	return c.getScript.CallUnary(ctx, req)
}

// AcsHandler is an implementation of the api.v0alpha.Acs service.
type AcsHandler interface {
	CreateScript(context.Context, *connect_go.Request[v0alpha.CreateScriptRequest]) (*connect_go.Response[v0alpha.CreateScriptResponse], error)
	UpdateScript(context.Context, *connect_go.Request[v0alpha.UpdateScriptRequest]) (*connect_go.Response[v0alpha.UpdateScriptResponse], error)
	GetScriptById(context.Context, *connect_go.Request[v0alpha.GetScriptByIdRequest]) (*connect_go.Response[v0alpha.GetScriptByIdResponse], error)
	ListScripts(context.Context, *connect_go.Request[v0alpha.ListScriptsRequest], *connect_go.ServerStream[v0alpha.ScriptDetails]) error
	SaveResponses(context.Context, *connect_go.Request[v0alpha.SaveResponsesRequest]) (*connect_go.Response[v0alpha.SaveResponsesResponse], error)
	GetResponses(context.Context, *connect_go.Request[v0alpha.GetResponsesRequest]) (*connect_go.Response[v0alpha.GetResponsesResponse], error)
	CreateRule(context.Context, *connect_go.Request[v0alpha.CreateRuleRequest]) (*connect_go.Response[v0alpha.CreateRuleResponse], error)
	UpdateRules(context.Context, *connect_go.Request[v0alpha.UpdateRulesRequest]) (*connect_go.Response[v0alpha.UpdateRulesResponse], error)
	ListRules(context.Context, *connect_go.Request[v0alpha.ListRulesRequest]) (*connect_go.Response[v0alpha.ListRulesResponse], error)
	CreateRegexEvaluator(context.Context, *connect_go.Request[v0alpha.CreateRegexEvaluatorRequest]) (*connect_go.Response[v0alpha.CreateRegexEvaluatorResponse], error)
	UpdateRegexEvaluator(context.Context, *connect_go.Request[v0alpha.UpdateRegexEvaluatorRequest]) (*connect_go.Response[v0alpha.UpdateRegexEvaluatorResponse], error)
	ListRegexEvaluators(context.Context, *connect_go.Request[v0alpha.ListRegexEvaluatorsRequest]) (*connect_go.Response[v0alpha.ListRegexEvaluatorsResponse], error)
	DeleteRule(context.Context, *connect_go.Request[v0alpha.DeleteRuleRequest]) (*connect_go.Response[v0alpha.DeleteRuleResponse], error)
	GetScript(context.Context, *connect_go.Request[v0alpha.GetScriptRequest]) (*connect_go.Response[v0alpha.GetScriptResponse], error)
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
	return "/api.v0alpha.Acs/", mux
}

// UnimplementedAcsHandler returns CodeUnimplemented from all methods.
type UnimplementedAcsHandler struct{}

func (UnimplementedAcsHandler) CreateScript(context.Context, *connect_go.Request[v0alpha.CreateScriptRequest]) (*connect_go.Response[v0alpha.CreateScriptResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Acs.CreateScript is not implemented"))
}

func (UnimplementedAcsHandler) UpdateScript(context.Context, *connect_go.Request[v0alpha.UpdateScriptRequest]) (*connect_go.Response[v0alpha.UpdateScriptResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Acs.UpdateScript is not implemented"))
}

func (UnimplementedAcsHandler) GetScriptById(context.Context, *connect_go.Request[v0alpha.GetScriptByIdRequest]) (*connect_go.Response[v0alpha.GetScriptByIdResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Acs.GetScriptById is not implemented"))
}

func (UnimplementedAcsHandler) ListScripts(context.Context, *connect_go.Request[v0alpha.ListScriptsRequest], *connect_go.ServerStream[v0alpha.ScriptDetails]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Acs.ListScripts is not implemented"))
}

func (UnimplementedAcsHandler) SaveResponses(context.Context, *connect_go.Request[v0alpha.SaveResponsesRequest]) (*connect_go.Response[v0alpha.SaveResponsesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Acs.SaveResponses is not implemented"))
}

func (UnimplementedAcsHandler) GetResponses(context.Context, *connect_go.Request[v0alpha.GetResponsesRequest]) (*connect_go.Response[v0alpha.GetResponsesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Acs.GetResponses is not implemented"))
}

func (UnimplementedAcsHandler) CreateRule(context.Context, *connect_go.Request[v0alpha.CreateRuleRequest]) (*connect_go.Response[v0alpha.CreateRuleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Acs.CreateRule is not implemented"))
}

func (UnimplementedAcsHandler) UpdateRules(context.Context, *connect_go.Request[v0alpha.UpdateRulesRequest]) (*connect_go.Response[v0alpha.UpdateRulesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Acs.UpdateRules is not implemented"))
}

func (UnimplementedAcsHandler) ListRules(context.Context, *connect_go.Request[v0alpha.ListRulesRequest]) (*connect_go.Response[v0alpha.ListRulesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Acs.ListRules is not implemented"))
}

func (UnimplementedAcsHandler) CreateRegexEvaluator(context.Context, *connect_go.Request[v0alpha.CreateRegexEvaluatorRequest]) (*connect_go.Response[v0alpha.CreateRegexEvaluatorResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Acs.CreateRegexEvaluator is not implemented"))
}

func (UnimplementedAcsHandler) UpdateRegexEvaluator(context.Context, *connect_go.Request[v0alpha.UpdateRegexEvaluatorRequest]) (*connect_go.Response[v0alpha.UpdateRegexEvaluatorResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Acs.UpdateRegexEvaluator is not implemented"))
}

func (UnimplementedAcsHandler) ListRegexEvaluators(context.Context, *connect_go.Request[v0alpha.ListRegexEvaluatorsRequest]) (*connect_go.Response[v0alpha.ListRegexEvaluatorsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Acs.ListRegexEvaluators is not implemented"))
}

func (UnimplementedAcsHandler) DeleteRule(context.Context, *connect_go.Request[v0alpha.DeleteRuleRequest]) (*connect_go.Response[v0alpha.DeleteRuleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Acs.DeleteRule is not implemented"))
}

func (UnimplementedAcsHandler) GetScript(context.Context, *connect_go.Request[v0alpha.GetScriptRequest]) (*connect_go.Response[v0alpha.GetScriptResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v0alpha.Acs.GetScript is not implemented"))
}
