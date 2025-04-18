// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: services/translations/v1alpha1/service.proto

package translationsv1alpha1

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
	TranslationsService_TranslateTemplate_FullMethodName            = "/services.translations.v1alpha1.TranslationsService/TranslateTemplate"
	TranslationsService_ListTranslations_FullMethodName             = "/services.translations.v1alpha1.TranslationsService/ListTranslations"
	TranslationsService_ListLanguages_FullMethodName                = "/services.translations.v1alpha1.TranslationsService/ListLanguages"
	TranslationsService_ListContexts_FullMethodName                 = "/services.translations.v1alpha1.TranslationsService/ListContexts"
	TranslationsService_CreateTranslation_FullMethodName            = "/services.translations.v1alpha1.TranslationsService/CreateTranslation"
	TranslationsService_UpdateTranslation_FullMethodName            = "/services.translations.v1alpha1.TranslationsService/UpdateTranslation"
	TranslationsService_TriggerLLMTranslation_FullMethodName        = "/services.translations.v1alpha1.TranslationsService/TriggerLLMTranslation"
	TranslationsService_TriggerLLMTranslations_FullMethodName       = "/services.translations.v1alpha1.TranslationsService/TriggerLLMTranslations"
	TranslationsService_SetSystemMessage_FullMethodName             = "/services.translations.v1alpha1.TranslationsService/SetSystemMessage"
	TranslationsService_GetSystemMessage_FullMethodName             = "/services.translations.v1alpha1.TranslationsService/GetSystemMessage"
	TranslationsService_TestSystemMessage_FullMethodName            = "/services.translations.v1alpha1.TranslationsService/TestSystemMessage"
	TranslationsService_EnableContext_FullMethodName                = "/services.translations.v1alpha1.TranslationsService/EnableContext"
	TranslationsService_DisableContext_FullMethodName               = "/services.translations.v1alpha1.TranslationsService/DisableContext"
	TranslationsService_DeleteTranslationsByTemplate_FullMethodName = "/services.translations.v1alpha1.TranslationsService/DeleteTranslationsByTemplate"
	TranslationsService_BulkDeleteTranslations_FullMethodName       = "/services.translations.v1alpha1.TranslationsService/BulkDeleteTranslations"
)

// TranslationsServiceClient is the client API for TranslationsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TranslationsServiceClient interface {
	// Translate a template for a given context and language.
	// Required permissions:
	//
	//	Any Authenticated User
	//
	// Errors:
	//   - grpc.AlreadyExists : This template is already translated for the given context and language.
	//   - grpc.InvalidArgument: The request is not valid.
	TranslateTemplate(ctx context.Context, in *TranslateTemplateRequest, opts ...grpc.CallOption) (*TranslateTemplateResponse, error)
	// Lists translations by context/language
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	//   - grpc.NotFound: No templates found for the given context and language.
	ListTranslations(ctx context.Context, in *ListTranslationsRequest, opts ...grpc.CallOption) (*ListTranslationsResponse, error)
	// Lists localization languages
	// Required permissions:
	//   - Any Authenticated User
	ListLanguages(ctx context.Context, in *ListLanguagesRequest, opts ...grpc.CallOption) (*ListLanguagesResponse, error)
	// Lists localization contexts
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	ListContexts(ctx context.Context, in *ListContextsRequest, opts ...grpc.CallOption) (*ListContextsResponse, error)
	// Creates a new Translation
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	//   - grpc.AlreadyExists: The template already exists for the given context and language (use override method).
	CreateTranslation(ctx context.Context, in *CreateTranslationRequest, opts ...grpc.CallOption) (*CreateTranslationResponse, error)
	// Overrides the translation for a given translationID
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	UpdateTranslation(ctx context.Context, in *UpdateTranslationRequest, opts ...grpc.CallOption) (*UpdateTranslationResponse, error)
	// Re-run the LLM translation for a given translationID
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	TriggerLLMTranslation(ctx context.Context, in *TriggerLLMTranslationRequest, opts ...grpc.CallOption) (*TriggerLLMTranslationResponse, error)
	// re-run all translations for a given context (WARNING - this should be ran sparingly as it is a heavy operation and costs money)
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	TriggerLLMTranslations(ctx context.Context, in *TriggerLLMTranslationsRequest, opts ...grpc.CallOption) (*TriggerLLMTranslationsResponse, error)
	// set/get context system message to give more tuned LLMs when translating for that context (WARNING - this overrides the previous system message for the context if exists)
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	SetSystemMessage(ctx context.Context, in *SetSystemMessageRequest, opts ...grpc.CallOption) (*SetSystemMessageResponse, error)
	GetSystemMessage(ctx context.Context, in *GetSystemMessageRequest, opts ...grpc.CallOption) (*GetSystemMessageResponse, error)
	// Gives a translation for a system message, template and language with no side effects (Used for testing system messages)
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	TestSystemMessage(ctx context.Context, in *TestSystemMessageRequest, opts ...grpc.CallOption) (*TestSystemMessageResponse, error)
	// enable a context for LLM translations
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	//   - grpc.NotFound: No context found
	EnableContext(ctx context.Context, in *EnableContextRequest, opts ...grpc.CallOption) (*EnableContextResponse, error)
	// disable a context for LLM translations
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	DisableContext(ctx context.Context, in *DisableContextRequest, opts ...grpc.CallOption) (*DisableContextResponse, error)
	// Delete translations by template and context
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	//   - grpc.NotFound: No translations found for the given template and context.
	DeleteTranslationsByTemplate(ctx context.Context, in *DeleteTranslationsByTemplateRequest, opts ...grpc.CallOption) (*DeleteTranslationsByTemplateResponse, error)
	// Bulk delete translations
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	BulkDeleteTranslations(ctx context.Context, in *BulkDeleteTranslationsRequest, opts ...grpc.CallOption) (*BulkDeleteTranslationsResponse, error)
}

type translationsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTranslationsServiceClient(cc grpc.ClientConnInterface) TranslationsServiceClient {
	return &translationsServiceClient{cc}
}

func (c *translationsServiceClient) TranslateTemplate(ctx context.Context, in *TranslateTemplateRequest, opts ...grpc.CallOption) (*TranslateTemplateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TranslateTemplateResponse)
	err := c.cc.Invoke(ctx, TranslationsService_TranslateTemplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationsServiceClient) ListTranslations(ctx context.Context, in *ListTranslationsRequest, opts ...grpc.CallOption) (*ListTranslationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTranslationsResponse)
	err := c.cc.Invoke(ctx, TranslationsService_ListTranslations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationsServiceClient) ListLanguages(ctx context.Context, in *ListLanguagesRequest, opts ...grpc.CallOption) (*ListLanguagesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListLanguagesResponse)
	err := c.cc.Invoke(ctx, TranslationsService_ListLanguages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationsServiceClient) ListContexts(ctx context.Context, in *ListContextsRequest, opts ...grpc.CallOption) (*ListContextsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListContextsResponse)
	err := c.cc.Invoke(ctx, TranslationsService_ListContexts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationsServiceClient) CreateTranslation(ctx context.Context, in *CreateTranslationRequest, opts ...grpc.CallOption) (*CreateTranslationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTranslationResponse)
	err := c.cc.Invoke(ctx, TranslationsService_CreateTranslation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationsServiceClient) UpdateTranslation(ctx context.Context, in *UpdateTranslationRequest, opts ...grpc.CallOption) (*UpdateTranslationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTranslationResponse)
	err := c.cc.Invoke(ctx, TranslationsService_UpdateTranslation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationsServiceClient) TriggerLLMTranslation(ctx context.Context, in *TriggerLLMTranslationRequest, opts ...grpc.CallOption) (*TriggerLLMTranslationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TriggerLLMTranslationResponse)
	err := c.cc.Invoke(ctx, TranslationsService_TriggerLLMTranslation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationsServiceClient) TriggerLLMTranslations(ctx context.Context, in *TriggerLLMTranslationsRequest, opts ...grpc.CallOption) (*TriggerLLMTranslationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TriggerLLMTranslationsResponse)
	err := c.cc.Invoke(ctx, TranslationsService_TriggerLLMTranslations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationsServiceClient) SetSystemMessage(ctx context.Context, in *SetSystemMessageRequest, opts ...grpc.CallOption) (*SetSystemMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetSystemMessageResponse)
	err := c.cc.Invoke(ctx, TranslationsService_SetSystemMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationsServiceClient) GetSystemMessage(ctx context.Context, in *GetSystemMessageRequest, opts ...grpc.CallOption) (*GetSystemMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSystemMessageResponse)
	err := c.cc.Invoke(ctx, TranslationsService_GetSystemMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationsServiceClient) TestSystemMessage(ctx context.Context, in *TestSystemMessageRequest, opts ...grpc.CallOption) (*TestSystemMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TestSystemMessageResponse)
	err := c.cc.Invoke(ctx, TranslationsService_TestSystemMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationsServiceClient) EnableContext(ctx context.Context, in *EnableContextRequest, opts ...grpc.CallOption) (*EnableContextResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EnableContextResponse)
	err := c.cc.Invoke(ctx, TranslationsService_EnableContext_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationsServiceClient) DisableContext(ctx context.Context, in *DisableContextRequest, opts ...grpc.CallOption) (*DisableContextResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DisableContextResponse)
	err := c.cc.Invoke(ctx, TranslationsService_DisableContext_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationsServiceClient) DeleteTranslationsByTemplate(ctx context.Context, in *DeleteTranslationsByTemplateRequest, opts ...grpc.CallOption) (*DeleteTranslationsByTemplateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteTranslationsByTemplateResponse)
	err := c.cc.Invoke(ctx, TranslationsService_DeleteTranslationsByTemplate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationsServiceClient) BulkDeleteTranslations(ctx context.Context, in *BulkDeleteTranslationsRequest, opts ...grpc.CallOption) (*BulkDeleteTranslationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BulkDeleteTranslationsResponse)
	err := c.cc.Invoke(ctx, TranslationsService_BulkDeleteTranslations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranslationsServiceServer is the server API for TranslationsService service.
// All implementations must embed UnimplementedTranslationsServiceServer
// for forward compatibility.
type TranslationsServiceServer interface {
	// Translate a template for a given context and language.
	// Required permissions:
	//
	//	Any Authenticated User
	//
	// Errors:
	//   - grpc.AlreadyExists : This template is already translated for the given context and language.
	//   - grpc.InvalidArgument: The request is not valid.
	TranslateTemplate(context.Context, *TranslateTemplateRequest) (*TranslateTemplateResponse, error)
	// Lists translations by context/language
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	//   - grpc.NotFound: No templates found for the given context and language.
	ListTranslations(context.Context, *ListTranslationsRequest) (*ListTranslationsResponse, error)
	// Lists localization languages
	// Required permissions:
	//   - Any Authenticated User
	ListLanguages(context.Context, *ListLanguagesRequest) (*ListLanguagesResponse, error)
	// Lists localization contexts
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	ListContexts(context.Context, *ListContextsRequest) (*ListContextsResponse, error)
	// Creates a new Translation
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	//   - grpc.AlreadyExists: The template already exists for the given context and language (use override method).
	CreateTranslation(context.Context, *CreateTranslationRequest) (*CreateTranslationResponse, error)
	// Overrides the translation for a given translationID
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	UpdateTranslation(context.Context, *UpdateTranslationRequest) (*UpdateTranslationResponse, error)
	// Re-run the LLM translation for a given translationID
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	TriggerLLMTranslation(context.Context, *TriggerLLMTranslationRequest) (*TriggerLLMTranslationResponse, error)
	// re-run all translations for a given context (WARNING - this should be ran sparingly as it is a heavy operation and costs money)
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	TriggerLLMTranslations(context.Context, *TriggerLLMTranslationsRequest) (*TriggerLLMTranslationsResponse, error)
	// set/get context system message to give more tuned LLMs when translating for that context (WARNING - this overrides the previous system message for the context if exists)
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	SetSystemMessage(context.Context, *SetSystemMessageRequest) (*SetSystemMessageResponse, error)
	GetSystemMessage(context.Context, *GetSystemMessageRequest) (*GetSystemMessageResponse, error)
	// Gives a translation for a system message, template and language with no side effects (Used for testing system messages)
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	TestSystemMessage(context.Context, *TestSystemMessageRequest) (*TestSystemMessageResponse, error)
	// enable a context for LLM translations
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	//   - grpc.NotFound: No context found
	EnableContext(context.Context, *EnableContextRequest) (*EnableContextResponse, error)
	// disable a context for LLM translations
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	DisableContext(context.Context, *DisableContextRequest) (*DisableContextResponse, error)
	// Delete translations by template and context
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	//   - grpc.NotFound: No translations found for the given template and context.
	DeleteTranslationsByTemplate(context.Context, *DeleteTranslationsByTemplateRequest) (*DeleteTranslationsByTemplateResponse, error)
	// Bulk delete translations
	// Required permissions:
	//   - PERMISSION_CUSTOMER_SUPPORT
	//
	// Errors:
	//   - grpc.InvalidArgument: The request is not valid.
	BulkDeleteTranslations(context.Context, *BulkDeleteTranslationsRequest) (*BulkDeleteTranslationsResponse, error)
	mustEmbedUnimplementedTranslationsServiceServer()
}

// UnimplementedTranslationsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTranslationsServiceServer struct{}

func (UnimplementedTranslationsServiceServer) TranslateTemplate(context.Context, *TranslateTemplateRequest) (*TranslateTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TranslateTemplate not implemented")
}
func (UnimplementedTranslationsServiceServer) ListTranslations(context.Context, *ListTranslationsRequest) (*ListTranslationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTranslations not implemented")
}
func (UnimplementedTranslationsServiceServer) ListLanguages(context.Context, *ListLanguagesRequest) (*ListLanguagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLanguages not implemented")
}
func (UnimplementedTranslationsServiceServer) ListContexts(context.Context, *ListContextsRequest) (*ListContextsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListContexts not implemented")
}
func (UnimplementedTranslationsServiceServer) CreateTranslation(context.Context, *CreateTranslationRequest) (*CreateTranslationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTranslation not implemented")
}
func (UnimplementedTranslationsServiceServer) UpdateTranslation(context.Context, *UpdateTranslationRequest) (*UpdateTranslationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTranslation not implemented")
}
func (UnimplementedTranslationsServiceServer) TriggerLLMTranslation(context.Context, *TriggerLLMTranslationRequest) (*TriggerLLMTranslationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerLLMTranslation not implemented")
}
func (UnimplementedTranslationsServiceServer) TriggerLLMTranslations(context.Context, *TriggerLLMTranslationsRequest) (*TriggerLLMTranslationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TriggerLLMTranslations not implemented")
}
func (UnimplementedTranslationsServiceServer) SetSystemMessage(context.Context, *SetSystemMessageRequest) (*SetSystemMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSystemMessage not implemented")
}
func (UnimplementedTranslationsServiceServer) GetSystemMessage(context.Context, *GetSystemMessageRequest) (*GetSystemMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSystemMessage not implemented")
}
func (UnimplementedTranslationsServiceServer) TestSystemMessage(context.Context, *TestSystemMessageRequest) (*TestSystemMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestSystemMessage not implemented")
}
func (UnimplementedTranslationsServiceServer) EnableContext(context.Context, *EnableContextRequest) (*EnableContextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnableContext not implemented")
}
func (UnimplementedTranslationsServiceServer) DisableContext(context.Context, *DisableContextRequest) (*DisableContextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisableContext not implemented")
}
func (UnimplementedTranslationsServiceServer) DeleteTranslationsByTemplate(context.Context, *DeleteTranslationsByTemplateRequest) (*DeleteTranslationsByTemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTranslationsByTemplate not implemented")
}
func (UnimplementedTranslationsServiceServer) BulkDeleteTranslations(context.Context, *BulkDeleteTranslationsRequest) (*BulkDeleteTranslationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkDeleteTranslations not implemented")
}
func (UnimplementedTranslationsServiceServer) mustEmbedUnimplementedTranslationsServiceServer() {}
func (UnimplementedTranslationsServiceServer) testEmbeddedByValue()                             {}

// UnsafeTranslationsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TranslationsServiceServer will
// result in compilation errors.
type UnsafeTranslationsServiceServer interface {
	mustEmbedUnimplementedTranslationsServiceServer()
}

func RegisterTranslationsServiceServer(s grpc.ServiceRegistrar, srv TranslationsServiceServer) {
	// If the following call pancis, it indicates UnimplementedTranslationsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TranslationsService_ServiceDesc, srv)
}

func _TranslationsService_TranslateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationsServiceServer).TranslateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslationsService_TranslateTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationsServiceServer).TranslateTemplate(ctx, req.(*TranslateTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationsService_ListTranslations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTranslationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationsServiceServer).ListTranslations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslationsService_ListTranslations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationsServiceServer).ListTranslations(ctx, req.(*ListTranslationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationsService_ListLanguages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLanguagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationsServiceServer).ListLanguages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslationsService_ListLanguages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationsServiceServer).ListLanguages(ctx, req.(*ListLanguagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationsService_ListContexts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListContextsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationsServiceServer).ListContexts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslationsService_ListContexts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationsServiceServer).ListContexts(ctx, req.(*ListContextsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationsService_CreateTranslation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTranslationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationsServiceServer).CreateTranslation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslationsService_CreateTranslation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationsServiceServer).CreateTranslation(ctx, req.(*CreateTranslationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationsService_UpdateTranslation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTranslationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationsServiceServer).UpdateTranslation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslationsService_UpdateTranslation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationsServiceServer).UpdateTranslation(ctx, req.(*UpdateTranslationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationsService_TriggerLLMTranslation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerLLMTranslationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationsServiceServer).TriggerLLMTranslation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslationsService_TriggerLLMTranslation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationsServiceServer).TriggerLLMTranslation(ctx, req.(*TriggerLLMTranslationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationsService_TriggerLLMTranslations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TriggerLLMTranslationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationsServiceServer).TriggerLLMTranslations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslationsService_TriggerLLMTranslations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationsServiceServer).TriggerLLMTranslations(ctx, req.(*TriggerLLMTranslationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationsService_SetSystemMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSystemMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationsServiceServer).SetSystemMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslationsService_SetSystemMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationsServiceServer).SetSystemMessage(ctx, req.(*SetSystemMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationsService_GetSystemMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSystemMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationsServiceServer).GetSystemMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslationsService_GetSystemMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationsServiceServer).GetSystemMessage(ctx, req.(*GetSystemMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationsService_TestSystemMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestSystemMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationsServiceServer).TestSystemMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslationsService_TestSystemMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationsServiceServer).TestSystemMessage(ctx, req.(*TestSystemMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationsService_EnableContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationsServiceServer).EnableContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslationsService_EnableContext_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationsServiceServer).EnableContext(ctx, req.(*EnableContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationsService_DisableContext_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableContextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationsServiceServer).DisableContext(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslationsService_DisableContext_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationsServiceServer).DisableContext(ctx, req.(*DisableContextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationsService_DeleteTranslationsByTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTranslationsByTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationsServiceServer).DeleteTranslationsByTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslationsService_DeleteTranslationsByTemplate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationsServiceServer).DeleteTranslationsByTemplate(ctx, req.(*DeleteTranslationsByTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationsService_BulkDeleteTranslations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkDeleteTranslationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationsServiceServer).BulkDeleteTranslations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TranslationsService_BulkDeleteTranslations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationsServiceServer).BulkDeleteTranslations(ctx, req.(*BulkDeleteTranslationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TranslationsService_ServiceDesc is the grpc.ServiceDesc for TranslationsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TranslationsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.translations.v1alpha1.TranslationsService",
	HandlerType: (*TranslationsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TranslateTemplate",
			Handler:    _TranslationsService_TranslateTemplate_Handler,
		},
		{
			MethodName: "ListTranslations",
			Handler:    _TranslationsService_ListTranslations_Handler,
		},
		{
			MethodName: "ListLanguages",
			Handler:    _TranslationsService_ListLanguages_Handler,
		},
		{
			MethodName: "ListContexts",
			Handler:    _TranslationsService_ListContexts_Handler,
		},
		{
			MethodName: "CreateTranslation",
			Handler:    _TranslationsService_CreateTranslation_Handler,
		},
		{
			MethodName: "UpdateTranslation",
			Handler:    _TranslationsService_UpdateTranslation_Handler,
		},
		{
			MethodName: "TriggerLLMTranslation",
			Handler:    _TranslationsService_TriggerLLMTranslation_Handler,
		},
		{
			MethodName: "TriggerLLMTranslations",
			Handler:    _TranslationsService_TriggerLLMTranslations_Handler,
		},
		{
			MethodName: "SetSystemMessage",
			Handler:    _TranslationsService_SetSystemMessage_Handler,
		},
		{
			MethodName: "GetSystemMessage",
			Handler:    _TranslationsService_GetSystemMessage_Handler,
		},
		{
			MethodName: "TestSystemMessage",
			Handler:    _TranslationsService_TestSystemMessage_Handler,
		},
		{
			MethodName: "EnableContext",
			Handler:    _TranslationsService_EnableContext_Handler,
		},
		{
			MethodName: "DisableContext",
			Handler:    _TranslationsService_DisableContext_Handler,
		},
		{
			MethodName: "DeleteTranslationsByTemplate",
			Handler:    _TranslationsService_DeleteTranslationsByTemplate_Handler,
		},
		{
			MethodName: "BulkDeleteTranslations",
			Handler:    _TranslationsService_BulkDeleteTranslations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/translations/v1alpha1/service.proto",
}
