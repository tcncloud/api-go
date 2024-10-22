// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: services/translations/v1alpha1/service.proto

package translationsv1alpha1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1alpha1 "github.com/tcncloud/api-go/services/translations/v1alpha1"
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
	// TranslationsServiceName is the fully-qualified name of the TranslationsService service.
	TranslationsServiceName = "services.translations.v1alpha1.TranslationsService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TranslationsServiceTranslateTemplateProcedure is the fully-qualified name of the
	// TranslationsService's TranslateTemplate RPC.
	TranslationsServiceTranslateTemplateProcedure = "/services.translations.v1alpha1.TranslationsService/TranslateTemplate"
	// TranslationsServiceListTranslationsProcedure is the fully-qualified name of the
	// TranslationsService's ListTranslations RPC.
	TranslationsServiceListTranslationsProcedure = "/services.translations.v1alpha1.TranslationsService/ListTranslations"
	// TranslationsServiceListLanguagesProcedure is the fully-qualified name of the
	// TranslationsService's ListLanguages RPC.
	TranslationsServiceListLanguagesProcedure = "/services.translations.v1alpha1.TranslationsService/ListLanguages"
	// TranslationsServiceListContextsProcedure is the fully-qualified name of the TranslationsService's
	// ListContexts RPC.
	TranslationsServiceListContextsProcedure = "/services.translations.v1alpha1.TranslationsService/ListContexts"
	// TranslationsServiceCreateTranslationProcedure is the fully-qualified name of the
	// TranslationsService's CreateTranslation RPC.
	TranslationsServiceCreateTranslationProcedure = "/services.translations.v1alpha1.TranslationsService/CreateTranslation"
	// TranslationsServiceUpdateTranslationProcedure is the fully-qualified name of the
	// TranslationsService's UpdateTranslation RPC.
	TranslationsServiceUpdateTranslationProcedure = "/services.translations.v1alpha1.TranslationsService/UpdateTranslation"
	// TranslationsServiceTriggerLLMTranslationProcedure is the fully-qualified name of the
	// TranslationsService's TriggerLLMTranslation RPC.
	TranslationsServiceTriggerLLMTranslationProcedure = "/services.translations.v1alpha1.TranslationsService/TriggerLLMTranslation"
	// TranslationsServiceTriggerLLMTranslationsProcedure is the fully-qualified name of the
	// TranslationsService's TriggerLLMTranslations RPC.
	TranslationsServiceTriggerLLMTranslationsProcedure = "/services.translations.v1alpha1.TranslationsService/TriggerLLMTranslations"
	// TranslationsServiceSetSystemMessageProcedure is the fully-qualified name of the
	// TranslationsService's SetSystemMessage RPC.
	TranslationsServiceSetSystemMessageProcedure = "/services.translations.v1alpha1.TranslationsService/SetSystemMessage"
	// TranslationsServiceGetSystemMessageProcedure is the fully-qualified name of the
	// TranslationsService's GetSystemMessage RPC.
	TranslationsServiceGetSystemMessageProcedure = "/services.translations.v1alpha1.TranslationsService/GetSystemMessage"
	// TranslationsServiceTestSystemMessageProcedure is the fully-qualified name of the
	// TranslationsService's TestSystemMessage RPC.
	TranslationsServiceTestSystemMessageProcedure = "/services.translations.v1alpha1.TranslationsService/TestSystemMessage"
	// TranslationsServiceEnableContextProcedure is the fully-qualified name of the
	// TranslationsService's EnableContext RPC.
	TranslationsServiceEnableContextProcedure = "/services.translations.v1alpha1.TranslationsService/EnableContext"
	// TranslationsServiceDisableContextProcedure is the fully-qualified name of the
	// TranslationsService's DisableContext RPC.
	TranslationsServiceDisableContextProcedure = "/services.translations.v1alpha1.TranslationsService/DisableContext"
	// TranslationsServiceDeleteTranslationsByTemplateProcedure is the fully-qualified name of the
	// TranslationsService's DeleteTranslationsByTemplate RPC.
	TranslationsServiceDeleteTranslationsByTemplateProcedure = "/services.translations.v1alpha1.TranslationsService/DeleteTranslationsByTemplate"
	// TranslationsServiceBulkDeleteTranslationsProcedure is the fully-qualified name of the
	// TranslationsService's BulkDeleteTranslations RPC.
	TranslationsServiceBulkDeleteTranslationsProcedure = "/services.translations.v1alpha1.TranslationsService/BulkDeleteTranslations"
)

// TranslationsServiceClient is a client for the services.translations.v1alpha1.TranslationsService
// service.
type TranslationsServiceClient interface {
	// Translate a template for a given context and language.
	// Required permissions:
	//
	//	Any Authenticated User
	//
	// Errors:
	//   - grpc.AlreadyExists : This template is already translated for the given context and language.
	//   - grpc.InvalidArgument: The request is not valid.
	TranslateTemplate(context.Context, *connect_go.Request[v1alpha1.TranslateTemplateRequest]) (*connect_go.Response[v1alpha1.TranslateTemplateResponse], error)
	// Lists translations by context/language
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	//   - grpc.NotFound: No templates found for the given context and language.
	ListTranslations(context.Context, *connect_go.Request[v1alpha1.ListTranslationsRequest]) (*connect_go.Response[v1alpha1.ListTranslationsResponse], error)
	// Lists localization languages
	// Required permissions:
	//   - Any Authenticated User
	ListLanguages(context.Context, *connect_go.Request[v1alpha1.ListLanguagesRequest]) (*connect_go.Response[v1alpha1.ListLanguagesResponse], error)
	// Lists localization contexts
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	ListContexts(context.Context, *connect_go.Request[v1alpha1.ListContextsRequest]) (*connect_go.Response[v1alpha1.ListContextsResponse], error)
	// Creates a new Translation
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	//   - grpc.AlreadyExists: The template already exists for the given context and language (use override method).
	CreateTranslation(context.Context, *connect_go.Request[v1alpha1.CreateTranslationRequest]) (*connect_go.Response[v1alpha1.CreateTranslationResponse], error)
	// Overrides the translation for a given translationID
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	UpdateTranslation(context.Context, *connect_go.Request[v1alpha1.UpdateTranslationRequest]) (*connect_go.Response[v1alpha1.UpdateTranslationResponse], error)
	// Re-run the LLM translation for a given translationID
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	TriggerLLMTranslation(context.Context, *connect_go.Request[v1alpha1.TriggerLLMTranslationRequest]) (*connect_go.Response[v1alpha1.TriggerLLMTranslationResponse], error)
	// re-run all translations for a given context (WARNING - this should be ran sparingly as it is a heavy operation and costs money)
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	TriggerLLMTranslations(context.Context, *connect_go.Request[v1alpha1.TriggerLLMTranslationsRequest]) (*connect_go.Response[v1alpha1.TriggerLLMTranslationsResponse], error)
	// set/get context system message to give more tuned LLMs when translating for that context (WARNING - this overrides the previous system message for the context if exists)
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	SetSystemMessage(context.Context, *connect_go.Request[v1alpha1.SetSystemMessageRequest]) (*connect_go.Response[v1alpha1.SetSystemMessageResponse], error)
	GetSystemMessage(context.Context, *connect_go.Request[v1alpha1.GetSystemMessageRequest]) (*connect_go.Response[v1alpha1.GetSystemMessageResponse], error)
	// Gives a translation for a system message, template and language with no side effects (Used for testing system messages)
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	TestSystemMessage(context.Context, *connect_go.Request[v1alpha1.TestSystemMessageRequest]) (*connect_go.Response[v1alpha1.TestSystemMessageResponse], error)
	// enable a context for LLM translations
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	//   - grpc.NotFound: No context found
	EnableContext(context.Context, *connect_go.Request[v1alpha1.EnableContextRequest]) (*connect_go.Response[v1alpha1.EnableContextResponse], error)
	// disable a context for LLM translations
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	DisableContext(context.Context, *connect_go.Request[v1alpha1.DisableContextRequest]) (*connect_go.Response[v1alpha1.DisableContextResponse], error)
	// Delete translations by template and context
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	//   - grpc.NotFound: No translations found for the given template and context.
	DeleteTranslationsByTemplate(context.Context, *connect_go.Request[v1alpha1.DeleteTranslationsByTemplateRequest]) (*connect_go.Response[v1alpha1.DeleteTranslationsByTemplateResponse], error)
	// Bulk delete translations
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	BulkDeleteTranslations(context.Context, *connect_go.Request[v1alpha1.BulkDeleteTranslationsRequest]) (*connect_go.Response[v1alpha1.BulkDeleteTranslationsResponse], error)
}

// NewTranslationsServiceClient constructs a client for the
// services.translations.v1alpha1.TranslationsService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTranslationsServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TranslationsServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &translationsServiceClient{
		translateTemplate: connect_go.NewClient[v1alpha1.TranslateTemplateRequest, v1alpha1.TranslateTemplateResponse](
			httpClient,
			baseURL+TranslationsServiceTranslateTemplateProcedure,
			opts...,
		),
		listTranslations: connect_go.NewClient[v1alpha1.ListTranslationsRequest, v1alpha1.ListTranslationsResponse](
			httpClient,
			baseURL+TranslationsServiceListTranslationsProcedure,
			opts...,
		),
		listLanguages: connect_go.NewClient[v1alpha1.ListLanguagesRequest, v1alpha1.ListLanguagesResponse](
			httpClient,
			baseURL+TranslationsServiceListLanguagesProcedure,
			opts...,
		),
		listContexts: connect_go.NewClient[v1alpha1.ListContextsRequest, v1alpha1.ListContextsResponse](
			httpClient,
			baseURL+TranslationsServiceListContextsProcedure,
			opts...,
		),
		createTranslation: connect_go.NewClient[v1alpha1.CreateTranslationRequest, v1alpha1.CreateTranslationResponse](
			httpClient,
			baseURL+TranslationsServiceCreateTranslationProcedure,
			opts...,
		),
		updateTranslation: connect_go.NewClient[v1alpha1.UpdateTranslationRequest, v1alpha1.UpdateTranslationResponse](
			httpClient,
			baseURL+TranslationsServiceUpdateTranslationProcedure,
			opts...,
		),
		triggerLLMTranslation: connect_go.NewClient[v1alpha1.TriggerLLMTranslationRequest, v1alpha1.TriggerLLMTranslationResponse](
			httpClient,
			baseURL+TranslationsServiceTriggerLLMTranslationProcedure,
			opts...,
		),
		triggerLLMTranslations: connect_go.NewClient[v1alpha1.TriggerLLMTranslationsRequest, v1alpha1.TriggerLLMTranslationsResponse](
			httpClient,
			baseURL+TranslationsServiceTriggerLLMTranslationsProcedure,
			opts...,
		),
		setSystemMessage: connect_go.NewClient[v1alpha1.SetSystemMessageRequest, v1alpha1.SetSystemMessageResponse](
			httpClient,
			baseURL+TranslationsServiceSetSystemMessageProcedure,
			opts...,
		),
		getSystemMessage: connect_go.NewClient[v1alpha1.GetSystemMessageRequest, v1alpha1.GetSystemMessageResponse](
			httpClient,
			baseURL+TranslationsServiceGetSystemMessageProcedure,
			opts...,
		),
		testSystemMessage: connect_go.NewClient[v1alpha1.TestSystemMessageRequest, v1alpha1.TestSystemMessageResponse](
			httpClient,
			baseURL+TranslationsServiceTestSystemMessageProcedure,
			opts...,
		),
		enableContext: connect_go.NewClient[v1alpha1.EnableContextRequest, v1alpha1.EnableContextResponse](
			httpClient,
			baseURL+TranslationsServiceEnableContextProcedure,
			opts...,
		),
		disableContext: connect_go.NewClient[v1alpha1.DisableContextRequest, v1alpha1.DisableContextResponse](
			httpClient,
			baseURL+TranslationsServiceDisableContextProcedure,
			opts...,
		),
		deleteTranslationsByTemplate: connect_go.NewClient[v1alpha1.DeleteTranslationsByTemplateRequest, v1alpha1.DeleteTranslationsByTemplateResponse](
			httpClient,
			baseURL+TranslationsServiceDeleteTranslationsByTemplateProcedure,
			opts...,
		),
		bulkDeleteTranslations: connect_go.NewClient[v1alpha1.BulkDeleteTranslationsRequest, v1alpha1.BulkDeleteTranslationsResponse](
			httpClient,
			baseURL+TranslationsServiceBulkDeleteTranslationsProcedure,
			opts...,
		),
	}
}

// translationsServiceClient implements TranslationsServiceClient.
type translationsServiceClient struct {
	translateTemplate            *connect_go.Client[v1alpha1.TranslateTemplateRequest, v1alpha1.TranslateTemplateResponse]
	listTranslations             *connect_go.Client[v1alpha1.ListTranslationsRequest, v1alpha1.ListTranslationsResponse]
	listLanguages                *connect_go.Client[v1alpha1.ListLanguagesRequest, v1alpha1.ListLanguagesResponse]
	listContexts                 *connect_go.Client[v1alpha1.ListContextsRequest, v1alpha1.ListContextsResponse]
	createTranslation            *connect_go.Client[v1alpha1.CreateTranslationRequest, v1alpha1.CreateTranslationResponse]
	updateTranslation            *connect_go.Client[v1alpha1.UpdateTranslationRequest, v1alpha1.UpdateTranslationResponse]
	triggerLLMTranslation        *connect_go.Client[v1alpha1.TriggerLLMTranslationRequest, v1alpha1.TriggerLLMTranslationResponse]
	triggerLLMTranslations       *connect_go.Client[v1alpha1.TriggerLLMTranslationsRequest, v1alpha1.TriggerLLMTranslationsResponse]
	setSystemMessage             *connect_go.Client[v1alpha1.SetSystemMessageRequest, v1alpha1.SetSystemMessageResponse]
	getSystemMessage             *connect_go.Client[v1alpha1.GetSystemMessageRequest, v1alpha1.GetSystemMessageResponse]
	testSystemMessage            *connect_go.Client[v1alpha1.TestSystemMessageRequest, v1alpha1.TestSystemMessageResponse]
	enableContext                *connect_go.Client[v1alpha1.EnableContextRequest, v1alpha1.EnableContextResponse]
	disableContext               *connect_go.Client[v1alpha1.DisableContextRequest, v1alpha1.DisableContextResponse]
	deleteTranslationsByTemplate *connect_go.Client[v1alpha1.DeleteTranslationsByTemplateRequest, v1alpha1.DeleteTranslationsByTemplateResponse]
	bulkDeleteTranslations       *connect_go.Client[v1alpha1.BulkDeleteTranslationsRequest, v1alpha1.BulkDeleteTranslationsResponse]
}

// TranslateTemplate calls services.translations.v1alpha1.TranslationsService.TranslateTemplate.
func (c *translationsServiceClient) TranslateTemplate(ctx context.Context, req *connect_go.Request[v1alpha1.TranslateTemplateRequest]) (*connect_go.Response[v1alpha1.TranslateTemplateResponse], error) {
	return c.translateTemplate.CallUnary(ctx, req)
}

// ListTranslations calls services.translations.v1alpha1.TranslationsService.ListTranslations.
func (c *translationsServiceClient) ListTranslations(ctx context.Context, req *connect_go.Request[v1alpha1.ListTranslationsRequest]) (*connect_go.Response[v1alpha1.ListTranslationsResponse], error) {
	return c.listTranslations.CallUnary(ctx, req)
}

// ListLanguages calls services.translations.v1alpha1.TranslationsService.ListLanguages.
func (c *translationsServiceClient) ListLanguages(ctx context.Context, req *connect_go.Request[v1alpha1.ListLanguagesRequest]) (*connect_go.Response[v1alpha1.ListLanguagesResponse], error) {
	return c.listLanguages.CallUnary(ctx, req)
}

// ListContexts calls services.translations.v1alpha1.TranslationsService.ListContexts.
func (c *translationsServiceClient) ListContexts(ctx context.Context, req *connect_go.Request[v1alpha1.ListContextsRequest]) (*connect_go.Response[v1alpha1.ListContextsResponse], error) {
	return c.listContexts.CallUnary(ctx, req)
}

// CreateTranslation calls services.translations.v1alpha1.TranslationsService.CreateTranslation.
func (c *translationsServiceClient) CreateTranslation(ctx context.Context, req *connect_go.Request[v1alpha1.CreateTranslationRequest]) (*connect_go.Response[v1alpha1.CreateTranslationResponse], error) {
	return c.createTranslation.CallUnary(ctx, req)
}

// UpdateTranslation calls services.translations.v1alpha1.TranslationsService.UpdateTranslation.
func (c *translationsServiceClient) UpdateTranslation(ctx context.Context, req *connect_go.Request[v1alpha1.UpdateTranslationRequest]) (*connect_go.Response[v1alpha1.UpdateTranslationResponse], error) {
	return c.updateTranslation.CallUnary(ctx, req)
}

// TriggerLLMTranslation calls
// services.translations.v1alpha1.TranslationsService.TriggerLLMTranslation.
func (c *translationsServiceClient) TriggerLLMTranslation(ctx context.Context, req *connect_go.Request[v1alpha1.TriggerLLMTranslationRequest]) (*connect_go.Response[v1alpha1.TriggerLLMTranslationResponse], error) {
	return c.triggerLLMTranslation.CallUnary(ctx, req)
}

// TriggerLLMTranslations calls
// services.translations.v1alpha1.TranslationsService.TriggerLLMTranslations.
func (c *translationsServiceClient) TriggerLLMTranslations(ctx context.Context, req *connect_go.Request[v1alpha1.TriggerLLMTranslationsRequest]) (*connect_go.Response[v1alpha1.TriggerLLMTranslationsResponse], error) {
	return c.triggerLLMTranslations.CallUnary(ctx, req)
}

// SetSystemMessage calls services.translations.v1alpha1.TranslationsService.SetSystemMessage.
func (c *translationsServiceClient) SetSystemMessage(ctx context.Context, req *connect_go.Request[v1alpha1.SetSystemMessageRequest]) (*connect_go.Response[v1alpha1.SetSystemMessageResponse], error) {
	return c.setSystemMessage.CallUnary(ctx, req)
}

// GetSystemMessage calls services.translations.v1alpha1.TranslationsService.GetSystemMessage.
func (c *translationsServiceClient) GetSystemMessage(ctx context.Context, req *connect_go.Request[v1alpha1.GetSystemMessageRequest]) (*connect_go.Response[v1alpha1.GetSystemMessageResponse], error) {
	return c.getSystemMessage.CallUnary(ctx, req)
}

// TestSystemMessage calls services.translations.v1alpha1.TranslationsService.TestSystemMessage.
func (c *translationsServiceClient) TestSystemMessage(ctx context.Context, req *connect_go.Request[v1alpha1.TestSystemMessageRequest]) (*connect_go.Response[v1alpha1.TestSystemMessageResponse], error) {
	return c.testSystemMessage.CallUnary(ctx, req)
}

// EnableContext calls services.translations.v1alpha1.TranslationsService.EnableContext.
func (c *translationsServiceClient) EnableContext(ctx context.Context, req *connect_go.Request[v1alpha1.EnableContextRequest]) (*connect_go.Response[v1alpha1.EnableContextResponse], error) {
	return c.enableContext.CallUnary(ctx, req)
}

// DisableContext calls services.translations.v1alpha1.TranslationsService.DisableContext.
func (c *translationsServiceClient) DisableContext(ctx context.Context, req *connect_go.Request[v1alpha1.DisableContextRequest]) (*connect_go.Response[v1alpha1.DisableContextResponse], error) {
	return c.disableContext.CallUnary(ctx, req)
}

// DeleteTranslationsByTemplate calls
// services.translations.v1alpha1.TranslationsService.DeleteTranslationsByTemplate.
func (c *translationsServiceClient) DeleteTranslationsByTemplate(ctx context.Context, req *connect_go.Request[v1alpha1.DeleteTranslationsByTemplateRequest]) (*connect_go.Response[v1alpha1.DeleteTranslationsByTemplateResponse], error) {
	return c.deleteTranslationsByTemplate.CallUnary(ctx, req)
}

// BulkDeleteTranslations calls
// services.translations.v1alpha1.TranslationsService.BulkDeleteTranslations.
func (c *translationsServiceClient) BulkDeleteTranslations(ctx context.Context, req *connect_go.Request[v1alpha1.BulkDeleteTranslationsRequest]) (*connect_go.Response[v1alpha1.BulkDeleteTranslationsResponse], error) {
	return c.bulkDeleteTranslations.CallUnary(ctx, req)
}

// TranslationsServiceHandler is an implementation of the
// services.translations.v1alpha1.TranslationsService service.
type TranslationsServiceHandler interface {
	// Translate a template for a given context and language.
	// Required permissions:
	//
	//	Any Authenticated User
	//
	// Errors:
	//   - grpc.AlreadyExists : This template is already translated for the given context and language.
	//   - grpc.InvalidArgument: The request is not valid.
	TranslateTemplate(context.Context, *connect_go.Request[v1alpha1.TranslateTemplateRequest]) (*connect_go.Response[v1alpha1.TranslateTemplateResponse], error)
	// Lists translations by context/language
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	//   - grpc.NotFound: No templates found for the given context and language.
	ListTranslations(context.Context, *connect_go.Request[v1alpha1.ListTranslationsRequest]) (*connect_go.Response[v1alpha1.ListTranslationsResponse], error)
	// Lists localization languages
	// Required permissions:
	//   - Any Authenticated User
	ListLanguages(context.Context, *connect_go.Request[v1alpha1.ListLanguagesRequest]) (*connect_go.Response[v1alpha1.ListLanguagesResponse], error)
	// Lists localization contexts
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	ListContexts(context.Context, *connect_go.Request[v1alpha1.ListContextsRequest]) (*connect_go.Response[v1alpha1.ListContextsResponse], error)
	// Creates a new Translation
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	//   - grpc.AlreadyExists: The template already exists for the given context and language (use override method).
	CreateTranslation(context.Context, *connect_go.Request[v1alpha1.CreateTranslationRequest]) (*connect_go.Response[v1alpha1.CreateTranslationResponse], error)
	// Overrides the translation for a given translationID
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	UpdateTranslation(context.Context, *connect_go.Request[v1alpha1.UpdateTranslationRequest]) (*connect_go.Response[v1alpha1.UpdateTranslationResponse], error)
	// Re-run the LLM translation for a given translationID
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	TriggerLLMTranslation(context.Context, *connect_go.Request[v1alpha1.TriggerLLMTranslationRequest]) (*connect_go.Response[v1alpha1.TriggerLLMTranslationResponse], error)
	// re-run all translations for a given context (WARNING - this should be ran sparingly as it is a heavy operation and costs money)
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	TriggerLLMTranslations(context.Context, *connect_go.Request[v1alpha1.TriggerLLMTranslationsRequest]) (*connect_go.Response[v1alpha1.TriggerLLMTranslationsResponse], error)
	// set/get context system message to give more tuned LLMs when translating for that context (WARNING - this overrides the previous system message for the context if exists)
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	SetSystemMessage(context.Context, *connect_go.Request[v1alpha1.SetSystemMessageRequest]) (*connect_go.Response[v1alpha1.SetSystemMessageResponse], error)
	GetSystemMessage(context.Context, *connect_go.Request[v1alpha1.GetSystemMessageRequest]) (*connect_go.Response[v1alpha1.GetSystemMessageResponse], error)
	// Gives a translation for a system message, template and language with no side effects (Used for testing system messages)
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	TestSystemMessage(context.Context, *connect_go.Request[v1alpha1.TestSystemMessageRequest]) (*connect_go.Response[v1alpha1.TestSystemMessageResponse], error)
	// enable a context for LLM translations
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	//   - grpc.NotFound: No context found
	EnableContext(context.Context, *connect_go.Request[v1alpha1.EnableContextRequest]) (*connect_go.Response[v1alpha1.EnableContextResponse], error)
	// disable a context for LLM translations
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	DisableContext(context.Context, *connect_go.Request[v1alpha1.DisableContextRequest]) (*connect_go.Response[v1alpha1.DisableContextResponse], error)
	// Delete translations by template and context
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	//   - grpc.NotFound: No translations found for the given template and context.
	DeleteTranslationsByTemplate(context.Context, *connect_go.Request[v1alpha1.DeleteTranslationsByTemplateRequest]) (*connect_go.Response[v1alpha1.DeleteTranslationsByTemplateResponse], error)
	// Bulk delete translations
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	BulkDeleteTranslations(context.Context, *connect_go.Request[v1alpha1.BulkDeleteTranslationsRequest]) (*connect_go.Response[v1alpha1.BulkDeleteTranslationsResponse], error)
}

// NewTranslationsServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTranslationsServiceHandler(svc TranslationsServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	translationsServiceTranslateTemplateHandler := connect_go.NewUnaryHandler(
		TranslationsServiceTranslateTemplateProcedure,
		svc.TranslateTemplate,
		opts...,
	)
	translationsServiceListTranslationsHandler := connect_go.NewUnaryHandler(
		TranslationsServiceListTranslationsProcedure,
		svc.ListTranslations,
		opts...,
	)
	translationsServiceListLanguagesHandler := connect_go.NewUnaryHandler(
		TranslationsServiceListLanguagesProcedure,
		svc.ListLanguages,
		opts...,
	)
	translationsServiceListContextsHandler := connect_go.NewUnaryHandler(
		TranslationsServiceListContextsProcedure,
		svc.ListContexts,
		opts...,
	)
	translationsServiceCreateTranslationHandler := connect_go.NewUnaryHandler(
		TranslationsServiceCreateTranslationProcedure,
		svc.CreateTranslation,
		opts...,
	)
	translationsServiceUpdateTranslationHandler := connect_go.NewUnaryHandler(
		TranslationsServiceUpdateTranslationProcedure,
		svc.UpdateTranslation,
		opts...,
	)
	translationsServiceTriggerLLMTranslationHandler := connect_go.NewUnaryHandler(
		TranslationsServiceTriggerLLMTranslationProcedure,
		svc.TriggerLLMTranslation,
		opts...,
	)
	translationsServiceTriggerLLMTranslationsHandler := connect_go.NewUnaryHandler(
		TranslationsServiceTriggerLLMTranslationsProcedure,
		svc.TriggerLLMTranslations,
		opts...,
	)
	translationsServiceSetSystemMessageHandler := connect_go.NewUnaryHandler(
		TranslationsServiceSetSystemMessageProcedure,
		svc.SetSystemMessage,
		opts...,
	)
	translationsServiceGetSystemMessageHandler := connect_go.NewUnaryHandler(
		TranslationsServiceGetSystemMessageProcedure,
		svc.GetSystemMessage,
		opts...,
	)
	translationsServiceTestSystemMessageHandler := connect_go.NewUnaryHandler(
		TranslationsServiceTestSystemMessageProcedure,
		svc.TestSystemMessage,
		opts...,
	)
	translationsServiceEnableContextHandler := connect_go.NewUnaryHandler(
		TranslationsServiceEnableContextProcedure,
		svc.EnableContext,
		opts...,
	)
	translationsServiceDisableContextHandler := connect_go.NewUnaryHandler(
		TranslationsServiceDisableContextProcedure,
		svc.DisableContext,
		opts...,
	)
	translationsServiceDeleteTranslationsByTemplateHandler := connect_go.NewUnaryHandler(
		TranslationsServiceDeleteTranslationsByTemplateProcedure,
		svc.DeleteTranslationsByTemplate,
		opts...,
	)
	translationsServiceBulkDeleteTranslationsHandler := connect_go.NewUnaryHandler(
		TranslationsServiceBulkDeleteTranslationsProcedure,
		svc.BulkDeleteTranslations,
		opts...,
	)
	return "/services.translations.v1alpha1.TranslationsService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TranslationsServiceTranslateTemplateProcedure:
			translationsServiceTranslateTemplateHandler.ServeHTTP(w, r)
		case TranslationsServiceListTranslationsProcedure:
			translationsServiceListTranslationsHandler.ServeHTTP(w, r)
		case TranslationsServiceListLanguagesProcedure:
			translationsServiceListLanguagesHandler.ServeHTTP(w, r)
		case TranslationsServiceListContextsProcedure:
			translationsServiceListContextsHandler.ServeHTTP(w, r)
		case TranslationsServiceCreateTranslationProcedure:
			translationsServiceCreateTranslationHandler.ServeHTTP(w, r)
		case TranslationsServiceUpdateTranslationProcedure:
			translationsServiceUpdateTranslationHandler.ServeHTTP(w, r)
		case TranslationsServiceTriggerLLMTranslationProcedure:
			translationsServiceTriggerLLMTranslationHandler.ServeHTTP(w, r)
		case TranslationsServiceTriggerLLMTranslationsProcedure:
			translationsServiceTriggerLLMTranslationsHandler.ServeHTTP(w, r)
		case TranslationsServiceSetSystemMessageProcedure:
			translationsServiceSetSystemMessageHandler.ServeHTTP(w, r)
		case TranslationsServiceGetSystemMessageProcedure:
			translationsServiceGetSystemMessageHandler.ServeHTTP(w, r)
		case TranslationsServiceTestSystemMessageProcedure:
			translationsServiceTestSystemMessageHandler.ServeHTTP(w, r)
		case TranslationsServiceEnableContextProcedure:
			translationsServiceEnableContextHandler.ServeHTTP(w, r)
		case TranslationsServiceDisableContextProcedure:
			translationsServiceDisableContextHandler.ServeHTTP(w, r)
		case TranslationsServiceDeleteTranslationsByTemplateProcedure:
			translationsServiceDeleteTranslationsByTemplateHandler.ServeHTTP(w, r)
		case TranslationsServiceBulkDeleteTranslationsProcedure:
			translationsServiceBulkDeleteTranslationsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTranslationsServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTranslationsServiceHandler struct{}

func (UnimplementedTranslationsServiceHandler) TranslateTemplate(context.Context, *connect_go.Request[v1alpha1.TranslateTemplateRequest]) (*connect_go.Response[v1alpha1.TranslateTemplateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.translations.v1alpha1.TranslationsService.TranslateTemplate is not implemented"))
}

func (UnimplementedTranslationsServiceHandler) ListTranslations(context.Context, *connect_go.Request[v1alpha1.ListTranslationsRequest]) (*connect_go.Response[v1alpha1.ListTranslationsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.translations.v1alpha1.TranslationsService.ListTranslations is not implemented"))
}

func (UnimplementedTranslationsServiceHandler) ListLanguages(context.Context, *connect_go.Request[v1alpha1.ListLanguagesRequest]) (*connect_go.Response[v1alpha1.ListLanguagesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.translations.v1alpha1.TranslationsService.ListLanguages is not implemented"))
}

func (UnimplementedTranslationsServiceHandler) ListContexts(context.Context, *connect_go.Request[v1alpha1.ListContextsRequest]) (*connect_go.Response[v1alpha1.ListContextsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.translations.v1alpha1.TranslationsService.ListContexts is not implemented"))
}

func (UnimplementedTranslationsServiceHandler) CreateTranslation(context.Context, *connect_go.Request[v1alpha1.CreateTranslationRequest]) (*connect_go.Response[v1alpha1.CreateTranslationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.translations.v1alpha1.TranslationsService.CreateTranslation is not implemented"))
}

func (UnimplementedTranslationsServiceHandler) UpdateTranslation(context.Context, *connect_go.Request[v1alpha1.UpdateTranslationRequest]) (*connect_go.Response[v1alpha1.UpdateTranslationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.translations.v1alpha1.TranslationsService.UpdateTranslation is not implemented"))
}

func (UnimplementedTranslationsServiceHandler) TriggerLLMTranslation(context.Context, *connect_go.Request[v1alpha1.TriggerLLMTranslationRequest]) (*connect_go.Response[v1alpha1.TriggerLLMTranslationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.translations.v1alpha1.TranslationsService.TriggerLLMTranslation is not implemented"))
}

func (UnimplementedTranslationsServiceHandler) TriggerLLMTranslations(context.Context, *connect_go.Request[v1alpha1.TriggerLLMTranslationsRequest]) (*connect_go.Response[v1alpha1.TriggerLLMTranslationsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.translations.v1alpha1.TranslationsService.TriggerLLMTranslations is not implemented"))
}

func (UnimplementedTranslationsServiceHandler) SetSystemMessage(context.Context, *connect_go.Request[v1alpha1.SetSystemMessageRequest]) (*connect_go.Response[v1alpha1.SetSystemMessageResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.translations.v1alpha1.TranslationsService.SetSystemMessage is not implemented"))
}

func (UnimplementedTranslationsServiceHandler) GetSystemMessage(context.Context, *connect_go.Request[v1alpha1.GetSystemMessageRequest]) (*connect_go.Response[v1alpha1.GetSystemMessageResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.translations.v1alpha1.TranslationsService.GetSystemMessage is not implemented"))
}

func (UnimplementedTranslationsServiceHandler) TestSystemMessage(context.Context, *connect_go.Request[v1alpha1.TestSystemMessageRequest]) (*connect_go.Response[v1alpha1.TestSystemMessageResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.translations.v1alpha1.TranslationsService.TestSystemMessage is not implemented"))
}

func (UnimplementedTranslationsServiceHandler) EnableContext(context.Context, *connect_go.Request[v1alpha1.EnableContextRequest]) (*connect_go.Response[v1alpha1.EnableContextResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.translations.v1alpha1.TranslationsService.EnableContext is not implemented"))
}

func (UnimplementedTranslationsServiceHandler) DisableContext(context.Context, *connect_go.Request[v1alpha1.DisableContextRequest]) (*connect_go.Response[v1alpha1.DisableContextResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.translations.v1alpha1.TranslationsService.DisableContext is not implemented"))
}

func (UnimplementedTranslationsServiceHandler) DeleteTranslationsByTemplate(context.Context, *connect_go.Request[v1alpha1.DeleteTranslationsByTemplateRequest]) (*connect_go.Response[v1alpha1.DeleteTranslationsByTemplateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.translations.v1alpha1.TranslationsService.DeleteTranslationsByTemplate is not implemented"))
}

func (UnimplementedTranslationsServiceHandler) BulkDeleteTranslations(context.Context, *connect_go.Request[v1alpha1.BulkDeleteTranslationsRequest]) (*connect_go.Response[v1alpha1.BulkDeleteTranslationsResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("services.translations.v1alpha1.TranslationsService.BulkDeleteTranslations is not implemented"))
}
