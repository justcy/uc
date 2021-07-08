// Code generated by protoc-gen-go. DO NOT EDIT.
// source: myproto.proto

//指定所在包名

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

//定义枚举
type Week int32

const (
	Week_Monday  Week = 0
	Week_Tuesday Week = 1
)

var Week_name = map[int32]string{
	0: "Monday",
	1: "Tuesday",
}

var Week_value = map[string]int32{
	"Monday":  0,
	"Tuesday": 1,
}

func (x Week) String() string {
	return proto.EnumName(Week_name, int32(x))
}

func (Week) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_04877df457807402, []int{0}
}

type Student struct {
	Age  int32   `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	Name string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	P    *People `protobuf:"bytes,3,opt,name=p,proto3" json:"p,omitempty"`
	//数组
	Score []int32 `protobuf:"varint,4,rep,packed,name=score,proto3" json:"score,omitempty"`
	//枚举
	W Week `protobuf:"varint,5,opt,name=w,proto3,enum=pb.Week" json:"w,omitempty"`
	//联合体
	//
	// Types that are valid to be assigned to Data:
	//	*Student_Teacher
	//	*Student_Class
	Data                 isStudent_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Student) Reset()         { *m = Student{} }
func (m *Student) String() string { return proto.CompactTextString(m) }
func (*Student) ProtoMessage()    {}
func (*Student) Descriptor() ([]byte, []int) {
	return fileDescriptor_04877df457807402, []int{0}
}

func (m *Student) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Student.Unmarshal(m, b)
}
func (m *Student) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Student.Marshal(b, m, deterministic)
}
func (m *Student) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Student.Merge(m, src)
}
func (m *Student) XXX_Size() int {
	return xxx_messageInfo_Student.Size(m)
}
func (m *Student) XXX_DiscardUnknown() {
	xxx_messageInfo_Student.DiscardUnknown(m)
}

var xxx_messageInfo_Student proto.InternalMessageInfo

func (m *Student) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Student) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Student) GetP() *People {
	if m != nil {
		return m.P
	}
	return nil
}

func (m *Student) GetScore() []int32 {
	if m != nil {
		return m.Score
	}
	return nil
}

func (m *Student) GetW() Week {
	if m != nil {
		return m.W
	}
	return Week_Monday
}

type isStudent_Data interface {
	isStudent_Data()
}

type Student_Teacher struct {
	Teacher string `protobuf:"bytes,6,opt,name=teacher,proto3,oneof"`
}

type Student_Class struct {
	Class string `protobuf:"bytes,7,opt,name=class,proto3,oneof"`
}

func (*Student_Teacher) isStudent_Data() {}

func (*Student_Class) isStudent_Data() {}

func (m *Student) GetData() isStudent_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Student) GetTeacher() string {
	if x, ok := m.GetData().(*Student_Teacher); ok {
		return x.Teacher
	}
	return ""
}

func (m *Student) GetClass() string {
	if x, ok := m.GetData().(*Student_Class); ok {
		return x.Class
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Student) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Student_Teacher)(nil),
		(*Student_Class)(nil),
	}
}

type People struct {
	Weight               int32    `protobuf:"varint,1,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *People) Reset()         { *m = People{} }
func (m *People) String() string { return proto.CompactTextString(m) }
func (*People) ProtoMessage()    {}
func (*People) Descriptor() ([]byte, []int) {
	return fileDescriptor_04877df457807402, []int{1}
}

func (m *People) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_People.Unmarshal(m, b)
}
func (m *People) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_People.Marshal(b, m, deterministic)
}
func (m *People) XXX_Merge(src proto.Message) {
	xxx_messageInfo_People.Merge(m, src)
}
func (m *People) XXX_Size() int {
	return xxx_messageInfo_People.Size(m)
}
func (m *People) XXX_DiscardUnknown() {
	xxx_messageInfo_People.DiscardUnknown(m)
}

var xxx_messageInfo_People proto.InternalMessageInfo

func (m *People) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func init() {
	proto.RegisterEnum("pb.Week", Week_name, Week_value)
	proto.RegisterType((*Student)(nil), "pb.Student")
	proto.RegisterType((*People)(nil), "pb.People")
}

func init() { proto.RegisterFile("myproto.proto", fileDescriptor_04877df457807402) }

var fileDescriptor_04877df457807402 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0x3b, 0x4d, 0xb2, 0xf9, 0xff, 0xa7, 0x28, 0x65, 0x90, 0xb0, 0xf4, 0xa0, 0x4b, 0x0f,
	0x12, 0x3c, 0xe4, 0x90, 0x5e, 0x3c, 0x7b, 0xf2, 0x22, 0xc8, 0x56, 0xf0, 0xbc, 0x49, 0x86, 0x56,
	0x6d, 0x93, 0x25, 0xd9, 0x12, 0xf2, 0xb9, 0xfc, 0x82, 0xb2, 0x49, 0x04, 0x2f, 0xcb, 0xfb, 0xed,
	0x0c, 0xf3, 0x66, 0x1e, 0x5e, 0x9d, 0x07, 0xdb, 0x36, 0xae, 0xc9, 0xc6, 0x97, 0x96, 0xb6, 0xd8,
	0x7e, 0x03, 0xc6, 0x7b, 0x77, 0xa9, 0xb8, 0x76, 0xb4, 0xc6, 0xc0, 0x1c, 0x58, 0x82, 0x82, 0x34,
	0xd2, 0x5e, 0x12, 0x61, 0x58, 0x9b, 0x33, 0xcb, 0xa5, 0x82, 0xf4, 0xbf, 0x1e, 0x35, 0x49, 0x04,
	0x2b, 0x03, 0x05, 0xe9, 0x2a, 0xc7, 0xcc, 0x16, 0xd9, 0x2b, 0x37, 0xf6, 0xc4, 0x1a, 0x2c, 0xdd,
	0x60, 0xd4, 0x95, 0x4d, 0xcb, 0x32, 0x54, 0x41, 0x1a, 0xe9, 0x09, 0x28, 0x41, 0xe8, 0x65, 0xa4,
	0x20, 0xbd, 0xce, 0xff, 0xf9, 0xfe, 0x77, 0xe6, 0x2f, 0x0d, 0x3d, 0x6d, 0x30, 0x76, 0x6c, 0xca,
	0x23, 0xb7, 0x52, 0xf8, 0xf1, 0xcf, 0x0b, 0xfd, 0xfb, 0x41, 0x09, 0x46, 0xe5, 0xc9, 0x74, 0x9d,
	0x8c, 0xe7, 0xca, 0x84, 0x4f, 0x02, 0xc3, 0xca, 0x38, 0xb3, 0x55, 0x28, 0x26, 0x5b, 0x4a, 0x50,
	0xf4, 0xfc, 0x71, 0x38, 0xba, 0x79, 0xed, 0x99, 0x1e, 0xee, 0x30, 0xf4, 0x46, 0x84, 0x28, 0x5e,
	0x9a, 0xba, 0x32, 0xc3, 0x7a, 0x41, 0x2b, 0x8c, 0xdf, 0x2e, 0xdc, 0x79, 0x80, 0xfc, 0x1e, 0xc3,
	0xe2, 0x73, 0xf7, 0x48, 0xb7, 0x18, 0xec, 0xcd, 0x40, 0x7f, 0x4e, 0xd9, 0xac, 0xbc, 0x9e, 0x43,
	0x29, 0xc4, 0x98, 0xd5, 0xee, 0x27, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x72, 0x39, 0x5f, 0x3c, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Bj38Client is the client API for Bj38 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Bj38Client interface {
	Say(ctx context.Context, in *People, opts ...grpc.CallOption) (*Student, error)
}

type bj38Client struct {
	cc *grpc.ClientConn
}

func NewBj38Client(cc *grpc.ClientConn) Bj38Client {
	return &bj38Client{cc}
}

func (c *bj38Client) Say(ctx context.Context, in *People, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/pb.bj38/Say", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Bj38Server is the server API for Bj38 service.
type Bj38Server interface {
	Say(context.Context, *People) (*Student, error)
}

// UnimplementedBj38Server can be embedded to have forward compatible implementations.
type UnimplementedBj38Server struct {
}

func (*UnimplementedBj38Server) Say(ctx context.Context, req *People) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Say not implemented")
}

func RegisterBj38Server(s *grpc.Server, srv Bj38Server) {
	s.RegisterService(&_Bj38_serviceDesc, srv)
}

func _Bj38_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(People)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Bj38Server).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.bj38/Say",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Bj38Server).Say(ctx, req.(*People))
	}
	return interceptor(ctx, in, info, handler)
}

var _Bj38_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.bj38",
	HandlerType: (*Bj38Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _Bj38_Say_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "myproto.proto",
}