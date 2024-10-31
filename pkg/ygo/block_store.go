package ygo

type BlockStore struct {
	clients map[ClientID]ClientBlockList
}

func (b *BlockStore) GetStateVector() StateVector {
	panic("not implemented")
}

type ClientBlockList struct {
	list []*Block
}

func (c *ClientBlockList) GetState() uint32 {
	panic("not implemented")
	//item := c.Get(len(c.list) - 1)
	//return item.Id()
}

func (c *ClientBlockList) Get(index int) *Block {
	return c.list[index]
}

type StateVector struct {
	Vector map[ClientID]uint32
}

func NewStateVector() StateVector {
	return StateVector{
		Vector: make(map[ClientID]uint32),
	}
}
func NewStateVectorFrom(ss *BlockStore) StateVector {
	sv := NewStateVector()
	for clientId, clientStructList := range ss.clients {
		sv.Vector[clientId] = clientStructList.GetState()
	}
	return sv
}
