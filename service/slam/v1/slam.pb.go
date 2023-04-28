// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: service/slam/v1/slam.proto

package v1

import (
	v1 "go.viam.com/api/common/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetPositionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of slam service
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPositionRequest) Reset() {
	*x = GetPositionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_slam_v1_slam_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPositionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPositionRequest) ProtoMessage() {}

func (x *GetPositionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_slam_v1_slam_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPositionRequest.ProtoReflect.Descriptor instead.
func (*GetPositionRequest) Descriptor() ([]byte, []int) {
	return file_service_slam_v1_slam_proto_rawDescGZIP(), []int{0}
}

func (x *GetPositionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetPositionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Current position of the specified component in the SLAM Map
	Pose *v1.Pose `protobuf:"bytes,1,opt,name=pose,proto3" json:"pose,omitempty"`
	// This is usually the name of the camera that is in the SLAM config
	ComponentReference string `protobuf:"bytes,2,opt,name=component_reference,json=componentReference,proto3" json:"component_reference,omitempty"`
	// Additional information in the response
	Extra *structpb.Struct `protobuf:"bytes,99,opt,name=extra,proto3" json:"extra,omitempty"`
}

func (x *GetPositionResponse) Reset() {
	*x = GetPositionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_slam_v1_slam_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPositionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPositionResponse) ProtoMessage() {}

func (x *GetPositionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_slam_v1_slam_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPositionResponse.ProtoReflect.Descriptor instead.
func (*GetPositionResponse) Descriptor() ([]byte, []int) {
	return file_service_slam_v1_slam_proto_rawDescGZIP(), []int{1}
}

func (x *GetPositionResponse) GetPose() *v1.Pose {
	if x != nil {
		return x.Pose
	}
	return nil
}

func (x *GetPositionResponse) GetComponentReference() string {
	if x != nil {
		return x.ComponentReference
	}
	return ""
}

func (x *GetPositionResponse) GetExtra() *structpb.Struct {
	if x != nil {
		return x.Extra
	}
	return nil
}

type GetPointCloudMapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of slam service
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPointCloudMapRequest) Reset() {
	*x = GetPointCloudMapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_slam_v1_slam_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPointCloudMapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPointCloudMapRequest) ProtoMessage() {}

func (x *GetPointCloudMapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_slam_v1_slam_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPointCloudMapRequest.ProtoReflect.Descriptor instead.
func (*GetPointCloudMapRequest) Descriptor() ([]byte, []int) {
	return file_service_slam_v1_slam_proto_rawDescGZIP(), []int{2}
}

func (x *GetPointCloudMapRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetPointCloudMapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// One chunk of the PointCloud.
	// For a given GetPointCloudMap request, concatenating all
	// GetPointCloudMapResponse.point_cloud_pcd_chunk values in the
	// order received result in the complete pointcloud in standard PCD
	// format where XY is the ground plane and positive Z is up, following
	// the Right Hand Rule.
	//
	// Read more about the pointcloud format here:
	// https://pointclouds.org/documentation/tutorials/pcd_file_format.html
	//
	// Viam expects pointcloud data with fields "x y z" or "x y z rgb", and for
	// this to be specified in the pointcloud header in the FIELDS entry. If color
	// data is included in the pointcloud, Viam's services assume that the color
	// value encodes a confidence score for that data point. Viam expects the
	// confidence score to be encoded in the blue parameter of the RGB value, on a
	// scale from 1-100.
	PointCloudPcdChunk []byte `protobuf:"bytes,1,opt,name=point_cloud_pcd_chunk,json=pointCloudPcdChunk,proto3" json:"point_cloud_pcd_chunk,omitempty"`
}

func (x *GetPointCloudMapResponse) Reset() {
	*x = GetPointCloudMapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_slam_v1_slam_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPointCloudMapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPointCloudMapResponse) ProtoMessage() {}

func (x *GetPointCloudMapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_slam_v1_slam_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPointCloudMapResponse.ProtoReflect.Descriptor instead.
func (*GetPointCloudMapResponse) Descriptor() ([]byte, []int) {
	return file_service_slam_v1_slam_proto_rawDescGZIP(), []int{3}
}

func (x *GetPointCloudMapResponse) GetPointCloudPcdChunk() []byte {
	if x != nil {
		return x.PointCloudPcdChunk
	}
	return nil
}

type GetInternalStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of slam service
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetInternalStateRequest) Reset() {
	*x = GetInternalStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_slam_v1_slam_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInternalStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInternalStateRequest) ProtoMessage() {}

func (x *GetInternalStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_slam_v1_slam_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInternalStateRequest.ProtoReflect.Descriptor instead.
func (*GetInternalStateRequest) Descriptor() ([]byte, []int) {
	return file_service_slam_v1_slam_proto_rawDescGZIP(), []int{4}
}

func (x *GetInternalStateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetInternalStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chunk of the internal state of the SLAM algorithm required to continue
	// mapping/localization
	InternalStateChunk []byte `protobuf:"bytes,1,opt,name=internal_state_chunk,json=internalStateChunk,proto3" json:"internal_state_chunk,omitempty"`
}

func (x *GetInternalStateResponse) Reset() {
	*x = GetInternalStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_slam_v1_slam_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInternalStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInternalStateResponse) ProtoMessage() {}

func (x *GetInternalStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_slam_v1_slam_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInternalStateResponse.ProtoReflect.Descriptor instead.
func (*GetInternalStateResponse) Descriptor() ([]byte, []int) {
	return file_service_slam_v1_slam_proto_rawDescGZIP(), []int{5}
}

func (x *GetInternalStateResponse) GetInternalStateChunk() []byte {
	if x != nil {
		return x.InternalStateChunk
	}
	return nil
}

var File_service_slam_v1_slam_proto protoreflect.FileDescriptor

var file_service_slam_v1_slam_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x6c, 0x61, 0x6d, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x6c, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x76, 0x69,
	0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x6c, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x1a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x9f, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x65, 0x52, 0x04, 0x70, 0x6f,
	0x73, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x5f,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x12, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x65, 0x78, 0x74, 0x72, 0x61, 0x18, 0x63, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x65, 0x78, 0x74,
	0x72, 0x61, 0x22, 0x2d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x4d, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a,
	0x15, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x63, 0x64,
	0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x63, 0x64, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x22, 0x2d, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x4c, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x32, 0x8c, 0x05,
	0x0a, 0x0b, 0x53, 0x4c, 0x41, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x95, 0x01,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x2e,
	0x76, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x6c, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x6c, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x12, 0x29, 0x2f, 0x76, 0x69, 0x61,
	0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x73, 0x6c, 0x61, 0x6d, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xad, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4d, 0x61, 0x70, 0x12, 0x2d, 0x2e, 0x76, 0x69, 0x61,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x6c, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4d,
	0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x76, 0x69, 0x61, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x6c, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x4d, 0x61,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x38, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x32, 0x12, 0x30, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x6c, 0x61, 0x6d, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0x2f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f,
	0x6d, 0x61, 0x70, 0x30, 0x01, 0x12, 0xac, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x76, 0x69, 0x61,
	0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x6c, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x76, 0x69, 0x61, 0x6d,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x6c, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x31, 0x12, 0x2f, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x6c, 0x61, 0x6d, 0x2f, 0x7b, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x30, 0x01, 0x12, 0x85, 0x01, 0x0a, 0x09, 0x44, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x20, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x22,
	0x2b, 0x2f, 0x76, 0x69, 0x61, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x73, 0x6c, 0x61, 0x6d, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65,
	0x7d, 0x2f, 0x64, 0x6f, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x42, 0x3b, 0x0a, 0x18,
	0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x69, 0x61, 0x6d, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x73, 0x6c, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x5a, 0x1f, 0x67, 0x6f, 0x2e, 0x76, 0x69, 0x61,
	0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x73, 0x6c, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_service_slam_v1_slam_proto_rawDescOnce sync.Once
	file_service_slam_v1_slam_proto_rawDescData = file_service_slam_v1_slam_proto_rawDesc
)

func file_service_slam_v1_slam_proto_rawDescGZIP() []byte {
	file_service_slam_v1_slam_proto_rawDescOnce.Do(func() {
		file_service_slam_v1_slam_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_slam_v1_slam_proto_rawDescData)
	})
	return file_service_slam_v1_slam_proto_rawDescData
}

var file_service_slam_v1_slam_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_slam_v1_slam_proto_goTypes = []interface{}{
	(*GetPositionRequest)(nil),       // 0: viam.service.slam.v1.GetPositionRequest
	(*GetPositionResponse)(nil),      // 1: viam.service.slam.v1.GetPositionResponse
	(*GetPointCloudMapRequest)(nil),  // 2: viam.service.slam.v1.GetPointCloudMapRequest
	(*GetPointCloudMapResponse)(nil), // 3: viam.service.slam.v1.GetPointCloudMapResponse
	(*GetInternalStateRequest)(nil),  // 4: viam.service.slam.v1.GetInternalStateRequest
	(*GetInternalStateResponse)(nil), // 5: viam.service.slam.v1.GetInternalStateResponse
	(*v1.Pose)(nil),                  // 6: viam.common.v1.Pose
	(*structpb.Struct)(nil),          // 7: google.protobuf.Struct
	(*v1.DoCommandRequest)(nil),      // 8: viam.common.v1.DoCommandRequest
	(*v1.DoCommandResponse)(nil),     // 9: viam.common.v1.DoCommandResponse
}
var file_service_slam_v1_slam_proto_depIdxs = []int32{
	6, // 0: viam.service.slam.v1.GetPositionResponse.pose:type_name -> viam.common.v1.Pose
	7, // 1: viam.service.slam.v1.GetPositionResponse.extra:type_name -> google.protobuf.Struct
	0, // 2: viam.service.slam.v1.SLAMService.GetPosition:input_type -> viam.service.slam.v1.GetPositionRequest
	2, // 3: viam.service.slam.v1.SLAMService.GetPointCloudMap:input_type -> viam.service.slam.v1.GetPointCloudMapRequest
	4, // 4: viam.service.slam.v1.SLAMService.GetInternalState:input_type -> viam.service.slam.v1.GetInternalStateRequest
	8, // 5: viam.service.slam.v1.SLAMService.DoCommand:input_type -> viam.common.v1.DoCommandRequest
	1, // 6: viam.service.slam.v1.SLAMService.GetPosition:output_type -> viam.service.slam.v1.GetPositionResponse
	3, // 7: viam.service.slam.v1.SLAMService.GetPointCloudMap:output_type -> viam.service.slam.v1.GetPointCloudMapResponse
	5, // 8: viam.service.slam.v1.SLAMService.GetInternalState:output_type -> viam.service.slam.v1.GetInternalStateResponse
	9, // 9: viam.service.slam.v1.SLAMService.DoCommand:output_type -> viam.common.v1.DoCommandResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_slam_v1_slam_proto_init() }
func file_service_slam_v1_slam_proto_init() {
	if File_service_slam_v1_slam_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_slam_v1_slam_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPositionRequest); i {
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
		file_service_slam_v1_slam_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPositionResponse); i {
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
		file_service_slam_v1_slam_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPointCloudMapRequest); i {
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
		file_service_slam_v1_slam_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPointCloudMapResponse); i {
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
		file_service_slam_v1_slam_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInternalStateRequest); i {
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
		file_service_slam_v1_slam_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInternalStateResponse); i {
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
			RawDescriptor: file_service_slam_v1_slam_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_slam_v1_slam_proto_goTypes,
		DependencyIndexes: file_service_slam_v1_slam_proto_depIdxs,
		MessageInfos:      file_service_slam_v1_slam_proto_msgTypes,
	}.Build()
	File_service_slam_v1_slam_proto = out.File
	file_service_slam_v1_slam_proto_rawDesc = nil
	file_service_slam_v1_slam_proto_goTypes = nil
	file_service_slam_v1_slam_proto_depIdxs = nil
}
