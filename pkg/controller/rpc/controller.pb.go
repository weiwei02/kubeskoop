// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.17.3
// source: controller.proto

package rpc

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

type TaskType int32

const (
	TaskType_Capture TaskType = 0
)

// Enum value maps for TaskType.
var (
	TaskType_name = map[int32]string{
		0: "Capture",
	}
	TaskType_value = map[string]int32{
		"Capture": 0,
	}
)

func (x TaskType) Enum() *TaskType {
	p := new(TaskType)
	*p = x
	return p
}

func (x TaskType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskType) Descriptor() protoreflect.EnumDescriptor {
	return file_controller_proto_enumTypes[0].Descriptor()
}

func (TaskType) Type() protoreflect.EnumType {
	return &file_controller_proto_enumTypes[0]
}

func (x TaskType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskType.Descriptor instead.
func (TaskType) EnumDescriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{0}
}

type AgentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName         string     `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Version          string     `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	SupportTaskTypes []TaskType `protobuf:"varint,3,rep,packed,name=support_task_types,json=supportTaskTypes,proto3,enum=controller_rpc.TaskType" json:"support_task_types,omitempty"`
}

func (x *AgentInfo) Reset() {
	*x = AgentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentInfo) ProtoMessage() {}

func (x *AgentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentInfo.ProtoReflect.Descriptor instead.
func (*AgentInfo) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{0}
}

func (x *AgentInfo) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *AgentInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *AgentInfo) GetSupportTaskTypes() []TaskType {
	if x != nil {
		return x.SupportTaskTypes
	}
	return nil
}

type ControllerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ControllerInfo) Reset() {
	*x = ControllerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControllerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerInfo) ProtoMessage() {}

func (x *ControllerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControllerInfo.ProtoReflect.Descriptor instead.
func (*ControllerInfo) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{1}
}

func (x *ControllerInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp int64  `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Event) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EventReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EventReply) Reset() {
	*x = EventReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventReply) ProtoMessage() {}

func (x *EventReply) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventReply.ProtoReflect.Descriptor instead.
func (*EventReply) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{3}
}

func (x *EventReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *EventReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type TaskFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string   `protobuf:"bytes,1,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	Type     TaskType `protobuf:"varint,2,opt,name=type,proto3,enum=controller_rpc.TaskType" json:"type,omitempty"`
}

func (x *TaskFilter) Reset() {
	*x = TaskFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskFilter) ProtoMessage() {}

func (x *TaskFilter) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskFilter.ProtoReflect.Descriptor instead.
func (*TaskFilter) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{4}
}

func (x *TaskFilter) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *TaskFilter) GetType() TaskType {
	if x != nil {
		return x.Type
	}
	return TaskType_Capture
}

type ServerTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server *ControllerInfo `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Task   *Task           `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *ServerTask) Reset() {
	*x = ServerTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerTask) ProtoMessage() {}

func (x *ServerTask) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerTask.ProtoReflect.Descriptor instead.
func (*ServerTask) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{5}
}

func (x *ServerTask) GetServer() *ControllerInfo {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *ServerTask) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type TaskType `protobuf:"varint,1,opt,name=type,proto3,enum=controller_rpc.TaskType" json:"type,omitempty"`
	Id   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to TaskInfo:
	//
	//	*Task_Capture
	TaskInfo isTask_TaskInfo `protobuf_oneof:"TaskInfo"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{6}
}

func (x *Task) GetType() TaskType {
	if x != nil {
		return x.Type
	}
	return TaskType_Capture
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *Task) GetTaskInfo() isTask_TaskInfo {
	if m != nil {
		return m.TaskInfo
	}
	return nil
}

func (x *Task) GetCapture() *CaptureInfo {
	if x, ok := x.GetTaskInfo().(*Task_Capture); ok {
		return x.Capture
	}
	return nil
}

type isTask_TaskInfo interface {
	isTask_TaskInfo()
}

type Task_Capture struct {
	Capture *CaptureInfo `protobuf:"bytes,3,opt,name=capture,proto3,oneof"`
}

func (*Task_Capture) isTask_TaskInfo() {}

type PodInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Ipv4      string `protobuf:"bytes,3,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	Ipv6      string `protobuf:"bytes,4,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
}

func (x *PodInfo) Reset() {
	*x = PodInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodInfo) ProtoMessage() {}

func (x *PodInfo) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodInfo.ProtoReflect.Descriptor instead.
func (*PodInfo) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{7}
}

func (x *PodInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PodInfo) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *PodInfo) GetIpv4() string {
	if x != nil {
		return x.Ipv4
	}
	return ""
}

func (x *PodInfo) GetIpv6() string {
	if x != nil {
		return x.Ipv6
	}
	return ""
}

type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{8}
}

func (x *NodeInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CaptureInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pod           *PodInfo  `protobuf:"bytes,1,opt,name=pod,proto3" json:"pod,omitempty"`
	Node          *NodeInfo `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	CaptureHostNs bool      `protobuf:"varint,4,opt,name=capture_host_ns,json=captureHostNs,proto3" json:"capture_host_ns,omitempty"`
}

func (x *CaptureInfo) Reset() {
	*x = CaptureInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaptureInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaptureInfo) ProtoMessage() {}

func (x *CaptureInfo) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaptureInfo.ProtoReflect.Descriptor instead.
func (*CaptureInfo) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{9}
}

func (x *CaptureInfo) GetPod() *PodInfo {
	if x != nil {
		return x.Pod
	}
	return nil
}

func (x *CaptureInfo) GetNode() *NodeInfo {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *CaptureInfo) GetCaptureHostNs() bool {
	if x != nil {
		return x.CaptureHostNs
	}
	return false
}

type CaptureResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message []byte `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CaptureResult) Reset() {
	*x = CaptureResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaptureResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaptureResult) ProtoMessage() {}

func (x *CaptureResult) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaptureResult.ProtoReflect.Descriptor instead.
func (*CaptureResult) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{10}
}

func (x *CaptureResult) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

type TaskResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type    TaskType `protobuf:"varint,2,opt,name=type,proto3,enum=controller_rpc.TaskType" json:"type,omitempty"`
	Success bool     `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	Message string   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	// Types that are assignable to TaskResultInfo:
	//
	//	*TaskResult_Capture
	TaskResultInfo isTaskResult_TaskResultInfo `protobuf_oneof:"TaskResultInfo"`
}

func (x *TaskResult) Reset() {
	*x = TaskResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResult) ProtoMessage() {}

func (x *TaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResult.ProtoReflect.Descriptor instead.
func (*TaskResult) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{11}
}

func (x *TaskResult) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TaskResult) GetType() TaskType {
	if x != nil {
		return x.Type
	}
	return TaskType_Capture
}

func (x *TaskResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *TaskResult) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (m *TaskResult) GetTaskResultInfo() isTaskResult_TaskResultInfo {
	if m != nil {
		return m.TaskResultInfo
	}
	return nil
}

func (x *TaskResult) GetCapture() *CaptureResult {
	if x, ok := x.GetTaskResultInfo().(*TaskResult_Capture); ok {
		return x.Capture
	}
	return nil
}

type isTaskResult_TaskResultInfo interface {
	isTaskResult_TaskResultInfo()
}

type TaskResult_Capture struct {
	Capture *CaptureResult `protobuf:"bytes,5,opt,name=capture,proto3,oneof"`
}

func (*TaskResult_Capture) isTaskResult_TaskResultInfo() {}

type TaskResultReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *TaskResultReply) Reset() {
	*x = TaskResultReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_controller_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResultReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResultReply) ProtoMessage() {}

func (x *TaskResultReply) ProtoReflect() protoreflect.Message {
	mi := &file_controller_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResultReply.ProtoReflect.Descriptor instead.
func (*TaskResultReply) Descriptor() ([]byte, []int) {
	return file_controller_proto_rawDescGZIP(), []int{12}
}

func (x *TaskResultReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *TaskResultReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_controller_proto protoreflect.FileDescriptor

var file_controller_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72,
	0x70, 0x63, 0x22, 0x8a, 0x01, 0x0a, 0x09, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x12, 0x73, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x5f, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x10, 0x73,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22,
	0x2a, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x53, 0x0a, 0x05, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x40, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x57, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x6e, 0x0a, 0x0a, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x89, 0x01, 0x0a, 0x04,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f,
	0x72, 0x70, 0x63, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x37, 0x0a, 0x07, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x48, 0x00, 0x52, 0x07, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x54,
	0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x63, 0x0a, 0x07, 0x50, 0x6f, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x76, 0x34, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x36,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x70, 0x76, 0x36, 0x22, 0x1e, 0x0a, 0x08,
	0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x8e, 0x01, 0x0a,
	0x0b, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x29, 0x0a, 0x03,
	0x70, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x03, 0x70, 0x6f, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x6e, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65,
	0x5f, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x63, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x4e, 0x73, 0x22, 0x29, 0x0a,
	0x0d, 0x43, 0x61, 0x70, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xcb, 0x01, 0x0a, 0x0a, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x63, 0x61, 0x70,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x61, 0x70, 0x74,
	0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x61, 0x70,
	0x74, 0x75, 0x72, 0x65, 0x42, 0x10, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x45, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x17, 0x0a,
	0x08, 0x54, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x61, 0x70,
	0x74, 0x75, 0x72, 0x65, 0x10, 0x00, 0x32, 0xc5, 0x02, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c,
	0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70,
	0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x43, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x15, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70,
	0x63, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x28, 0x01, 0x12, 0x46, 0x0a, 0x0a, 0x57, 0x61, 0x74, 0x63, 0x68, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x5f, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x1a,
	0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x61, 0x73, 0x6b, 0x30, 0x01, 0x12, 0x4f, 0x0a,
	0x10, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x1a, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72,
	0x70, 0x63, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a, 0x1f, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x08,
	0x5a, 0x06, 0x2e, 0x2f, 0x3b, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_controller_proto_rawDescOnce sync.Once
	file_controller_proto_rawDescData = file_controller_proto_rawDesc
)

func file_controller_proto_rawDescGZIP() []byte {
	file_controller_proto_rawDescOnce.Do(func() {
		file_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_controller_proto_rawDescData)
	})
	return file_controller_proto_rawDescData
}

var file_controller_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_controller_proto_goTypes = []interface{}{
	(TaskType)(0),           // 0: controller_rpc.TaskType
	(*AgentInfo)(nil),       // 1: controller_rpc.AgentInfo
	(*ControllerInfo)(nil),  // 2: controller_rpc.ControllerInfo
	(*Event)(nil),           // 3: controller_rpc.Event
	(*EventReply)(nil),      // 4: controller_rpc.EventReply
	(*TaskFilter)(nil),      // 5: controller_rpc.TaskFilter
	(*ServerTask)(nil),      // 6: controller_rpc.ServerTask
	(*Task)(nil),            // 7: controller_rpc.Task
	(*PodInfo)(nil),         // 8: controller_rpc.PodInfo
	(*NodeInfo)(nil),        // 9: controller_rpc.NodeInfo
	(*CaptureInfo)(nil),     // 10: controller_rpc.CaptureInfo
	(*CaptureResult)(nil),   // 11: controller_rpc.CaptureResult
	(*TaskResult)(nil),      // 12: controller_rpc.TaskResult
	(*TaskResultReply)(nil), // 13: controller_rpc.TaskResultReply
}
var file_controller_proto_depIdxs = []int32{
	0,  // 0: controller_rpc.AgentInfo.support_task_types:type_name -> controller_rpc.TaskType
	0,  // 1: controller_rpc.TaskFilter.type:type_name -> controller_rpc.TaskType
	2,  // 2: controller_rpc.ServerTask.server:type_name -> controller_rpc.ControllerInfo
	7,  // 3: controller_rpc.ServerTask.task:type_name -> controller_rpc.Task
	0,  // 4: controller_rpc.Task.type:type_name -> controller_rpc.TaskType
	10, // 5: controller_rpc.Task.capture:type_name -> controller_rpc.CaptureInfo
	8,  // 6: controller_rpc.CaptureInfo.pod:type_name -> controller_rpc.PodInfo
	9,  // 7: controller_rpc.CaptureInfo.node:type_name -> controller_rpc.NodeInfo
	0,  // 8: controller_rpc.TaskResult.type:type_name -> controller_rpc.TaskType
	11, // 9: controller_rpc.TaskResult.capture:type_name -> controller_rpc.CaptureResult
	1,  // 10: controller_rpc.ControllerRegisterService.RegisterAgent:input_type -> controller_rpc.AgentInfo
	3,  // 11: controller_rpc.ControllerRegisterService.ReportEvents:input_type -> controller_rpc.Event
	5,  // 12: controller_rpc.ControllerRegisterService.WatchTasks:input_type -> controller_rpc.TaskFilter
	12, // 13: controller_rpc.ControllerRegisterService.UploadTaskResult:input_type -> controller_rpc.TaskResult
	2,  // 14: controller_rpc.ControllerRegisterService.RegisterAgent:output_type -> controller_rpc.ControllerInfo
	4,  // 15: controller_rpc.ControllerRegisterService.ReportEvents:output_type -> controller_rpc.EventReply
	6,  // 16: controller_rpc.ControllerRegisterService.WatchTasks:output_type -> controller_rpc.ServerTask
	13, // 17: controller_rpc.ControllerRegisterService.UploadTaskResult:output_type -> controller_rpc.TaskResultReply
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_controller_proto_init() }
func file_controller_proto_init() {
	if File_controller_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_controller_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentInfo); i {
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
		file_controller_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControllerInfo); i {
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
		file_controller_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_controller_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventReply); i {
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
		file_controller_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskFilter); i {
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
		file_controller_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerTask); i {
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
		file_controller_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_controller_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodInfo); i {
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
		file_controller_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfo); i {
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
		file_controller_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaptureInfo); i {
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
		file_controller_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaptureResult); i {
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
		file_controller_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResult); i {
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
		file_controller_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskResultReply); i {
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
	file_controller_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*Task_Capture)(nil),
	}
	file_controller_proto_msgTypes[11].OneofWrappers = []interface{}{
		(*TaskResult_Capture)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_controller_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_controller_proto_goTypes,
		DependencyIndexes: file_controller_proto_depIdxs,
		EnumInfos:         file_controller_proto_enumTypes,
		MessageInfos:      file_controller_proto_msgTypes,
	}.Build()
	File_controller_proto = out.File
	file_controller_proto_rawDesc = nil
	file_controller_proto_goTypes = nil
	file_controller_proto_depIdxs = nil
}
