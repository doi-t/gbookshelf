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
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x4f, 0xca, 0xcf,
	0xcf, 0x2e, 0xce, 0x48, 0xcd, 0x49, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x79, 0x72, 0x71, 0x3a, 0xe5, 0xe7, 0x67, 0x07, 0x97, 0x24, 0x96, 0xa4, 0x0a, 0x89, 0x70,
	0xb1, 0x96, 0x64, 0x96, 0xe4, 0xa4, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x42,
	0x42, 0x5c, 0x2c, 0x05, 0x89, 0xe9, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x60, 0x36,
	0x48, 0x2c, 0x25, 0x3f, 0x2f, 0x55, 0x82, 0x59, 0x81, 0x51, 0x83, 0x23, 0x08, 0xcc, 0x56, 0x32,
	0xe1, 0x62, 0x05, 0x19, 0x55, 0x2c, 0xa4, 0xcd, 0xc5, 0x0a, 0xb6, 0x40, 0x82, 0x51, 0x81, 0x59,
	0x83, 0xdb, 0x48, 0x54, 0x0f, 0xc9, 0x05, 0x70, 0xcb, 0x82, 0x20, 0x6a, 0x94, 0x0c, 0xb8, 0x58,
	0x40, 0x62, 0xc4, 0xdb, 0xad, 0xc4, 0xc6, 0xc5, 0x12, 0x96, 0x9f, 0x99, 0x62, 0x34, 0x87, 0x11,
	0xea, 0x76, 0x90, 0xc1, 0x42, 0xba, 0x5c, 0x2c, 0x3e, 0x99, 0xc5, 0x25, 0x42, 0x02, 0xc8, 0xb6,
	0x81, 0xd4, 0x49, 0x09, 0xa2, 0xdb, 0x5f, 0xac, 0xc4, 0x20, 0x64, 0xc0, 0xc5, 0xec, 0x98, 0x92,
	0x82, 0xaa, 0x1a, 0x24, 0x27, 0x85, 0xdd, 0xb5, 0x4a, 0x0c, 0x42, 0x7a, 0x5c, 0x6c, 0x41, 0xa9,
	0xb9, 0xf9, 0x65, 0xa9, 0x58, 0x34, 0x61, 0x58, 0xaa, 0xc4, 0x90, 0xc4, 0x06, 0x0e, 0x6c, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x15, 0x23, 0xd2, 0x80, 0x01, 0x00, 0x00,
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

// BookShelfServer is the server API for BookShelf service.
type BookShelfServer interface {
	List(context.Context, *Void) (*Books, error)
	Add(context.Context, *Book) (*BookState, error)
	Remove(context.Context, *Book) (*Void, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gbookshelf.proto",
}
