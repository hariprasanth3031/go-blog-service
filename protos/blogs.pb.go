// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: protos/blogs.proto

package go_blog_service

import (
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

type PostDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              uint64   `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title           string   `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Content         string   `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
	Author          string   `protobuf:"bytes,4,opt,name=Author,proto3" json:"Author,omitempty"`
	PublicationDate uint64   `protobuf:"varint,5,opt,name=PublicationDate,proto3" json:"PublicationDate,omitempty"`
	Tags            []string `protobuf:"bytes,6,rep,name=Tags,proto3" json:"Tags,omitempty"`
}

func (x *PostDetails) Reset() {
	*x = PostDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_blogs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostDetails) ProtoMessage() {}

func (x *PostDetails) ProtoReflect() protoreflect.Message {
	mi := &file_protos_blogs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostDetails.ProtoReflect.Descriptor instead.
func (*PostDetails) Descriptor() ([]byte, []int) {
	return file_protos_blogs_proto_rawDescGZIP(), []int{0}
}

func (x *PostDetails) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PostDetails) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostDetails) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *PostDetails) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *PostDetails) GetPublicationDate() uint64 {
	if x != nil {
		return x.PublicationDate
	}
	return 0
}

func (x *PostDetails) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type CreatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *PostDetails `protobuf:"bytes,1,opt,name=Post,proto3" json:"Post,omitempty"`
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_blogs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_blogs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_protos_blogs_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePostRequest) GetPost() *PostDetails {
	if x != nil {
		return x.Post
	}
	return nil
}

type PostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post   *PostDetails `protobuf:"bytes,1,opt,name=Post,proto3" json:"Post,omitempty"`
	Status string       `protobuf:"bytes,2,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *PostResponse) Reset() {
	*x = PostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_blogs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostResponse) ProtoMessage() {}

func (x *PostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_blogs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostResponse.ProtoReflect.Descriptor instead.
func (*PostResponse) Descriptor() ([]byte, []int) {
	return file_protos_blogs_proto_rawDescGZIP(), []int{2}
}

func (x *PostResponse) GetPost() *PostDetails {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *PostResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type PostId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId int64 `protobuf:"varint,1,opt,name=PostId,proto3" json:"PostId,omitempty"`
}

func (x *PostId) Reset() {
	*x = PostId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_blogs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostId) ProtoMessage() {}

func (x *PostId) ProtoReflect() protoreflect.Message {
	mi := &file_protos_blogs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostId.ProtoReflect.Descriptor instead.
func (*PostId) Descriptor() ([]byte, []int) {
	return file_protos_blogs_proto_rawDescGZIP(), []int{3}
}

func (x *PostId) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

type UpdatePost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *PostId      `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	UpdateInput *PostDetails `protobuf:"bytes,2,opt,name=UpdateInput,proto3" json:"UpdateInput,omitempty"`
}

func (x *UpdatePost) Reset() {
	*x = UpdatePost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_blogs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePost) ProtoMessage() {}

func (x *UpdatePost) ProtoReflect() protoreflect.Message {
	mi := &file_protos_blogs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePost.ProtoReflect.Descriptor instead.
func (*UpdatePost) Descriptor() ([]byte, []int) {
	return file_protos_blogs_proto_rawDescGZIP(), []int{4}
}

func (x *UpdatePost) GetId() *PostId {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *UpdatePost) GetUpdateInput() *PostDetails {
	if x != nil {
		return x.UpdateInput
	}
	return nil
}

type DeletePost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (x *DeletePost) Reset() {
	*x = DeletePost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_blogs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePost) ProtoMessage() {}

func (x *DeletePost) ProtoReflect() protoreflect.Message {
	mi := &file_protos_blogs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePost.ProtoReflect.Descriptor instead.
func (*DeletePost) Descriptor() ([]byte, []int) {
	return file_protos_blogs_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePost) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_protos_blogs_proto protoreflect.FileDescriptor

var file_protos_blogs_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x28, 0x0a, 0x0f,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x61, 0x67, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x54, 0x61, 0x67, 0x73, 0x22, 0x35, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x04, 0x50, 0x6f, 0x73,
	0x74, 0x22, 0x48, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x04, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x20, 0x0a, 0x06, 0x50,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x55, 0x0a,
	0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x22, 0x24, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x99, 0x01, 0x0a, 0x05, 0x42,
	0x6c, 0x6f, 0x67, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x07, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x1a, 0x0d, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x24, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x12, 0x07, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x1a, 0x0b, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x72, 0x69, 0x70, 0x72, 0x61, 0x73, 0x61, 0x6e, 0x74,
	0x68, 0x33, 0x30, 0x33, 0x31, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x6c, 0x6f, 0x67, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_blogs_proto_rawDescOnce sync.Once
	file_protos_blogs_proto_rawDescData = file_protos_blogs_proto_rawDesc
)

func file_protos_blogs_proto_rawDescGZIP() []byte {
	file_protos_blogs_proto_rawDescOnce.Do(func() {
		file_protos_blogs_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_blogs_proto_rawDescData)
	})
	return file_protos_blogs_proto_rawDescData
}

var file_protos_blogs_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_blogs_proto_goTypes = []interface{}{
	(*PostDetails)(nil),       // 0: PostDetails
	(*CreatePostRequest)(nil), // 1: CreatePostRequest
	(*PostResponse)(nil),      // 2: PostResponse
	(*PostId)(nil),            // 3: PostId
	(*UpdatePost)(nil),        // 4: UpdatePost
	(*DeletePost)(nil),        // 5: DeletePost
}
var file_protos_blogs_proto_depIdxs = []int32{
	0, // 0: CreatePostRequest.Post:type_name -> PostDetails
	0, // 1: PostResponse.Post:type_name -> PostDetails
	3, // 2: UpdatePost.Id:type_name -> PostId
	0, // 3: UpdatePost.UpdateInput:type_name -> PostDetails
	1, // 4: Blogs.Create:input_type -> CreatePostRequest
	3, // 5: Blogs.Get:input_type -> PostId
	4, // 6: Blogs.Update:input_type -> UpdatePost
	3, // 7: Blogs.Delete:input_type -> PostId
	2, // 8: Blogs.Create:output_type -> PostResponse
	2, // 9: Blogs.Get:output_type -> PostResponse
	2, // 10: Blogs.Update:output_type -> PostResponse
	5, // 11: Blogs.Delete:output_type -> DeletePost
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protos_blogs_proto_init() }
func file_protos_blogs_proto_init() {
	if File_protos_blogs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_blogs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostDetails); i {
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
		file_protos_blogs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostRequest); i {
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
		file_protos_blogs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostResponse); i {
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
		file_protos_blogs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostId); i {
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
		file_protos_blogs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePost); i {
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
		file_protos_blogs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePost); i {
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
			RawDescriptor: file_protos_blogs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_blogs_proto_goTypes,
		DependencyIndexes: file_protos_blogs_proto_depIdxs,
		MessageInfos:      file_protos_blogs_proto_msgTypes,
	}.Build()
	File_protos_blogs_proto = out.File
	file_protos_blogs_proto_rawDesc = nil
	file_protos_blogs_proto_goTypes = nil
	file_protos_blogs_proto_depIdxs = nil
}
