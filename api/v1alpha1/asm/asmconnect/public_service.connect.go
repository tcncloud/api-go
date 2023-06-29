// Copyright (c) 2020, TCN Inc.
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
// Source: api/v1alpha1/asm/public_service.proto

package asmconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	commons "github.com/tcncloud/api-go/api/commons"
	asm "github.com/tcncloud/api-go/api/v1alpha1/asm"
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
	// AsmName is the fully-qualified name of the Asm service.
	AsmName = "api.v1alpha1.asm.Asm"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AsmStreamAgentStateProcedure is the fully-qualified name of the Asm's StreamAgentState RPC.
	AsmStreamAgentStateProcedure = "/api.v1alpha1.asm.Asm/StreamAgentState"
	// AsmManagerStreamAgentStateProcedure is the fully-qualified name of the Asm's
	// ManagerStreamAgentState RPC.
	AsmManagerStreamAgentStateProcedure = "/api.v1alpha1.asm.Asm/ManagerStreamAgentState"
	// AsmPushEventsProcedure is the fully-qualified name of the Asm's PushEvents RPC.
	AsmPushEventsProcedure = "/api.v1alpha1.asm.Asm/PushEvents"
	// AsmCreateSessionProcedure is the fully-qualified name of the Asm's CreateSession RPC.
	AsmCreateSessionProcedure = "/api.v1alpha1.asm.Asm/CreateSession"
	// AsmEndSessionProcedure is the fully-qualified name of the Asm's EndSession RPC.
	AsmEndSessionProcedure = "/api.v1alpha1.asm.Asm/EndSession"
	// AsmGetCurrentSessionProcedure is the fully-qualified name of the Asm's GetCurrentSession RPC.
	AsmGetCurrentSessionProcedure = "/api.v1alpha1.asm.Asm/GetCurrentSession"
	// AsmEnableVoiceProcedure is the fully-qualified name of the Asm's EnableVoice RPC.
	AsmEnableVoiceProcedure = "/api.v1alpha1.asm.Asm/EnableVoice"
	// AsmDisableVoiceProcedure is the fully-qualified name of the Asm's DisableVoice RPC.
	AsmDisableVoiceProcedure = "/api.v1alpha1.asm.Asm/DisableVoice"
	// AsmListConversationsProcedure is the fully-qualified name of the Asm's ListConversations RPC.
	AsmListConversationsProcedure = "/api.v1alpha1.asm.Asm/ListConversations"
	// AsmAssignNewConversationProcedure is the fully-qualified name of the Asm's AssignNewConversation
	// RPC.
	AsmAssignNewConversationProcedure = "/api.v1alpha1.asm.Asm/AssignNewConversation"
	// AsmListAgentsProcedure is the fully-qualified name of the Asm's ListAgents RPC.
	AsmListAgentsProcedure = "/api.v1alpha1.asm.Asm/ListAgents"
	// AsmSetConversationCollectedDataProcedure is the fully-qualified name of the Asm's
	// SetConversationCollectedData RPC.
	AsmSetConversationCollectedDataProcedure = "/api.v1alpha1.asm.Asm/SetConversationCollectedData"
	// AsmGetQueuesDetailsProcedure is the fully-qualified name of the Asm's GetQueuesDetails RPC.
	AsmGetQueuesDetailsProcedure = "/api.v1alpha1.asm.Asm/GetQueuesDetails"
)

// AsmClient is a client for the api.v1alpha1.asm.Asm service.
type AsmClient interface {
	// Streams back status updates for the given asm session
	// only the asm session sid filter is allowed
	StreamAgentState(context.Context, *connect_go.Request[asm.StreamAgentStateReq]) (*connect_go.ServerStreamForClient[commons.StreamAgentStateRes], error)
	// Streams back statuses for the desired filter
	ManagerStreamAgentState(context.Context, *connect_go.Request[asm.ManagerStreamAgentStateReq]) (*connect_go.ServerStreamForClient[commons.ManagerStreamAgentStateRes], error)
	PushEvents(context.Context, *connect_go.Request[asm.PushEventsReq]) (*connect_go.Response[asm.PushEventsRes], error)
	// Creates an agent session and enables the voice channel
	CreateSession(context.Context, *connect_go.Request[asm.CreateSessionReq]) (*connect_go.Response[asm.CreateSessionRes], error)
	// Closes an asm session and all sub sessions
	EndSession(context.Context, *connect_go.Request[asm.EndSessionReq]) (*connect_go.Response[asm.EndSessionRes], error)
	// Gets an agent's current asm session
	GetCurrentSession(context.Context, *connect_go.Request[asm.GetCurrentSessionReq]) (*connect_go.Response[asm.AsmSession], error)
	// Updates the currently active subsession
	EnableVoice(context.Context, *connect_go.Request[asm.EnableVoiceReq]) (*connect_go.Response[asm.EnableVoiceRes], error)
	DisableVoice(context.Context, *connect_go.Request[asm.DisableVoiceReq]) (*connect_go.Response[asm.DisableVoiceRes], error)
	// Lists the conversations for an assigned agent
	ListConversations(context.Context, *connect_go.Request[asm.ListConversationsReq]) (*connect_go.Response[asm.ListConversationsRes], error)
	// Assign agent to matched conversation based on skills and channelTypes requested
	AssignNewConversation(context.Context, *connect_go.Request[asm.AssignNewConversationReq]) (*connect_go.Response[asm.AssignNewConversationRes], error)
	// List all agents for the given user. Contains statistical enrichments for each agent and their conversations.
	ListAgents(context.Context, *connect_go.Request[asm.ListAgentsReq]) (*connect_go.Response[asm.ListAgentsRes], error)
	// Set collected data per conversation
	SetConversationCollectedData(context.Context, *connect_go.Request[asm.SetConversationCollectedDataReq]) (*connect_go.Response[asm.SetConversationCollectedDataRes], error)
	// Set queue details
	GetQueuesDetails(context.Context, *connect_go.Request[asm.GetQueuesDetailsReq]) (*connect_go.Response[commons.GetQueuesDetailsRes], error)
}

// NewAsmClient constructs a client for the api.v1alpha1.asm.Asm service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAsmClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AsmClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &asmClient{
		streamAgentState: connect_go.NewClient[asm.StreamAgentStateReq, commons.StreamAgentStateRes](
			httpClient,
			baseURL+AsmStreamAgentStateProcedure,
			opts...,
		),
		managerStreamAgentState: connect_go.NewClient[asm.ManagerStreamAgentStateReq, commons.ManagerStreamAgentStateRes](
			httpClient,
			baseURL+AsmManagerStreamAgentStateProcedure,
			opts...,
		),
		pushEvents: connect_go.NewClient[asm.PushEventsReq, asm.PushEventsRes](
			httpClient,
			baseURL+AsmPushEventsProcedure,
			opts...,
		),
		createSession: connect_go.NewClient[asm.CreateSessionReq, asm.CreateSessionRes](
			httpClient,
			baseURL+AsmCreateSessionProcedure,
			opts...,
		),
		endSession: connect_go.NewClient[asm.EndSessionReq, asm.EndSessionRes](
			httpClient,
			baseURL+AsmEndSessionProcedure,
			opts...,
		),
		getCurrentSession: connect_go.NewClient[asm.GetCurrentSessionReq, asm.AsmSession](
			httpClient,
			baseURL+AsmGetCurrentSessionProcedure,
			opts...,
		),
		enableVoice: connect_go.NewClient[asm.EnableVoiceReq, asm.EnableVoiceRes](
			httpClient,
			baseURL+AsmEnableVoiceProcedure,
			opts...,
		),
		disableVoice: connect_go.NewClient[asm.DisableVoiceReq, asm.DisableVoiceRes](
			httpClient,
			baseURL+AsmDisableVoiceProcedure,
			opts...,
		),
		listConversations: connect_go.NewClient[asm.ListConversationsReq, asm.ListConversationsRes](
			httpClient,
			baseURL+AsmListConversationsProcedure,
			opts...,
		),
		assignNewConversation: connect_go.NewClient[asm.AssignNewConversationReq, asm.AssignNewConversationRes](
			httpClient,
			baseURL+AsmAssignNewConversationProcedure,
			opts...,
		),
		listAgents: connect_go.NewClient[asm.ListAgentsReq, asm.ListAgentsRes](
			httpClient,
			baseURL+AsmListAgentsProcedure,
			opts...,
		),
		setConversationCollectedData: connect_go.NewClient[asm.SetConversationCollectedDataReq, asm.SetConversationCollectedDataRes](
			httpClient,
			baseURL+AsmSetConversationCollectedDataProcedure,
			opts...,
		),
		getQueuesDetails: connect_go.NewClient[asm.GetQueuesDetailsReq, commons.GetQueuesDetailsRes](
			httpClient,
			baseURL+AsmGetQueuesDetailsProcedure,
			opts...,
		),
	}
}

// asmClient implements AsmClient.
type asmClient struct {
	streamAgentState             *connect_go.Client[asm.StreamAgentStateReq, commons.StreamAgentStateRes]
	managerStreamAgentState      *connect_go.Client[asm.ManagerStreamAgentStateReq, commons.ManagerStreamAgentStateRes]
	pushEvents                   *connect_go.Client[asm.PushEventsReq, asm.PushEventsRes]
	createSession                *connect_go.Client[asm.CreateSessionReq, asm.CreateSessionRes]
	endSession                   *connect_go.Client[asm.EndSessionReq, asm.EndSessionRes]
	getCurrentSession            *connect_go.Client[asm.GetCurrentSessionReq, asm.AsmSession]
	enableVoice                  *connect_go.Client[asm.EnableVoiceReq, asm.EnableVoiceRes]
	disableVoice                 *connect_go.Client[asm.DisableVoiceReq, asm.DisableVoiceRes]
	listConversations            *connect_go.Client[asm.ListConversationsReq, asm.ListConversationsRes]
	assignNewConversation        *connect_go.Client[asm.AssignNewConversationReq, asm.AssignNewConversationRes]
	listAgents                   *connect_go.Client[asm.ListAgentsReq, asm.ListAgentsRes]
	setConversationCollectedData *connect_go.Client[asm.SetConversationCollectedDataReq, asm.SetConversationCollectedDataRes]
	getQueuesDetails             *connect_go.Client[asm.GetQueuesDetailsReq, commons.GetQueuesDetailsRes]
}

// StreamAgentState calls api.v1alpha1.asm.Asm.StreamAgentState.
func (c *asmClient) StreamAgentState(ctx context.Context, req *connect_go.Request[asm.StreamAgentStateReq]) (*connect_go.ServerStreamForClient[commons.StreamAgentStateRes], error) {
	return c.streamAgentState.CallServerStream(ctx, req)
}

// ManagerStreamAgentState calls api.v1alpha1.asm.Asm.ManagerStreamAgentState.
func (c *asmClient) ManagerStreamAgentState(ctx context.Context, req *connect_go.Request[asm.ManagerStreamAgentStateReq]) (*connect_go.ServerStreamForClient[commons.ManagerStreamAgentStateRes], error) {
	return c.managerStreamAgentState.CallServerStream(ctx, req)
}

// PushEvents calls api.v1alpha1.asm.Asm.PushEvents.
func (c *asmClient) PushEvents(ctx context.Context, req *connect_go.Request[asm.PushEventsReq]) (*connect_go.Response[asm.PushEventsRes], error) {
	return c.pushEvents.CallUnary(ctx, req)
}

// CreateSession calls api.v1alpha1.asm.Asm.CreateSession.
func (c *asmClient) CreateSession(ctx context.Context, req *connect_go.Request[asm.CreateSessionReq]) (*connect_go.Response[asm.CreateSessionRes], error) {
	return c.createSession.CallUnary(ctx, req)
}

// EndSession calls api.v1alpha1.asm.Asm.EndSession.
func (c *asmClient) EndSession(ctx context.Context, req *connect_go.Request[asm.EndSessionReq]) (*connect_go.Response[asm.EndSessionRes], error) {
	return c.endSession.CallUnary(ctx, req)
}

// GetCurrentSession calls api.v1alpha1.asm.Asm.GetCurrentSession.
func (c *asmClient) GetCurrentSession(ctx context.Context, req *connect_go.Request[asm.GetCurrentSessionReq]) (*connect_go.Response[asm.AsmSession], error) {
	return c.getCurrentSession.CallUnary(ctx, req)
}

// EnableVoice calls api.v1alpha1.asm.Asm.EnableVoice.
func (c *asmClient) EnableVoice(ctx context.Context, req *connect_go.Request[asm.EnableVoiceReq]) (*connect_go.Response[asm.EnableVoiceRes], error) {
	return c.enableVoice.CallUnary(ctx, req)
}

// DisableVoice calls api.v1alpha1.asm.Asm.DisableVoice.
func (c *asmClient) DisableVoice(ctx context.Context, req *connect_go.Request[asm.DisableVoiceReq]) (*connect_go.Response[asm.DisableVoiceRes], error) {
	return c.disableVoice.CallUnary(ctx, req)
}

// ListConversations calls api.v1alpha1.asm.Asm.ListConversations.
func (c *asmClient) ListConversations(ctx context.Context, req *connect_go.Request[asm.ListConversationsReq]) (*connect_go.Response[asm.ListConversationsRes], error) {
	return c.listConversations.CallUnary(ctx, req)
}

// AssignNewConversation calls api.v1alpha1.asm.Asm.AssignNewConversation.
func (c *asmClient) AssignNewConversation(ctx context.Context, req *connect_go.Request[asm.AssignNewConversationReq]) (*connect_go.Response[asm.AssignNewConversationRes], error) {
	return c.assignNewConversation.CallUnary(ctx, req)
}

// ListAgents calls api.v1alpha1.asm.Asm.ListAgents.
func (c *asmClient) ListAgents(ctx context.Context, req *connect_go.Request[asm.ListAgentsReq]) (*connect_go.Response[asm.ListAgentsRes], error) {
	return c.listAgents.CallUnary(ctx, req)
}

// SetConversationCollectedData calls api.v1alpha1.asm.Asm.SetConversationCollectedData.
func (c *asmClient) SetConversationCollectedData(ctx context.Context, req *connect_go.Request[asm.SetConversationCollectedDataReq]) (*connect_go.Response[asm.SetConversationCollectedDataRes], error) {
	return c.setConversationCollectedData.CallUnary(ctx, req)
}

// GetQueuesDetails calls api.v1alpha1.asm.Asm.GetQueuesDetails.
func (c *asmClient) GetQueuesDetails(ctx context.Context, req *connect_go.Request[asm.GetQueuesDetailsReq]) (*connect_go.Response[commons.GetQueuesDetailsRes], error) {
	return c.getQueuesDetails.CallUnary(ctx, req)
}

// AsmHandler is an implementation of the api.v1alpha1.asm.Asm service.
type AsmHandler interface {
	// Streams back status updates for the given asm session
	// only the asm session sid filter is allowed
	StreamAgentState(context.Context, *connect_go.Request[asm.StreamAgentStateReq], *connect_go.ServerStream[commons.StreamAgentStateRes]) error
	// Streams back statuses for the desired filter
	ManagerStreamAgentState(context.Context, *connect_go.Request[asm.ManagerStreamAgentStateReq], *connect_go.ServerStream[commons.ManagerStreamAgentStateRes]) error
	PushEvents(context.Context, *connect_go.Request[asm.PushEventsReq]) (*connect_go.Response[asm.PushEventsRes], error)
	// Creates an agent session and enables the voice channel
	CreateSession(context.Context, *connect_go.Request[asm.CreateSessionReq]) (*connect_go.Response[asm.CreateSessionRes], error)
	// Closes an asm session and all sub sessions
	EndSession(context.Context, *connect_go.Request[asm.EndSessionReq]) (*connect_go.Response[asm.EndSessionRes], error)
	// Gets an agent's current asm session
	GetCurrentSession(context.Context, *connect_go.Request[asm.GetCurrentSessionReq]) (*connect_go.Response[asm.AsmSession], error)
	// Updates the currently active subsession
	EnableVoice(context.Context, *connect_go.Request[asm.EnableVoiceReq]) (*connect_go.Response[asm.EnableVoiceRes], error)
	DisableVoice(context.Context, *connect_go.Request[asm.DisableVoiceReq]) (*connect_go.Response[asm.DisableVoiceRes], error)
	// Lists the conversations for an assigned agent
	ListConversations(context.Context, *connect_go.Request[asm.ListConversationsReq]) (*connect_go.Response[asm.ListConversationsRes], error)
	// Assign agent to matched conversation based on skills and channelTypes requested
	AssignNewConversation(context.Context, *connect_go.Request[asm.AssignNewConversationReq]) (*connect_go.Response[asm.AssignNewConversationRes], error)
	// List all agents for the given user. Contains statistical enrichments for each agent and their conversations.
	ListAgents(context.Context, *connect_go.Request[asm.ListAgentsReq]) (*connect_go.Response[asm.ListAgentsRes], error)
	// Set collected data per conversation
	SetConversationCollectedData(context.Context, *connect_go.Request[asm.SetConversationCollectedDataReq]) (*connect_go.Response[asm.SetConversationCollectedDataRes], error)
	// Set queue details
	GetQueuesDetails(context.Context, *connect_go.Request[asm.GetQueuesDetailsReq]) (*connect_go.Response[commons.GetQueuesDetailsRes], error)
}

// NewAsmHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAsmHandler(svc AsmHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	asmStreamAgentStateHandler := connect_go.NewServerStreamHandler(
		AsmStreamAgentStateProcedure,
		svc.StreamAgentState,
		opts...,
	)
	asmManagerStreamAgentStateHandler := connect_go.NewServerStreamHandler(
		AsmManagerStreamAgentStateProcedure,
		svc.ManagerStreamAgentState,
		opts...,
	)
	asmPushEventsHandler := connect_go.NewUnaryHandler(
		AsmPushEventsProcedure,
		svc.PushEvents,
		opts...,
	)
	asmCreateSessionHandler := connect_go.NewUnaryHandler(
		AsmCreateSessionProcedure,
		svc.CreateSession,
		opts...,
	)
	asmEndSessionHandler := connect_go.NewUnaryHandler(
		AsmEndSessionProcedure,
		svc.EndSession,
		opts...,
	)
	asmGetCurrentSessionHandler := connect_go.NewUnaryHandler(
		AsmGetCurrentSessionProcedure,
		svc.GetCurrentSession,
		opts...,
	)
	asmEnableVoiceHandler := connect_go.NewUnaryHandler(
		AsmEnableVoiceProcedure,
		svc.EnableVoice,
		opts...,
	)
	asmDisableVoiceHandler := connect_go.NewUnaryHandler(
		AsmDisableVoiceProcedure,
		svc.DisableVoice,
		opts...,
	)
	asmListConversationsHandler := connect_go.NewUnaryHandler(
		AsmListConversationsProcedure,
		svc.ListConversations,
		opts...,
	)
	asmAssignNewConversationHandler := connect_go.NewUnaryHandler(
		AsmAssignNewConversationProcedure,
		svc.AssignNewConversation,
		opts...,
	)
	asmListAgentsHandler := connect_go.NewUnaryHandler(
		AsmListAgentsProcedure,
		svc.ListAgents,
		opts...,
	)
	asmSetConversationCollectedDataHandler := connect_go.NewUnaryHandler(
		AsmSetConversationCollectedDataProcedure,
		svc.SetConversationCollectedData,
		opts...,
	)
	asmGetQueuesDetailsHandler := connect_go.NewUnaryHandler(
		AsmGetQueuesDetailsProcedure,
		svc.GetQueuesDetails,
		opts...,
	)
	return "/api.v1alpha1.asm.Asm/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AsmStreamAgentStateProcedure:
			asmStreamAgentStateHandler.ServeHTTP(w, r)
		case AsmManagerStreamAgentStateProcedure:
			asmManagerStreamAgentStateHandler.ServeHTTP(w, r)
		case AsmPushEventsProcedure:
			asmPushEventsHandler.ServeHTTP(w, r)
		case AsmCreateSessionProcedure:
			asmCreateSessionHandler.ServeHTTP(w, r)
		case AsmEndSessionProcedure:
			asmEndSessionHandler.ServeHTTP(w, r)
		case AsmGetCurrentSessionProcedure:
			asmGetCurrentSessionHandler.ServeHTTP(w, r)
		case AsmEnableVoiceProcedure:
			asmEnableVoiceHandler.ServeHTTP(w, r)
		case AsmDisableVoiceProcedure:
			asmDisableVoiceHandler.ServeHTTP(w, r)
		case AsmListConversationsProcedure:
			asmListConversationsHandler.ServeHTTP(w, r)
		case AsmAssignNewConversationProcedure:
			asmAssignNewConversationHandler.ServeHTTP(w, r)
		case AsmListAgentsProcedure:
			asmListAgentsHandler.ServeHTTP(w, r)
		case AsmSetConversationCollectedDataProcedure:
			asmSetConversationCollectedDataHandler.ServeHTTP(w, r)
		case AsmGetQueuesDetailsProcedure:
			asmGetQueuesDetailsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAsmHandler returns CodeUnimplemented from all methods.
type UnimplementedAsmHandler struct{}

func (UnimplementedAsmHandler) StreamAgentState(context.Context, *connect_go.Request[asm.StreamAgentStateReq], *connect_go.ServerStream[commons.StreamAgentStateRes]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.asm.Asm.StreamAgentState is not implemented"))
}

func (UnimplementedAsmHandler) ManagerStreamAgentState(context.Context, *connect_go.Request[asm.ManagerStreamAgentStateReq], *connect_go.ServerStream[commons.ManagerStreamAgentStateRes]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.asm.Asm.ManagerStreamAgentState is not implemented"))
}

func (UnimplementedAsmHandler) PushEvents(context.Context, *connect_go.Request[asm.PushEventsReq]) (*connect_go.Response[asm.PushEventsRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.asm.Asm.PushEvents is not implemented"))
}

func (UnimplementedAsmHandler) CreateSession(context.Context, *connect_go.Request[asm.CreateSessionReq]) (*connect_go.Response[asm.CreateSessionRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.asm.Asm.CreateSession is not implemented"))
}

func (UnimplementedAsmHandler) EndSession(context.Context, *connect_go.Request[asm.EndSessionReq]) (*connect_go.Response[asm.EndSessionRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.asm.Asm.EndSession is not implemented"))
}

func (UnimplementedAsmHandler) GetCurrentSession(context.Context, *connect_go.Request[asm.GetCurrentSessionReq]) (*connect_go.Response[asm.AsmSession], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.asm.Asm.GetCurrentSession is not implemented"))
}

func (UnimplementedAsmHandler) EnableVoice(context.Context, *connect_go.Request[asm.EnableVoiceReq]) (*connect_go.Response[asm.EnableVoiceRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.asm.Asm.EnableVoice is not implemented"))
}

func (UnimplementedAsmHandler) DisableVoice(context.Context, *connect_go.Request[asm.DisableVoiceReq]) (*connect_go.Response[asm.DisableVoiceRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.asm.Asm.DisableVoice is not implemented"))
}

func (UnimplementedAsmHandler) ListConversations(context.Context, *connect_go.Request[asm.ListConversationsReq]) (*connect_go.Response[asm.ListConversationsRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.asm.Asm.ListConversations is not implemented"))
}

func (UnimplementedAsmHandler) AssignNewConversation(context.Context, *connect_go.Request[asm.AssignNewConversationReq]) (*connect_go.Response[asm.AssignNewConversationRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.asm.Asm.AssignNewConversation is not implemented"))
}

func (UnimplementedAsmHandler) ListAgents(context.Context, *connect_go.Request[asm.ListAgentsReq]) (*connect_go.Response[asm.ListAgentsRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.asm.Asm.ListAgents is not implemented"))
}

func (UnimplementedAsmHandler) SetConversationCollectedData(context.Context, *connect_go.Request[asm.SetConversationCollectedDataReq]) (*connect_go.Response[asm.SetConversationCollectedDataRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.asm.Asm.SetConversationCollectedData is not implemented"))
}

func (UnimplementedAsmHandler) GetQueuesDetails(context.Context, *connect_go.Request[asm.GetQueuesDetailsReq]) (*connect_go.Response[commons.GetQueuesDetailsRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.asm.Asm.GetQueuesDetails is not implemented"))
}
