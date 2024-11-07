package lib0_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"gotest.tools/assert"
	"riguz.com/ygo/pkg/lib0"
)

/*
	import * as enc from "lib0/encoding";

	const toHexString = (bytes) => {
	return Array.from(bytes, (byte) => {
		return ("0" + (byte & 0xff).toString(16)).slice(-2);
	}).join("");
	};

	var encoder = enc.createEncoder();
	enc.writeUint8(encoder, 1);
	var result = enc.toUint8Array(encoder);
	console.log(toHexString(result));
*/

func TestWrite_uint8Array(t *testing.T) {
	w := lib0.NewBufferWrite()

	err := w.WriteUint8Array([]byte{0, 1, 128, 254})

	assert.NilError(t, err)
	assert.Equal(t, "000180fe", hex.EncodeToString(w.ToBytes()))
}

func TestWrite_varUint8Array(t *testing.T) {
	w := lib0.NewBufferWrite()

	err := w.WriteVarUint8Array([]byte{0, 1, 128, 254})

	assert.NilError(t, err)
	assert.Equal(t, "04000180fe", hex.EncodeToString(w.ToBytes()))
}

func TestWrite_uint8(t *testing.T) {
	var tests = []struct {
		number   uint8
		expected string
	}{
		{1, "01"},
		{123, "7b"},
		{255, "ff"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("write uint8:%d", tt.number), func(t *testing.T) {
			w := lib0.NewBufferWrite()

			err := w.WriteUint8(tt.number)

			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(w.ToBytes()))
		})
	}
}

func TestWrite_uint16(t *testing.T) {
	var tests = []struct {
		number   uint16
		expected string
	}{
		{0, "0000"},
		{1, "0100"},
		{255, "ff00"},
		{65535, "ffff"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("write uint16:%d", tt.number), func(t *testing.T) {
			w := lib0.NewBufferWrite()

			err := w.WriteUint16(tt.number)

			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(w.ToBytes()))
		})
	}
}

func TestWrite_uint32(t *testing.T) {
	var tests = []struct {
		number   uint32
		expected string
	}{
		{0, "00000000"},
		{1, "01000000"},
		{255, "ff000000"},
		{65535, "ffff0000"},
		{4294967295, "ffffffff"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("write uint32:%d", tt.number), func(t *testing.T) {
			w := lib0.NewBufferWrite()

			err := w.WriteUint32(tt.number)

			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(w.ToBytes()))
		})
	}
}

func TestWrite_uint32_be(t *testing.T) {
	var tests = []struct {
		number   uint32
		expected string
	}{
		{0, "00000000"},
		{1, "00000001"},
		{255, "000000ff"},
		{65535, "0000ffff"},
		{4294967295, "ffffffff"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("write uint32be:%d", tt.number), func(t *testing.T) {
			w := lib0.NewBufferWrite()

			err := w.WriteUint32BigEndian(tt.number)

			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(w.ToBytes()))
		})
	}
}

func TestWrite_float32(t *testing.T) {
	var tests = []struct {
		number   float32
		expected string
	}{
		{0, "00000000"},
		{1, "3f800000"},
		{255, "437f0000"},
		{65535, "477fff00"},
		{4294967295, "4f800000"},
		{123.45, "42f6e666"},
		{982990.1, "496ffce2"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("write float32:%f", tt.number), func(t *testing.T) {
			w := lib0.NewBufferWrite()

			err := w.WriteFloat32(tt.number)

			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(w.ToBytes()))
		})
	}
}

func TestWrite_float64(t *testing.T) {
	var tests = []struct {
		number   float64
		expected string
	}{
		{0, "0000000000000000"},
		{1, "3ff0000000000000"},
		{255, "406fe00000000000"},
		{65535, "40efffe000000000"},
		{4294967295, "41efffffffe00000"},
		{18446744073709552000, "43f0000000000000"},
		{123.45, "405edccccccccccd"},
		{982990.1, "412dff9c33333333"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("write float32:%f", tt.number), func(t *testing.T) {
			w := lib0.NewBufferWrite()

			err := w.WriteFloat64(tt.number)

			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(w.ToBytes()))
		})
	}
}

func TestWrite_uint64(t *testing.T) {
	var tests = []struct {
		number   uint64
		expected string
	}{
		{0, "0000000000000000"},
		{1, "0000000000000001"},
		{255, "00000000000000ff"},
		{65535, "000000000000ffff"},
		{4294967295, "00000000ffffffff"},
		{18446744073709551615, "ffffffffffffffff"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("write uint64:%d", tt.number), func(t *testing.T) {
			w := lib0.NewBufferWrite()

			err := w.WriteUint64(tt.number)

			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(w.ToBytes()))
		})
	}
}

func TestWrite_int64(t *testing.T) {
	var tests = []struct {
		number   int64
		expected string
	}{
		{0, "0000000000000000"},
		{1, "0000000000000001"},
		{-1, "ffffffffffffffff"},
		{255, "00000000000000ff"},
		{65535, "000000000000ffff"},
		{-65535, "ffffffffffff0001"},
		{-9223372036854775808, "8000000000000000"},
		{9223372036854775807, "7fffffffffffffff"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("write int64:%d", tt.number), func(t *testing.T) {
			w := lib0.NewBufferWrite()

			err := w.WriteInt64(tt.number)

			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(w.ToBytes()))
		})
	}
}

func TestWrite_varUint(t *testing.T) {
	var tests = []struct {
		number   uint
		expected string
	}{
		{0, "00"},
		{1, "01"},
		{255, "ff01"},
		{65535, "ffff03"},
		{4294967295, "ffffffff0f"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("writeVarUint:%d", tt.number), func(t *testing.T) {
			w := lib0.NewBufferWrite()

			err := w.WriteVarUint(tt.number)

			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(w.ToBytes()))
		})
	}
}

func TestWrite_varUint8(t *testing.T) {
	var tests = []struct {
		number   uint8
		expected string
	}{
		{0, "00"},
		{1, "01"},
		{255, "ff01"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("writeVarUint8:%d", tt.number), func(t *testing.T) {
			w := lib0.NewBufferWrite()

			err := w.WriteVarUint8(tt.number)

			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(w.ToBytes()))
		})
	}
}

func TestWrite_varUint16(t *testing.T) {
	var tests = []struct {
		number   uint16
		expected string
	}{
		{0, "00"},
		{1, "01"},
		{255, "ff01"},
		{65535, "ffff03"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("writeVarUint16:%d", tt.number), func(t *testing.T) {
			w := lib0.NewBufferWrite()

			err := w.WriteVarUint16(tt.number)

			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(w.ToBytes()))
		})
	}
}

func TestWrite_varUint32(t *testing.T) {
	var tests = []struct {
		number   uint32
		expected string
	}{
		{0, "00"},
		{1, "01"},
		{255, "ff01"},
		{65535, "ffff03"},
		{4294967295, "ffffffff0f"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("writeVarUint32:%d", tt.number), func(t *testing.T) {
			w := lib0.NewBufferWrite()

			err := w.WriteVarUint32(tt.number)

			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(w.ToBytes()))
		})
	}
}

func TestWrite_varUint64(t *testing.T) {
	var tests = []struct {
		number   uint64
		expected string
	}{
		{0, "00"},
		{1, "01"},
		{255, "ff01"},
		{65535, "ffff03"},
		{4294967295, "ffffffff0f"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("writeVarUint64:%d", tt.number), func(t *testing.T) {
			w := lib0.NewBufferWrite()

			err := w.WriteVarUint64(tt.number)

			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(w.ToBytes()))
		})
	}
}

func TestWrite_varInt(t *testing.T) {
	var tests = []struct {
		number   int
		expected string
	}{
		{0, "00"},
		{1, "01"},
		{255, "bf03"},
		{65535, "bfff07"},
		{4294967295, "bfffffff1f"},
		{-1, "41"},
		{-255, "ff03"},
		{-65535, "ffff07"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("writeVarInt:%d", tt.number), func(t *testing.T) {
			w := lib0.NewBufferWrite()

			err := w.WriteVarInt(tt.number)

			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(w.ToBytes()))
		})
	}
}

func TestWrite_varInt8(t *testing.T) {
	var tests = []struct {
		number   int8
		expected string
	}{
		{0, "00"},
		{1, "01"},
		{-1, "41"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("writeVarInt8:%d", tt.number), func(t *testing.T) {
			w := lib0.NewBufferWrite()

			err := w.WriteVarInt8(tt.number)

			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(w.ToBytes()))
		})
	}
}

func TestWrite_varInt16(t *testing.T) {
	var tests = []struct {
		number   int16
		expected string
	}{
		{0, "00"},
		{1, "01"},
		{255, "bf03"},
		{-1, "41"},
		{-255, "ff03"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("writeVarInt16:%d", tt.number), func(t *testing.T) {
			w := lib0.NewBufferWrite()

			err := w.WriteVarInt16(tt.number)

			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(w.ToBytes()))
		})
	}
}

func TestWrite_varInt32(t *testing.T) {
	var tests = []struct {
		number   int32
		expected string
	}{
		{0, "00"},
		{1, "01"},
		{255, "bf03"},
		{65535, "bfff07"},
		{-1, "41"},
		{-255, "ff03"},
		{-65535, "ffff07"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("writeVarInt32:%d", tt.number), func(t *testing.T) {
			w := lib0.NewBufferWrite()

			err := w.WriteVarInt32(tt.number)

			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(w.ToBytes()))
		})
	}
}

func TestWrite_varInt64(t *testing.T) {
	var tests = []struct {
		number   int64
		expected string
	}{
		{0, "00"},
		{1, "01"},
		{255, "bf03"},
		{65535, "bfff07"},
		{4294967295, "bfffffff1f"},
		{-1, "41"},
		{-255, "ff03"},
		{-65535, "ffff07"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("writeVarInt64:%d", tt.number), func(t *testing.T) {
			w := lib0.NewBufferWrite()

			err := w.WriteVarInt64(tt.number)

			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(w.ToBytes()))
		})
	}
}

func TestWrite_varString(t *testing.T) {
	var tests = []struct {
		str      string
		expected string
	}{
		{"Hello World!", "0c48656c6c6f20576f726c6421"},
		{"", "00"},
		{"你好，世界！", "12e4bda0e5a5bdefbc8ce4b896e7958cefbc81"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("write var string:%s", tt.str), func(t *testing.T) {
			w := lib0.NewBufferWrite()

			err := w.WriteVarString(&tt.str)

			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(w.ToBytes()))
		})
	}
}
