package ygo_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"gotest.tools/assert"
	"riguz.com/ygo/pkg/ygo"
)

func TestIntDiffOptRleEncoder_write(t *testing.T) {
	var tests = []struct {
		numbers  []uint32
		expected string
	}{
		{[]uint32{}, ""},
		{[]uint32{1, 2, 3, 2}, "030142"},
		{[]uint32{1, 2, 3, 2, 2, 2, 2, 2, 2, 2}, "0301420104"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("write IntDiffOptRleEncoder:%s", tt.expected), func(t *testing.T) {
			encoder := ygo.NewIntDiffOptRleEncoder()
			for _, e := range tt.numbers {
				err := encoder.Write(e)
				assert.NilError(t, err)
			}
			result, err := encoder.ToBytes()
			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(result))
		})
	}
}

func TestUIntOptRleEncoder_write(t *testing.T) {
	var tests = []struct {
		numbers  []uint64
		expected string
	}{
		{[]uint64{}, ""},
		{[]uint64{1, 2, 3, 3, 3}, "01024301"},
		{[]uint64{1, 2, 3, 65535, 18273719133}, "010203bfff079dcd96938801"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("write UIntOptRleEncoder:%s", tt.expected), func(t *testing.T) {
			encoder := ygo.NewUIntOptRleEncoder()
			for _, e := range tt.numbers {
				err := encoder.Write(e)
				assert.NilError(t, err)
			}
			result, err := encoder.ToBytes()
			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(result))
		})
	}
}

func TestRleEncoder_write(t *testing.T) {
	var tests = []struct {
		numbers  []uint8
		expected string
	}{
		{[]uint8{}, ""},
		{[]uint8{1, 1, 1, 7}, "010207"},
		{[]uint8{1, 2, 3, 255}, "010002000300ff"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("write RleEncoder:%s", tt.expected), func(t *testing.T) {
			encoder := ygo.NewRleEncoder()
			for _, e := range tt.numbers {
				err := encoder.Write(e)
				assert.NilError(t, err)
			}
			result := encoder.ToBytes()
			assert.Equal(t, tt.expected, hex.EncodeToString(result))
		})
	}
}

func TestStringEncoder_write(t *testing.T) {
	var tests = []struct {
		str      string
		expected string
	}{
		{"", "0000"},
		{"abc", "0361626303"},
		{"Hello world!", "0c48656c6c6f20776f726c64210c"},
		{"𐐷", "04f09090b702"},
		{"Hello,中国！𐐷𐐷𐐷", "1b48656c6c6f2ce4b8ade59bbdefbc81f09090b7f09090b7f09090b70f"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("write StringEncoder:%s", tt.expected), func(t *testing.T) {
			encoder := ygo.NewStringEncoder()
			err := encoder.Write(&tt.str)
			assert.NilError(t, err)
			result, err := encoder.ToBytes()
			assert.NilError(t, err)
			assert.Equal(t, tt.expected, hex.EncodeToString(result))
		})
	}
}
