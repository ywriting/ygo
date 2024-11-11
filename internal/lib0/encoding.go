package lib0

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
)

type Undefined struct{}

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
	WriteVarString(str *string) error
	WriteAny(a any) error
	ToBytes() []byte
}

var _ Write = &BufferWrite{}

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
	return binary.Write(w.buffer, binary.LittleEndian, num)
}

func (w *BufferWrite) WriteUint32(num uint32) error {
	return binary.Write(w.buffer, binary.LittleEndian, num)
}

func (w *BufferWrite) WriteUint32BigEndian(num uint32) error {
	return binary.Write(w.buffer, binary.BigEndian, num)
}

func (w *BufferWrite) WriteUint64(num uint64) error {
	return binary.Write(w.buffer, binary.BigEndian, num)
}

func (w *BufferWrite) WriteInt64(num int64) error {
	return binary.Write(w.buffer, binary.BigEndian, num)
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

func (w *BufferWrite) WriteVarString(str *string) error {
	return w.WriteVarUint8Array([]byte(*str))
}

func (w *BufferWrite) WriteAny(a any) error {
	switch t := a.(type) {
	case string:
		return w.writeString(t)
	case float32:
		return w.writeFloat32(t)
	case float64:
		return w.writeFloat64(t)
	case int8:
		return w.writeVarInt(int32(t))
	case int16:
		return w.writeVarInt(int32(t))
	case int32:
		return w.writeVarInt(t)
	case int:
		if t <= math.MaxInt32 && t >= math.MinInt32 {
			return w.writeVarInt(int32(t))
		} else {
			return w.writeBigInt(int64(t))
		}
	case int64:
		if t <= math.MaxInt32 && t >= math.MinInt32 {
			return w.writeVarInt(int32(t))
		} else {
			return w.writeBigInt(int64(t))
		}
	case bool:
		return w.writeBool(t)
	case []uint8:
		return w.writeUint8Array(t)
	case []any:
		return w.writeArray(t)
	case map[string]any:
		return w.writeObject(t)
	case nil:
		return w.WriteUint8(126)
	case Undefined:
		return w.WriteUint8(127)
	default:
		return fmt.Errorf("unrecognized any payload type:%v", reflect.TypeOf(a))
	}
}

func (w *BufferWrite) writeVarInt(num int32) error {
	if err := w.WriteUint8(125); err != nil {
		return err
	}
	return w.WriteVarInt32(num)
}

func (w *BufferWrite) writeFloat32(num float32) error {
	if err := w.WriteUint8(124); err != nil {
		return err
	}
	return w.WriteFloat32(num)
}

var F64_MAX_SAFE_INTEGER float64 = math.Pow(2, 53) - 1
var F64_MIN_SAFE_INTEGER float64 = -F64_MAX_SAFE_INTEGER

func (w *BufferWrite) writeFloat64(num float64) error {
	truncated := math.Trunc(num)
	if truncated == num &&
		truncated <= F64_MAX_SAFE_INTEGER &&
		truncated >= F64_MIN_SAFE_INTEGER {
		return w.WriteAny(int64(truncated))
	} else if float64(float32(num)) == num {
		return w.writeFloat32(float32(num))
	} else {
		if err := w.WriteUint8(123); err != nil {
			return err
		}
		return w.WriteFloat64(num)
	}
}

func (w *BufferWrite) writeBigInt(num int64) error {
	if err := w.WriteUint8(122); err != nil {
		return err
	}
	return w.WriteInt64(num)
}

func (w *BufferWrite) writeBool(val bool) error {
	var t uint8 = 121
	if val {
		t = 120
	}

	return w.WriteUint8(t)
}

func (w *BufferWrite) writeString(str string) error {
	if err := w.WriteUint8(119); err != nil {
		return err
	}
	return w.WriteVarString(&str)
}

func (w *BufferWrite) writeObject(obj map[string]any) error {
	if err := w.WriteUint8(118); err != nil {
		return err
	}
	if err := w.WriteVarUint(uint(len(obj))); err != nil {
		return err
	}
	for k, v := range obj {
		if err := w.WriteVarString(&k); err != nil {
			return err
		}
		if err := w.WriteAny(v); err != nil {
			return err
		}
	}
	return nil
}

func (w *BufferWrite) writeArray(arr []any) error {
	if err := w.WriteUint8(117); err != nil {
		return err
	}
	if err := w.WriteVarUint(uint(len(arr))); err != nil {
		return err
	}
	for _, a := range arr {
		if err := w.WriteAny(a); err != nil {
			return err
		}
	}
	return nil
}

func (w *BufferWrite) writeUint8Array(arr []byte) error {
	if err := w.WriteUint8(116); err != nil {
		return err
	}
	return w.WriteVarUint8Array(arr)
}
