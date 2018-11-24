// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gbookshelf.proto

package gbookshelf

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BookState struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Page                 int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Done                 bool     `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BookState) Reset()         { *m = BookState{} }
func (m *BookState) String() string { return proto.CompactTextString(m) }
func (*BookState) ProtoMessage()    {}
func (*BookState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ba0627d911171ab, []int{0}
}

func (m *BookState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BookState.Unmarshal(m, b)
}
func (m *BookState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BookState.Marshal(b, m, deterministic)
}
func (m *BookState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookState.Merge(m, src)
}
func (m *BookState) XXX_Size() int {
	return xxx_messageInfo_BookState.Size(m)
}
func (m *BookState) XXX_DiscardUnknown() {
	xxx_messageInfo_BookState.DiscardUnknown(m)
}

var xxx_messageInfo_BookState proto.InternalMessageInfo

func (m *BookState) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *BookState) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *BookState) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

type Books struct {
	Books                []*BookState `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Books) Reset()         { *m = Books{} }
func (m *Books) String() string { return proto.CompactTextString(m) }
func (*Books) ProtoMessage()    {}
func (*Books) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ba0627d911171ab, []int{1}
}

func (m *Books) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Books.Unmarshal(m, b)
}
func (m *Books) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Books.Marshal(b, m, deterministic)
}
func (m *Books) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Books.Merge(m, src)
}
func (m *Books) XXX_Size() int {
	return xxx_messageInfo_Books.Size(m)
}
func (m *Books) XXX_DiscardUnknown() {
	xxx_messageInfo_Books.DiscardUnknown(m)
}

var xxx_messageInfo_Books proto.InternalMessageInfo

func (m *Books) GetBooks() []*BookState {
	if m != nil {
		return m.Books
	}
	return nil
}

type Book struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Page                 int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ba0627d911171ab, []int{2}
}

func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ba0627d911171ab, []int{3}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BookState)(nil), "gbookshelf.BookState")
	proto.RegisterType((*Books)(nil), "gbookshelf.Books")
	proto.RegisterType((*Book)(nil), "gbookshelf.Book")
	proto.RegisterType((*Void)(nil), "gbookshelf.Void")
}

func init() { proto.RegisterFile("gbookshelf.proto", fileDescriptor_7ba0627d911171ab) }

var fileDescriptor_7ba0627d911171ab = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x31, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0x37, 0x6e, 0x12, 0xbc, 0xb1, 0x39, 0x07, 0x85, 0x70, 0x55, 0x48, 0x15, 0x10, 0x97,
	0xe3, 0xb4, 0xb0, 0xd5, 0x4e, 0xb0, 0x8a, 0x68, 0x7f, 0x47, 0xe2, 0xb9, 0xdc, 0xea, 0x2c, 0x26,
	0xf8, 0x53, 0xfd, 0x3d, 0x92, 0xac, 0xe8, 0xaa, 0x2b, 0xd8, 0x4d, 0x5e, 0x5e, 0xde, 0xf7, 0xc8,
	0xc0, 0x7c, 0xbb, 0x21, 0xda, 0xc5, 0xc7, 0xd0, 0x3d, 0x34, 0xfd, 0x0b, 0x25, 0x42, 0xf8, 0x52,
	0xcc, 0x35, 0xcc, 0xae, 0x88, 0x76, 0xb7, 0x69, 0x9d, 0x02, 0x1e, 0x81, 0x48, 0x6d, 0xea, 0x82,
	0x62, 0x9a, 0xd9, 0x99, 0x1b, 0x0e, 0x88, 0xc0, 0xfb, 0xf5, 0x36, 0xa8, 0x3d, 0xcd, 0xac, 0x70,
	0x65, 0xce, 0x9a, 0xa7, 0xe7, 0xa0, 0x6a, 0xcd, 0xec, 0xbe, 0x2b, 0xb3, 0x39, 0x07, 0x91, 0xa3,
	0x22, 0x9e, 0x80, 0x28, 0x00, 0xc5, 0x74, 0x6d, 0x0f, 0x56, 0xc7, 0xcd, 0xa8, 0xc1, 0x27, 0xcc,
	0x0d, 0x1e, 0xb3, 0x04, 0x9e, 0xb5, 0xff, 0xb3, 0x8d, 0x04, 0x7e, 0x4f, 0xad, 0x5f, 0xbd, 0xb1,
	0x8f, 0xee, 0x39, 0x18, 0x4f, 0x81, 0xdf, 0xb4, 0x31, 0xe1, 0x7c, 0x4c, 0xcb, 0xbe, 0xc5, 0xe1,
	0x4f, 0x7e, 0x34, 0x15, 0x2e, 0xa1, 0xbe, 0xf4, 0xfe, 0xbb, 0x3b, 0xdf, 0x2d, 0xa6, 0xdb, 0x9a,
	0x0a, 0x1b, 0x90, 0x2e, 0x3c, 0xd1, 0x6b, 0x98, 0x78, 0xf4, 0x0b, 0x6a, 0x2a, 0xbc, 0x00, 0x79,
	0xd7, 0xfb, 0xfc, 0xad, 0xd3, 0x91, 0x7f, 0x92, 0x36, 0xb2, 0xac, 0xe9, 0xec, 0x3d, 0x00, 0x00,
	0xff, 0xff, 0x17, 0xcf, 0xeb, 0xa1, 0xba, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BookShelfClient is the client API for BookShelf service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookShelfClient interface {
	List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Books, error)
	Add(ctx context.Context, in *Book, opts ...grpc.CallOption) (*BookState, error)
	Remove(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Void, error)
	Update(ctx context.Context, in *BookState, opts ...grpc.CallOption) (*BookState, error)
}

type bookShelfClient struct {
	cc *grpc.ClientConn
}

func NewBookShelfClient(cc *grpc.ClientConn) BookShelfClient {
	return &bookShelfClient{cc}
}

func (c *bookShelfClient) List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Books, error) {
	out := new(Books)
	err := c.cc.Invoke(ctx, "/gbookshelf.BookShelf/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookShelfClient) Add(ctx context.Context, in *Book, opts ...grpc.CallOption) (*BookState, error) {
	out := new(BookState)
	err := c.cc.Invoke(ctx, "/gbookshelf.BookShelf/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookShelfClient) Remove(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/gbookshelf.BookShelf/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookShelfClient) Update(ctx context.Context, in *BookState, opts ...grpc.CallOption) (*BookState, error) {
	out := new(BookState)
	err := c.cc.Invoke(ctx, "/gbookshelf.BookShelf/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookShelfServer is the server API for BookShelf service.
type BookShelfServer interface {
	List(context.Context, *Void) (*Books, error)
	Add(context.Context, *Book) (*BookState, error)
	Remove(context.Context, *Book) (*Void, error)
	Update(context.Context, *BookState) (*BookState, error)
}

func RegisterBookShelfServer(s *grpc.Server, srv BookShelfServer) {
	s.RegisterService(&_BookShelf_serviceDesc, srv)
}

func _BookShelf_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookShelfServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gbookshelf.BookShelf/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookShelfServer).List(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookShelf_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookShelfServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gbookshelf.BookShelf/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookShelfServer).Add(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookShelf_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Book)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookShelfServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gbookshelf.BookShelf/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookShelfServer).Remove(ctx, req.(*Book))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookShelf_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookState)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookShelfServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gbookshelf.BookShelf/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookShelfServer).Update(ctx, req.(*BookState))
	}
	return interceptor(ctx, in, info, handler)
}

var _BookShelf_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gbookshelf.BookShelf",
	HandlerType: (*BookShelfServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _BookShelf_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _BookShelf_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _BookShelf_Remove_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BookShelf_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gbookshelf.proto",
}
