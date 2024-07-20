// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: repairorder/repairorder.proto

package repairorder

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request and Response messages
type CreateRepairOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string             `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	VehicleId  string             `protobuf:"bytes,2,opt,name=vehicle_id,json=vehicleId,proto3" json:"vehicle_id,omitempty"`
	Items      []*RepairOrderItem `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"` // Other fields as needed
}

func (x *CreateRepairOrderRequest) Reset() {
	*x = CreateRepairOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repairorder_repairorder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRepairOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRepairOrderRequest) ProtoMessage() {}

func (x *CreateRepairOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_repairorder_repairorder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRepairOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateRepairOrderRequest) Descriptor() ([]byte, []int) {
	return file_repairorder_repairorder_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRepairOrderRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CreateRepairOrderRequest) GetVehicleId() string {
	if x != nil {
		return x.VehicleId
	}
	return ""
}

func (x *CreateRepairOrderRequest) GetItems() []*RepairOrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type CreateRepairOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepairOrderId string `protobuf:"bytes,1,opt,name=repair_order_id,json=repairOrderId,proto3" json:"repair_order_id,omitempty"`
}

func (x *CreateRepairOrderResponse) Reset() {
	*x = CreateRepairOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repairorder_repairorder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRepairOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRepairOrderResponse) ProtoMessage() {}

func (x *CreateRepairOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_repairorder_repairorder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRepairOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateRepairOrderResponse) Descriptor() ([]byte, []int) {
	return file_repairorder_repairorder_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRepairOrderResponse) GetRepairOrderId() string {
	if x != nil {
		return x.RepairOrderId
	}
	return ""
}

type GetRepairOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepairOrderId string `protobuf:"bytes,1,opt,name=repair_order_id,json=repairOrderId,proto3" json:"repair_order_id,omitempty"`
}

func (x *GetRepairOrderRequest) Reset() {
	*x = GetRepairOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repairorder_repairorder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRepairOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRepairOrderRequest) ProtoMessage() {}

func (x *GetRepairOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_repairorder_repairorder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRepairOrderRequest.ProtoReflect.Descriptor instead.
func (*GetRepairOrderRequest) Descriptor() ([]byte, []int) {
	return file_repairorder_repairorder_proto_rawDescGZIP(), []int{2}
}

func (x *GetRepairOrderRequest) GetRepairOrderId() string {
	if x != nil {
		return x.RepairOrderId
	}
	return ""
}

type ListRepairOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *ListRepairOrdersRequest) Reset() {
	*x = ListRepairOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repairorder_repairorder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRepairOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRepairOrdersRequest) ProtoMessage() {}

func (x *ListRepairOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_repairorder_repairorder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRepairOrdersRequest.ProtoReflect.Descriptor instead.
func (*ListRepairOrdersRequest) Descriptor() ([]byte, []int) {
	return file_repairorder_repairorder_proto_rawDescGZIP(), []int{3}
}

func (x *ListRepairOrdersRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type ListRepairOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RepairOrders []*RepairOrder `protobuf:"bytes,1,rep,name=repair_orders,json=repairOrders,proto3" json:"repair_orders,omitempty"`
}

func (x *ListRepairOrdersResponse) Reset() {
	*x = ListRepairOrdersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repairorder_repairorder_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRepairOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRepairOrdersResponse) ProtoMessage() {}

func (x *ListRepairOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_repairorder_repairorder_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRepairOrdersResponse.ProtoReflect.Descriptor instead.
func (*ListRepairOrdersResponse) Descriptor() ([]byte, []int) {
	return file_repairorder_repairorder_proto_rawDescGZIP(), []int{4}
}

func (x *ListRepairOrdersResponse) GetRepairOrders() []*RepairOrder {
	if x != nil {
		return x.RepairOrders
	}
	return nil
}

type RepairOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId string             `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	VehicleId  string             `protobuf:"bytes,3,opt,name=vehicle_id,json=vehicleId,proto3" json:"vehicle_id,omitempty"`
	Items      []*RepairOrderItem `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	Status     string             `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"` // e.g., Pending, Completed
}

func (x *RepairOrder) Reset() {
	*x = RepairOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repairorder_repairorder_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepairOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepairOrder) ProtoMessage() {}

func (x *RepairOrder) ProtoReflect() protoreflect.Message {
	mi := &file_repairorder_repairorder_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepairOrder.ProtoReflect.Descriptor instead.
func (*RepairOrder) Descriptor() ([]byte, []int) {
	return file_repairorder_repairorder_proto_rawDescGZIP(), []int{5}
}

func (x *RepairOrder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RepairOrder) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *RepairOrder) GetVehicleId() string {
	if x != nil {
		return x.VehicleId
	}
	return ""
}

func (x *RepairOrder) GetItems() []*RepairOrderItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *RepairOrder) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type RepairOrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	RepairOrderId string `protobuf:"bytes,2,opt,name=repair_order_id,json=repairOrderId,proto3" json:"repair_order_id,omitempty"`
	ItemType      string `protobuf:"bytes,3,opt,name=item_type,json=itemType,proto3" json:"item_type,omitempty"` // e.g., Part, Labor
	Part          *Part  `protobuf:"bytes,4,opt,name=part,proto3" json:"part,omitempty"`
	Labor         *Labor `protobuf:"bytes,5,opt,name=labor,proto3" json:"labor,omitempty"`
}

func (x *RepairOrderItem) Reset() {
	*x = RepairOrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repairorder_repairorder_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RepairOrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RepairOrderItem) ProtoMessage() {}

func (x *RepairOrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_repairorder_repairorder_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RepairOrderItem.ProtoReflect.Descriptor instead.
func (*RepairOrderItem) Descriptor() ([]byte, []int) {
	return file_repairorder_repairorder_proto_rawDescGZIP(), []int{6}
}

func (x *RepairOrderItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RepairOrderItem) GetRepairOrderId() string {
	if x != nil {
		return x.RepairOrderId
	}
	return ""
}

func (x *RepairOrderItem) GetItemType() string {
	if x != nil {
		return x.ItemType
	}
	return ""
}

func (x *RepairOrderItem) GetPart() *Part {
	if x != nil {
		return x.Part
	}
	return nil
}

func (x *RepairOrderItem) GetLabor() *Labor {
	if x != nil {
		return x.Labor
	}
	return nil
}

type Part struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Quantity    int32   `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Cost        float64 `protobuf:"fixed64,4,opt,name=cost,proto3" json:"cost,omitempty"`
	Markup      float64 `protobuf:"fixed64,5,opt,name=markup,proto3" json:"markup,omitempty"`
	Subtotal    float64 `protobuf:"fixed64,6,opt,name=subtotal,proto3" json:"subtotal,omitempty"`
	Total       float64 `protobuf:"fixed64,7,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Part) Reset() {
	*x = Part{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repairorder_repairorder_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Part) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Part) ProtoMessage() {}

func (x *Part) ProtoReflect() protoreflect.Message {
	mi := &file_repairorder_repairorder_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Part.ProtoReflect.Descriptor instead.
func (*Part) Descriptor() ([]byte, []int) {
	return file_repairorder_repairorder_proto_rawDescGZIP(), []int{7}
}

func (x *Part) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Part) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Part) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Part) GetCost() float64 {
	if x != nil {
		return x.Cost
	}
	return 0
}

func (x *Part) GetMarkup() float64 {
	if x != nil {
		return x.Markup
	}
	return 0
}

func (x *Part) GetSubtotal() float64 {
	if x != nil {
		return x.Subtotal
	}
	return 0
}

func (x *Part) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type Labor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description     string             `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Hours           float64            `protobuf:"fixed64,3,opt,name=hours,proto3" json:"hours,omitempty"`
	Subtotal        float64            `protobuf:"fixed64,4,opt,name=subtotal,proto3" json:"subtotal,omitempty"`
	Total           float64            `protobuf:"fixed64,5,opt,name=total,proto3" json:"total,omitempty"`
	TechnicianLabor []*TechnicianLabor `protobuf:"bytes,6,rep,name=technician_labor,json=technicianLabor,proto3" json:"technician_labor,omitempty"`
}

func (x *Labor) Reset() {
	*x = Labor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repairorder_repairorder_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Labor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Labor) ProtoMessage() {}

func (x *Labor) ProtoReflect() protoreflect.Message {
	mi := &file_repairorder_repairorder_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Labor.ProtoReflect.Descriptor instead.
func (*Labor) Descriptor() ([]byte, []int) {
	return file_repairorder_repairorder_proto_rawDescGZIP(), []int{8}
}

func (x *Labor) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Labor) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Labor) GetHours() float64 {
	if x != nil {
		return x.Hours
	}
	return 0
}

func (x *Labor) GetSubtotal() float64 {
	if x != nil {
		return x.Subtotal
	}
	return 0
}

func (x *Labor) GetTotal() float64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *Labor) GetTechnicianLabor() []*TechnicianLabor {
	if x != nil {
		return x.TechnicianLabor
	}
	return nil
}

type TechnicianLabor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TechnicianId string  `protobuf:"bytes,1,opt,name=technician_id,json=technicianId,proto3" json:"technician_id,omitempty"`
	Hours        float64 `protobuf:"fixed64,2,opt,name=hours,proto3" json:"hours,omitempty"`
}

func (x *TechnicianLabor) Reset() {
	*x = TechnicianLabor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_repairorder_repairorder_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TechnicianLabor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TechnicianLabor) ProtoMessage() {}

func (x *TechnicianLabor) ProtoReflect() protoreflect.Message {
	mi := &file_repairorder_repairorder_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TechnicianLabor.ProtoReflect.Descriptor instead.
func (*TechnicianLabor) Descriptor() ([]byte, []int) {
	return file_repairorder_repairorder_proto_rawDescGZIP(), []int{9}
}

func (x *TechnicianLabor) GetTechnicianId() string {
	if x != nil {
		return x.TechnicianId
	}
	return ""
}

func (x *TechnicianLabor) GetHours() float64 {
	if x != nil {
		return x.Hours
	}
	return 0
}

var File_repairorder_repairorder_proto protoreflect.FileDescriptor

var file_repairorder_repairorder_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x72, 0x65, 0x70, 0x61, 0x69, 0x72, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x72, 0x65,
	0x70, 0x61, 0x69, 0x72, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x82, 0x01, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x52, 0x65,
	0x70, 0x61, 0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0x43, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x70, 0x61, 0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x61, 0x69, 0x72, 0x5f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x61,
	0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x61, 0x69, 0x72, 0x5f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x70,
	0x61, 0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3a, 0x0a, 0x17, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x61, 0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x31, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x61, 0x69, 0x72, 0x5f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x52, 0x65, 0x70, 0x61,
	0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x61, 0x69, 0x72, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x65, 0x70,
	0x61, 0x69, 0x72, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x70, 0x61, 0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19,
	0x0a, 0x04, 0x70, 0x61, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50,
	0x61, 0x72, 0x74, 0x52, 0x04, 0x70, 0x61, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x05, 0x6c, 0x61, 0x62,
	0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4c, 0x61, 0x62, 0x6f, 0x72,
	0x52, 0x05, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x22, 0xb2, 0x01, 0x0a, 0x04, 0x50, 0x61, 0x72, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x63, 0x6f,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x75, 0x70, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x75, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x75,
	0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x73, 0x75,
	0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0xbe, 0x01, 0x0a,
	0x05, 0x4c, 0x61, 0x62, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x75, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x08, 0x73, 0x75, 0x62, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x3b, 0x0a, 0x10, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x5f, 0x6c,
	0x61, 0x62, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x54, 0x65, 0x63,
	0x68, 0x6e, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x4c, 0x61, 0x62, 0x6f, 0x72, 0x52, 0x0f, 0x74, 0x65,
	0x63, 0x68, 0x6e, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x4c, 0x61, 0x62, 0x6f, 0x72, 0x22, 0x4c, 0x0a,
	0x0f, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x4c, 0x61, 0x62, 0x6f, 0x72,
	0x12, 0x23, 0x0a, 0x0d, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x69, 0x63,
	0x69, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x32, 0xe1, 0x01, 0x0a, 0x12,
	0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x4a, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x61,
	0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x61, 0x69,
	0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x52, 0x65, 0x70, 0x61, 0x69,
	0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x47, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x61, 0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x18, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x61, 0x69, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x61, 0x69,
	0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61,
	0x62, 0x69, 0x61, 0x6e, 0x72, 0x77, 0x78, 0x2f, 0x6d, 0x79, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x70, 0x61, 0x69, 0x72, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_repairorder_repairorder_proto_rawDescOnce sync.Once
	file_repairorder_repairorder_proto_rawDescData = file_repairorder_repairorder_proto_rawDesc
)

func file_repairorder_repairorder_proto_rawDescGZIP() []byte {
	file_repairorder_repairorder_proto_rawDescOnce.Do(func() {
		file_repairorder_repairorder_proto_rawDescData = protoimpl.X.CompressGZIP(file_repairorder_repairorder_proto_rawDescData)
	})
	return file_repairorder_repairorder_proto_rawDescData
}

var file_repairorder_repairorder_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_repairorder_repairorder_proto_goTypes = []interface{}{
	(*CreateRepairOrderRequest)(nil),  // 0: CreateRepairOrderRequest
	(*CreateRepairOrderResponse)(nil), // 1: CreateRepairOrderResponse
	(*GetRepairOrderRequest)(nil),     // 2: GetRepairOrderRequest
	(*ListRepairOrdersRequest)(nil),   // 3: ListRepairOrdersRequest
	(*ListRepairOrdersResponse)(nil),  // 4: ListRepairOrdersResponse
	(*RepairOrder)(nil),               // 5: RepairOrder
	(*RepairOrderItem)(nil),           // 6: RepairOrderItem
	(*Part)(nil),                      // 7: Part
	(*Labor)(nil),                     // 8: Labor
	(*TechnicianLabor)(nil),           // 9: TechnicianLabor
}
var file_repairorder_repairorder_proto_depIdxs = []int32{
	6, // 0: CreateRepairOrderRequest.items:type_name -> RepairOrderItem
	5, // 1: ListRepairOrdersResponse.repair_orders:type_name -> RepairOrder
	6, // 2: RepairOrder.items:type_name -> RepairOrderItem
	7, // 3: RepairOrderItem.part:type_name -> Part
	8, // 4: RepairOrderItem.labor:type_name -> Labor
	9, // 5: Labor.technician_labor:type_name -> TechnicianLabor
	0, // 6: RepairOrderService.CreateRepairOrder:input_type -> CreateRepairOrderRequest
	2, // 7: RepairOrderService.GetRepairOrder:input_type -> GetRepairOrderRequest
	3, // 8: RepairOrderService.ListRepairOrders:input_type -> ListRepairOrdersRequest
	1, // 9: RepairOrderService.CreateRepairOrder:output_type -> CreateRepairOrderResponse
	5, // 10: RepairOrderService.GetRepairOrder:output_type -> RepairOrder
	4, // 11: RepairOrderService.ListRepairOrders:output_type -> ListRepairOrdersResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_repairorder_repairorder_proto_init() }
func file_repairorder_repairorder_proto_init() {
	if File_repairorder_repairorder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_repairorder_repairorder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRepairOrderRequest); i {
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
		file_repairorder_repairorder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRepairOrderResponse); i {
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
		file_repairorder_repairorder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRepairOrderRequest); i {
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
		file_repairorder_repairorder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRepairOrdersRequest); i {
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
		file_repairorder_repairorder_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRepairOrdersResponse); i {
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
		file_repairorder_repairorder_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepairOrder); i {
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
		file_repairorder_repairorder_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RepairOrderItem); i {
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
		file_repairorder_repairorder_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Part); i {
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
		file_repairorder_repairorder_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Labor); i {
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
		file_repairorder_repairorder_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TechnicianLabor); i {
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
			RawDescriptor: file_repairorder_repairorder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_repairorder_repairorder_proto_goTypes,
		DependencyIndexes: file_repairorder_repairorder_proto_depIdxs,
		MessageInfos:      file_repairorder_repairorder_proto_msgTypes,
	}.Build()
	File_repairorder_repairorder_proto = out.File
	file_repairorder_repairorder_proto_rawDesc = nil
	file_repairorder_repairorder_proto_goTypes = nil
	file_repairorder_repairorder_proto_depIdxs = nil
}
