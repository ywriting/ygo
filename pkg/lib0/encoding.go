package lib0

import (
	"bytes"
	"encoding/binary"
)

type Write interface {
	WriteAll(buf []uint8) error
	WriteUint8(num uint8) error
	WriteUint16(num uint16) error
	WriteUint32(num uint32) error
	WriteUint32BigEndian(num uint32) error
	WriteFloat32(num float32) error
	WriteFloat64(num float64) error
	WriteI64(num int64) error
	WriteU64(num uint64) error
	WriteVarUint(num uint) error
	WriteVarU8(num uint8) error
	WriteVarU16(num uint16) error
	WriteVarU32(num uint32) error
	WriteVarU64(num uint64) error
	WriteVarInt(num int) error
	WriteVarI8(num int8) error
	WriteVarI16(num int16) error
	WriteVarI32(num int32) error
	WriteVarI64(num int64) error
	WriteBuf(buf []uint8) error
	WriteString(str *string) error
	ToBytes() []byte
}

type BufferWrite struct {
	buffer *bytes.Buffer
}

func NewBufferWrite() BufferWrite {
	writer := BufferWrite{
		buffer: &bytes.Buffer{},
	}
	return writer
}

func (w *BufferWrite) ToBytes() []byte {
	return w.buffer.Bytes()
}

func (w *BufferWrite) WriteAll(buf []uint8) error {
	_, err := w.buffer.Write(buf)
	return err
}

func (w *BufferWrite) WriteUint8(value uint8) error {
	_, err := w.buffer.Write([]byte{value})
	return err
}

func (w *BufferWrite) WriteUint16(num uint16) error {
	return w.WriteAll([]uint8{uint8(num), uint8(num >> 8)})
}

func (w *BufferWrite) WriteUint32(num uint32) error {
	return w.WriteAll([]uint8{
		uint8(num),
		uint8(num >> 8),
		uint8(num >> 16),
		uint8(num >> 24)})
}

func (w *BufferWrite) WriteUint32BigEndian(num uint32) error {
	return w.WriteAll([]uint8{
		uint8(num >> 24),
		uint8(num >> 16),
		uint8(num >> 8),
		uint8(num)})
}

func (w *BufferWrite) WriteFloat32(num float32) error {
	return binary.Write(w.buffer, binary.BigEndian, num)
}

func (w *BufferWrite) WriteFloat64(num float64) error {
	return binary.Write(w.buffer, binary.BigEndian, num)
}

func (w *BufferWrite) WriteVarUint(num uint) error {
	return w.WriteU64(uint64(num))
}

func (w *BufferWrite) WriteVarU8(num uint8) error {
	return w.WriteUint32(uint32(num))
}

func (w *BufferWrite) WriteVarU16(num uint16) error {
	return w.WriteUint32(uint32(num))
}

func (w *BufferWrite) WriteVarU32(num uint32) error {
	return w.WriteVarU64(uint64(num))
}

func (w *BufferWrite) WriteVarU64(num uint64) error {
	for num >= 0b1000_0000 {
		var b uint8 = uint8(num&0b0111_1111) | 0b1000_0000
		err := w.WriteUint8(b)
		if err != nil {
			return err
		}
		num = num >> 7
	}
	return w.WriteUint8(uint8(num & 0b0111_1111))
}

func (w *BufferWrite) WriteVarInt(num int) error {
	return w.WriteVarI64(int64(num))
}

func (w *BufferWrite) WriteVarI8(num int8) error {
	return w.WriteVarI64(int64(num))
}

func (w *BufferWrite) WriteVarI16(num int16) error {
	return w.WriteVarI64(int64(num))
}

func (w *BufferWrite) WriteVarI32(num int32) error {
	return w.WriteVarI64(int64(num))
}

func (w *BufferWrite) WriteVarI64(num int64) error {
	isNegative := num < 0
	if isNegative {
		num = -num
	}
	firstByte := uint8(int64(0b0011_1111) & num)
	if num > 0b0011_1111 {
		firstByte |= uint8(0b1000_0000) // continue reading or not
	}
	if isNegative {
		firstByte |= uint8(0b0100_0000) // is negative or not
	}
	if err := w.WriteUint8(firstByte); err != nil {
		return err
	}

	num >>= 6
	for num > 0 {
		var b uint8 = 0
		if num > int64(0b0111_1111) {
			b |= uint8(0b1000_0000)
		}
		b |= uint8(int64(0b0111_111) & num)
		if err := w.WriteUint8(b); err != nil {
			return err
		}
	}
	return nil
}

func (w *BufferWrite) WriteVarBuf(buf []uint8) error {
	if err := w.WriteVarUint(uint(len(buf))); err != nil {
		return err
	}
	return w.WriteAll(buf)
}

func (w *BufferWrite) WriteString(str *string) error {
	return w.WriteVarBuf([]byte(*str))
}

func (w *BufferWrite) WriteI64(num int64) error {
	return binary.Write(w.buffer, binary.BigEndian, num)
}

func (w *BufferWrite) WriteU64(num uint64) error {
	return binary.Write(w.buffer, binary.BigEndian, num)
}
