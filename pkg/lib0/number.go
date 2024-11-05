package lib0

import "math"

var (
	F64_MAX_SAFE_INTEGER float64 = float64(int64(math.Pow(2, 53)) - 1)
	F64_MIN_SAFE_INTEGER float64 = -F64_MAX_SAFE_INTEGER
)

type VarIntWrite interface {
	WriteVarUint(value uint) error
	WriteVarU8(value uint8) error
	WriteVarU16(num uint16) error
	WriteVarU32(num uint32) error
	WriteVarU64(num uint64) error
	WriteVarint(value uint) error
	WriteVarI8(value uint8) error
	WriteVarI16(num uint16) error
	WriteVarI32(num uint32) error
	WriteVarI64(num uint64) error
}
