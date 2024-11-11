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
	vector map[ClientID]uint32
}

func NewStateVector() StateVector {
	return StateVector{
		vector: make(map[ClientID]uint32),
	}
}
func NewStateVectorFrom(ss *BlockStore) StateVector {
	sv := NewStateVector()
	for clientId, clientStructList := range ss.clients {
		sv.vector[clientId] = clientStructList.GetState()
	}
	return sv
}

func (s *StateVector) IsEmpty() bool {
	return len(s.vector) == 0
}

func (s *StateVector) Len() int {
	return len(s.vector)
}

func (s *StateVector) Get(clientId ClientID) uint32 {
	return s.vector[clientId]
}

func (s *StateVector) Contains(id ID) bool {
	return id.Clock <= s.Get(id.Client)
}

func (s *StateVector) IncreaseBy(clientId ClientID, delta uint32) {
	if delta > 0 {
		e := s.vector[clientId]
		e += delta
		s.vector[clientId] = e
	}
}

func (s *StateVector) SetMin(clientId ClientID, clock uint32) {
	c, exists := s.vector[clientId]
	if !exists || clock < c {
		s.vector[clientId] = clock
	}
}

func (s *StateVector) SetMax(clientId ClientID, clock uint32) {
	c, exists := s.vector[clientId]
	if !exists || clock > c {
		s.vector[clientId] = clock
	}
}

func (s *StateVector) Merge(other *StateVector) {
	for client, clock := range other.vector {
		s.SetMax(client, clock)
	}
}

func (s *StateVector) Encode(encoder Encoder) error {
	if err := encoder.WriteVarInt(len(s.vector)); err != nil {
		return err
	}
	for client, clock := range s.vector {
		if err := encoder.WriteVarUint64(uint64(client)); err != nil {
			return err
		}
		if err := encoder.WriteVarUint32(clock); err != nil {
			return err
		}
	}
	return nil
}
