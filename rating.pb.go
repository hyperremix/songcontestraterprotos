// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.29.3
// source: rating.proto

package songcontestraterprotos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListUserRatingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *ListUserRatingsRequest) Reset() {
	*x = ListUserRatingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserRatingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserRatingsRequest) ProtoMessage() {}

func (x *ListUserRatingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rating_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserRatingsRequest.ProtoReflect.Descriptor instead.
func (*ListUserRatingsRequest) Descriptor() ([]byte, []int) {
	return file_rating_proto_rawDescGZIP(), []int{0}
}

func (x *ListUserRatingsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetRatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRatingRequest) Reset() {
	*x = GetRatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRatingRequest) ProtoMessage() {}

func (x *GetRatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rating_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRatingRequest.ProtoReflect.Descriptor instead.
func (*GetRatingRequest) Descriptor() ([]byte, []int) {
	return file_rating_proto_rawDescGZIP(), []int{1}
}

func (x *GetRatingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateRatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContestId string `protobuf:"bytes,1,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	ActId     string `protobuf:"bytes,2,opt,name=act_id,json=actId,proto3" json:"act_id,omitempty"`
	Song      int32  `protobuf:"varint,3,opt,name=song,proto3" json:"song,omitempty"`
	Singing   int32  `protobuf:"varint,4,opt,name=singing,proto3" json:"singing,omitempty"`
	Show      int32  `protobuf:"varint,5,opt,name=show,proto3" json:"show,omitempty"`
	Looks     int32  `protobuf:"varint,6,opt,name=looks,proto3" json:"looks,omitempty"`
	Clothes   int32  `protobuf:"varint,7,opt,name=clothes,proto3" json:"clothes,omitempty"`
}

func (x *CreateRatingRequest) Reset() {
	*x = CreateRatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRatingRequest) ProtoMessage() {}

func (x *CreateRatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rating_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRatingRequest.ProtoReflect.Descriptor instead.
func (*CreateRatingRequest) Descriptor() ([]byte, []int) {
	return file_rating_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRatingRequest) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *CreateRatingRequest) GetActId() string {
	if x != nil {
		return x.ActId
	}
	return ""
}

func (x *CreateRatingRequest) GetSong() int32 {
	if x != nil {
		return x.Song
	}
	return 0
}

func (x *CreateRatingRequest) GetSinging() int32 {
	if x != nil {
		return x.Singing
	}
	return 0
}

func (x *CreateRatingRequest) GetShow() int32 {
	if x != nil {
		return x.Show
	}
	return 0
}

func (x *CreateRatingRequest) GetLooks() int32 {
	if x != nil {
		return x.Looks
	}
	return 0
}

func (x *CreateRatingRequest) GetClothes() int32 {
	if x != nil {
		return x.Clothes
	}
	return 0
}

type UpdateRatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Song    int32  `protobuf:"varint,2,opt,name=song,proto3" json:"song,omitempty"`
	Singing int32  `protobuf:"varint,3,opt,name=singing,proto3" json:"singing,omitempty"`
	Show    int32  `protobuf:"varint,4,opt,name=show,proto3" json:"show,omitempty"`
	Looks   int32  `protobuf:"varint,5,opt,name=looks,proto3" json:"looks,omitempty"`
	Clothes int32  `protobuf:"varint,6,opt,name=clothes,proto3" json:"clothes,omitempty"`
}

func (x *UpdateRatingRequest) Reset() {
	*x = UpdateRatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRatingRequest) ProtoMessage() {}

func (x *UpdateRatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rating_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRatingRequest.ProtoReflect.Descriptor instead.
func (*UpdateRatingRequest) Descriptor() ([]byte, []int) {
	return file_rating_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRatingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateRatingRequest) GetSong() int32 {
	if x != nil {
		return x.Song
	}
	return 0
}

func (x *UpdateRatingRequest) GetSinging() int32 {
	if x != nil {
		return x.Singing
	}
	return 0
}

func (x *UpdateRatingRequest) GetShow() int32 {
	if x != nil {
		return x.Show
	}
	return 0
}

func (x *UpdateRatingRequest) GetLooks() int32 {
	if x != nil {
		return x.Looks
	}
	return 0
}

func (x *UpdateRatingRequest) GetClothes() int32 {
	if x != nil {
		return x.Clothes
	}
	return 0
}

type DeleteRatingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRatingRequest) Reset() {
	*x = DeleteRatingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRatingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRatingRequest) ProtoMessage() {}

func (x *DeleteRatingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rating_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRatingRequest.ProtoReflect.Descriptor instead.
func (*DeleteRatingRequest) Descriptor() ([]byte, []int) {
	return file_rating_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRatingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RatingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ContestId string                 `protobuf:"bytes,2,opt,name=contest_id,json=contestId,proto3" json:"contest_id,omitempty"`
	ActId     string                 `protobuf:"bytes,3,opt,name=act_id,json=actId,proto3" json:"act_id,omitempty"`
	Song      int32                  `protobuf:"varint,4,opt,name=song,proto3" json:"song,omitempty"`
	Singing   int32                  `protobuf:"varint,5,opt,name=singing,proto3" json:"singing,omitempty"`
	Show      int32                  `protobuf:"varint,6,opt,name=show,proto3" json:"show,omitempty"`
	Looks     int32                  `protobuf:"varint,7,opt,name=looks,proto3" json:"looks,omitempty"`
	Clothes   int32                  `protobuf:"varint,8,opt,name=clothes,proto3" json:"clothes,omitempty"`
	Total     int32                  `protobuf:"varint,9,opt,name=total,proto3" json:"total,omitempty"`
	User      *UserResponse          `protobuf:"bytes,10,opt,name=user,proto3" json:"user,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *RatingResponse) Reset() {
	*x = RatingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RatingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RatingResponse) ProtoMessage() {}

func (x *RatingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rating_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RatingResponse.ProtoReflect.Descriptor instead.
func (*RatingResponse) Descriptor() ([]byte, []int) {
	return file_rating_proto_rawDescGZIP(), []int{5}
}

func (x *RatingResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RatingResponse) GetContestId() string {
	if x != nil {
		return x.ContestId
	}
	return ""
}

func (x *RatingResponse) GetActId() string {
	if x != nil {
		return x.ActId
	}
	return ""
}

func (x *RatingResponse) GetSong() int32 {
	if x != nil {
		return x.Song
	}
	return 0
}

func (x *RatingResponse) GetSinging() int32 {
	if x != nil {
		return x.Singing
	}
	return 0
}

func (x *RatingResponse) GetShow() int32 {
	if x != nil {
		return x.Show
	}
	return 0
}

func (x *RatingResponse) GetLooks() int32 {
	if x != nil {
		return x.Looks
	}
	return 0
}

func (x *RatingResponse) GetClothes() int32 {
	if x != nil {
		return x.Clothes
	}
	return 0
}

func (x *RatingResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *RatingResponse) GetUser() *UserResponse {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *RatingResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *RatingResponse) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ListRatingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ratings []*RatingResponse `protobuf:"bytes,1,rep,name=ratings,proto3" json:"ratings,omitempty"`
}

func (x *ListRatingsResponse) Reset() {
	*x = ListRatingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rating_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRatingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRatingsResponse) ProtoMessage() {}

func (x *ListRatingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rating_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRatingsResponse.ProtoReflect.Descriptor instead.
func (*ListRatingsResponse) Descriptor() ([]byte, []int) {
	return file_rating_proto_rawDescGZIP(), []int{6}
}

func (x *ListRatingsResponse) GetRatings() []*RatingResponse {
	if x != nil {
		return x.Ratings
	}
	return nil
}

var File_rating_proto protoreflect.FileDescriptor

var file_rating_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x73, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x31, 0x0a,
	0x16, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0xbd, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x6f, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x69, 0x6e, 0x67, 0x69, 0x6e,
	0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x69, 0x6e, 0x67, 0x69, 0x6e, 0x67,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x68, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x73, 0x68, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c,
	0x6f, 0x74, 0x68, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6c, 0x6f,
	0x74, 0x68, 0x65, 0x73, 0x22, 0x97, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x6f, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x6e, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x69, 0x6e, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x73, 0x69, 0x6e, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x68,
	0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x77, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c,
	0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x6f, 0x74, 0x68, 0x65, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6c, 0x6f, 0x74, 0x68, 0x65, 0x73, 0x22, 0x25,
	0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x88, 0x03, 0x0a, 0x0e, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x6f, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x6f,
	0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x69, 0x6e, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x69, 0x6e, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x68, 0x6f, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x77,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x6c, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x6f, 0x74, 0x68, 0x65,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x63, 0x6c, 0x6f, 0x74, 0x68, 0x65, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x51, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x42, 0x1a, 0x5a, 0x18, 0x2e, 0x3b, 0x73, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rating_proto_rawDescOnce sync.Once
	file_rating_proto_rawDescData = file_rating_proto_rawDesc
)

func file_rating_proto_rawDescGZIP() []byte {
	file_rating_proto_rawDescOnce.Do(func() {
		file_rating_proto_rawDescData = protoimpl.X.CompressGZIP(file_rating_proto_rawDescData)
	})
	return file_rating_proto_rawDescData
}

var file_rating_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_rating_proto_goTypes = []any{
	(*ListUserRatingsRequest)(nil), // 0: songcontestrater.ListUserRatingsRequest
	(*GetRatingRequest)(nil),       // 1: songcontestrater.GetRatingRequest
	(*CreateRatingRequest)(nil),    // 2: songcontestrater.CreateRatingRequest
	(*UpdateRatingRequest)(nil),    // 3: songcontestrater.UpdateRatingRequest
	(*DeleteRatingRequest)(nil),    // 4: songcontestrater.DeleteRatingRequest
	(*RatingResponse)(nil),         // 5: songcontestrater.RatingResponse
	(*ListRatingsResponse)(nil),    // 6: songcontestrater.ListRatingsResponse
	(*UserResponse)(nil),           // 7: songcontestrater.UserResponse
	(*timestamppb.Timestamp)(nil),  // 8: google.protobuf.Timestamp
}
var file_rating_proto_depIdxs = []int32{
	7, // 0: songcontestrater.RatingResponse.user:type_name -> songcontestrater.UserResponse
	8, // 1: songcontestrater.RatingResponse.created_at:type_name -> google.protobuf.Timestamp
	8, // 2: songcontestrater.RatingResponse.updated_at:type_name -> google.protobuf.Timestamp
	5, // 3: songcontestrater.ListRatingsResponse.ratings:type_name -> songcontestrater.RatingResponse
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_rating_proto_init() }
func file_rating_proto_init() {
	if File_rating_proto != nil {
		return
	}
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rating_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListUserRatingsRequest); i {
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
		file_rating_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetRatingRequest); i {
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
		file_rating_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateRatingRequest); i {
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
		file_rating_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateRatingRequest); i {
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
		file_rating_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteRatingRequest); i {
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
		file_rating_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*RatingResponse); i {
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
		file_rating_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListRatingsResponse); i {
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
			RawDescriptor: file_rating_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rating_proto_goTypes,
		DependencyIndexes: file_rating_proto_depIdxs,
		MessageInfos:      file_rating_proto_msgTypes,
	}.Build()
	File_rating_proto = out.File
	file_rating_proto_rawDesc = nil
	file_rating_proto_goTypes = nil
	file_rating_proto_depIdxs = nil
}
