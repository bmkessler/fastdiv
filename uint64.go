package fastdiv

import "math/bits"

// Uint64 calculates division by using a pre-computed inverse.
type Uint64 struct {
	d      uint64
	hi, lo uint64
}

// NewUint64 initializes a new pre-computed inverse.
func NewUint64(d uint64) Uint64 {
	hi, r := ^uint64(0)/d, ^uint64(0)%d
	lo, _ := bits.Div64(r, ^uint64(0), d)
	var c uint64
	lo, c = bits.Add64(lo, 1, 0)
	hi, _ = bits.Add64(hi, 0, c)
	return Uint64{
		d:  d,
		hi: hi,
		lo: lo,
	}
}

// Div calculates n / d using the pre-computed inverse.
func (d Uint64) Div(n uint64) uint64 {
	divlo1, _ := bits.Mul64(d.lo, n)
	div, divlo2 := bits.Mul64(d.hi, n)
	var c uint64
	_, c = bits.Add64(divlo1, divlo2, 0)
	div, _ = bits.Add64(div, 0, c)
	return div
}

// Mod calculates n % d using the pre-computed inverse.
func (d Uint64) Mod(n uint64) uint64 {
	hi, lo := bits.Mul64(d.lo, n)
	hi += d.hi * n
	modlo1, _ := bits.Mul64(lo, d.d)
	mod, modlo2 := bits.Mul64(hi, d.d)
	var c uint64
	_, c = bits.Add64(modlo1, modlo2, 0)
	mod, _ = bits.Add64(mod, 0, c)
	return mod
}

// DivMod calculates n / d and n % d using the pre-computed inverse.
// Note must have d > 1.
func (d Uint64) DivMod(n uint64) (q, r uint64) {
	divlo1, lo := bits.Mul64(d.lo, n)
	div, divlo2 := bits.Mul64(d.hi, n)

	hi, c := bits.Add64(divlo1, divlo2, 0)
	div, _ = bits.Add64(div, 0, c)
	q = uint64(div)

	modlo1, _ := bits.Mul64(lo, d.d)
	mod, modlo2 := bits.Mul64(hi, d.d)

	_, c = bits.Add64(modlo1, modlo2, 0)
	mod, _ = bits.Add64(mod, 0, c)
	r = uint64(mod)

	return q, r
}

// Divisible determines whether n is exactly divisible by d using the pre-computed inverse.
func (d Uint64) Divisible(n uint64) bool {
	var hicheck, locheck, b uint64
	locheck, b = bits.Sub64(d.lo, 1, b)
	hicheck, _ = bits.Sub64(d.hi, 0, b)
	hi, lo := bits.Mul64(d.lo, n)
	hi += d.hi * n
	return (hi < hicheck) || ((hi == hicheck) && (lo <= locheck))
}
