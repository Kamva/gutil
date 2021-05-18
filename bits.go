package gutil

type Bits int32

func (b Bits) Has(f Bits) bool {
	return b&f != 0
}

func (b Bits) Set(f Bits) Bits {
	return b | f
}

func (b Bits) Toggle(f Bits) Bits {
	return b ^ f
}

func (b Bits) Clear(f Bits) Bits {
	return b &^ f
}

func (b Bits) HasInt32(f int32) bool {
	return b.Has(Bits(f))
}

func (b Bits) SetInt32(f int32) Bits {
	return b.Set(Bits(f))
}

func (b Bits) ToggleInt32(f int32) Bits {
	return b.Toggle(Bits(f))
}

func (b Bits) ClearInt32(f int32) Bits {
	return b.Clear(Bits(f))
}
