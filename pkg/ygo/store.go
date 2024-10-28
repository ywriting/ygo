package ygo

import "riguz.com/ygo/pkg/ygo/types"

type Store struct {
	Options Options
	Types   map[string]types.Branch
}

func NewStore(options Options) Store {
	return Store{
		Options: options,
	}
}
