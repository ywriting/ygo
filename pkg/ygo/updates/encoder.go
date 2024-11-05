package updates

import (
	"bufio"
	"bytes"
	"io"

	"riguz.com/ygo/pkg/ygo"
)

type Encode interface {
	Encode(encoder *Encoder)
	EncodeV1() ([]uint8, error)
	EncodeV2() ([]uint8, error)
}

type Encoder interface {
	ToBytes() []uint8
	ResetDsCurVal()
	WriteDsClock(clock uint32)
	WriteDsLen(len uint32)
	WriteLeftId(id ygo.ID)
	WriteRightId(id ygo.ID)
	WriteClient(client ygo.ClientID)
	WriteInfo(info uint8) error
	WriteParentInfo(isYKey bool) error
	WriteTypeRef(info uint8) error
	WriteLen(len uint32) error
	// WriteAny(any *Any)
	// WriteJson(any *Any)
	WriteKey(key *string) error
}

type EncoderV1 struct {
	buf io.Writer
}

func NewEncoderV1() EncoderV1 {
	var b bytes.Buffer

	return EncoderV1{
		buf: bufio.NewWriter(&b),
	}
}

func (e *EncoderV1) WriteId(id ygo.ID) {
	//bufio.
}
