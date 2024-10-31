package ygo

type ClientID uint64

type ID struct {
	Client ClientID
	Clock  uint32
}

type Block interface {
	LastId() ID
	AsItem() (*Item, error)
	IsDeleted() bool
	Id() ID
	Len() uint32
	SameType(other *Block) bool
	IsGc() bool
	IsItem() bool
	Contains(id ID) bool
}

type Item struct {
	ID          ID
	Len         uint32
	Left        *Block
	Right       *Block
	Origin      *ID
	RightOrigin *ID
	Content     ItemContent
	Parent      TypePtr
	ParentSub   *string
	Moved       *Block
	Info        ItemFlags
}

func NewItem(id ID, left *Block, origin *ID, right *Block, rightOrigin *ID,
	parent TypePtr, parentSub *string,
	content ItemContent) Item {
	panic("not implemented")
}

type GC struct {
}

type BlockRange struct {
	ID  ID
	Len uint32
}

type ItemContent interface{}

type AnyContent struct{}

type BinaryContent struct{}

type DeletedContent struct{}

type DocContent struct{}

type JsonContent struct{}

type EmbedContent struct{}

type FormatContent struct{}

type StringContent struct{}

type TypeContent struct{}

type MoveContent struct{}

const (
	ITEM_FLAG_MARKED    uint8 = 0b0000_1000
	ITEM_FLAG_DELETED   uint8 = 0b0000_0100
	ITEM_FLAG_COUNTABLE uint8 = 0b0000_0010
	ITEM_FLAG_KEEP      uint8 = 0b0000_0001
)

type ItemFlags struct {
	flags uint8
}

func NewItemFlags(source uint8) ItemFlags {
	return ItemFlags{flags: source}
}

func (i *ItemFlags) Into() uint8 {
	return i.flags
}

func (i *ItemFlags) Set(value uint8) {
	i.flags |= value
}

func (i *ItemFlags) Clear(value uint8) {
	i.flags &= ^value
}

func (i *ItemFlags) Check(value uint8) bool {
	return i.flags&value == value
}

func (i *ItemFlags) IsKeep() bool {
	return i.Check(ITEM_FLAG_KEEP)
}

func (i *ItemFlags) IsCountable() bool {
	return i.Check(ITEM_FLAG_COUNTABLE)
}

func (i *ItemFlags) IsDeleted() bool {
	return i.Check(ITEM_FLAG_DELETED)
}

func (i *ItemFlags) IsMarked() bool {
	return i.Check(ITEM_FLAG_MARKED)
}

func (i *ItemFlags) SetCountable() {
	i.Set(ITEM_FLAG_COUNTABLE)
}

func (i *ItemFlags) ClearCountable() {
	i.Clear(ITEM_FLAG_COUNTABLE)
}

func (i *ItemFlags) SetDeleted() {
	i.Set(ITEM_FLAG_DELETED)
}
