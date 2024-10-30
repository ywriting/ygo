package ygo

type IdRange struct {
	Continuous *Continuous
	Fragmented *Fragmented
}

type Continuous struct {
}

type Fragmented struct {
}

type IdSet struct {
	Set map[ClientID]IdRange
}

type DeleteSet struct {
}
