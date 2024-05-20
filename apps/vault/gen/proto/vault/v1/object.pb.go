// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: proto/vault/v1/object.proto

package vaultv1

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

type Algorithm int32

const (
	Algorithm_AES_256_GCM Algorithm = 0
)

// Enum value maps for Algorithm.
var (
	Algorithm_name = map[int32]string{
		0: "AES_256_GCM",
	}
	Algorithm_value = map[string]int32{
		"AES_256_GCM": 0,
	}
)

func (x Algorithm) Enum() *Algorithm {
	p := new(Algorithm)
	*p = x
	return p
}

func (x Algorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Algorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_vault_v1_object_proto_enumTypes[0].Descriptor()
}

func (Algorithm) Type() protoreflect.EnumType {
	return &file_proto_vault_v1_object_proto_enumTypes[0]
}

func (x Algorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Algorithm.Descriptor instead.
func (Algorithm) EnumDescriptor() ([]byte, []int) {
	return file_proto_vault_v1_object_proto_rawDescGZIP(), []int{0}
}

type DataEncryptionKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Linux milliseconds since epoch
	CreatedAt int64  `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Key       []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DataEncryptionKey) Reset() {
	*x = DataEncryptionKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vault_v1_object_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataEncryptionKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataEncryptionKey) ProtoMessage() {}

func (x *DataEncryptionKey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_v1_object_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataEncryptionKey.ProtoReflect.Descriptor instead.
func (*DataEncryptionKey) Descriptor() ([]byte, []int) {
	return file_proto_vault_v1_object_proto_rawDescGZIP(), []int{0}
}

func (x *DataEncryptionKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataEncryptionKey) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *DataEncryptionKey) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

// This is stored in the database in whatever format the database uses
type EncryptedDataEncryptionKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Linux milliseconds since epoch
	CreatedAt int64      `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Encrypted *Encrypted `protobuf:"bytes,3,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
}

func (x *EncryptedDataEncryptionKey) Reset() {
	*x = EncryptedDataEncryptionKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vault_v1_object_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptedDataEncryptionKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptedDataEncryptionKey) ProtoMessage() {}

func (x *EncryptedDataEncryptionKey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_v1_object_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptedDataEncryptionKey.ProtoReflect.Descriptor instead.
func (*EncryptedDataEncryptionKey) Descriptor() ([]byte, []int) {
	return file_proto_vault_v1_object_proto_rawDescGZIP(), []int{1}
}

func (x *EncryptedDataEncryptionKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EncryptedDataEncryptionKey) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *EncryptedDataEncryptionKey) GetEncrypted() *Encrypted {
	if x != nil {
		return x.Encrypted
	}
	return nil
}

// KeyEncryptionKey is a key used to encrypt data encryption keys
type KeyEncryptionKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt int64  `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Key       []byte `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *KeyEncryptionKey) Reset() {
	*x = KeyEncryptionKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vault_v1_object_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyEncryptionKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyEncryptionKey) ProtoMessage() {}

func (x *KeyEncryptionKey) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_v1_object_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyEncryptionKey.ProtoReflect.Descriptor instead.
func (*KeyEncryptionKey) Descriptor() ([]byte, []int) {
	return file_proto_vault_v1_object_proto_rawDescGZIP(), []int{2}
}

func (x *KeyEncryptionKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *KeyEncryptionKey) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *KeyEncryptionKey) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

// Encrypted contains the output of the encryption and all of the metadata required to decrypt it
type Encrypted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Algorithm  Algorithm `protobuf:"varint,1,opt,name=algorithm,proto3,enum=vault.v1.Algorithm" json:"algorithm,omitempty"`
	Nonce      []byte    `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Ciphertext []byte    `protobuf:"bytes,3,opt,name=ciphertext,proto3" json:"ciphertext,omitempty"`
	// key id of the key that encrypted this data
	EncryptionKeyId string `protobuf:"bytes,4,opt,name=encryption_key_id,json=encryptionKeyId,proto3" json:"encryption_key_id,omitempty"`
	// time of encryption
	// we can use this later to figure out if a piece of data should be re-encrypted
	Time int64 `protobuf:"varint,5,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *Encrypted) Reset() {
	*x = Encrypted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_vault_v1_object_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Encrypted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Encrypted) ProtoMessage() {}

func (x *Encrypted) ProtoReflect() protoreflect.Message {
	mi := &file_proto_vault_v1_object_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Encrypted.ProtoReflect.Descriptor instead.
func (*Encrypted) Descriptor() ([]byte, []int) {
	return file_proto_vault_v1_object_proto_rawDescGZIP(), []int{3}
}

func (x *Encrypted) GetAlgorithm() Algorithm {
	if x != nil {
		return x.Algorithm
	}
	return Algorithm_AES_256_GCM
}

func (x *Encrypted) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *Encrypted) GetCiphertext() []byte {
	if x != nil {
		return x.Ciphertext
	}
	return nil
}

func (x *Encrypted) GetEncryptionKeyId() string {
	if x != nil {
		return x.EncryptionKeyId
	}
	return ""
}

func (x *Encrypted) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

var File_proto_vault_v1_object_proto protoreflect.FileDescriptor

var file_proto_vault_v1_object_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x76,
	0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x54, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x7e, 0x0a,
	0x1a, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x65, 0x64, 0x52, 0x09, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x22, 0x53, 0x0a,
	0x10, 0x4b, 0x65, 0x79, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4b, 0x65,
	0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x22, 0xb4, 0x01, 0x0a, 0x09, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64,
	0x12, 0x31, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x69, 0x70,
	0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x63,
	0x69, 0x70, 0x68, 0x65, 0x72, 0x74, 0x65, 0x78, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x2a, 0x1c, 0x0a, 0x09, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x45, 0x53, 0x5f, 0x32, 0x35,
	0x36, 0x5f, 0x47, 0x43, 0x4d, 0x10, 0x00, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x6e, 0x6b, 0x65, 0x79, 0x65, 0x64, 0x2f, 0x75, 0x6e,
	0x6b, 0x65, 0x79, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x76,
	0x31, 0x3b, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_vault_v1_object_proto_rawDescOnce sync.Once
	file_proto_vault_v1_object_proto_rawDescData = file_proto_vault_v1_object_proto_rawDesc
)

func file_proto_vault_v1_object_proto_rawDescGZIP() []byte {
	file_proto_vault_v1_object_proto_rawDescOnce.Do(func() {
		file_proto_vault_v1_object_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_vault_v1_object_proto_rawDescData)
	})
	return file_proto_vault_v1_object_proto_rawDescData
}

var file_proto_vault_v1_object_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_vault_v1_object_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_vault_v1_object_proto_goTypes = []interface{}{
	(Algorithm)(0),                     // 0: vault.v1.Algorithm
	(*DataEncryptionKey)(nil),          // 1: vault.v1.DataEncryptionKey
	(*EncryptedDataEncryptionKey)(nil), // 2: vault.v1.EncryptedDataEncryptionKey
	(*KeyEncryptionKey)(nil),           // 3: vault.v1.KeyEncryptionKey
	(*Encrypted)(nil),                  // 4: vault.v1.Encrypted
}
var file_proto_vault_v1_object_proto_depIdxs = []int32{
	4, // 0: vault.v1.EncryptedDataEncryptionKey.encrypted:type_name -> vault.v1.Encrypted
	0, // 1: vault.v1.Encrypted.algorithm:type_name -> vault.v1.Algorithm
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_vault_v1_object_proto_init() }
func file_proto_vault_v1_object_proto_init() {
	if File_proto_vault_v1_object_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_vault_v1_object_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataEncryptionKey); i {
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
		file_proto_vault_v1_object_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptedDataEncryptionKey); i {
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
		file_proto_vault_v1_object_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyEncryptionKey); i {
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
		file_proto_vault_v1_object_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Encrypted); i {
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
			RawDescriptor: file_proto_vault_v1_object_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_vault_v1_object_proto_goTypes,
		DependencyIndexes: file_proto_vault_v1_object_proto_depIdxs,
		EnumInfos:         file_proto_vault_v1_object_proto_enumTypes,
		MessageInfos:      file_proto_vault_v1_object_proto_msgTypes,
	}.Build()
	File_proto_vault_v1_object_proto = out.File
	file_proto_vault_v1_object_proto_rawDesc = nil
	file_proto_vault_v1_object_proto_goTypes = nil
	file_proto_vault_v1_object_proto_depIdxs = nil
}
