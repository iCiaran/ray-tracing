package maths

type TextureSolid struct {
	value *Colour
}

func NewTextureSolid(r, g, b float64) *TextureSolid {
	return &TextureSolid{NewVec3(r, g, b)}
}

func (ts *TextureSolid) Value(u, v float64, p *Point3) *Colour {
	return ts.value
}
