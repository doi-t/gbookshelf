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

type Book struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Page                 int32    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	Done                 bool     `protobuf:"varint,3,opt,name=done,proto3" json:"done,omitempty"`
	Current              int32    `protobuf:"varint,4,opt,name=current,proto3" json:"current,omitempty"`
	Id                   int32    `protobuf:"varint,5,opt,name=id,proto3" json:"id,omitempty"`
	BookNumber           int32    `protobuf:"varint,6,opt,name=book_number,json=bookNumber,proto3" json:"book_number,omitempty"`
	Why                  string   `protobuf:"bytes,7,opt,name=why,proto3" json:"why,omitempty"`
	What                 string   `protobuf:"bytes,8,opt,name=what,proto3" json:"what,omitempty"`
	Review               string   `protobuf:"bytes,9,opt,name=review,proto3" json:"review,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ba0627d911171ab, []int{0}
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

func (m *Book) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *Book) GetCurrent() int32 {
	if m != nil {
		return m.Current
	}
	return 0
}

func (m *Book) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Book) GetBookNumber() int32 {
	if m != nil {
		return m.BookNumber
	}
	return 0
}

func (m *Book) GetWhy() string {
	if m != nil {
		return m.Why
	}
	return ""
}

func (m *Book) GetWhat() string {
	if m != nil {
		return m.What
	}
	return ""
}

func (m *Book) GetReview() string {
	if m != nil {
		return m.Review
	}
	return ""
}

type Books struct {
	Books                []*Book  `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
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

func (m *Books) GetBooks() []*Book {
	if m != nil {
		return m.Books
	}
	return nil
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
	return fileDescriptor_7ba0627d911171ab, []int{2}
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
	proto.RegisterType((*Book)(nil), "gbookshelf.Book")
	proto.RegisterType((*Books)(nil), "gbookshelf.Books")
	proto.RegisterType((*Void)(nil), "gbookshelf.Void")
}

func init() { proto.RegisterFile("gbookshelf.proto", fileDescriptor_7ba0627d911171ab) }

var fileDescriptor_7ba0627d911171ab = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0xbb, 0xcd, 0x47, 0xdb, 0x29, 0x48, 0x1c, 0x44, 0x86, 0x5e, 0x0c, 0x39, 0x48, 0x40,
	0x8c, 0x50, 0x9f, 0x40, 0xcf, 0xe2, 0x21, 0xa2, 0x57, 0x49, 0xdc, 0xb5, 0x59, 0xda, 0x66, 0x43,
	0xb2, 0x6d, 0xf0, 0x09, 0x3d, 0xf9, 0x4e, 0xb2, 0x53, 0xc5, 0x8f, 0x5e, 0x7a, 0xfb, 0xcf, 0x6f,
	0x7e, 0xc9, 0xec, 0xec, 0x42, 0xb4, 0x28, 0x8d, 0x59, 0x76, 0x95, 0x5a, 0xbd, 0x66, 0x4d, 0x6b,
	0xac, 0x41, 0xf8, 0x21, 0xc9, 0x87, 0x00, 0xff, 0xd6, 0x98, 0x25, 0x9e, 0x40, 0x60, 0xb5, 0x5d,
	0x29, 0x12, 0xb1, 0x48, 0x27, 0xf9, 0xae, 0x40, 0x04, 0xbf, 0x29, 0x16, 0x8a, 0x86, 0xb1, 0x48,
	0x83, 0x9c, 0xb3, 0x63, 0xd2, 0xd4, 0x8a, 0xbc, 0x58, 0xa4, 0xe3, 0x9c, 0x33, 0x12, 0x8c, 0x5e,
	0x36, 0x6d, 0xab, 0x6a, 0x4b, 0x3e, 0xab, 0xdf, 0x25, 0x1e, 0xc1, 0x50, 0x4b, 0x0a, 0x18, 0x0e,
	0xb5, 0xc4, 0x33, 0x98, 0xba, 0xe9, 0xcf, 0xf5, 0x66, 0x5d, 0xaa, 0x96, 0x42, 0x6e, 0x80, 0x43,
	0xf7, 0x4c, 0x30, 0x02, 0xaf, 0xaf, 0xde, 0x68, 0xc4, 0xc7, 0x70, 0xd1, 0x0d, 0xec, 0xab, 0xc2,
	0xd2, 0x98, 0x11, 0x67, 0x3c, 0x85, 0xb0, 0x55, 0x5b, 0xad, 0x7a, 0x9a, 0x30, 0xfd, 0xaa, 0x92,
	0x2b, 0x08, 0xdc, 0x3a, 0x1d, 0x9e, 0x43, 0xc0, 0x5b, 0x92, 0x88, 0xbd, 0x74, 0x3a, 0x8f, 0xb2,
	0x5f, 0xd7, 0xe0, 0x8c, 0x7c, 0xd7, 0x4e, 0x42, 0xf0, 0x9f, 0x8c, 0x96, 0xf3, 0x77, 0x01, 0x13,
	0xc7, 0x1f, 0x9c, 0x81, 0x97, 0xe0, 0xdf, 0xe9, 0xce, 0xe2, 0x9f, 0xcf, 0x9c, 0x37, 0x3b, 0xfe,
	0xff, 0xa3, 0x2e, 0x19, 0xe0, 0x05, 0x78, 0x37, 0x52, 0xe2, 0xde, 0x90, 0xd9, 0x1e, 0x49, 0x06,
	0x98, 0x41, 0x98, 0xab, 0xb5, 0xd9, 0xaa, 0xc3, 0xfd, 0xc7, 0x46, 0x16, 0xf6, 0x40, 0xbf, 0x0c,
	0xf9, 0x95, 0xaf, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xc2, 0xd5, 0x54, 0xf9, 0x01, 0x00,
	0x00,
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
	Add(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Book, error)
	Remove(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Book, error)
	Update(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Book, error)
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

func (c *bookShelfClient) Add(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/gbookshelf.BookShelf/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookShelfClient) Remove(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/gbookshelf.BookShelf/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookShelfClient) Update(ctx context.Context, in *Book, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/gbookshelf.BookShelf/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookShelfServer is the server API for BookShelf service.
type BookShelfServer interface {
	List(context.Context, *Void) (*Books, error)
	Add(context.Context, *Book) (*Book, error)
	Remove(context.Context, *Book) (*Book, error)
	Update(context.Context, *Book) (*Book, error)
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
	in := new(Book)
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
		return srv.(BookShelfServer).Update(ctx, req.(*Book))
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
