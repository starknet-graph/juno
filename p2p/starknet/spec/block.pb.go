// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: p2p/proto/block.proto

package spec

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// for now, we assume a small consensus, so this fits in 1M. Else, these will be repeated
type Signatures struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block      *BlockID              `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Signatures []*ConsensusSignature `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"` //
}

func (x *Signatures) Reset() {
	*x = Signatures{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signatures) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signatures) ProtoMessage() {}

func (x *Signatures) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signatures.ProtoReflect.Descriptor instead.
func (*Signatures) Descriptor() ([]byte, []int) {
	return file_p2p_proto_block_proto_rawDescGZIP(), []int{0}
}

func (x *Signatures) GetBlock() *BlockID {
	if x != nil {
		return x.Block
	}
	return nil
}

func (x *Signatures) GetSignatures() []*ConsensusSignature {
	if x != nil {
		return x.Signatures
	}
	return nil
}

// Note: commitments may change to be for the previous blocks like comet/tendermint
// hash of block header sent to L1
type BlockHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ParentHash       *Hash                  `protobuf:"bytes,1,opt,name=parent_hash,json=parentHash,proto3" json:"parent_hash,omitempty"`
	Number           uint64                 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Time             *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"` // TODO: see if this needs to be Felt252 or can be converted
	SequencerAddress *Address               `protobuf:"bytes,4,opt,name=sequencer_address,json=sequencerAddress,proto3" json:"sequencer_address,omitempty"`
	StateDiffs       *Merkle                `protobuf:"bytes,5,opt,name=state_diffs,json=stateDiffs,proto3" json:"state_diffs,omitempty"` //  By order of (contract, key), taking last in case of duplicates.
	// This means the proposer needs to sort after finishing the block (TBD: patricia? )
	// State is optional and appears every X blocks for the last block. This is to support
	// snapshot sync and also so that light nodes can sync on state without state diffs.
	State     *Patricia `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`                          // hash of contract and class patricia tries. Same as in L1. Later more trees will be included
	ProofFact *Hash     `protobuf:"bytes,7,opt,name=proof_fact,json=proofFact,proto3" json:"proof_fact,omitempty"` // for Kth block behind. A hash of the output of the proof
	// The following merkles can be built on the fly while sequencing/validating txs.
	Transactions    *Merkle  `protobuf:"bytes,8,opt,name=transactions,proto3" json:"transactions,omitempty"`                               // By order of execution. TBD: required? the client can execute (powerful machine) and match state diff
	Events          *Merkle  `protobuf:"bytes,9,opt,name=events,proto3" json:"events,omitempty"`                                           // By order of issuance. TBD: in receipts?
	Receipts        *Merkle  `protobuf:"bytes,10,opt,name=receipts,proto3" json:"receipts,omitempty"`                                      // By order of issuance.
	ProtocolVersion string   `protobuf:"bytes,11,opt,name=protocol_version,json=protocolVersion,proto3" json:"protocol_version,omitempty"` // Starknet version
	GasPrice        *Felt252 `protobuf:"bytes,12,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
}

func (x *BlockHeader) Reset() {
	*x = BlockHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHeader) ProtoMessage() {}

func (x *BlockHeader) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHeader.ProtoReflect.Descriptor instead.
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return file_p2p_proto_block_proto_rawDescGZIP(), []int{1}
}

func (x *BlockHeader) GetParentHash() *Hash {
	if x != nil {
		return x.ParentHash
	}
	return nil
}

func (x *BlockHeader) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *BlockHeader) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *BlockHeader) GetSequencerAddress() *Address {
	if x != nil {
		return x.SequencerAddress
	}
	return nil
}

func (x *BlockHeader) GetStateDiffs() *Merkle {
	if x != nil {
		return x.StateDiffs
	}
	return nil
}

func (x *BlockHeader) GetState() *Patricia {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *BlockHeader) GetProofFact() *Hash {
	if x != nil {
		return x.ProofFact
	}
	return nil
}

func (x *BlockHeader) GetTransactions() *Merkle {
	if x != nil {
		return x.Transactions
	}
	return nil
}

func (x *BlockHeader) GetEvents() *Merkle {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *BlockHeader) GetReceipts() *Merkle {
	if x != nil {
		return x.Receipts
	}
	return nil
}

func (x *BlockHeader) GetProtocolVersion() string {
	if x != nil {
		return x.ProtocolVersion
	}
	return ""
}

func (x *BlockHeader) GetGasPrice() *Felt252 {
	if x != nil {
		return x.GasPrice
	}
	return nil
}

type BlockProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proof []byte `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"` // proof size is currently 142K
}

func (x *BlockProof) Reset() {
	*x = BlockProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_block_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockProof) ProtoMessage() {}

func (x *BlockProof) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_block_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockProof.ProtoReflect.Descriptor instead.
func (*BlockProof) Descriptor() ([]byte, []int) {
	return file_p2p_proto_block_proto_rawDescGZIP(), []int{2}
}

func (x *BlockProof) GetProof() []byte {
	if x != nil {
		return x.Proof
	}
	return nil
}

// sent to all peers (except the ones this was received from, if any).
// for a fraction of peers, also send the GetBlockHeaders and GetBlockBodies response (as if they asked for it for this block)
type NewBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MaybeFull:
	//
	//	*NewBlock_Id
	//	*NewBlock_Header
	//	*NewBlock_Body
	MaybeFull isNewBlock_MaybeFull `protobuf_oneof:"maybe_full"`
}

func (x *NewBlock) Reset() {
	*x = NewBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_block_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewBlock) ProtoMessage() {}

func (x *NewBlock) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_block_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewBlock.ProtoReflect.Descriptor instead.
func (*NewBlock) Descriptor() ([]byte, []int) {
	return file_p2p_proto_block_proto_rawDescGZIP(), []int{3}
}

func (m *NewBlock) GetMaybeFull() isNewBlock_MaybeFull {
	if m != nil {
		return m.MaybeFull
	}
	return nil
}

func (x *NewBlock) GetId() *BlockID {
	if x, ok := x.GetMaybeFull().(*NewBlock_Id); ok {
		return x.Id
	}
	return nil
}

func (x *NewBlock) GetHeader() *BlockHeadersResponse {
	if x, ok := x.GetMaybeFull().(*NewBlock_Header); ok {
		return x.Header
	}
	return nil
}

func (x *NewBlock) GetBody() *BlockBodiesResponse {
	if x, ok := x.GetMaybeFull().(*NewBlock_Body); ok {
		return x.Body
	}
	return nil
}

type isNewBlock_MaybeFull interface {
	isNewBlock_MaybeFull()
}

type NewBlock_Id struct {
	Id *BlockID `protobuf:"bytes,1,opt,name=id,proto3,oneof"`
}

type NewBlock_Header struct {
	Header *BlockHeadersResponse `protobuf:"bytes,2,opt,name=header,proto3,oneof"`
}

type NewBlock_Body struct {
	Body *BlockBodiesResponse `protobuf:"bytes,3,opt,name=body,proto3,oneof"`
}

func (*NewBlock_Id) isNewBlock_MaybeFull() {}

func (*NewBlock_Header) isNewBlock_MaybeFull() {}

func (*NewBlock_Body) isNewBlock_MaybeFull() {}

// Requests a peer's CurrentBlockHeader
type CurrentBlockHeaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CurrentBlockHeaderRequest) Reset() {
	*x = CurrentBlockHeaderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_block_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrentBlockHeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrentBlockHeaderRequest) ProtoMessage() {}

func (x *CurrentBlockHeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_block_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrentBlockHeaderRequest.ProtoReflect.Descriptor instead.
func (*CurrentBlockHeaderRequest) Descriptor() ([]byte, []int) {
	return file_p2p_proto_block_proto_rawDescGZIP(), []int{4}
}

// result is (BlockHeader, Signature?)* in order of creation (incr/dec)
type BlockHeadersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iteration *Iteration `protobuf:"bytes,1,opt,name=iteration,proto3" json:"iteration,omitempty"`
}

func (x *BlockHeadersRequest) Reset() {
	*x = BlockHeadersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_block_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockHeadersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHeadersRequest) ProtoMessage() {}

func (x *BlockHeadersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_block_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHeadersRequest.ProtoReflect.Descriptor instead.
func (*BlockHeadersRequest) Descriptor() ([]byte, []int) {
	return file_p2p_proto_block_proto_rawDescGZIP(), []int{5}
}

func (x *BlockHeadersRequest) GetIteration() *Iteration {
	if x != nil {
		return x.Iteration
	}
	return nil
}

type BlockHeadersResponsePart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to HeaderMessage:
	//
	//	*BlockHeadersResponsePart_Header
	//	*BlockHeadersResponsePart_Signatures
	//	*BlockHeadersResponsePart_Fin
	HeaderMessage isBlockHeadersResponsePart_HeaderMessage `protobuf_oneof:"header_message"`
}

func (x *BlockHeadersResponsePart) Reset() {
	*x = BlockHeadersResponsePart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_block_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockHeadersResponsePart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHeadersResponsePart) ProtoMessage() {}

func (x *BlockHeadersResponsePart) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_block_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHeadersResponsePart.ProtoReflect.Descriptor instead.
func (*BlockHeadersResponsePart) Descriptor() ([]byte, []int) {
	return file_p2p_proto_block_proto_rawDescGZIP(), []int{6}
}

func (m *BlockHeadersResponsePart) GetHeaderMessage() isBlockHeadersResponsePart_HeaderMessage {
	if m != nil {
		return m.HeaderMessage
	}
	return nil
}

func (x *BlockHeadersResponsePart) GetHeader() *BlockHeader {
	if x, ok := x.GetHeaderMessage().(*BlockHeadersResponsePart_Header); ok {
		return x.Header
	}
	return nil
}

func (x *BlockHeadersResponsePart) GetSignatures() *Signatures {
	if x, ok := x.GetHeaderMessage().(*BlockHeadersResponsePart_Signatures); ok {
		return x.Signatures
	}
	return nil
}

func (x *BlockHeadersResponsePart) GetFin() *Fin {
	if x, ok := x.GetHeaderMessage().(*BlockHeadersResponsePart_Fin); ok {
		return x.Fin
	}
	return nil
}

type isBlockHeadersResponsePart_HeaderMessage interface {
	isBlockHeadersResponsePart_HeaderMessage()
}

type BlockHeadersResponsePart_Header struct {
	Header *BlockHeader `protobuf:"bytes,1,opt,name=header,proto3,oneof"`
}

type BlockHeadersResponsePart_Signatures struct {
	Signatures *Signatures `protobuf:"bytes,2,opt,name=signatures,proto3,oneof"`
}

type BlockHeadersResponsePart_Fin struct {
	Fin *Fin `protobuf:"bytes,3,opt,name=fin,proto3,oneof"` // no support for interleaving for now
}

func (*BlockHeadersResponsePart_Header) isBlockHeadersResponsePart_HeaderMessage() {}

func (*BlockHeadersResponsePart_Signatures) isBlockHeadersResponsePart_HeaderMessage() {}

func (*BlockHeadersResponsePart_Fin) isBlockHeadersResponsePart_HeaderMessage() {}

type BlockHeadersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Part []*BlockHeadersResponsePart `protobuf:"bytes,1,rep,name=part,proto3" json:"part,omitempty"`
}

func (x *BlockHeadersResponse) Reset() {
	*x = BlockHeadersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_block_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockHeadersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHeadersResponse) ProtoMessage() {}

func (x *BlockHeadersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_block_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHeadersResponse.ProtoReflect.Descriptor instead.
func (*BlockHeadersResponse) Descriptor() ([]byte, []int) {
	return file_p2p_proto_block_proto_rawDescGZIP(), []int{7}
}

func (x *BlockHeadersResponse) GetPart() []*BlockHeadersResponsePart {
	if x != nil {
		return x.Part
	}
	return nil
}

// result is (StateDiff*, Classes*, BlockProof?)* currently in creation order (incr/dec), but may change in the future
type BlockBodiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Iteration *Iteration `protobuf:"bytes,1,opt,name=iteration,proto3" json:"iteration,omitempty"`
}

func (x *BlockBodiesRequest) Reset() {
	*x = BlockBodiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_block_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockBodiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockBodiesRequest) ProtoMessage() {}

func (x *BlockBodiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_block_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockBodiesRequest.ProtoReflect.Descriptor instead.
func (*BlockBodiesRequest) Descriptor() ([]byte, []int) {
	return file_p2p_proto_block_proto_rawDescGZIP(), []int{8}
}

func (x *BlockBodiesRequest) GetIteration() *Iteration {
	if x != nil {
		return x.Iteration
	}
	return nil
}

type BlockBodiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *BlockID `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"` // may not appear if Fin is sent to end the whole response
	// Types that are assignable to BodyMessage:
	//
	//	*BlockBodiesResponse_Diff
	//	*BlockBodiesResponse_Classes
	//	*BlockBodiesResponse_Proof
	//	*BlockBodiesResponse_Fin
	BodyMessage isBlockBodiesResponse_BodyMessage `protobuf_oneof:"body_message"`
}

func (x *BlockBodiesResponse) Reset() {
	*x = BlockBodiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_p2p_proto_block_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockBodiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockBodiesResponse) ProtoMessage() {}

func (x *BlockBodiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_p2p_proto_block_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockBodiesResponse.ProtoReflect.Descriptor instead.
func (*BlockBodiesResponse) Descriptor() ([]byte, []int) {
	return file_p2p_proto_block_proto_rawDescGZIP(), []int{9}
}

func (x *BlockBodiesResponse) GetId() *BlockID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (m *BlockBodiesResponse) GetBodyMessage() isBlockBodiesResponse_BodyMessage {
	if m != nil {
		return m.BodyMessage
	}
	return nil
}

func (x *BlockBodiesResponse) GetDiff() *StateDiff {
	if x, ok := x.GetBodyMessage().(*BlockBodiesResponse_Diff); ok {
		return x.Diff
	}
	return nil
}

func (x *BlockBodiesResponse) GetClasses() *Classes {
	if x, ok := x.GetBodyMessage().(*BlockBodiesResponse_Classes); ok {
		return x.Classes
	}
	return nil
}

func (x *BlockBodiesResponse) GetProof() *BlockProof {
	if x, ok := x.GetBodyMessage().(*BlockBodiesResponse_Proof); ok {
		return x.Proof
	}
	return nil
}

func (x *BlockBodiesResponse) GetFin() *Fin {
	if x, ok := x.GetBodyMessage().(*BlockBodiesResponse_Fin); ok {
		return x.Fin
	}
	return nil
}

type isBlockBodiesResponse_BodyMessage interface {
	isBlockBodiesResponse_BodyMessage()
}

type BlockBodiesResponse_Diff struct {
	Diff *StateDiff `protobuf:"bytes,2,opt,name=diff,proto3,oneof"`
}

type BlockBodiesResponse_Classes struct {
	Classes *Classes `protobuf:"bytes,3,opt,name=classes,proto3,oneof"`
}

type BlockBodiesResponse_Proof struct {
	Proof *BlockProof `protobuf:"bytes,4,opt,name=proof,proto3,oneof"`
}

type BlockBodiesResponse_Fin struct {
	Fin *Fin `protobuf:"bytes,5,opt,name=fin,proto3,oneof"`
}

func (*BlockBodiesResponse_Diff) isBlockBodiesResponse_BodyMessage() {}

func (*BlockBodiesResponse_Classes) isBlockBodiesResponse_BodyMessage() {}

func (*BlockBodiesResponse_Proof) isBlockBodiesResponse_BodyMessage() {}

func (*BlockBodiesResponse_Fin) isBlockBodiesResponse_BodyMessage() {}

var File_p2p_proto_block_proto protoreflect.FileDescriptor

var file_p2p_proto_block_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x32, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x32, 0x70, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x70, 0x32, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x61, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x52, 0x05,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x33, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0a,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0xea, 0x03, 0x0a, 0x0b, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0b, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61,
	0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x11, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52,
	0x10, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x44, 0x69, 0x66, 0x66, 0x73, 0x12, 0x1f, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x50, 0x61, 0x74,
	0x72, 0x69, 0x63, 0x69, 0x61, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x6f, 0x66, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x05, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x46, 0x61,
	0x63, 0x74, 0x12, 0x2b, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4d, 0x65, 0x72, 0x6b, 0x6c,
	0x65, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x1f, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x23, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x08, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x25, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x46, 0x65, 0x6c, 0x74, 0x32, 0x35, 0x32, 0x52, 0x08, 0x67,
	0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x91, 0x01, 0x0a, 0x08,
	0x4e, 0x65, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x48, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x06, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x6f, 0x64, 0x69, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x42, 0x0c, 0x0a, 0x0a, 0x6d, 0x61, 0x79, 0x62, 0x65, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x22,
	0x1b, 0x0a, 0x19, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x13,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x09, 0x69, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9d, 0x01,
	0x0a, 0x18, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72, 0x74, 0x12, 0x26, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x12, 0x2d, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x03, 0x66, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x04,
	0x2e, 0x46, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x03, 0x66, 0x69, 0x6e, 0x42, 0x10, 0x0a, 0x0e, 0x68,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x45, 0x0a,
	0x14, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x70, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x72, 0x74, 0x52, 0x04,
	0x70, 0x61, 0x72, 0x74, 0x22, 0x3e, 0x0a, 0x12, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x6f, 0x64,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x09, 0x69, 0x74,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x49, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x69, 0x74, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xd2, 0x01, 0x0a, 0x13, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x6f,
	0x64, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x49, 0x44, 0x48, 0x01, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x04, 0x64,
	0x69, 0x66, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x44, 0x69, 0x66, 0x66, 0x48, 0x00, 0x52, 0x04, 0x64, 0x69, 0x66, 0x66, 0x12, 0x24, 0x0a,
	0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x48,
	0x00, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x18, 0x0a, 0x03, 0x66, 0x69, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x46, 0x69, 0x6e, 0x48, 0x00, 0x52, 0x03, 0x66,
	0x69, 0x6e, 0x42, 0x0e, 0x0a, 0x0c, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_p2p_proto_block_proto_rawDescOnce sync.Once
	file_p2p_proto_block_proto_rawDescData = file_p2p_proto_block_proto_rawDesc
)

func file_p2p_proto_block_proto_rawDescGZIP() []byte {
	file_p2p_proto_block_proto_rawDescOnce.Do(func() {
		file_p2p_proto_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_p2p_proto_block_proto_rawDescData)
	})
	return file_p2p_proto_block_proto_rawDescData
}

var (
	file_p2p_proto_block_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
	file_p2p_proto_block_proto_goTypes  = []interface{}{
		(*Signatures)(nil),                // 0: Signatures
		(*BlockHeader)(nil),               // 1: BlockHeader
		(*BlockProof)(nil),                // 2: BlockProof
		(*NewBlock)(nil),                  // 3: NewBlock
		(*CurrentBlockHeaderRequest)(nil), // 4: CurrentBlockHeaderRequest
		(*BlockHeadersRequest)(nil),       // 5: BlockHeadersRequest
		(*BlockHeadersResponsePart)(nil),  // 6: BlockHeadersResponsePart
		(*BlockHeadersResponse)(nil),      // 7: BlockHeadersResponse
		(*BlockBodiesRequest)(nil),        // 8: BlockBodiesRequest
		(*BlockBodiesResponse)(nil),       // 9: BlockBodiesResponse
		(*BlockID)(nil),                   // 10: BlockID
		(*ConsensusSignature)(nil),        // 11: ConsensusSignature
		(*Hash)(nil),                      // 12: Hash
		(*timestamppb.Timestamp)(nil),     // 13: google.protobuf.Timestamp
		(*Address)(nil),                   // 14: Address
		(*Merkle)(nil),                    // 15: Merkle
		(*Patricia)(nil),                  // 16: Patricia
		(*Felt252)(nil),                   // 17: Felt252
		(*Iteration)(nil),                 // 18: Iteration
		(*Fin)(nil),                       // 19: Fin
		(*StateDiff)(nil),                 // 20: StateDiff
		(*Classes)(nil),                   // 21: Classes
	}
)

var file_p2p_proto_block_proto_depIdxs = []int32{
	10, // 0: Signatures.block:type_name -> BlockID
	11, // 1: Signatures.signatures:type_name -> ConsensusSignature
	12, // 2: BlockHeader.parent_hash:type_name -> Hash
	13, // 3: BlockHeader.time:type_name -> google.protobuf.Timestamp
	14, // 4: BlockHeader.sequencer_address:type_name -> Address
	15, // 5: BlockHeader.state_diffs:type_name -> Merkle
	16, // 6: BlockHeader.state:type_name -> Patricia
	12, // 7: BlockHeader.proof_fact:type_name -> Hash
	15, // 8: BlockHeader.transactions:type_name -> Merkle
	15, // 9: BlockHeader.events:type_name -> Merkle
	15, // 10: BlockHeader.receipts:type_name -> Merkle
	17, // 11: BlockHeader.gas_price:type_name -> Felt252
	10, // 12: NewBlock.id:type_name -> BlockID
	7,  // 13: NewBlock.header:type_name -> BlockHeadersResponse
	9,  // 14: NewBlock.body:type_name -> BlockBodiesResponse
	18, // 15: BlockHeadersRequest.iteration:type_name -> Iteration
	1,  // 16: BlockHeadersResponsePart.header:type_name -> BlockHeader
	0,  // 17: BlockHeadersResponsePart.signatures:type_name -> Signatures
	19, // 18: BlockHeadersResponsePart.fin:type_name -> Fin
	6,  // 19: BlockHeadersResponse.part:type_name -> BlockHeadersResponsePart
	18, // 20: BlockBodiesRequest.iteration:type_name -> Iteration
	10, // 21: BlockBodiesResponse.id:type_name -> BlockID
	20, // 22: BlockBodiesResponse.diff:type_name -> StateDiff
	21, // 23: BlockBodiesResponse.classes:type_name -> Classes
	2,  // 24: BlockBodiesResponse.proof:type_name -> BlockProof
	19, // 25: BlockBodiesResponse.fin:type_name -> Fin
	26, // [26:26] is the sub-list for method output_type
	26, // [26:26] is the sub-list for method input_type
	26, // [26:26] is the sub-list for extension type_name
	26, // [26:26] is the sub-list for extension extendee
	0,  // [0:26] is the sub-list for field type_name
}

func init() { file_p2p_proto_block_proto_init() }
func file_p2p_proto_block_proto_init() {
	if File_p2p_proto_block_proto != nil {
		return
	}
	file_p2p_proto_common_proto_init()
	file_p2p_proto_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_p2p_proto_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signatures); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_p2p_proto_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockHeader); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_p2p_proto_block_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockProof); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_p2p_proto_block_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewBlock); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_p2p_proto_block_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrentBlockHeaderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_p2p_proto_block_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockHeadersRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_p2p_proto_block_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockHeadersResponsePart); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_p2p_proto_block_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockHeadersResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_p2p_proto_block_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockBodiesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_p2p_proto_block_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockBodiesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_p2p_proto_block_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*NewBlock_Id)(nil),
		(*NewBlock_Header)(nil),
		(*NewBlock_Body)(nil),
	}
	file_p2p_proto_block_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*BlockHeadersResponsePart_Header)(nil),
		(*BlockHeadersResponsePart_Signatures)(nil),
		(*BlockHeadersResponsePart_Fin)(nil),
	}
	file_p2p_proto_block_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*BlockBodiesResponse_Diff)(nil),
		(*BlockBodiesResponse_Classes)(nil),
		(*BlockBodiesResponse_Proof)(nil),
		(*BlockBodiesResponse_Fin)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_p2p_proto_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_p2p_proto_block_proto_goTypes,
		DependencyIndexes: file_p2p_proto_block_proto_depIdxs,
		MessageInfos:      file_p2p_proto_block_proto_msgTypes,
	}.Build()
	File_p2p_proto_block_proto = out.File
	file_p2p_proto_block_proto_rawDesc = nil
	file_p2p_proto_block_proto_goTypes = nil
	file_p2p_proto_block_proto_depIdxs = nil
}