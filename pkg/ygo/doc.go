package ygo

type Doc struct {
	ClientID ClientID
}

type Options struct {
	ClientID ClientID
	OffsetKind OffsetKind
	SkipGC bool
}

type OffsetKind int
const (
	Bytes OffsetKind = 0
	Utf16 OffsetKind = 1
	Utf32 OffsetKind = 2
)