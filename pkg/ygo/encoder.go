package ygo

import (
	"riguz.com/ygo/internal/lib0"
)

type Encode interface {
	Encode(encoder *Encoder)
	EncodeV1() ([]uint8, error)
	EncodeV2() ([]uint8, error)
}

type Encoder interface {
	lib0.Write
	ResetDsCurVal()
	WriteDsClock(clock uint32) error
	WriteDsLen(len uint32) error
	WriteLeftId(id ID) error
	WriteRightId(id ID) error
	WriteClient(client ClientID) error
	WriteInfo(info uint8) error
	WriteParentInfo(isYKey bool) error
	WriteTypeRef(info uint8) error
	WriteLen(len uint32) error
	WriteJson(data any) error
	WriteKey(key *string) error
}

var _ Encoder = &EncoderV1{}

type EncoderV1 struct {
	buf lib0.Write
}

func NewEncoderV1() EncoderV1 {
	w := lib0.NewBufferWrite()
	return EncoderV1{
		buf: &w,
	}
}

func (e *EncoderV1) WriteUint8Array(buf []uint8) error     { return e.buf.WriteUint8Array(buf) }
func (e *EncoderV1) WriteUint8(num uint8) error            { return e.buf.WriteUint8(num) }
func (e *EncoderV1) WriteUint16(num uint16) error          { return e.buf.WriteUint16(num) }
func (e *EncoderV1) WriteUint32(num uint32) error          { return e.buf.WriteUint32(num) }
func (e *EncoderV1) WriteUint32BigEndian(num uint32) error { return e.buf.WriteUint32BigEndian(num) }
func (e *EncoderV1) WriteUint64(num uint64) error          { return e.buf.WriteUint64(num) }
func (e *EncoderV1) WriteFloat32(num float32) error        { return e.buf.WriteFloat32(num) }
func (e *EncoderV1) WriteFloat64(num float64) error        { return e.buf.WriteFloat64(num) }
func (e *EncoderV1) WriteInt64(num int64) error            { return e.buf.WriteInt64(num) }
func (e *EncoderV1) WriteVarUint(num uint) error           { return e.buf.WriteVarUint(num) }
func (e *EncoderV1) WriteVarUint8(num uint8) error         { return e.buf.WriteVarUint8(num) }
func (e *EncoderV1) WriteVarUint16(num uint16) error       { return e.buf.WriteVarUint16(num) }
func (e *EncoderV1) WriteVarUint32(num uint32) error       { return e.buf.WriteVarUint32(num) }
func (e *EncoderV1) WriteVarUint64(num uint64) error       { return e.buf.WriteVarUint64(num) }
func (e *EncoderV1) WriteVarInt(num int) error             { return e.buf.WriteVarInt(num) }
func (e *EncoderV1) WriteVarInt8(num int8) error           { return e.buf.WriteVarInt8(num) }
func (e *EncoderV1) WriteVarInt16(num int16) error         { return e.buf.WriteVarInt16(num) }
func (e *EncoderV1) WriteVarInt32(num int32) error         { return e.buf.WriteVarInt32(num) }
func (e *EncoderV1) WriteVarInt64(num int64) error         { return e.buf.WriteVarInt64(num) }
func (e *EncoderV1) WriteVarUint8Array(buf []uint8) error  { return e.buf.WriteVarUint8Array(buf) }
func (e *EncoderV1) WriteVarString(str *string) error      { return e.buf.WriteVarString(str) }
func (e *EncoderV1) WriteAny(a any) error                  { return e.buf.WriteAny(a) }
func (e *EncoderV1) ToBytes() []uint8                      { return e.buf.ToBytes() }

func (e *EncoderV1) WriteId(id ID) error {
	if err := e.buf.WriteVarInt(int(id.Client)); err != nil {
		return err
	}
	return e.buf.WriteVarInt(int(id.Clock))
}

func (e *EncoderV1) ResetDsCurVal() {
	/* no op */
}

func (e *EncoderV1) WriteDsClock(clock uint32) error {
	return e.buf.WriteVarUint32(clock)
}

func (e *EncoderV1) WriteDsLen(len uint32) error {
	return e.buf.WriteVarUint32(len)
}

func (e *EncoderV1) WriteLeftId(id ID) error {
	return e.WriteId(id)
}

func (e *EncoderV1) WriteRightId(id ID) error {
	return e.WriteId(id)
}

func (e *EncoderV1) WriteClient(client ClientID) error {
	return e.buf.WriteVarInt64(int64(client))
}

func (e *EncoderV1) WriteInfo(info uint8) error {
	return e.buf.WriteUint8(info)
}

func (e *EncoderV1) WriteParentInfo(isYKey bool) error {
	var i uint32 = 0
	if isYKey {
		i = 1
	}
	return e.buf.WriteVarUint32(i)
}

func (e *EncoderV1) WriteTypeRef(info uint8) error {
	return e.buf.WriteUint8(info)
}

func (e *EncoderV1) WriteLen(len uint32) error {
	return e.buf.WriteVarUint32(len)
}

func (e *EncoderV1) WriteJson(data any) error {
	panic("todo")
}

func (e *EncoderV1) WriteKey(key *string) error {
	return e.buf.WriteVarString(key)
}
