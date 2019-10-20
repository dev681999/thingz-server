// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mqtt.proto

package mqttproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Mqtt struct {
	// @inject_tag: bson:"_id"
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func MqttFromBytes(data []byte) (*Mqtt, error) {
	t := &Mqtt{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *Mqtt) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *Mqtt) Reset()                   { *m = Mqtt{} }
func (m *Mqtt) String() string           { return proto.CompactTextString(m) }
func (*Mqtt) ProtoMessage()              {}
func (*Mqtt) Descriptor() ([]byte, []int) {
	return fileDescriptor_35327d90702720f6, []int{0}
}

func (m *Mqtt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mqtt.Unmarshal(m, b)
}
func (m *Mqtt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mqtt.Marshal(b, m, deterministic)
}
func (m *Mqtt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mqtt.Merge(m, src)
}
func (m *Mqtt) XXX_Size() int {
	return xxx_messageInfo_Mqtt.Size(m)
}
func (m *Mqtt) XXX_DiscardUnknown() {
	xxx_messageInfo_Mqtt.DiscardUnknown(m)
}

var xxx_messageInfo_Mqtt proto.InternalMessageInfo

func (m *Mqtt) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Mqtt) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Mqtt) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Mqtt) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Channel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FloatValue           float64  `protobuf:"fixed64,2,opt,name=floatValue,proto3" json:"floatValue,omitempty"`
	StringValue          string   `protobuf:"bytes,3,opt,name=stringValue,proto3" json:"stringValue,omitempty"`
	BoolValue            bool     `protobuf:"varint,4,opt,name=boolValue,proto3" json:"boolValue,omitempty"`
	DataValue            string   `protobuf:"bytes,5,opt,name=dataValue,proto3" json:"dataValue,omitempty"`
	Unit                 int32    `protobuf:"varint,6,opt,name=unit,proto3" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func ChannelFromBytes(data []byte) (*Channel, error) {
	t := &Channel{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *Channel) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *Channel) Reset()                   { *m = Channel{} }
func (m *Channel) String() string           { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()              {}
func (*Channel) Descriptor() ([]byte, []int) {
	return fileDescriptor_35327d90702720f6, []int{1}
}

func (m *Channel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Channel.Unmarshal(m, b)
}
func (m *Channel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Channel.Marshal(b, m, deterministic)
}
func (m *Channel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Channel.Merge(m, src)
}
func (m *Channel) XXX_Size() int {
	return xxx_messageInfo_Channel.Size(m)
}
func (m *Channel) XXX_DiscardUnknown() {
	xxx_messageInfo_Channel.DiscardUnknown(m)
}

var xxx_messageInfo_Channel proto.InternalMessageInfo

func (m *Channel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Channel) GetFloatValue() float64 {
	if m != nil {
		return m.FloatValue
	}
	return 0
}

func (m *Channel) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *Channel) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *Channel) GetDataValue() string {
	if m != nil {
		return m.DataValue
	}
	return ""
}

func (m *Channel) GetUnit() int32 {
	if m != nil {
		return m.Unit
	}
	return 0
}

type Thing struct {
	Id                   string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Channels             []*Channel `protobuf:"bytes,2,rep,name=channels,proto3" json:"channels,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func ThingFromBytes(data []byte) (*Thing, error) {
	t := &Thing{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *Thing) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *Thing) Reset()                   { *m = Thing{} }
func (m *Thing) String() string           { return proto.CompactTextString(m) }
func (*Thing) ProtoMessage()              {}
func (*Thing) Descriptor() ([]byte, []int) {
	return fileDescriptor_35327d90702720f6, []int{2}
}

func (m *Thing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Thing.Unmarshal(m, b)
}
func (m *Thing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Thing.Marshal(b, m, deterministic)
}
func (m *Thing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Thing.Merge(m, src)
}
func (m *Thing) XXX_Size() int {
	return xxx_messageInfo_Thing.Size(m)
}
func (m *Thing) XXX_DiscardUnknown() {
	xxx_messageInfo_Thing.DiscardUnknown(m)
}

var xxx_messageInfo_Thing proto.InternalMessageInfo

func (m *Thing) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Thing) GetChannels() []*Channel {
	if m != nil {
		return m.Channels
	}
	return nil
}

type UpdateThingRequest struct {
	Thing                *Thing   `protobuf:"bytes,1,opt,name=thing,proto3" json:"thing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func UpdateThingRequestFromBytes(data []byte) (*UpdateThingRequest, error) {
	t := &UpdateThingRequest{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *UpdateThingRequest) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *UpdateThingRequest) Reset()                   { *m = UpdateThingRequest{} }
func (m *UpdateThingRequest) String() string           { return proto.CompactTextString(m) }
func (*UpdateThingRequest) ProtoMessage()              {}
func (*UpdateThingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_35327d90702720f6, []int{3}
}

func (m *UpdateThingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateThingRequest.Unmarshal(m, b)
}
func (m *UpdateThingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateThingRequest.Marshal(b, m, deterministic)
}
func (m *UpdateThingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateThingRequest.Merge(m, src)
}
func (m *UpdateThingRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateThingRequest.Size(m)
}
func (m *UpdateThingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateThingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateThingRequest proto.InternalMessageInfo

func (m *UpdateThingRequest) GetThing() *Thing {
	if m != nil {
		return m.Thing
	}
	return nil
}

type UpdateThingResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func UpdateThingResponseFromBytes(data []byte) (*UpdateThingResponse, error) {
	t := &UpdateThingResponse{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *UpdateThingResponse) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *UpdateThingResponse) Reset()                   { *m = UpdateThingResponse{} }
func (m *UpdateThingResponse) String() string           { return proto.CompactTextString(m) }
func (*UpdateThingResponse) ProtoMessage()              {}
func (*UpdateThingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_35327d90702720f6, []int{4}
}

func (m *UpdateThingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateThingResponse.Unmarshal(m, b)
}
func (m *UpdateThingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateThingResponse.Marshal(b, m, deterministic)
}
func (m *UpdateThingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateThingResponse.Merge(m, src)
}
func (m *UpdateThingResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateThingResponse.Size(m)
}
func (m *UpdateThingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateThingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateThingResponse proto.InternalMessageInfo

func (m *UpdateThingResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *UpdateThingResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type CreateMqttRequest struct {
	Mqtt                 *Mqtt    `protobuf:"bytes,1,opt,name=mqtt,proto3" json:"mqtt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func CreateMqttRequestFromBytes(data []byte) (*CreateMqttRequest, error) {
	t := &CreateMqttRequest{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *CreateMqttRequest) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *CreateMqttRequest) Reset()                   { *m = CreateMqttRequest{} }
func (m *CreateMqttRequest) String() string           { return proto.CompactTextString(m) }
func (*CreateMqttRequest) ProtoMessage()              {}
func (*CreateMqttRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_35327d90702720f6, []int{5}
}

func (m *CreateMqttRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMqttRequest.Unmarshal(m, b)
}
func (m *CreateMqttRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMqttRequest.Marshal(b, m, deterministic)
}
func (m *CreateMqttRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMqttRequest.Merge(m, src)
}
func (m *CreateMqttRequest) XXX_Size() int {
	return xxx_messageInfo_CreateMqttRequest.Size(m)
}
func (m *CreateMqttRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMqttRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMqttRequest proto.InternalMessageInfo

func (m *CreateMqttRequest) GetMqtt() *Mqtt {
	if m != nil {
		return m.Mqtt
	}
	return nil
}

type CreateMqttResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func CreateMqttResponseFromBytes(data []byte) (*CreateMqttResponse, error) {
	t := &CreateMqttResponse{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *CreateMqttResponse) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *CreateMqttResponse) Reset()                   { *m = CreateMqttResponse{} }
func (m *CreateMqttResponse) String() string           { return proto.CompactTextString(m) }
func (*CreateMqttResponse) ProtoMessage()              {}
func (*CreateMqttResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_35327d90702720f6, []int{6}
}

func (m *CreateMqttResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMqttResponse.Unmarshal(m, b)
}
func (m *CreateMqttResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMqttResponse.Marshal(b, m, deterministic)
}
func (m *CreateMqttResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMqttResponse.Merge(m, src)
}
func (m *CreateMqttResponse) XXX_Size() int {
	return xxx_messageInfo_CreateMqttResponse.Size(m)
}
func (m *CreateMqttResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMqttResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMqttResponse proto.InternalMessageInfo

func (m *CreateMqttResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CreateMqttResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateMqttResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UserMqttsRequest struct {
	Owner                string   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func UserMqttsRequestFromBytes(data []byte) (*UserMqttsRequest, error) {
	t := &UserMqttsRequest{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *UserMqttsRequest) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *UserMqttsRequest) Reset()                   { *m = UserMqttsRequest{} }
func (m *UserMqttsRequest) String() string           { return proto.CompactTextString(m) }
func (*UserMqttsRequest) ProtoMessage()              {}
func (*UserMqttsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_35327d90702720f6, []int{7}
}

func (m *UserMqttsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserMqttsRequest.Unmarshal(m, b)
}
func (m *UserMqttsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserMqttsRequest.Marshal(b, m, deterministic)
}
func (m *UserMqttsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserMqttsRequest.Merge(m, src)
}
func (m *UserMqttsRequest) XXX_Size() int {
	return xxx_messageInfo_UserMqttsRequest.Size(m)
}
func (m *UserMqttsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserMqttsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserMqttsRequest proto.InternalMessageInfo

func (m *UserMqttsRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type UserMqttsResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Mqtts                []*Mqtt  `protobuf:"bytes,3,rep,name=mqtts,proto3" json:"mqtts,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func UserMqttsResponseFromBytes(data []byte) (*UserMqttsResponse, error) {
	t := &UserMqttsResponse{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *UserMqttsResponse) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *UserMqttsResponse) Reset()                   { *m = UserMqttsResponse{} }
func (m *UserMqttsResponse) String() string           { return proto.CompactTextString(m) }
func (*UserMqttsResponse) ProtoMessage()              {}
func (*UserMqttsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_35327d90702720f6, []int{8}
}

func (m *UserMqttsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserMqttsResponse.Unmarshal(m, b)
}
func (m *UserMqttsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserMqttsResponse.Marshal(b, m, deterministic)
}
func (m *UserMqttsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserMqttsResponse.Merge(m, src)
}
func (m *UserMqttsResponse) XXX_Size() int {
	return xxx_messageInfo_UserMqttsResponse.Size(m)
}
func (m *UserMqttsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserMqttsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserMqttsResponse proto.InternalMessageInfo

func (m *UserMqttsResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *UserMqttsResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *UserMqttsResponse) GetMqtts() []*Mqtt {
	if m != nil {
		return m.Mqtts
	}
	return nil
}

type GetMqttRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func GetMqttRequestFromBytes(data []byte) (*GetMqttRequest, error) {
	t := &GetMqttRequest{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *GetMqttRequest) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *GetMqttRequest) Reset()                   { *m = GetMqttRequest{} }
func (m *GetMqttRequest) String() string           { return proto.CompactTextString(m) }
func (*GetMqttRequest) ProtoMessage()              {}
func (*GetMqttRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_35327d90702720f6, []int{9}
}

func (m *GetMqttRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMqttRequest.Unmarshal(m, b)
}
func (m *GetMqttRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMqttRequest.Marshal(b, m, deterministic)
}
func (m *GetMqttRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMqttRequest.Merge(m, src)
}
func (m *GetMqttRequest) XXX_Size() int {
	return xxx_messageInfo_GetMqttRequest.Size(m)
}
func (m *GetMqttRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMqttRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMqttRequest proto.InternalMessageInfo

func (m *GetMqttRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetMqttRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type GetMqttResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Mqtt                 *Mqtt    `protobuf:"bytes,3,opt,name=mqtt,proto3" json:"mqtt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func GetMqttResponseFromBytes(data []byte) (*GetMqttResponse, error) {
	t := &GetMqttResponse{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *GetMqttResponse) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *GetMqttResponse) Reset()                   { *m = GetMqttResponse{} }
func (m *GetMqttResponse) String() string           { return proto.CompactTextString(m) }
func (*GetMqttResponse) ProtoMessage()              {}
func (*GetMqttResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_35327d90702720f6, []int{10}
}

func (m *GetMqttResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMqttResponse.Unmarshal(m, b)
}
func (m *GetMqttResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMqttResponse.Marshal(b, m, deterministic)
}
func (m *GetMqttResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMqttResponse.Merge(m, src)
}
func (m *GetMqttResponse) XXX_Size() int {
	return xxx_messageInfo_GetMqttResponse.Size(m)
}
func (m *GetMqttResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMqttResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMqttResponse proto.InternalMessageInfo

func (m *GetMqttResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *GetMqttResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *GetMqttResponse) GetMqtt() *Mqtt {
	if m != nil {
		return m.Mqtt
	}
	return nil
}

type DeleteMqttRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func DeleteMqttRequestFromBytes(data []byte) (*DeleteMqttRequest, error) {
	t := &DeleteMqttRequest{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *DeleteMqttRequest) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *DeleteMqttRequest) Reset()                   { *m = DeleteMqttRequest{} }
func (m *DeleteMqttRequest) String() string           { return proto.CompactTextString(m) }
func (*DeleteMqttRequest) ProtoMessage()              {}
func (*DeleteMqttRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_35327d90702720f6, []int{11}
}

func (m *DeleteMqttRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMqttRequest.Unmarshal(m, b)
}
func (m *DeleteMqttRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMqttRequest.Marshal(b, m, deterministic)
}
func (m *DeleteMqttRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMqttRequest.Merge(m, src)
}
func (m *DeleteMqttRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteMqttRequest.Size(m)
}
func (m *DeleteMqttRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMqttRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMqttRequest proto.InternalMessageInfo

func (m *DeleteMqttRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteMqttRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type DeleteMqttResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func DeleteMqttResponseFromBytes(data []byte) (*DeleteMqttResponse, error) {
	t := &DeleteMqttResponse{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *DeleteMqttResponse) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *DeleteMqttResponse) Reset()                   { *m = DeleteMqttResponse{} }
func (m *DeleteMqttResponse) String() string           { return proto.CompactTextString(m) }
func (*DeleteMqttResponse) ProtoMessage()              {}
func (*DeleteMqttResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_35327d90702720f6, []int{12}
}

func (m *DeleteMqttResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMqttResponse.Unmarshal(m, b)
}
func (m *DeleteMqttResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMqttResponse.Marshal(b, m, deterministic)
}
func (m *DeleteMqttResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMqttResponse.Merge(m, src)
}
func (m *DeleteMqttResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteMqttResponse.Size(m)
}
func (m *DeleteMqttResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMqttResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMqttResponse proto.InternalMessageInfo

func (m *DeleteMqttResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *DeleteMqttResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type UpdateMqttRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func UpdateMqttRequestFromBytes(data []byte) (*UpdateMqttRequest, error) {
	t := &UpdateMqttRequest{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *UpdateMqttRequest) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *UpdateMqttRequest) Reset()                   { *m = UpdateMqttRequest{} }
func (m *UpdateMqttRequest) String() string           { return proto.CompactTextString(m) }
func (*UpdateMqttRequest) ProtoMessage()              {}
func (*UpdateMqttRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_35327d90702720f6, []int{13}
}

func (m *UpdateMqttRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMqttRequest.Unmarshal(m, b)
}
func (m *UpdateMqttRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMqttRequest.Marshal(b, m, deterministic)
}
func (m *UpdateMqttRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMqttRequest.Merge(m, src)
}
func (m *UpdateMqttRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateMqttRequest.Size(m)
}
func (m *UpdateMqttRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMqttRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMqttRequest proto.InternalMessageInfo

func (m *UpdateMqttRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateMqttRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *UpdateMqttRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateMqttRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type UpdateMqttResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func UpdateMqttResponseFromBytes(data []byte) (*UpdateMqttResponse, error) {
	t := &UpdateMqttResponse{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *UpdateMqttResponse) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *UpdateMqttResponse) Reset()                   { *m = UpdateMqttResponse{} }
func (m *UpdateMqttResponse) String() string           { return proto.CompactTextString(m) }
func (*UpdateMqttResponse) ProtoMessage()              {}
func (*UpdateMqttResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_35327d90702720f6, []int{14}
}

func (m *UpdateMqttResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateMqttResponse.Unmarshal(m, b)
}
func (m *UpdateMqttResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateMqttResponse.Marshal(b, m, deterministic)
}
func (m *UpdateMqttResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMqttResponse.Merge(m, src)
}
func (m *UpdateMqttResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateMqttResponse.Size(m)
}
func (m *UpdateMqttResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMqttResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMqttResponse proto.InternalMessageInfo

func (m *UpdateMqttResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *UpdateMqttResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Mqtt)(nil), "mqttproto.Mqtt")
	proto.RegisterType((*Channel)(nil), "mqttproto.Channel")
	proto.RegisterType((*Thing)(nil), "mqttproto.Thing")
	proto.RegisterType((*UpdateThingRequest)(nil), "mqttproto.UpdateThingRequest")
	proto.RegisterType((*UpdateThingResponse)(nil), "mqttproto.UpdateThingResponse")
	proto.RegisterType((*CreateMqttRequest)(nil), "mqttproto.CreateMqttRequest")
	proto.RegisterType((*CreateMqttResponse)(nil), "mqttproto.CreateMqttResponse")
	proto.RegisterType((*UserMqttsRequest)(nil), "mqttproto.UserMqttsRequest")
	proto.RegisterType((*UserMqttsResponse)(nil), "mqttproto.UserMqttsResponse")
	proto.RegisterType((*GetMqttRequest)(nil), "mqttproto.GetMqttRequest")
	proto.RegisterType((*GetMqttResponse)(nil), "mqttproto.GetMqttResponse")
	proto.RegisterType((*DeleteMqttRequest)(nil), "mqttproto.DeleteMqttRequest")
	proto.RegisterType((*DeleteMqttResponse)(nil), "mqttproto.DeleteMqttResponse")
	proto.RegisterType((*UpdateMqttRequest)(nil), "mqttproto.UpdateMqttRequest")
	proto.RegisterType((*UpdateMqttResponse)(nil), "mqttproto.UpdateMqttResponse")
}

func init() { proto.RegisterFile("mqtt.proto", fileDescriptor_35327d90702720f6) }

var fileDescriptor_35327d90702720f6 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x55, 0x36, 0x9b, 0x76, 0x77, 0x56, 0x6a, 0xbb, 0x86, 0x83, 0x0f, 0x08, 0xad, 0x8c, 0x40,
	0x39, 0xe5, 0x50, 0x24, 0x04, 0x12, 0xb7, 0x2e, 0xea, 0x89, 0x4b, 0xd4, 0x72, 0xf7, 0x26, 0xd3,
	0x36, 0x28, 0xb5, 0xb3, 0xb6, 0x23, 0x7e, 0x14, 0x7f, 0x12, 0x79, 0xbc, 0x49, 0x5d, 0xad, 0x84,
	0x4a, 0xb8, 0x79, 0x3e, 0xdf, 0x9b, 0x79, 0x93, 0x00, 0x3c, 0xee, 0x9d, 0x2b, 0x3a, 0xa3, 0x9d,
	0x66, 0x4b, 0xff, 0xa6, 0xa7, 0xd8, 0xc1, 0xfc, 0xfb, 0xde, 0x39, 0x76, 0x06, 0xb3, 0xa6, 0xe6,
	0xc9, 0x26, 0xc9, 0x97, 0xe5, 0xac, 0xa9, 0xd9, 0x6b, 0xc8, 0xf4, 0x2f, 0x85, 0x86, 0xcf, 0xc8,
	0x15, 0x0c, 0xc6, 0x60, 0xae, 0xe4, 0x23, 0xf2, 0x94, 0x9c, 0xf4, 0x66, 0x1b, 0x58, 0xd5, 0x68,
	0x2b, 0xd3, 0x74, 0xae, 0xd1, 0x8a, 0xcf, 0x29, 0x14, 0xbb, 0xc4, 0xef, 0x04, 0x4e, 0xaf, 0x1e,
	0xa4, 0x52, 0xd8, 0x1e, 0xe1, 0xbc, 0x05, 0xb8, 0x6b, 0xb5, 0x74, 0x3f, 0x64, 0xdb, 0x23, 0x81,
	0x25, 0x65, 0xe4, 0xf1, 0xdd, 0xad, 0x33, 0x8d, 0xba, 0x0f, 0x09, 0x01, 0x38, 0x76, 0xb1, 0x37,
	0xb0, 0xdc, 0x69, 0xdd, 0x86, 0xb8, 0x47, 0x5f, 0x94, 0x4f, 0x0e, 0x1f, 0xad, 0xa5, 0x93, 0x21,
	0x9a, 0x51, 0xf5, 0x93, 0xc3, 0xcf, 0xd3, 0xab, 0xc6, 0xf1, 0x93, 0x4d, 0x92, 0x67, 0x25, 0xbd,
	0xc5, 0x35, 0x64, 0x37, 0x0f, 0x8d, 0xba, 0x3f, 0xa2, 0x5a, 0xc0, 0xa2, 0x0a, 0x53, 0x58, 0x3e,
	0xdb, 0xa4, 0xf9, 0xea, 0x92, 0x15, 0xe3, 0x22, 0x8b, 0xc3, 0x80, 0xe5, 0x98, 0x23, 0xbe, 0x02,
	0xbb, 0xed, 0x6a, 0xe9, 0x90, 0xda, 0x95, 0xb8, 0xef, 0xd1, 0x3a, 0xf6, 0x01, 0x32, 0xe7, 0x6d,
	0x6a, 0xbc, 0xba, 0xbc, 0x88, 0x5a, 0x84, 0xbc, 0x10, 0x16, 0xdf, 0xe0, 0xd5, 0xb3, 0x6a, 0xdb,
	0x69, 0x65, 0x91, 0x71, 0x38, 0xb5, 0x7d, 0x55, 0xa1, 0xb5, 0xd4, 0x60, 0x51, 0x0e, 0xa6, 0x57,
	0x0c, 0x8d, 0xd1, 0xa3, 0x62, 0x64, 0x88, 0xcf, 0xb0, 0xbe, 0x32, 0x28, 0x1d, 0x7a, 0x95, 0x07,
	0x0e, 0xef, 0x60, 0xee, 0x51, 0x0f, 0x14, 0xce, 0x23, 0x0a, 0x94, 0x45, 0x41, 0x71, 0x03, 0x2c,
	0xae, 0x9c, 0x86, 0x7f, 0x58, 0x62, 0x3a, 0x2c, 0x51, 0xe4, 0x70, 0x71, 0x6b, 0xd1, 0xf8, 0x9e,
	0x76, 0xa0, 0x33, 0xde, 0x5a, 0x12, 0xdd, 0x9a, 0xf8, 0x09, 0xeb, 0x28, 0x73, 0x22, 0xfc, 0x7b,
	0xc8, 0xfc, 0x30, 0x96, 0xa7, 0x24, 0xd8, 0xd1, 0xa8, 0x21, 0x2a, 0x3e, 0xc1, 0xd9, 0x35, 0xba,
	0x78, 0x45, 0x2f, 0xfa, 0x1e, 0xc4, 0x1d, 0x9c, 0x8f, 0x75, 0x13, 0x19, 0x0e, 0x5a, 0xa4, 0x7f,
	0xd3, 0xe2, 0x0b, 0xac, 0xb7, 0xd8, 0xe2, 0x73, 0x15, 0x5f, 0x46, 0x71, 0x0b, 0x2c, 0x2e, 0x9d,
	0x78, 0x46, 0x1a, 0xd6, 0xe1, 0x1a, 0xff, 0x99, 0xc0, 0xc4, 0x7f, 0xc6, 0x76, 0xf8, 0x78, 0xfe,
	0x87, 0xf6, 0xee, 0x84, 0x36, 0xf9, 0xf1, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xde, 0xc3, 0xa5,
	0x4a, 0xfd, 0x04, 0x00, 0x00,
}
