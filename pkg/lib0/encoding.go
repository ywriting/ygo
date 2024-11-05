package lib0

import (
	"bytes"
	"encoding/binary"
)

type Write interface {
	WriteUint8Array(buf []uint8) error
	WriteUint8(num uint8) error
	WriteUint16(num uint16) error
	WriteUint32(num uint32) error
	WriteUint32BigEndian(num uint32) error
	WriteUint64(num uint64) error
	WriteFloat32(num float32) error
	WriteFloat64(num float64) error
	WriteInt64(num int64) error
	WriteVarUint(num uint) error
	WriteVarUint8(num uint8) error
	WriteVarUint16(num uint16) error
	WriteVarUint32(num uint32) error
	WriteVarUint64(num uint64) error
	WriteVarInt(num int) error
	WriteVarInt8(num int8) error
	WriteVarInt16(num int16) error
	WriteVarInt32(num int32) error
	WriteVarInt64(num int64) error
	WriteVarUint8Array(buf []uint8) error
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

func (w *BufferWrite) WriteUint8Array(buf []uint8) error {
	_, err := w.buffer.Write(buf)
	return err
}

func (w *BufferWrite) WriteUint8(value uint8) error {
	_, err := w.buffer.Write([]byte{value})
	return err
}

func (w *BufferWrite) WriteUint16(num uint16) error {
	return w.WriteUint8Array([]uint8{uint8(num), uint8(num >> 8)})
}

func (w *BufferWrite) WriteUint32(num uint32) error {
	return w.WriteUint8Array([]uint8{
		uint8(num),
		uint8(num >> 8),
		uint8(num >> 16),
		uint8(num >> 24)})
}

func (w *BufferWrite) WriteUint32BigEndian(num uint32) error {
	return w.WriteUint8Array([]uint8{
		uint8(num >> 24),
		uint8(num >> 16),
		uint8(num >> 8),
		uint8(num)})
}

func (w *BufferWrite) WriteUint64(num uint64) error {
	return w.WriteUint8Array([]uint8{
		uint8(num >> 56),
		uint8(num >> 48),
		uint8(num >> 40),
		uint8(num >> 32),
		uint8(num >> 24),
		uint8(num >> 16),
		uint8(num >> 8),
		uint8(num),
	})
}

func (w *BufferWrite) WriteInt64(num int64) error {
	return w.WriteUint8Array([]uint8{
		uint8(num >> 56),
		uint8(num >> 48),
		uint8(num >> 40),
		uint8(num >> 32),
		uint8(num >> 24),
		uint8(num >> 16),
		uint8(num >> 8),
		uint8(num),
	})
}

func (w *BufferWrite) WriteFloat32(num float32) error {
	return binary.Write(w.buffer, binary.BigEndian, num)
}

func (w *BufferWrite) WriteFloat64(num float64) error {
	return binary.Write(w.buffer, binary.BigEndian, num)
}

func (w *BufferWrite) WriteVarUint(num uint) error {
	return w.WriteVarUint64(uint64(num))
}

func (w *BufferWrite) WriteVarUint8(num uint8) error {
	return w.WriteVarUint64(uint64(num))
}

func (w *BufferWrite) WriteVarUint16(num uint16) error {
	return w.WriteVarUint64(uint64(num))
}

func (w *BufferWrite) WriteVarUint32(num uint32) error {
	return w.WriteVarUint64(uint64(num))
}

func (w *BufferWrite) WriteVarUint64(num uint64) error {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(buf, uint64(num))
	return w.WriteUint8Array(buf[:n])
}

func (w *BufferWrite) WriteVarInt(num int) error {
	return w.WriteVarInt64(int64(num))
}

func (w *BufferWrite) WriteVarInt8(num int8) error {
	return w.WriteVarInt64(int64(num))
}

func (w *BufferWrite) WriteVarInt16(num int16) error {
	return w.WriteVarInt64(int64(num))
}

func (w *BufferWrite) WriteVarInt32(num int32) error {
	return w.WriteVarInt64(int64(num))
}

func (w *BufferWrite) WriteVarInt64(num int64) error {
	isNegative := num < 0
	if isNegative {
		num = -num
	}
	firstByte := uint8(int64(0b0011_1111) & num)
	if num > int64(0b0011_1111) {
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
		b |= uint8(int64(0b0111_1111) & num)
		if err := w.WriteUint8(b); err != nil {
			return err
		}
		num >>= 7
	}
	return nil
}

func (w *BufferWrite) WriteVarUint8Array(buf []uint8) error {
	if err := w.WriteVarUint(uint(len(buf))); err != nil {
		return err
	}
	return w.WriteUint8Array(buf)
}

func (w *BufferWrite) WriteString(str *string) error {
	return w.WriteVarUint8Array([]byte(*str))
}
