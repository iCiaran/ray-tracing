package maths

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const epsilon = 1e-6

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		a   *Vec3
		b   *Vec3
		res *Vec3
	}{
		{
			a:   NewVec3(1.0, 2.0, 3.0),
			b:   NewVec3(2.0, 3.0, 4.0),
			res: NewVec3(3.0, 5.0, 7.0),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := test.a.Add(test.b)
			assert.Equal(test.res, got)
		})
	}
}

func TestAddChain(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		a   *Vec3
		b   *Vec3
		c   *Vec3
		res *Vec3
	}{
		{
			a:   NewVec3(1.0, 2.0, 3.0),
			b:   NewVec3(2.0, 3.0, 4.0),
			c:   NewVec3(3.0, 4.0, 5.0),
			res: NewVec3(6.0, 9.0, 12.0),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := test.a.Add(test.b).Add(test.c)
			assert.Equal(test.res, got)
		})
	}
}

func TestMul(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		a   *Vec3
		b   float64
		res *Vec3
	}{
		{
			a:   NewVec3(1.0, 2.0, 3.0),
			b:   2.0,
			res: NewVec3(2.0, 4.0, 6.0),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := test.a.Mul(test.b)
			assert.Equal(test.res, got)
		})
	}
}

func TestMulChain(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		a   *Vec3
		b   float64
		c   float64
		res *Vec3
	}{
		{
			a:   NewVec3(1.0, 2.0, 3.0),
			b:   2.0,
			c:   3.0,
			res: NewVec3(6.0, 12.0, 18.0),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := test.a.Mul(test.b).Mul(test.c)
			assert.Equal(test.res, got)
		})
	}
}

func TestDiv(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		a   *Vec3
		b   float64
		res *Vec3
	}{
		{
			a:   NewVec3(1.0, 2.0, 3.0),
			b:   2.0,
			res: NewVec3(0.5, 1.0, 1.5),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := test.a.Div(test.b)
			assert.Equal(test.res, got)
		})
	}
}

func TestDivChain(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		a   *Vec3
		b   float64
		c   float64
		res *Vec3
	}{
		{
			a:   NewVec3(1.0, 2.0, 3.0),
			b:   2.0,
			c:   3.0,
			res: NewVec3(1.0/6.0, 2.0/6.0, 3.0/6.0),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := test.a.Div(test.b).Div(test.c)
			assert.Equal(test.res, got)
		})
	}
}

func TestLenSquared(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		a   *Vec3
		res float64
	}{
		{
			a:   NewVec3(1.0, 2.0, 3.0),
			res: 14.0,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := test.a.LenSquared()
			assert.Equal(test.res, got)
		})
	}
}

func TestLen(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		a   *Vec3
		res float64
	}{
		{
			a:   NewVec3(2.0, 10.0, 11.0),
			res: 15.0,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := test.a.Len()
			assert.Equal(test.res, got)
		})
	}
}

func TestAddTwo(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		a   *Vec3
		b   *Vec3
		res *Vec3
	}{
		{
			a:   NewVec3(1.0, 2.0, 3.0),
			b:   NewVec3(2.0, 3.0, 4.0),
			res: NewVec3(3.0, 5.0, 7.0),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := Add(test.a, test.b)
			assert.Equal(test.res, got)
		})
	}
}

func TestSubTwo(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		a   *Vec3
		b   *Vec3
		res *Vec3
	}{
		{
			a:   NewVec3(1.0, 2.0, 3.0),
			b:   NewVec3(2.0, 3.0, 4.0),
			res: NewVec3(-1.0, -1.0, -1.0),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := Sub(test.a, test.b)
			assert.Equal(test.res, got)
		})
	}
}

func TestMulTwo(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		a   *Vec3
		b   float64
		res *Vec3
	}{
		{
			a:   NewVec3(1.0, 2.0, 3.0),
			b:   3,
			res: NewVec3(3.0, 6.0, 9.0),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := Mul(test.a, test.b)
			assert.Equal(test.res, got)
		})
	}
}

func TestDivTwo(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		a   *Vec3
		b   float64
		res *Vec3
	}{
		{
			a:   NewVec3(1.0, 2.0, 3.0),
			b:   2,
			res: NewVec3(0.5, 1.0, 1.5),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := Div(test.a, test.b)
			assert.Equal(test.res, got)
		})
	}
}

func TestDot(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		a   *Vec3
		b   *Vec3
		res float64
	}{
		{
			a:   NewVec3(1.0, 2.0, 3.0),
			b:   NewVec3(2.0, 3.0, 4.0),
			res: 20.0,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := Dot(test.a, test.b)
			assert.Equal(test.res, got)
		})
	}
}

func TestCross(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		a   *Vec3
		b   *Vec3
		res *Vec3
	}{
		{
			a:   NewVec3(1.0, 2.0, 3.0),
			b:   NewVec3(2.0, 3.0, 4.0),
			res: NewVec3(-1.0, 2.0, -1.0),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := Cross(test.a, test.b)
			assert.Equal(test.res, got)
		})
	}
}

func TestNormalise(t *testing.T) {
	assert := assert.New(t)
	sqrt3 := math.Sqrt(3.0)
	tests := []struct {
		a   *Vec3
		b   *Vec3
		res *Vec3
	}{
		{
			a:   NewVec3(1.0, 1.0, 1.0),
			res: NewVec3(sqrt3/3, sqrt3/3, sqrt3/3),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprint("test_", i), func(t *testing.T) {
			got := Normalise(test.a)
			assert.InEpsilonSlice(test.res[:], got[:], epsilon)
		})
	}
}
