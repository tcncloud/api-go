// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/commons/org/user.proto

package org

import (
	commons "github.com/tcncloud/api-go/api/commons"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

// User represents a user of the system.
type User struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ID of the user
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// The ID of the org the user belongs to.
	OrgId string `protobuf:"bytes,2,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// The Username of the user.
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// The ID of the p3 permission group the user has.
	P3PermissionGroupId string `protobuf:"bytes,4,opt,name=p3_permission_group_id,json=p3PermissionGroupId,proto3" json:"p3_permission_group_id,omitempty"`
	// The p3 login sid of the user.
	// This is going to be deprecated, RegionSidMap should be used.
	LoginSid int64 `protobuf:"varint,5,opt,name=login_sid,json=loginSid,proto3" json:"login_sid,omitempty"`
	// The p3 agent sid of the user.
	// This is going to be deprecated, RegionSidMap should be used.
	AgentSid int64 `protobuf:"varint,6,opt,name=agent_sid,json=agentSid,proto3" json:"agent_sid,omitempty"`
	// The region the user belongs to.
	// This is going to be deprecated, RegionSidMap should be used.
	RegionId string `protobuf:"bytes,7,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// The User's partner agent id.
	PartnerAgentId string `protobuf:"bytes,8,opt,name=partner_agent_id,json=partnerAgentId,proto3" json:"partner_agent_id,omitempty"`
	// The p3 client sid of the user.
	ClientSid int64 `protobuf:"varint,11,opt,name=client_sid,json=clientSid,proto3" json:"client_sid,omitempty"`
	// The user's regional sids. The key of the map is a region id.
	RegionSidMap map[string]*User_RegionSids `protobuf:"bytes,15,rep,name=region_sid_map,json=regionSidMap,proto3" json:"region_sid_map,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// payload
	ApiKey string `protobuf:"bytes,101,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	// The User's email.
	Email string `protobuf:"bytes,104,opt,name=email,proto3" json:"email,omitempty"`
	// Whether or not the user is disabled.
	LoginDisabled bool `protobuf:"varint,112,opt,name=login_disabled,json=loginDisabled,proto3" json:"login_disabled,omitempty"`
	// login preferences
	CallerIds       []string `protobuf:"bytes,115,rep,name=caller_ids,json=callerIds,proto3" json:"caller_ids,omitempty"`
	LinkbackNumbers []string `protobuf:"bytes,116,rep,name=linkback_numbers,json=linkbackNumbers,proto3" json:"linkback_numbers,omitempty"`
	AuthUserId      string   `protobuf:"bytes,118,opt,name=auth_user_id,json=authUserId,proto3" json:"auth_user_id,omitempty"`
	EnableMfa       bool     `protobuf:"varint,123,opt,name=enable_mfa,json=enableMfa,proto3" json:"enable_mfa,omitempty"`
	// The user's first name.
	FirstName string `protobuf:"bytes,161,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	// THe user's last name.
	LastName string `protobuf:"bytes,162,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	// The date/time the user was created.
	Created *timestamppb.Timestamp `protobuf:"bytes,163,opt,name=created,proto3" json:"created,omitempty"`
	// The date/time the user was last updated.
	LastUpdated *timestamppb.Timestamp `protobuf:"bytes,164,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
	// Whether or not the user must reset their password on next login.
	PasswordResetRequired bool `protobuf:"varint,165,opt,name=password_reset_required,json=passwordResetRequired,proto3" json:"password_reset_required,omitempty"`
	// connection id is the id of the auth connection that the
	// user is coming from if it is a delgated user. If the user
	// is not delgated this will be nil.
	ConnectionId *wrapperspb.StringValue `protobuf:"bytes,166,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	// User TimeZone with wrapper for nullable timezone
	TimeZoneOverride *commons.TimeZoneWrapper `protobuf:"bytes,167,opt,name=time_zone_override,json=timeZoneOverride,proto3" json:"time_zone_override,omitempty"`
	// UserSettings
	PermissionGroupIds []string `protobuf:"bytes,200,rep,name=permission_group_ids,json=permissionGroupIds,proto3" json:"permission_group_ids,omitempty"`
	// List of trust ids the user has been assigned.
	TrustIds []string `protobuf:"bytes,209,rep,name=trust_ids,json=trustIds,proto3" json:"trust_ids,omitempty"`
	// List of regions for user
	// repeated string regions = 201;
	// region_id of a user entity
	DefaultRegion string `protobuf:"bytes,202,opt,name=default_region,json=defaultRegion,proto3" json:"default_region,omitempty"`
	// default_app for a user
	DefaultApplication commons.OperatorApplications `protobuf:"varint,203,opt,name=default_application,json=defaultApplication,proto3,enum=api.commons.OperatorApplications" json:"default_application,omitempty"`
	// P3 permission for user
	// User caller id
	UserCallerId string `protobuf:"bytes,205,opt,name=user_caller_id,json=userCallerId,proto3" json:"user_caller_id,omitempty"`
	// users assigned agent profile group
	AgentProfileGroupId string `protobuf:"bytes,207,opt,name=agent_profile_group_id,json=agentProfileGroupId,proto3" json:"agent_profile_group_id,omitempty"`
	// The skills this user has as an agent, if any.
	Skills []*Skill `protobuf:"bytes,208,rep,name=skills,proto3" json:"skills,omitempty"`
	// a bool signalling whether or not the user is also an agent
	Agent bool `protobuf:"varint,300,opt,name=agent,proto3" json:"agent,omitempty"`
	// a bool to determine whether or not user is account owner
	AccountOwner bool `protobuf:"varint,400,opt,name=account_owner,json=accountOwner,proto3" json:"account_owner,omitempty"`
	// Whether or not the user's email is verified.
	EmailVerified bool `protobuf:"varint,401,opt,name=email_verified,json=emailVerified,proto3" json:"email_verified,omitempty"`
	// Whitelisted IPs for the user
	WhitelistedIps []string `protobuf:"bytes,402,rep,name=whitelisted_ips,json=whitelistedIps,proto3" json:"whitelisted_ips,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_api_commons_org_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_commons_org_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetP3PermissionGroupId() string {
	if x != nil {
		return x.P3PermissionGroupId
	}
	return ""
}

func (x *User) GetLoginSid() int64 {
	if x != nil {
		return x.LoginSid
	}
	return 0
}

func (x *User) GetAgentSid() int64 {
	if x != nil {
		return x.AgentSid
	}
	return 0
}

func (x *User) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *User) GetPartnerAgentId() string {
	if x != nil {
		return x.PartnerAgentId
	}
	return ""
}

func (x *User) GetClientSid() int64 {
	if x != nil {
		return x.ClientSid
	}
	return 0
}

func (x *User) GetRegionSidMap() map[string]*User_RegionSids {
	if x != nil {
		return x.RegionSidMap
	}
	return nil
}

func (x *User) GetApiKey() string {
	if x != nil {
		return x.ApiKey
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetLoginDisabled() bool {
	if x != nil {
		return x.LoginDisabled
	}
	return false
}

func (x *User) GetCallerIds() []string {
	if x != nil {
		return x.CallerIds
	}
	return nil
}

func (x *User) GetLinkbackNumbers() []string {
	if x != nil {
		return x.LinkbackNumbers
	}
	return nil
}

func (x *User) GetAuthUserId() string {
	if x != nil {
		return x.AuthUserId
	}
	return ""
}

func (x *User) GetEnableMfa() bool {
	if x != nil {
		return x.EnableMfa
	}
	return false
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetCreated() *timestamppb.Timestamp {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *User) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

func (x *User) GetPasswordResetRequired() bool {
	if x != nil {
		return x.PasswordResetRequired
	}
	return false
}

func (x *User) GetConnectionId() *wrapperspb.StringValue {
	if x != nil {
		return x.ConnectionId
	}
	return nil
}

func (x *User) GetTimeZoneOverride() *commons.TimeZoneWrapper {
	if x != nil {
		return x.TimeZoneOverride
	}
	return nil
}

func (x *User) GetPermissionGroupIds() []string {
	if x != nil {
		return x.PermissionGroupIds
	}
	return nil
}

func (x *User) GetTrustIds() []string {
	if x != nil {
		return x.TrustIds
	}
	return nil
}

func (x *User) GetDefaultRegion() string {
	if x != nil {
		return x.DefaultRegion
	}
	return ""
}

func (x *User) GetDefaultApplication() commons.OperatorApplications {
	if x != nil {
		return x.DefaultApplication
	}
	return commons.OperatorApplications(0)
}

func (x *User) GetUserCallerId() string {
	if x != nil {
		return x.UserCallerId
	}
	return ""
}

func (x *User) GetAgentProfileGroupId() string {
	if x != nil {
		return x.AgentProfileGroupId
	}
	return ""
}

func (x *User) GetSkills() []*Skill {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *User) GetAgent() bool {
	if x != nil {
		return x.Agent
	}
	return false
}

func (x *User) GetAccountOwner() bool {
	if x != nil {
		return x.AccountOwner
	}
	return false
}

func (x *User) GetEmailVerified() bool {
	if x != nil {
		return x.EmailVerified
	}
	return false
}

func (x *User) GetWhitelistedIps() []string {
	if x != nil {
		return x.WhitelistedIps
	}
	return nil
}

// MFA/2FA Information
type MfaInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The org the user belongs to.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// The user that the MFA info belongs to.
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Whether or not MFA is enabled for the user.
	MfaEnabled bool `protobuf:"varint,3,opt,name=mfa_enabled,json=mfaEnabled,proto3" json:"mfa_enabled,omitempty"`
	// Types that are valid to be assigned to MfaType:
	//
	//	*MfaInfo_None
	//	*MfaInfo_Otp
	//	*MfaInfo_Duo_
	//	*MfaInfo_Totp_
	MfaType       isMfaInfo_MfaType `protobuf_oneof:"mfa_type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MfaInfo) Reset() {
	*x = MfaInfo{}
	mi := &file_api_commons_org_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MfaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MfaInfo) ProtoMessage() {}

func (x *MfaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MfaInfo.ProtoReflect.Descriptor instead.
func (*MfaInfo) Descriptor() ([]byte, []int) {
	return file_api_commons_org_user_proto_rawDescGZIP(), []int{1}
}

func (x *MfaInfo) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *MfaInfo) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *MfaInfo) GetMfaEnabled() bool {
	if x != nil {
		return x.MfaEnabled
	}
	return false
}

func (x *MfaInfo) GetMfaType() isMfaInfo_MfaType {
	if x != nil {
		return x.MfaType
	}
	return nil
}

func (x *MfaInfo) GetNone() *MfaInfo_NoneSelected {
	if x != nil {
		if x, ok := x.MfaType.(*MfaInfo_None); ok {
			return x.None
		}
	}
	return nil
}

func (x *MfaInfo) GetOtp() *MfaInfo_OtpType {
	if x != nil {
		if x, ok := x.MfaType.(*MfaInfo_Otp); ok {
			return x.Otp
		}
	}
	return nil
}

func (x *MfaInfo) GetDuo() *MfaInfo_Duo {
	if x != nil {
		if x, ok := x.MfaType.(*MfaInfo_Duo_); ok {
			return x.Duo
		}
	}
	return nil
}

func (x *MfaInfo) GetTotp() *MfaInfo_Totp {
	if x != nil {
		if x, ok := x.MfaType.(*MfaInfo_Totp_); ok {
			return x.Totp
		}
	}
	return nil
}

type isMfaInfo_MfaType interface {
	isMfaInfo_MfaType()
}

type MfaInfo_None struct {
	None *MfaInfo_NoneSelected `protobuf:"bytes,10,opt,name=none,proto3,oneof"`
}

type MfaInfo_Otp struct {
	Otp *MfaInfo_OtpType `protobuf:"bytes,11,opt,name=otp,proto3,oneof"`
}

type MfaInfo_Duo_ struct {
	Duo *MfaInfo_Duo `protobuf:"bytes,12,opt,name=duo,proto3,oneof"`
}

type MfaInfo_Totp_ struct {
	Totp *MfaInfo_Totp `protobuf:"bytes,13,opt,name=totp,proto3,oneof"`
}

func (*MfaInfo_None) isMfaInfo_MfaType() {}

func (*MfaInfo_Otp) isMfaInfo_MfaType() {}

func (*MfaInfo_Duo_) isMfaInfo_MfaType() {}

func (*MfaInfo_Totp_) isMfaInfo_MfaType() {}

// Agent skill
type Skill struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Level         int64                  `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	SkillSid      int64                  `protobuf:"varint,4,opt,name=skill_sid,json=skillSid,proto3" json:"skill_sid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Skill) Reset() {
	*x = Skill{}
	mi := &file_api_commons_org_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Skill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Skill) ProtoMessage() {}

func (x *Skill) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Skill.ProtoReflect.Descriptor instead.
func (*Skill) Descriptor() ([]byte, []int) {
	return file_api_commons_org_user_proto_rawDescGZIP(), []int{2}
}

func (x *Skill) GetLevel() int64 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Skill) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Skill) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Skill) GetSkillSid() int64 {
	if x != nil {
		return x.SkillSid
	}
	return 0
}

// PasswordResetLink defines a link for the given user to reset their password.
type PasswordResetLink struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LinkId        string                 `protobuf:"bytes,1,opt,name=link_id,json=linkId,proto3" json:"link_id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrgId         string                 `protobuf:"bytes,3,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	Expiration    *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expiration,proto3" json:"expiration,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PasswordResetLink) Reset() {
	*x = PasswordResetLink{}
	mi := &file_api_commons_org_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PasswordResetLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasswordResetLink) ProtoMessage() {}

func (x *PasswordResetLink) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasswordResetLink.ProtoReflect.Descriptor instead.
func (*PasswordResetLink) Descriptor() ([]byte, []int) {
	return file_api_commons_org_user_proto_rawDescGZIP(), []int{3}
}

func (x *PasswordResetLink) GetLinkId() string {
	if x != nil {
		return x.LinkId
	}
	return ""
}

func (x *PasswordResetLink) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PasswordResetLink) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *PasswordResetLink) GetExpiration() *timestamppb.Timestamp {
	if x != nil {
		return x.Expiration
	}
	return nil
}

// The entity object used in the region_sid_map.
type User_RegionSids struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The login sid for the region.
	LoginSid int64 `protobuf:"varint,1,opt,name=login_sid,json=loginSid,proto3" json:"login_sid,omitempty"`
	// THe agent sid for the region.
	AgentSid int64 `protobuf:"varint,2,opt,name=agent_sid,json=agentSid,proto3" json:"agent_sid,omitempty"`
	// THe client sid for the region.
	ClientSid     int64 `protobuf:"varint,3,opt,name=client_sid,json=clientSid,proto3" json:"client_sid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User_RegionSids) Reset() {
	*x = User_RegionSids{}
	mi := &file_api_commons_org_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User_RegionSids) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User_RegionSids) ProtoMessage() {}

func (x *User_RegionSids) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User_RegionSids.ProtoReflect.Descriptor instead.
func (*User_RegionSids) Descriptor() ([]byte, []int) {
	return file_api_commons_org_user_proto_rawDescGZIP(), []int{0, 1}
}

func (x *User_RegionSids) GetLoginSid() int64 {
	if x != nil {
		return x.LoginSid
	}
	return 0
}

func (x *User_RegionSids) GetAgentSid() int64 {
	if x != nil {
		return x.AgentSid
	}
	return 0
}

func (x *User_RegionSids) GetClientSid() int64 {
	if x != nil {
		return x.ClientSid
	}
	return 0
}

type MfaInfo_NoneSelected struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The time by wich the user must setup MFA
	// to avoid being locked out.
	Timeout       *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MfaInfo_NoneSelected) Reset() {
	*x = MfaInfo_NoneSelected{}
	mi := &file_api_commons_org_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MfaInfo_NoneSelected) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MfaInfo_NoneSelected) ProtoMessage() {}

func (x *MfaInfo_NoneSelected) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MfaInfo_NoneSelected.ProtoReflect.Descriptor instead.
func (*MfaInfo_NoneSelected) Descriptor() ([]byte, []int) {
	return file_api_commons_org_user_proto_rawDescGZIP(), []int{1, 0}
}

func (x *MfaInfo_NoneSelected) GetTimeout() *timestamppb.Timestamp {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type MfaInfo_OtpType struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to DeliveryMethod:
	//
	//	*MfaInfo_OtpType_Email
	DeliveryMethod isMfaInfo_OtpType_DeliveryMethod `protobuf_oneof:"delivery_method"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *MfaInfo_OtpType) Reset() {
	*x = MfaInfo_OtpType{}
	mi := &file_api_commons_org_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MfaInfo_OtpType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MfaInfo_OtpType) ProtoMessage() {}

func (x *MfaInfo_OtpType) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MfaInfo_OtpType.ProtoReflect.Descriptor instead.
func (*MfaInfo_OtpType) Descriptor() ([]byte, []int) {
	return file_api_commons_org_user_proto_rawDescGZIP(), []int{1, 1}
}

func (x *MfaInfo_OtpType) GetDeliveryMethod() isMfaInfo_OtpType_DeliveryMethod {
	if x != nil {
		return x.DeliveryMethod
	}
	return nil
}

func (x *MfaInfo_OtpType) GetEmail() *MfaInfo_OtpType_EmailDeliveryMethod {
	if x != nil {
		if x, ok := x.DeliveryMethod.(*MfaInfo_OtpType_Email); ok {
			return x.Email
		}
	}
	return nil
}

type isMfaInfo_OtpType_DeliveryMethod interface {
	isMfaInfo_OtpType_DeliveryMethod()
}

type MfaInfo_OtpType_Email struct {
	Email *MfaInfo_OtpType_EmailDeliveryMethod `protobuf:"bytes,1,opt,name=email,proto3,oneof"`
}

func (*MfaInfo_OtpType_Email) isMfaInfo_OtpType_DeliveryMethod() {}

type MfaInfo_Duo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DuoUsername   string                 `protobuf:"bytes,1,opt,name=duo_username,json=duoUsername,proto3" json:"duo_username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MfaInfo_Duo) Reset() {
	*x = MfaInfo_Duo{}
	mi := &file_api_commons_org_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MfaInfo_Duo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MfaInfo_Duo) ProtoMessage() {}

func (x *MfaInfo_Duo) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MfaInfo_Duo.ProtoReflect.Descriptor instead.
func (*MfaInfo_Duo) Descriptor() ([]byte, []int) {
	return file_api_commons_org_user_proto_rawDescGZIP(), []int{1, 2}
}

func (x *MfaInfo_Duo) GetDuoUsername() string {
	if x != nil {
		return x.DuoUsername
	}
	return ""
}

// Totp is the type of MFA that uses a time-based one-time password.
type MfaInfo_Totp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MfaInfo_Totp) Reset() {
	*x = MfaInfo_Totp{}
	mi := &file_api_commons_org_user_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MfaInfo_Totp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MfaInfo_Totp) ProtoMessage() {}

func (x *MfaInfo_Totp) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_user_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MfaInfo_Totp.ProtoReflect.Descriptor instead.
func (*MfaInfo_Totp) Descriptor() ([]byte, []int) {
	return file_api_commons_org_user_proto_rawDescGZIP(), []int{1, 3}
}

type MfaInfo_OtpType_EmailDeliveryMethod struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MfaInfo_OtpType_EmailDeliveryMethod) Reset() {
	*x = MfaInfo_OtpType_EmailDeliveryMethod{}
	mi := &file_api_commons_org_user_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MfaInfo_OtpType_EmailDeliveryMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MfaInfo_OtpType_EmailDeliveryMethod) ProtoMessage() {}

func (x *MfaInfo_OtpType_EmailDeliveryMethod) ProtoReflect() protoreflect.Message {
	mi := &file_api_commons_org_user_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MfaInfo_OtpType_EmailDeliveryMethod.ProtoReflect.Descriptor instead.
func (*MfaInfo_OtpType_EmailDeliveryMethod) Descriptor() ([]byte, []int) {
	return file_api_commons_org_user_proto_rawDescGZIP(), []int{1, 1, 0}
}

var File_api_commons_org_user_proto protoreflect.FileDescriptor

const file_api_commons_org_user_proto_rawDesc = "" +
	"\n" +
	"\x1aapi/commons/org/user.proto\x12\x0fapi.commons.org\x1a\x15api/commons/org.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\"\x8b\r\n" +
	"\x04User\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x15\n" +
	"\x06org_id\x18\x02 \x01(\tR\x05orgId\x12\x1a\n" +
	"\busername\x18\x03 \x01(\tR\busername\x123\n" +
	"\x16p3_permission_group_id\x18\x04 \x01(\tR\x13p3PermissionGroupId\x12\x1b\n" +
	"\tlogin_sid\x18\x05 \x01(\x03R\bloginSid\x12\x1b\n" +
	"\tagent_sid\x18\x06 \x01(\x03R\bagentSid\x12\x1b\n" +
	"\tregion_id\x18\a \x01(\tR\bregionId\x12(\n" +
	"\x10partner_agent_id\x18\b \x01(\tR\x0epartnerAgentId\x12\x1d\n" +
	"\n" +
	"client_sid\x18\v \x01(\x03R\tclientSid\x12M\n" +
	"\x0eregion_sid_map\x18\x0f \x03(\v2'.api.commons.org.User.RegionSidMapEntryR\fregionSidMap\x12\x17\n" +
	"\aapi_key\x18e \x01(\tR\x06apiKey\x12\x14\n" +
	"\x05email\x18h \x01(\tR\x05email\x12%\n" +
	"\x0elogin_disabled\x18p \x01(\bR\rloginDisabled\x12\x1d\n" +
	"\n" +
	"caller_ids\x18s \x03(\tR\tcallerIds\x12)\n" +
	"\x10linkback_numbers\x18t \x03(\tR\x0flinkbackNumbers\x12 \n" +
	"\fauth_user_id\x18v \x01(\tR\n" +
	"authUserId\x12\x1d\n" +
	"\n" +
	"enable_mfa\x18{ \x01(\bR\tenableMfa\x12\x1e\n" +
	"\n" +
	"first_name\x18\xa1\x01 \x01(\tR\tfirstName\x12\x1c\n" +
	"\tlast_name\x18\xa2\x01 \x01(\tR\blastName\x125\n" +
	"\acreated\x18\xa3\x01 \x01(\v2\x1a.google.protobuf.TimestampR\acreated\x12>\n" +
	"\flast_updated\x18\xa4\x01 \x01(\v2\x1a.google.protobuf.TimestampR\vlastUpdated\x127\n" +
	"\x17password_reset_required\x18\xa5\x01 \x01(\bR\x15passwordResetRequired\x12B\n" +
	"\rconnection_id\x18\xa6\x01 \x01(\v2\x1c.google.protobuf.StringValueR\fconnectionId\x12K\n" +
	"\x12time_zone_override\x18\xa7\x01 \x01(\v2\x1c.api.commons.TimeZoneWrapperR\x10timeZoneOverride\x121\n" +
	"\x14permission_group_ids\x18\xc8\x01 \x03(\tR\x12permissionGroupIds\x12\x1c\n" +
	"\ttrust_ids\x18\xd1\x01 \x03(\tR\btrustIds\x12&\n" +
	"\x0edefault_region\x18\xca\x01 \x01(\tR\rdefaultRegion\x12S\n" +
	"\x13default_application\x18\xcb\x01 \x01(\x0e2!.api.commons.OperatorApplicationsR\x12defaultApplication\x12%\n" +
	"\x0euser_caller_id\x18\xcd\x01 \x01(\tR\fuserCallerId\x124\n" +
	"\x16agent_profile_group_id\x18\xcf\x01 \x01(\tR\x13agentProfileGroupId\x12/\n" +
	"\x06skills\x18\xd0\x01 \x03(\v2\x16.api.commons.org.SkillR\x06skills\x12\x15\n" +
	"\x05agent\x18\xac\x02 \x01(\bR\x05agent\x12$\n" +
	"\raccount_owner\x18\x90\x03 \x01(\bR\faccountOwner\x12&\n" +
	"\x0eemail_verified\x18\x91\x03 \x01(\bR\remailVerified\x12(\n" +
	"\x0fwhitelisted_ips\x18\x92\x03 \x03(\tR\x0ewhitelistedIps\x1aa\n" +
	"\x11RegionSidMapEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x126\n" +
	"\x05value\x18\x02 \x01(\v2 .api.commons.org.User.RegionSidsR\x05value:\x028\x01\x1ae\n" +
	"\n" +
	"RegionSids\x12\x1b\n" +
	"\tlogin_sid\x18\x01 \x01(\x03R\bloginSid\x12\x1b\n" +
	"\tagent_sid\x18\x02 \x01(\x03R\bagentSid\x12\x1d\n" +
	"\n" +
	"client_sid\x18\x03 \x01(\x03R\tclientSid\"\xbc\x04\n" +
	"\aMfaInfo\x12\x15\n" +
	"\x06org_id\x18\x01 \x01(\tR\x05orgId\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x1f\n" +
	"\vmfa_enabled\x18\x03 \x01(\bR\n" +
	"mfaEnabled\x12;\n" +
	"\x04none\x18\n" +
	" \x01(\v2%.api.commons.org.MfaInfo.NoneSelectedH\x00R\x04none\x124\n" +
	"\x03otp\x18\v \x01(\v2 .api.commons.org.MfaInfo.OtpTypeH\x00R\x03otp\x120\n" +
	"\x03duo\x18\f \x01(\v2\x1c.api.commons.org.MfaInfo.DuoH\x00R\x03duo\x123\n" +
	"\x04totp\x18\r \x01(\v2\x1d.api.commons.org.MfaInfo.TotpH\x00R\x04totp\x1aD\n" +
	"\fNoneSelected\x124\n" +
	"\atimeout\x18\x01 \x01(\v2\x1a.google.protobuf.TimestampR\atimeout\x1a\x81\x01\n" +
	"\aOtpType\x12L\n" +
	"\x05email\x18\x01 \x01(\v24.api.commons.org.MfaInfo.OtpType.EmailDeliveryMethodH\x00R\x05email\x1a\x15\n" +
	"\x13EmailDeliveryMethodB\x11\n" +
	"\x0fdelivery_method\x1a(\n" +
	"\x03Duo\x12!\n" +
	"\fduo_username\x18\x01 \x01(\tR\vduoUsername\x1a\x06\n" +
	"\x04TotpB\n" +
	"\n" +
	"\bmfa_type\"p\n" +
	"\x05Skill\x12\x14\n" +
	"\x05level\x18\x01 \x01(\x03R\x05level\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x1b\n" +
	"\tskill_sid\x18\x04 \x01(\x03R\bskillSid\"\x98\x01\n" +
	"\x11PasswordResetLink\x12\x17\n" +
	"\alink_id\x18\x01 \x01(\tR\x06linkId\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\tR\x06userId\x12\x15\n" +
	"\x06org_id\x18\x03 \x01(\tR\x05orgId\x12:\n" +
	"\n" +
	"expiration\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"expirationB\xaa\x01\n" +
	"\x13com.api.commons.orgB\tUserProtoP\x01Z*github.com/tcncloud/api-go/api/commons/org\xa2\x02\x03ACO\xaa\x02\x0fApi.Commons.Org\xca\x02\x0fApi\\Commons\\Org\xe2\x02\x1bApi\\Commons\\Org\\GPBMetadata\xea\x02\x11Api::Commons::Orgb\x06proto3"

var (
	file_api_commons_org_user_proto_rawDescOnce sync.Once
	file_api_commons_org_user_proto_rawDescData []byte
)

func file_api_commons_org_user_proto_rawDescGZIP() []byte {
	file_api_commons_org_user_proto_rawDescOnce.Do(func() {
		file_api_commons_org_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_commons_org_user_proto_rawDesc), len(file_api_commons_org_user_proto_rawDesc)))
	})
	return file_api_commons_org_user_proto_rawDescData
}

var file_api_commons_org_user_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_commons_org_user_proto_goTypes = []any{
	(*User)(nil),                 // 0: api.commons.org.User
	(*MfaInfo)(nil),              // 1: api.commons.org.MfaInfo
	(*Skill)(nil),                // 2: api.commons.org.Skill
	(*PasswordResetLink)(nil),    // 3: api.commons.org.PasswordResetLink
	nil,                          // 4: api.commons.org.User.RegionSidMapEntry
	(*User_RegionSids)(nil),      // 5: api.commons.org.User.RegionSids
	(*MfaInfo_NoneSelected)(nil), // 6: api.commons.org.MfaInfo.NoneSelected
	(*MfaInfo_OtpType)(nil),      // 7: api.commons.org.MfaInfo.OtpType
	(*MfaInfo_Duo)(nil),          // 8: api.commons.org.MfaInfo.Duo
	(*MfaInfo_Totp)(nil),         // 9: api.commons.org.MfaInfo.Totp
	(*MfaInfo_OtpType_EmailDeliveryMethod)(nil), // 10: api.commons.org.MfaInfo.OtpType.EmailDeliveryMethod
	(*timestamppb.Timestamp)(nil),               // 11: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil),              // 12: google.protobuf.StringValue
	(*commons.TimeZoneWrapper)(nil),             // 13: api.commons.TimeZoneWrapper
	(commons.OperatorApplications)(0),           // 14: api.commons.OperatorApplications
}
var file_api_commons_org_user_proto_depIdxs = []int32{
	4,  // 0: api.commons.org.User.region_sid_map:type_name -> api.commons.org.User.RegionSidMapEntry
	11, // 1: api.commons.org.User.created:type_name -> google.protobuf.Timestamp
	11, // 2: api.commons.org.User.last_updated:type_name -> google.protobuf.Timestamp
	12, // 3: api.commons.org.User.connection_id:type_name -> google.protobuf.StringValue
	13, // 4: api.commons.org.User.time_zone_override:type_name -> api.commons.TimeZoneWrapper
	14, // 5: api.commons.org.User.default_application:type_name -> api.commons.OperatorApplications
	2,  // 6: api.commons.org.User.skills:type_name -> api.commons.org.Skill
	6,  // 7: api.commons.org.MfaInfo.none:type_name -> api.commons.org.MfaInfo.NoneSelected
	7,  // 8: api.commons.org.MfaInfo.otp:type_name -> api.commons.org.MfaInfo.OtpType
	8,  // 9: api.commons.org.MfaInfo.duo:type_name -> api.commons.org.MfaInfo.Duo
	9,  // 10: api.commons.org.MfaInfo.totp:type_name -> api.commons.org.MfaInfo.Totp
	11, // 11: api.commons.org.PasswordResetLink.expiration:type_name -> google.protobuf.Timestamp
	5,  // 12: api.commons.org.User.RegionSidMapEntry.value:type_name -> api.commons.org.User.RegionSids
	11, // 13: api.commons.org.MfaInfo.NoneSelected.timeout:type_name -> google.protobuf.Timestamp
	10, // 14: api.commons.org.MfaInfo.OtpType.email:type_name -> api.commons.org.MfaInfo.OtpType.EmailDeliveryMethod
	15, // [15:15] is the sub-list for method output_type
	15, // [15:15] is the sub-list for method input_type
	15, // [15:15] is the sub-list for extension type_name
	15, // [15:15] is the sub-list for extension extendee
	0,  // [0:15] is the sub-list for field type_name
}

func init() { file_api_commons_org_user_proto_init() }
func file_api_commons_org_user_proto_init() {
	if File_api_commons_org_user_proto != nil {
		return
	}
	file_api_commons_org_user_proto_msgTypes[1].OneofWrappers = []any{
		(*MfaInfo_None)(nil),
		(*MfaInfo_Otp)(nil),
		(*MfaInfo_Duo_)(nil),
		(*MfaInfo_Totp_)(nil),
	}
	file_api_commons_org_user_proto_msgTypes[7].OneofWrappers = []any{
		(*MfaInfo_OtpType_Email)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_commons_org_user_proto_rawDesc), len(file_api_commons_org_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_commons_org_user_proto_goTypes,
		DependencyIndexes: file_api_commons_org_user_proto_depIdxs,
		MessageInfos:      file_api_commons_org_user_proto_msgTypes,
	}.Build()
	File_api_commons_org_user_proto = out.File
	file_api_commons_org_user_proto_goTypes = nil
	file_api_commons_org_user_proto_depIdxs = nil
}
