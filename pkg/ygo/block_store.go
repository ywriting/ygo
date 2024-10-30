package ygo

type BlockStore struct {
	clients map[ClientID]ClientBlockList
}

type ClientBlockList struct {
	list []*Block
}

type StateVector struct {
	Vector map[ClientID]uint32
}
