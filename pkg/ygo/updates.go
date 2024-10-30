package ygo

type UpdateBlocks struct {
	clients map[ClientID][]BlockCarrier
}

type BlockCarrier struct {
	Block *Block
	Skip  *BlockRange
}

type Update struct {
	Blocks    UpdateBlocks
	DeleteSet DeleteSet
}

type PendingUpdate struct {
	Update  Update
	Missing StateVector
}
