// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: waves/node/grpc/assets_api.proto

package grpc

import (
	waves "github.com/wavesplatform/gowaves/pkg/grpc/generated/waves"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AssetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetId []byte `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
}

func (x *AssetRequest) Reset() {
	*x = AssetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waves_node_grpc_assets_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetRequest) ProtoMessage() {}

func (x *AssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_waves_node_grpc_assets_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetRequest.ProtoReflect.Descriptor instead.
func (*AssetRequest) Descriptor() ([]byte, []int) {
	return file_waves_node_grpc_assets_api_proto_rawDescGZIP(), []int{0}
}

func (x *AssetRequest) GetAssetId() []byte {
	if x != nil {
		return x.AssetId
	}
	return nil
}

type NFTRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address      []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Limit        int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	AfterAssetId []byte `protobuf:"bytes,3,opt,name=after_asset_id,json=afterAssetId,proto3" json:"after_asset_id,omitempty"`
}

func (x *NFTRequest) Reset() {
	*x = NFTRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waves_node_grpc_assets_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NFTRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NFTRequest) ProtoMessage() {}

func (x *NFTRequest) ProtoReflect() protoreflect.Message {
	mi := &file_waves_node_grpc_assets_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NFTRequest.ProtoReflect.Descriptor instead.
func (*NFTRequest) Descriptor() ([]byte, []int) {
	return file_waves_node_grpc_assets_api_proto_rawDescGZIP(), []int{1}
}

func (x *NFTRequest) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *NFTRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *NFTRequest) GetAfterAssetId() []byte {
	if x != nil {
		return x.AfterAssetId
	}
	return nil
}

type NFTResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetId   []byte             `protobuf:"bytes,1,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	AssetInfo *AssetInfoResponse `protobuf:"bytes,2,opt,name=asset_info,json=assetInfo,proto3" json:"asset_info,omitempty"`
}

func (x *NFTResponse) Reset() {
	*x = NFTResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waves_node_grpc_assets_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NFTResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NFTResponse) ProtoMessage() {}

func (x *NFTResponse) ProtoReflect() protoreflect.Message {
	mi := &file_waves_node_grpc_assets_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NFTResponse.ProtoReflect.Descriptor instead.
func (*NFTResponse) Descriptor() ([]byte, []int) {
	return file_waves_node_grpc_assets_api_proto_rawDescGZIP(), []int{2}
}

func (x *NFTResponse) GetAssetId() []byte {
	if x != nil {
		return x.AssetId
	}
	return nil
}

func (x *NFTResponse) GetAssetInfo() *AssetInfoResponse {
	if x != nil {
		return x.AssetInfo
	}
	return nil
}

type AssetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Issuer           []byte                   `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Name             string                   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description      string                   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Decimals         int32                    `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Reissuable       bool                     `protobuf:"varint,5,opt,name=reissuable,proto3" json:"reissuable,omitempty"`
	TotalVolume      int64                    `protobuf:"varint,6,opt,name=total_volume,json=totalVolume,proto3" json:"total_volume,omitempty"`
	Script           *ScriptData              `protobuf:"bytes,7,opt,name=script,proto3" json:"script,omitempty"`
	Sponsorship      int64                    `protobuf:"varint,8,opt,name=sponsorship,proto3" json:"sponsorship,omitempty"`
	IssueTransaction *waves.SignedTransaction `protobuf:"bytes,11,opt,name=issue_transaction,json=issueTransaction,proto3" json:"issue_transaction,omitempty"`
	SponsorBalance   int64                    `protobuf:"varint,10,opt,name=sponsor_balance,json=sponsorBalance,proto3" json:"sponsor_balance,omitempty"`
}

func (x *AssetInfoResponse) Reset() {
	*x = AssetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waves_node_grpc_assets_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AssetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetInfoResponse) ProtoMessage() {}

func (x *AssetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_waves_node_grpc_assets_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetInfoResponse.ProtoReflect.Descriptor instead.
func (*AssetInfoResponse) Descriptor() ([]byte, []int) {
	return file_waves_node_grpc_assets_api_proto_rawDescGZIP(), []int{3}
}

func (x *AssetInfoResponse) GetIssuer() []byte {
	if x != nil {
		return x.Issuer
	}
	return nil
}

func (x *AssetInfoResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AssetInfoResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AssetInfoResponse) GetDecimals() int32 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

func (x *AssetInfoResponse) GetReissuable() bool {
	if x != nil {
		return x.Reissuable
	}
	return false
}

func (x *AssetInfoResponse) GetTotalVolume() int64 {
	if x != nil {
		return x.TotalVolume
	}
	return 0
}

func (x *AssetInfoResponse) GetScript() *ScriptData {
	if x != nil {
		return x.Script
	}
	return nil
}

func (x *AssetInfoResponse) GetSponsorship() int64 {
	if x != nil {
		return x.Sponsorship
	}
	return 0
}

func (x *AssetInfoResponse) GetIssueTransaction() *waves.SignedTransaction {
	if x != nil {
		return x.IssueTransaction
	}
	return nil
}

func (x *AssetInfoResponse) GetSponsorBalance() int64 {
	if x != nil {
		return x.SponsorBalance
	}
	return 0
}

var File_waves_node_grpc_assets_api_proto protoreflect.FileDescriptor

var file_waves_node_grpc_assets_api_proto_rawDesc = []byte{
	0x0a, 0x20, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0f, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x1a, 0x17, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x77, 0x61,
	0x76, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x29, 0x0a, 0x0c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x22, 0x62, 0x0a, 0x0a, 0x4e,
	0x46, 0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x66, 0x74,
	0x65, 0x72, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0c, 0x61, 0x66, 0x74, 0x65, 0x72, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x22,
	0x6b, 0x0a, 0x0b, 0x4e, 0x46, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x0a, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x77, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x87, 0x03, 0x0a,
	0x11, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x69, 0x73, 0x73, 0x75, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x72, 0x65, 0x69, 0x73, 0x73, 0x75, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12,
	0x33, 0x0a, 0x06, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x6f, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x6f, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x45, 0x0a, 0x11, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x0a,
	0x0f, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x6f, 0x72, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x32, 0xa4, 0x01, 0x0a, 0x09, 0x41, 0x73, 0x73, 0x65, 0x74,
	0x73, 0x41, 0x70, 0x69, 0x12, 0x4c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1d, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4e, 0x46, 0x54, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1b, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x4e, 0x46, 0x54, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x77, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4e, 0x46, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x73, 0x0a,
	0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x61, 0x76, 0x65, 0x73, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5a, 0x43, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x61, 0x76, 0x65, 0x73, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x67, 0x6f, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x2f, 0x77, 0x61, 0x76, 0x65, 0x73, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0xaa, 0x02, 0x0f, 0x57, 0x61, 0x76, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x47, 0x72,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_waves_node_grpc_assets_api_proto_rawDescOnce sync.Once
	file_waves_node_grpc_assets_api_proto_rawDescData = file_waves_node_grpc_assets_api_proto_rawDesc
)

func file_waves_node_grpc_assets_api_proto_rawDescGZIP() []byte {
	file_waves_node_grpc_assets_api_proto_rawDescOnce.Do(func() {
		file_waves_node_grpc_assets_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_waves_node_grpc_assets_api_proto_rawDescData)
	})
	return file_waves_node_grpc_assets_api_proto_rawDescData
}

var file_waves_node_grpc_assets_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_waves_node_grpc_assets_api_proto_goTypes = []interface{}{
	(*AssetRequest)(nil),            // 0: waves.node.grpc.AssetRequest
	(*NFTRequest)(nil),              // 1: waves.node.grpc.NFTRequest
	(*NFTResponse)(nil),             // 2: waves.node.grpc.NFTResponse
	(*AssetInfoResponse)(nil),       // 3: waves.node.grpc.AssetInfoResponse
	(*ScriptData)(nil),              // 4: waves.node.grpc.ScriptData
	(*waves.SignedTransaction)(nil), // 5: waves.SignedTransaction
}
var file_waves_node_grpc_assets_api_proto_depIdxs = []int32{
	3, // 0: waves.node.grpc.NFTResponse.asset_info:type_name -> waves.node.grpc.AssetInfoResponse
	4, // 1: waves.node.grpc.AssetInfoResponse.script:type_name -> waves.node.grpc.ScriptData
	5, // 2: waves.node.grpc.AssetInfoResponse.issue_transaction:type_name -> waves.SignedTransaction
	0, // 3: waves.node.grpc.AssetsApi.GetInfo:input_type -> waves.node.grpc.AssetRequest
	1, // 4: waves.node.grpc.AssetsApi.GetNFTList:input_type -> waves.node.grpc.NFTRequest
	3, // 5: waves.node.grpc.AssetsApi.GetInfo:output_type -> waves.node.grpc.AssetInfoResponse
	2, // 6: waves.node.grpc.AssetsApi.GetNFTList:output_type -> waves.node.grpc.NFTResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_waves_node_grpc_assets_api_proto_init() }
func file_waves_node_grpc_assets_api_proto_init() {
	if File_waves_node_grpc_assets_api_proto != nil {
		return
	}
	file_waves_node_grpc_accounts_api_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_waves_node_grpc_assets_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetRequest); i {
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
		file_waves_node_grpc_assets_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NFTRequest); i {
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
		file_waves_node_grpc_assets_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NFTResponse); i {
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
		file_waves_node_grpc_assets_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AssetInfoResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_waves_node_grpc_assets_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_waves_node_grpc_assets_api_proto_goTypes,
		DependencyIndexes: file_waves_node_grpc_assets_api_proto_depIdxs,
		MessageInfos:      file_waves_node_grpc_assets_api_proto_msgTypes,
	}.Build()
	File_waves_node_grpc_assets_api_proto = out.File
	file_waves_node_grpc_assets_api_proto_rawDesc = nil
	file_waves_node_grpc_assets_api_proto_goTypes = nil
	file_waves_node_grpc_assets_api_proto_depIdxs = nil
}