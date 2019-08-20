// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: network.proto

package types

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type NetworkGetRequest struct {
	// Types that are valid to be assigned to Options:
	//	*NetworkGetRequest_Address
	//	*NetworkGetRequest_Name
	Options isNetworkGetRequest_Options `protobuf_oneof:"Options"`
}

func (m *NetworkGetRequest) Reset()         { *m = NetworkGetRequest{} }
func (m *NetworkGetRequest) String() string { return proto.CompactTextString(m) }
func (*NetworkGetRequest) ProtoMessage()    {}
func (*NetworkGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{0}
}
func (m *NetworkGetRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkGetRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkGetRequest.Merge(m, src)
}
func (m *NetworkGetRequest) XXX_Size() int {
	return m.Size()
}
func (m *NetworkGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkGetRequest proto.InternalMessageInfo

type isNetworkGetRequest_Options interface {
	isNetworkGetRequest_Options()
	MarshalTo([]byte) (int, error)
	Size() int
}

type NetworkGetRequest_Address struct {
	Address string `protobuf:"bytes,1,opt,name=Address,proto3,oneof"`
}
type NetworkGetRequest_Name struct {
	Name string `protobuf:"bytes,2,opt,name=Name,proto3,oneof"`
}

func (*NetworkGetRequest_Address) isNetworkGetRequest_Options() {}
func (*NetworkGetRequest_Name) isNetworkGetRequest_Options()    {}

func (m *NetworkGetRequest) GetOptions() isNetworkGetRequest_Options {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *NetworkGetRequest) GetAddress() string {
	if x, ok := m.GetOptions().(*NetworkGetRequest_Address); ok {
		return x.Address
	}
	return ""
}

func (m *NetworkGetRequest) GetName() string {
	if x, ok := m.GetOptions().(*NetworkGetRequest_Name); ok {
		return x.Name
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*NetworkGetRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _NetworkGetRequest_OneofMarshaler, _NetworkGetRequest_OneofUnmarshaler, _NetworkGetRequest_OneofSizer, []interface{}{
		(*NetworkGetRequest_Address)(nil),
		(*NetworkGetRequest_Name)(nil),
	}
}

func _NetworkGetRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*NetworkGetRequest)
	// Options
	switch x := m.Options.(type) {
	case *NetworkGetRequest_Address:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Address)
	case *NetworkGetRequest_Name:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		_ = b.EncodeStringBytes(x.Name)
	case nil:
	default:
		return fmt.Errorf("NetworkGetRequest.Options has unexpected type %T", x)
	}
	return nil
}

func _NetworkGetRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*NetworkGetRequest)
	switch tag {
	case 1: // Options.Address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Options = &NetworkGetRequest_Address{x}
		return true, err
	case 2: // Options.Name
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Options = &NetworkGetRequest_Name{x}
		return true, err
	default:
		return false, nil
	}
}

func _NetworkGetRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*NetworkGetRequest)
	// Options
	switch x := m.Options.(type) {
	case *NetworkGetRequest_Address:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Address)))
		n += len(x.Address)
	case *NetworkGetRequest_Name:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Name)))
		n += len(x.Name)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type NetworkGetResponse struct {
	Data *NetworkAdapter `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *NetworkGetResponse) Reset()         { *m = NetworkGetResponse{} }
func (m *NetworkGetResponse) String() string { return proto.CompactTextString(m) }
func (*NetworkGetResponse) ProtoMessage()    {}
func (*NetworkGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{1}
}
func (m *NetworkGetResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkGetResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkGetResponse.Merge(m, src)
}
func (m *NetworkGetResponse) XXX_Size() int {
	return m.Size()
}
func (m *NetworkGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkGetResponse proto.InternalMessageInfo

func (m *NetworkGetResponse) GetData() *NetworkAdapter {
	if m != nil {
		return m.Data
	}
	return nil
}

type NetworkAdapter struct {
	InterfaceIndex string `protobuf:"bytes,1,opt,name=InterfaceIndex,proto3" json:"InterfaceIndex,omitempty"`
	GatewayAddress string `protobuf:"bytes,2,opt,name=GatewayAddress,proto3" json:"GatewayAddress,omitempty"`
	SubnetCIDR     string `protobuf:"bytes,3,opt,name=SubnetCIDR,proto3" json:"SubnetCIDR,omitempty"`
	HostName       string `protobuf:"bytes,4,opt,name=HostName,proto3" json:"HostName,omitempty"`
	AddressCIDR    string `protobuf:"bytes,5,opt,name=AddressCIDR,proto3" json:"AddressCIDR,omitempty"`
}

func (m *NetworkAdapter) Reset()         { *m = NetworkAdapter{} }
func (m *NetworkAdapter) String() string { return proto.CompactTextString(m) }
func (*NetworkAdapter) ProtoMessage()    {}
func (*NetworkAdapter) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{2}
}
func (m *NetworkAdapter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NetworkAdapter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NetworkAdapter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NetworkAdapter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkAdapter.Merge(m, src)
}
func (m *NetworkAdapter) XXX_Size() int {
	return m.Size()
}
func (m *NetworkAdapter) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkAdapter.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkAdapter proto.InternalMessageInfo

func (m *NetworkAdapter) GetInterfaceIndex() string {
	if m != nil {
		return m.InterfaceIndex
	}
	return ""
}

func (m *NetworkAdapter) GetGatewayAddress() string {
	if m != nil {
		return m.GatewayAddress
	}
	return ""
}

func (m *NetworkAdapter) GetSubnetCIDR() string {
	if m != nil {
		return m.SubnetCIDR
	}
	return ""
}

func (m *NetworkAdapter) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *NetworkAdapter) GetAddressCIDR() string {
	if m != nil {
		return m.AddressCIDR
	}
	return ""
}

func init() {
	proto.RegisterType((*NetworkGetRequest)(nil), "wins.NetworkGetRequest")
	proto.RegisterType((*NetworkGetResponse)(nil), "wins.NetworkGetResponse")
	proto.RegisterType((*NetworkAdapter)(nil), "wins.NetworkAdapter")
}

func init() { proto.RegisterFile("network.proto", fileDescriptor_8571034d60397816) }

var fileDescriptor_8571034d60397816 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x31, 0x4f, 0x02, 0x41,
	0x10, 0x85, 0x6f, 0xe1, 0x10, 0x19, 0x22, 0x89, 0x1b, 0x12, 0x2f, 0x14, 0x2b, 0xb9, 0xc2, 0x50,
	0x51, 0x60, 0x67, 0x61, 0x02, 0x92, 0x00, 0x89, 0x41, 0x73, 0x74, 0x76, 0x0b, 0x37, 0x26, 0xc4,
	0xb8, 0x7b, 0xee, 0x0e, 0x22, 0xff, 0xc2, 0x5f, 0x64, 0x6d, 0x49, 0x69, 0x69, 0xe0, 0x8f, 0x18,
	0xf6, 0x00, 0x39, 0x2d, 0xe7, 0x7b, 0xef, 0xde, 0xbd, 0x99, 0x85, 0x13, 0x85, 0x34, 0xd7, 0xe6,
	0xa9, 0x99, 0x18, 0x4d, 0x9a, 0xfb, 0xf3, 0xa9, 0xb2, 0xe1, 0x3d, 0x9c, 0x0e, 0x53, 0xdc, 0x43,
	0x8a, 0xf0, 0x65, 0x86, 0x96, 0x78, 0x0d, 0x8a, 0xed, 0x38, 0x36, 0x68, 0x6d, 0xc0, 0xea, 0xac,
	0x51, 0xea, 0x7b, 0xd1, 0x0e, 0xf0, 0x2a, 0xf8, 0x43, 0xf9, 0x8c, 0x41, 0x6e, 0x2b, 0xb8, 0xa9,
	0x53, 0x82, 0xe2, 0x5d, 0x42, 0x53, 0xad, 0x6c, 0x78, 0x0d, 0xfc, 0x30, 0xd1, 0x26, 0x5a, 0x59,
	0xe4, 0x0d, 0xf0, 0xbb, 0x92, 0xa4, 0xcb, 0x2b, 0xb7, 0xaa, 0xcd, 0xcd, 0xcf, 0x9b, 0x5b, 0x5f,
	0x3b, 0x96, 0x09, 0xa1, 0x89, 0x9c, 0x23, 0xfc, 0x60, 0x50, 0xc9, 0x0a, 0xfc, 0x02, 0x2a, 0x03,
	0x45, 0x68, 0x1e, 0xe5, 0x04, 0x07, 0x2a, 0xc6, 0xb7, 0xb4, 0x56, 0xf4, 0x87, 0x6e, 0x7c, 0x3d,
	0x49, 0x38, 0x97, 0x8b, 0x5d, 0xfd, 0x5c, 0xea, 0xcb, 0x52, 0x2e, 0x00, 0x46, 0xb3, 0xb1, 0x42,
	0xba, 0x19, 0x74, 0xa3, 0x20, 0xef, 0x3c, 0x07, 0x84, 0xd7, 0xe0, 0xb8, 0xaf, 0x2d, 0xb9, 0x3d,
	0x7d, 0xa7, 0xee, 0x67, 0x5e, 0x87, 0xf2, 0x36, 0xc6, 0x7d, 0x5c, 0x70, 0xf2, 0x21, 0x6a, 0xdd,
	0xee, 0xfb, 0x8f, 0xd0, 0xbc, 0x4e, 0x27, 0xc8, 0xaf, 0x20, 0xdf, 0x43, 0xe2, 0x67, 0x99, 0xad,
	0x7f, 0xef, 0x5d, 0x0b, 0xfe, 0x0b, 0xe9, 0xd9, 0x42, 0xaf, 0x73, 0xfe, 0xb9, 0x12, 0x6c, 0xb9,
	0x12, 0xec, 0x7b, 0x25, 0xd8, 0xfb, 0x5a, 0x78, 0xcb, 0xb5, 0xf0, 0xbe, 0xd6, 0xc2, 0x7b, 0x28,
	0xd0, 0x22, 0x41, 0x3b, 0x3e, 0x72, 0xcf, 0x79, 0xf9, 0x13, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x85,
	0x78, 0xf3, 0xdf, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkServiceClient is the client API for NetworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkServiceClient interface {
	Get(ctx context.Context, in *NetworkGetRequest, opts ...grpc.CallOption) (*NetworkGetResponse, error)
}

type networkServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServiceClient(cc *grpc.ClientConn) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) Get(ctx context.Context, in *NetworkGetRequest, opts ...grpc.CallOption) (*NetworkGetResponse, error) {
	out := new(NetworkGetResponse)
	err := c.cc.Invoke(ctx, "/wins.NetworkService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServiceServer is the server API for NetworkService service.
type NetworkServiceServer interface {
	Get(context.Context, *NetworkGetRequest) (*NetworkGetResponse, error)
}

func RegisterNetworkServiceServer(s *grpc.Server, srv NetworkServiceServer) {
	s.RegisterService(&_NetworkService_serviceDesc, srv)
}

func _NetworkService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wins.NetworkService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Get(ctx, req.(*NetworkGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wins.NetworkService",
	HandlerType: (*NetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _NetworkService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "network.proto",
}

func (m *NetworkGetRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkGetRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Options != nil {
		nn1, err := m.Options.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += nn1
	}
	return i, nil
}

func (m *NetworkGetRequest_Address) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0xa
	i++
	i = encodeVarintNetwork(dAtA, i, uint64(len(m.Address)))
	i += copy(dAtA[i:], m.Address)
	return i, nil
}
func (m *NetworkGetRequest_Name) MarshalTo(dAtA []byte) (int, error) {
	i := 0
	dAtA[i] = 0x12
	i++
	i = encodeVarintNetwork(dAtA, i, uint64(len(m.Name)))
	i += copy(dAtA[i:], m.Name)
	return i, nil
}
func (m *NetworkGetResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkGetResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNetwork(dAtA, i, uint64(m.Data.Size()))
		n2, err := m.Data.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *NetworkAdapter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkAdapter) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.InterfaceIndex) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNetwork(dAtA, i, uint64(len(m.InterfaceIndex)))
		i += copy(dAtA[i:], m.InterfaceIndex)
	}
	if len(m.GatewayAddress) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNetwork(dAtA, i, uint64(len(m.GatewayAddress)))
		i += copy(dAtA[i:], m.GatewayAddress)
	}
	if len(m.SubnetCIDR) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintNetwork(dAtA, i, uint64(len(m.SubnetCIDR)))
		i += copy(dAtA[i:], m.SubnetCIDR)
	}
	if len(m.HostName) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintNetwork(dAtA, i, uint64(len(m.HostName)))
		i += copy(dAtA[i:], m.HostName)
	}
	if len(m.AddressCIDR) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintNetwork(dAtA, i, uint64(len(m.AddressCIDR)))
		i += copy(dAtA[i:], m.AddressCIDR)
	}
	return i, nil
}

func encodeVarintNetwork(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *NetworkGetRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Options != nil {
		n += m.Options.Size()
	}
	return n
}

func (m *NetworkGetRequest_Address) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	n += 1 + l + sovNetwork(uint64(l))
	return n
}
func (m *NetworkGetRequest_Name) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	n += 1 + l + sovNetwork(uint64(l))
	return n
}
func (m *NetworkGetResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovNetwork(uint64(l))
	}
	return n
}

func (m *NetworkAdapter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.InterfaceIndex)
	if l > 0 {
		n += 1 + l + sovNetwork(uint64(l))
	}
	l = len(m.GatewayAddress)
	if l > 0 {
		n += 1 + l + sovNetwork(uint64(l))
	}
	l = len(m.SubnetCIDR)
	if l > 0 {
		n += 1 + l + sovNetwork(uint64(l))
	}
	l = len(m.HostName)
	if l > 0 {
		n += 1 + l + sovNetwork(uint64(l))
	}
	l = len(m.AddressCIDR)
	if l > 0 {
		n += 1 + l + sovNetwork(uint64(l))
	}
	return n
}

func sovNetwork(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNetwork(x uint64) (n int) {
	return sovNetwork(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NetworkGetRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetwork
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NetworkGetRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkGetRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetwork
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetwork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Options = &NetworkGetRequest_Address{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetwork
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetwork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Options = &NetworkGetRequest_Name{string(dAtA[iNdEx:postIndex])}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetwork(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetwork
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetwork
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NetworkGetResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetwork
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NetworkGetResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkGetResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNetwork
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNetwork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &NetworkAdapter{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetwork(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetwork
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetwork
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NetworkAdapter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetwork
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NetworkAdapter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkAdapter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterfaceIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetwork
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetwork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InterfaceIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GatewayAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetwork
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetwork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GatewayAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubnetCIDR", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetwork
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetwork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubnetCIDR = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetwork
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetwork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressCIDR", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNetwork
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNetwork
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddressCIDR = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetwork(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetwork
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNetwork
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNetwork(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNetwork
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNetwork
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthNetwork
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthNetwork
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNetwork
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipNetwork(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthNetwork
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthNetwork = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNetwork   = fmt.Errorf("proto: integer overflow")
)