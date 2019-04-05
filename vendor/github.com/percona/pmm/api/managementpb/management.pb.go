// Code generated by protoc-gen-go. DO NOT EDIT.
// source: managementpb/management.proto

package managementpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	inventory "github.com/percona/pmm/api/inventory"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

type AddMySQLRequest struct {
	// Node identifier on which a service is been running. Required.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Unique across all Services user-defined name. Required.
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Node and Service access address (DNS name or IP). Required.
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	// Service Access port. Required.
	Port uint32 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	// The "pmm-agent" identifier which should run agents. Required.
	PmmAgentId string `protobuf:"bytes,5,opt,name=pmm_agent_id,json=pmmAgentId,proto3" json:"pmm_agent_id,omitempty"`
	// MySQL username for scraping metrics.
	Username string `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	// MySQL password for scraping metrics.
	Password string `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	// FIXME remove
	QanUsername string `protobuf:"bytes,8,opt,name=qan_username,json=qanUsername,proto3" json:"qan_username,omitempty"`
	// FIXME remove
	QanPassword string `protobuf:"bytes,9,opt,name=qan_password,json=qanPassword,proto3" json:"qan_password,omitempty"`
	// If true, adds mysqld_exporter for provided service.
	MysqldExporter bool `protobuf:"varint,10,opt,name=mysqld_exporter,json=mysqldExporter,proto3" json:"mysqld_exporter,omitempty"`
	// If true, adds qan-mysql-perfschema-agent for provided service.
	QanMysqlPerfschema   bool     `protobuf:"varint,11,opt,name=qan_mysql_perfschema,json=qanMysqlPerfschema,proto3" json:"qan_mysql_perfschema,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddMySQLRequest) Reset()         { *m = AddMySQLRequest{} }
func (m *AddMySQLRequest) String() string { return proto.CompactTextString(m) }
func (*AddMySQLRequest) ProtoMessage()    {}
func (*AddMySQLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3669328120c0a3b7, []int{0}
}

func (m *AddMySQLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMySQLRequest.Unmarshal(m, b)
}
func (m *AddMySQLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMySQLRequest.Marshal(b, m, deterministic)
}
func (m *AddMySQLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMySQLRequest.Merge(m, src)
}
func (m *AddMySQLRequest) XXX_Size() int {
	return xxx_messageInfo_AddMySQLRequest.Size(m)
}
func (m *AddMySQLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMySQLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddMySQLRequest proto.InternalMessageInfo

func (m *AddMySQLRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *AddMySQLRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *AddMySQLRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AddMySQLRequest) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *AddMySQLRequest) GetPmmAgentId() string {
	if m != nil {
		return m.PmmAgentId
	}
	return ""
}

func (m *AddMySQLRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AddMySQLRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AddMySQLRequest) GetQanUsername() string {
	if m != nil {
		return m.QanUsername
	}
	return ""
}

func (m *AddMySQLRequest) GetQanPassword() string {
	if m != nil {
		return m.QanPassword
	}
	return ""
}

func (m *AddMySQLRequest) GetMysqldExporter() bool {
	if m != nil {
		return m.MysqldExporter
	}
	return false
}

func (m *AddMySQLRequest) GetQanMysqlPerfschema() bool {
	if m != nil {
		return m.QanMysqlPerfschema
	}
	return false
}

type AddMySQLResponse struct {
	Service              *inventory.MySQLService            `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	MysqldExporter       *inventory.MySQLdExporter          `protobuf:"bytes,2,opt,name=mysqld_exporter,json=mysqldExporter,proto3" json:"mysqld_exporter,omitempty"`
	QanMysqlPerfschema   *inventory.QANMySQLPerfSchemaAgent `protobuf:"bytes,3,opt,name=qan_mysql_perfschema,json=qanMysqlPerfschema,proto3" json:"qan_mysql_perfschema,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *AddMySQLResponse) Reset()         { *m = AddMySQLResponse{} }
func (m *AddMySQLResponse) String() string { return proto.CompactTextString(m) }
func (*AddMySQLResponse) ProtoMessage()    {}
func (*AddMySQLResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3669328120c0a3b7, []int{1}
}

func (m *AddMySQLResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMySQLResponse.Unmarshal(m, b)
}
func (m *AddMySQLResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMySQLResponse.Marshal(b, m, deterministic)
}
func (m *AddMySQLResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMySQLResponse.Merge(m, src)
}
func (m *AddMySQLResponse) XXX_Size() int {
	return xxx_messageInfo_AddMySQLResponse.Size(m)
}
func (m *AddMySQLResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMySQLResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddMySQLResponse proto.InternalMessageInfo

func (m *AddMySQLResponse) GetService() *inventory.MySQLService {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *AddMySQLResponse) GetMysqldExporter() *inventory.MySQLdExporter {
	if m != nil {
		return m.MysqldExporter
	}
	return nil
}

func (m *AddMySQLResponse) GetQanMysqlPerfschema() *inventory.QANMySQLPerfSchemaAgent {
	if m != nil {
		return m.QanMysqlPerfschema
	}
	return nil
}

type RegisterNodeRequest struct {
	// Type of node which will be registered.
	NodeType inventory.NodeType `protobuf:"varint,1,opt,name=node_type,json=nodeType,proto3,enum=inventory.NodeType" json:"node_type,omitempty"`
	// Unique across all Nodes user-defined name, can be changed.
	NodeName string `protobuf:"bytes,2,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Linux machine-id. Can't be changed. Must be unique across all Generic Nodes if specified (optional).
	MachineId string `protobuf:"bytes,3,opt,name=machine_id,json=machineId,proto3" json:"machine_id,omitempty"`
	// Linux distribution. For Generic nodes (optional).
	Distro string `protobuf:"bytes,4,opt,name=distro,proto3" json:"distro,omitempty"`
	// Linux distribution version. For Generic nodes (optional).
	DistroVersion string `protobuf:"bytes,5,opt,name=distro_version,json=distroVersion,proto3" json:"distro_version,omitempty"`
	// Сontainer identifier. Only for Container Nodes (optional).
	ContainerId string `protobuf:"bytes,6,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	// Сontainer name. Only for Container Nodes (optional).
	ContainerName string `protobuf:"bytes,7,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	// Custom user-assigned labels for node (optional).
	CustomLabels map[string]string `protobuf:"bytes,10,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Address FIXME https://jira.percona.com/browse/PMM-3786
	Address              string   `protobuf:"bytes,42,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterNodeRequest) Reset()         { *m = RegisterNodeRequest{} }
func (m *RegisterNodeRequest) String() string { return proto.CompactTextString(m) }
func (*RegisterNodeRequest) ProtoMessage()    {}
func (*RegisterNodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3669328120c0a3b7, []int{2}
}

func (m *RegisterNodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterNodeRequest.Unmarshal(m, b)
}
func (m *RegisterNodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterNodeRequest.Marshal(b, m, deterministic)
}
func (m *RegisterNodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterNodeRequest.Merge(m, src)
}
func (m *RegisterNodeRequest) XXX_Size() int {
	return xxx_messageInfo_RegisterNodeRequest.Size(m)
}
func (m *RegisterNodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterNodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterNodeRequest proto.InternalMessageInfo

func (m *RegisterNodeRequest) GetNodeType() inventory.NodeType {
	if m != nil {
		return m.NodeType
	}
	return inventory.NodeType_NODE_TYPE_INVALID
}

func (m *RegisterNodeRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *RegisterNodeRequest) GetMachineId() string {
	if m != nil {
		return m.MachineId
	}
	return ""
}

func (m *RegisterNodeRequest) GetDistro() string {
	if m != nil {
		return m.Distro
	}
	return ""
}

func (m *RegisterNodeRequest) GetDistroVersion() string {
	if m != nil {
		return m.DistroVersion
	}
	return ""
}

func (m *RegisterNodeRequest) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func (m *RegisterNodeRequest) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *RegisterNodeRequest) GetCustomLabels() map[string]string {
	if m != nil {
		return m.CustomLabels
	}
	return nil
}

func (m *RegisterNodeRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type RegisterNodeResponse struct {
	GenericNode          *inventory.GenericNode   `protobuf:"bytes,1,opt,name=generic_node,json=genericNode,proto3" json:"generic_node,omitempty"`
	ContainerNode        *inventory.ContainerNode `protobuf:"bytes,2,opt,name=container_node,json=containerNode,proto3" json:"container_node,omitempty"`
	PmmAgent             *inventory.PMMAgent      `protobuf:"bytes,3,opt,name=pmm_agent,json=pmmAgent,proto3" json:"pmm_agent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *RegisterNodeResponse) Reset()         { *m = RegisterNodeResponse{} }
func (m *RegisterNodeResponse) String() string { return proto.CompactTextString(m) }
func (*RegisterNodeResponse) ProtoMessage()    {}
func (*RegisterNodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3669328120c0a3b7, []int{3}
}

func (m *RegisterNodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterNodeResponse.Unmarshal(m, b)
}
func (m *RegisterNodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterNodeResponse.Marshal(b, m, deterministic)
}
func (m *RegisterNodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterNodeResponse.Merge(m, src)
}
func (m *RegisterNodeResponse) XXX_Size() int {
	return xxx_messageInfo_RegisterNodeResponse.Size(m)
}
func (m *RegisterNodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterNodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterNodeResponse proto.InternalMessageInfo

func (m *RegisterNodeResponse) GetGenericNode() *inventory.GenericNode {
	if m != nil {
		return m.GenericNode
	}
	return nil
}

func (m *RegisterNodeResponse) GetContainerNode() *inventory.ContainerNode {
	if m != nil {
		return m.ContainerNode
	}
	return nil
}

func (m *RegisterNodeResponse) GetPmmAgent() *inventory.PMMAgent {
	if m != nil {
		return m.PmmAgent
	}
	return nil
}

func init() {
	proto.RegisterType((*AddMySQLRequest)(nil), "management.AddMySQLRequest")
	proto.RegisterType((*AddMySQLResponse)(nil), "management.AddMySQLResponse")
	proto.RegisterType((*RegisterNodeRequest)(nil), "management.RegisterNodeRequest")
	proto.RegisterMapType((map[string]string)(nil), "management.RegisterNodeRequest.CustomLabelsEntry")
	proto.RegisterType((*RegisterNodeResponse)(nil), "management.RegisterNodeResponse")
}

func init() { proto.RegisterFile("managementpb/management.proto", fileDescriptor_3669328120c0a3b7) }

var fileDescriptor_3669328120c0a3b7 = []byte{
	// 902 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xdd, 0x6e, 0x23, 0x35,
	0x18, 0x65, 0x92, 0x36, 0x3f, 0x5f, 0xd2, 0x6e, 0xf1, 0x96, 0x32, 0x84, 0x56, 0xcd, 0xa6, 0x42,
	0xcd, 0x16, 0x9a, 0x69, 0x83, 0x84, 0xa0, 0x37, 0xab, 0x74, 0x55, 0xa1, 0x4a, 0x6d, 0xd5, 0x9d,
	0x2e, 0x2b, 0xc4, 0xcd, 0xc8, 0x8d, 0xdd, 0xe9, 0x68, 0x33, 0xf6, 0xc4, 0x76, 0x52, 0x72, 0xc3,
	0x05, 0x8f, 0x00, 0xf7, 0x3c, 0x0d, 0x12, 0x0f, 0xc0, 0x03, 0xac, 0x84, 0x78, 0x10, 0x64, 0x7b,
	0x66, 0x32, 0xe9, 0x8f, 0xb8, 0xca, 0xf8, 0x9c, 0xf3, 0xfd, 0xf8, 0xf3, 0x89, 0x0d, 0x5b, 0x31,
	0x66, 0x38, 0xa4, 0x31, 0x65, 0x2a, 0xb9, 0xf6, 0xe6, 0x8b, 0x5e, 0x22, 0xb8, 0xe2, 0x08, 0xe6,
	0x48, 0xeb, 0x9b, 0x30, 0x52, 0xb7, 0x93, 0xeb, 0xde, 0x90, 0xc7, 0x5e, 0x7c, 0x17, 0xa9, 0xf7,
	0xfc, 0xce, 0x0b, 0xf9, 0xbe, 0x11, 0xee, 0x4f, 0xf1, 0x28, 0x22, 0x58, 0x71, 0x21, 0xbd, 0xfc,
	0xd3, 0xe6, 0x68, 0x6d, 0x86, 0x9c, 0x87, 0x23, 0xea, 0xe1, 0x24, 0xf2, 0x30, 0x63, 0x5c, 0x61,
	0x15, 0x71, 0x26, 0x53, 0x76, 0x23, 0x62, 0x53, 0xca, 0x14, 0x17, 0x33, 0x0f, 0x87, 0x94, 0xa9,
	0x0c, 0xff, 0x64, 0x8e, 0x33, 0x4e, 0x68, 0x06, 0xbb, 0x73, 0x58, 0x52, 0x31, 0x8d, 0x86, 0x39,
	0xf3, 0x95, 0xf9, 0x19, 0xee, 0x87, 0x94, 0xed, 0xcb, 0x3b, 0x1c, 0x86, 0x54, 0x78, 0x3c, 0x31,
	0xa5, 0x1e, 0x96, 0xed, 0xfc, 0x51, 0x86, 0x67, 0x03, 0x42, 0xce, 0x67, 0x57, 0x6f, 0xce, 0x7c,
	0x3a, 0x9e, 0x50, 0xa9, 0xd0, 0x36, 0x54, 0x75, 0xa9, 0x20, 0x22, 0xae, 0xd3, 0x76, 0xba, 0xf5,
	0xe3, 0xca, 0x3f, 0x1f, 0xb6, 0x4b, 0x3f, 0x3a, 0x7e, 0x45, 0xc3, 0xa7, 0x04, 0xbd, 0x84, 0x66,
	0x5a, 0x34, 0x60, 0x38, 0xa6, 0x6e, 0x69, 0x41, 0xd5, 0x48, 0xb9, 0x0b, 0x1c, 0x53, 0xd4, 0x86,
	0x2a, 0x26, 0x44, 0x50, 0x29, 0xdd, 0xf2, 0x82, 0x2a, 0x83, 0x51, 0x0b, 0x96, 0x12, 0x2e, 0x94,
	0xbb, 0xd4, 0x76, 0xba, 0x2b, 0x96, 0x5e, 0xfb, 0xc8, 0x37, 0x18, 0xea, 0x42, 0x33, 0x89, 0xe3,
	0xc0, 0x0c, 0x44, 0xb7, 0xb3, 0xbc, 0x90, 0x02, 0x92, 0x38, 0x1e, 0x68, 0xea, 0x94, 0xa0, 0x16,
	0xd4, 0x26, 0x92, 0x0a, 0xd3, 0x4e, 0x45, 0xab, 0xfc, 0x7c, 0xad, 0xb9, 0x04, 0x4b, 0x79, 0xc7,
	0x05, 0x71, 0xab, 0x96, 0xcb, 0xd6, 0xe8, 0x05, 0x34, 0xc7, 0x98, 0x05, 0x79, 0x6c, 0xcd, 0xf0,
	0x8d, 0x31, 0x66, 0x3f, 0x64, 0xe1, 0xa9, 0x24, 0x4f, 0x51, 0xcf, 0x25, 0x97, 0x59, 0x96, 0x5d,
	0x78, 0x16, 0xcf, 0xe4, 0x78, 0x44, 0x02, 0xfa, 0xb3, 0x6e, 0x9c, 0x0a, 0x17, 0xda, 0x4e, 0xb7,
	0xe6, 0xaf, 0x5a, 0xf8, 0x24, 0x45, 0xd1, 0x01, 0xac, 0xeb, 0x5c, 0x06, 0x0d, 0x12, 0x2a, 0x6e,
	0xe4, 0xf0, 0x96, 0xc6, 0xd8, 0x6d, 0x18, 0x35, 0x1a, 0x63, 0x76, 0xae, 0xa9, 0xcb, 0x9c, 0xe9,
	0x7c, 0x70, 0x60, 0x6d, 0x7e, 0x40, 0x32, 0xe1, 0x4c, 0x52, 0x74, 0x08, 0xd5, 0x74, 0xc8, 0xe6,
	0x84, 0x1a, 0xfd, 0x4f, 0x7b, 0xb9, 0x1f, 0x7a, 0x46, 0x7a, 0x65, 0x69, 0x3f, 0xd3, 0xa1, 0xe3,
	0x87, 0x2d, 0x96, 0x4c, 0xe8, 0x67, 0xf7, 0x43, 0xf3, 0x6e, 0x1f, 0x74, 0xff, 0xf6, 0x89, 0xee,
	0xcb, 0x26, 0x51, 0xa7, 0x90, 0xe8, 0xcd, 0xe0, 0xc2, 0xe4, 0xd2, 0x1b, 0xb9, 0x32, 0x22, 0x73,
	0x50, 0x8f, 0xee, 0xf0, 0xaf, 0x32, 0x3c, 0xf7, 0x69, 0x18, 0x49, 0x45, 0xc5, 0x05, 0x27, 0x34,
	0xb3, 0xe1, 0x01, 0xd4, 0x8d, 0x0d, 0xd5, 0x2c, 0xb1, 0xdb, 0x5c, 0xed, 0x3f, 0x2f, 0x94, 0xd0,
	0xd2, 0xb7, 0xb3, 0x84, 0xfa, 0x35, 0x96, 0x7e, 0xa1, 0x9d, 0x34, 0xe2, 0x11, 0x53, 0x1a, 0x91,
	0x71, 0xe4, 0x16, 0x40, 0x8c, 0x87, 0xb7, 0x11, 0x33, 0x06, 0x37, 0xa6, 0xf4, 0xeb, 0x29, 0x72,
	0x4a, 0xd0, 0x06, 0x54, 0x48, 0x24, 0x95, 0xe0, 0xc6, 0x90, 0x75, 0x3f, 0x5d, 0xa1, 0x2f, 0x60,
	0xd5, 0x7e, 0x05, 0x53, 0x2a, 0x64, 0xc4, 0x99, 0x35, 0xa3, 0xbf, 0x62, 0xd1, 0x77, 0x16, 0xd4,
	0x66, 0x19, 0x72, 0xa6, 0x70, 0xc4, 0xa8, 0xd0, 0xf9, 0xad, 0x17, 0x1b, 0x39, 0x76, 0x4a, 0x74,
	0xa6, 0xb9, 0xc4, 0xb4, 0x6a, 0x4d, 0xb9, 0x92, 0xa3, 0xa6, 0xcf, 0x77, 0xb0, 0x32, 0x9c, 0x48,
	0xc5, 0xe3, 0x60, 0x84, 0xaf, 0xe9, 0x48, 0xba, 0xd0, 0x2e, 0x77, 0x1b, 0xfd, 0xc3, 0x5e, 0xe1,
	0x72, 0x7a, 0x64, 0x6c, 0xbd, 0xd7, 0x26, 0xe8, 0xcc, 0xc4, 0x9c, 0x30, 0x25, 0x66, 0x7e, 0x73,
	0x58, 0x80, 0x90, 0x3b, 0xff, 0x47, 0xee, 0x99, 0xba, 0xd9, 0xb2, 0xf5, 0x0a, 0x3e, 0x7e, 0x10,
	0x8c, 0xd6, 0xa0, 0xfc, 0x9e, 0xce, 0xec, 0x45, 0xe0, 0xeb, 0x4f, 0xb4, 0x0e, 0xcb, 0x53, 0x3c,
	0x9a, 0xa4, 0x13, 0xf6, 0xed, 0xe2, 0xa8, 0xf4, 0xad, 0xd3, 0xf9, 0xd3, 0x81, 0xf5, 0xc5, 0x96,
	0x52, 0xbf, 0x7e, 0x07, 0xcd, 0x90, 0x32, 0x2a, 0xa2, 0x61, 0xa0, 0xcf, 0x21, 0x35, 0xed, 0x46,
	0xe1, 0x34, 0xbf, 0xb7, 0xb4, 0x89, 0x6a, 0x84, 0xf3, 0x05, 0x7a, 0xb5, 0x30, 0x2d, 0x1d, 0x6c,
	0x6d, 0xeb, 0x16, 0x82, 0x5f, 0xe7, 0x83, 0xd3, 0xe1, 0x85, 0x39, 0xea, 0x04, 0x07, 0x50, 0xcf,
	0xef, 0x90, 0xd4, 0xa9, 0x45, 0x1b, 0x5d, 0x9e, 0x9f, 0x5b, 0x6b, 0xd6, 0xb2, 0xdb, 0xa4, 0xcf,
	0x60, 0xd9, 0x98, 0x17, 0x51, 0x28, 0x0f, 0x08, 0x41, 0x9f, 0x17, 0x47, 0x7e, 0xef, 0xb2, 0x6c,
	0x6d, 0x3e, 0x4e, 0xda, 0x8d, 0x77, 0x76, 0x7e, 0xfd, 0xfb, 0xdf, 0xdf, 0x4b, 0x5b, 0x1d, 0xd7,
	0x9b, 0x1e, 0x16, 0x5e, 0x15, 0xcf, 0xa8, 0xbc, 0x01, 0x21, 0x47, 0xce, 0x5e, 0xff, 0x17, 0x58,
	0x32, 0x9d, 0x4e, 0xa1, 0x96, 0x4d, 0x0f, 0x6d, 0xff, 0xcf, 0x31, 0xb7, 0xda, 0x4f, 0x0b, 0xd2,
	0xda, 0xbb, 0xa6, 0xf6, 0x8b, 0xce, 0xe6, 0xbd, 0xda, 0x5a, 0xe4, 0x65, 0x11, 0x47, 0xce, 0xde,
	0x71, 0xf0, 0xdb, 0xe0, 0xc2, 0x3f, 0x83, 0x2a, 0xa1, 0x37, 0x78, 0x32, 0x52, 0x68, 0x00, 0x68,
	0xc0, 0xda, 0x54, 0x08, 0x2e, 0xda, 0x22, 0xcd, 0xd6, 0x43, 0x5f, 0xc2, 0xcb, 0xd6, 0xee, 0x8e,
	0x47, 0xe8, 0x4d, 0xc4, 0x22, 0xfb, 0x9c, 0x14, 0xdf, 0xcc, 0x13, 0x2d, 0xcf, 0x6a, 0xff, 0xd4,
	0x2c, 0x52, 0xd7, 0x15, 0xf3, 0xd6, 0x7c, 0xfd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x32, 0x43,
	0xa1, 0x6c, 0x65, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MySQLClient is the client API for MySQL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MySQLClient interface {
	// Add MySQL adds MySQL Service and starts several Agents.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "mysqld_exporter", and "qan_mysql_perfschema" agents
	// with provided "pmm_agent_id" and other parameters.
	Add(ctx context.Context, in *AddMySQLRequest, opts ...grpc.CallOption) (*AddMySQLResponse, error)
}

type mySQLClient struct {
	cc *grpc.ClientConn
}

func NewMySQLClient(cc *grpc.ClientConn) MySQLClient {
	return &mySQLClient{cc}
}

func (c *mySQLClient) Add(ctx context.Context, in *AddMySQLRequest, opts ...grpc.CallOption) (*AddMySQLResponse, error) {
	out := new(AddMySQLResponse)
	err := c.cc.Invoke(ctx, "/management.MySQL/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MySQLServer is the server API for MySQL service.
type MySQLServer interface {
	// Add MySQL adds MySQL Service and starts several Agents.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "mysqld_exporter", and "qan_mysql_perfschema" agents
	// with provided "pmm_agent_id" and other parameters.
	Add(context.Context, *AddMySQLRequest) (*AddMySQLResponse, error)
}

func RegisterMySQLServer(s *grpc.Server, srv MySQLServer) {
	s.RegisterService(&_MySQL_serviceDesc, srv)
}

func _MySQL_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMySQLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MySQLServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.MySQL/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MySQLServer).Add(ctx, req.(*AddMySQLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MySQL_serviceDesc = grpc.ServiceDesc{
	ServiceName: "management.MySQL",
	HandlerType: (*MySQLServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _MySQL_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/management.proto",
}

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClient interface {
	// Register do node registration.
	Register(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeResponse, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Register(ctx context.Context, in *RegisterNodeRequest, opts ...grpc.CallOption) (*RegisterNodeResponse, error) {
	out := new(RegisterNodeResponse)
	err := c.cc.Invoke(ctx, "/management.Node/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
type NodeServer interface {
	// Register do node registration.
	Register(context.Context, *RegisterNodeRequest) (*RegisterNodeResponse, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Node/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Register(ctx, req.(*RegisterNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "management.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Node_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/management.proto",
}
