// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.20.3
// source: dict/system_dict.proto

// system_dict 字典数据表

package dict

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SystemDictObject 数据对象
type SystemDictObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: db:"id" json:"id"
	Id *int64 `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id" db:"id"` //字典编码
	// @inject_tag: db:"sort" json:"sort"
	Sort *int32 `protobuf:"varint,2,opt,name=sort,proto3,oneof" json:"sort" db:"sort"` //字典排序
	// @inject_tag: db:"label" json:"label"
	Label *string `protobuf:"bytes,3,opt,name=label,proto3,oneof" json:"label" db:"label"` //字典标签
	// @inject_tag: db:"value" json:"value"
	Value *string `protobuf:"bytes,4,opt,name=value,proto3,oneof" json:"value" db:"value"` //字典键值
	// @inject_tag: db:"dict_type" json:"dictType"
	DictType *string `protobuf:"bytes,5,opt,name=dict_type,json=dictType,proto3,oneof" json:"dictType" db:"dict_type"` //字典类型
	// @inject_tag: db:"status" json:"status"
	Status *int32 `protobuf:"varint,6,opt,name=status,proto3,oneof" json:"status" db:"status"` //状态（0正常 1停用）
	// @inject_tag: db:"color_type" json:"colorType"
	ColorType *string `protobuf:"bytes,7,opt,name=color_type,json=colorType,proto3,oneof" json:"colorType" db:"color_type"` //颜色类型
	// @inject_tag: db:"css_class" json:"cssClass"
	CssClass *string `protobuf:"bytes,8,opt,name=css_class,json=cssClass,proto3,oneof" json:"cssClass" db:"css_class"` //css 样式
	// @inject_tag: db:"remark" json:"remark"
	Remark *string `protobuf:"bytes,9,opt,name=remark,proto3,oneof" json:"remark" db:"remark"` //备注
	// @inject_tag: db:"creator" json:"creator"
	Creator *string `protobuf:"bytes,10,opt,name=creator,proto3,oneof" json:"creator" db:"creator"` //创建人
	// @inject_tag: db:"create_time" json:"createTime"
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=create_time,json=createTime,proto3" json:"createTime" db:"create_time"` //创建时间
	// @inject_tag: db:"updater" json:"updater"
	Updater *string `protobuf:"bytes,12,opt,name=updater,proto3,oneof" json:"updater" db:"updater"` //更新人
	// @inject_tag: db:"update_time" json:"updateTime"
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=update_time,json=updateTime,proto3" json:"updateTime" db:"update_time"` //更新时间
}

func (x *SystemDictObject) Reset() {
	*x = SystemDictObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dict_system_dict_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemDictObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemDictObject) ProtoMessage() {}

func (x *SystemDictObject) ProtoReflect() protoreflect.Message {
	mi := &file_dict_system_dict_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemDictObject.ProtoReflect.Descriptor instead.
func (*SystemDictObject) Descriptor() ([]byte, []int) {
	return file_dict_system_dict_proto_rawDescGZIP(), []int{0}
}

func (x *SystemDictObject) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *SystemDictObject) GetSort() int32 {
	if x != nil && x.Sort != nil {
		return *x.Sort
	}
	return 0
}

func (x *SystemDictObject) GetLabel() string {
	if x != nil && x.Label != nil {
		return *x.Label
	}
	return ""
}

func (x *SystemDictObject) GetValue() string {
	if x != nil && x.Value != nil {
		return *x.Value
	}
	return ""
}

func (x *SystemDictObject) GetDictType() string {
	if x != nil && x.DictType != nil {
		return *x.DictType
	}
	return ""
}

func (x *SystemDictObject) GetStatus() int32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *SystemDictObject) GetColorType() string {
	if x != nil && x.ColorType != nil {
		return *x.ColorType
	}
	return ""
}

func (x *SystemDictObject) GetCssClass() string {
	if x != nil && x.CssClass != nil {
		return *x.CssClass
	}
	return ""
}

func (x *SystemDictObject) GetRemark() string {
	if x != nil && x.Remark != nil {
		return *x.Remark
	}
	return ""
}

func (x *SystemDictObject) GetCreator() string {
	if x != nil && x.Creator != nil {
		return *x.Creator
	}
	return ""
}

func (x *SystemDictObject) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *SystemDictObject) GetUpdater() string {
	if x != nil && x.Updater != nil {
		return *x.Updater
	}
	return ""
}

func (x *SystemDictObject) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

// SystemDictCreateRequest 创建数据请求
type SystemDictCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *SystemDictObject `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SystemDictCreateRequest) Reset() {
	*x = SystemDictCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dict_system_dict_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemDictCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemDictCreateRequest) ProtoMessage() {}

func (x *SystemDictCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dict_system_dict_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemDictCreateRequest.ProtoReflect.Descriptor instead.
func (*SystemDictCreateRequest) Descriptor() ([]byte, []int) {
	return file_dict_system_dict_proto_rawDescGZIP(), []int{1}
}

func (x *SystemDictCreateRequest) GetData() *SystemDictObject {
	if x != nil {
		return x.Data
	}
	return nil
}

// SystemDictCreateResponse 创建数据响应
type SystemDictCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data int64  `protobuf:"varint,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SystemDictCreateResponse) Reset() {
	*x = SystemDictCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dict_system_dict_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemDictCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemDictCreateResponse) ProtoMessage() {}

func (x *SystemDictCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dict_system_dict_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemDictCreateResponse.ProtoReflect.Descriptor instead.
func (*SystemDictCreateResponse) Descriptor() ([]byte, []int) {
	return file_dict_system_dict_proto_rawDescGZIP(), []int{2}
}

func (x *SystemDictCreateResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SystemDictCreateResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SystemDictCreateResponse) GetData() int64 {
	if x != nil {
		return x.Data
	}
	return 0
}

// SystemDictUpdateRequest 更新数据请求
type SystemDictUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: db:"system_dict_id" json:"systemDictId"
	SystemDictId int64             `protobuf:"varint,1,opt,name=system_dict_id,json=systemDictId,proto3" json:"systemDictId" db:"system_dict_id"` //字典编码
	Data         *SystemDictObject `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SystemDictUpdateRequest) Reset() {
	*x = SystemDictUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dict_system_dict_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemDictUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemDictUpdateRequest) ProtoMessage() {}

func (x *SystemDictUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dict_system_dict_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemDictUpdateRequest.ProtoReflect.Descriptor instead.
func (*SystemDictUpdateRequest) Descriptor() ([]byte, []int) {
	return file_dict_system_dict_proto_rawDescGZIP(), []int{3}
}

func (x *SystemDictUpdateRequest) GetSystemDictId() int64 {
	if x != nil {
		return x.SystemDictId
	}
	return 0
}

func (x *SystemDictUpdateRequest) GetData() *SystemDictObject {
	if x != nil {
		return x.Data
	}
	return nil
}

// SystemDictUpdateResponse 更新数据响应
type SystemDictUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SystemDictUpdateResponse) Reset() {
	*x = SystemDictUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dict_system_dict_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemDictUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemDictUpdateResponse) ProtoMessage() {}

func (x *SystemDictUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dict_system_dict_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemDictUpdateResponse.ProtoReflect.Descriptor instead.
func (*SystemDictUpdateResponse) Descriptor() ([]byte, []int) {
	return file_dict_system_dict_proto_rawDescGZIP(), []int{4}
}

func (x *SystemDictUpdateResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SystemDictUpdateResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// SystemDictDeleteRequest 删除数据请求
type SystemDictDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: db:"system_dict_id" json:"systemDictId"
	SystemDictId int64 `protobuf:"varint,1,opt,name=system_dict_id,json=systemDictId,proto3" json:"systemDictId" db:"system_dict_id"` //字典编码
}

func (x *SystemDictDeleteRequest) Reset() {
	*x = SystemDictDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dict_system_dict_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemDictDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemDictDeleteRequest) ProtoMessage() {}

func (x *SystemDictDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dict_system_dict_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemDictDeleteRequest.ProtoReflect.Descriptor instead.
func (*SystemDictDeleteRequest) Descriptor() ([]byte, []int) {
	return file_dict_system_dict_proto_rawDescGZIP(), []int{5}
}

func (x *SystemDictDeleteRequest) GetSystemDictId() int64 {
	if x != nil {
		return x.SystemDictId
	}
	return 0
}

// SystemDictDeleteResponse 删除数据响应
type SystemDictDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SystemDictDeleteResponse) Reset() {
	*x = SystemDictDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dict_system_dict_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemDictDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemDictDeleteResponse) ProtoMessage() {}

func (x *SystemDictDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dict_system_dict_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemDictDeleteResponse.ProtoReflect.Descriptor instead.
func (*SystemDictDeleteResponse) Descriptor() ([]byte, []int) {
	return file_dict_system_dict_proto_rawDescGZIP(), []int{6}
}

func (x *SystemDictDeleteResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SystemDictDeleteResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

// SystemDictRequest 查询单条数据请求
type SystemDictRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: db:"system_dict_id" json:"systemDictId"
	SystemDictId int64 `protobuf:"varint,1,opt,name=system_dict_id,json=systemDictId,proto3" json:"systemDictId" db:"system_dict_id"` //字典编码
}

func (x *SystemDictRequest) Reset() {
	*x = SystemDictRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dict_system_dict_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemDictRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemDictRequest) ProtoMessage() {}

func (x *SystemDictRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dict_system_dict_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemDictRequest.ProtoReflect.Descriptor instead.
func (*SystemDictRequest) Descriptor() ([]byte, []int) {
	return file_dict_system_dict_proto_rawDescGZIP(), []int{7}
}

func (x *SystemDictRequest) GetSystemDictId() int64 {
	if x != nil {
		return x.SystemDictId
	}
	return 0
}

// SystemDictResponse 查询单条数据响应
type SystemDictResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string            `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *SystemDictObject `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SystemDictResponse) Reset() {
	*x = SystemDictResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dict_system_dict_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemDictResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemDictResponse) ProtoMessage() {}

func (x *SystemDictResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dict_system_dict_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemDictResponse.ProtoReflect.Descriptor instead.
func (*SystemDictResponse) Descriptor() ([]byte, []int) {
	return file_dict_system_dict_proto_rawDescGZIP(), []int{8}
}

func (x *SystemDictResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SystemDictResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SystemDictResponse) GetData() *SystemDictObject {
	if x != nil {
		return x.Data
	}
	return nil
}

// SystemDictListRequest 列表数据
type SystemDictListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @inject_tag: db:"dict_type" json:"dictType"
	DictType *string `protobuf:"bytes,1,opt,name=dict_type,json=dictType,proto3,oneof" json:"dictType" db:"dict_type"` // 字典类型
	// @inject_tag: db:"status" json:"status"
	Status *int32 `protobuf:"varint,2,opt,name=status,proto3,oneof" json:"status" db:"status"` // 状态（0正常 1停用）
}

func (x *SystemDictListRequest) Reset() {
	*x = SystemDictListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dict_system_dict_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemDictListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemDictListRequest) ProtoMessage() {}

func (x *SystemDictListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dict_system_dict_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemDictListRequest.ProtoReflect.Descriptor instead.
func (*SystemDictListRequest) Descriptor() ([]byte, []int) {
	return file_dict_system_dict_proto_rawDescGZIP(), []int{9}
}

func (x *SystemDictListRequest) GetDictType() string {
	if x != nil && x.DictType != nil {
		return *x.DictType
	}
	return ""
}

func (x *SystemDictListRequest) GetStatus() int32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

// SystemDictListResponse 数据响应
type SystemDictListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64               `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string              `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data []*SystemDictObject `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SystemDictListResponse) Reset() {
	*x = SystemDictListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dict_system_dict_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemDictListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemDictListResponse) ProtoMessage() {}

func (x *SystemDictListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dict_system_dict_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemDictListResponse.ProtoReflect.Descriptor instead.
func (*SystemDictListResponse) Descriptor() ([]byte, []int) {
	return file_dict_system_dict_proto_rawDescGZIP(), []int{10}
}

func (x *SystemDictListResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SystemDictListResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SystemDictListResponse) GetData() []*SystemDictObject {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_dict_system_dict_proto protoreflect.FileDescriptor

var file_dict_system_dict_proto_rawDesc = []byte{
	0x0a, 0x16, 0x64, 0x69, 0x63, 0x74, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x64, 0x69,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x64, 0x69, 0x63, 0x74, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xcd, 0x04, 0x0a, 0x10, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x6f, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x64, 0x69, 0x63, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x08, 0x64,
	0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x48, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6f, 0x72,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x06, 0x52, 0x09, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x63,
	0x73, 0x73, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x07,
	0x52, 0x08, 0x63, 0x73, 0x73, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a,
	0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x08, 0x52,
	0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x09, 0x52, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x48, 0x0a, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x6f,
	0x72, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0x08, 0x0a, 0x06,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x64, 0x69, 0x63, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x63, 0x73, 0x73, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x22,
	0x45, 0x0a, 0x17, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x54, 0x0a, 0x18, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x44, 0x69, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6b, 0x0a, 0x17,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x5f, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x69,
	0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x40, 0x0a, 0x18, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x3f, 0x0a, 0x17, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x5f, 0x64, 0x69, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x18,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x39,
	0x0a, 0x11, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x64, 0x69,
	0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x49, 0x64, 0x22, 0x66, 0x0a, 0x12, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x44, 0x69, 0x63, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x6f, 0x0a, 0x15, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x09, 0x64, 0x69,
	0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x08, 0x64, 0x69, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x64, 0x69,
	0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x6a, 0x0a, 0x16, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x2a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69,
	0x63, 0x74, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x9a,
	0x03, 0x0a, 0x11, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69,
	0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x44, 0x69, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x64, 0x69,
	0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x64, 0x69, 0x63,
	0x74, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x10, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d,
	0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x64, 0x69, 0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a,
	0x0a, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x12, 0x17, 0x2e, 0x64, 0x69,
	0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b,
	0x0a, 0x0e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1b, 0x2e, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69,
	0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x64, 0x69, 0x63, 0x74, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x44, 0x69, 0x63, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2e,
	0x2f, 0x64, 0x69, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dict_system_dict_proto_rawDescOnce sync.Once
	file_dict_system_dict_proto_rawDescData = file_dict_system_dict_proto_rawDesc
)

func file_dict_system_dict_proto_rawDescGZIP() []byte {
	file_dict_system_dict_proto_rawDescOnce.Do(func() {
		file_dict_system_dict_proto_rawDescData = protoimpl.X.CompressGZIP(file_dict_system_dict_proto_rawDescData)
	})
	return file_dict_system_dict_proto_rawDescData
}

var file_dict_system_dict_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_dict_system_dict_proto_goTypes = []interface{}{
	(*SystemDictObject)(nil),         // 0: dict.SystemDictObject
	(*SystemDictCreateRequest)(nil),  // 1: dict.SystemDictCreateRequest
	(*SystemDictCreateResponse)(nil), // 2: dict.SystemDictCreateResponse
	(*SystemDictUpdateRequest)(nil),  // 3: dict.SystemDictUpdateRequest
	(*SystemDictUpdateResponse)(nil), // 4: dict.SystemDictUpdateResponse
	(*SystemDictDeleteRequest)(nil),  // 5: dict.SystemDictDeleteRequest
	(*SystemDictDeleteResponse)(nil), // 6: dict.SystemDictDeleteResponse
	(*SystemDictRequest)(nil),        // 7: dict.SystemDictRequest
	(*SystemDictResponse)(nil),       // 8: dict.SystemDictResponse
	(*SystemDictListRequest)(nil),    // 9: dict.SystemDictListRequest
	(*SystemDictListResponse)(nil),   // 10: dict.SystemDictListResponse
	(*timestamppb.Timestamp)(nil),    // 11: google.protobuf.Timestamp
}
var file_dict_system_dict_proto_depIdxs = []int32{
	11, // 0: dict.SystemDictObject.create_time:type_name -> google.protobuf.Timestamp
	11, // 1: dict.SystemDictObject.update_time:type_name -> google.protobuf.Timestamp
	0,  // 2: dict.SystemDictCreateRequest.data:type_name -> dict.SystemDictObject
	0,  // 3: dict.SystemDictUpdateRequest.data:type_name -> dict.SystemDictObject
	0,  // 4: dict.SystemDictResponse.data:type_name -> dict.SystemDictObject
	0,  // 5: dict.SystemDictListResponse.data:type_name -> dict.SystemDictObject
	1,  // 6: dict.SystemDictService.SystemDictCreate:input_type -> dict.SystemDictCreateRequest
	3,  // 7: dict.SystemDictService.SystemDictUpdate:input_type -> dict.SystemDictUpdateRequest
	5,  // 8: dict.SystemDictService.SystemDictDelete:input_type -> dict.SystemDictDeleteRequest
	7,  // 9: dict.SystemDictService.SystemDict:input_type -> dict.SystemDictRequest
	9,  // 10: dict.SystemDictService.SystemDictList:input_type -> dict.SystemDictListRequest
	2,  // 11: dict.SystemDictService.SystemDictCreate:output_type -> dict.SystemDictCreateResponse
	4,  // 12: dict.SystemDictService.SystemDictUpdate:output_type -> dict.SystemDictUpdateResponse
	6,  // 13: dict.SystemDictService.SystemDictDelete:output_type -> dict.SystemDictDeleteResponse
	8,  // 14: dict.SystemDictService.SystemDict:output_type -> dict.SystemDictResponse
	10, // 15: dict.SystemDictService.SystemDictList:output_type -> dict.SystemDictListResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_dict_system_dict_proto_init() }
func file_dict_system_dict_proto_init() {
	if File_dict_system_dict_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dict_system_dict_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemDictObject); i {
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
		file_dict_system_dict_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemDictCreateRequest); i {
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
		file_dict_system_dict_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemDictCreateResponse); i {
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
		file_dict_system_dict_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemDictUpdateRequest); i {
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
		file_dict_system_dict_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemDictUpdateResponse); i {
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
		file_dict_system_dict_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemDictDeleteRequest); i {
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
		file_dict_system_dict_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemDictDeleteResponse); i {
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
		file_dict_system_dict_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemDictRequest); i {
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
		file_dict_system_dict_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemDictResponse); i {
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
		file_dict_system_dict_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemDictListRequest); i {
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
		file_dict_system_dict_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemDictListResponse); i {
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
	file_dict_system_dict_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_dict_system_dict_proto_msgTypes[9].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dict_system_dict_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dict_system_dict_proto_goTypes,
		DependencyIndexes: file_dict_system_dict_proto_depIdxs,
		MessageInfos:      file_dict_system_dict_proto_msgTypes,
	}.Build()
	File_dict_system_dict_proto = out.File
	file_dict_system_dict_proto_rawDesc = nil
	file_dict_system_dict_proto_goTypes = nil
	file_dict_system_dict_proto_depIdxs = nil
}
