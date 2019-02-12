package fastdiv

import "math/bits"

// Int64 calculates division by using a pre-computed inverse.
type Int64 struct {
	absd   uint64
	hi, lo uint64
	neg    bool
}

// NewInt64 initializes a new pre-computed inverse.
func NewInt64(d int64) Int64 {
	var neg bool
	if d < 0 {
		neg = true
		d = -d
	}
	absd := uint64(d)

	hi, r := ^uint64(0)/absd, ^uint64(0)%absd
	lo, _ := bits.Div64(r, ^uint64(0), absd)

	var c uint64 = 1
	if absd&(absd-1) == 0 {
		c++
	}
	lo, c = bits.Add64(lo, c, 0)
	hi, _ = bits.Add64(hi, 0, c)
	return Int64{
		absd: absd,
		hi:   hi,
		lo:   lo,
		neg:  neg,
	}
}

// Div calculates n / d using the pre-computed inverse.
// Note, must have d != 1, -1, 0, or math.MinInt32
func (d Int64) Div(n int64) int64 {
	neg := d.neg
	if n < 0 {
		n = -n
		neg = !neg
	}

	divlo1, _ := bits.Mul64(d.lo, uint64(n))
	div, divlo2 := bits.Mul64(d.hi, uint64(n))
	var c uint64
	_, c = bits.Add64(divlo1, divlo2, 0)
	div, _ = bits.Add64(div, 0, c)

	if neg {
		return -int64(div)
	}
	return int64(div)
}

// Mod calculates n % d using the pre-computed inverse.
func (d Int64) Mod(n int64) int64 {
	var neg bool
	if n < 0 {
		n = -n
		neg = true
	}
	hi, lo := bits.Mul64(d.lo, uint64(n))
	hi += d.hi * uint64(n)
	modlo1, _ := bits.Mul64(lo, d.absd)
	mod, modlo2 := bits.Mul64(hi, d.absd)
	var c uint64
	_, c = bits.Add64(modlo1, modlo2, 0)
	mod, _ = bits.Add64(mod, 0, c)
	if neg {
		return -int64(mod)
	}
	return int64(mod)
}

// DivMod calculates n / d and n % d using the pre-computed inverse.
// Note, must have d != 1, -1, 0, or math.MinInt32
func (d Int64) DivMod(n int64) (q, r int64) {
	var neg bool
	if n < 0 {
		n = -n
		neg = true
	}

	divlo1, lo := bits.Mul64(d.lo, uint64(n))
	div, divlo2 := bits.Mul64(d.hi, uint64(n))

	hi, c := bits.Add64(divlo1, divlo2, 0)
	div, _ = bits.Add64(div, 0, c)
	q = int64(div)

	modlo1, _ := bits.Mul64(lo, d.absd)
	mod, modlo2 := bits.Mul64(hi, d.absd)

	_, c = bits.Add64(modlo1, modlo2, 0)
	mod, _ = bits.Add64(mod, 0, c)
	r = int64(mod)

	if neg {
		q = -q
		r = -r
	}
	if d.neg {
		q = -q
	}

	return q, r
}
