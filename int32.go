package fastdiv

import "math/bits"

// Int32 calculates division by using a pre-computed inverse.
type Int32 struct {
	absd uint64
	m    uint64
	neg  bool
}

// NewInt32 initializes a new pre-computed inverse for d != 0.
// If d == 0, a runtime divide-by-zero panic is raised.
func NewInt32(d int32) Int32 {
	var neg bool
	if d < 0 {
		neg = true
		d = -d
	}
	absd := uint64(d)
	m := ^uint64(0)/absd + 1
	if absd&(absd-1) == 0 {
		m++
	}
	return Int32{
		absd: absd,
		m:    m,
		neg:  neg,
	}
}

// Div calculates n / d using the pre-computed inverse.
// Note, must have d != 1, -1, 0, or math.MinInt32
func (d Int32) Div(n int32) int32 {
	neg := d.neg
	if n < 0 {
		n = -n
		neg = !neg
	}
	div, _ := bits.Mul64(d.m, uint64(n))
	if neg {
		return -int32(div)
	}
	return int32(div)
}

// Mod calculates n % d using the pre-computed inverse.
func (d Int32) Mod(n int32) int32 {
	fraction := d.m * uint64(n)
	mod, _ := bits.Mul64(fraction, d.absd)
	return int32(mod) - (int32(d.absd)-1)&(n>>31)
}

// DivMod calculates n / d and n % d using the pre-computed inverse.
// Note, must have d != 1, -1, 0, or math.MinInt32
func (d Int32) DivMod(n int32) (q, r int32) {
	var neg bool
	if n < 0 {
		n = -n
		neg = !neg
	}
	div, fraction := bits.Mul64(d.m, uint64(n))
	q = int32(div)
	mod, _ := bits.Mul64(fraction, d.absd)
	r = int32(mod)
	if neg {
		q = -q
		r = -r
	}
	if d.neg {
		q = -q
	}
	return q, r
}
