package ygo

type Store struct {
	Options   Options
	Types     map[string]Branch
	Blocks    BlockStore
	Pending   *PendingUpdate
	PendingDs *DeleteSet
	// todo: event handlers
}

func NewStore(options Options) Store {
	return Store{
		Options: options,
	}
}
