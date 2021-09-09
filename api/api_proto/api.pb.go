// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api.proto

package api

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserAuth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Role     string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
	Userid   string `protobuf:"bytes,4,opt,name=userid,proto3" json:"userid,omitempty"`
}

func (x *UserAuth) Reset() {
	*x = UserAuth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAuth) ProtoMessage() {}

func (x *UserAuth) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAuth.ProtoReflect.Descriptor instead.
func (*UserAuth) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *UserAuth) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserAuth) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserAuth) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *UserAuth) GetUserid() string {
	if x != nil {
		return x.Userid
	}
	return ""
}

type JWTResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Exp   string `protobuf:"bytes,2,opt,name=exp,proto3" json:"exp,omitempty"`
}

func (x *JWTResponse) Reset() {
	*x = JWTResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JWTResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JWTResponse) ProtoMessage() {}

func (x *JWTResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JWTResponse.ProtoReflect.Descriptor instead.
func (*JWTResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *JWTResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *JWTResponse) GetExp() string {
	if x != nil {
		return x.Exp
	}
	return ""
}

type Station struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Capacity     float32 `protobuf:"fixed32,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	UsedCapacity float32 `protobuf:"fixed32,3,opt,name=usedCapacity,proto3" json:"usedCapacity,omitempty"`
	Docks        []*Dock `protobuf:"bytes,4,rep,name=docks,proto3" json:"docks,omitempty"`
	IsRegistered bool    `protobuf:"varint,5,opt,name=isRegistered,proto3" json:"isRegistered,omitempty"`
}

func (x *Station) Reset() {
	*x = Station{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Station) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Station) ProtoMessage() {}

func (x *Station) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Station.ProtoReflect.Descriptor instead.
func (*Station) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *Station) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Station) GetCapacity() float32 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *Station) GetUsedCapacity() float32 {
	if x != nil {
		return x.UsedCapacity
	}
	return 0
}

func (x *Station) GetDocks() []*Dock {
	if x != nil {
		return x.Docks
	}
	return nil
}

func (x *Station) GetIsRegistered() bool {
	if x != nil {
		return x.IsRegistered
	}
	return false
}

type Stations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stations []*Station `protobuf:"bytes,1,rep,name=stations,proto3" json:"stations,omitempty"`
}

func (x *Stations) Reset() {
	*x = Stations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stations) ProtoMessage() {}

func (x *Stations) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stations.ProtoReflect.Descriptor instead.
func (*Stations) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *Stations) GetStations() []*Station {
	if x != nil {
		return x.Stations
	}
	return nil
}

type Ship struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status string  `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Weight float32 `protobuf:"fixed32,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Time   int32   `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Ship) Reset() {
	*x = Ship{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ship) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ship) ProtoMessage() {}

func (x *Ship) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ship.ProtoReflect.Descriptor instead.
func (*Ship) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *Ship) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Ship) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Ship) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Ship) GetTime() int32 {
	if x != nil {
		return x.Time
	}
	return 0
}

type Ships struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ships []*Ship `protobuf:"bytes,1,rep,name=ships,proto3" json:"ships,omitempty"`
}

func (x *Ships) Reset() {
	*x = Ships{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ships) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ships) ProtoMessage() {}

func (x *Ships) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ships.ProtoReflect.Descriptor instead.
func (*Ships) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *Ships) GetShips() []*Ship {
	if x != nil {
		return x.Ships
	}
	return nil
}

type Dock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NumDockingPorts int32   `protobuf:"varint,2,opt,name=numDockingPorts,proto3" json:"numDockingPorts,omitempty"`
	Occupied        int32   `protobuf:"varint,3,opt,name=occupied,proto3" json:"occupied,omitempty"`
	Weight          float32 `protobuf:"fixed32,4,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (x *Dock) Reset() {
	*x = Dock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dock) ProtoMessage() {}

func (x *Dock) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dock.ProtoReflect.Descriptor instead.
func (*Dock) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *Dock) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Dock) GetNumDockingPorts() int32 {
	if x != nil {
		return x.NumDockingPorts
	}
	return 0
}

func (x *Dock) GetOccupied() int32 {
	if x != nil {
		return x.Occupied
	}
	return 0
}

func (x *Dock) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Command        string `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Duration       int32  `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	DockingStation string `protobuf:"bytes,3,opt,name=dockingStation,proto3" json:"dockingStation,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *Command) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Command) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Command) GetDockingStation() string {
	if x != nil {
		return x.DockingStation
	}
	return ""
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6e, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x69, 0x64, 0x22,
	0x35, 0x0a, 0x0b, 0x4a, 0x57, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x65, 0x78, 0x70, 0x22, 0xb0, 0x01, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x22,
	0x0a, 0x0c, 0x75, 0x73, 0x65, 0x64, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x64, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x64, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x44, 0x6f, 0x63, 0x6b, 0x52, 0x05,
	0x64, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64, 0x22, 0x46, 0x0a, 0x08, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f,
	0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x5a, 0x0a, 0x04, 0x53, 0x68, 0x69, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x3a, 0x0a,
	0x05, 0x53, 0x68, 0x69, 0x70, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x74, 0x72,
	0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x53, 0x68,
	0x69, 0x70, 0x52, 0x05, 0x73, 0x68, 0x69, 0x70, 0x73, 0x22, 0x74, 0x0a, 0x04, 0x44, 0x6f, 0x63,
	0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x75, 0x6d, 0x44, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x50,
	0x6f, 0x72, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x6e, 0x75, 0x6d, 0x44,
	0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6f,
	0x63, 0x63, 0x75, 0x70, 0x69, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6f,
	0x63, 0x63, 0x75, 0x70, 0x69, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22,
	0x67, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x26, 0x0a, 0x0e, 0x64, 0x6f, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x6f, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xad, 0x01, 0x0a, 0x12, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x4e, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x22, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x4a, 0x57, 0x54, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x47, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x1f, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x00, 0x32, 0xbb, 0x02, 0x0a, 0x09, 0x43, 0x43, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1e, 0x2e, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0b, 0x61,
	0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1f, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x5f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x73,
	0x68, 0x69, 0x70, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c,
	0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x42, 0x0a, 0x08, 0x61, 0x6c, 0x6c, 0x53, 0x68, 0x69, 0x70, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1c, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x74,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x53,
	0x68, 0x69, 0x70, 0x73, 0x22, 0x00, 0x32, 0xa1, 0x01, 0x0a, 0x0f, 0x53, 0x68, 0x69, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x0b, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x61, 0x6e, 0x64, 0x12, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x1e, 0x2e, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x74,
	0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x07, 0x6c, 0x61, 0x6e, 0x64,
	0x69, 0x6e, 0x67, 0x12, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x1d, 0x5a, 0x1b, 0x61, 0x62,
	0x6f, 0x75, 0x72, 0x6f, 0x75, 0x6d, 0x69, 0x6e, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74,
	0x63, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_proto_goTypes = []interface{}{
	(*UserAuth)(nil),               // 0: space_traffic_control.UserAuth
	(*JWTResponse)(nil),            // 1: space_traffic_control.JWTResponse
	(*Station)(nil),                // 2: space_traffic_control.Station
	(*Stations)(nil),               // 3: space_traffic_control.Stations
	(*Ship)(nil),                   // 4: space_traffic_control.Ship
	(*Ships)(nil),                  // 5: space_traffic_control.Ships
	(*Dock)(nil),                   // 6: space_traffic_control.Dock
	(*Command)(nil),                // 7: space_traffic_control.Command
	(*wrapperspb.StringValue)(nil), // 8: google.protobuf.StringValue
	(*wrapperspb.FloatValue)(nil),  // 9: google.protobuf.FloatValue
	(*emptypb.Empty)(nil),          // 10: google.protobuf.Empty
	(*wrapperspb.Int32Value)(nil),  // 11: google.protobuf.Int32Value
	(*wrapperspb.BoolValue)(nil),   // 12: google.protobuf.BoolValue
}
var file_api_proto_depIdxs = []int32{
	6,  // 0: space_traffic_control.Station.docks:type_name -> space_traffic_control.Dock
	2,  // 1: space_traffic_control.Stations.stations:type_name -> space_traffic_control.Station
	4,  // 2: space_traffic_control.Ships.ships:type_name -> space_traffic_control.Ship
	0,  // 3: space_traffic_control.AuthenticationInfo.login:input_type -> space_traffic_control.UserAuth
	0,  // 4: space_traffic_control.AuthenticationInfo.signUp:input_type -> space_traffic_control.UserAuth
	2,  // 5: space_traffic_control.CCService.stationRegister:input_type -> space_traffic_control.Station
	8,  // 6: space_traffic_control.CCService.allStations:input_type -> google.protobuf.StringValue
	9,  // 7: space_traffic_control.CCService.shipRegister:input_type -> google.protobuf.FloatValue
	10, // 8: space_traffic_control.CCService.allShips:input_type -> google.protobuf.Empty
	11, // 9: space_traffic_control.ShippingStation.requestLand:input_type -> google.protobuf.Int32Value
	11, // 10: space_traffic_control.ShippingStation.landing:input_type -> google.protobuf.Int32Value
	1,  // 11: space_traffic_control.AuthenticationInfo.login:output_type -> space_traffic_control.JWTResponse
	12, // 12: space_traffic_control.AuthenticationInfo.signUp:output_type -> google.protobuf.BoolValue
	2,  // 13: space_traffic_control.CCService.stationRegister:output_type -> space_traffic_control.Station
	3,  // 14: space_traffic_control.CCService.allStations:output_type -> space_traffic_control.Stations
	10, // 15: space_traffic_control.CCService.shipRegister:output_type -> google.protobuf.Empty
	5,  // 16: space_traffic_control.CCService.allShips:output_type -> space_traffic_control.Ships
	7,  // 17: space_traffic_control.ShippingStation.requestLand:output_type -> space_traffic_control.Command
	10, // 18: space_traffic_control.ShippingStation.landing:output_type -> google.protobuf.Empty
	11, // [11:19] is the sub-list for method output_type
	3,  // [3:11] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAuth); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JWTResponse); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Station); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stations); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ship); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ships); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dock); i {
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
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
