// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fire_note.proto

package fire_note

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

type FireNote struct {
	NoteId               int32    `protobuf:"varint,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FireNote) Reset()         { *m = FireNote{} }
func (m *FireNote) String() string { return proto.CompactTextString(m) }
func (*FireNote) ProtoMessage()    {}
func (*FireNote) Descriptor() ([]byte, []int) {
	return fileDescriptor_52ae37f9d381b716, []int{0}
}

func (m *FireNote) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FireNote.Unmarshal(m, b)
}
func (m *FireNote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FireNote.Marshal(b, m, deterministic)
}
func (m *FireNote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FireNote.Merge(m, src)
}
func (m *FireNote) XXX_Size() int {
	return xxx_messageInfo_FireNote.Size(m)
}
func (m *FireNote) XXX_DiscardUnknown() {
	xxx_messageInfo_FireNote.DiscardUnknown(m)
}

var xxx_messageInfo_FireNote proto.InternalMessageInfo

func (m *FireNote) GetNoteId() int32 {
	if m != nil {
		return m.NoteId
	}
	return 0
}

type FireNote_Note struct {
	// noteの配列id
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FireNote_Note) Reset()         { *m = FireNote_Note{} }
func (m *FireNote_Note) String() string { return proto.CompactTextString(m) }
func (*FireNote_Note) ProtoMessage()    {}
func (*FireNote_Note) Descriptor() ([]byte, []int) {
	return fileDescriptor_52ae37f9d381b716, []int{0, 0}
}

func (m *FireNote_Note) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FireNote_Note.Unmarshal(m, b)
}
func (m *FireNote_Note) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FireNote_Note.Marshal(b, m, deterministic)
}
func (m *FireNote_Note) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FireNote_Note.Merge(m, src)
}
func (m *FireNote_Note) XXX_Size() int {
	return xxx_messageInfo_FireNote_Note.Size(m)
}
func (m *FireNote_Note) XXX_DiscardUnknown() {
	xxx_messageInfo_FireNote_Note.DiscardUnknown(m)
}

var xxx_messageInfo_FireNote_Note proto.InternalMessageInfo

func (m *FireNote_Note) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*FireNote)(nil), "FireNote")
	proto.RegisterType((*FireNote_Note)(nil), "FireNote.Note")
}

func init() {
	proto.RegisterFile("fire_note.proto", fileDescriptor_52ae37f9d381b716)
}

var fileDescriptor_52ae37f9d381b716 = []byte{
	// 93 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xcb, 0x2c, 0x4a,
	0x8d, 0xcf, 0xcb, 0x2f, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xb2, 0xe6, 0xe2, 0x70,
	0xcb, 0x2c, 0x4a, 0xf5, 0xcb, 0x2f, 0x49, 0x15, 0x12, 0xe7, 0x62, 0x07, 0xc9, 0xc4, 0x67, 0xa6,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0xb1, 0x81, 0xb8, 0x9e, 0x29, 0x52, 0x62, 0x5c, 0x2c,
	0x60, 0x05, 0x7c, 0x5c, 0x4c, 0x70, 0x39, 0xa6, 0xcc, 0x94, 0x24, 0x36, 0xb0, 0x19, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xe5, 0xbc, 0x60, 0x56, 0x00, 0x00, 0x00,
}