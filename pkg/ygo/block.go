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
	info := NewItemFlags(0)
	if content.IsCountable() {
		info.SetCountable()
	}
	len := content.Len(Utf16)
	item := Item{
		ID:          id,
		Len:         len,
		Left:        left,
		Origin:      origin,
		Right:       right,
		RightOrigin: rightOrigin,
		Content:     content,
		Parent:      parent,
		ParentSub:   parentSub,
		Info:        info,
		Moved:       nil,
	}
	// todo:
	return item
}

func (i *Item) Contains(id ID) bool {
	return i.ID.Client == id.Client &&
		id.Clock >= i.ID.Clock &&
		id.Clock < i.ID.Clock+i.Len
}

func (i *Item) IsDeleted() bool {
	return i.Info.IsDeleted()
}

func (i *Item) IsCountable() bool {
	return i.Info.IsCountable()
}

func (i *Item) MarkAsDeleted() {
	i.Info.SetDeleted()
}

func (i *Item) Repair(store *Store) {
	panic("todo")
}

func (i *Item) ContentLen(kind OffsetKind) uint32 {
	return i.Content.Len(kind)
}

func (i *Item) LastId() ID {
	return ID{
		Client: i.ID.Client,
		Clock:  i.ID.Clock + i.Len - 1,
	}
}

const (
	HAS_ORIGIN       uint8 = 0b1000_0000
	HAS_RIGHT_ORIGIN uint8 = 0b0100_0000
	HAS_PARENT_SUB   uint8 = 0b0010_0000
)

func (i *Item) ItemInfo() uint8 {
	var info uint8 = 0
	if i.Origin != nil {
		info |= HAS_ORIGIN
	}

	if i.RightOrigin != nil {
		info |= HAS_RIGHT_ORIGIN
	}

	if i.ParentSub != nil {
		info |= HAS_PARENT_SUB
	}
	info |= i.Content.GetRefNumber() & 0b1111
	return info
}

type GC struct {
}

type BlockRange struct {
	ID  ID
	Len uint32
}

type ItemContent interface {
	GetRefNumber() uint8
	IsCountable() bool
	Len(kind OffsetKind) uint32
}

const (
	BLOCK_GC_REF_NUMBER           uint8 = 0
	BLOCK_ITEM_DELETED_REF_NUMBER uint8 = 1
	BLOCK_ITEM_JSON_REF_NUMBER    uint8 = 2
	BLOCK_ITEM_BINARY_REF_NUMBER  uint8 = 3
	BLOCK_ITEM_STRING_REF_NUMBER  uint8 = 4
	BLOCK_ITEM_EMBED_REF_NUMBER   uint8 = 5
	BLOCK_ITEM_FORMAT_REF_NUMBER  uint8 = 6
	BLOCK_ITEM_TYPE_REF_NUMBER    uint8 = 7
	BLOCK_ITEM_ANY_REF_NUMBER     uint8 = 8
	BLOCK_ITEM_DOC_REF_NUMBER     uint8 = 9
	BLOCK_SKIP_REF_NUMBER         uint8 = 10
	BLOCK_ITEM_MOVE_REF_NUMBER    uint8 = 11
)

type AnyContent struct {
}

func (c *AnyContent) GetRefNumber() uint8 {
	return BLOCK_ITEM_ANY_REF_NUMBER
}

func (c *AnyContent) IsCountable() bool {
	return true
}

func (c *AnyContent) Len(kind OffsetKind) uint32 {
	return 0
}

type BinaryContent struct{}

func (c *BinaryContent) GetRefNumber() uint8 {
	return BLOCK_ITEM_BINARY_REF_NUMBER
}
func (c *BinaryContent) IsCountable() bool {
	return true
}

type DeletedContent struct{}

func (c *DeletedContent) GetRefNumber() uint8 {
	return BLOCK_ITEM_DELETED_REF_NUMBER
}

func (c *DeletedContent) IsCountable() bool {
	return false
}

type DocContent struct{}

func (c *DocContent) GetRefNumber() uint8 {
	return BLOCK_ITEM_DOC_REF_NUMBER
}

func (c *DocContent) IsCountable() bool {
	return true
}

type JsonContent struct{}

func (c *JsonContent) GetRefNumber() uint8 {
	return BLOCK_ITEM_DOC_REF_NUMBER
}

func (c *JsonContent) IsCountable() bool {
	return true
}

type EmbedContent struct{}

func (c *EmbedContent) GetRefNumber() uint8 {
	return BLOCK_ITEM_EMBED_REF_NUMBER
}

func (c *EmbedContent) IsCountable() bool {
	return true
}

type FormatContent struct{}

func (c *FormatContent) GetRefNumber() uint8 {
	return BLOCK_ITEM_FORMAT_REF_NUMBER
}

func (c *FormatContent) IsCountable() bool {
	return false
}

type StringContent struct{}

func (c *StringContent) GetRefNumber() uint8 {
	return BLOCK_ITEM_STRING_REF_NUMBER
}

func (c *StringContent) IsCountable() bool {
	return true
}

type TypeContent struct{}

func (c *TypeContent) GetRefNumber() uint8 {
	return BLOCK_ITEM_TYPE_REF_NUMBER
}

func (c *TypeContent) IsCountable() bool {
	return true
}

type MoveContent struct{}

func (c *MoveContent) GetRefNumber() uint8 {
	return BLOCK_ITEM_MOVE_REF_NUMBER
}

func (c *MoveContent) IsCountable() bool {
	return false
}

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
