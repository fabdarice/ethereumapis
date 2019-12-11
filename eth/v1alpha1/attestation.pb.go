// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eth/v1alpha1/attestation.proto

package eth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Attestation struct {
	AggregationBits      []byte           `protobuf:"bytes,1,opt,name=aggregation_bits,json=aggregationBits,proto3" json:"aggregation_bits,omitempty"`
	Data                 *AttestationData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Signature            []byte           `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Attestation) Reset()         { *m = Attestation{} }
func (m *Attestation) String() string { return proto.CompactTextString(m) }
func (*Attestation) ProtoMessage()    {}
func (*Attestation) Descriptor() ([]byte, []int) {
	return fileDescriptor_7894c78397fc93a1, []int{0}
}

func (m *Attestation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attestation.Unmarshal(m, b)
}
func (m *Attestation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attestation.Marshal(b, m, deterministic)
}
func (m *Attestation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attestation.Merge(m, src)
}
func (m *Attestation) XXX_Size() int {
	return xxx_messageInfo_Attestation.Size(m)
}
func (m *Attestation) XXX_DiscardUnknown() {
	xxx_messageInfo_Attestation.DiscardUnknown(m)
}

var xxx_messageInfo_Attestation proto.InternalMessageInfo

func (m *Attestation) GetAggregationBits() []byte {
	if m != nil {
		return m.AggregationBits
	}
	return nil
}

func (m *Attestation) GetData() *AttestationData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Attestation) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type AggregateAttestationAndProof struct {
	Index                uint64       `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	SelectionProof       []byte       `protobuf:"bytes,2,opt,name=selection_proof,json=selectionProof,proto3" json:"selection_proof,omitempty"`
	Aggregate            *Attestation `protobuf:"bytes,3,opt,name=aggregate,proto3" json:"aggregate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AggregateAttestationAndProof) Reset()         { *m = AggregateAttestationAndProof{} }
func (m *AggregateAttestationAndProof) String() string { return proto.CompactTextString(m) }
func (*AggregateAttestationAndProof) ProtoMessage()    {}
func (*AggregateAttestationAndProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_7894c78397fc93a1, []int{1}
}

func (m *AggregateAttestationAndProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregateAttestationAndProof.Unmarshal(m, b)
}
func (m *AggregateAttestationAndProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregateAttestationAndProof.Marshal(b, m, deterministic)
}
func (m *AggregateAttestationAndProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregateAttestationAndProof.Merge(m, src)
}
func (m *AggregateAttestationAndProof) XXX_Size() int {
	return xxx_messageInfo_AggregateAttestationAndProof.Size(m)
}
func (m *AggregateAttestationAndProof) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregateAttestationAndProof.DiscardUnknown(m)
}

var xxx_messageInfo_AggregateAttestationAndProof proto.InternalMessageInfo

func (m *AggregateAttestationAndProof) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *AggregateAttestationAndProof) GetSelectionProof() []byte {
	if m != nil {
		return m.SelectionProof
	}
	return nil
}

func (m *AggregateAttestationAndProof) GetAggregate() *Attestation {
	if m != nil {
		return m.Aggregate
	}
	return nil
}

type AttestationData struct {
	Slot                 uint64      `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty"`
	CommitteeIndex       uint64      `protobuf:"varint,2,opt,name=committee_index,json=committeeIndex,proto3" json:"committee_index,omitempty"`
	BeaconBlockRoot      []byte      `protobuf:"bytes,3,opt,name=beacon_block_root,json=beaconBlockRoot,proto3" json:"beacon_block_root,omitempty"`
	Source               *Checkpoint `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	Target               *Checkpoint `protobuf:"bytes,5,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AttestationData) Reset()         { *m = AttestationData{} }
func (m *AttestationData) String() string { return proto.CompactTextString(m) }
func (*AttestationData) ProtoMessage()    {}
func (*AttestationData) Descriptor() ([]byte, []int) {
	return fileDescriptor_7894c78397fc93a1, []int{2}
}

func (m *AttestationData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestationData.Unmarshal(m, b)
}
func (m *AttestationData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestationData.Marshal(b, m, deterministic)
}
func (m *AttestationData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestationData.Merge(m, src)
}
func (m *AttestationData) XXX_Size() int {
	return xxx_messageInfo_AttestationData.Size(m)
}
func (m *AttestationData) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestationData.DiscardUnknown(m)
}

var xxx_messageInfo_AttestationData proto.InternalMessageInfo

func (m *AttestationData) GetSlot() uint64 {
	if m != nil {
		return m.Slot
	}
	return 0
}

func (m *AttestationData) GetCommitteeIndex() uint64 {
	if m != nil {
		return m.CommitteeIndex
	}
	return 0
}

func (m *AttestationData) GetBeaconBlockRoot() []byte {
	if m != nil {
		return m.BeaconBlockRoot
	}
	return nil
}

func (m *AttestationData) GetSource() *Checkpoint {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *AttestationData) GetTarget() *Checkpoint {
	if m != nil {
		return m.Target
	}
	return nil
}

type Crosslink struct {
	Shard                uint64   `protobuf:"varint,1,opt,name=shard,proto3" json:"shard,omitempty"`
	ParentRoot           []byte   `protobuf:"bytes,2,opt,name=parent_root,json=parentRoot,proto3" json:"parent_root,omitempty"`
	StartEpoch           uint64   `protobuf:"varint,3,opt,name=start_epoch,json=startEpoch,proto3" json:"start_epoch,omitempty"`
	EndEpoch             uint64   `protobuf:"varint,4,opt,name=end_epoch,json=endEpoch,proto3" json:"end_epoch,omitempty"`
	DataRoot             []byte   `protobuf:"bytes,5,opt,name=data_root,json=dataRoot,proto3" json:"data_root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Crosslink) Reset()         { *m = Crosslink{} }
func (m *Crosslink) String() string { return proto.CompactTextString(m) }
func (*Crosslink) ProtoMessage()    {}
func (*Crosslink) Descriptor() ([]byte, []int) {
	return fileDescriptor_7894c78397fc93a1, []int{3}
}

func (m *Crosslink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Crosslink.Unmarshal(m, b)
}
func (m *Crosslink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Crosslink.Marshal(b, m, deterministic)
}
func (m *Crosslink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Crosslink.Merge(m, src)
}
func (m *Crosslink) XXX_Size() int {
	return xxx_messageInfo_Crosslink.Size(m)
}
func (m *Crosslink) XXX_DiscardUnknown() {
	xxx_messageInfo_Crosslink.DiscardUnknown(m)
}

var xxx_messageInfo_Crosslink proto.InternalMessageInfo

func (m *Crosslink) GetShard() uint64 {
	if m != nil {
		return m.Shard
	}
	return 0
}

func (m *Crosslink) GetParentRoot() []byte {
	if m != nil {
		return m.ParentRoot
	}
	return nil
}

func (m *Crosslink) GetStartEpoch() uint64 {
	if m != nil {
		return m.StartEpoch
	}
	return 0
}

func (m *Crosslink) GetEndEpoch() uint64 {
	if m != nil {
		return m.EndEpoch
	}
	return 0
}

func (m *Crosslink) GetDataRoot() []byte {
	if m != nil {
		return m.DataRoot
	}
	return nil
}

type Checkpoint struct {
	Epoch                uint64   `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	Root                 []byte   `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Checkpoint) Reset()         { *m = Checkpoint{} }
func (m *Checkpoint) String() string { return proto.CompactTextString(m) }
func (*Checkpoint) ProtoMessage()    {}
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_7894c78397fc93a1, []int{4}
}

func (m *Checkpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Checkpoint.Unmarshal(m, b)
}
func (m *Checkpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Checkpoint.Marshal(b, m, deterministic)
}
func (m *Checkpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Checkpoint.Merge(m, src)
}
func (m *Checkpoint) XXX_Size() int {
	return xxx_messageInfo_Checkpoint.Size(m)
}
func (m *Checkpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Checkpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Checkpoint proto.InternalMessageInfo

func (m *Checkpoint) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func (m *Checkpoint) GetRoot() []byte {
	if m != nil {
		return m.Root
	}
	return nil
}

func init() {
	proto.RegisterType((*Attestation)(nil), "ethereum.eth.v1alpha1.Attestation")
	proto.RegisterType((*AggregateAttestationAndProof)(nil), "ethereum.eth.v1alpha1.AggregateAttestationAndProof")
	proto.RegisterType((*AttestationData)(nil), "ethereum.eth.v1alpha1.AttestationData")
	proto.RegisterType((*Crosslink)(nil), "ethereum.eth.v1alpha1.Crosslink")
	proto.RegisterType((*Checkpoint)(nil), "ethereum.eth.v1alpha1.Checkpoint")
}

func init() { proto.RegisterFile("eth/v1alpha1/attestation.proto", fileDescriptor_7894c78397fc93a1) }

var fileDescriptor_7894c78397fc93a1 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xe5, 0xe0, 0x54, 0xcd, 0xa4, 0x6a, 0xca, 0x8a, 0x4a, 0x41, 0x54, 0xb4, 0xf8, 0x00,
	0x85, 0x83, 0xa3, 0x16, 0xa9, 0x12, 0x70, 0x21, 0x29, 0x39, 0x70, 0x8b, 0x7c, 0x44, 0x95, 0xa2,
	0x8d, 0x3d, 0x78, 0x57, 0x71, 0xbc, 0xd6, 0xee, 0x04, 0xc1, 0x6b, 0x70, 0xe1, 0x84, 0xb8, 0xf3,
	0x28, 0x3c, 0x0d, 0x8f, 0x80, 0x76, 0xd7, 0x8e, 0x2d, 0x54, 0x44, 0x6f, 0x9e, 0x7f, 0x76, 0xfe,
	0xf9, 0x76, 0xc6, 0x0b, 0x8f, 0x91, 0xc4, 0xe4, 0xd3, 0x05, 0x2f, 0x2a, 0xc1, 0x2f, 0x26, 0x9c,
	0x08, 0x0d, 0x71, 0x92, 0xaa, 0x8c, 0x2b, 0xad, 0x48, 0xb1, 0x63, 0x24, 0x81, 0x1a, 0xb7, 0x9b,
	0x18, 0x49, 0xc4, 0xcd, 0xc1, 0xe8, 0x6b, 0x00, 0xc3, 0x69, 0x7b, 0x98, 0x3d, 0x87, 0x23, 0x9e,
	0xe7, 0x1a, 0x73, 0x17, 0x2e, 0x57, 0x92, 0xcc, 0x38, 0x38, 0x0b, 0xce, 0x0f, 0x92, 0x51, 0x47,
	0x9f, 0x49, 0x32, 0xec, 0x35, 0x84, 0x19, 0x27, 0x3e, 0xee, 0x9d, 0x05, 0xe7, 0xc3, 0xcb, 0xa7,
	0xf1, 0xad, 0x0d, 0xe2, 0x8e, 0xf9, 0x3b, 0x4e, 0x3c, 0x71, 0x35, 0xec, 0x04, 0x06, 0x46, 0xe6,
	0x25, 0xa7, 0xad, 0xc6, 0xf1, 0x3d, 0xe7, 0xdf, 0x0a, 0xd1, 0x8f, 0x00, 0x4e, 0xa6, 0x75, 0x37,
	0xec, 0x18, 0x4c, 0xcb, 0x6c, 0xa1, 0x95, 0xfa, 0xc8, 0x1e, 0x40, 0x5f, 0x96, 0x19, 0x7e, 0x76,
	0x68, 0x61, 0xe2, 0x03, 0xf6, 0x0c, 0x46, 0x06, 0x0b, 0x4c, 0x1d, 0x79, 0x65, 0x0f, 0x3a, 0xb6,
	0x83, 0xe4, 0x70, 0x27, 0xfb, 0xf2, 0xb7, 0x30, 0x68, 0x2e, 0xe3, 0xbb, 0x0f, 0x2f, 0xa3, 0xff,
	0xe3, 0x27, 0x6d, 0x51, 0xf4, 0x3b, 0x80, 0xd1, 0x5f, 0x37, 0x63, 0x0c, 0x42, 0x53, 0x28, 0xaa,
	0x99, 0xdc, 0xb7, 0x45, 0x4a, 0xd5, 0x66, 0x23, 0x89, 0x10, 0x97, 0x1e, 0xb9, 0xe7, 0xd2, 0x87,
	0x3b, 0xf9, 0xbd, 0x63, 0x7f, 0x01, 0xf7, 0x57, 0xc8, 0x53, 0x3b, 0xf2, 0x42, 0xa5, 0xeb, 0xa5,
	0x56, 0x8a, 0xea, 0xc1, 0x8c, 0x7c, 0x62, 0x66, 0xf5, 0x44, 0x29, 0x62, 0xaf, 0x60, 0xcf, 0xa8,
	0xad, 0x4e, 0x71, 0x1c, 0x3a, 0xf6, 0x27, 0xff, 0x60, 0xbf, 0x16, 0x98, 0xae, 0x2b, 0x25, 0x4b,
	0x4a, 0xea, 0x02, 0x5b, 0x4a, 0x5c, 0xe7, 0x48, 0xe3, 0xfe, 0x9d, 0x4b, 0x7d, 0x41, 0xf4, 0x3d,
	0x80, 0xc1, 0xb5, 0x56, 0xc6, 0x14, 0xb2, 0x5c, 0xdb, 0x0d, 0x18, 0xc1, 0x75, 0xd6, 0x6c, 0xc0,
	0x05, 0xec, 0x14, 0x86, 0x15, 0xd7, 0x58, 0x92, 0xe7, 0xf7, 0xd3, 0x07, 0x2f, 0x39, 0xf4, 0x53,
	0x18, 0x1a, 0xe2, 0x9a, 0x96, 0x58, 0xa9, 0x54, 0xb8, 0x0b, 0x86, 0x09, 0x38, 0x69, 0x6e, 0x15,
	0xf6, 0x08, 0x06, 0x58, 0x66, 0x75, 0x3a, 0x74, 0xe9, 0x7d, 0x2c, 0xb3, 0x5d, 0xd2, 0xfe, 0x3d,
	0xde, 0xbc, 0xef, 0xcc, 0xf7, 0xad, 0x60, 0xad, 0xa3, 0x2b, 0x80, 0x96, 0xda, 0xf2, 0x79, 0x8f,
	0x9a, 0xcf, 0x05, 0x76, 0x45, 0x1d, 0x30, 0xf7, 0x3d, 0xfb, 0x16, 0xc0, 0x43, 0xa5, 0xf3, 0xdb,
	0x07, 0x31, 0x3b, 0xea, 0x6c, 0x79, 0x61, 0x1f, 0xd2, 0x22, 0xf8, 0x70, 0x95, 0x4b, 0x12, 0xdb,
	0x55, 0x9c, 0xaa, 0xcd, 0xa4, 0xd2, 0x5f, 0xcc, 0x86, 0x93, 0x4c, 0x0b, 0xbe, 0x32, 0x93, 0xc6,
	0x83, 0x57, 0xd2, 0x05, 0xbb, 0x07, 0xf9, 0x06, 0x49, 0xfc, 0xec, 0x1d, 0xcf, 0x9b, 0x1e, 0xf3,
	0x4e, 0x8f, 0x5f, 0xad, 0x7e, 0x33, 0x27, 0x71, 0xd3, 0xe8, 0xab, 0x3d, 0xf7, 0x72, 0x5f, 0xfe,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x2c, 0x8f, 0x07, 0x85, 0xdb, 0x03, 0x00, 0x00,
}
