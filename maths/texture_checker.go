package maths

import "math"

type TextureChecker struct {
	odd   Texture
	even  Texture
	width float64
}

func NewTextureChecker(t0, t1 Texture, w float64) *TextureChecker {
	return &TextureChecker{t0, t1, 1.0 / w}
}

func (tc *TextureChecker) Value(u, v float64, p *Vec3) *Colour {
	sines := math.Sin(tc.width*p.X()) * math.Sin(tc.width*p.Y()) * math.Sin(tc.width*p.Z())
	if sines < 0 {
		return tc.odd.Value(u, v, p)
	}
	return tc.even.Value(u, v, p)
}
