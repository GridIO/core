// Code generated by protoc-gen-go. DO NOT EDIT.
// source: node.proto

package sonm

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

type DealStatus int32

const (
	DealStatus_PENDING  DealStatus = 0
	DealStatus_APPROVED DealStatus = 1
	DealStatus_FINISHED DealStatus = 2
)

var DealStatus_name = map[int32]string{
	0: "PENDING",
	1: "APPROVED",
	2: "FINISHED",
}
var DealStatus_value = map[string]int32{
	"PENDING":  0,
	"APPROVED": 1,
	"FINISHED": 2,
}

func (x DealStatus) String() string {
	return proto.EnumName(DealStatus_name, int32(x))
}
func (DealStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

type TaskListRequest struct {
	// HubID is hub eth id;
	// If empty - collect task info from all hubs
	HubID string `protobuf:"bytes,1,opt,name=hubID" json:"hubID,omitempty"`
}

func (m *TaskListRequest) Reset()                    { *m = TaskListRequest{} }
func (m *TaskListRequest) String() string            { return proto.CompactTextString(m) }
func (*TaskListRequest) ProtoMessage()               {}
func (*TaskListRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *TaskListRequest) GetHubID() string {
	if m != nil {
		return m.HubID
	}
	return ""
}

type Deal struct {
	BidID  string     `protobuf:"bytes,1,opt,name=BidID" json:"BidID,omitempty"`
	AskID  string     `protobuf:"bytes,2,opt,name=AskID" json:"AskID,omitempty"`
	Status DealStatus `protobuf:"varint,3,opt,name=status,enum=sonm.DealStatus" json:"status,omitempty"`
}

func (m *Deal) Reset()                    { *m = Deal{} }
func (m *Deal) String() string            { return proto.CompactTextString(m) }
func (*Deal) ProtoMessage()               {}
func (*Deal) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *Deal) GetBidID() string {
	if m != nil {
		return m.BidID
	}
	return ""
}

func (m *Deal) GetAskID() string {
	if m != nil {
		return m.AskID
	}
	return ""
}

func (m *Deal) GetStatus() DealStatus {
	if m != nil {
		return m.Status
	}
	return DealStatus_PENDING
}

type DealListRequest struct {
	Owner  string     `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	Status DealStatus `protobuf:"varint,2,opt,name=status,enum=sonm.DealStatus" json:"status,omitempty"`
}

func (m *DealListRequest) Reset()                    { *m = DealListRequest{} }
func (m *DealListRequest) String() string            { return proto.CompactTextString(m) }
func (*DealListRequest) ProtoMessage()               {}
func (*DealListRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *DealListRequest) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *DealListRequest) GetStatus() DealStatus {
	if m != nil {
		return m.Status
	}
	return DealStatus_PENDING
}

type DealListReply struct {
	Deal []*Deal `protobuf:"bytes,1,rep,name=deal" json:"deal,omitempty"`
}

func (m *DealListReply) Reset()                    { *m = DealListReply{} }
func (m *DealListReply) String() string            { return proto.CompactTextString(m) }
func (*DealListReply) ProtoMessage()               {}
func (*DealListReply) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *DealListReply) GetDeal() []*Deal {
	if m != nil {
		return m.Deal
	}
	return nil
}

func init() {
	proto.RegisterType((*TaskListRequest)(nil), "sonm.TaskListRequest")
	proto.RegisterType((*Deal)(nil), "sonm.Deal")
	proto.RegisterType((*DealListRequest)(nil), "sonm.DealListRequest")
	proto.RegisterType((*DealListReply)(nil), "sonm.DealListReply")
	proto.RegisterEnum("sonm.DealStatus", DealStatus_name, DealStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TaskManagement service

type TaskManagementClient interface {
	// List produces a list of all tasks running on different SONM nodes
	List(ctx context.Context, in *TaskListRequest, opts ...grpc.CallOption) (*TaskListReply, error)
	// Start starts a task on given resource
	Start(ctx context.Context, in *HubStartTaskRequest, opts ...grpc.CallOption) (*TaskInfo, error)
	// Status produces a task status by their ID
	Status(ctx context.Context, in *ID, opts ...grpc.CallOption) (*TaskInfo, error)
	// Logs retrieves a task log (stdin/stderr) from given task
	Logs(ctx context.Context, in *ID, opts ...grpc.CallOption) (TaskManagement_LogsClient, error)
	// Stop stops a task by their IDd
	Stop(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
}

type taskManagementClient struct {
	cc *grpc.ClientConn
}

func NewTaskManagementClient(cc *grpc.ClientConn) TaskManagementClient {
	return &taskManagementClient{cc}
}

func (c *taskManagementClient) List(ctx context.Context, in *TaskListRequest, opts ...grpc.CallOption) (*TaskListReply, error) {
	out := new(TaskListReply)
	err := grpc.Invoke(ctx, "/sonm.TaskManagement/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagementClient) Start(ctx context.Context, in *HubStartTaskRequest, opts ...grpc.CallOption) (*TaskInfo, error) {
	out := new(TaskInfo)
	err := grpc.Invoke(ctx, "/sonm.TaskManagement/Start", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagementClient) Status(ctx context.Context, in *ID, opts ...grpc.CallOption) (*TaskInfo, error) {
	out := new(TaskInfo)
	err := grpc.Invoke(ctx, "/sonm.TaskManagement/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskManagementClient) Logs(ctx context.Context, in *ID, opts ...grpc.CallOption) (TaskManagement_LogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TaskManagement_serviceDesc.Streams[0], c.cc, "/sonm.TaskManagement/Logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskManagementLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TaskManagement_LogsClient interface {
	Recv() (*TaskLogsChunk, error)
	grpc.ClientStream
}

type taskManagementLogsClient struct {
	grpc.ClientStream
}

func (x *taskManagementLogsClient) Recv() (*TaskLogsChunk, error) {
	m := new(TaskLogsChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskManagementClient) Stop(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.TaskManagement/Stop", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskManagement service

type TaskManagementServer interface {
	// List produces a list of all tasks running on different SONM nodes
	List(context.Context, *TaskListRequest) (*TaskListReply, error)
	// Start starts a task on given resource
	Start(context.Context, *HubStartTaskRequest) (*TaskInfo, error)
	// Status produces a task status by their ID
	Status(context.Context, *ID) (*TaskInfo, error)
	// Logs retrieves a task log (stdin/stderr) from given task
	Logs(*ID, TaskManagement_LogsServer) error
	// Stop stops a task by their IDd
	Stop(context.Context, *ID) (*Empty, error)
}

func RegisterTaskManagementServer(s *grpc.Server, srv TaskManagementServer) {
	s.RegisterService(&_TaskManagement_serviceDesc, srv)
}

func _TaskManagement_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagementServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.TaskManagement/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagementServer).List(ctx, req.(*TaskListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManagement_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HubStartTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagementServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.TaskManagement/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagementServer).Start(ctx, req.(*HubStartTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManagement_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagementServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.TaskManagement/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagementServer).Status(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskManagement_Logs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskManagementServer).Logs(m, &taskManagementLogsServer{stream})
}

type TaskManagement_LogsServer interface {
	Send(*TaskLogsChunk) error
	grpc.ServerStream
}

type taskManagementLogsServer struct {
	grpc.ServerStream
}

func (x *taskManagementLogsServer) Send(m *TaskLogsChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _TaskManagement_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskManagementServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.TaskManagement/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskManagementServer).Stop(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.TaskManagement",
	HandlerType: (*TaskManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _TaskManagement_List_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _TaskManagement_Start_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _TaskManagement_Status_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _TaskManagement_Stop_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Logs",
			Handler:       _TaskManagement_Logs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "node.proto",
}

// Client API for DealManagement service

type DealManagementClient interface {
	// List produces a list of all deals made by client with given ID
	List(ctx context.Context, in *DealListRequest, opts ...grpc.CallOption) (*DealListReply, error)
	// Status produces a detailed info about deal with given ID
	Status(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Deal, error)
	// Finish finishes a deal with given ID
	Finish(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
}

type dealManagementClient struct {
	cc *grpc.ClientConn
}

func NewDealManagementClient(cc *grpc.ClientConn) DealManagementClient {
	return &dealManagementClient{cc}
}

func (c *dealManagementClient) List(ctx context.Context, in *DealListRequest, opts ...grpc.CallOption) (*DealListReply, error) {
	out := new(DealListReply)
	err := grpc.Invoke(ctx, "/sonm.DealManagement/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dealManagementClient) Status(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Deal, error) {
	out := new(Deal)
	err := grpc.Invoke(ctx, "/sonm.DealManagement/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dealManagementClient) Finish(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.DealManagement/Finish", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DealManagement service

type DealManagementServer interface {
	// List produces a list of all deals made by client with given ID
	List(context.Context, *DealListRequest) (*DealListReply, error)
	// Status produces a detailed info about deal with given ID
	Status(context.Context, *ID) (*Deal, error)
	// Finish finishes a deal with given ID
	Finish(context.Context, *ID) (*Empty, error)
}

func RegisterDealManagementServer(s *grpc.Server, srv DealManagementServer) {
	s.RegisterService(&_DealManagement_serviceDesc, srv)
}

func _DealManagement_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DealListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DealManagementServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.DealManagement/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DealManagementServer).List(ctx, req.(*DealListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DealManagement_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DealManagementServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.DealManagement/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DealManagementServer).Status(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DealManagement_Finish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DealManagementServer).Finish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.DealManagement/Finish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DealManagementServer).Finish(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _DealManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.DealManagement",
	HandlerType: (*DealManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _DealManagement_List_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _DealManagement_Status_Handler,
		},
		{
			MethodName: "Finish",
			Handler:    _DealManagement_Finish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}

// Client API for HubManagement service

type HubManagementClient interface {
	// Status produse a detailed info about Hub
	Status(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HubStatusReply, error)
	// WorkersList prouces a list of connected Workers
	WorkersList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListReply, error)
	// WorkersStatus produces a detailed info about a Worker with given ID
	WorkerStatus(ctx context.Context, in *ID, opts ...grpc.CallOption) (*InfoReply, error)
	// GetRegistredWorkers produce a list of Workers IDs allowed
	// to connect to this hub
	GetRegistredWorkers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetRegistredWorkersReply, error)
	// RegisterWorkers allows Worker with given ID connect to Hub
	RegisterWorker(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
	// RegisterWorkers deny Worker with given ID connect to Hub
	UnregisterWorker(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
	// GetMinerProperties allows to obtain previously assigned resource
	// properties for a given miner
	GetWorkerProperties(ctx context.Context, in *ID, opts ...grpc.CallOption) (*GetDevicePropertiesReply, error)
	// SetMinerProperties method allows to specify additional resource
	// properties for a miner specified by its ID
	// Note, that this method overrides all previously specified properties.
	SetWorkerProperties(ctx context.Context, in *SetDevicePropertiesRequest, opts ...grpc.CallOption) (*Empty, error)
	// GetAskPlans allows to obtain previously assigned Ask Plans from for a given worker.
	GetAskPlans(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SlotsReply, error)
	GetAskPlan(ctx context.Context, in *ID, opts ...grpc.CallOption) (*SlotsReply, error)
	// CreateAskPlan allows to create rules
	// for creating Ask orders on Marketplace
	CreateAskPlan(ctx context.Context, in *Slot, opts ...grpc.CallOption) (*Empty, error)
	// RemoveAskPlan allows to remove rules
	// for creating Ask orders on Marketplace
	RemoveAskPlan(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
	// List produces a list of all running tasks on the Hub
	TaskList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TaskListReply, error)
	// Status produces a detailed info about task on the Hub
	TaskStatus(ctx context.Context, in *ID, opts ...grpc.CallOption) (*TaskStatusReply, error)
}

type hubManagementClient struct {
	cc *grpc.ClientConn
}

func NewHubManagementClient(cc *grpc.ClientConn) HubManagementClient {
	return &hubManagementClient{cc}
}

func (c *hubManagementClient) Status(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*HubStatusReply, error) {
	out := new(HubStatusReply)
	err := grpc.Invoke(ctx, "/sonm.HubManagement/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubManagementClient) WorkersList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := grpc.Invoke(ctx, "/sonm.HubManagement/WorkersList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubManagementClient) WorkerStatus(ctx context.Context, in *ID, opts ...grpc.CallOption) (*InfoReply, error) {
	out := new(InfoReply)
	err := grpc.Invoke(ctx, "/sonm.HubManagement/WorkerStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubManagementClient) GetRegistredWorkers(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetRegistredWorkersReply, error) {
	out := new(GetRegistredWorkersReply)
	err := grpc.Invoke(ctx, "/sonm.HubManagement/GetRegistredWorkers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubManagementClient) RegisterWorker(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.HubManagement/RegisterWorker", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubManagementClient) UnregisterWorker(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.HubManagement/UnregisterWorker", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubManagementClient) GetWorkerProperties(ctx context.Context, in *ID, opts ...grpc.CallOption) (*GetDevicePropertiesReply, error) {
	out := new(GetDevicePropertiesReply)
	err := grpc.Invoke(ctx, "/sonm.HubManagement/GetWorkerProperties", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubManagementClient) SetWorkerProperties(ctx context.Context, in *SetDevicePropertiesRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.HubManagement/SetWorkerProperties", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubManagementClient) GetAskPlans(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SlotsReply, error) {
	out := new(SlotsReply)
	err := grpc.Invoke(ctx, "/sonm.HubManagement/GetAskPlans", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubManagementClient) GetAskPlan(ctx context.Context, in *ID, opts ...grpc.CallOption) (*SlotsReply, error) {
	out := new(SlotsReply)
	err := grpc.Invoke(ctx, "/sonm.HubManagement/GetAskPlan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubManagementClient) CreateAskPlan(ctx context.Context, in *Slot, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.HubManagement/CreateAskPlan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubManagementClient) RemoveAskPlan(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.HubManagement/RemoveAskPlan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubManagementClient) TaskList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TaskListReply, error) {
	out := new(TaskListReply)
	err := grpc.Invoke(ctx, "/sonm.HubManagement/TaskList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hubManagementClient) TaskStatus(ctx context.Context, in *ID, opts ...grpc.CallOption) (*TaskStatusReply, error) {
	out := new(TaskStatusReply)
	err := grpc.Invoke(ctx, "/sonm.HubManagement/TaskStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HubManagement service

type HubManagementServer interface {
	// Status produse a detailed info about Hub
	Status(context.Context, *Empty) (*HubStatusReply, error)
	// WorkersList prouces a list of connected Workers
	WorkersList(context.Context, *Empty) (*ListReply, error)
	// WorkersStatus produces a detailed info about a Worker with given ID
	WorkerStatus(context.Context, *ID) (*InfoReply, error)
	// GetRegistredWorkers produce a list of Workers IDs allowed
	// to connect to this hub
	GetRegistredWorkers(context.Context, *Empty) (*GetRegistredWorkersReply, error)
	// RegisterWorkers allows Worker with given ID connect to Hub
	RegisterWorker(context.Context, *ID) (*Empty, error)
	// RegisterWorkers deny Worker with given ID connect to Hub
	UnregisterWorker(context.Context, *ID) (*Empty, error)
	// GetMinerProperties allows to obtain previously assigned resource
	// properties for a given miner
	GetWorkerProperties(context.Context, *ID) (*GetDevicePropertiesReply, error)
	// SetMinerProperties method allows to specify additional resource
	// properties for a miner specified by its ID
	// Note, that this method overrides all previously specified properties.
	SetWorkerProperties(context.Context, *SetDevicePropertiesRequest) (*Empty, error)
	// GetAskPlans allows to obtain previously assigned Ask Plans from for a given worker.
	GetAskPlans(context.Context, *Empty) (*SlotsReply, error)
	GetAskPlan(context.Context, *ID) (*SlotsReply, error)
	// CreateAskPlan allows to create rules
	// for creating Ask orders on Marketplace
	CreateAskPlan(context.Context, *Slot) (*Empty, error)
	// RemoveAskPlan allows to remove rules
	// for creating Ask orders on Marketplace
	RemoveAskPlan(context.Context, *ID) (*Empty, error)
	// List produces a list of all running tasks on the Hub
	TaskList(context.Context, *Empty) (*TaskListReply, error)
	// Status produces a detailed info about task on the Hub
	TaskStatus(context.Context, *ID) (*TaskStatusReply, error)
}

func RegisterHubManagementServer(s *grpc.Server, srv HubManagementServer) {
	s.RegisterService(&_HubManagement_serviceDesc, srv)
}

func _HubManagement_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubManagementServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.HubManagement/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubManagementServer).Status(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HubManagement_WorkersList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubManagementServer).WorkersList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.HubManagement/WorkersList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubManagementServer).WorkersList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HubManagement_WorkerStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubManagementServer).WorkerStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.HubManagement/WorkerStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubManagementServer).WorkerStatus(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _HubManagement_GetRegistredWorkers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubManagementServer).GetRegistredWorkers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.HubManagement/GetRegistredWorkers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubManagementServer).GetRegistredWorkers(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HubManagement_RegisterWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubManagementServer).RegisterWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.HubManagement/RegisterWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubManagementServer).RegisterWorker(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _HubManagement_UnregisterWorker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubManagementServer).UnregisterWorker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.HubManagement/UnregisterWorker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubManagementServer).UnregisterWorker(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _HubManagement_GetWorkerProperties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubManagementServer).GetWorkerProperties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.HubManagement/GetWorkerProperties",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubManagementServer).GetWorkerProperties(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _HubManagement_SetWorkerProperties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDevicePropertiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubManagementServer).SetWorkerProperties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.HubManagement/SetWorkerProperties",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubManagementServer).SetWorkerProperties(ctx, req.(*SetDevicePropertiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HubManagement_GetAskPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubManagementServer).GetAskPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.HubManagement/GetAskPlans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubManagementServer).GetAskPlans(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HubManagement_GetAskPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubManagementServer).GetAskPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.HubManagement/GetAskPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubManagementServer).GetAskPlan(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _HubManagement_CreateAskPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Slot)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubManagementServer).CreateAskPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.HubManagement/CreateAskPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubManagementServer).CreateAskPlan(ctx, req.(*Slot))
	}
	return interceptor(ctx, in, info, handler)
}

func _HubManagement_RemoveAskPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubManagementServer).RemoveAskPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.HubManagement/RemoveAskPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubManagementServer).RemoveAskPlan(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _HubManagement_TaskList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubManagementServer).TaskList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.HubManagement/TaskList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubManagementServer).TaskList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _HubManagement_TaskStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HubManagementServer).TaskStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.HubManagement/TaskStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HubManagementServer).TaskStatus(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _HubManagement_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.HubManagement",
	HandlerType: (*HubManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _HubManagement_Status_Handler,
		},
		{
			MethodName: "WorkersList",
			Handler:    _HubManagement_WorkersList_Handler,
		},
		{
			MethodName: "WorkerStatus",
			Handler:    _HubManagement_WorkerStatus_Handler,
		},
		{
			MethodName: "GetRegistredWorkers",
			Handler:    _HubManagement_GetRegistredWorkers_Handler,
		},
		{
			MethodName: "RegisterWorker",
			Handler:    _HubManagement_RegisterWorker_Handler,
		},
		{
			MethodName: "UnregisterWorker",
			Handler:    _HubManagement_UnregisterWorker_Handler,
		},
		{
			MethodName: "GetWorkerProperties",
			Handler:    _HubManagement_GetWorkerProperties_Handler,
		},
		{
			MethodName: "SetWorkerProperties",
			Handler:    _HubManagement_SetWorkerProperties_Handler,
		},
		{
			MethodName: "GetAskPlans",
			Handler:    _HubManagement_GetAskPlans_Handler,
		},
		{
			MethodName: "GetAskPlan",
			Handler:    _HubManagement_GetAskPlan_Handler,
		},
		{
			MethodName: "CreateAskPlan",
			Handler:    _HubManagement_CreateAskPlan_Handler,
		},
		{
			MethodName: "RemoveAskPlan",
			Handler:    _HubManagement_RemoveAskPlan_Handler,
		},
		{
			MethodName: "TaskList",
			Handler:    _HubManagement_TaskList_Handler,
		},
		{
			MethodName: "TaskStatus",
			Handler:    _HubManagement_TaskStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}

func init() { proto.RegisterFile("node.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 619 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xb5, 0xdb, 0x34, 0x5f, 0x3b, 0x69, 0x7e, 0xb4, 0xed, 0x27, 0x85, 0x5c, 0x94, 0x60, 0x21,
	0x91, 0xb6, 0x6a, 0x8a, 0x42, 0x79, 0x80, 0x50, 0xa7, 0xad, 0xa5, 0x52, 0x82, 0xcd, 0xcf, 0x0d,
	0x37, 0x0e, 0x59, 0x12, 0x2b, 0xce, 0xae, 0xd9, 0x5d, 0x17, 0xf5, 0x19, 0x78, 0x3f, 0x1e, 0x83,
	0x67, 0x40, 0xeb, 0xb5, 0xeb, 0x8d, 0xe3, 0x08, 0xee, 0x3c, 0x67, 0xce, 0x9c, 0x39, 0x3b, 0x3b,
	0x6b, 0x00, 0x42, 0xa7, 0xb8, 0x1f, 0x31, 0x2a, 0x28, 0xaa, 0x70, 0x4a, 0x96, 0x9d, 0xbd, 0x49,
	0x30, 0x55, 0x40, 0xa7, 0x19, 0x10, 0x09, 0x91, 0xc0, 0x4f, 0x81, 0xda, 0x32, 0x20, 0x98, 0xa5,
	0xc1, 0xde, 0x3c, 0x9e, 0xa8, 0x4f, 0xeb, 0x05, 0x34, 0x3f, 0xf8, 0x7c, 0x71, 0x1b, 0x70, 0xe1,
	0xe2, 0xef, 0x31, 0xe6, 0x02, 0x1d, 0xc2, 0xce, 0x3c, 0x9e, 0x38, 0x76, 0xdb, 0xec, 0x9a, 0xbd,
	0x3d, 0x57, 0x05, 0xd6, 0x17, 0xa8, 0xd8, 0xd8, 0x0f, 0x65, 0xf6, 0x4d, 0x30, 0xcd, 0xb3, 0x49,
	0x20, 0xd1, 0x21, 0x5f, 0x38, 0x76, 0x7b, 0x4b, 0xa1, 0x49, 0x80, 0x7a, 0x50, 0xe5, 0xc2, 0x17,
	0x31, 0x6f, 0x6f, 0x77, 0xcd, 0x5e, 0x63, 0xd0, 0xea, 0x4b, 0x53, 0x7d, 0xa9, 0xe3, 0x25, 0xb8,
	0x9b, 0xe6, 0xad, 0xf7, 0xd0, 0x94, 0x68, 0xc1, 0x06, 0xfd, 0x41, 0x30, 0xcb, 0x1a, 0x25, 0x81,
	0x26, 0xb9, 0xf5, 0x17, 0xc9, 0x73, 0xa8, 0xe7, 0x92, 0x51, 0xf8, 0x80, 0x8e, 0xa0, 0x32, 0xc5,
	0x7e, 0xd8, 0x36, 0xbb, 0xdb, 0xbd, 0xda, 0x00, 0xf2, 0x42, 0x37, 0xc1, 0x4f, 0x5e, 0x03, 0xe4,
	0x32, 0xa8, 0x06, 0xff, 0x8d, 0x47, 0x77, 0xb6, 0x73, 0x77, 0xdd, 0x32, 0xd0, 0x3e, 0xec, 0x0e,
	0xc7, 0x63, 0xf7, 0xdd, 0xa7, 0x91, 0xdd, 0x32, 0x65, 0x74, 0xe5, 0xdc, 0x39, 0xde, 0xcd, 0xc8,
	0x6e, 0x6d, 0x0d, 0x7e, 0x9b, 0xd0, 0x90, 0x23, 0x7c, 0xeb, 0x13, 0x7f, 0x86, 0x97, 0x98, 0x08,
	0x74, 0x01, 0x15, 0xd9, 0x16, 0xfd, 0xaf, 0x7a, 0x14, 0x06, 0xdc, 0x39, 0x28, 0xc2, 0x51, 0xf8,
	0x60, 0x19, 0xe8, 0x02, 0x76, 0x3c, 0xe1, 0x33, 0x81, 0x9e, 0xa8, 0xfc, 0x4d, 0x3c, 0x49, 0x62,
	0xc9, 0xcb, 0x4a, 0x1b, 0x79, 0xa9, 0x43, 0xbe, 0x51, 0xcb, 0x40, 0xcf, 0xa1, 0x9a, 0x3a, 0xde,
	0x55, 0x39, 0xc7, 0x2e, 0x61, 0x1d, 0x43, 0xe5, 0x96, 0xce, 0x74, 0x8e, 0x6e, 0x82, 0xce, 0xf8,
	0xe5, 0x3c, 0x26, 0x0b, 0xcb, 0x78, 0x69, 0xa2, 0xa7, 0x50, 0xf1, 0x04, 0x8d, 0x34, 0x6a, 0x4d,
	0x7d, 0x8d, 0x96, 0x91, 0x78, 0xb0, 0x8c, 0xc1, 0x4f, 0x13, 0x1a, 0x72, 0x50, 0x9b, 0x0f, 0x5c,
	0xb8, 0xca, 0xac, 0xd7, 0xca, 0x75, 0x58, 0x06, 0xea, 0x96, 0x58, 0xd7, 0xae, 0xc5, 0x32, 0xd0,
	0x33, 0xa8, 0x5e, 0x05, 0x24, 0xe0, 0xf3, 0xcd, 0x6e, 0x7e, 0xed, 0x40, 0xfd, 0x26, 0x9e, 0x68,
	0x66, 0xce, 0x1e, 0x65, 0x75, 0x6a, 0xe7, 0x50, 0x9f, 0xaa, 0x5c, 0x94, 0xd4, 0xc5, 0x19, 0xd4,
	0x3e, 0x53, 0xb6, 0xc0, 0x8c, 0x27, 0x47, 0x58, 0xa9, 0x69, 0xaa, 0x40, 0x37, 0x7d, 0x0a, 0xfb,
	0x8a, 0xbe, 0x66, 0x3d, 0x25, 0xcb, 0x89, 0x67, 0x64, 0x1b, 0x0e, 0xae, 0xb1, 0x70, 0xf1, 0x2c,
	0xe0, 0x82, 0xe1, 0x69, 0xda, 0x67, 0xb5, 0xc7, 0x91, 0x0a, 0x4a, 0x78, 0x99, 0xca, 0x31, 0x34,
	0x54, 0x0a, 0x33, 0x95, 0xd9, 0x38, 0x0d, 0x74, 0x0a, 0xad, 0x8f, 0x84, 0xfd, 0x23, 0x79, 0x98,
	0xb8, 0x53, 0xac, 0x31, 0xa3, 0x11, 0x66, 0x22, 0xc0, 0xfa, 0x89, 0x72, 0x6b, 0x36, 0xbe, 0x0f,
	0xbe, 0xe2, 0x9c, 0x94, 0x59, 0xbb, 0x82, 0x03, 0xaf, 0x44, 0xa2, 0xab, 0x0a, 0xbd, 0xb2, 0x42,
	0xb5, 0x12, 0x05, 0x2b, 0x7d, 0xa8, 0x5d, 0x63, 0x31, 0xe4, 0x8b, 0x71, 0xe8, 0x93, 0xc2, 0x80,
	0xd2, 0x27, 0xee, 0x85, 0x54, 0x3c, 0xf6, 0x3d, 0x01, 0xc8, 0xf9, 0x9a, 0xe3, 0x72, 0x6e, 0xfd,
	0x92, 0x61, 0x5f, 0xe0, 0x8c, 0x0e, 0x39, 0xa9, 0xe8, 0xa3, 0x07, 0x75, 0x17, 0x2f, 0xe9, 0x3d,
	0x5e, 0x97, 0x5e, 0x73, 0xbc, 0x9b, 0x3d, 0xe0, 0x55, 0xbb, 0x1b, 0x5e, 0xf7, 0x39, 0x80, 0x84,
	0xd6, 0xb6, 0x46, 0xfb, 0x47, 0xac, 0xec, 0xe5, 0xa4, 0x9a, 0xfc, 0xa0, 0x5f, 0xfd, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x5c, 0xc2, 0xd2, 0xec, 0xe8, 0x05, 0x00, 0x00,
}