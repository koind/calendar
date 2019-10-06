// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/transport/grpc/pb/event.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type EventRequest struct {
	Title                string               `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Datetime             *timestamp.Timestamp `protobuf:"bytes,2,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Duration             *duration.Duration   `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Description          string               `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	UserId               int32                `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TimeSendNotify       *timestamp.Timestamp `protobuf:"bytes,6,opt,name=time_send_notify,json=timeSendNotify,proto3" json:"time_send_notify,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EventRequest) Reset()         { *m = EventRequest{} }
func (m *EventRequest) String() string { return proto.CompactTextString(m) }
func (*EventRequest) ProtoMessage()    {}
func (*EventRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_feac913fc7642a85, []int{0}
}

func (m *EventRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventRequest.Unmarshal(m, b)
}
func (m *EventRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventRequest.Marshal(b, m, deterministic)
}
func (m *EventRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRequest.Merge(m, src)
}
func (m *EventRequest) XXX_Size() int {
	return xxx_messageInfo_EventRequest.Size(m)
}
func (m *EventRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventRequest proto.InternalMessageInfo

func (m *EventRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *EventRequest) GetDatetime() *timestamp.Timestamp {
	if m != nil {
		return m.Datetime
	}
	return nil
}

func (m *EventRequest) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *EventRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EventRequest) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *EventRequest) GetTimeSendNotify() *timestamp.Timestamp {
	if m != nil {
		return m.TimeSendNotify
	}
	return nil
}

type EventResponse struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Datetime             *timestamp.Timestamp `protobuf:"bytes,3,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Duration             *duration.Duration   `protobuf:"bytes,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Description          string               `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	UserId               int32                `protobuf:"varint,6,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TimeSendNotify       *timestamp.Timestamp `protobuf:"bytes,7,opt,name=time_send_notify,json=timeSendNotify,proto3" json:"time_send_notify,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EventResponse) Reset()         { *m = EventResponse{} }
func (m *EventResponse) String() string { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()    {}
func (*EventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_feac913fc7642a85, []int{1}
}

func (m *EventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventResponse.Unmarshal(m, b)
}
func (m *EventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventResponse.Marshal(b, m, deterministic)
}
func (m *EventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventResponse.Merge(m, src)
}
func (m *EventResponse) XXX_Size() int {
	return xxx_messageInfo_EventResponse.Size(m)
}
func (m *EventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventResponse proto.InternalMessageInfo

func (m *EventResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EventResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *EventResponse) GetDatetime() *timestamp.Timestamp {
	if m != nil {
		return m.Datetime
	}
	return nil
}

func (m *EventResponse) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *EventResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EventResponse) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *EventResponse) GetTimeSendNotify() *timestamp.Timestamp {
	if m != nil {
		return m.TimeSendNotify
	}
	return nil
}

type EventID struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventID) Reset()         { *m = EventID{} }
func (m *EventID) String() string { return proto.CompactTextString(m) }
func (*EventID) ProtoMessage()    {}
func (*EventID) Descriptor() ([]byte, []int) {
	return fileDescriptor_feac913fc7642a85, []int{2}
}

func (m *EventID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventID.Unmarshal(m, b)
}
func (m *EventID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventID.Marshal(b, m, deterministic)
}
func (m *EventID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventID.Merge(m, src)
}
func (m *EventID) XXX_Size() int {
	return xxx_messageInfo_EventID.Size(m)
}
func (m *EventID) XXX_DiscardUnknown() {
	xxx_messageInfo_EventID.DiscardUnknown(m)
}

var xxx_messageInfo_EventID proto.InternalMessageInfo

func (m *EventID) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type EventStatus struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventStatus) Reset()         { *m = EventStatus{} }
func (m *EventStatus) String() string { return proto.CompactTextString(m) }
func (*EventStatus) ProtoMessage()    {}
func (*EventStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_feac913fc7642a85, []int{3}
}

func (m *EventStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventStatus.Unmarshal(m, b)
}
func (m *EventStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventStatus.Marshal(b, m, deterministic)
}
func (m *EventStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventStatus.Merge(m, src)
}
func (m *EventStatus) XXX_Size() int {
	return xxx_messageInfo_EventStatus.Size(m)
}
func (m *EventStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_EventStatus.DiscardUnknown(m)
}

var xxx_messageInfo_EventStatus proto.InternalMessageInfo

func (m *EventStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type EventChange struct {
	EventID              *EventID `protobuf:"bytes,1,opt,name=eventID,proto3" json:"eventID,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventChange) Reset()         { *m = EventChange{} }
func (m *EventChange) String() string { return proto.CompactTextString(m) }
func (*EventChange) ProtoMessage()    {}
func (*EventChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_feac913fc7642a85, []int{4}
}

func (m *EventChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventChange.Unmarshal(m, b)
}
func (m *EventChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventChange.Marshal(b, m, deterministic)
}
func (m *EventChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventChange.Merge(m, src)
}
func (m *EventChange) XXX_Size() int {
	return xxx_messageInfo_EventChange.Size(m)
}
func (m *EventChange) XXX_DiscardUnknown() {
	xxx_messageInfo_EventChange.DiscardUnknown(m)
}

var xxx_messageInfo_EventChange proto.InternalMessageInfo

func (m *EventChange) GetEventID() *EventID {
	if m != nil {
		return m.EventID
	}
	return nil
}

func (m *EventChange) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func init() {
	proto.RegisterType((*EventRequest)(nil), "pb.EventRequest")
	proto.RegisterType((*EventResponse)(nil), "pb.EventResponse")
	proto.RegisterType((*EventID)(nil), "pb.EventID")
	proto.RegisterType((*EventStatus)(nil), "pb.EventStatus")
	proto.RegisterType((*EventChange)(nil), "pb.EventChange")
}

func init() { proto.RegisterFile("app/transport/grpc/pb/event.proto", fileDescriptor_feac913fc7642a85) }

var fileDescriptor_feac913fc7642a85 = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0x95, 0x74, 0x93, 0x2d, 0x13, 0x28, 0xc5, 0x42, 0x90, 0xe6, 0x00, 0xcb, 0x4a, 0x45,
	0x7b, 0x80, 0x44, 0x5a, 0x04, 0x0f, 0x40, 0x97, 0x43, 0x39, 0x70, 0x48, 0xe1, 0xc2, 0x65, 0xe5,
	0xd4, 0xd3, 0xc5, 0xd2, 0xd6, 0x36, 0xf6, 0x04, 0x89, 0x77, 0xe0, 0xc8, 0x2b, 0xf0, 0x9e, 0x28,
	0x8e, 0x37, 0x5d, 0x95, 0x42, 0xc5, 0xde, 0xe2, 0x99, 0xff, 0xff, 0x67, 0xf2, 0xd9, 0xf0, 0x8c,
	0x1b, 0x53, 0x91, 0xe5, 0xca, 0x19, 0x6d, 0xa9, 0x5a, 0x59, 0x73, 0x5e, 0x99, 0xa6, 0xc2, 0x6f,
	0xa8, 0xa8, 0x34, 0x56, 0x93, 0x66, 0xb1, 0x69, 0x8a, 0xa7, 0x2b, 0xad, 0x57, 0x6b, 0xac, 0x7c,
	0xa5, 0x69, 0x2f, 0x2a, 0x92, 0x97, 0xe8, 0x88, 0x5f, 0x9a, 0x5e, 0x54, 0x3c, 0xb9, 0x2e, 0x10,
	0xad, 0xe5, 0x24, 0xb5, 0xea, 0xfb, 0xd3, 0x9f, 0x31, 0xdc, 0x7d, 0xd7, 0x85, 0xd6, 0xf8, 0xb5,
	0x45, 0x47, 0xec, 0x21, 0x24, 0x24, 0x69, 0x8d, 0x79, 0x34, 0x89, 0x66, 0x77, 0xea, 0xfe, 0xc0,
	0xde, 0xc0, 0xbe, 0xe0, 0x84, 0x5d, 0x7a, 0x1e, 0x4f, 0xa2, 0x59, 0x36, 0x2f, 0xca, 0x3e, 0xb9,
	0xdc, 0x24, 0x97, 0x1f, 0x37, 0xa3, 0xeb, 0x41, 0xcb, 0x5e, 0xc3, 0xfe, 0x66, 0x60, 0xbe, 0xe7,
	0x7d, 0x47, 0x7f, 0xf8, 0x16, 0x41, 0x50, 0x0f, 0x52, 0x36, 0x81, 0x4c, 0xa0, 0x3b, 0xb7, 0xd2,
	0x78, 0xe7, 0xc8, 0xaf, 0xb2, 0x5d, 0x62, 0x8f, 0x61, 0xdc, 0x3a, 0xb4, 0x4b, 0x29, 0xf2, 0x64,
	0x12, 0xcd, 0x92, 0x3a, 0xed, 0x8e, 0xa7, 0x82, 0x2d, 0xe0, 0xb0, 0x9b, 0xbc, 0x74, 0xa8, 0xc4,
	0x52, 0x69, 0x92, 0x17, 0xdf, 0xf3, 0xf4, 0xd6, 0x8d, 0x0f, 0x3a, 0xcf, 0x19, 0x2a, 0xf1, 0xc1,
	0x3b, 0xa6, 0xbf, 0x62, 0xb8, 0x17, 0xb0, 0x38, 0xa3, 0x95, 0x43, 0x76, 0x00, 0xb1, 0x14, 0x1e,
	0x4a, 0x52, 0xc7, 0x52, 0x5c, 0x71, 0x8a, 0xff, 0xc6, 0x69, 0x6f, 0x47, 0x4e, 0xa3, 0x9d, 0x39,
	0x25, 0xff, 0xe4, 0x94, 0xde, 0xca, 0x69, 0xfc, 0xdf, 0x9c, 0x8e, 0x60, 0xec, 0x31, 0x9d, 0x2e,
	0xae, 0x03, 0x9a, 0x1e, 0x43, 0xe6, 0x5b, 0x67, 0xc4, 0xa9, 0x75, 0xec, 0x11, 0xa4, 0xce, 0x7f,
	0x85, 0x87, 0x15, 0x4e, 0xd3, 0xf7, 0x41, 0x76, 0xf2, 0x85, 0xab, 0x15, 0xb2, 0x63, 0x18, 0x63,
	0x1f, 0xe8, 0x75, 0xd9, 0x3c, 0x2b, 0x4d, 0x53, 0x86, 0x19, 0xf5, 0xa6, 0x77, 0x33, 0xfd, 0xf9,
	0x8f, 0x08, 0x12, 0x2f, 0x65, 0x2f, 0x21, 0x3d, 0xb1, 0xc8, 0x09, 0xd9, 0xe1, 0xe0, 0x0f, 0x2f,
	0xbc, 0x78, 0xb0, 0x55, 0x09, 0x97, 0xfb, 0x02, 0xd2, 0x4f, 0xa6, 0xbb, 0x0c, 0x76, 0x7f, 0x68,
	0xf6, 0x0b, 0xdd, 0xa4, 0x7e, 0x0e, 0xe9, 0x02, 0xd7, 0x48, 0xc8, 0xb6, 0x97, 0x2b, 0xae, 0xac,
	0xfd, 0x2f, 0xbf, 0x1d, 0x7d, 0x8e, 0x4d, 0xd3, 0xa4, 0x1e, 0xe3, 0xab, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x31, 0xae, 0x47, 0x7e, 0xd2, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventClient is the client API for Event service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventClient interface {
	// Создает новое событие
	Create(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
	// Обновляет событие
	Update(ctx context.Context, in *EventChange, opts ...grpc.CallOption) (*EventResponse, error)
	// Удаляет событие
	Delete(ctx context.Context, in *EventID, opts ...grpc.CallOption) (*EventStatus, error)
}

type eventClient struct {
	cc *grpc.ClientConn
}

func NewEventClient(cc *grpc.ClientConn) EventClient {
	return &eventClient{cc}
}

func (c *eventClient) Create(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/pb.Event/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) Update(ctx context.Context, in *EventChange, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/pb.Event/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventClient) Delete(ctx context.Context, in *EventID, opts ...grpc.CallOption) (*EventStatus, error) {
	out := new(EventStatus)
	err := c.cc.Invoke(ctx, "/pb.Event/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServer is the server API for Event service.
type EventServer interface {
	// Создает новое событие
	Create(context.Context, *EventRequest) (*EventResponse, error)
	// Обновляет событие
	Update(context.Context, *EventChange) (*EventResponse, error)
	// Удаляет событие
	Delete(context.Context, *EventID) (*EventStatus, error)
}

// UnimplementedEventServer can be embedded to have forward compatible implementations.
type UnimplementedEventServer struct {
}

func (*UnimplementedEventServer) Create(ctx context.Context, req *EventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedEventServer) Update(ctx context.Context, req *EventChange) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedEventServer) Delete(ctx context.Context, req *EventID) (*EventStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterEventServer(s *grpc.Server, srv EventServer) {
	s.RegisterService(&_Event_serviceDesc, srv)
}

func _Event_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Event/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).Create(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventChange)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Event/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).Update(ctx, req.(*EventChange))
	}
	return interceptor(ctx, in, info, handler)
}

func _Event_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Event/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServer).Delete(ctx, req.(*EventID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Event_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Event",
	HandlerType: (*EventServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Event_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Event_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Event_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/transport/grpc/pb/event.proto",
}