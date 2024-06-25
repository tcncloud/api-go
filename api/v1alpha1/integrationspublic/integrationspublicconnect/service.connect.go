// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1alpha1/integrationspublic/service.proto

package integrationspublicconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	integrationspublic "github.com/tcncloud/api-go/api/v1alpha1/integrationspublic"
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
	// IntegrationsPublicName is the fully-qualified name of the IntegrationsPublic service.
	IntegrationsPublicName = "api.v1alpha1.integrationspublic.IntegrationsPublic"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// IntegrationsPublicGetLinkDataProcedure is the fully-qualified name of the IntegrationsPublic's
	// GetLinkData RPC.
	IntegrationsPublicGetLinkDataProcedure = "/api.v1alpha1.integrationspublic.IntegrationsPublic/GetLinkData"
	// IntegrationsPublicSubmitVerificationProcedure is the fully-qualified name of the
	// IntegrationsPublic's SubmitVerification RPC.
	IntegrationsPublicSubmitVerificationProcedure = "/api.v1alpha1.integrationspublic.IntegrationsPublic/SubmitVerification"
	// IntegrationsPublicSessionKeepAliveProcedure is the fully-qualified name of the
	// IntegrationsPublic's SessionKeepAlive RPC.
	IntegrationsPublicSessionKeepAliveProcedure = "/api.v1alpha1.integrationspublic.IntegrationsPublic/SessionKeepAlive"
	// IntegrationsPublicGetInvoiceProcedure is the fully-qualified name of the IntegrationsPublic's
	// GetInvoice RPC.
	IntegrationsPublicGetInvoiceProcedure = "/api.v1alpha1.integrationspublic.IntegrationsPublic/GetInvoice"
	// IntegrationsPublicSubmitPaymentProcedure is the fully-qualified name of the IntegrationsPublic's
	// SubmitPayment RPC.
	IntegrationsPublicSubmitPaymentProcedure = "/api.v1alpha1.integrationspublic.IntegrationsPublic/SubmitPayment"
	// IntegrationsPublicGetReceiptProcedure is the fully-qualified name of the IntegrationsPublic's
	// GetReceipt RPC.
	IntegrationsPublicGetReceiptProcedure = "/api.v1alpha1.integrationspublic.IntegrationsPublic/GetReceipt"
)

// IntegrationsPublicClient is a client for the api.v1alpha1.integrationspublic.IntegrationsPublic
// service.
type IntegrationsPublicClient interface {
	GetLinkData(context.Context, *connect_go.Request[integrationspublic.GetLinkDataReq]) (*connect_go.Response[integrationspublic.GetLinkDataRes], error)
	SubmitVerification(context.Context, *connect_go.Request[integrationspublic.SubmitVerificationReq]) (*connect_go.Response[integrationspublic.SubmitVerificationRes], error)
	SessionKeepAlive(context.Context, *connect_go.Request[integrationspublic.SessionKeepAliveReq]) (*connect_go.Response[integrationspublic.SessionKeepAliveRes], error)
	GetInvoice(context.Context, *connect_go.Request[integrationspublic.GetInvoiceReq]) (*connect_go.Response[integrationspublic.GetInvoiceRes], error)
	SubmitPayment(context.Context, *connect_go.Request[integrationspublic.SubmitPaymentReq]) (*connect_go.Response[integrationspublic.SubmitPaymentRes], error)
	GetReceipt(context.Context, *connect_go.Request[integrationspublic.GetReceiptReq]) (*connect_go.Response[integrationspublic.GetReceiptRes], error)
}

// NewIntegrationsPublicClient constructs a client for the
// api.v1alpha1.integrationspublic.IntegrationsPublic service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewIntegrationsPublicClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) IntegrationsPublicClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &integrationsPublicClient{
		getLinkData: connect_go.NewClient[integrationspublic.GetLinkDataReq, integrationspublic.GetLinkDataRes](
			httpClient,
			baseURL+IntegrationsPublicGetLinkDataProcedure,
			opts...,
		),
		submitVerification: connect_go.NewClient[integrationspublic.SubmitVerificationReq, integrationspublic.SubmitVerificationRes](
			httpClient,
			baseURL+IntegrationsPublicSubmitVerificationProcedure,
			opts...,
		),
		sessionKeepAlive: connect_go.NewClient[integrationspublic.SessionKeepAliveReq, integrationspublic.SessionKeepAliveRes](
			httpClient,
			baseURL+IntegrationsPublicSessionKeepAliveProcedure,
			opts...,
		),
		getInvoice: connect_go.NewClient[integrationspublic.GetInvoiceReq, integrationspublic.GetInvoiceRes](
			httpClient,
			baseURL+IntegrationsPublicGetInvoiceProcedure,
			opts...,
		),
		submitPayment: connect_go.NewClient[integrationspublic.SubmitPaymentReq, integrationspublic.SubmitPaymentRes](
			httpClient,
			baseURL+IntegrationsPublicSubmitPaymentProcedure,
			opts...,
		),
		getReceipt: connect_go.NewClient[integrationspublic.GetReceiptReq, integrationspublic.GetReceiptRes](
			httpClient,
			baseURL+IntegrationsPublicGetReceiptProcedure,
			opts...,
		),
	}
}

// integrationsPublicClient implements IntegrationsPublicClient.
type integrationsPublicClient struct {
	getLinkData        *connect_go.Client[integrationspublic.GetLinkDataReq, integrationspublic.GetLinkDataRes]
	submitVerification *connect_go.Client[integrationspublic.SubmitVerificationReq, integrationspublic.SubmitVerificationRes]
	sessionKeepAlive   *connect_go.Client[integrationspublic.SessionKeepAliveReq, integrationspublic.SessionKeepAliveRes]
	getInvoice         *connect_go.Client[integrationspublic.GetInvoiceReq, integrationspublic.GetInvoiceRes]
	submitPayment      *connect_go.Client[integrationspublic.SubmitPaymentReq, integrationspublic.SubmitPaymentRes]
	getReceipt         *connect_go.Client[integrationspublic.GetReceiptReq, integrationspublic.GetReceiptRes]
}

// GetLinkData calls api.v1alpha1.integrationspublic.IntegrationsPublic.GetLinkData.
func (c *integrationsPublicClient) GetLinkData(ctx context.Context, req *connect_go.Request[integrationspublic.GetLinkDataReq]) (*connect_go.Response[integrationspublic.GetLinkDataRes], error) {
	return c.getLinkData.CallUnary(ctx, req)
}

// SubmitVerification calls api.v1alpha1.integrationspublic.IntegrationsPublic.SubmitVerification.
func (c *integrationsPublicClient) SubmitVerification(ctx context.Context, req *connect_go.Request[integrationspublic.SubmitVerificationReq]) (*connect_go.Response[integrationspublic.SubmitVerificationRes], error) {
	return c.submitVerification.CallUnary(ctx, req)
}

// SessionKeepAlive calls api.v1alpha1.integrationspublic.IntegrationsPublic.SessionKeepAlive.
func (c *integrationsPublicClient) SessionKeepAlive(ctx context.Context, req *connect_go.Request[integrationspublic.SessionKeepAliveReq]) (*connect_go.Response[integrationspublic.SessionKeepAliveRes], error) {
	return c.sessionKeepAlive.CallUnary(ctx, req)
}

// GetInvoice calls api.v1alpha1.integrationspublic.IntegrationsPublic.GetInvoice.
func (c *integrationsPublicClient) GetInvoice(ctx context.Context, req *connect_go.Request[integrationspublic.GetInvoiceReq]) (*connect_go.Response[integrationspublic.GetInvoiceRes], error) {
	return c.getInvoice.CallUnary(ctx, req)
}

// SubmitPayment calls api.v1alpha1.integrationspublic.IntegrationsPublic.SubmitPayment.
func (c *integrationsPublicClient) SubmitPayment(ctx context.Context, req *connect_go.Request[integrationspublic.SubmitPaymentReq]) (*connect_go.Response[integrationspublic.SubmitPaymentRes], error) {
	return c.submitPayment.CallUnary(ctx, req)
}

// GetReceipt calls api.v1alpha1.integrationspublic.IntegrationsPublic.GetReceipt.
func (c *integrationsPublicClient) GetReceipt(ctx context.Context, req *connect_go.Request[integrationspublic.GetReceiptReq]) (*connect_go.Response[integrationspublic.GetReceiptRes], error) {
	return c.getReceipt.CallUnary(ctx, req)
}

// IntegrationsPublicHandler is an implementation of the
// api.v1alpha1.integrationspublic.IntegrationsPublic service.
type IntegrationsPublicHandler interface {
	GetLinkData(context.Context, *connect_go.Request[integrationspublic.GetLinkDataReq]) (*connect_go.Response[integrationspublic.GetLinkDataRes], error)
	SubmitVerification(context.Context, *connect_go.Request[integrationspublic.SubmitVerificationReq]) (*connect_go.Response[integrationspublic.SubmitVerificationRes], error)
	SessionKeepAlive(context.Context, *connect_go.Request[integrationspublic.SessionKeepAliveReq]) (*connect_go.Response[integrationspublic.SessionKeepAliveRes], error)
	GetInvoice(context.Context, *connect_go.Request[integrationspublic.GetInvoiceReq]) (*connect_go.Response[integrationspublic.GetInvoiceRes], error)
	SubmitPayment(context.Context, *connect_go.Request[integrationspublic.SubmitPaymentReq]) (*connect_go.Response[integrationspublic.SubmitPaymentRes], error)
	GetReceipt(context.Context, *connect_go.Request[integrationspublic.GetReceiptReq]) (*connect_go.Response[integrationspublic.GetReceiptRes], error)
}

// NewIntegrationsPublicHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewIntegrationsPublicHandler(svc IntegrationsPublicHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	integrationsPublicGetLinkDataHandler := connect_go.NewUnaryHandler(
		IntegrationsPublicGetLinkDataProcedure,
		svc.GetLinkData,
		opts...,
	)
	integrationsPublicSubmitVerificationHandler := connect_go.NewUnaryHandler(
		IntegrationsPublicSubmitVerificationProcedure,
		svc.SubmitVerification,
		opts...,
	)
	integrationsPublicSessionKeepAliveHandler := connect_go.NewUnaryHandler(
		IntegrationsPublicSessionKeepAliveProcedure,
		svc.SessionKeepAlive,
		opts...,
	)
	integrationsPublicGetInvoiceHandler := connect_go.NewUnaryHandler(
		IntegrationsPublicGetInvoiceProcedure,
		svc.GetInvoice,
		opts...,
	)
	integrationsPublicSubmitPaymentHandler := connect_go.NewUnaryHandler(
		IntegrationsPublicSubmitPaymentProcedure,
		svc.SubmitPayment,
		opts...,
	)
	integrationsPublicGetReceiptHandler := connect_go.NewUnaryHandler(
		IntegrationsPublicGetReceiptProcedure,
		svc.GetReceipt,
		opts...,
	)
	return "/api.v1alpha1.integrationspublic.IntegrationsPublic/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case IntegrationsPublicGetLinkDataProcedure:
			integrationsPublicGetLinkDataHandler.ServeHTTP(w, r)
		case IntegrationsPublicSubmitVerificationProcedure:
			integrationsPublicSubmitVerificationHandler.ServeHTTP(w, r)
		case IntegrationsPublicSessionKeepAliveProcedure:
			integrationsPublicSessionKeepAliveHandler.ServeHTTP(w, r)
		case IntegrationsPublicGetInvoiceProcedure:
			integrationsPublicGetInvoiceHandler.ServeHTTP(w, r)
		case IntegrationsPublicSubmitPaymentProcedure:
			integrationsPublicSubmitPaymentHandler.ServeHTTP(w, r)
		case IntegrationsPublicGetReceiptProcedure:
			integrationsPublicGetReceiptHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedIntegrationsPublicHandler returns CodeUnimplemented from all methods.
type UnimplementedIntegrationsPublicHandler struct{}

func (UnimplementedIntegrationsPublicHandler) GetLinkData(context.Context, *connect_go.Request[integrationspublic.GetLinkDataReq]) (*connect_go.Response[integrationspublic.GetLinkDataRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.integrationspublic.IntegrationsPublic.GetLinkData is not implemented"))
}

func (UnimplementedIntegrationsPublicHandler) SubmitVerification(context.Context, *connect_go.Request[integrationspublic.SubmitVerificationReq]) (*connect_go.Response[integrationspublic.SubmitVerificationRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.integrationspublic.IntegrationsPublic.SubmitVerification is not implemented"))
}

func (UnimplementedIntegrationsPublicHandler) SessionKeepAlive(context.Context, *connect_go.Request[integrationspublic.SessionKeepAliveReq]) (*connect_go.Response[integrationspublic.SessionKeepAliveRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.integrationspublic.IntegrationsPublic.SessionKeepAlive is not implemented"))
}

func (UnimplementedIntegrationsPublicHandler) GetInvoice(context.Context, *connect_go.Request[integrationspublic.GetInvoiceReq]) (*connect_go.Response[integrationspublic.GetInvoiceRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.integrationspublic.IntegrationsPublic.GetInvoice is not implemented"))
}

func (UnimplementedIntegrationsPublicHandler) SubmitPayment(context.Context, *connect_go.Request[integrationspublic.SubmitPaymentReq]) (*connect_go.Response[integrationspublic.SubmitPaymentRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.integrationspublic.IntegrationsPublic.SubmitPayment is not implemented"))
}

func (UnimplementedIntegrationsPublicHandler) GetReceipt(context.Context, *connect_go.Request[integrationspublic.GetReceiptReq]) (*connect_go.Response[integrationspublic.GetReceiptRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.integrationspublic.IntegrationsPublic.GetReceipt is not implemented"))
}
