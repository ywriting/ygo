package ygo

import (
	"math/rand"
	"time"
)

type Doc struct {
	ClientId ClientId
	store    Store
}

type Options struct {
	ClientId   ClientId
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
		ClientId:   ClientId(rng.Uint32()),
		OffsetKind: Bytes,
		SkipGc:     false,
	}
}
func NewOptionsWithClientId(clientId ClientId) Options {
	return Options{
		ClientId:   clientId,
		OffsetKind: Bytes,
		SkipGc:     false,
	}
}

func NewDoc() Doc {
	return NewDocWithOptions(NewDefaultOptions())
}

func NewDocWithClientId(clientId ClientId) Doc {
	return NewDocWithOptions(NewOptionsWithClientId(clientId))
}

func NewDocWithOptions(options Options) Doc {
	return Doc{
		ClientId: options.ClientId,
	}
}
