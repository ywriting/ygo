package ygo

type Transaction struct {
	Store       *Store
	BeforeState StateVector
	AfterState  StateVector
	MergeBlocks []ID
	DeleteSet   DeleteSet
	PrevMoved   map[*Block]*Block
	changed     map[TypePtr][]string
	committed   bool
}

func NewTransaction(store *Store) Transaction {
	panic("not implemented")
}
