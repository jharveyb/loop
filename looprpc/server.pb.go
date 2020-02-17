// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package looprpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type ServerLoopOutRequest struct {
	ReceiverKey []byte `protobuf:"bytes,1,opt,name=receiver_key,json=receiverKey,proto3" json:"receiver_key,omitempty"`
	SwapHash    []byte `protobuf:"bytes,2,opt,name=swap_hash,json=swapHash,proto3" json:"swap_hash,omitempty"`
	Amt         uint64 `protobuf:"varint,3,opt,name=amt,proto3" json:"amt,omitempty"`
	/// The unix time in seconds we want the on-chain swap to be published by.
	SwapPublicationDeadline int64    `protobuf:"varint,4,opt,name=swap_publication_deadline,json=swapPublicationDeadline,proto3" json:"swap_publication_deadline,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *ServerLoopOutRequest) Reset()         { *m = ServerLoopOutRequest{} }
func (m *ServerLoopOutRequest) String() string { return proto.CompactTextString(m) }
func (*ServerLoopOutRequest) ProtoMessage()    {}
func (*ServerLoopOutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *ServerLoopOutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerLoopOutRequest.Unmarshal(m, b)
}
func (m *ServerLoopOutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerLoopOutRequest.Marshal(b, m, deterministic)
}
func (m *ServerLoopOutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerLoopOutRequest.Merge(m, src)
}
func (m *ServerLoopOutRequest) XXX_Size() int {
	return xxx_messageInfo_ServerLoopOutRequest.Size(m)
}
func (m *ServerLoopOutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerLoopOutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerLoopOutRequest proto.InternalMessageInfo

func (m *ServerLoopOutRequest) GetReceiverKey() []byte {
	if m != nil {
		return m.ReceiverKey
	}
	return nil
}

func (m *ServerLoopOutRequest) GetSwapHash() []byte {
	if m != nil {
		return m.SwapHash
	}
	return nil
}

func (m *ServerLoopOutRequest) GetAmt() uint64 {
	if m != nil {
		return m.Amt
	}
	return 0
}

func (m *ServerLoopOutRequest) GetSwapPublicationDeadline() int64 {
	if m != nil {
		return m.SwapPublicationDeadline
	}
	return 0
}

type ServerLoopOutResponse struct {
	SwapInvoice          string   `protobuf:"bytes,1,opt,name=swap_invoice,json=swapInvoice,proto3" json:"swap_invoice,omitempty"`
	PrepayInvoice        string   `protobuf:"bytes,2,opt,name=prepay_invoice,json=prepayInvoice,proto3" json:"prepay_invoice,omitempty"`
	SenderKey            []byte   `protobuf:"bytes,3,opt,name=sender_key,json=senderKey,proto3" json:"sender_key,omitempty"`
	Expiry               int32    `protobuf:"varint,4,opt,name=expiry,proto3" json:"expiry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerLoopOutResponse) Reset()         { *m = ServerLoopOutResponse{} }
func (m *ServerLoopOutResponse) String() string { return proto.CompactTextString(m) }
func (*ServerLoopOutResponse) ProtoMessage()    {}
func (*ServerLoopOutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *ServerLoopOutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerLoopOutResponse.Unmarshal(m, b)
}
func (m *ServerLoopOutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerLoopOutResponse.Marshal(b, m, deterministic)
}
func (m *ServerLoopOutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerLoopOutResponse.Merge(m, src)
}
func (m *ServerLoopOutResponse) XXX_Size() int {
	return xxx_messageInfo_ServerLoopOutResponse.Size(m)
}
func (m *ServerLoopOutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerLoopOutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerLoopOutResponse proto.InternalMessageInfo

func (m *ServerLoopOutResponse) GetSwapInvoice() string {
	if m != nil {
		return m.SwapInvoice
	}
	return ""
}

func (m *ServerLoopOutResponse) GetPrepayInvoice() string {
	if m != nil {
		return m.PrepayInvoice
	}
	return ""
}

func (m *ServerLoopOutResponse) GetSenderKey() []byte {
	if m != nil {
		return m.SenderKey
	}
	return nil
}

func (m *ServerLoopOutResponse) GetExpiry() int32 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

type ServerLoopOutQuoteRequest struct {
	/// The swap amount. If zero, a quote for a maximum amt swap will be given.
	Amt uint64 `protobuf:"varint,1,opt,name=amt,proto3" json:"amt,omitempty"`
	/// The unix time in seconds we want the on-chain swap to be published by.
	SwapPublicationDeadline int64    `protobuf:"varint,2,opt,name=swap_publication_deadline,json=swapPublicationDeadline,proto3" json:"swap_publication_deadline,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *ServerLoopOutQuoteRequest) Reset()         { *m = ServerLoopOutQuoteRequest{} }
func (m *ServerLoopOutQuoteRequest) String() string { return proto.CompactTextString(m) }
func (*ServerLoopOutQuoteRequest) ProtoMessage()    {}
func (*ServerLoopOutQuoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{2}
}

func (m *ServerLoopOutQuoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerLoopOutQuoteRequest.Unmarshal(m, b)
}
func (m *ServerLoopOutQuoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerLoopOutQuoteRequest.Marshal(b, m, deterministic)
}
func (m *ServerLoopOutQuoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerLoopOutQuoteRequest.Merge(m, src)
}
func (m *ServerLoopOutQuoteRequest) XXX_Size() int {
	return xxx_messageInfo_ServerLoopOutQuoteRequest.Size(m)
}
func (m *ServerLoopOutQuoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerLoopOutQuoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerLoopOutQuoteRequest proto.InternalMessageInfo

func (m *ServerLoopOutQuoteRequest) GetAmt() uint64 {
	if m != nil {
		return m.Amt
	}
	return 0
}

func (m *ServerLoopOutQuoteRequest) GetSwapPublicationDeadline() int64 {
	if m != nil {
		return m.SwapPublicationDeadline
	}
	return 0
}

type ServerLoopOutQuote struct {
	SwapPaymentDest string `protobuf:"bytes,1,opt,name=swap_payment_dest,json=swapPaymentDest,proto3" json:"swap_payment_dest,omitempty"`
	/// The total estimated swap fee given the quote amt.
	SwapFee int64 `protobuf:"varint,2,opt,name=swap_fee,json=swapFee,proto3" json:"swap_fee,omitempty"`
	/// Deprecated, total swap fee given quote amt is calculated in swap_fee.
	SwapFeeRate          int64    `protobuf:"varint,3,opt,name=swap_fee_rate,json=swapFeeRate,proto3" json:"swap_fee_rate,omitempty"` // Deprecated: Do not use.
	PrepayAmt            uint64   `protobuf:"varint,4,opt,name=prepay_amt,json=prepayAmt,proto3" json:"prepay_amt,omitempty"`
	MinSwapAmount        uint64   `protobuf:"varint,5,opt,name=min_swap_amount,json=minSwapAmount,proto3" json:"min_swap_amount,omitempty"` // Deprecated: Do not use.
	MaxSwapAmount        uint64   `protobuf:"varint,6,opt,name=max_swap_amount,json=maxSwapAmount,proto3" json:"max_swap_amount,omitempty"` // Deprecated: Do not use.
	CltvDelta            int32    `protobuf:"varint,7,opt,name=cltv_delta,json=cltvDelta,proto3" json:"cltv_delta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerLoopOutQuote) Reset()         { *m = ServerLoopOutQuote{} }
func (m *ServerLoopOutQuote) String() string { return proto.CompactTextString(m) }
func (*ServerLoopOutQuote) ProtoMessage()    {}
func (*ServerLoopOutQuote) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{3}
}

func (m *ServerLoopOutQuote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerLoopOutQuote.Unmarshal(m, b)
}
func (m *ServerLoopOutQuote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerLoopOutQuote.Marshal(b, m, deterministic)
}
func (m *ServerLoopOutQuote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerLoopOutQuote.Merge(m, src)
}
func (m *ServerLoopOutQuote) XXX_Size() int {
	return xxx_messageInfo_ServerLoopOutQuote.Size(m)
}
func (m *ServerLoopOutQuote) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerLoopOutQuote.DiscardUnknown(m)
}

var xxx_messageInfo_ServerLoopOutQuote proto.InternalMessageInfo

func (m *ServerLoopOutQuote) GetSwapPaymentDest() string {
	if m != nil {
		return m.SwapPaymentDest
	}
	return ""
}

func (m *ServerLoopOutQuote) GetSwapFee() int64 {
	if m != nil {
		return m.SwapFee
	}
	return 0
}

// Deprecated: Do not use.
func (m *ServerLoopOutQuote) GetSwapFeeRate() int64 {
	if m != nil {
		return m.SwapFeeRate
	}
	return 0
}

func (m *ServerLoopOutQuote) GetPrepayAmt() uint64 {
	if m != nil {
		return m.PrepayAmt
	}
	return 0
}

// Deprecated: Do not use.
func (m *ServerLoopOutQuote) GetMinSwapAmount() uint64 {
	if m != nil {
		return m.MinSwapAmount
	}
	return 0
}

// Deprecated: Do not use.
func (m *ServerLoopOutQuote) GetMaxSwapAmount() uint64 {
	if m != nil {
		return m.MaxSwapAmount
	}
	return 0
}

func (m *ServerLoopOutQuote) GetCltvDelta() int32 {
	if m != nil {
		return m.CltvDelta
	}
	return 0
}

type ServerLoopOutTermsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerLoopOutTermsRequest) Reset()         { *m = ServerLoopOutTermsRequest{} }
func (m *ServerLoopOutTermsRequest) String() string { return proto.CompactTextString(m) }
func (*ServerLoopOutTermsRequest) ProtoMessage()    {}
func (*ServerLoopOutTermsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{4}
}

func (m *ServerLoopOutTermsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerLoopOutTermsRequest.Unmarshal(m, b)
}
func (m *ServerLoopOutTermsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerLoopOutTermsRequest.Marshal(b, m, deterministic)
}
func (m *ServerLoopOutTermsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerLoopOutTermsRequest.Merge(m, src)
}
func (m *ServerLoopOutTermsRequest) XXX_Size() int {
	return xxx_messageInfo_ServerLoopOutTermsRequest.Size(m)
}
func (m *ServerLoopOutTermsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerLoopOutTermsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerLoopOutTermsRequest proto.InternalMessageInfo

type ServerLoopOutTerms struct {
	MinSwapAmount        uint64   `protobuf:"varint,1,opt,name=min_swap_amount,json=minSwapAmount,proto3" json:"min_swap_amount,omitempty"`
	MaxSwapAmount        uint64   `protobuf:"varint,2,opt,name=max_swap_amount,json=maxSwapAmount,proto3" json:"max_swap_amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerLoopOutTerms) Reset()         { *m = ServerLoopOutTerms{} }
func (m *ServerLoopOutTerms) String() string { return proto.CompactTextString(m) }
func (*ServerLoopOutTerms) ProtoMessage()    {}
func (*ServerLoopOutTerms) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{5}
}

func (m *ServerLoopOutTerms) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerLoopOutTerms.Unmarshal(m, b)
}
func (m *ServerLoopOutTerms) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerLoopOutTerms.Marshal(b, m, deterministic)
}
func (m *ServerLoopOutTerms) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerLoopOutTerms.Merge(m, src)
}
func (m *ServerLoopOutTerms) XXX_Size() int {
	return xxx_messageInfo_ServerLoopOutTerms.Size(m)
}
func (m *ServerLoopOutTerms) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerLoopOutTerms.DiscardUnknown(m)
}

var xxx_messageInfo_ServerLoopOutTerms proto.InternalMessageInfo

func (m *ServerLoopOutTerms) GetMinSwapAmount() uint64 {
	if m != nil {
		return m.MinSwapAmount
	}
	return 0
}

func (m *ServerLoopOutTerms) GetMaxSwapAmount() uint64 {
	if m != nil {
		return m.MaxSwapAmount
	}
	return 0
}

type ServerLoopInRequest struct {
	SenderKey            []byte   `protobuf:"bytes,1,opt,name=sender_key,json=senderKey,proto3" json:"sender_key,omitempty"`
	SwapHash             []byte   `protobuf:"bytes,2,opt,name=swap_hash,json=swapHash,proto3" json:"swap_hash,omitempty"`
	Amt                  uint64   `protobuf:"varint,3,opt,name=amt,proto3" json:"amt,omitempty"`
	SwapInvoice          string   `protobuf:"bytes,4,opt,name=swap_invoice,json=swapInvoice,proto3" json:"swap_invoice,omitempty"`
	LastHop              []byte   `protobuf:"bytes,5,opt,name=last_hop,json=lastHop,proto3" json:"last_hop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerLoopInRequest) Reset()         { *m = ServerLoopInRequest{} }
func (m *ServerLoopInRequest) String() string { return proto.CompactTextString(m) }
func (*ServerLoopInRequest) ProtoMessage()    {}
func (*ServerLoopInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{6}
}

func (m *ServerLoopInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerLoopInRequest.Unmarshal(m, b)
}
func (m *ServerLoopInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerLoopInRequest.Marshal(b, m, deterministic)
}
func (m *ServerLoopInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerLoopInRequest.Merge(m, src)
}
func (m *ServerLoopInRequest) XXX_Size() int {
	return xxx_messageInfo_ServerLoopInRequest.Size(m)
}
func (m *ServerLoopInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerLoopInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerLoopInRequest proto.InternalMessageInfo

func (m *ServerLoopInRequest) GetSenderKey() []byte {
	if m != nil {
		return m.SenderKey
	}
	return nil
}

func (m *ServerLoopInRequest) GetSwapHash() []byte {
	if m != nil {
		return m.SwapHash
	}
	return nil
}

func (m *ServerLoopInRequest) GetAmt() uint64 {
	if m != nil {
		return m.Amt
	}
	return 0
}

func (m *ServerLoopInRequest) GetSwapInvoice() string {
	if m != nil {
		return m.SwapInvoice
	}
	return ""
}

func (m *ServerLoopInRequest) GetLastHop() []byte {
	if m != nil {
		return m.LastHop
	}
	return nil
}

type ServerLoopInResponse struct {
	ReceiverKey          []byte   `protobuf:"bytes,1,opt,name=receiver_key,json=receiverKey,proto3" json:"receiver_key,omitempty"`
	Expiry               int32    `protobuf:"varint,2,opt,name=expiry,proto3" json:"expiry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerLoopInResponse) Reset()         { *m = ServerLoopInResponse{} }
func (m *ServerLoopInResponse) String() string { return proto.CompactTextString(m) }
func (*ServerLoopInResponse) ProtoMessage()    {}
func (*ServerLoopInResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{7}
}

func (m *ServerLoopInResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerLoopInResponse.Unmarshal(m, b)
}
func (m *ServerLoopInResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerLoopInResponse.Marshal(b, m, deterministic)
}
func (m *ServerLoopInResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerLoopInResponse.Merge(m, src)
}
func (m *ServerLoopInResponse) XXX_Size() int {
	return xxx_messageInfo_ServerLoopInResponse.Size(m)
}
func (m *ServerLoopInResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerLoopInResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerLoopInResponse proto.InternalMessageInfo

func (m *ServerLoopInResponse) GetReceiverKey() []byte {
	if m != nil {
		return m.ReceiverKey
	}
	return nil
}

func (m *ServerLoopInResponse) GetExpiry() int32 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

type ServerLoopInQuoteRequest struct {
	/// The swap amount. If zero, a quote for a maximum amt swap will be given.
	Amt                  uint64   `protobuf:"varint,1,opt,name=amt,proto3" json:"amt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerLoopInQuoteRequest) Reset()         { *m = ServerLoopInQuoteRequest{} }
func (m *ServerLoopInQuoteRequest) String() string { return proto.CompactTextString(m) }
func (*ServerLoopInQuoteRequest) ProtoMessage()    {}
func (*ServerLoopInQuoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{8}
}

func (m *ServerLoopInQuoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerLoopInQuoteRequest.Unmarshal(m, b)
}
func (m *ServerLoopInQuoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerLoopInQuoteRequest.Marshal(b, m, deterministic)
}
func (m *ServerLoopInQuoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerLoopInQuoteRequest.Merge(m, src)
}
func (m *ServerLoopInQuoteRequest) XXX_Size() int {
	return xxx_messageInfo_ServerLoopInQuoteRequest.Size(m)
}
func (m *ServerLoopInQuoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerLoopInQuoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerLoopInQuoteRequest proto.InternalMessageInfo

func (m *ServerLoopInQuoteRequest) GetAmt() uint64 {
	if m != nil {
		return m.Amt
	}
	return 0
}

type ServerLoopInQuoteResponse struct {
	SwapFee              int64    `protobuf:"varint,1,opt,name=swap_fee,json=swapFee,proto3" json:"swap_fee,omitempty"`
	SwapFeeRate          int64    `protobuf:"varint,2,opt,name=swap_fee_rate,json=swapFeeRate,proto3" json:"swap_fee_rate,omitempty"`       // Deprecated: Do not use.
	MinSwapAmount        uint64   `protobuf:"varint,4,opt,name=min_swap_amount,json=minSwapAmount,proto3" json:"min_swap_amount,omitempty"` // Deprecated: Do not use.
	MaxSwapAmount        uint64   `protobuf:"varint,5,opt,name=max_swap_amount,json=maxSwapAmount,proto3" json:"max_swap_amount,omitempty"` // Deprecated: Do not use.
	CltvDelta            int32    `protobuf:"varint,6,opt,name=cltv_delta,json=cltvDelta,proto3" json:"cltv_delta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerLoopInQuoteResponse) Reset()         { *m = ServerLoopInQuoteResponse{} }
func (m *ServerLoopInQuoteResponse) String() string { return proto.CompactTextString(m) }
func (*ServerLoopInQuoteResponse) ProtoMessage()    {}
func (*ServerLoopInQuoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{9}
}

func (m *ServerLoopInQuoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerLoopInQuoteResponse.Unmarshal(m, b)
}
func (m *ServerLoopInQuoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerLoopInQuoteResponse.Marshal(b, m, deterministic)
}
func (m *ServerLoopInQuoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerLoopInQuoteResponse.Merge(m, src)
}
func (m *ServerLoopInQuoteResponse) XXX_Size() int {
	return xxx_messageInfo_ServerLoopInQuoteResponse.Size(m)
}
func (m *ServerLoopInQuoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerLoopInQuoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ServerLoopInQuoteResponse proto.InternalMessageInfo

func (m *ServerLoopInQuoteResponse) GetSwapFee() int64 {
	if m != nil {
		return m.SwapFee
	}
	return 0
}

// Deprecated: Do not use.
func (m *ServerLoopInQuoteResponse) GetSwapFeeRate() int64 {
	if m != nil {
		return m.SwapFeeRate
	}
	return 0
}

// Deprecated: Do not use.
func (m *ServerLoopInQuoteResponse) GetMinSwapAmount() uint64 {
	if m != nil {
		return m.MinSwapAmount
	}
	return 0
}

// Deprecated: Do not use.
func (m *ServerLoopInQuoteResponse) GetMaxSwapAmount() uint64 {
	if m != nil {
		return m.MaxSwapAmount
	}
	return 0
}

func (m *ServerLoopInQuoteResponse) GetCltvDelta() int32 {
	if m != nil {
		return m.CltvDelta
	}
	return 0
}

type ServerLoopInTermsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerLoopInTermsRequest) Reset()         { *m = ServerLoopInTermsRequest{} }
func (m *ServerLoopInTermsRequest) String() string { return proto.CompactTextString(m) }
func (*ServerLoopInTermsRequest) ProtoMessage()    {}
func (*ServerLoopInTermsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{10}
}

func (m *ServerLoopInTermsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerLoopInTermsRequest.Unmarshal(m, b)
}
func (m *ServerLoopInTermsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerLoopInTermsRequest.Marshal(b, m, deterministic)
}
func (m *ServerLoopInTermsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerLoopInTermsRequest.Merge(m, src)
}
func (m *ServerLoopInTermsRequest) XXX_Size() int {
	return xxx_messageInfo_ServerLoopInTermsRequest.Size(m)
}
func (m *ServerLoopInTermsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerLoopInTermsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ServerLoopInTermsRequest proto.InternalMessageInfo

type ServerLoopInTerms struct {
	MinSwapAmount        uint64   `protobuf:"varint,1,opt,name=min_swap_amount,json=minSwapAmount,proto3" json:"min_swap_amount,omitempty"`
	MaxSwapAmount        uint64   `protobuf:"varint,2,opt,name=max_swap_amount,json=maxSwapAmount,proto3" json:"max_swap_amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServerLoopInTerms) Reset()         { *m = ServerLoopInTerms{} }
func (m *ServerLoopInTerms) String() string { return proto.CompactTextString(m) }
func (*ServerLoopInTerms) ProtoMessage()    {}
func (*ServerLoopInTerms) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{11}
}

func (m *ServerLoopInTerms) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerLoopInTerms.Unmarshal(m, b)
}
func (m *ServerLoopInTerms) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerLoopInTerms.Marshal(b, m, deterministic)
}
func (m *ServerLoopInTerms) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerLoopInTerms.Merge(m, src)
}
func (m *ServerLoopInTerms) XXX_Size() int {
	return xxx_messageInfo_ServerLoopInTerms.Size(m)
}
func (m *ServerLoopInTerms) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerLoopInTerms.DiscardUnknown(m)
}

var xxx_messageInfo_ServerLoopInTerms proto.InternalMessageInfo

func (m *ServerLoopInTerms) GetMinSwapAmount() uint64 {
	if m != nil {
		return m.MinSwapAmount
	}
	return 0
}

func (m *ServerLoopInTerms) GetMaxSwapAmount() uint64 {
	if m != nil {
		return m.MaxSwapAmount
	}
	return 0
}

func init() {
	proto.RegisterType((*ServerLoopOutRequest)(nil), "looprpc.ServerLoopOutRequest")
	proto.RegisterType((*ServerLoopOutResponse)(nil), "looprpc.ServerLoopOutResponse")
	proto.RegisterType((*ServerLoopOutQuoteRequest)(nil), "looprpc.ServerLoopOutQuoteRequest")
	proto.RegisterType((*ServerLoopOutQuote)(nil), "looprpc.ServerLoopOutQuote")
	proto.RegisterType((*ServerLoopOutTermsRequest)(nil), "looprpc.ServerLoopOutTermsRequest")
	proto.RegisterType((*ServerLoopOutTerms)(nil), "looprpc.ServerLoopOutTerms")
	proto.RegisterType((*ServerLoopInRequest)(nil), "looprpc.ServerLoopInRequest")
	proto.RegisterType((*ServerLoopInResponse)(nil), "looprpc.ServerLoopInResponse")
	proto.RegisterType((*ServerLoopInQuoteRequest)(nil), "looprpc.ServerLoopInQuoteRequest")
	proto.RegisterType((*ServerLoopInQuoteResponse)(nil), "looprpc.ServerLoopInQuoteResponse")
	proto.RegisterType((*ServerLoopInTermsRequest)(nil), "looprpc.ServerLoopInTermsRequest")
	proto.RegisterType((*ServerLoopInTerms)(nil), "looprpc.ServerLoopInTerms")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x4e, 0xdb, 0x4c,
	0x10, 0x95, 0x9d, 0x90, 0x90, 0x21, 0x81, 0x8f, 0xfd, 0xfa, 0xe3, 0x04, 0x52, 0x81, 0xa5, 0x22,
	0x84, 0x2a, 0x90, 0xda, 0xbb, 0xde, 0x51, 0x21, 0x04, 0x2a, 0x2a, 0xc5, 0x70, 0x6f, 0x2d, 0xc9,
	0x94, 0x58, 0xb5, 0xbd, 0x5b, 0x7b, 0x13, 0xc8, 0x9b, 0xb4, 0x97, 0x95, 0xfa, 0x4a, 0x7d, 0x90,
	0xbe, 0x41, 0xb5, 0x3f, 0x26, 0xb6, 0xf3, 0x03, 0x48, 0xbd, 0x63, 0x67, 0x0e, 0x33, 0x67, 0xce,
	0x9c, 0x89, 0xa1, 0x99, 0x62, 0x32, 0xc2, 0x64, 0x9f, 0x27, 0x4c, 0x30, 0x52, 0x0f, 0x19, 0xe3,
	0x09, 0xef, 0x75, 0x36, 0x6f, 0x18, 0xbb, 0x09, 0xf1, 0x80, 0xf2, 0xe0, 0x80, 0xc6, 0x31, 0x13,
	0x54, 0x04, 0x2c, 0x4e, 0x35, 0xcc, 0xfd, 0x65, 0xc1, 0xb3, 0x4b, 0xf5, 0x7f, 0x67, 0x8c, 0xf1,
	0xf3, 0xa1, 0xf0, 0xf0, 0xdb, 0x10, 0x53, 0x41, 0xb6, 0xa1, 0x99, 0x60, 0x0f, 0x83, 0x11, 0x26,
	0xfe, 0x57, 0x1c, 0x3b, 0xd6, 0x96, 0xb5, 0xdb, 0xf4, 0x56, 0xb2, 0xd8, 0x47, 0x1c, 0x93, 0x0d,
	0x68, 0xa4, 0xb7, 0x94, 0xfb, 0x03, 0x9a, 0x0e, 0x1c, 0x5b, 0xe5, 0x97, 0x65, 0xe0, 0x84, 0xa6,
	0x03, 0xf2, 0x1f, 0x54, 0x68, 0x24, 0x9c, 0xca, 0x96, 0xb5, 0x5b, 0xf5, 0xe4, 0x9f, 0xe4, 0x3d,
	0xb4, 0x15, 0x9c, 0x0f, 0xaf, 0xc3, 0xa0, 0xa7, 0x58, 0xf8, 0x7d, 0xa4, 0xfd, 0x30, 0x88, 0xd1,
	0xa9, 0x6e, 0x59, 0xbb, 0x15, 0xef, 0xa5, 0x04, 0x7c, 0x9e, 0xe4, 0x8f, 0x4c, 0xda, 0xfd, 0x6e,
	0xc1, 0xf3, 0x12, 0xcd, 0x94, 0xb3, 0x38, 0x45, 0xc9, 0x53, 0x55, 0x0d, 0xe2, 0x11, 0x0b, 0x7a,
	0xa8, 0x78, 0x36, 0xbc, 0x15, 0x19, 0x3b, 0xd5, 0x21, 0xf2, 0x1a, 0x56, 0x79, 0x82, 0x9c, 0x8e,
	0xef, 0x41, 0xb6, 0x02, 0xb5, 0x74, 0x34, 0x83, 0x75, 0x01, 0x52, 0x8c, 0xfb, 0x66, 0xde, 0x8a,
	0x9a, 0xa7, 0xa1, 0x23, 0x72, 0xda, 0x17, 0x50, 0xc3, 0x3b, 0x1e, 0x24, 0x63, 0xc5, 0x75, 0xc9,
	0x33, 0x2f, 0x37, 0x80, 0x76, 0x81, 0xd9, 0xc5, 0x90, 0x09, 0xcc, 0x54, 0x34, 0x2a, 0x58, 0x8f,
	0x54, 0xc1, 0x5e, 0xac, 0xc2, 0x0f, 0x1b, 0xc8, 0x74, 0x2f, 0xb2, 0x07, 0xeb, 0xba, 0x24, 0x1d,
	0x47, 0x18, 0x0b, 0xbf, 0x8f, 0xa9, 0x30, 0x3a, 0xac, 0xa9, 0x52, 0x3a, 0x7e, 0x24, 0x09, 0xb5,
	0x41, 0xad, 0xc8, 0xff, 0x82, 0x59, 0xb7, 0xba, 0x7c, 0x1f, 0x23, 0x92, 0x1d, 0x68, 0x65, 0x29,
	0x3f, 0xa1, 0x02, 0x95, 0x04, 0x95, 0x0f, 0xb6, 0x63, 0x69, 0x39, 0x8f, 0x11, 0x3d, 0x2a, 0x94,
	0x4e, 0x46, 0x4e, 0x39, 0x5a, 0x55, 0x8d, 0xd6, 0xd0, 0x91, 0xc3, 0x48, 0x90, 0x3d, 0x58, 0x8b,
	0x82, 0xd8, 0x57, 0xa5, 0x68, 0xc4, 0x86, 0xb1, 0x70, 0x96, 0x24, 0x46, 0x15, 0x6a, 0x45, 0x41,
	0x7c, 0x79, 0x4b, 0xf9, 0xa1, 0x4a, 0x28, 0x2c, 0xbd, 0x2b, 0x60, 0x6b, 0x39, 0x2c, 0xbd, 0xcb,
	0x61, 0xbb, 0x00, 0xbd, 0x50, 0x8c, 0xfc, 0x3e, 0x86, 0x82, 0x3a, 0x75, 0xb5, 0x83, 0x86, 0x8c,
	0x1c, 0xc9, 0x80, 0xbb, 0x51, 0x5a, 0xc3, 0x15, 0x26, 0x51, 0x6a, 0xd6, 0xe0, 0xf6, 0x4b, 0xba,
	0xa9, 0x24, 0xd9, 0x99, 0x66, 0xaa, 0x17, 0x55, 0x62, 0xb9, 0x33, 0xcd, 0xd2, 0x36, 0xb8, 0x3c,
	0x43, 0xf7, 0xa7, 0x05, 0xff, 0x4f, 0xda, 0x9c, 0xc6, 0x99, 0x09, 0x8a, 0xc6, 0xb2, 0xca, 0xc6,
	0x7a, 0xe2, 0x19, 0x95, 0x0d, 0x5f, 0x9d, 0x36, 0x7c, 0x1b, 0x96, 0x43, 0x9a, 0x0a, 0x7f, 0xc0,
	0xb8, 0xd2, 0xbe, 0xe9, 0xd5, 0xe5, 0xfb, 0x84, 0x71, 0xf7, 0x22, 0x7f, 0xee, 0x92, 0xe2, 0xe4,
	0x8c, 0x1e, 0x3a, 0xf7, 0xc9, 0x01, 0xd8, 0x85, 0x03, 0x78, 0x03, 0x4e, 0xbe, 0xe4, 0x62, 0xff,
	0xbb, 0xbf, 0xad, 0xfc, 0xa2, 0xee, 0xe1, 0x86, 0x46, 0xde, 0x9e, 0xd6, 0x03, 0xf6, 0xb4, 0x67,
	0xdb, 0x73, 0x86, 0xff, 0xaa, 0x4f, 0xf0, 0xdf, 0xd2, 0xe3, 0xfc, 0x57, 0x2b, 0xfb, 0xaf, 0x53,
	0x54, 0xa1, 0x60, 0xbf, 0x1e, 0xac, 0x4f, 0xe5, 0xfe, 0xb5, 0xfb, 0xde, 0xfe, 0xa9, 0x00, 0xc8,
	0xa7, 0xee, 0x44, 0xce, 0xa1, 0x59, 0x30, 0xbb, 0xbb, 0x6f, 0x3e, 0x08, 0xfb, 0x73, 0xcf, 0xa4,
	0xb3, 0xb1, 0x00, 0x43, 0xce, 0x61, 0xf5, 0x13, 0xde, 0x9a, 0x90, 0x6c, 0x44, 0xba, 0xb3, 0xe1,
	0x59, 0xb5, 0x57, 0xf3, 0xd2, 0x66, 0xd7, 0x13, 0x86, 0xfa, 0x67, 0x6c, 0x0e, 0xc3, 0xbc, 0x9f,
	0xe6, 0x31, 0xd4, 0x05, 0xce, 0x60, 0x25, 0x2f, 0xf0, 0xf6, 0x0c, 0x6c, 0x71, 0x31, 0x9d, 0xce,
	0x7c, 0x08, 0x39, 0x83, 0x96, 0x99, 0xf7, 0x54, 0xad, 0x83, 0x6c, 0xce, 0x04, 0x67, 0xa5, 0xba,
	0x73, 0xb2, 0x66, 0xd8, 0xab, 0x8c, 0x9b, 0xa6, 0x3a, 0x9b, 0x5b, 0x61, 0x54, 0x77, 0x11, 0x44,
	0x57, 0xbd, 0xae, 0xa9, 0x8f, 0xf8, 0xbb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x99, 0x4f,
	0x5f, 0xfb, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SwapServerClient is the client API for SwapServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SwapServerClient interface {
	LoopOutTerms(ctx context.Context, in *ServerLoopOutTermsRequest, opts ...grpc.CallOption) (*ServerLoopOutTerms, error)
	NewLoopOutSwap(ctx context.Context, in *ServerLoopOutRequest, opts ...grpc.CallOption) (*ServerLoopOutResponse, error)
	LoopOutQuote(ctx context.Context, in *ServerLoopOutQuoteRequest, opts ...grpc.CallOption) (*ServerLoopOutQuote, error)
	LoopInTerms(ctx context.Context, in *ServerLoopInTermsRequest, opts ...grpc.CallOption) (*ServerLoopInTerms, error)
	NewLoopInSwap(ctx context.Context, in *ServerLoopInRequest, opts ...grpc.CallOption) (*ServerLoopInResponse, error)
	LoopInQuote(ctx context.Context, in *ServerLoopInQuoteRequest, opts ...grpc.CallOption) (*ServerLoopInQuoteResponse, error)
}

type swapServerClient struct {
	cc *grpc.ClientConn
}

func NewSwapServerClient(cc *grpc.ClientConn) SwapServerClient {
	return &swapServerClient{cc}
}

func (c *swapServerClient) LoopOutTerms(ctx context.Context, in *ServerLoopOutTermsRequest, opts ...grpc.CallOption) (*ServerLoopOutTerms, error) {
	out := new(ServerLoopOutTerms)
	err := c.cc.Invoke(ctx, "/looprpc.SwapServer/LoopOutTerms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapServerClient) NewLoopOutSwap(ctx context.Context, in *ServerLoopOutRequest, opts ...grpc.CallOption) (*ServerLoopOutResponse, error) {
	out := new(ServerLoopOutResponse)
	err := c.cc.Invoke(ctx, "/looprpc.SwapServer/NewLoopOutSwap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapServerClient) LoopOutQuote(ctx context.Context, in *ServerLoopOutQuoteRequest, opts ...grpc.CallOption) (*ServerLoopOutQuote, error) {
	out := new(ServerLoopOutQuote)
	err := c.cc.Invoke(ctx, "/looprpc.SwapServer/LoopOutQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapServerClient) LoopInTerms(ctx context.Context, in *ServerLoopInTermsRequest, opts ...grpc.CallOption) (*ServerLoopInTerms, error) {
	out := new(ServerLoopInTerms)
	err := c.cc.Invoke(ctx, "/looprpc.SwapServer/LoopInTerms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapServerClient) NewLoopInSwap(ctx context.Context, in *ServerLoopInRequest, opts ...grpc.CallOption) (*ServerLoopInResponse, error) {
	out := new(ServerLoopInResponse)
	err := c.cc.Invoke(ctx, "/looprpc.SwapServer/NewLoopInSwap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapServerClient) LoopInQuote(ctx context.Context, in *ServerLoopInQuoteRequest, opts ...grpc.CallOption) (*ServerLoopInQuoteResponse, error) {
	out := new(ServerLoopInQuoteResponse)
	err := c.cc.Invoke(ctx, "/looprpc.SwapServer/LoopInQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SwapServerServer is the server API for SwapServer service.
type SwapServerServer interface {
	LoopOutTerms(context.Context, *ServerLoopOutTermsRequest) (*ServerLoopOutTerms, error)
	NewLoopOutSwap(context.Context, *ServerLoopOutRequest) (*ServerLoopOutResponse, error)
	LoopOutQuote(context.Context, *ServerLoopOutQuoteRequest) (*ServerLoopOutQuote, error)
	LoopInTerms(context.Context, *ServerLoopInTermsRequest) (*ServerLoopInTerms, error)
	NewLoopInSwap(context.Context, *ServerLoopInRequest) (*ServerLoopInResponse, error)
	LoopInQuote(context.Context, *ServerLoopInQuoteRequest) (*ServerLoopInQuoteResponse, error)
}

func RegisterSwapServerServer(s *grpc.Server, srv SwapServerServer) {
	s.RegisterService(&_SwapServer_serviceDesc, srv)
}

func _SwapServer_LoopOutTerms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerLoopOutTermsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapServerServer).LoopOutTerms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/looprpc.SwapServer/LoopOutTerms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapServerServer).LoopOutTerms(ctx, req.(*ServerLoopOutTermsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapServer_NewLoopOutSwap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerLoopOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapServerServer).NewLoopOutSwap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/looprpc.SwapServer/NewLoopOutSwap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapServerServer).NewLoopOutSwap(ctx, req.(*ServerLoopOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapServer_LoopOutQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerLoopOutQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapServerServer).LoopOutQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/looprpc.SwapServer/LoopOutQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapServerServer).LoopOutQuote(ctx, req.(*ServerLoopOutQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapServer_LoopInTerms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerLoopInTermsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapServerServer).LoopInTerms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/looprpc.SwapServer/LoopInTerms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapServerServer).LoopInTerms(ctx, req.(*ServerLoopInTermsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapServer_NewLoopInSwap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerLoopInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapServerServer).NewLoopInSwap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/looprpc.SwapServer/NewLoopInSwap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapServerServer).NewLoopInSwap(ctx, req.(*ServerLoopInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapServer_LoopInQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerLoopInQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapServerServer).LoopInQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/looprpc.SwapServer/LoopInQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapServerServer).LoopInQuote(ctx, req.(*ServerLoopInQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SwapServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "looprpc.SwapServer",
	HandlerType: (*SwapServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoopOutTerms",
			Handler:    _SwapServer_LoopOutTerms_Handler,
		},
		{
			MethodName: "NewLoopOutSwap",
			Handler:    _SwapServer_NewLoopOutSwap_Handler,
		},
		{
			MethodName: "LoopOutQuote",
			Handler:    _SwapServer_LoopOutQuote_Handler,
		},
		{
			MethodName: "LoopInTerms",
			Handler:    _SwapServer_LoopInTerms_Handler,
		},
		{
			MethodName: "NewLoopInSwap",
			Handler:    _SwapServer_NewLoopInSwap_Handler,
		},
		{
			MethodName: "LoopInQuote",
			Handler:    _SwapServer_LoopInQuote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
