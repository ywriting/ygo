package ygo

type ClientID uint64;

type ID struct {
	Client ClientID
	Clock uint32
}