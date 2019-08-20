// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: host.proto

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

type HostGetVersionResponse struct {
	Data *HostVersion `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *HostGetVersionResponse) Reset()         { *m = HostGetVersionResponse{} }
func (m *HostGetVersionResponse) String() string { return proto.CompactTextString(m) }
func (*HostGetVersionResponse) ProtoMessage()    {}
func (*HostGetVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85e40b83b4d50a8d, []int{0}
}
func (m *HostGetVersionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HostGetVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HostGetVersionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HostGetVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostGetVersionResponse.Merge(m, src)
}
func (m *HostGetVersionResponse) XXX_Size() int {
	return m.Size()
}
func (m *HostGetVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HostGetVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HostGetVersionResponse proto.InternalMessageInfo

func (m *HostGetVersionResponse) GetData() *HostVersion {
	if m != nil {
		return m.Data
	}
	return nil
}

type HostVersion struct {
	CurrentMajorVersionNumber string `protobuf:"bytes,1,opt,name=CurrentMajorVersionNumber,proto3" json:"CurrentMajorVersionNumber,omitempty"`
	CurrentMinorVersionNumber string `protobuf:"bytes,2,opt,name=CurrentMinorVersionNumber,proto3" json:"CurrentMinorVersionNumber,omitempty"`
	CurrentBuildNumber        string `protobuf:"bytes,3,opt,name=CurrentBuildNumber,proto3" json:"CurrentBuildNumber,omitempty"`
	UBR                       string `protobuf:"bytes,4,opt,name=UBR,proto3" json:"UBR,omitempty"`
	ReleaseId                 string `protobuf:"bytes,5,opt,name=ReleaseId,proto3" json:"ReleaseId,omitempty"`
	BuildLabEx                string `protobuf:"bytes,6,opt,name=BuildLabEx,proto3" json:"BuildLabEx,omitempty"`
	CurrentBuild              string `protobuf:"bytes,7,opt,name=CurrentBuild,proto3" json:"CurrentBuild,omitempty"`
}

func (m *HostVersion) Reset()         { *m = HostVersion{} }
func (m *HostVersion) String() string { return proto.CompactTextString(m) }
func (*HostVersion) ProtoMessage()    {}
func (*HostVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_85e40b83b4d50a8d, []int{1}
}
func (m *HostVersion) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HostVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HostVersion.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *HostVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HostVersion.Merge(m, src)
}
func (m *HostVersion) XXX_Size() int {
	return m.Size()
}
func (m *HostVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_HostVersion.DiscardUnknown(m)
}

var xxx_messageInfo_HostVersion proto.InternalMessageInfo

func (m *HostVersion) GetCurrentMajorVersionNumber() string {
	if m != nil {
		return m.CurrentMajorVersionNumber
	}
	return ""
}

func (m *HostVersion) GetCurrentMinorVersionNumber() string {
	if m != nil {
		return m.CurrentMinorVersionNumber
	}
	return ""
}

func (m *HostVersion) GetCurrentBuildNumber() string {
	if m != nil {
		return m.CurrentBuildNumber
	}
	return ""
}

func (m *HostVersion) GetUBR() string {
	if m != nil {
		return m.UBR
	}
	return ""
}

func (m *HostVersion) GetReleaseId() string {
	if m != nil {
		return m.ReleaseId
	}
	return ""
}

func (m *HostVersion) GetBuildLabEx() string {
	if m != nil {
		return m.BuildLabEx
	}
	return ""
}

func (m *HostVersion) GetCurrentBuild() string {
	if m != nil {
		return m.CurrentBuild
	}
	return ""
}

func init() {
	proto.RegisterType((*HostGetVersionResponse)(nil), "wins.HostGetVersionResponse")
	proto.RegisterType((*HostVersion)(nil), "wins.HostVersion")
}

func init() { proto.RegisterFile("host.proto", fileDescriptor_85e40b83b4d50a8d) }

var fileDescriptor_85e40b83b4d50a8d = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x93, 0xfe, 0x49, 0x6f, 0xbb, 0xd0, 0x59, 0x48, 0x2c, 0x65, 0x94, 0x80, 0xe0, 0x2a,
	0x8b, 0xba, 0x71, 0x21, 0x08, 0x55, 0xa9, 0x82, 0xba, 0x88, 0xd8, 0x85, 0xbb, 0xb4, 0xb9, 0xe0,
	0x48, 0x33, 0x13, 0x66, 0x26, 0xfe, 0xbc, 0x85, 0x2f, 0xe2, 0x7b, 0xb8, 0xec, 0xd2, 0xa5, 0x24,
	0x2f, 0x22, 0x99, 0x44, 0x92, 0x5a, 0xdd, 0x85, 0xef, 0x9c, 0x2f, 0x03, 0xe7, 0x02, 0x3c, 0x08,
	0xa5, 0xbd, 0x58, 0x0a, 0x2d, 0x48, 0xeb, 0x99, 0x71, 0x35, 0xe8, 0xcf, 0x45, 0x14, 0x09, 0x5e,
	0x30, 0xf7, 0x04, 0xb6, 0x2f, 0x84, 0xd2, 0x13, 0xd4, 0x53, 0x94, 0x8a, 0x09, 0xee, 0xa3, 0x8a,
	0x05, 0x57, 0x48, 0xf6, 0xa1, 0x75, 0x16, 0xe8, 0xc0, 0xb1, 0xf7, 0xec, 0x83, 0xde, 0x68, 0xcb,
	0xcb, 0x65, 0x2f, 0xef, 0xfe, 0x14, 0x4d, 0xec, 0xbe, 0x37, 0xa0, 0x57, 0xa3, 0xe4, 0x18, 0x76,
	0x4e, 0x13, 0x29, 0x91, 0xeb, 0xeb, 0xe0, 0x51, 0xc8, 0x12, 0xdf, 0x24, 0xd1, 0x0c, 0xa5, 0xf9,
	0x57, 0xd7, 0xff, 0xbf, 0x50, 0xb7, 0x19, 0xff, 0x6d, 0x37, 0x56, 0xed, 0xb5, 0x02, 0xf1, 0x80,
	0x94, 0xe1, 0x38, 0x61, 0x8b, 0xb0, 0xd4, 0x9a, 0x46, 0xfb, 0x23, 0x21, 0x9b, 0xd0, 0xbc, 0x1b,
	0xfb, 0x4e, 0xcb, 0x14, 0xf2, 0x4f, 0x32, 0x84, 0xae, 0x8f, 0x0b, 0x0c, 0x14, 0x5e, 0x86, 0x4e,
	0xdb, 0xf0, 0x0a, 0x10, 0x0a, 0x60, 0xf4, 0xab, 0x60, 0x76, 0xfe, 0xe2, 0x74, 0x4c, 0x5c, 0x23,
	0xc4, 0x85, 0x7e, 0xfd, 0x15, 0x67, 0xc3, 0x34, 0x56, 0xd8, 0x68, 0x52, 0xcc, 0x75, 0x8b, 0xf2,
	0x89, 0xcd, 0x91, 0x1c, 0x01, 0x54, 0xdb, 0x13, 0x28, 0x56, 0x9e, 0x0a, 0x16, 0x0e, 0x86, 0xd5,
	0xe2, 0xeb, 0xd7, 0x71, 0xad, 0xf1, 0xee, 0x47, 0x4a, 0xed, 0x65, 0x4a, 0xed, 0xaf, 0x94, 0xda,
	0x6f, 0x19, 0xb5, 0x96, 0x19, 0xb5, 0x3e, 0x33, 0x6a, 0xdd, 0xb7, 0xf5, 0x6b, 0x8c, 0x6a, 0xd6,
	0x31, 0x17, 0x3e, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x33, 0xd0, 0x4e, 0x03, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HostServiceClient is the client API for HostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HostServiceClient interface {
	GetVersion(ctx context.Context, in *Void, opts ...grpc.CallOption) (*HostGetVersionResponse, error)
}

type hostServiceClient struct {
	cc *grpc.ClientConn
}

func NewHostServiceClient(cc *grpc.ClientConn) HostServiceClient {
	return &hostServiceClient{cc}
}

func (c *hostServiceClient) GetVersion(ctx context.Context, in *Void, opts ...grpc.CallOption) (*HostGetVersionResponse, error) {
	out := new(HostGetVersionResponse)
	err := c.cc.Invoke(ctx, "/wins.HostService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HostServiceServer is the server API for HostService service.
type HostServiceServer interface {
	GetVersion(context.Context, *Void) (*HostGetVersionResponse, error)
}

func RegisterHostServiceServer(s *grpc.Server, srv HostServiceServer) {
	s.RegisterService(&_HostService_serviceDesc, srv)
}

func _HostService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HostServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wins.HostService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HostServiceServer).GetVersion(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _HostService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wins.HostService",
	HandlerType: (*HostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _HostService_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "host.proto",
}

func (m *HostGetVersionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HostGetVersionResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHost(dAtA, i, uint64(m.Data.Size()))
		n1, err := m.Data.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *HostVersion) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HostVersion) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.CurrentMajorVersionNumber) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintHost(dAtA, i, uint64(len(m.CurrentMajorVersionNumber)))
		i += copy(dAtA[i:], m.CurrentMajorVersionNumber)
	}
	if len(m.CurrentMinorVersionNumber) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHost(dAtA, i, uint64(len(m.CurrentMinorVersionNumber)))
		i += copy(dAtA[i:], m.CurrentMinorVersionNumber)
	}
	if len(m.CurrentBuildNumber) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintHost(dAtA, i, uint64(len(m.CurrentBuildNumber)))
		i += copy(dAtA[i:], m.CurrentBuildNumber)
	}
	if len(m.UBR) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintHost(dAtA, i, uint64(len(m.UBR)))
		i += copy(dAtA[i:], m.UBR)
	}
	if len(m.ReleaseId) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintHost(dAtA, i, uint64(len(m.ReleaseId)))
		i += copy(dAtA[i:], m.ReleaseId)
	}
	if len(m.BuildLabEx) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintHost(dAtA, i, uint64(len(m.BuildLabEx)))
		i += copy(dAtA[i:], m.BuildLabEx)
	}
	if len(m.CurrentBuild) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintHost(dAtA, i, uint64(len(m.CurrentBuild)))
		i += copy(dAtA[i:], m.CurrentBuild)
	}
	return i, nil
}

func encodeVarintHost(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HostGetVersionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovHost(uint64(l))
	}
	return n
}

func (m *HostVersion) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CurrentMajorVersionNumber)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	l = len(m.CurrentMinorVersionNumber)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	l = len(m.CurrentBuildNumber)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	l = len(m.UBR)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	l = len(m.ReleaseId)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	l = len(m.BuildLabEx)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	l = len(m.CurrentBuild)
	if l > 0 {
		n += 1 + l + sovHost(uint64(l))
	}
	return n
}

func sovHost(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHost(x uint64) (n int) {
	return sovHost(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HostGetVersionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHost
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
			return fmt.Errorf("proto: HostGetVersionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HostGetVersionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &HostVersion{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHost
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHost
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
func (m *HostVersion) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHost
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
			return fmt.Errorf("proto: HostVersion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HostVersion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentMajorVersionNumber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrentMajorVersionNumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentMinorVersionNumber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrentMinorVersionNumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentBuildNumber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrentBuildNumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UBR", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UBR = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReleaseId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReleaseId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BuildLabEx", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BuildLabEx = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentBuild", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHost
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
				return ErrInvalidLengthHost
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthHost
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CurrentBuild = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHost(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHost
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthHost
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
func skipHost(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHost
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
					return 0, ErrIntOverflowHost
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
					return 0, ErrIntOverflowHost
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
				return 0, ErrInvalidLengthHost
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthHost
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHost
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
				next, err := skipHost(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthHost
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
	ErrInvalidLengthHost = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHost   = fmt.Errorf("proto: integer overflow")
)