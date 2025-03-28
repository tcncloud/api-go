// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/v1alpha1/org/trusts/entities.proto

package trusts

import (
	auth "github.com/tcncloud/api-go/api/commons/auth"
	org "github.com/tcncloud/api-go/api/commons/org"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for the CreateTrust RPC.
type CreateTrustRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Org ID of the org recieving the trust.
	Grantee string `protobuf:"bytes,1,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// Name of the trust.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the trust.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Permissions given through the trust
	Permissions []auth.Permission `protobuf:"varint,4,rep,packed,name=permissions,proto3,enum=api.commons.auth.Permission" json:"permissions,omitempty"`
	// IDs of labels associated with the given permissions.
	// NOTE: All given labels are applied to ALL given permissions.
	LabelIds      []string `protobuf:"bytes,5,rep,name=label_ids,json=labelIds,proto3" json:"label_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTrustRequest) Reset() {
	*x = CreateTrustRequest{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTrustRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTrustRequest) ProtoMessage() {}

func (x *CreateTrustRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTrustRequest.ProtoReflect.Descriptor instead.
func (*CreateTrustRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTrustRequest) GetGrantee() string {
	if x != nil {
		return x.Grantee
	}
	return ""
}

func (x *CreateTrustRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTrustRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateTrustRequest) GetPermissions() []auth.Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *CreateTrustRequest) GetLabelIds() []string {
	if x != nil {
		return x.LabelIds
	}
	return nil
}

// Response message for the CreateTrust RPC.
type CreateTrustResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the created trust.
	TrustId       string `protobuf:"bytes,1,opt,name=trust_id,json=trustId,proto3" json:"trust_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateTrustResponse) Reset() {
	*x = CreateTrustResponse{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTrustResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTrustResponse) ProtoMessage() {}

func (x *CreateTrustResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTrustResponse.ProtoReflect.Descriptor instead.
func (*CreateTrustResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTrustResponse) GetTrustId() string {
	if x != nil {
		return x.TrustId
	}
	return ""
}

// Request message for the AcceptTrust RPC.
type AcceptTrustRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the trust being accepted.
	TrustId       string `protobuf:"bytes,1,opt,name=trust_id,json=trustId,proto3" json:"trust_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AcceptTrustRequest) Reset() {
	*x = AcceptTrustRequest{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AcceptTrustRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptTrustRequest) ProtoMessage() {}

func (x *AcceptTrustRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptTrustRequest.ProtoReflect.Descriptor instead.
func (*AcceptTrustRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{2}
}

func (x *AcceptTrustRequest) GetTrustId() string {
	if x != nil {
		return x.TrustId
	}
	return ""
}

// Response message for the AcceptTrust RPC.
type AcceptTrustResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AcceptTrustResponse) Reset() {
	*x = AcceptTrustResponse{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AcceptTrustResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptTrustResponse) ProtoMessage() {}

func (x *AcceptTrustResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptTrustResponse.ProtoReflect.Descriptor instead.
func (*AcceptTrustResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{3}
}

// Request message for the RejectTrust RPC.
type RejectTrustRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the trust being rejected.
	TrustId       string `protobuf:"bytes,1,opt,name=trust_id,json=trustId,proto3" json:"trust_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RejectTrustRequest) Reset() {
	*x = RejectTrustRequest{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RejectTrustRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectTrustRequest) ProtoMessage() {}

func (x *RejectTrustRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectTrustRequest.ProtoReflect.Descriptor instead.
func (*RejectTrustRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{4}
}

func (x *RejectTrustRequest) GetTrustId() string {
	if x != nil {
		return x.TrustId
	}
	return ""
}

// Response message for the RejectTrust RPC.
type RejectTrustResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RejectTrustResponse) Reset() {
	*x = RejectTrustResponse{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RejectTrustResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectTrustResponse) ProtoMessage() {}

func (x *RejectTrustResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectTrustResponse.ProtoReflect.Descriptor instead.
func (*RejectTrustResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{5}
}

// Response message for the GetTrust RPC.
type GetTrustRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the trust being retrieved.
	TrustId       string `protobuf:"bytes,1,opt,name=trust_id,json=trustId,proto3" json:"trust_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTrustRequest) Reset() {
	*x = GetTrustRequest{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTrustRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrustRequest) ProtoMessage() {}

func (x *GetTrustRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrustRequest.ProtoReflect.Descriptor instead.
func (*GetTrustRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{6}
}

func (x *GetTrustRequest) GetTrustId() string {
	if x != nil {
		return x.TrustId
	}
	return ""
}

// Response message for the GetTrust RPC.
type GetTrustResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Retrieved trust.
	Trust *org.Trust `protobuf:"bytes,1,opt,name=trust,proto3" json:"trust,omitempty"`
	// Name of the Grantor org.
	GrantorName string `protobuf:"bytes,2,opt,name=grantor_name,json=grantorName,proto3" json:"grantor_name,omitempty"`
	// Name of the Grantee org.
	GranteeName   string `protobuf:"bytes,3,opt,name=grantee_name,json=granteeName,proto3" json:"grantee_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTrustResponse) Reset() {
	*x = GetTrustResponse{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTrustResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTrustResponse) ProtoMessage() {}

func (x *GetTrustResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTrustResponse.ProtoReflect.Descriptor instead.
func (*GetTrustResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{7}
}

func (x *GetTrustResponse) GetTrust() *org.Trust {
	if x != nil {
		return x.Trust
	}
	return nil
}

func (x *GetTrustResponse) GetGrantorName() string {
	if x != nil {
		return x.GrantorName
	}
	return ""
}

func (x *GetTrustResponse) GetGranteeName() string {
	if x != nil {
		return x.GranteeName
	}
	return ""
}

// Payload for list trust responses
type ListTrustsResponsePayload struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Trust response
	Trust *org.Trust `protobuf:"bytes,1,opt,name=trust,proto3" json:"trust,omitempty"`
	// Grantor of Trust
	GrantorName string `protobuf:"bytes,2,opt,name=grantor_name,json=grantorName,proto3" json:"grantor_name,omitempty"`
	// Grantee of Trust
	GranteeName   string `protobuf:"bytes,3,opt,name=grantee_name,json=granteeName,proto3" json:"grantee_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListTrustsResponsePayload) Reset() {
	*x = ListTrustsResponsePayload{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListTrustsResponsePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTrustsResponsePayload) ProtoMessage() {}

func (x *ListTrustsResponsePayload) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTrustsResponsePayload.ProtoReflect.Descriptor instead.
func (*ListTrustsResponsePayload) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{8}
}

func (x *ListTrustsResponsePayload) GetTrust() *org.Trust {
	if x != nil {
		return x.Trust
	}
	return nil
}

func (x *ListTrustsResponsePayload) GetGrantorName() string {
	if x != nil {
		return x.GrantorName
	}
	return ""
}

func (x *ListTrustsResponsePayload) GetGranteeName() string {
	if x != nil {
		return x.GranteeName
	}
	return ""
}

// Request message for the ListIncomingTrusts RPC.
type ListIncomingTrustsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListIncomingTrustsRequest) Reset() {
	*x = ListIncomingTrustsRequest{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListIncomingTrustsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIncomingTrustsRequest) ProtoMessage() {}

func (x *ListIncomingTrustsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIncomingTrustsRequest.ProtoReflect.Descriptor instead.
func (*ListIncomingTrustsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{9}
}

// Response message for the ListIncomingTrusts RPC.
type ListIncomingTrustsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of trusts.
	Trusts        []*ListTrustsResponsePayload `protobuf:"bytes,1,rep,name=trusts,proto3" json:"trusts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListIncomingTrustsResponse) Reset() {
	*x = ListIncomingTrustsResponse{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListIncomingTrustsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIncomingTrustsResponse) ProtoMessage() {}

func (x *ListIncomingTrustsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIncomingTrustsResponse.ProtoReflect.Descriptor instead.
func (*ListIncomingTrustsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{10}
}

func (x *ListIncomingTrustsResponse) GetTrusts() []*ListTrustsResponsePayload {
	if x != nil {
		return x.Trusts
	}
	return nil
}

// Request message for the ListGivenTrusts RPC.
type ListGivenTrustsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListGivenTrustsRequest) Reset() {
	*x = ListGivenTrustsRequest{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListGivenTrustsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGivenTrustsRequest) ProtoMessage() {}

func (x *ListGivenTrustsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGivenTrustsRequest.ProtoReflect.Descriptor instead.
func (*ListGivenTrustsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{11}
}

// Response message for the ListGivenTrusts RPC.
type ListGivenTrustsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of trusts.
	Trusts        []*ListTrustsResponsePayload `protobuf:"bytes,1,rep,name=trusts,proto3" json:"trusts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListGivenTrustsResponse) Reset() {
	*x = ListGivenTrustsResponse{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListGivenTrustsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGivenTrustsResponse) ProtoMessage() {}

func (x *ListGivenTrustsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGivenTrustsResponse.ProtoReflect.Descriptor instead.
func (*ListGivenTrustsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{12}
}

func (x *ListGivenTrustsResponse) GetTrusts() []*ListTrustsResponsePayload {
	if x != nil {
		return x.Trusts
	}
	return nil
}

// Request message for the CreateTrust RPC.
type ListAssignableTrustsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAssignableTrustsRequest) Reset() {
	*x = ListAssignableTrustsRequest{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAssignableTrustsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssignableTrustsRequest) ProtoMessage() {}

func (x *ListAssignableTrustsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssignableTrustsRequest.ProtoReflect.Descriptor instead.
func (*ListAssignableTrustsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{13}
}

// Response message for the CreateTrust RPC.
type ListAssignableTrustsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of trusts.
	Trusts        []*ListTrustsResponsePayload `protobuf:"bytes,1,rep,name=trusts,proto3" json:"trusts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAssignableTrustsResponse) Reset() {
	*x = ListAssignableTrustsResponse{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAssignableTrustsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssignableTrustsResponse) ProtoMessage() {}

func (x *ListAssignableTrustsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssignableTrustsResponse.ProtoReflect.Descriptor instead.
func (*ListAssignableTrustsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{14}
}

func (x *ListAssignableTrustsResponse) GetTrusts() []*ListTrustsResponsePayload {
	if x != nil {
		return x.Trusts
	}
	return nil
}

// Request message for the DeleteTrust RPC.
type DeleteTrustRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the trust to be deleted.
	TrustId       string `protobuf:"bytes,1,opt,name=trust_id,json=trustId,proto3" json:"trust_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTrustRequest) Reset() {
	*x = DeleteTrustRequest{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTrustRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTrustRequest) ProtoMessage() {}

func (x *DeleteTrustRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTrustRequest.ProtoReflect.Descriptor instead.
func (*DeleteTrustRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{15}
}

func (x *DeleteTrustRequest) GetTrustId() string {
	if x != nil {
		return x.TrustId
	}
	return ""
}

// Response message for the DeleteTrust RPC.
type DeleteTrustResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteTrustResponse) Reset() {
	*x = DeleteTrustResponse{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTrustResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTrustResponse) ProtoMessage() {}

func (x *DeleteTrustResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTrustResponse.ProtoReflect.Descriptor instead.
func (*DeleteTrustResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{16}
}

// Request message for the AssignTrust RPC.
type AssignTrustRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the trust being assigned.
	TrustId string `protobuf:"bytes,1,opt,name=trust_id,json=trustId,proto3" json:"trust_id,omitempty"`
	// List of user IDs to assign the trust to.
	UserIds       []string `protobuf:"bytes,2,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssignTrustRequest) Reset() {
	*x = AssignTrustRequest{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[17]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssignTrustRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignTrustRequest) ProtoMessage() {}

func (x *AssignTrustRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[17]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignTrustRequest.ProtoReflect.Descriptor instead.
func (*AssignTrustRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{17}
}

func (x *AssignTrustRequest) GetTrustId() string {
	if x != nil {
		return x.TrustId
	}
	return ""
}

func (x *AssignTrustRequest) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

// Response message for the AssignTrust RPC.
type AssignTrustResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssignTrustResponse) Reset() {
	*x = AssignTrustResponse{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[18]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssignTrustResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignTrustResponse) ProtoMessage() {}

func (x *AssignTrustResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[18]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignTrustResponse.ProtoReflect.Descriptor instead.
func (*AssignTrustResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{18}
}

// Request message for the UnassignTrust RPC.
type UnassignTrustRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the trust being unassigned.
	TrustId string `protobuf:"bytes,1,opt,name=trust_id,json=trustId,proto3" json:"trust_id,omitempty"`
	// User ID to unassign the trust from.
	UserId        string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnassignTrustRequest) Reset() {
	*x = UnassignTrustRequest{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[19]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnassignTrustRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnassignTrustRequest) ProtoMessage() {}

func (x *UnassignTrustRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[19]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnassignTrustRequest.ProtoReflect.Descriptor instead.
func (*UnassignTrustRequest) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{19}
}

func (x *UnassignTrustRequest) GetTrustId() string {
	if x != nil {
		return x.TrustId
	}
	return ""
}

func (x *UnassignTrustRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Response message for the UnassignTrust RPC.
type UnassignTrustResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UnassignTrustResponse) Reset() {
	*x = UnassignTrustResponse{}
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[20]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnassignTrustResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnassignTrustResponse) ProtoMessage() {}

func (x *UnassignTrustResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1alpha1_org_trusts_entities_proto_msgTypes[20]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnassignTrustResponse.ProtoReflect.Descriptor instead.
func (*UnassignTrustResponse) Descriptor() ([]byte, []int) {
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP(), []int{20}
}

var File_api_v1alpha1_org_trusts_entities_proto protoreflect.FileDescriptor

const file_api_v1alpha1_org_trusts_entities_proto_rawDesc = "" +
	"\n" +
	"&api/v1alpha1/org/trusts/entities.proto\x12\x17api.v1alpha1.org.trusts\x1a\x1capi/commons/auth/perms.proto\x1a\x1capi/commons/org/trusts.proto\"\xc1\x01\n" +
	"\x12CreateTrustRequest\x12\x18\n" +
	"\agrantee\x18\x01 \x01(\tR\agrantee\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12>\n" +
	"\vpermissions\x18\x04 \x03(\x0e2\x1c.api.commons.auth.PermissionR\vpermissions\x12\x1b\n" +
	"\tlabel_ids\x18\x05 \x03(\tR\blabelIds\"0\n" +
	"\x13CreateTrustResponse\x12\x19\n" +
	"\btrust_id\x18\x01 \x01(\tR\atrustId\"/\n" +
	"\x12AcceptTrustRequest\x12\x19\n" +
	"\btrust_id\x18\x01 \x01(\tR\atrustId\"\x15\n" +
	"\x13AcceptTrustResponse\"/\n" +
	"\x12RejectTrustRequest\x12\x19\n" +
	"\btrust_id\x18\x01 \x01(\tR\atrustId\"\x15\n" +
	"\x13RejectTrustResponse\",\n" +
	"\x0fGetTrustRequest\x12\x19\n" +
	"\btrust_id\x18\x01 \x01(\tR\atrustId\"\x86\x01\n" +
	"\x10GetTrustResponse\x12,\n" +
	"\x05trust\x18\x01 \x01(\v2\x16.api.commons.org.TrustR\x05trust\x12!\n" +
	"\fgrantor_name\x18\x02 \x01(\tR\vgrantorName\x12!\n" +
	"\fgrantee_name\x18\x03 \x01(\tR\vgranteeName\"\x8f\x01\n" +
	"\x19ListTrustsResponsePayload\x12,\n" +
	"\x05trust\x18\x01 \x01(\v2\x16.api.commons.org.TrustR\x05trust\x12!\n" +
	"\fgrantor_name\x18\x02 \x01(\tR\vgrantorName\x12!\n" +
	"\fgrantee_name\x18\x03 \x01(\tR\vgranteeName\"\x1b\n" +
	"\x19ListIncomingTrustsRequest\"h\n" +
	"\x1aListIncomingTrustsResponse\x12J\n" +
	"\x06trusts\x18\x01 \x03(\v22.api.v1alpha1.org.trusts.ListTrustsResponsePayloadR\x06trusts\"\x18\n" +
	"\x16ListGivenTrustsRequest\"e\n" +
	"\x17ListGivenTrustsResponse\x12J\n" +
	"\x06trusts\x18\x01 \x03(\v22.api.v1alpha1.org.trusts.ListTrustsResponsePayloadR\x06trusts\"\x1d\n" +
	"\x1bListAssignableTrustsRequest\"j\n" +
	"\x1cListAssignableTrustsResponse\x12J\n" +
	"\x06trusts\x18\x01 \x03(\v22.api.v1alpha1.org.trusts.ListTrustsResponsePayloadR\x06trusts\"/\n" +
	"\x12DeleteTrustRequest\x12\x19\n" +
	"\btrust_id\x18\x01 \x01(\tR\atrustId\"\x15\n" +
	"\x13DeleteTrustResponse\"J\n" +
	"\x12AssignTrustRequest\x12\x19\n" +
	"\btrust_id\x18\x01 \x01(\tR\atrustId\x12\x19\n" +
	"\buser_ids\x18\x02 \x03(\tR\auserIds\"\x15\n" +
	"\x13AssignTrustResponse\"J\n" +
	"\x14UnassignTrustRequest\x12\x19\n" +
	"\btrust_id\x18\x01 \x01(\tR\atrustId\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\"\x17\n" +
	"\x15UnassignTrustResponseB\xe0\x01\n" +
	"\x1bcom.api.v1alpha1.org.trustsB\rEntitiesProtoP\x01Z2github.com/tcncloud/api-go/api/v1alpha1/org/trusts\xa2\x02\x04AVOT\xaa\x02\x17Api.V1alpha1.Org.Trusts\xca\x02\x17Api\\V1alpha1\\Org\\Trusts\xe2\x02#Api\\V1alpha1\\Org\\Trusts\\GPBMetadata\xea\x02\x1aApi::V1alpha1::Org::Trustsb\x06proto3"

var (
	file_api_v1alpha1_org_trusts_entities_proto_rawDescOnce sync.Once
	file_api_v1alpha1_org_trusts_entities_proto_rawDescData []byte
)

func file_api_v1alpha1_org_trusts_entities_proto_rawDescGZIP() []byte {
	file_api_v1alpha1_org_trusts_entities_proto_rawDescOnce.Do(func() {
		file_api_v1alpha1_org_trusts_entities_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1alpha1_org_trusts_entities_proto_rawDesc), len(file_api_v1alpha1_org_trusts_entities_proto_rawDesc)))
	})
	return file_api_v1alpha1_org_trusts_entities_proto_rawDescData
}

var file_api_v1alpha1_org_trusts_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 21)
var file_api_v1alpha1_org_trusts_entities_proto_goTypes = []any{
	(*CreateTrustRequest)(nil),           // 0: api.v1alpha1.org.trusts.CreateTrustRequest
	(*CreateTrustResponse)(nil),          // 1: api.v1alpha1.org.trusts.CreateTrustResponse
	(*AcceptTrustRequest)(nil),           // 2: api.v1alpha1.org.trusts.AcceptTrustRequest
	(*AcceptTrustResponse)(nil),          // 3: api.v1alpha1.org.trusts.AcceptTrustResponse
	(*RejectTrustRequest)(nil),           // 4: api.v1alpha1.org.trusts.RejectTrustRequest
	(*RejectTrustResponse)(nil),          // 5: api.v1alpha1.org.trusts.RejectTrustResponse
	(*GetTrustRequest)(nil),              // 6: api.v1alpha1.org.trusts.GetTrustRequest
	(*GetTrustResponse)(nil),             // 7: api.v1alpha1.org.trusts.GetTrustResponse
	(*ListTrustsResponsePayload)(nil),    // 8: api.v1alpha1.org.trusts.ListTrustsResponsePayload
	(*ListIncomingTrustsRequest)(nil),    // 9: api.v1alpha1.org.trusts.ListIncomingTrustsRequest
	(*ListIncomingTrustsResponse)(nil),   // 10: api.v1alpha1.org.trusts.ListIncomingTrustsResponse
	(*ListGivenTrustsRequest)(nil),       // 11: api.v1alpha1.org.trusts.ListGivenTrustsRequest
	(*ListGivenTrustsResponse)(nil),      // 12: api.v1alpha1.org.trusts.ListGivenTrustsResponse
	(*ListAssignableTrustsRequest)(nil),  // 13: api.v1alpha1.org.trusts.ListAssignableTrustsRequest
	(*ListAssignableTrustsResponse)(nil), // 14: api.v1alpha1.org.trusts.ListAssignableTrustsResponse
	(*DeleteTrustRequest)(nil),           // 15: api.v1alpha1.org.trusts.DeleteTrustRequest
	(*DeleteTrustResponse)(nil),          // 16: api.v1alpha1.org.trusts.DeleteTrustResponse
	(*AssignTrustRequest)(nil),           // 17: api.v1alpha1.org.trusts.AssignTrustRequest
	(*AssignTrustResponse)(nil),          // 18: api.v1alpha1.org.trusts.AssignTrustResponse
	(*UnassignTrustRequest)(nil),         // 19: api.v1alpha1.org.trusts.UnassignTrustRequest
	(*UnassignTrustResponse)(nil),        // 20: api.v1alpha1.org.trusts.UnassignTrustResponse
	(auth.Permission)(0),                 // 21: api.commons.auth.Permission
	(*org.Trust)(nil),                    // 22: api.commons.org.Trust
}
var file_api_v1alpha1_org_trusts_entities_proto_depIdxs = []int32{
	21, // 0: api.v1alpha1.org.trusts.CreateTrustRequest.permissions:type_name -> api.commons.auth.Permission
	22, // 1: api.v1alpha1.org.trusts.GetTrustResponse.trust:type_name -> api.commons.org.Trust
	22, // 2: api.v1alpha1.org.trusts.ListTrustsResponsePayload.trust:type_name -> api.commons.org.Trust
	8,  // 3: api.v1alpha1.org.trusts.ListIncomingTrustsResponse.trusts:type_name -> api.v1alpha1.org.trusts.ListTrustsResponsePayload
	8,  // 4: api.v1alpha1.org.trusts.ListGivenTrustsResponse.trusts:type_name -> api.v1alpha1.org.trusts.ListTrustsResponsePayload
	8,  // 5: api.v1alpha1.org.trusts.ListAssignableTrustsResponse.trusts:type_name -> api.v1alpha1.org.trusts.ListTrustsResponsePayload
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_v1alpha1_org_trusts_entities_proto_init() }
func file_api_v1alpha1_org_trusts_entities_proto_init() {
	if File_api_v1alpha1_org_trusts_entities_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1alpha1_org_trusts_entities_proto_rawDesc), len(file_api_v1alpha1_org_trusts_entities_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   21,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_v1alpha1_org_trusts_entities_proto_goTypes,
		DependencyIndexes: file_api_v1alpha1_org_trusts_entities_proto_depIdxs,
		MessageInfos:      file_api_v1alpha1_org_trusts_entities_proto_msgTypes,
	}.Build()
	File_api_v1alpha1_org_trusts_entities_proto = out.File
	file_api_v1alpha1_org_trusts_entities_proto_goTypes = nil
	file_api_v1alpha1_org_trusts_entities_proto_depIdxs = nil
}
