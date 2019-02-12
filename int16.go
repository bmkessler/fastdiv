package fastdiv

import "math/bits"

// Int16 calculates division by using a pre-computed inverse.
type Int16 struct {
	absd uint32
	m    uint32
	neg  bool
}

// NewInt16 initializes a new pre-computed inverse for d != 0.
// If d == 0, a runtime divide-by-zero panic is raised.
func NewInt16(d int16) Int16 {
	var neg bool
	if d < 0 {
		neg = true
		d = -d
	}
	absd := uint32(d)
	m := ^uint32(0)/absd + 1
	if absd&(absd-1) == 0 {
		m++
	}
	return Int16{
		absd: absd,
		m:    m,
		neg:  neg,
	}
}

// Div calculates n / d using the pre-computed inverse.
// Note, must have d != 1, -1, 0, or math.MinInt16
func (d Int16) Div(n int16) int16 {
	neg := d.neg
	if n < 0 {
		n = -n
		neg = !neg
	}
	div, _ := bits.Mul32(d.m, uint32(n))
	if neg {
		return -int16(div)
	}
	return int16(div)
}

// Mod calculates n % d using the pre-computed inverse.
func (d Int16) Mod(n int16) int16 {
	fraction := d.m * uint32(n)
	mod, _ := bits.Mul32(fraction, d.absd)
	return int16(mod) - (int16(d.absd)-1)&(n>>31)
}

// DivMod calculates n / d and n % d using the pre-computed inverse.
// Note, must have d != 1, -1, 0, or math.MinInt16
func (d Int16) DivMod(n int16) (q, r int16) {
	var neg bool
	if n < 0 {
		n = -n
		neg = !neg
	}
	div, fraction := bits.Mul32(d.m, uint32(n))
	q = int16(div)
	mod, _ := bits.Mul32(fraction, d.absd)
	r = int16(mod)
	if neg {
		q = -q
		r = -r
	}
	if d.neg {
		q = -q
	}
	return q, r
}
