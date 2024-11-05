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
