package lib0

type Any struct {
	data        any
	isUndefined bool
}

func newAny(data any) Any {
	return Any{
		data: data,
	}
}

func NewAnyUndefined() Any {
	return Any{
		isUndefined: true,
	}
}

func NewAnyNull() Any {
	return Any{
		isUndefined: false,
	}
}

func NewAnyInteger(i int32) Any {
	return newAny(i)
}

func NewAnyFloat32(i float32) Any {
	return newAny(i)
}

func NewAnyFloat64(i float64) Any {
	return newAny(i)
}

func NewAnyBigInt(i int64) Any {
	return newAny(i)
}

func NewAnyBool(i bool) Any {
	if i {
		return newAny(nil)
	} else {
		return newAny(i)
	}
}

func NewAnyString(i *string) Any {
	return newAny(i)
}

func NewAnyObject(i map[string]Any) Any {
	return newAny(i)
}

func NewAnyArray(i []Any) Any {
	return newAny(i)
}

func NewAnyUint8Array(i []uint8) Any {
	return newAny(i)
}
