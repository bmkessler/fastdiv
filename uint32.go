package fastdiv

import "math/bits"

// Uint32 calculates division by using a pre-computed inverse.
type Uint32 struct {
	d uint64
	m uint64
}

// NewUint32 initializes a new pre-computed inverse for d != 0.
// If d == 0, a runtime divide-by-zero panic is raised.
func NewUint32(d uint32) Uint32 {
	return Uint32{
		d: uint64(d),
		m: ^uint64(0)/uint64(d) + 1,
	}
}

// Div calculates n / d using the pre-computed inverse.
// Note must have d > 1.
func (d Uint32) Div(n uint32) uint32 {
	div, _ := bits.Mul64(d.m, uint64(n))
	return uint32(div)
}

// Mod calculates n % d using the pre-computed inverse.
func (d Uint32) Mod(n uint32) uint32 {
	fraction := d.m * uint64(n)
	mod, _ := bits.Mul64(fraction, d.d)
	return uint32(mod)
}

// DivMod calculates n / d and n % d using the pre-computed inverse.
// Note must have d > 1.
func (d Uint32) DivMod(n uint32) (uint32, uint32) {
	div, fraction := bits.Mul64(d.m, uint64(n))
	mod, _ := bits.Mul64(fraction, d.d)
	return uint32(div), uint32(mod)
}

// Divisible determines whether n is exactly divisible by d using the pre-computed inverse.
func (d Uint32) Divisible(n uint32) bool {
	return d.m*uint64(n) <= d.m-1
}
