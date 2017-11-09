// Code generated by protoc-gen-go. DO NOT EDIT.
// source: synse.proto

/*
Package synse is a generated protocol buffer package.

It is generated from these files:
	synse.proto

It has these top-level messages:
	ReadRequest
	WriteRequest
	MetainfoRequest
	TransactionId
	ReadResponse
	Transactions
	MetainfoResponse
	WriteResponse
	WriteData
	MetaOutputUnit
	MetaOutputRange
	MetaOutput
	MetaLocation
*/
package synse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type WriteResponse_WriteStatus int32

const (
	WriteResponse_UNKNOWN WriteResponse_WriteStatus = 0
	WriteResponse_PENDING WriteResponse_WriteStatus = 1
	WriteResponse_WRITING WriteResponse_WriteStatus = 2
	WriteResponse_DONE    WriteResponse_WriteStatus = 3
)

var WriteResponse_WriteStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "PENDING",
	2: "WRITING",
	3: "DONE",
}
var WriteResponse_WriteStatus_value = map[string]int32{
	"UNKNOWN": 0,
	"PENDING": 1,
	"WRITING": 2,
	"DONE":    3,
}

func (x WriteResponse_WriteStatus) String() string {
	return proto.EnumName(WriteResponse_WriteStatus_name, int32(x))
}
func (WriteResponse_WriteStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

type WriteResponse_WriteState int32

const (
	WriteResponse_OK    WriteResponse_WriteState = 0
	WriteResponse_ERROR WriteResponse_WriteState = 1
)

var WriteResponse_WriteState_name = map[int32]string{
	0: "OK",
	1: "ERROR",
}
var WriteResponse_WriteState_value = map[string]int32{
	"OK":    0,
	"ERROR": 1,
}

func (x WriteResponse_WriteState) String() string {
	return proto.EnumName(WriteResponse_WriteState_name, int32(x))
}
func (WriteResponse_WriteState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 1} }

// Read
// ~~~~
// the read request message contains the uuid of the device that
// we desire to read. the uuid of the device should be generated
// by the owning background process and should be returned to the
// synse application in the MetainfoResponse, which Synse will
// cache and use as a lookup table for routing requests.
type ReadRequest struct {
	// the id of the device to read. this is generated by the plugin
	// and returned via the `Metainfo` request.
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReadRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

// Write
// ~~~~~
// the write request message contains the uuid of the device that
// we desire to write to, as well as a repeated string (e.g. a
// list of strings in Python) which makes up the data that we
// which to write to that device.
type WriteRequest struct {
	// the id of the device to write to. this is generated by the
	// plugin and returned via the `Metainfo` request.
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	// the data to write. a given synse-server write request could
	// actually be a composite of writes. for example, one can turn
	// an LED on and change its color simultaneously via the Synse
	// JSON API. each `WriteData` will get its own transaction id.
	Data []*WriteData `protobuf:"bytes,2,rep,name=data" json:"data,omitempty"`
}

func (m *WriteRequest) Reset()                    { *m = WriteRequest{} }
func (m *WriteRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()               {}
func (*WriteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *WriteRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *WriteRequest) GetData() []*WriteData {
	if m != nil {
		return m.Data
	}
	return nil
}

// Metainfo
// ~~~~~~~~
// the metainfo request message contains a field for rack and board,
// but neither are required. if specified, the response will contain
// only information relating to the rack/board filter applied. if
// they are left unspecified, the response will contain the entirety
// of the metainfo scan information.
type MetainfoRequest struct {
	// the rack filter for the meta information response.
	Rack string `protobuf:"bytes,1,opt,name=rack" json:"rack,omitempty"`
	// the board filter for the meta information response.
	Board string `protobuf:"bytes,2,opt,name=board" json:"board,omitempty"`
}

func (m *MetainfoRequest) Reset()                    { *m = MetainfoRequest{} }
func (m *MetainfoRequest) String() string            { return proto.CompactTextString(m) }
func (*MetainfoRequest) ProtoMessage()               {}
func (*MetainfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *MetainfoRequest) GetRack() string {
	if m != nil {
		return m.Rack
	}
	return ""
}

func (m *MetainfoRequest) GetBoard() string {
	if m != nil {
		return m.Board
	}
	return ""
}

// TransactionCheck
// ~~~~~~~~~~~~~~~~
// the transaction id gives identity to a single 'write' action. since
// device writes are handled asynchronously, the background process
// returns the transaction id when a write is registered, which the
// caller can later pass back to `TransactionCheck` to get the status
// of that write.
type TransactionId struct {
	// the id of a write transaction. this is returned by the write
	// commands `Transactions` response.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *TransactionId) Reset()                    { *m = TransactionId{} }
func (m *TransactionId) String() string            { return proto.CompactTextString(m) }
func (*TransactionId) ProtoMessage()               {}
func (*TransactionId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TransactionId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Read
// ~~~~
// the read response provides the timestamp at which the reading was
// taken, the type of the reading (e.g. temperature, humidity, led
// state, etc.), and the value of that reading. read responses are
// returned to the client as a stream, so a single device can return
// multiple readings. (e.g. a humidity sensor can return a %humidity
// reading and a temperature reading).
type ReadResponse struct {
	// the time which the reading was taken.
	Timestamp string `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	// the type of reading.
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// the value of the reading.
	Value string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ReadResponse) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *ReadResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ReadResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Write
// ~~~~~
// the transactions message specifies the asynchronous transactions for
// each of the given write actions. each transaction identifies a single
// write action with a unique transaction id and context to help identify
// which transaction that id is associated with. the transaction id can
// later be passed back to `TransactionCheck` to get the status of that
// write.
type Transactions struct {
	// a map where the key is the transaction id for a given `WriteData`
	// message that has been processed, and the value is that same
	// `WriteData` message. the `WriteData` message is passed back in
	// order to provide some context and make identifying transactions
	// possible. the number of entries in the transactions map corresponds
	// to the number of `WriteData` recieved in a `WriteRequest`.
	Transactions map[string]*WriteData `protobuf:"bytes,1,rep,name=transactions" json:"transactions,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Transactions) Reset()                    { *m = Transactions{} }
func (m *Transactions) String() string            { return proto.CompactTextString(m) }
func (*Transactions) ProtoMessage()               {}
func (*Transactions) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Transactions) GetTransactions() map[string]*WriteData {
	if m != nil {
		return m.Transactions
	}
	return nil
}

// Metainfo
// ~~~~~~~~
// the metainfo response represents a single device that is owned by
// the process. metainfo responses are returned to the client as a stream
// so a background process can support any number of devices. the response
// itself contains a timestamp for when the response was generated, an
// for the device, and all other meta-information we have pertaining to
// that device. the caller, Synse, will cache this information and use it
// to route requests to the appropriate device as well as provide responses
// for scan and info requests.
type MetainfoResponse struct {
	// the time at which the metainfo was gathered.
	Timestamp string `protobuf:"bytes,1,opt,name=timestamp" json:"timestamp,omitempty"`
	// the unique id for the device this response represents.
	Uid string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
	// the device type.
	Type string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	// the device model.
	Model string `protobuf:"bytes,4,opt,name=model" json:"model,omitempty"`
	// the device manufacturer.
	Manufacturer string `protobuf:"bytes,5,opt,name=manufacturer" json:"manufacturer,omitempty"`
	// the protocol that the device is configured to use.
	Protocol string `protobuf:"bytes,6,opt,name=protocol" json:"protocol,omitempty"`
	// any additional information specified for the device.
	Info string `protobuf:"bytes,7,opt,name=info" json:"info,omitempty"`
	// any comment specified for the device.
	Comment string `protobuf:"bytes,8,opt,name=comment" json:"comment,omitempty"`
	// the location of the device, as specified by rack and board
	// identifiers.
	Location *MetaLocation `protobuf:"bytes,9,opt,name=location" json:"location,omitempty"`
	// the reading output of the device. this specifies all of the
	// outputs a device will generate when read. most devices will have
	// a single output, but some devices (e.g. a humidity sensor) could
	// return multiple data points from a single reading. (e.g.
	// temperature and humidity).
	Output []*MetaOutput `protobuf:"bytes,10,rep,name=output" json:"output,omitempty"`
}

func (m *MetainfoResponse) Reset()                    { *m = MetainfoResponse{} }
func (m *MetainfoResponse) String() string            { return proto.CompactTextString(m) }
func (*MetainfoResponse) ProtoMessage()               {}
func (*MetainfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *MetainfoResponse) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *MetainfoResponse) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *MetainfoResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MetainfoResponse) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *MetainfoResponse) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *MetainfoResponse) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *MetainfoResponse) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *MetainfoResponse) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *MetainfoResponse) GetLocation() *MetaLocation {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *MetainfoResponse) GetOutput() []*MetaOutput {
	if m != nil {
		return m.Output
	}
	return nil
}

// TransactionCheck
// ~~~~~~~~~~~~~~~~
// the response for a transaction check command gives the status of the
// transaction. transactions correspond to write requests. since writes
// are performed asynchronously, the transaction id is used to track the
// progress of that transaction.
type WriteResponse struct {
	// the time at which the write transaction was created.
	Created string `protobuf:"bytes,1,opt,name=created" json:"created,omitempty"`
	// the time at which the write transaction was last updated (either
	// for an update to state or status).
	Updated string `protobuf:"bytes,2,opt,name=updated" json:"updated,omitempty"`
	// the status of the transaction. this describes what stage of
	// processing the transaction is at.
	Status WriteResponse_WriteStatus `protobuf:"varint,3,opt,name=status,enum=synse.WriteResponse_WriteStatus" json:"status,omitempty"`
	// the state of the transaction. this describes the so called "health"
	// of the transaction.
	State WriteResponse_WriteState `protobuf:"varint,4,opt,name=state,enum=synse.WriteResponse_WriteState" json:"state,omitempty"`
	// the message field will be used to specify any context information
	// when the state is in ERROR. if the state is OK, this field will
	// remain empty.
	Message string `protobuf:"bytes,5,opt,name=message" json:"message,omitempty"`
}

func (m *WriteResponse) Reset()                    { *m = WriteResponse{} }
func (m *WriteResponse) String() string            { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()               {}
func (*WriteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *WriteResponse) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *WriteResponse) GetUpdated() string {
	if m != nil {
		return m.Updated
	}
	return ""
}

func (m *WriteResponse) GetStatus() WriteResponse_WriteStatus {
	if m != nil {
		return m.Status
	}
	return WriteResponse_UNKNOWN
}

func (m *WriteResponse) GetState() WriteResponse_WriteState {
	if m != nil {
		return m.State
	}
	return WriteResponse_OK
}

func (m *WriteResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// Specifies the data that gets written on a `WriteRequest`. This is a
// composite of raw bytes and an action string. Both of the fields can
// be specified, or only one of them. In cases where no action is supplied,
// the plugin can take the raw data as raw bytes to write to the device.
// There may be cases when only an action is supplied, such as "on" for
// turning an LED on (which the plugin can then interpret for its protocol).
// Both can be specified, such as changing the color of an LED where the
// raw bytes can specify the color itself and the action could be "color",
// differentiating it from a write to LED blink state, etc.
type WriteData struct {
	// raw bytes to send to the plugin for writing.
	Raw [][]byte `protobuf:"bytes,1,rep,name=raw,proto3" json:"raw,omitempty"`
	// a (well-known) action identifier.
	Action string `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
}

func (m *WriteData) Reset()                    { *m = WriteData{} }
func (m *WriteData) String() string            { return proto.CompactTextString(m) }
func (*WriteData) ProtoMessage()               {}
func (*WriteData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *WriteData) GetRaw() [][]byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *WriteData) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

// the unit specification for a reading output.
type MetaOutputUnit struct {
	// the full name of the unit, e.g. "degrees celsius".
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// the symbol for the unit, e.g. "C".
	Symbol string `protobuf:"bytes,2,opt,name=symbol" json:"symbol,omitempty"`
}

func (m *MetaOutputUnit) Reset()                    { *m = MetaOutputUnit{} }
func (m *MetaOutputUnit) String() string            { return proto.CompactTextString(m) }
func (*MetaOutputUnit) ProtoMessage()               {}
func (*MetaOutputUnit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *MetaOutputUnit) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetaOutputUnit) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

// the value range specification for a reading output.
type MetaOutputRange struct {
	// the minimum allowable value for a numeric reading.
	Min int32 `protobuf:"varint,1,opt,name=min" json:"min,omitempty"`
	// the maximum allowable value for a numeric reading.
	Max int32 `protobuf:"varint,2,opt,name=max" json:"max,omitempty"`
}

func (m *MetaOutputRange) Reset()                    { *m = MetaOutputRange{} }
func (m *MetaOutputRange) String() string            { return proto.CompactTextString(m) }
func (*MetaOutputRange) ProtoMessage()               {}
func (*MetaOutputRange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *MetaOutputRange) GetMin() int32 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *MetaOutputRange) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

// the specification for one of a device's reading outputs.
type MetaOutput struct {
	// the type of the reading output (e.g. 'temperature',
	// 'humidity', etc).
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	// the data type of the output (e.g. int, string, bool).
	DataType string `protobuf:"bytes,2,opt,name=data_type,json=dataType" json:"data_type,omitempty"`
	// the decimal precision of the output. if the output is
	// non-numeric, this can be left unspecified and will be
	// ignored.
	Precision int32 `protobuf:"varint,3,opt,name=precision" json:"precision,omitempty"`
	// the unit of measure for the reading.
	Unit *MetaOutputUnit `protobuf:"bytes,4,opt,name=unit" json:"unit,omitempty"`
	// the acceptable range of values for the reading.
	Range *MetaOutputRange `protobuf:"bytes,5,opt,name=range" json:"range,omitempty"`
}

func (m *MetaOutput) Reset()                    { *m = MetaOutput{} }
func (m *MetaOutput) String() string            { return proto.CompactTextString(m) }
func (*MetaOutput) ProtoMessage()               {}
func (*MetaOutput) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *MetaOutput) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MetaOutput) GetDataType() string {
	if m != nil {
		return m.DataType
	}
	return ""
}

func (m *MetaOutput) GetPrecision() int32 {
	if m != nil {
		return m.Precision
	}
	return 0
}

func (m *MetaOutput) GetUnit() *MetaOutputUnit {
	if m != nil {
		return m.Unit
	}
	return nil
}

func (m *MetaOutput) GetRange() *MetaOutputRange {
	if m != nil {
		return m.Range
	}
	return nil
}

// the location specification for a device.
type MetaLocation struct {
	// the rack which the device belongs to.
	Rack string `protobuf:"bytes,1,opt,name=rack" json:"rack,omitempty"`
	// the board which the device belongs to.
	Board string `protobuf:"bytes,2,opt,name=board" json:"board,omitempty"`
}

func (m *MetaLocation) Reset()                    { *m = MetaLocation{} }
func (m *MetaLocation) String() string            { return proto.CompactTextString(m) }
func (*MetaLocation) ProtoMessage()               {}
func (*MetaLocation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *MetaLocation) GetRack() string {
	if m != nil {
		return m.Rack
	}
	return ""
}

func (m *MetaLocation) GetBoard() string {
	if m != nil {
		return m.Board
	}
	return ""
}

func init() {
	proto.RegisterType((*ReadRequest)(nil), "synse.ReadRequest")
	proto.RegisterType((*WriteRequest)(nil), "synse.WriteRequest")
	proto.RegisterType((*MetainfoRequest)(nil), "synse.MetainfoRequest")
	proto.RegisterType((*TransactionId)(nil), "synse.TransactionId")
	proto.RegisterType((*ReadResponse)(nil), "synse.ReadResponse")
	proto.RegisterType((*Transactions)(nil), "synse.Transactions")
	proto.RegisterType((*MetainfoResponse)(nil), "synse.MetainfoResponse")
	proto.RegisterType((*WriteResponse)(nil), "synse.WriteResponse")
	proto.RegisterType((*WriteData)(nil), "synse.WriteData")
	proto.RegisterType((*MetaOutputUnit)(nil), "synse.MetaOutputUnit")
	proto.RegisterType((*MetaOutputRange)(nil), "synse.MetaOutputRange")
	proto.RegisterType((*MetaOutput)(nil), "synse.MetaOutput")
	proto.RegisterType((*MetaLocation)(nil), "synse.MetaLocation")
	proto.RegisterEnum("synse.WriteResponse_WriteStatus", WriteResponse_WriteStatus_name, WriteResponse_WriteStatus_value)
	proto.RegisterEnum("synse.WriteResponse_WriteState", WriteResponse_WriteState_name, WriteResponse_WriteState_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for InternalApi service

type InternalApiClient interface {
	// Read from the specified device(s).
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (InternalApi_ReadClient, error)
	// Write to the specified device(s).
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*Transactions, error)
	// Get the metainformation from the background process that describes
	// all of the available devices which that process owns
	Metainfo(ctx context.Context, in *MetainfoRequest, opts ...grpc.CallOption) (InternalApi_MetainfoClient, error)
	// Check on the state of a write transaction.
	TransactionCheck(ctx context.Context, in *TransactionId, opts ...grpc.CallOption) (*WriteResponse, error)
}

type internalApiClient struct {
	cc *grpc.ClientConn
}

func NewInternalApiClient(cc *grpc.ClientConn) InternalApiClient {
	return &internalApiClient{cc}
}

func (c *internalApiClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (InternalApi_ReadClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_InternalApi_serviceDesc.Streams[0], c.cc, "/synse.InternalApi/Read", opts...)
	if err != nil {
		return nil, err
	}
	x := &internalApiReadClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InternalApi_ReadClient interface {
	Recv() (*ReadResponse, error)
	grpc.ClientStream
}

type internalApiReadClient struct {
	grpc.ClientStream
}

func (x *internalApiReadClient) Recv() (*ReadResponse, error) {
	m := new(ReadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *internalApiClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*Transactions, error) {
	out := new(Transactions)
	err := grpc.Invoke(ctx, "/synse.InternalApi/Write", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalApiClient) Metainfo(ctx context.Context, in *MetainfoRequest, opts ...grpc.CallOption) (InternalApi_MetainfoClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_InternalApi_serviceDesc.Streams[1], c.cc, "/synse.InternalApi/Metainfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &internalApiMetainfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type InternalApi_MetainfoClient interface {
	Recv() (*MetainfoResponse, error)
	grpc.ClientStream
}

type internalApiMetainfoClient struct {
	grpc.ClientStream
}

func (x *internalApiMetainfoClient) Recv() (*MetainfoResponse, error) {
	m := new(MetainfoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *internalApiClient) TransactionCheck(ctx context.Context, in *TransactionId, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := grpc.Invoke(ctx, "/synse.InternalApi/TransactionCheck", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InternalApi service

type InternalApiServer interface {
	// Read from the specified device(s).
	Read(*ReadRequest, InternalApi_ReadServer) error
	// Write to the specified device(s).
	Write(context.Context, *WriteRequest) (*Transactions, error)
	// Get the metainformation from the background process that describes
	// all of the available devices which that process owns
	Metainfo(*MetainfoRequest, InternalApi_MetainfoServer) error
	// Check on the state of a write transaction.
	TransactionCheck(context.Context, *TransactionId) (*WriteResponse, error)
}

func RegisterInternalApiServer(s *grpc.Server, srv InternalApiServer) {
	s.RegisterService(&_InternalApi_serviceDesc, srv)
}

func _InternalApi_Read_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InternalApiServer).Read(m, &internalApiReadServer{stream})
}

type InternalApi_ReadServer interface {
	Send(*ReadResponse) error
	grpc.ServerStream
}

type internalApiReadServer struct {
	grpc.ServerStream
}

func (x *internalApiReadServer) Send(m *ReadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _InternalApi_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalApiServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/synse.InternalApi/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalApiServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalApi_Metainfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MetainfoRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(InternalApiServer).Metainfo(m, &internalApiMetainfoServer{stream})
}

type InternalApi_MetainfoServer interface {
	Send(*MetainfoResponse) error
	grpc.ServerStream
}

type internalApiMetainfoServer struct {
	grpc.ServerStream
}

func (x *internalApiMetainfoServer) Send(m *MetainfoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _InternalApi_TransactionCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalApiServer).TransactionCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/synse.InternalApi/TransactionCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalApiServer).TransactionCheck(ctx, req.(*TransactionId))
	}
	return interceptor(ctx, in, info, handler)
}

var _InternalApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "synse.InternalApi",
	HandlerType: (*InternalApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _InternalApi_Write_Handler,
		},
		{
			MethodName: "TransactionCheck",
			Handler:    _InternalApi_TransactionCheck_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Read",
			Handler:       _InternalApi_Read_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Metainfo",
			Handler:       _InternalApi_Metainfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "synse.proto",
}

func init() { proto.RegisterFile("synse.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 799 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0x5b, 0x6f, 0xd3, 0x4a,
	0x10, 0x8e, 0x9d, 0x38, 0x4d, 0xc6, 0x69, 0x8e, 0xbb, 0xed, 0xe9, 0xb1, 0x72, 0x2a, 0xa5, 0x5a,
	0x9d, 0x83, 0x5a, 0x09, 0x15, 0x94, 0xaa, 0x52, 0x05, 0x88, 0x8b, 0x68, 0x40, 0x51, 0x21, 0x81,
	0xa5, 0xa5, 0x8f, 0x68, 0x1b, 0x6f, 0x8b, 0xd5, 0xf8, 0x82, 0xbd, 0x86, 0xe6, 0x3f, 0x21, 0xf1,
	0x82, 0xf8, 0x71, 0x3c, 0xa1, 0x59, 0xaf, 0x13, 0x47, 0x29, 0xa8, 0x6f, 0xf3, 0xcd, 0xc5, 0xdf,
	0xce, 0x37, 0xe3, 0x01, 0x3b, 0x9d, 0x86, 0xa9, 0xd8, 0x8b, 0x93, 0x48, 0x46, 0xc4, 0x52, 0x80,
	0x76, 0xc1, 0x66, 0x82, 0x7b, 0x4c, 0x7c, 0xca, 0x44, 0x2a, 0x89, 0x03, 0xd5, 0xcc, 0xf7, 0x5c,
	0x63, 0xdb, 0xd8, 0x69, 0x32, 0x34, 0xe9, 0x0b, 0x68, 0x9d, 0x25, 0xbe, 0x14, 0xbf, 0xcd, 0x20,
	0xff, 0x41, 0xcd, 0xe3, 0x92, 0xbb, 0xe6, 0x76, 0x75, 0xc7, 0xee, 0x39, 0x7b, 0x39, 0x8b, 0x2a,
	0x3a, 0xe2, 0x92, 0x33, 0x15, 0xa5, 0x0f, 0xe1, 0xaf, 0xd7, 0x42, 0x72, 0x3f, 0xbc, 0x88, 0x8a,
	0x4f, 0x11, 0xa8, 0x25, 0x7c, 0x7c, 0xa5, 0xbf, 0xa5, 0x6c, 0xb2, 0x01, 0xd6, 0x79, 0xc4, 0x13,
	0xcf, 0x35, 0x95, 0x33, 0x07, 0xb4, 0x0b, 0xab, 0x27, 0x09, 0x0f, 0x53, 0x3e, 0x96, 0x7e, 0x14,
	0x0e, 0x3c, 0xd2, 0x06, 0x73, 0xf6, 0x08, 0xd3, 0xf7, 0xe8, 0x7b, 0x68, 0xe5, 0x6d, 0xa4, 0x71,
	0x14, 0xa6, 0x82, 0x6c, 0x41, 0x53, 0xfa, 0x81, 0x48, 0x25, 0x0f, 0x62, 0x9d, 0x36, 0x77, 0x20,
	0xb1, 0x9c, 0xc6, 0x42, 0x73, 0x28, 0x1b, 0x89, 0x3f, 0xf3, 0x49, 0x26, 0xdc, 0x6a, 0x4e, 0xac,
	0x00, 0xfd, 0x6a, 0x40, 0xab, 0xc4, 0x9c, 0x92, 0x01, 0xb4, 0x64, 0x09, 0xbb, 0x86, 0x6a, 0xfa,
	0x7f, 0xdd, 0x74, 0x39, 0x75, 0x01, 0xf4, 0x43, 0x99, 0x4c, 0xd9, 0x42, 0x69, 0xe7, 0x2d, 0xac,
	0x2d, 0xa5, 0xa0, 0xbc, 0x57, 0x62, 0x5a, 0xc8, 0x7b, 0x25, 0xa6, 0xe4, 0x4e, 0xf1, 0x30, 0x7c,
	0xed, 0x4d, 0xfa, 0xe6, 0xe1, 0x07, 0xe6, 0xa1, 0x41, 0x7f, 0x98, 0xe0, 0xcc, 0x55, 0xbe, 0x95,
	0x16, 0x7a, 0x9e, 0xe6, 0x7c, 0x9e, 0x85, 0x3a, 0xd5, 0x45, 0x75, 0x82, 0xc8, 0x13, 0x13, 0xb7,
	0x96, 0xab, 0xa3, 0x00, 0xa1, 0xd0, 0x0a, 0x78, 0x98, 0x5d, 0xf0, 0xb1, 0xcc, 0x12, 0x91, 0xb8,
	0x96, 0x0a, 0x2e, 0xf8, 0x48, 0x07, 0x1a, 0x6a, 0xe1, 0xc6, 0xd1, 0xc4, 0xad, 0xab, 0xf8, 0x0c,
	0x23, 0x13, 0xbe, 0xd4, 0x5d, 0xc9, 0x99, 0xd0, 0x26, 0x2e, 0xac, 0x8c, 0xa3, 0x20, 0x10, 0xa1,
	0x74, 0x1b, 0xca, 0x5d, 0x40, 0x72, 0x0f, 0x1a, 0x93, 0x68, 0xcc, 0x51, 0x2c, 0xb7, 0xa9, 0xb4,
	0x58, 0xd7, 0x5a, 0x60, 0xcb, 0xaf, 0x74, 0x88, 0xcd, 0x92, 0xc8, 0x2e, 0xd4, 0xa3, 0x4c, 0xc6,
	0x99, 0x74, 0x41, 0x4d, 0x69, 0xad, 0x94, 0x3e, 0x52, 0x01, 0xa6, 0x13, 0xe8, 0x37, 0x13, 0x56,
	0xf5, 0x9a, 0x6b, 0xd5, 0xf0, 0x1d, 0x89, 0xe0, 0x52, 0x14, 0x6b, 0x56, 0x40, 0x8c, 0x64, 0xb1,
	0xa7, 0x22, 0xb9, 0x6a, 0x05, 0x24, 0x87, 0x50, 0x4f, 0x25, 0x97, 0x59, 0xaa, 0xb4, 0x6b, 0xf7,
	0xb6, 0xcb, 0xb3, 0x2a, 0xbe, 0x9c, 0xa3, 0x77, 0x2a, 0x8f, 0xe9, 0x7c, 0x72, 0x00, 0x16, 0x5a,
	0x42, 0xe9, 0xdb, 0xee, 0x75, 0xff, 0x5c, 0x28, 0x58, 0x9e, 0x8d, 0x4f, 0x09, 0x44, 0x9a, 0xf2,
	0x4b, 0xa1, 0xb5, 0x2f, 0x20, 0x7d, 0x0c, 0x76, 0x89, 0x87, 0xd8, 0xb0, 0x72, 0x3a, 0x3c, 0x1e,
	0x8e, 0xce, 0x86, 0x4e, 0x05, 0xc1, 0x9b, 0xfe, 0xf0, 0x68, 0x30, 0x7c, 0xe9, 0x18, 0x08, 0xce,
	0xd8, 0xe0, 0x04, 0x81, 0x49, 0x1a, 0x50, 0x3b, 0x1a, 0x0d, 0xfb, 0x4e, 0x95, 0x76, 0x01, 0xe6,
	0x74, 0xa4, 0x0e, 0xe6, 0xe8, 0xd8, 0xa9, 0x90, 0x26, 0x58, 0x7d, 0xc6, 0x46, 0xcc, 0x31, 0xe8,
	0x01, 0x34, 0x67, 0x2b, 0x88, 0x4b, 0x94, 0xf0, 0x2f, 0xea, 0x67, 0x68, 0x31, 0x34, 0xc9, 0x26,
	0xd4, 0xf3, 0xbd, 0xd6, 0x1a, 0x69, 0x44, 0x1f, 0x41, 0x7b, 0x2e, 0xff, 0x69, 0xe8, 0xab, 0x2b,
	0x10, 0xf2, 0x40, 0x14, 0x57, 0x00, 0x6d, 0xac, 0x4e, 0xa7, 0xc1, 0x79, 0x34, 0x29, 0xaa, 0x73,
	0x44, 0x0f, 0xf2, 0x23, 0xa2, 0x87, 0xc7, 0xc3, 0x4b, 0x81, 0xd4, 0x81, 0x1f, 0xaa, 0x6a, 0x8b,
	0xa1, 0xa9, 0x3c, 0xfc, 0x5a, 0x55, 0xa2, 0x87, 0x5f, 0xd3, 0xef, 0x06, 0xc0, 0xbc, 0x6e, 0xb6,
	0xe0, 0x46, 0x69, 0xc1, 0xff, 0x85, 0x26, 0x9e, 0xa9, 0x0f, 0xa5, 0xbb, 0xd0, 0x40, 0xc7, 0x09,
	0x06, 0xb7, 0xa0, 0x19, 0x27, 0x62, 0xec, 0xa7, 0xd8, 0x4f, 0x55, 0x7d, 0x77, 0xee, 0x20, 0xbb,
	0x50, 0xcb, 0x42, 0x5f, 0xaa, 0xd1, 0xd9, 0xbd, 0xbf, 0x97, 0x96, 0x0c, 0xbb, 0x64, 0x2a, 0x85,
	0xdc, 0x05, 0x2b, 0xc1, 0x57, 0xab, 0x69, 0xd9, 0xbd, 0xcd, 0xe5, 0x85, 0xc4, 0x28, 0xcb, 0x93,
	0xe8, 0x21, 0xb4, 0xca, 0x9b, 0x7d, 0xfb, 0x7b, 0xd9, 0xfb, 0x69, 0x80, 0x3d, 0x08, 0xa5, 0x48,
	0x42, 0x3e, 0x79, 0x16, 0xfb, 0x64, 0x1f, 0x6a, 0x78, 0x1e, 0x09, 0xd1, 0x84, 0xa5, 0x93, 0xdf,
	0x59, 0x5f, 0xf0, 0xe5, 0xab, 0x46, 0x2b, 0xf7, 0x0d, 0xb2, 0x0f, 0x96, 0x9a, 0x30, 0x59, 0x5f,
	0xdc, 0xc6, 0xc5, 0xb2, 0xf2, 0x09, 0xa3, 0x15, 0xf2, 0x04, 0x1a, 0xc5, 0x01, 0x22, 0xe5, 0xf6,
	0x4a, 0x77, 0xbf, 0xf3, 0xcf, 0x92, 0xbf, 0xc4, 0xfa, 0x14, 0x9c, 0xd2, 0x27, 0x9f, 0x7f, 0x14,
	0xd8, 0xe4, 0x32, 0xd7, 0xc0, 0xeb, 0x6c, 0xdc, 0xf4, 0x93, 0xd0, 0xca, 0x79, 0x5d, 0xdd, 0x97,
	0xfd, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x33, 0xd3, 0x34, 0xc7, 0xef, 0x06, 0x00, 0x00,
}
