// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.0
// source: contents/movie.proto

package contents_grpc

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

type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year   string `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	ImdbID string `protobuf:"bytes,3,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Poster string `protobuf:"bytes,5,opt,name=poster,proto3" json:"poster,omitempty"`
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contents_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_contents_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_contents_movie_proto_rawDescGZIP(), []int{0}
}

func (x *Movie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Movie) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *Movie) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *Movie) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Movie) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

type MovieGetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Searchword string `protobuf:"bytes,1,opt,name=searchword,proto3" json:"searchword,omitempty"`
	Pagination int32  `protobuf:"varint,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *MovieGetListRequest) Reset() {
	*x = MovieGetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contents_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieGetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieGetListRequest) ProtoMessage() {}

func (x *MovieGetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contents_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieGetListRequest.ProtoReflect.Descriptor instead.
func (*MovieGetListRequest) Descriptor() ([]byte, []int) {
	return file_contents_movie_proto_rawDescGZIP(), []int{1}
}

func (x *MovieGetListRequest) GetSearchword() string {
	if x != nil {
		return x.Searchword
	}
	return ""
}

func (x *MovieGetListRequest) GetPagination() int32 {
	if x != nil {
		return x.Pagination
	}
	return 0
}

type MovieGetListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Movie `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *MovieGetListResponse) Reset() {
	*x = MovieGetListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contents_movie_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieGetListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieGetListResponse) ProtoMessage() {}

func (x *MovieGetListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contents_movie_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieGetListResponse.ProtoReflect.Descriptor instead.
func (*MovieGetListResponse) Descriptor() ([]byte, []int) {
	return file_contents_movie_proto_rawDescGZIP(), []int{2}
}

func (x *MovieGetListResponse) GetData() []*Movie {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_contents_movie_proto protoreflect.FileDescriptor

var file_contents_movie_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x22, 0x75, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x22, 0x55, 0x0a, 0x13, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3b,
	0x0a, 0x14, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x19, 0x5a, 0x17, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contents_movie_proto_rawDescOnce sync.Once
	file_contents_movie_proto_rawDescData = file_contents_movie_proto_rawDesc
)

func file_contents_movie_proto_rawDescGZIP() []byte {
	file_contents_movie_proto_rawDescOnce.Do(func() {
		file_contents_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_contents_movie_proto_rawDescData)
	})
	return file_contents_movie_proto_rawDescData
}

var file_contents_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_contents_movie_proto_goTypes = []interface{}{
	(*Movie)(nil),                // 0: contents.Movie
	(*MovieGetListRequest)(nil),  // 1: contents.MovieGetListRequest
	(*MovieGetListResponse)(nil), // 2: contents.MovieGetListResponse
}
var file_contents_movie_proto_depIdxs = []int32{
	0, // 0: contents.MovieGetListResponse.data:type_name -> contents.Movie
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_contents_movie_proto_init() }
func file_contents_movie_proto_init() {
	if File_contents_movie_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contents_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
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
		file_contents_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieGetListRequest); i {
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
		file_contents_movie_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieGetListResponse); i {
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
			RawDescriptor: file_contents_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contents_movie_proto_goTypes,
		DependencyIndexes: file_contents_movie_proto_depIdxs,
		MessageInfos:      file_contents_movie_proto_msgTypes,
	}.Build()
	File_contents_movie_proto = out.File
	file_contents_movie_proto_rawDesc = nil
	file_contents_movie_proto_goTypes = nil
	file_contents_movie_proto_depIdxs = nil
}
