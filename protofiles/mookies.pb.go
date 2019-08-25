// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protofiles/mookies.proto

package mookiespb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Item struct {
	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id    int64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Price float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	// @inject_tag: db:"category_id"
	CategoryID int64 `protobuf:"varint,4,opt,name=categoryID,proto3" json:"categoryID,omitempty" db:"category_id"`
	// @inject_tag: db:"order_item_id"
	OrderItemID          int64     `protobuf:"varint,5,opt,name=orderItemID,proto3" json:"orderItemID,omitempty" db:"order_item_id"`
	Options              []*Option `protobuf:"bytes,6,rep,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_a42f9bdfdd6c12cf, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Item) GetCategoryID() int64 {
	if m != nil {
		return m.CategoryID
	}
	return 0
}

func (m *Item) GetOrderItemID() int64 {
	if m != nil {
		return m.OrderItemID
	}
	return 0
}

func (m *Item) GetOptions() []*Option {
	if m != nil {
		return m.Options
	}
	return nil
}

type Option struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Price                float32  `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Selected             bool     `protobuf:"varint,4,opt,name=selected,proto3" json:"selected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Option) Reset()         { *m = Option{} }
func (m *Option) String() string { return proto.CompactTextString(m) }
func (*Option) ProtoMessage()    {}
func (*Option) Descriptor() ([]byte, []int) {
	return fileDescriptor_a42f9bdfdd6c12cf, []int{1}
}

func (m *Option) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Option.Unmarshal(m, b)
}
func (m *Option) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Option.Marshal(b, m, deterministic)
}
func (m *Option) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Option.Merge(m, src)
}
func (m *Option) XXX_Size() int {
	return xxx_messageInfo_Option.Size(m)
}
func (m *Option) XXX_DiscardUnknown() {
	xxx_messageInfo_Option.DiscardUnknown(m)
}

var xxx_messageInfo_Option proto.InternalMessageInfo

func (m *Option) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Option) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Option) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Option) GetSelected() bool {
	if m != nil {
		return m.Selected
	}
	return false
}

type Category struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Items                []*Item  `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_a42f9bdfdd6c12cf, []int{2}
}

func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (m *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(m, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Category) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Category) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type Menu struct {
	Categories           []*Category `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Menu) Reset()         { *m = Menu{} }
func (m *Menu) String() string { return proto.CompactTextString(m) }
func (*Menu) ProtoMessage()    {}
func (*Menu) Descriptor() ([]byte, []int) {
	return fileDescriptor_a42f9bdfdd6c12cf, []int{3}
}

func (m *Menu) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Menu.Unmarshal(m, b)
}
func (m *Menu) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Menu.Marshal(b, m, deterministic)
}
func (m *Menu) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Menu.Merge(m, src)
}
func (m *Menu) XXX_Size() int {
	return xxx_messageInfo_Menu.Size(m)
}
func (m *Menu) XXX_DiscardUnknown() {
	xxx_messageInfo_Menu.DiscardUnknown(m)
}

var xxx_messageInfo_Menu proto.InternalMessageInfo

func (m *Menu) GetCategories() []*Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

type Order struct {
	Id     int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Items  []*Item `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
	Total  float32 `protobuf:"fixed32,4,opt,name=total,proto3" json:"total,omitempty"`
	Status string  `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	// @inject_tag: db:"time_ordered"
	TimeOrdered string `protobuf:"bytes,6,opt,name=time_ordered,json=timeOrdered,proto3" json:"time_ordered,omitempty" db:"time_ordered"`
	// @inject_tag: db:"time_complete"
	TimeComplete         string   `protobuf:"bytes,7,opt,name=time_complete,json=timeComplete,proto3" json:"time_complete,omitempty" db:"time_complete"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_a42f9bdfdd6c12cf, []int{4}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Order) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Order) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Order) GetTotal() float32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Order) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Order) GetTimeOrdered() string {
	if m != nil {
		return m.TimeOrdered
	}
	return ""
}

func (m *Order) GetTimeComplete() string {
	if m != nil {
		return m.TimeComplete
	}
	return ""
}

type Response struct {
	Response             string   `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_a42f9bdfdd6c12cf, []int{5}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

type CreateMenuItemResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateMenuItemResponse) Reset()         { *m = CreateMenuItemResponse{} }
func (m *CreateMenuItemResponse) String() string { return proto.CompactTextString(m) }
func (*CreateMenuItemResponse) ProtoMessage()    {}
func (*CreateMenuItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a42f9bdfdd6c12cf, []int{6}
}

func (m *CreateMenuItemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateMenuItemResponse.Unmarshal(m, b)
}
func (m *CreateMenuItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateMenuItemResponse.Marshal(b, m, deterministic)
}
func (m *CreateMenuItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateMenuItemResponse.Merge(m, src)
}
func (m *CreateMenuItemResponse) XXX_Size() int {
	return xxx_messageInfo_CreateMenuItemResponse.Size(m)
}
func (m *CreateMenuItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateMenuItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateMenuItemResponse proto.InternalMessageInfo

func (m *CreateMenuItemResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CompleteOrderRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompleteOrderRequest) Reset()         { *m = CompleteOrderRequest{} }
func (m *CompleteOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CompleteOrderRequest) ProtoMessage()    {}
func (*CompleteOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a42f9bdfdd6c12cf, []int{7}
}

func (m *CompleteOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompleteOrderRequest.Unmarshal(m, b)
}
func (m *CompleteOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompleteOrderRequest.Marshal(b, m, deterministic)
}
func (m *CompleteOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompleteOrderRequest.Merge(m, src)
}
func (m *CompleteOrderRequest) XXX_Size() int {
	return xxx_messageInfo_CompleteOrderRequest.Size(m)
}
func (m *CompleteOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CompleteOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CompleteOrderRequest proto.InternalMessageInfo

func (m *CompleteOrderRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type OrdersRequest struct {
	Request              string   `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrdersRequest) Reset()         { *m = OrdersRequest{} }
func (m *OrdersRequest) String() string { return proto.CompactTextString(m) }
func (*OrdersRequest) ProtoMessage()    {}
func (*OrdersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a42f9bdfdd6c12cf, []int{8}
}

func (m *OrdersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrdersRequest.Unmarshal(m, b)
}
func (m *OrdersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrdersRequest.Marshal(b, m, deterministic)
}
func (m *OrdersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrdersRequest.Merge(m, src)
}
func (m *OrdersRequest) XXX_Size() int {
	return xxx_messageInfo_OrdersRequest.Size(m)
}
func (m *OrdersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrdersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrdersRequest proto.InternalMessageInfo

func (m *OrdersRequest) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

type OrdersResponse struct {
	Orders               []*Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrdersResponse) Reset()         { *m = OrdersResponse{} }
func (m *OrdersResponse) String() string { return proto.CompactTextString(m) }
func (*OrdersResponse) ProtoMessage()    {}
func (*OrdersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a42f9bdfdd6c12cf, []int{9}
}

func (m *OrdersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrdersResponse.Unmarshal(m, b)
}
func (m *OrdersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrdersResponse.Marshal(b, m, deterministic)
}
func (m *OrdersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrdersResponse.Merge(m, src)
}
func (m *OrdersResponse) XXX_Size() int {
	return xxx_messageInfo_OrdersResponse.Size(m)
}
func (m *OrdersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrdersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrdersResponse proto.InternalMessageInfo

func (m *OrdersResponse) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_a42f9bdfdd6c12cf, []int{10}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type DeleteMenuItemRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteMenuItemRequest) Reset()         { *m = DeleteMenuItemRequest{} }
func (m *DeleteMenuItemRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteMenuItemRequest) ProtoMessage()    {}
func (*DeleteMenuItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a42f9bdfdd6c12cf, []int{11}
}

func (m *DeleteMenuItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteMenuItemRequest.Unmarshal(m, b)
}
func (m *DeleteMenuItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteMenuItemRequest.Marshal(b, m, deterministic)
}
func (m *DeleteMenuItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteMenuItemRequest.Merge(m, src)
}
func (m *DeleteMenuItemRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteMenuItemRequest.Size(m)
}
func (m *DeleteMenuItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteMenuItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteMenuItemRequest proto.InternalMessageInfo

func (m *DeleteMenuItemRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Item)(nil), "mookiespb.Item")
	proto.RegisterType((*Option)(nil), "mookiespb.Option")
	proto.RegisterType((*Category)(nil), "mookiespb.Category")
	proto.RegisterType((*Menu)(nil), "mookiespb.Menu")
	proto.RegisterType((*Order)(nil), "mookiespb.Order")
	proto.RegisterType((*Response)(nil), "mookiespb.Response")
	proto.RegisterType((*CreateMenuItemResponse)(nil), "mookiespb.CreateMenuItemResponse")
	proto.RegisterType((*CompleteOrderRequest)(nil), "mookiespb.CompleteOrderRequest")
	proto.RegisterType((*OrdersRequest)(nil), "mookiespb.OrdersRequest")
	proto.RegisterType((*OrdersResponse)(nil), "mookiespb.OrdersResponse")
	proto.RegisterType((*Empty)(nil), "mookiespb.Empty")
	proto.RegisterType((*DeleteMenuItemRequest)(nil), "mookiespb.DeleteMenuItemRequest")
}

func init() { proto.RegisterFile("protofiles/mookies.proto", fileDescriptor_a42f9bdfdd6c12cf) }

var fileDescriptor_a42f9bdfdd6c12cf = []byte{
	// 687 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcf, 0x4e, 0xdb, 0x4e,
	0x10, 0xc6, 0xce, 0x5f, 0x26, 0x24, 0xc0, 0xfe, 0x02, 0x5a, 0xa2, 0x9f, 0x4a, 0xd8, 0xaa, 0x34,
	0xa5, 0x12, 0x69, 0x41, 0xea, 0x01, 0xd4, 0x4a, 0x15, 0x54, 0x88, 0x43, 0x45, 0x65, 0xca, 0x15,
	0xe4, 0x24, 0x53, 0xb4, 0x6a, 0x9c, 0x75, 0xbd, 0x1b, 0x24, 0xae, 0xbd, 0xf6, 0xd8, 0x07, 0xe9,
	0x2b, 0xf4, 0xd8, 0x7b, 0x5f, 0xa1, 0x0f, 0x52, 0x79, 0x6c, 0xa7, 0x9b, 0xc4, 0x54, 0x48, 0xbd,
	0x79, 0x66, 0xbf, 0xf9, 0xbe, 0x6f, 0x67, 0xc6, 0x36, 0xf0, 0x30, 0x52, 0x46, 0x7d, 0x90, 0x43,
	0xd4, 0xdd, 0x40, 0xa9, 0x8f, 0x12, 0xf5, 0x2e, 0xa5, 0xd8, 0x62, 0x1a, 0x86, 0xbd, 0xd6, 0xff,
	0xd7, 0x4a, 0x5d, 0x0f, 0xb1, 0xeb, 0x87, 0xb2, 0xeb, 0x8f, 0x46, 0xca, 0xf8, 0x46, 0xaa, 0x51,
	0x0a, 0x14, 0xdf, 0x1c, 0x28, 0x9e, 0x1a, 0x0c, 0x18, 0x83, 0xe2, 0xc8, 0x0f, 0x90, 0x3b, 0x6d,
	0xa7, 0xb3, 0xe8, 0xd1, 0x33, 0x6b, 0x80, 0x2b, 0x07, 0xdc, 0x6d, 0x3b, 0x9d, 0x82, 0xe7, 0xca,
	0x01, 0x6b, 0x42, 0x29, 0x8c, 0x64, 0x1f, 0x79, 0xa1, 0xed, 0x74, 0x5c, 0x2f, 0x09, 0xd8, 0x03,
	0x80, 0xbe, 0x6f, 0xf0, 0x5a, 0x45, 0xb7, 0xa7, 0xc7, 0xbc, 0x48, 0x68, 0x2b, 0xc3, 0xda, 0x50,
	0x53, 0xd1, 0x00, 0xa3, 0x58, 0xe6, 0xf4, 0x98, 0x97, 0x08, 0x60, 0xa7, 0xd8, 0x53, 0xa8, 0xa8,
	0x90, 0x5c, 0xf1, 0x72, 0xbb, 0xd0, 0xa9, 0xed, 0xad, 0xee, 0x4e, 0xfc, 0xef, 0x9e, 0xd1, 0x89,
	0x97, 0x21, 0xc4, 0x25, 0x94, 0x93, 0xd4, 0x3f, 0x58, 0x6e, 0x41, 0x55, 0xe3, 0x10, 0xfb, 0x06,
	0x07, 0x64, 0xb8, 0xea, 0x4d, 0x62, 0x71, 0x01, 0xd5, 0xa3, 0xd4, 0xfc, 0xbd, 0x14, 0x1e, 0x41,
	0x49, 0x1a, 0x0c, 0x34, 0x2f, 0x90, 0xf5, 0x65, 0xcb, 0x7a, 0x7c, 0x3d, 0x2f, 0x39, 0x15, 0x87,
	0x50, 0x7c, 0x8b, 0xa3, 0x31, 0xdb, 0x9f, 0x74, 0x4b, 0xa2, 0xe6, 0x0e, 0xd5, 0xfc, 0x67, 0xd5,
	0x64, 0xda, 0x9e, 0x05, 0x13, 0x3f, 0x1c, 0x28, 0x9d, 0xc5, 0x0d, 0x4b, 0xd5, 0x9d, 0x89, 0x7a,
	0xe6, 0xd0, 0xb5, 0x1c, 0xde, 0xcf, 0x51, 0xdc, 0x1a, 0xa3, 0x8c, 0x3f, 0xa4, 0x0e, 0xb8, 0x5e,
	0x12, 0xb0, 0x75, 0x28, 0x6b, 0xe3, 0x9b, 0xb1, 0xa6, 0x41, 0x2d, 0x7a, 0x69, 0xc4, 0xb6, 0x60,
	0xc9, 0xc8, 0x00, 0xaf, 0x68, 0x6e, 0x38, 0xe0, 0x65, 0x3a, 0xad, 0xc5, 0xb9, 0xb3, 0x24, 0xc5,
	0x1e, 0x42, 0x9d, 0x20, 0x7d, 0x15, 0x84, 0x43, 0x34, 0xc8, 0x2b, 0x84, 0xa1, 0xba, 0xa3, 0x34,
	0x27, 0xb6, 0xa1, 0xea, 0xa1, 0x0e, 0xd5, 0x48, 0xd3, 0x18, 0xa2, 0xf4, 0x39, 0x6d, 0xf1, 0x24,
	0x16, 0x1d, 0x58, 0x3f, 0x8a, 0xd0, 0x37, 0x18, 0x77, 0x8d, 0x6c, 0x67, 0x55, 0x33, 0x2d, 0x10,
	0xdb, 0xd0, 0xcc, 0xd8, 0xc9, 0x89, 0x87, 0x9f, 0xc6, 0xa8, 0xcd, 0x1c, 0xee, 0x09, 0xd4, 0xe9,
	0x5c, 0x67, 0x00, 0x0e, 0x95, 0x28, 0x79, 0x4c, 0xd5, 0xb3, 0x50, 0x1c, 0x40, 0x23, 0x83, 0xa6,
	0xa2, 0x1d, 0x28, 0xd3, 0xcd, 0xb3, 0x91, 0xad, 0xd8, 0x1b, 0x4a, 0xaa, 0xe9, 0xb9, 0xa8, 0x40,
	0xe9, 0x4d, 0x10, 0x9a, 0x5b, 0xf1, 0x18, 0xd6, 0x8e, 0x31, 0x76, 0xf5, 0xe7, 0x06, 0xb9, 0xc6,
	0xf6, 0xbe, 0x17, 0xa0, 0x16, 0x63, 0xce, 0x31, 0xba, 0x89, 0xb7, 0xf3, 0x15, 0x54, 0x4e, 0xd0,
	0xd0, 0xb6, 0xd8, 0x32, 0xc4, 0xda, 0xb2, 0xa7, 0x19, 0x43, 0xc4, 0xca, 0xe7, 0x9f, 0xbf, 0xbe,
	0xba, 0xc0, 0xaa, 0xdd, 0x9b, 0xe7, 0xdd, 0x20, 0x2e, 0xba, 0x84, 0xc6, 0x74, 0xeb, 0xd8, 0xec,
	0x0a, 0xb4, 0xb6, 0xec, 0x8d, 0xcb, 0x6d, 0xb3, 0xe0, 0xc4, 0xcb, 0x44, 0x3d, 0xe3, 0xed, 0xc6,
	0x6b, 0x73, 0xe0, 0xec, 0xb0, 0x77, 0xd0, 0xb8, 0x08, 0x07, 0x7f, 0xe5, 0xb7, 0x37, 0x7a, 0x96,
	0xb1, 0x35, 0xcf, 0x78, 0x09, 0x8d, 0xe9, 0x56, 0xb1, 0xb6, 0x45, 0x90, 0xdb, 0xc5, 0x7c, 0x89,
	0x35, 0x92, 0x58, 0xde, 0x99, 0x96, 0x60, 0x57, 0xd0, 0x9c, 0xbe, 0x65, 0xfa, 0x05, 0x99, 0xff,
	0xce, 0xe4, 0xd3, 0x6e, 0x12, 0xed, 0x86, 0x68, 0x4e, 0xd1, 0x76, 0x93, 0x4f, 0xd2, 0x81, 0xb3,
	0xb3, 0xf7, 0xc5, 0x85, 0x25, 0x5a, 0x83, 0x6c, 0x86, 0x2f, 0xa0, 0x76, 0x3e, 0xee, 0x05, 0xd2,
	0x24, 0xaf, 0xed, 0xdc, 0xba, 0xe4, 0xeb, 0x2c, 0xb0, 0x97, 0xb0, 0xf4, 0xba, 0x6f, 0xe4, 0x4d,
	0xb2, 0xca, 0x3a, 0x67, 0x01, 0x36, 0x66, 0xa9, 0xb4, 0x55, 0x7e, 0x02, 0xf5, 0xa9, 0x77, 0x81,
	0x6d, 0xda, 0x83, 0xce, 0x79, 0x4b, 0xee, 0xf2, 0x71, 0x08, 0xab, 0xe7, 0xe3, 0x9e, 0xee, 0x47,
	0xb2, 0x87, 0xef, 0xd5, 0x9d, 0x66, 0xe6, 0xee, 0x25, 0x16, 0x9e, 0x39, 0xbd, 0x32, 0xfd, 0x5b,
	0xf6, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x77, 0x31, 0xb5, 0x9e, 0xa0, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MenuServiceClient is the client API for MenuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MenuServiceClient interface {
	GetMenu(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Menu, error)
	CreateMenuItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*CreateMenuItemResponse, error)
	UpdateMenuItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Response, error)
	DeleteMenuItem(ctx context.Context, in *DeleteMenuItemRequest, opts ...grpc.CallOption) (*Response, error)
	CreateMenuItemOption(ctx context.Context, in *Option, opts ...grpc.CallOption) (*Response, error)
}

type menuServiceClient struct {
	cc *grpc.ClientConn
}

func NewMenuServiceClient(cc *grpc.ClientConn) MenuServiceClient {
	return &menuServiceClient{cc}
}

func (c *menuServiceClient) GetMenu(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Menu, error) {
	out := new(Menu)
	err := c.cc.Invoke(ctx, "/mookiespb.MenuService/GetMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) CreateMenuItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*CreateMenuItemResponse, error) {
	out := new(CreateMenuItemResponse)
	err := c.cc.Invoke(ctx, "/mookiespb.MenuService/CreateMenuItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) UpdateMenuItem(ctx context.Context, in *Item, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/mookiespb.MenuService/UpdateMenuItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) DeleteMenuItem(ctx context.Context, in *DeleteMenuItemRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/mookiespb.MenuService/DeleteMenuItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *menuServiceClient) CreateMenuItemOption(ctx context.Context, in *Option, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/mookiespb.MenuService/CreateMenuItemOption", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MenuServiceServer is the server API for MenuService service.
type MenuServiceServer interface {
	GetMenu(context.Context, *Empty) (*Menu, error)
	CreateMenuItem(context.Context, *Item) (*CreateMenuItemResponse, error)
	UpdateMenuItem(context.Context, *Item) (*Response, error)
	DeleteMenuItem(context.Context, *DeleteMenuItemRequest) (*Response, error)
	CreateMenuItemOption(context.Context, *Option) (*Response, error)
}

// UnimplementedMenuServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMenuServiceServer struct {
}

func (*UnimplementedMenuServiceServer) GetMenu(ctx context.Context, req *Empty) (*Menu, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMenu not implemented")
}
func (*UnimplementedMenuServiceServer) CreateMenuItem(ctx context.Context, req *Item) (*CreateMenuItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMenuItem not implemented")
}
func (*UnimplementedMenuServiceServer) UpdateMenuItem(ctx context.Context, req *Item) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMenuItem not implemented")
}
func (*UnimplementedMenuServiceServer) DeleteMenuItem(ctx context.Context, req *DeleteMenuItemRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMenuItem not implemented")
}
func (*UnimplementedMenuServiceServer) CreateMenuItemOption(ctx context.Context, req *Option) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMenuItemOption not implemented")
}

func RegisterMenuServiceServer(s *grpc.Server, srv MenuServiceServer) {
	s.RegisterService(&_MenuService_serviceDesc, srv)
}

func _MenuService_GetMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).GetMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mookiespb.MenuService/GetMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).GetMenu(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_CreateMenuItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).CreateMenuItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mookiespb.MenuService/CreateMenuItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).CreateMenuItem(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_UpdateMenuItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Item)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).UpdateMenuItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mookiespb.MenuService/UpdateMenuItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).UpdateMenuItem(ctx, req.(*Item))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_DeleteMenuItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMenuItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).DeleteMenuItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mookiespb.MenuService/DeleteMenuItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).DeleteMenuItem(ctx, req.(*DeleteMenuItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MenuService_CreateMenuItemOption_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Option)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MenuServiceServer).CreateMenuItemOption(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mookiespb.MenuService/CreateMenuItemOption",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MenuServiceServer).CreateMenuItemOption(ctx, req.(*Option))
	}
	return interceptor(ctx, in, info, handler)
}

var _MenuService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mookiespb.MenuService",
	HandlerType: (*MenuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMenu",
			Handler:    _MenuService_GetMenu_Handler,
		},
		{
			MethodName: "CreateMenuItem",
			Handler:    _MenuService_CreateMenuItem_Handler,
		},
		{
			MethodName: "UpdateMenuItem",
			Handler:    _MenuService_UpdateMenuItem_Handler,
		},
		{
			MethodName: "DeleteMenuItem",
			Handler:    _MenuService_DeleteMenuItem_Handler,
		},
		{
			MethodName: "CreateMenuItemOption",
			Handler:    _MenuService_CreateMenuItemOption_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protofiles/mookies.proto",
}

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	// Unary
	SubmitOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Response, error)
	ActiveOrders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*OrdersResponse, error)
	CompleteOrder(ctx context.Context, in *CompleteOrderRequest, opts ...grpc.CallOption) (*Response, error)
	// server streaming
	SubscribeToOrders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (OrderService_SubscribeToOrdersClient, error)
}

type orderServiceClient struct {
	cc *grpc.ClientConn
}

func NewOrderServiceClient(cc *grpc.ClientConn) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) SubmitOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/mookiespb.OrderService/SubmitOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ActiveOrders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*OrdersResponse, error) {
	out := new(OrdersResponse)
	err := c.cc.Invoke(ctx, "/mookiespb.OrderService/ActiveOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CompleteOrder(ctx context.Context, in *CompleteOrderRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/mookiespb.OrderService/CompleteOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) SubscribeToOrders(ctx context.Context, in *Empty, opts ...grpc.CallOption) (OrderService_SubscribeToOrdersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OrderService_serviceDesc.Streams[0], "/mookiespb.OrderService/SubscribeToOrders", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderServiceSubscribeToOrdersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type OrderService_SubscribeToOrdersClient interface {
	Recv() (*Order, error)
	grpc.ClientStream
}

type orderServiceSubscribeToOrdersClient struct {
	grpc.ClientStream
}

func (x *orderServiceSubscribeToOrdersClient) Recv() (*Order, error) {
	m := new(Order)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	// Unary
	SubmitOrder(context.Context, *Order) (*Response, error)
	ActiveOrders(context.Context, *Empty) (*OrdersResponse, error)
	CompleteOrder(context.Context, *CompleteOrderRequest) (*Response, error)
	// server streaming
	SubscribeToOrders(*Empty, OrderService_SubscribeToOrdersServer) error
}

// UnimplementedOrderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (*UnimplementedOrderServiceServer) SubmitOrder(ctx context.Context, req *Order) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitOrder not implemented")
}
func (*UnimplementedOrderServiceServer) ActiveOrders(ctx context.Context, req *Empty) (*OrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActiveOrders not implemented")
}
func (*UnimplementedOrderServiceServer) CompleteOrder(ctx context.Context, req *CompleteOrderRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteOrder not implemented")
}
func (*UnimplementedOrderServiceServer) SubscribeToOrders(req *Empty, srv OrderService_SubscribeToOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToOrders not implemented")
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_SubmitOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).SubmitOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mookiespb.OrderService/SubmitOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).SubmitOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ActiveOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ActiveOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mookiespb.OrderService/ActiveOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ActiveOrders(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CompleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CompleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mookiespb.OrderService/CompleteOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CompleteOrder(ctx, req.(*CompleteOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_SubscribeToOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OrderServiceServer).SubscribeToOrders(m, &orderServiceSubscribeToOrdersServer{stream})
}

type OrderService_SubscribeToOrdersServer interface {
	Send(*Order) error
	grpc.ServerStream
}

type orderServiceSubscribeToOrdersServer struct {
	grpc.ServerStream
}

func (x *orderServiceSubscribeToOrdersServer) Send(m *Order) error {
	return x.ServerStream.SendMsg(m)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mookiespb.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitOrder",
			Handler:    _OrderService_SubmitOrder_Handler,
		},
		{
			MethodName: "ActiveOrders",
			Handler:    _OrderService_ActiveOrders_Handler,
		},
		{
			MethodName: "CompleteOrder",
			Handler:    _OrderService_CompleteOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToOrders",
			Handler:       _OrderService_SubscribeToOrders_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protofiles/mookies.proto",
}
