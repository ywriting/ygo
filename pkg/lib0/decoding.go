package lib0

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

type Read interface {
	ReadUint8Array(len uint) ([]uint8, error)
	ReadUint8() (uint8, error)
	ReadUint16() (uint16, error)
	ReadUint32() (uint32, error)
	ReadUint32BigEndian(uint32, error)
	ReadUint64() (uint64, error)
	ReadFloat32() (float32, error)
	ReadFloat64() (float64, error)
	ReadInt64() (int64, error)
	ReadVarUint8Array() ([]uint8, error)
	ReadVarUint() (uint64, error)
	ReadVarInt() (int64, error)
	ReadVarString() (string, error)
}

type BufferRead struct {
	reader *bufio.Reader
}

func NewBufferRead(reader io.Reader) BufferRead {
	return BufferRead{
		reader: bufio.NewReader(reader),
	}
}

func (r *BufferRead) ReadUint8Array(len uint) ([]uint8, error) {
	buf := make([]uint8, len)
	n, err := r.reader.Read(buf)
	if err != nil {
		return nil, err
	}
	if n != int(len) {
		return nil, fmt.Errorf("unexpected EOF: expected to read %v but only got %v", len, n)
	}
	return buf, nil
}

func (r *BufferRead) ReadUint8() (uint8, error) {
	buf, err := r.ReadUint8Array(1)
	if err != nil {
		return 0, err
	}
	return buf[0], nil
}

func (r *BufferRead) ReadUint16() (uint16, error) {
	var value uint16
	err := binary.Read(r.reader, binary.LittleEndian, &value)
	return value, err
}

func (r *BufferRead) ReadUint32() (uint32, error) {
	var value uint32
	err := binary.Read(r.reader, binary.LittleEndian, &value)
	return value, err
}

func (r *BufferRead) ReadUint32BigEndian() (uint32, error) {
	var value uint32
	err := binary.Read(r.reader, binary.BigEndian, &value)
	return value, err
}

func (r *BufferRead) ReadUint64() (uint64, error) {
	var value uint64
	err := binary.Read(r.reader, binary.BigEndian, &value)
	return value, err
}

func (r *BufferRead) ReadFloat32() (float32, error) {
	var value float32
	err := binary.Read(r.reader, binary.BigEndian, &value)
	return value, err
}

func (r *BufferRead) ReadFloat64() (float64, error) {
	var value float64
	err := binary.Read(r.reader, binary.BigEndian, &value)
	return value, err
}

func (r *BufferRead) ReadInt64() (int64, error) {
	var value int64
	err := binary.Read(r.reader, binary.BigEndian, &value)
	return value, err
}

func (r *BufferRead) ReadVarUint() (uint64, error) {
	return binary.ReadUvarint(r.reader)
}

func (r *BufferRead) ReadVarUint8Array() ([]uint8, error) {
	len, err := r.ReadVarUint()
	if err != nil {
		return nil, err
	}
	return r.ReadUint8Array(uint(len))
}

func (r *BufferRead) ReadVarInt() (int64, error) {
	firstByte, err := r.ReadUint8()
	if err != nil {
		return 0, err
	}
	var num int64 = int64(firstByte & uint8(0b0011_1111))
	isNegative := false
	if firstByte&uint8(0b0100_0000) > 0 {
		isNegative = true
	}
	if firstByte&uint8(0b1000_0000) == 0 {
		if isNegative {
			return -num, nil
		} else {
			return num, nil
		}
	}
	len := 6
	for {
		byte, err := r.ReadUint8()
		if err != nil {
			return 0, err
		}
		num |= (int64(byte) & int64(0b0111_1111)) << len
		len += 7
		if byte < uint8(0b1000_0000) {
			if isNegative {
				return -num, nil
			} else {
				return num, nil
			}
		}
		if len > 70 {
			return 0, errors.New("varint size exceeded length of 70 bits")
		}
	}
}

func (r *BufferRead) ReadVarString() (string, error) {
	buf, err := r.ReadVarUint8Array()
	if err != nil {
		return "", err
	}
	return string(buf), nil
}
