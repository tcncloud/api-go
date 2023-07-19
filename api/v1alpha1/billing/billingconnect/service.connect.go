// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1alpha1/billing/service.proto

package billingconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	billing "github.com/tcncloud/api-go/api/v1alpha1/billing"
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
	// BillingName is the fully-qualified name of the Billing service.
	BillingName = "api.v1alpha1.billing.Billing"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// BillingGetBillingPlanProcedure is the fully-qualified name of the Billing's GetBillingPlan RPC.
	BillingGetBillingPlanProcedure = "/api.v1alpha1.billing.Billing/GetBillingPlan"
	// BillingUpdateBillingPlanProcedure is the fully-qualified name of the Billing's UpdateBillingPlan
	// RPC.
	BillingUpdateBillingPlanProcedure = "/api.v1alpha1.billing.Billing/UpdateBillingPlan"
	// BillingGetInvoiceProcedure is the fully-qualified name of the Billing's GetInvoice RPC.
	BillingGetInvoiceProcedure = "/api.v1alpha1.billing.Billing/GetInvoice"
	// BillingExportGeneratedInvoiceProcedure is the fully-qualified name of the Billing's
	// ExportGeneratedInvoice RPC.
	BillingExportGeneratedInvoiceProcedure = "/api.v1alpha1.billing.Billing/ExportGeneratedInvoice"
)

// BillingClient is a client for the api.v1alpha1.billing.Billing service.
type BillingClient interface {
	// GetBillingPlan - returns the billing plan for the provided organization.
	GetBillingPlan(context.Context, *connect_go.Request[billing.GetBillingPlanReq]) (*connect_go.Response[billing.GetBillingPlanRes], error)
	// UpdateBillingPlan - updates the provided billing plan and it's details.
	// If some details are not provided, they will be left as is. However, if
	// deletion is desired, the DeleteBillingDetails method should be used. The
	// billing plan still follows the constraint of only having one billing detail
	// with a specific config type and event type, and so if the request contains
	// more than one billing detail with a config type and event type, the request
	// is malformed and will result in potentially unexpected behavior.
	UpdateBillingPlan(context.Context, *connect_go.Request[billing.UpdateBillingPlanReq]) (*connect_go.Response[billing.UpdateBillingPlanRes], error)
	// GetInvoice - returns the invoice for the organization. If a date is
	// provided, this will return the invoice for the organization that
	// corresponds to the billing cycle that contains the provided date. If
	// no date is provided, this will return the invoice as it currently
	// stands for the current billing cycle.
	GetInvoice(context.Context, *connect_go.Request[billing.GetInvoiceReq]) (*connect_go.Response[billing.GetInvoiceRes], error)
	ExportGeneratedInvoice(context.Context, *connect_go.Request[billing.ExportGeneratedInvoiceReq]) (*connect_go.Response[billing.ExportGeneratedInvoiceRes], error)
}

// NewBillingClient constructs a client for the api.v1alpha1.billing.Billing service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBillingClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) BillingClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &billingClient{
		getBillingPlan: connect_go.NewClient[billing.GetBillingPlanReq, billing.GetBillingPlanRes](
			httpClient,
			baseURL+BillingGetBillingPlanProcedure,
			opts...,
		),
		updateBillingPlan: connect_go.NewClient[billing.UpdateBillingPlanReq, billing.UpdateBillingPlanRes](
			httpClient,
			baseURL+BillingUpdateBillingPlanProcedure,
			opts...,
		),
		getInvoice: connect_go.NewClient[billing.GetInvoiceReq, billing.GetInvoiceRes](
			httpClient,
			baseURL+BillingGetInvoiceProcedure,
			opts...,
		),
		exportGeneratedInvoice: connect_go.NewClient[billing.ExportGeneratedInvoiceReq, billing.ExportGeneratedInvoiceRes](
			httpClient,
			baseURL+BillingExportGeneratedInvoiceProcedure,
			opts...,
		),
	}
}

// billingClient implements BillingClient.
type billingClient struct {
	getBillingPlan         *connect_go.Client[billing.GetBillingPlanReq, billing.GetBillingPlanRes]
	updateBillingPlan      *connect_go.Client[billing.UpdateBillingPlanReq, billing.UpdateBillingPlanRes]
	getInvoice             *connect_go.Client[billing.GetInvoiceReq, billing.GetInvoiceRes]
	exportGeneratedInvoice *connect_go.Client[billing.ExportGeneratedInvoiceReq, billing.ExportGeneratedInvoiceRes]
}

// GetBillingPlan calls api.v1alpha1.billing.Billing.GetBillingPlan.
func (c *billingClient) GetBillingPlan(ctx context.Context, req *connect_go.Request[billing.GetBillingPlanReq]) (*connect_go.Response[billing.GetBillingPlanRes], error) {
	return c.getBillingPlan.CallUnary(ctx, req)
}

// UpdateBillingPlan calls api.v1alpha1.billing.Billing.UpdateBillingPlan.
func (c *billingClient) UpdateBillingPlan(ctx context.Context, req *connect_go.Request[billing.UpdateBillingPlanReq]) (*connect_go.Response[billing.UpdateBillingPlanRes], error) {
	return c.updateBillingPlan.CallUnary(ctx, req)
}

// GetInvoice calls api.v1alpha1.billing.Billing.GetInvoice.
func (c *billingClient) GetInvoice(ctx context.Context, req *connect_go.Request[billing.GetInvoiceReq]) (*connect_go.Response[billing.GetInvoiceRes], error) {
	return c.getInvoice.CallUnary(ctx, req)
}

// ExportGeneratedInvoice calls api.v1alpha1.billing.Billing.ExportGeneratedInvoice.
func (c *billingClient) ExportGeneratedInvoice(ctx context.Context, req *connect_go.Request[billing.ExportGeneratedInvoiceReq]) (*connect_go.Response[billing.ExportGeneratedInvoiceRes], error) {
	return c.exportGeneratedInvoice.CallUnary(ctx, req)
}

// BillingHandler is an implementation of the api.v1alpha1.billing.Billing service.
type BillingHandler interface {
	// GetBillingPlan - returns the billing plan for the provided organization.
	GetBillingPlan(context.Context, *connect_go.Request[billing.GetBillingPlanReq]) (*connect_go.Response[billing.GetBillingPlanRes], error)
	// UpdateBillingPlan - updates the provided billing plan and it's details.
	// If some details are not provided, they will be left as is. However, if
	// deletion is desired, the DeleteBillingDetails method should be used. The
	// billing plan still follows the constraint of only having one billing detail
	// with a specific config type and event type, and so if the request contains
	// more than one billing detail with a config type and event type, the request
	// is malformed and will result in potentially unexpected behavior.
	UpdateBillingPlan(context.Context, *connect_go.Request[billing.UpdateBillingPlanReq]) (*connect_go.Response[billing.UpdateBillingPlanRes], error)
	// GetInvoice - returns the invoice for the organization. If a date is
	// provided, this will return the invoice for the organization that
	// corresponds to the billing cycle that contains the provided date. If
	// no date is provided, this will return the invoice as it currently
	// stands for the current billing cycle.
	GetInvoice(context.Context, *connect_go.Request[billing.GetInvoiceReq]) (*connect_go.Response[billing.GetInvoiceRes], error)
	ExportGeneratedInvoice(context.Context, *connect_go.Request[billing.ExportGeneratedInvoiceReq]) (*connect_go.Response[billing.ExportGeneratedInvoiceRes], error)
}

// NewBillingHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBillingHandler(svc BillingHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	billingGetBillingPlanHandler := connect_go.NewUnaryHandler(
		BillingGetBillingPlanProcedure,
		svc.GetBillingPlan,
		opts...,
	)
	billingUpdateBillingPlanHandler := connect_go.NewUnaryHandler(
		BillingUpdateBillingPlanProcedure,
		svc.UpdateBillingPlan,
		opts...,
	)
	billingGetInvoiceHandler := connect_go.NewUnaryHandler(
		BillingGetInvoiceProcedure,
		svc.GetInvoice,
		opts...,
	)
	billingExportGeneratedInvoiceHandler := connect_go.NewUnaryHandler(
		BillingExportGeneratedInvoiceProcedure,
		svc.ExportGeneratedInvoice,
		opts...,
	)
	return "/api.v1alpha1.billing.Billing/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case BillingGetBillingPlanProcedure:
			billingGetBillingPlanHandler.ServeHTTP(w, r)
		case BillingUpdateBillingPlanProcedure:
			billingUpdateBillingPlanHandler.ServeHTTP(w, r)
		case BillingGetInvoiceProcedure:
			billingGetInvoiceHandler.ServeHTTP(w, r)
		case BillingExportGeneratedInvoiceProcedure:
			billingExportGeneratedInvoiceHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedBillingHandler returns CodeUnimplemented from all methods.
type UnimplementedBillingHandler struct{}

func (UnimplementedBillingHandler) GetBillingPlan(context.Context, *connect_go.Request[billing.GetBillingPlanReq]) (*connect_go.Response[billing.GetBillingPlanRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.billing.Billing.GetBillingPlan is not implemented"))
}

func (UnimplementedBillingHandler) UpdateBillingPlan(context.Context, *connect_go.Request[billing.UpdateBillingPlanReq]) (*connect_go.Response[billing.UpdateBillingPlanRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.billing.Billing.UpdateBillingPlan is not implemented"))
}

func (UnimplementedBillingHandler) GetInvoice(context.Context, *connect_go.Request[billing.GetInvoiceReq]) (*connect_go.Response[billing.GetInvoiceRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.billing.Billing.GetInvoice is not implemented"))
}

func (UnimplementedBillingHandler) ExportGeneratedInvoice(context.Context, *connect_go.Request[billing.ExportGeneratedInvoiceReq]) (*connect_go.Response[billing.ExportGeneratedInvoiceRes], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.billing.Billing.ExportGeneratedInvoice is not implemented"))
}
