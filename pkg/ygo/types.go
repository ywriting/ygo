package ygo

type Branch struct {
}

type TypePtr struct {
	Unknown *Unknown
	Branch  *Branch
	Named   *string
	ID      *ID
}

type Unknown struct{}
