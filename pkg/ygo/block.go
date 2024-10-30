package ygo

type ClientID uint64

type ID struct {
	Client ClientID
	Clock  uint32
}

type Block struct {
	Item *Item
	GC   *GC
}
type Item struct {
}

type GC struct {
}

type BlockRange struct {
	ID  ID
	Len uint32
}
