package lib0

type Read interface {
	ReadExact(len uint) ([]uint8, error)
	ReadU8() (uint8, error)
	ReadBuf() ([]uint8, error)
	ReadU16() (uint16, error)
	ReadU32() (uint32, error)
	ReadU32Be(uint32, error)
	ReadVar()
}
