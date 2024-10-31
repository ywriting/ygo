package ygo

import (
	"math/rand"
	"time"
)

type Doc struct {
	ClientID ClientID
	store    Store
}

type Options struct {
	ClientID   ClientID
	OffsetKind OffsetKind
	SkipGc     bool
}

type OffsetKind int

const (
	Bytes OffsetKind = 0
	Utf16 OffsetKind = 1
	Utf32 OffsetKind = 2
)

func NewDefaultOptions() Options {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	return Options{
		ClientID:   ClientID(rng.Uint32()),
		OffsetKind: Bytes,
		SkipGc:     false,
	}
}
func NewOptionsWithClientID(ClientID ClientID) Options {
	return Options{
		ClientID:   ClientID,
		OffsetKind: Bytes,
		SkipGc:     false,
	}
}

func NewDoc() Doc {
	return NewDocWithOptions(NewDefaultOptions())
}

func NewDocWithClientID(ClientID ClientID) Doc {
	return NewDocWithOptions(NewOptionsWithClientID(ClientID))
}

func NewDocWithOptions(options Options) Doc {
	return Doc{
		ClientID: options.ClientID,
	}
}

func (d *Doc) Transact() Transaction {
	panic("not implemented")
}
