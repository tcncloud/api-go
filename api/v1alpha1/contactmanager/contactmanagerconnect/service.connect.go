// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/v1alpha1/contactmanager/service.proto

package contactmanagerconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	contactmanager "github.com/tcncloud/api-go/api/v1alpha1/contactmanager"
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
	// ContactManagerName is the fully-qualified name of the ContactManager service.
	ContactManagerName = "api.v1alpha1.contactmanager.ContactManager"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ContactManagerGetContactListProcedure is the fully-qualified name of the ContactManager's
	// GetContactList RPC.
	ContactManagerGetContactListProcedure = "/api.v1alpha1.contactmanager.ContactManager/GetContactList"
	// ContactManagerListContactEntryListProcedure is the fully-qualified name of the ContactManager's
	// ListContactEntryList RPC.
	ContactManagerListContactEntryListProcedure = "/api.v1alpha1.contactmanager.ContactManager/ListContactEntryList"
	// ContactManagerGetEncContactEntryProcedure is the fully-qualified name of the ContactManager's
	// GetEncContactEntry RPC.
	ContactManagerGetEncContactEntryProcedure = "/api.v1alpha1.contactmanager.ContactManager/GetEncContactEntry"
	// ContactManagerGetKYCEncContactEntryProcedure is the fully-qualified name of the ContactManager's
	// GetKYCEncContactEntry RPC.
	ContactManagerGetKYCEncContactEntryProcedure = "/api.v1alpha1.contactmanager.ContactManager/GetKYCEncContactEntry"
	// ContactManagerGetKYCKeysProcedure is the fully-qualified name of the ContactManager's GetKYCKeys
	// RPC.
	ContactManagerGetKYCKeysProcedure = "/api.v1alpha1.contactmanager.ContactManager/GetKYCKeys"
	// ContactManagerAddContactEntryProcedure is the fully-qualified name of the ContactManager's
	// AddContactEntry RPC.
	ContactManagerAddContactEntryProcedure = "/api.v1alpha1.contactmanager.ContactManager/AddContactEntry"
	// ContactManagerEditContactEntryProcedure is the fully-qualified name of the ContactManager's
	// EditContactEntry RPC.
	ContactManagerEditContactEntryProcedure = "/api.v1alpha1.contactmanager.ContactManager/EditContactEntry"
	// ContactManagerListContactsByEntityProcedure is the fully-qualified name of the ContactManager's
	// ListContactsByEntity RPC.
	ContactManagerListContactsByEntityProcedure = "/api.v1alpha1.contactmanager.ContactManager/ListContactsByEntity"
	// ContactManagerGetContactFieldTypeProcedure is the fully-qualified name of the ContactManager's
	// GetContactFieldType RPC.
	ContactManagerGetContactFieldTypeProcedure = "/api.v1alpha1.contactmanager.ContactManager/GetContactFieldType"
)

// ContactManagerClient is a client for the api.v1alpha1.contactmanager.ContactManager service.
type ContactManagerClient interface {
	GetContactList(context.Context, *connect_go.Request[contactmanager.GetContactListRequest]) (*connect_go.Response[contactmanager.GetContactListResponse], error)
	ListContactEntryList(context.Context, *connect_go.Request[contactmanager.ListContactEntryListRequest]) (*connect_go.Response[contactmanager.ListContactEntryListResponse], error)
	GetEncContactEntry(context.Context, *connect_go.Request[contactmanager.GetEncContactEntryRequest]) (*connect_go.Response[contactmanager.GetEncContactEntryResponse], error)
	GetKYCEncContactEntry(context.Context, *connect_go.Request[contactmanager.GetKYCEncContactEntryRequest]) (*connect_go.Response[contactmanager.GetKYCEncContactEntryResponse], error)
	GetKYCKeys(context.Context, *connect_go.Request[contactmanager.GetKYCKeysRequest]) (*connect_go.Response[contactmanager.GetKYCKeysResponse], error)
	// *
	// Adds a new contact entry based on the provided request.
	AddContactEntry(context.Context, *connect_go.Request[contactmanager.AddContactEntryRequest]) (*connect_go.Response[contactmanager.AddContactEntryResponse], error)
	// *
	// Edits the fields of an existing contact entry...
	EditContactEntry(context.Context, *connect_go.Request[contactmanager.EditContactEntryRequest]) (*connect_go.Response[contactmanager.EditContactEntryResponse], error)
	// *
	// List contacts for entity
	ListContactsByEntity(context.Context, *connect_go.Request[contactmanager.ListContactsByEntityRequest]) (*connect_go.Response[contactmanager.ListContactsByEntityResponse], error)
	// *
	// Get Contact Field Type
	GetContactFieldType(context.Context, *connect_go.Request[contactmanager.GetContactFieldTypeRequest]) (*connect_go.Response[contactmanager.GetContactFieldTypeResponse], error)
}

// NewContactManagerClient constructs a client for the api.v1alpha1.contactmanager.ContactManager
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewContactManagerClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ContactManagerClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &contactManagerClient{
		getContactList: connect_go.NewClient[contactmanager.GetContactListRequest, contactmanager.GetContactListResponse](
			httpClient,
			baseURL+ContactManagerGetContactListProcedure,
			opts...,
		),
		listContactEntryList: connect_go.NewClient[contactmanager.ListContactEntryListRequest, contactmanager.ListContactEntryListResponse](
			httpClient,
			baseURL+ContactManagerListContactEntryListProcedure,
			opts...,
		),
		getEncContactEntry: connect_go.NewClient[contactmanager.GetEncContactEntryRequest, contactmanager.GetEncContactEntryResponse](
			httpClient,
			baseURL+ContactManagerGetEncContactEntryProcedure,
			opts...,
		),
		getKYCEncContactEntry: connect_go.NewClient[contactmanager.GetKYCEncContactEntryRequest, contactmanager.GetKYCEncContactEntryResponse](
			httpClient,
			baseURL+ContactManagerGetKYCEncContactEntryProcedure,
			opts...,
		),
		getKYCKeys: connect_go.NewClient[contactmanager.GetKYCKeysRequest, contactmanager.GetKYCKeysResponse](
			httpClient,
			baseURL+ContactManagerGetKYCKeysProcedure,
			opts...,
		),
		addContactEntry: connect_go.NewClient[contactmanager.AddContactEntryRequest, contactmanager.AddContactEntryResponse](
			httpClient,
			baseURL+ContactManagerAddContactEntryProcedure,
			opts...,
		),
		editContactEntry: connect_go.NewClient[contactmanager.EditContactEntryRequest, contactmanager.EditContactEntryResponse](
			httpClient,
			baseURL+ContactManagerEditContactEntryProcedure,
			opts...,
		),
		listContactsByEntity: connect_go.NewClient[contactmanager.ListContactsByEntityRequest, contactmanager.ListContactsByEntityResponse](
			httpClient,
			baseURL+ContactManagerListContactsByEntityProcedure,
			opts...,
		),
		getContactFieldType: connect_go.NewClient[contactmanager.GetContactFieldTypeRequest, contactmanager.GetContactFieldTypeResponse](
			httpClient,
			baseURL+ContactManagerGetContactFieldTypeProcedure,
			opts...,
		),
	}
}

// contactManagerClient implements ContactManagerClient.
type contactManagerClient struct {
	getContactList        *connect_go.Client[contactmanager.GetContactListRequest, contactmanager.GetContactListResponse]
	listContactEntryList  *connect_go.Client[contactmanager.ListContactEntryListRequest, contactmanager.ListContactEntryListResponse]
	getEncContactEntry    *connect_go.Client[contactmanager.GetEncContactEntryRequest, contactmanager.GetEncContactEntryResponse]
	getKYCEncContactEntry *connect_go.Client[contactmanager.GetKYCEncContactEntryRequest, contactmanager.GetKYCEncContactEntryResponse]
	getKYCKeys            *connect_go.Client[contactmanager.GetKYCKeysRequest, contactmanager.GetKYCKeysResponse]
	addContactEntry       *connect_go.Client[contactmanager.AddContactEntryRequest, contactmanager.AddContactEntryResponse]
	editContactEntry      *connect_go.Client[contactmanager.EditContactEntryRequest, contactmanager.EditContactEntryResponse]
	listContactsByEntity  *connect_go.Client[contactmanager.ListContactsByEntityRequest, contactmanager.ListContactsByEntityResponse]
	getContactFieldType   *connect_go.Client[contactmanager.GetContactFieldTypeRequest, contactmanager.GetContactFieldTypeResponse]
}

// GetContactList calls api.v1alpha1.contactmanager.ContactManager.GetContactList.
func (c *contactManagerClient) GetContactList(ctx context.Context, req *connect_go.Request[contactmanager.GetContactListRequest]) (*connect_go.Response[contactmanager.GetContactListResponse], error) {
	return c.getContactList.CallUnary(ctx, req)
}

// ListContactEntryList calls api.v1alpha1.contactmanager.ContactManager.ListContactEntryList.
func (c *contactManagerClient) ListContactEntryList(ctx context.Context, req *connect_go.Request[contactmanager.ListContactEntryListRequest]) (*connect_go.Response[contactmanager.ListContactEntryListResponse], error) {
	return c.listContactEntryList.CallUnary(ctx, req)
}

// GetEncContactEntry calls api.v1alpha1.contactmanager.ContactManager.GetEncContactEntry.
func (c *contactManagerClient) GetEncContactEntry(ctx context.Context, req *connect_go.Request[contactmanager.GetEncContactEntryRequest]) (*connect_go.Response[contactmanager.GetEncContactEntryResponse], error) {
	return c.getEncContactEntry.CallUnary(ctx, req)
}

// GetKYCEncContactEntry calls api.v1alpha1.contactmanager.ContactManager.GetKYCEncContactEntry.
func (c *contactManagerClient) GetKYCEncContactEntry(ctx context.Context, req *connect_go.Request[contactmanager.GetKYCEncContactEntryRequest]) (*connect_go.Response[contactmanager.GetKYCEncContactEntryResponse], error) {
	return c.getKYCEncContactEntry.CallUnary(ctx, req)
}

// GetKYCKeys calls api.v1alpha1.contactmanager.ContactManager.GetKYCKeys.
func (c *contactManagerClient) GetKYCKeys(ctx context.Context, req *connect_go.Request[contactmanager.GetKYCKeysRequest]) (*connect_go.Response[contactmanager.GetKYCKeysResponse], error) {
	return c.getKYCKeys.CallUnary(ctx, req)
}

// AddContactEntry calls api.v1alpha1.contactmanager.ContactManager.AddContactEntry.
func (c *contactManagerClient) AddContactEntry(ctx context.Context, req *connect_go.Request[contactmanager.AddContactEntryRequest]) (*connect_go.Response[contactmanager.AddContactEntryResponse], error) {
	return c.addContactEntry.CallUnary(ctx, req)
}

// EditContactEntry calls api.v1alpha1.contactmanager.ContactManager.EditContactEntry.
func (c *contactManagerClient) EditContactEntry(ctx context.Context, req *connect_go.Request[contactmanager.EditContactEntryRequest]) (*connect_go.Response[contactmanager.EditContactEntryResponse], error) {
	return c.editContactEntry.CallUnary(ctx, req)
}

// ListContactsByEntity calls api.v1alpha1.contactmanager.ContactManager.ListContactsByEntity.
func (c *contactManagerClient) ListContactsByEntity(ctx context.Context, req *connect_go.Request[contactmanager.ListContactsByEntityRequest]) (*connect_go.Response[contactmanager.ListContactsByEntityResponse], error) {
	return c.listContactsByEntity.CallUnary(ctx, req)
}

// GetContactFieldType calls api.v1alpha1.contactmanager.ContactManager.GetContactFieldType.
func (c *contactManagerClient) GetContactFieldType(ctx context.Context, req *connect_go.Request[contactmanager.GetContactFieldTypeRequest]) (*connect_go.Response[contactmanager.GetContactFieldTypeResponse], error) {
	return c.getContactFieldType.CallUnary(ctx, req)
}

// ContactManagerHandler is an implementation of the api.v1alpha1.contactmanager.ContactManager
// service.
type ContactManagerHandler interface {
	GetContactList(context.Context, *connect_go.Request[contactmanager.GetContactListRequest]) (*connect_go.Response[contactmanager.GetContactListResponse], error)
	ListContactEntryList(context.Context, *connect_go.Request[contactmanager.ListContactEntryListRequest]) (*connect_go.Response[contactmanager.ListContactEntryListResponse], error)
	GetEncContactEntry(context.Context, *connect_go.Request[contactmanager.GetEncContactEntryRequest]) (*connect_go.Response[contactmanager.GetEncContactEntryResponse], error)
	GetKYCEncContactEntry(context.Context, *connect_go.Request[contactmanager.GetKYCEncContactEntryRequest]) (*connect_go.Response[contactmanager.GetKYCEncContactEntryResponse], error)
	GetKYCKeys(context.Context, *connect_go.Request[contactmanager.GetKYCKeysRequest]) (*connect_go.Response[contactmanager.GetKYCKeysResponse], error)
	// *
	// Adds a new contact entry based on the provided request.
	AddContactEntry(context.Context, *connect_go.Request[contactmanager.AddContactEntryRequest]) (*connect_go.Response[contactmanager.AddContactEntryResponse], error)
	// *
	// Edits the fields of an existing contact entry...
	EditContactEntry(context.Context, *connect_go.Request[contactmanager.EditContactEntryRequest]) (*connect_go.Response[contactmanager.EditContactEntryResponse], error)
	// *
	// List contacts for entity
	ListContactsByEntity(context.Context, *connect_go.Request[contactmanager.ListContactsByEntityRequest]) (*connect_go.Response[contactmanager.ListContactsByEntityResponse], error)
	// *
	// Get Contact Field Type
	GetContactFieldType(context.Context, *connect_go.Request[contactmanager.GetContactFieldTypeRequest]) (*connect_go.Response[contactmanager.GetContactFieldTypeResponse], error)
}

// NewContactManagerHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewContactManagerHandler(svc ContactManagerHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	contactManagerGetContactListHandler := connect_go.NewUnaryHandler(
		ContactManagerGetContactListProcedure,
		svc.GetContactList,
		opts...,
	)
	contactManagerListContactEntryListHandler := connect_go.NewUnaryHandler(
		ContactManagerListContactEntryListProcedure,
		svc.ListContactEntryList,
		opts...,
	)
	contactManagerGetEncContactEntryHandler := connect_go.NewUnaryHandler(
		ContactManagerGetEncContactEntryProcedure,
		svc.GetEncContactEntry,
		opts...,
	)
	contactManagerGetKYCEncContactEntryHandler := connect_go.NewUnaryHandler(
		ContactManagerGetKYCEncContactEntryProcedure,
		svc.GetKYCEncContactEntry,
		opts...,
	)
	contactManagerGetKYCKeysHandler := connect_go.NewUnaryHandler(
		ContactManagerGetKYCKeysProcedure,
		svc.GetKYCKeys,
		opts...,
	)
	contactManagerAddContactEntryHandler := connect_go.NewUnaryHandler(
		ContactManagerAddContactEntryProcedure,
		svc.AddContactEntry,
		opts...,
	)
	contactManagerEditContactEntryHandler := connect_go.NewUnaryHandler(
		ContactManagerEditContactEntryProcedure,
		svc.EditContactEntry,
		opts...,
	)
	contactManagerListContactsByEntityHandler := connect_go.NewUnaryHandler(
		ContactManagerListContactsByEntityProcedure,
		svc.ListContactsByEntity,
		opts...,
	)
	contactManagerGetContactFieldTypeHandler := connect_go.NewUnaryHandler(
		ContactManagerGetContactFieldTypeProcedure,
		svc.GetContactFieldType,
		opts...,
	)
	return "/api.v1alpha1.contactmanager.ContactManager/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ContactManagerGetContactListProcedure:
			contactManagerGetContactListHandler.ServeHTTP(w, r)
		case ContactManagerListContactEntryListProcedure:
			contactManagerListContactEntryListHandler.ServeHTTP(w, r)
		case ContactManagerGetEncContactEntryProcedure:
			contactManagerGetEncContactEntryHandler.ServeHTTP(w, r)
		case ContactManagerGetKYCEncContactEntryProcedure:
			contactManagerGetKYCEncContactEntryHandler.ServeHTTP(w, r)
		case ContactManagerGetKYCKeysProcedure:
			contactManagerGetKYCKeysHandler.ServeHTTP(w, r)
		case ContactManagerAddContactEntryProcedure:
			contactManagerAddContactEntryHandler.ServeHTTP(w, r)
		case ContactManagerEditContactEntryProcedure:
			contactManagerEditContactEntryHandler.ServeHTTP(w, r)
		case ContactManagerListContactsByEntityProcedure:
			contactManagerListContactsByEntityHandler.ServeHTTP(w, r)
		case ContactManagerGetContactFieldTypeProcedure:
			contactManagerGetContactFieldTypeHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedContactManagerHandler returns CodeUnimplemented from all methods.
type UnimplementedContactManagerHandler struct{}

func (UnimplementedContactManagerHandler) GetContactList(context.Context, *connect_go.Request[contactmanager.GetContactListRequest]) (*connect_go.Response[contactmanager.GetContactListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.contactmanager.ContactManager.GetContactList is not implemented"))
}

func (UnimplementedContactManagerHandler) ListContactEntryList(context.Context, *connect_go.Request[contactmanager.ListContactEntryListRequest]) (*connect_go.Response[contactmanager.ListContactEntryListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.contactmanager.ContactManager.ListContactEntryList is not implemented"))
}

func (UnimplementedContactManagerHandler) GetEncContactEntry(context.Context, *connect_go.Request[contactmanager.GetEncContactEntryRequest]) (*connect_go.Response[contactmanager.GetEncContactEntryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.contactmanager.ContactManager.GetEncContactEntry is not implemented"))
}

func (UnimplementedContactManagerHandler) GetKYCEncContactEntry(context.Context, *connect_go.Request[contactmanager.GetKYCEncContactEntryRequest]) (*connect_go.Response[contactmanager.GetKYCEncContactEntryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.contactmanager.ContactManager.GetKYCEncContactEntry is not implemented"))
}

func (UnimplementedContactManagerHandler) GetKYCKeys(context.Context, *connect_go.Request[contactmanager.GetKYCKeysRequest]) (*connect_go.Response[contactmanager.GetKYCKeysResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.contactmanager.ContactManager.GetKYCKeys is not implemented"))
}

func (UnimplementedContactManagerHandler) AddContactEntry(context.Context, *connect_go.Request[contactmanager.AddContactEntryRequest]) (*connect_go.Response[contactmanager.AddContactEntryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.contactmanager.ContactManager.AddContactEntry is not implemented"))
}

func (UnimplementedContactManagerHandler) EditContactEntry(context.Context, *connect_go.Request[contactmanager.EditContactEntryRequest]) (*connect_go.Response[contactmanager.EditContactEntryResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.contactmanager.ContactManager.EditContactEntry is not implemented"))
}

func (UnimplementedContactManagerHandler) ListContactsByEntity(context.Context, *connect_go.Request[contactmanager.ListContactsByEntityRequest]) (*connect_go.Response[contactmanager.ListContactsByEntityResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.contactmanager.ContactManager.ListContactsByEntity is not implemented"))
}

func (UnimplementedContactManagerHandler) GetContactFieldType(context.Context, *connect_go.Request[contactmanager.GetContactFieldTypeRequest]) (*connect_go.Response[contactmanager.GetContactFieldTypeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.v1alpha1.contactmanager.ContactManager.GetContactFieldType is not implemented"))
}
