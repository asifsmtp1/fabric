// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/lifecycle/lifecycle.proto

package lifecycle // import "github.com/hyperledger/fabric/protos/peer/lifecycle"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// InstallChaincodeArgs is the message used as the argument to
// '+lifecycle.InstallChaincode'
type InstallChaincodeArgs struct {
	Name                    string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version                 string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	ChaincodeInstallPackage []byte   `protobuf:"bytes,3,opt,name=chaincode_install_package,json=chaincodeInstallPackage,proto3" json:"chaincode_install_package,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *InstallChaincodeArgs) Reset()         { *m = InstallChaincodeArgs{} }
func (m *InstallChaincodeArgs) String() string { return proto.CompactTextString(m) }
func (*InstallChaincodeArgs) ProtoMessage()    {}
func (*InstallChaincodeArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_lifecycle_7e2bec26dbec8f7e, []int{0}
}
func (m *InstallChaincodeArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallChaincodeArgs.Unmarshal(m, b)
}
func (m *InstallChaincodeArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallChaincodeArgs.Marshal(b, m, deterministic)
}
func (dst *InstallChaincodeArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallChaincodeArgs.Merge(dst, src)
}
func (m *InstallChaincodeArgs) XXX_Size() int {
	return xxx_messageInfo_InstallChaincodeArgs.Size(m)
}
func (m *InstallChaincodeArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallChaincodeArgs.DiscardUnknown(m)
}

var xxx_messageInfo_InstallChaincodeArgs proto.InternalMessageInfo

func (m *InstallChaincodeArgs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InstallChaincodeArgs) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *InstallChaincodeArgs) GetChaincodeInstallPackage() []byte {
	if m != nil {
		return m.ChaincodeInstallPackage
	}
	return nil
}

// InstallChaincodeArgs is the message returned by
// '+lifecycle.InstallChaincode'
type InstallChaincodeResult struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstallChaincodeResult) Reset()         { *m = InstallChaincodeResult{} }
func (m *InstallChaincodeResult) String() string { return proto.CompactTextString(m) }
func (*InstallChaincodeResult) ProtoMessage()    {}
func (*InstallChaincodeResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_lifecycle_7e2bec26dbec8f7e, []int{1}
}
func (m *InstallChaincodeResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstallChaincodeResult.Unmarshal(m, b)
}
func (m *InstallChaincodeResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstallChaincodeResult.Marshal(b, m, deterministic)
}
func (dst *InstallChaincodeResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstallChaincodeResult.Merge(dst, src)
}
func (m *InstallChaincodeResult) XXX_Size() int {
	return xxx_messageInfo_InstallChaincodeResult.Size(m)
}
func (m *InstallChaincodeResult) XXX_DiscardUnknown() {
	xxx_messageInfo_InstallChaincodeResult.DiscardUnknown(m)
}

var xxx_messageInfo_InstallChaincodeResult proto.InternalMessageInfo

func (m *InstallChaincodeResult) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// QueryInstalledChaincodeArgs is the message used as arguemtns
// '+lifecycle.QueryInstalledChaincode'
type QueryInstalledChaincodeArgs struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryInstalledChaincodeArgs) Reset()         { *m = QueryInstalledChaincodeArgs{} }
func (m *QueryInstalledChaincodeArgs) String() string { return proto.CompactTextString(m) }
func (*QueryInstalledChaincodeArgs) ProtoMessage()    {}
func (*QueryInstalledChaincodeArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_lifecycle_7e2bec26dbec8f7e, []int{2}
}
func (m *QueryInstalledChaincodeArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryInstalledChaincodeArgs.Unmarshal(m, b)
}
func (m *QueryInstalledChaincodeArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryInstalledChaincodeArgs.Marshal(b, m, deterministic)
}
func (dst *QueryInstalledChaincodeArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryInstalledChaincodeArgs.Merge(dst, src)
}
func (m *QueryInstalledChaincodeArgs) XXX_Size() int {
	return xxx_messageInfo_QueryInstalledChaincodeArgs.Size(m)
}
func (m *QueryInstalledChaincodeArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryInstalledChaincodeArgs.DiscardUnknown(m)
}

var xxx_messageInfo_QueryInstalledChaincodeArgs proto.InternalMessageInfo

func (m *QueryInstalledChaincodeArgs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *QueryInstalledChaincodeArgs) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// QueryInstalledChaincodeResult is the message returned by
// '+lifecycle.QueryInstalledChaincode'
type QueryInstalledChaincodeResult struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryInstalledChaincodeResult) Reset()         { *m = QueryInstalledChaincodeResult{} }
func (m *QueryInstalledChaincodeResult) String() string { return proto.CompactTextString(m) }
func (*QueryInstalledChaincodeResult) ProtoMessage()    {}
func (*QueryInstalledChaincodeResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_lifecycle_7e2bec26dbec8f7e, []int{3}
}
func (m *QueryInstalledChaincodeResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryInstalledChaincodeResult.Unmarshal(m, b)
}
func (m *QueryInstalledChaincodeResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryInstalledChaincodeResult.Marshal(b, m, deterministic)
}
func (dst *QueryInstalledChaincodeResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryInstalledChaincodeResult.Merge(dst, src)
}
func (m *QueryInstalledChaincodeResult) XXX_Size() int {
	return xxx_messageInfo_QueryInstalledChaincodeResult.Size(m)
}
func (m *QueryInstalledChaincodeResult) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryInstalledChaincodeResult.DiscardUnknown(m)
}

var xxx_messageInfo_QueryInstalledChaincodeResult proto.InternalMessageInfo

func (m *QueryInstalledChaincodeResult) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// QueryInstalledChaincodesArgs currently is an empty argument to
// '+lifecycle.QueryInstalledChaincodes'.   In the future, it may be
// extended to have parameters.
type QueryInstalledChaincodesArgs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryInstalledChaincodesArgs) Reset()         { *m = QueryInstalledChaincodesArgs{} }
func (m *QueryInstalledChaincodesArgs) String() string { return proto.CompactTextString(m) }
func (*QueryInstalledChaincodesArgs) ProtoMessage()    {}
func (*QueryInstalledChaincodesArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_lifecycle_7e2bec26dbec8f7e, []int{4}
}
func (m *QueryInstalledChaincodesArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryInstalledChaincodesArgs.Unmarshal(m, b)
}
func (m *QueryInstalledChaincodesArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryInstalledChaincodesArgs.Marshal(b, m, deterministic)
}
func (dst *QueryInstalledChaincodesArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryInstalledChaincodesArgs.Merge(dst, src)
}
func (m *QueryInstalledChaincodesArgs) XXX_Size() int {
	return xxx_messageInfo_QueryInstalledChaincodesArgs.Size(m)
}
func (m *QueryInstalledChaincodesArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryInstalledChaincodesArgs.DiscardUnknown(m)
}

var xxx_messageInfo_QueryInstalledChaincodesArgs proto.InternalMessageInfo

// QueryInstalledChaincodesResult is the message returned by
// '+lifecycle.QueryInstalledChaincodes'.  It returns a list of
// installed chaincodes.
type QueryInstalledChaincodesResult struct {
	InstalledChaincodes  []*QueryInstalledChaincodesResult_InstalledChaincode `protobuf:"bytes,1,rep,name=installed_chaincodes,json=installedChaincodes,proto3" json:"installed_chaincodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                             `json:"-"`
	XXX_unrecognized     []byte                                               `json:"-"`
	XXX_sizecache        int32                                                `json:"-"`
}

func (m *QueryInstalledChaincodesResult) Reset()         { *m = QueryInstalledChaincodesResult{} }
func (m *QueryInstalledChaincodesResult) String() string { return proto.CompactTextString(m) }
func (*QueryInstalledChaincodesResult) ProtoMessage()    {}
func (*QueryInstalledChaincodesResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_lifecycle_7e2bec26dbec8f7e, []int{5}
}
func (m *QueryInstalledChaincodesResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryInstalledChaincodesResult.Unmarshal(m, b)
}
func (m *QueryInstalledChaincodesResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryInstalledChaincodesResult.Marshal(b, m, deterministic)
}
func (dst *QueryInstalledChaincodesResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryInstalledChaincodesResult.Merge(dst, src)
}
func (m *QueryInstalledChaincodesResult) XXX_Size() int {
	return xxx_messageInfo_QueryInstalledChaincodesResult.Size(m)
}
func (m *QueryInstalledChaincodesResult) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryInstalledChaincodesResult.DiscardUnknown(m)
}

var xxx_messageInfo_QueryInstalledChaincodesResult proto.InternalMessageInfo

func (m *QueryInstalledChaincodesResult) GetInstalledChaincodes() []*QueryInstalledChaincodesResult_InstalledChaincode {
	if m != nil {
		return m.InstalledChaincodes
	}
	return nil
}

type QueryInstalledChaincodesResult_InstalledChaincode struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Hash                 []byte   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryInstalledChaincodesResult_InstalledChaincode) Reset() {
	*m = QueryInstalledChaincodesResult_InstalledChaincode{}
}
func (m *QueryInstalledChaincodesResult_InstalledChaincode) String() string {
	return proto.CompactTextString(m)
}
func (*QueryInstalledChaincodesResult_InstalledChaincode) ProtoMessage() {}
func (*QueryInstalledChaincodesResult_InstalledChaincode) Descriptor() ([]byte, []int) {
	return fileDescriptor_lifecycle_7e2bec26dbec8f7e, []int{5, 0}
}
func (m *QueryInstalledChaincodesResult_InstalledChaincode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryInstalledChaincodesResult_InstalledChaincode.Unmarshal(m, b)
}
func (m *QueryInstalledChaincodesResult_InstalledChaincode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryInstalledChaincodesResult_InstalledChaincode.Marshal(b, m, deterministic)
}
func (dst *QueryInstalledChaincodesResult_InstalledChaincode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryInstalledChaincodesResult_InstalledChaincode.Merge(dst, src)
}
func (m *QueryInstalledChaincodesResult_InstalledChaincode) XXX_Size() int {
	return xxx_messageInfo_QueryInstalledChaincodesResult_InstalledChaincode.Size(m)
}
func (m *QueryInstalledChaincodesResult_InstalledChaincode) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryInstalledChaincodesResult_InstalledChaincode.DiscardUnknown(m)
}

var xxx_messageInfo_QueryInstalledChaincodesResult_InstalledChaincode proto.InternalMessageInfo

func (m *QueryInstalledChaincodesResult_InstalledChaincode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *QueryInstalledChaincodesResult_InstalledChaincode) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *QueryInstalledChaincodesResult_InstalledChaincode) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// DefineChaincodeForMyOrgArgs is the message used as arguments to
// `+lifecycle.DefineChaincodeForMyOrg`.
type DefineChaincodeForMyOrgArgs struct {
	Sequence             int64    `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Hash                 []byte   `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	EndorsementPlugin    string   `protobuf:"bytes,5,opt,name=endorsement_plugin,json=endorsementPlugin,proto3" json:"endorsement_plugin,omitempty"`
	ValidationPlugin     string   `protobuf:"bytes,6,opt,name=validation_plugin,json=validationPlugin,proto3" json:"validation_plugin,omitempty"`
	ValidationParameter  []byte   `protobuf:"bytes,7,opt,name=validation_parameter,json=validationParameter,proto3" json:"validation_parameter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DefineChaincodeForMyOrgArgs) Reset()         { *m = DefineChaincodeForMyOrgArgs{} }
func (m *DefineChaincodeForMyOrgArgs) String() string { return proto.CompactTextString(m) }
func (*DefineChaincodeForMyOrgArgs) ProtoMessage()    {}
func (*DefineChaincodeForMyOrgArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_lifecycle_7e2bec26dbec8f7e, []int{6}
}
func (m *DefineChaincodeForMyOrgArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefineChaincodeForMyOrgArgs.Unmarshal(m, b)
}
func (m *DefineChaincodeForMyOrgArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefineChaincodeForMyOrgArgs.Marshal(b, m, deterministic)
}
func (dst *DefineChaincodeForMyOrgArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefineChaincodeForMyOrgArgs.Merge(dst, src)
}
func (m *DefineChaincodeForMyOrgArgs) XXX_Size() int {
	return xxx_messageInfo_DefineChaincodeForMyOrgArgs.Size(m)
}
func (m *DefineChaincodeForMyOrgArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_DefineChaincodeForMyOrgArgs.DiscardUnknown(m)
}

var xxx_messageInfo_DefineChaincodeForMyOrgArgs proto.InternalMessageInfo

func (m *DefineChaincodeForMyOrgArgs) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *DefineChaincodeForMyOrgArgs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DefineChaincodeForMyOrgArgs) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DefineChaincodeForMyOrgArgs) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *DefineChaincodeForMyOrgArgs) GetEndorsementPlugin() string {
	if m != nil {
		return m.EndorsementPlugin
	}
	return ""
}

func (m *DefineChaincodeForMyOrgArgs) GetValidationPlugin() string {
	if m != nil {
		return m.ValidationPlugin
	}
	return ""
}

func (m *DefineChaincodeForMyOrgArgs) GetValidationParameter() []byte {
	if m != nil {
		return m.ValidationParameter
	}
	return nil
}

// DefineChaincodeForMyOrgResult is the message returned by
// `+lifecycle.DefineChaincodeForMyOrg`. Currently it returns
// nothing, but may be extended in the future.
type DefineChaincodeForMyOrgResult struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DefineChaincodeForMyOrgResult) Reset()         { *m = DefineChaincodeForMyOrgResult{} }
func (m *DefineChaincodeForMyOrgResult) String() string { return proto.CompactTextString(m) }
func (*DefineChaincodeForMyOrgResult) ProtoMessage()    {}
func (*DefineChaincodeForMyOrgResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_lifecycle_7e2bec26dbec8f7e, []int{7}
}
func (m *DefineChaincodeForMyOrgResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefineChaincodeForMyOrgResult.Unmarshal(m, b)
}
func (m *DefineChaincodeForMyOrgResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefineChaincodeForMyOrgResult.Marshal(b, m, deterministic)
}
func (dst *DefineChaincodeForMyOrgResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefineChaincodeForMyOrgResult.Merge(dst, src)
}
func (m *DefineChaincodeForMyOrgResult) XXX_Size() int {
	return xxx_messageInfo_DefineChaincodeForMyOrgResult.Size(m)
}
func (m *DefineChaincodeForMyOrgResult) XXX_DiscardUnknown() {
	xxx_messageInfo_DefineChaincodeForMyOrgResult.DiscardUnknown(m)
}

var xxx_messageInfo_DefineChaincodeForMyOrgResult proto.InternalMessageInfo

// DefineChaincodeArgs is the message used as arguments to
// `+lifecycle.DefineChaincode`.
type DefineChaincodeArgs struct {
	Sequence             int64    `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Hash                 []byte   `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	EndorsementPlugin    string   `protobuf:"bytes,5,opt,name=endorsement_plugin,json=endorsementPlugin,proto3" json:"endorsement_plugin,omitempty"`
	ValidationPlugin     string   `protobuf:"bytes,6,opt,name=validation_plugin,json=validationPlugin,proto3" json:"validation_plugin,omitempty"`
	ValidationParameter  []byte   `protobuf:"bytes,7,opt,name=validation_parameter,json=validationParameter,proto3" json:"validation_parameter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DefineChaincodeArgs) Reset()         { *m = DefineChaincodeArgs{} }
func (m *DefineChaincodeArgs) String() string { return proto.CompactTextString(m) }
func (*DefineChaincodeArgs) ProtoMessage()    {}
func (*DefineChaincodeArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_lifecycle_7e2bec26dbec8f7e, []int{8}
}
func (m *DefineChaincodeArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefineChaincodeArgs.Unmarshal(m, b)
}
func (m *DefineChaincodeArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefineChaincodeArgs.Marshal(b, m, deterministic)
}
func (dst *DefineChaincodeArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefineChaincodeArgs.Merge(dst, src)
}
func (m *DefineChaincodeArgs) XXX_Size() int {
	return xxx_messageInfo_DefineChaincodeArgs.Size(m)
}
func (m *DefineChaincodeArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_DefineChaincodeArgs.DiscardUnknown(m)
}

var xxx_messageInfo_DefineChaincodeArgs proto.InternalMessageInfo

func (m *DefineChaincodeArgs) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *DefineChaincodeArgs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DefineChaincodeArgs) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *DefineChaincodeArgs) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *DefineChaincodeArgs) GetEndorsementPlugin() string {
	if m != nil {
		return m.EndorsementPlugin
	}
	return ""
}

func (m *DefineChaincodeArgs) GetValidationPlugin() string {
	if m != nil {
		return m.ValidationPlugin
	}
	return ""
}

func (m *DefineChaincodeArgs) GetValidationParameter() []byte {
	if m != nil {
		return m.ValidationParameter
	}
	return nil
}

// DefineChaincodeResult is the message returned by
// `+lifecycle.DefineChaincode`. Currently it returns
// nothing, but may be extended in the future.
type DefineChaincodeResult struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DefineChaincodeResult) Reset()         { *m = DefineChaincodeResult{} }
func (m *DefineChaincodeResult) String() string { return proto.CompactTextString(m) }
func (*DefineChaincodeResult) ProtoMessage()    {}
func (*DefineChaincodeResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_lifecycle_7e2bec26dbec8f7e, []int{9}
}
func (m *DefineChaincodeResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DefineChaincodeResult.Unmarshal(m, b)
}
func (m *DefineChaincodeResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DefineChaincodeResult.Marshal(b, m, deterministic)
}
func (dst *DefineChaincodeResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DefineChaincodeResult.Merge(dst, src)
}
func (m *DefineChaincodeResult) XXX_Size() int {
	return xxx_messageInfo_DefineChaincodeResult.Size(m)
}
func (m *DefineChaincodeResult) XXX_DiscardUnknown() {
	xxx_messageInfo_DefineChaincodeResult.DiscardUnknown(m)
}

var xxx_messageInfo_DefineChaincodeResult proto.InternalMessageInfo

// QueryDefinedChaincode is the message used as arguments to
// `+lifecycle.QueryDefinedChaincode`.
type QueryDefinedChaincodeArgs struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryDefinedChaincodeArgs) Reset()         { *m = QueryDefinedChaincodeArgs{} }
func (m *QueryDefinedChaincodeArgs) String() string { return proto.CompactTextString(m) }
func (*QueryDefinedChaincodeArgs) ProtoMessage()    {}
func (*QueryDefinedChaincodeArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_lifecycle_7e2bec26dbec8f7e, []int{10}
}
func (m *QueryDefinedChaincodeArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryDefinedChaincodeArgs.Unmarshal(m, b)
}
func (m *QueryDefinedChaincodeArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryDefinedChaincodeArgs.Marshal(b, m, deterministic)
}
func (dst *QueryDefinedChaincodeArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDefinedChaincodeArgs.Merge(dst, src)
}
func (m *QueryDefinedChaincodeArgs) XXX_Size() int {
	return xxx_messageInfo_QueryDefinedChaincodeArgs.Size(m)
}
func (m *QueryDefinedChaincodeArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDefinedChaincodeArgs.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDefinedChaincodeArgs proto.InternalMessageInfo

func (m *QueryDefinedChaincodeArgs) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// DefineChaincodeResult is the message returned by
// `+lifecycle.QueryDefinedChaincode`.
type QueryDefinedChaincodeResult struct {
	Sequence             int64    `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Hash                 []byte   `protobuf:"bytes,4,opt,name=hash,proto3" json:"hash,omitempty"`
	EndorsementPlugin    string   `protobuf:"bytes,5,opt,name=endorsement_plugin,json=endorsementPlugin,proto3" json:"endorsement_plugin,omitempty"`
	ValidationPlugin     string   `protobuf:"bytes,6,opt,name=validation_plugin,json=validationPlugin,proto3" json:"validation_plugin,omitempty"`
	ValidationParameter  []byte   `protobuf:"bytes,7,opt,name=validation_parameter,json=validationParameter,proto3" json:"validation_parameter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryDefinedChaincodeResult) Reset()         { *m = QueryDefinedChaincodeResult{} }
func (m *QueryDefinedChaincodeResult) String() string { return proto.CompactTextString(m) }
func (*QueryDefinedChaincodeResult) ProtoMessage()    {}
func (*QueryDefinedChaincodeResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_lifecycle_7e2bec26dbec8f7e, []int{11}
}
func (m *QueryDefinedChaincodeResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryDefinedChaincodeResult.Unmarshal(m, b)
}
func (m *QueryDefinedChaincodeResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryDefinedChaincodeResult.Marshal(b, m, deterministic)
}
func (dst *QueryDefinedChaincodeResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDefinedChaincodeResult.Merge(dst, src)
}
func (m *QueryDefinedChaincodeResult) XXX_Size() int {
	return xxx_messageInfo_QueryDefinedChaincodeResult.Size(m)
}
func (m *QueryDefinedChaincodeResult) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDefinedChaincodeResult.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDefinedChaincodeResult proto.InternalMessageInfo

func (m *QueryDefinedChaincodeResult) GetSequence() int64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *QueryDefinedChaincodeResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *QueryDefinedChaincodeResult) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *QueryDefinedChaincodeResult) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *QueryDefinedChaincodeResult) GetEndorsementPlugin() string {
	if m != nil {
		return m.EndorsementPlugin
	}
	return ""
}

func (m *QueryDefinedChaincodeResult) GetValidationPlugin() string {
	if m != nil {
		return m.ValidationPlugin
	}
	return ""
}

func (m *QueryDefinedChaincodeResult) GetValidationParameter() []byte {
	if m != nil {
		return m.ValidationParameter
	}
	return nil
}

// QueryDefinedNamespaces is the message used as arguments to
// `+lifecycle.QueryDefinedNamespaces`.
type QueryDefinedNamespacesArgs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryDefinedNamespacesArgs) Reset()         { *m = QueryDefinedNamespacesArgs{} }
func (m *QueryDefinedNamespacesArgs) String() string { return proto.CompactTextString(m) }
func (*QueryDefinedNamespacesArgs) ProtoMessage()    {}
func (*QueryDefinedNamespacesArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_lifecycle_7e2bec26dbec8f7e, []int{12}
}
func (m *QueryDefinedNamespacesArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryDefinedNamespacesArgs.Unmarshal(m, b)
}
func (m *QueryDefinedNamespacesArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryDefinedNamespacesArgs.Marshal(b, m, deterministic)
}
func (dst *QueryDefinedNamespacesArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDefinedNamespacesArgs.Merge(dst, src)
}
func (m *QueryDefinedNamespacesArgs) XXX_Size() int {
	return xxx_messageInfo_QueryDefinedNamespacesArgs.Size(m)
}
func (m *QueryDefinedNamespacesArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDefinedNamespacesArgs.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDefinedNamespacesArgs proto.InternalMessageInfo

// QueryDefinedNamespaces is the message returned by
// `+lifecycle.QueryDefinedNamespaces`.
type QueryDefinedNamespacesResult struct {
	Namespaces           map[string]*QueryDefinedNamespacesResult_Namespace `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                                           `json:"-"`
	XXX_unrecognized     []byte                                             `json:"-"`
	XXX_sizecache        int32                                              `json:"-"`
}

func (m *QueryDefinedNamespacesResult) Reset()         { *m = QueryDefinedNamespacesResult{} }
func (m *QueryDefinedNamespacesResult) String() string { return proto.CompactTextString(m) }
func (*QueryDefinedNamespacesResult) ProtoMessage()    {}
func (*QueryDefinedNamespacesResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_lifecycle_7e2bec26dbec8f7e, []int{13}
}
func (m *QueryDefinedNamespacesResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryDefinedNamespacesResult.Unmarshal(m, b)
}
func (m *QueryDefinedNamespacesResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryDefinedNamespacesResult.Marshal(b, m, deterministic)
}
func (dst *QueryDefinedNamespacesResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDefinedNamespacesResult.Merge(dst, src)
}
func (m *QueryDefinedNamespacesResult) XXX_Size() int {
	return xxx_messageInfo_QueryDefinedNamespacesResult.Size(m)
}
func (m *QueryDefinedNamespacesResult) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDefinedNamespacesResult.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDefinedNamespacesResult proto.InternalMessageInfo

func (m *QueryDefinedNamespacesResult) GetNamespaces() map[string]*QueryDefinedNamespacesResult_Namespace {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

type QueryDefinedNamespacesResult_Namespace struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryDefinedNamespacesResult_Namespace) Reset() {
	*m = QueryDefinedNamespacesResult_Namespace{}
}
func (m *QueryDefinedNamespacesResult_Namespace) String() string { return proto.CompactTextString(m) }
func (*QueryDefinedNamespacesResult_Namespace) ProtoMessage()    {}
func (*QueryDefinedNamespacesResult_Namespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_lifecycle_7e2bec26dbec8f7e, []int{13, 0}
}
func (m *QueryDefinedNamespacesResult_Namespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryDefinedNamespacesResult_Namespace.Unmarshal(m, b)
}
func (m *QueryDefinedNamespacesResult_Namespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryDefinedNamespacesResult_Namespace.Marshal(b, m, deterministic)
}
func (dst *QueryDefinedNamespacesResult_Namespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDefinedNamespacesResult_Namespace.Merge(dst, src)
}
func (m *QueryDefinedNamespacesResult_Namespace) XXX_Size() int {
	return xxx_messageInfo_QueryDefinedNamespacesResult_Namespace.Size(m)
}
func (m *QueryDefinedNamespacesResult_Namespace) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDefinedNamespacesResult_Namespace.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDefinedNamespacesResult_Namespace proto.InternalMessageInfo

func (m *QueryDefinedNamespacesResult_Namespace) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*InstallChaincodeArgs)(nil), "lifecycle.InstallChaincodeArgs")
	proto.RegisterType((*InstallChaincodeResult)(nil), "lifecycle.InstallChaincodeResult")
	proto.RegisterType((*QueryInstalledChaincodeArgs)(nil), "lifecycle.QueryInstalledChaincodeArgs")
	proto.RegisterType((*QueryInstalledChaincodeResult)(nil), "lifecycle.QueryInstalledChaincodeResult")
	proto.RegisterType((*QueryInstalledChaincodesArgs)(nil), "lifecycle.QueryInstalledChaincodesArgs")
	proto.RegisterType((*QueryInstalledChaincodesResult)(nil), "lifecycle.QueryInstalledChaincodesResult")
	proto.RegisterType((*QueryInstalledChaincodesResult_InstalledChaincode)(nil), "lifecycle.QueryInstalledChaincodesResult.InstalledChaincode")
	proto.RegisterType((*DefineChaincodeForMyOrgArgs)(nil), "lifecycle.DefineChaincodeForMyOrgArgs")
	proto.RegisterType((*DefineChaincodeForMyOrgResult)(nil), "lifecycle.DefineChaincodeForMyOrgResult")
	proto.RegisterType((*DefineChaincodeArgs)(nil), "lifecycle.DefineChaincodeArgs")
	proto.RegisterType((*DefineChaincodeResult)(nil), "lifecycle.DefineChaincodeResult")
	proto.RegisterType((*QueryDefinedChaincodeArgs)(nil), "lifecycle.QueryDefinedChaincodeArgs")
	proto.RegisterType((*QueryDefinedChaincodeResult)(nil), "lifecycle.QueryDefinedChaincodeResult")
	proto.RegisterType((*QueryDefinedNamespacesArgs)(nil), "lifecycle.QueryDefinedNamespacesArgs")
	proto.RegisterType((*QueryDefinedNamespacesResult)(nil), "lifecycle.QueryDefinedNamespacesResult")
	proto.RegisterMapType((map[string]*QueryDefinedNamespacesResult_Namespace)(nil), "lifecycle.QueryDefinedNamespacesResult.NamespacesEntry")
	proto.RegisterType((*QueryDefinedNamespacesResult_Namespace)(nil), "lifecycle.QueryDefinedNamespacesResult.Namespace")
}

func init() {
	proto.RegisterFile("peer/lifecycle/lifecycle.proto", fileDescriptor_lifecycle_7e2bec26dbec8f7e)
}

var fileDescriptor_lifecycle_7e2bec26dbec8f7e = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x55, 0xcb, 0x6f, 0xd3, 0x4e,
	0x10, 0x96, 0xed, 0x3e, 0x7e, 0x9d, 0x56, 0xfa, 0xb5, 0x6e, 0xa0, 0x6e, 0xda, 0xa6, 0x91, 0x4f,
	0x91, 0x28, 0xb6, 0xda, 0x1c, 0x40, 0x15, 0x17, 0xde, 0x42, 0x08, 0x28, 0x3e, 0x80, 0xc4, 0x25,
	0xda, 0x38, 0x13, 0x67, 0x55, 0x67, 0x6d, 0x76, 0xed, 0x48, 0xbe, 0x21, 0x8e, 0xfd, 0x4f, 0xf9,
	0x2f, 0x90, 0xd7, 0xeb, 0x07, 0x79, 0x14, 0x55, 0x1c, 0x7b, 0x1b, 0xcf, 0xf7, 0x7d, 0x33, 0xa3,
	0x6f, 0xbc, 0xbb, 0xd0, 0x89, 0x11, 0xb9, 0x1b, 0xd2, 0x31, 0xfa, 0x99, 0x1f, 0x62, 0x1d, 0x39,
	0x31, 0x8f, 0x92, 0xc8, 0xdc, 0xaa, 0x12, 0xf6, 0x0f, 0x0d, 0x5a, 0xef, 0x98, 0x48, 0x48, 0x18,
	0xbe, 0x9c, 0x10, 0xca, 0xfc, 0x68, 0x84, 0xcf, 0x79, 0x20, 0x4c, 0x13, 0xd6, 0x18, 0x99, 0xa2,
	0xa5, 0x75, 0xb5, 0xde, 0x96, 0x27, 0x63, 0xd3, 0x82, 0xcd, 0x19, 0x72, 0x41, 0x23, 0x66, 0xe9,
	0x32, 0x5d, 0x7e, 0x9a, 0x97, 0x70, 0xe8, 0x97, 0xf2, 0x01, 0x2d, 0xea, 0x0d, 0x62, 0xe2, 0x5f,
	0x93, 0x00, 0x2d, 0xa3, 0xab, 0xf5, 0x76, 0xbc, 0x83, 0x8a, 0xa0, 0xfa, 0x5d, 0x15, 0xb0, 0x7d,
	0x06, 0x0f, 0xe7, 0x27, 0xf0, 0x50, 0xa4, 0x61, 0x92, 0xcf, 0x30, 0x21, 0x62, 0x22, 0x67, 0xd8,
	0xf1, 0x64, 0x6c, 0xbf, 0x87, 0xa3, 0xcf, 0x29, 0xf2, 0x4c, 0x49, 0x70, 0xf4, 0x0f, 0x63, 0xdb,
	0x7d, 0x38, 0x59, 0x51, 0xec, 0x96, 0x09, 0x3a, 0x70, 0xbc, 0x42, 0x24, 0xf2, 0x11, 0xec, 0x5f,
	0x1a, 0x74, 0x56, 0x11, 0x54, 0xd9, 0x08, 0x5a, 0xb4, 0x04, 0x07, 0x95, 0x2f, 0xc2, 0xd2, 0xba,
	0x46, 0x6f, 0xfb, 0xe2, 0x99, 0x53, 0x2f, 0xec, 0xf6, 0x42, 0xce, 0x92, 0xc1, 0xf7, 0xe9, 0x22,
	0xbb, 0xfd, 0x05, 0xcc, 0x45, 0xea, 0x1d, 0x77, 0x5c, 0x7a, 0x61, 0x34, 0xbc, 0xb8, 0xd1, 0xe1,
	0xe8, 0x15, 0x8e, 0x29, 0xc3, 0xaa, 0xea, 0x9b, 0x88, 0x7f, 0xc8, 0x3e, 0xf1, 0x40, 0xae, 0xa3,
	0x0d, 0xff, 0x09, 0xfc, 0x9e, 0x22, 0xf3, 0x8b, 0x2e, 0x86, 0x57, 0x7d, 0x57, 0xdd, 0xf5, 0xe5,
	0xdd, 0x8d, 0xe5, 0xdd, 0xd7, 0xea, 0xee, 0xe6, 0x63, 0x30, 0x91, 0x8d, 0x22, 0x2e, 0x70, 0x8a,
	0x2c, 0x19, 0xc4, 0x61, 0x1a, 0x50, 0x66, 0xad, 0x4b, 0xe1, 0x5e, 0x03, 0xb9, 0x92, 0x80, 0xf9,
	0x08, 0xf6, 0x66, 0x24, 0xa4, 0x23, 0x92, 0xd0, 0x88, 0x95, 0xec, 0x0d, 0xc9, 0xde, 0xad, 0x01,
	0x45, 0x3e, 0x87, 0x56, 0x93, 0x4c, 0x38, 0x99, 0x62, 0x82, 0xdc, 0xda, 0x94, 0xfd, 0xf7, 0x1b,
	0xfc, 0x12, 0xb2, 0x4f, 0xe1, 0x64, 0x85, 0x17, 0xc5, 0xb6, 0xec, 0x9f, 0x3a, 0xec, 0xcf, 0x31,
	0xee, 0x9f, 0x4b, 0x07, 0xf0, 0x60, 0xce, 0x03, 0xe5, 0x8e, 0x0b, 0x87, 0xf2, 0x6f, 0x2f, 0xd0,
	0xbf, 0x9f, 0x6b, 0xf9, 0xf3, 0x2d, 0x55, 0xa8, 0x53, 0x76, 0xaf, 0x6c, 0x3d, 0x86, 0x76, 0xd3,
	0x8b, 0x8f, 0x64, 0x8a, 0x22, 0x26, 0xbe, 0xba, 0x93, 0x6e, 0x74, 0x75, 0x69, 0x2d, 0xc0, 0xca,
	0xab, 0xaf, 0x00, 0xac, 0xca, 0xa9, 0x7b, 0xe8, 0xc9, 0xfc, 0x3d, 0xb4, 0x42, 0xec, 0xd4, 0x89,
	0xd7, 0x2c, 0xe1, 0x99, 0xd7, 0x28, 0xd5, 0x3e, 0x85, 0xad, 0x0a, 0xce, 0x7d, 0x4c, 0xb2, 0xb8,
	0xda, 0x62, 0x1e, 0xb7, 0x63, 0xf8, 0x7f, 0x4e, 0x6f, 0xee, 0x82, 0x71, 0x8d, 0x99, 0x62, 0xe5,
	0xa1, 0xf9, 0x16, 0xd6, 0x67, 0x24, 0x4c, 0x8b, 0x7d, 0x6d, 0x5f, 0x9c, 0xdf, 0x79, 0x32, 0xaf,
	0xd0, 0x5f, 0xea, 0x4f, 0xb5, 0x17, 0x3e, 0x9c, 0x45, 0x3c, 0x70, 0x26, 0x59, 0x8c, 0x3c, 0xc4,
	0x51, 0x80, 0xdc, 0x19, 0x93, 0x21, 0xa7, 0x7e, 0xf1, 0x3c, 0x0a, 0x27, 0x7f, 0x3e, 0xeb, 0x0e,
	0xdf, 0xfa, 0x01, 0x4d, 0x26, 0xe9, 0xd0, 0xf1, 0xa3, 0xa9, 0xdb, 0x10, 0xb9, 0x85, 0xc8, 0x2d,
	0x44, 0xee, 0x9f, 0x6f, 0xee, 0x70, 0x43, 0xa6, 0xfb, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x1a,
	0x19, 0x56, 0xca, 0x8c, 0x07, 0x00, 0x00,
}
