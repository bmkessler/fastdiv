package fastdiv

import "math/bits"

// Uint16 calculates division by using a pre-computed inverse.
type Uint16 struct {
	d uint32
	m uint32
}

// NewUint16 initializes a new pre-computed inverse for d != 0.
// If d == 0, a runtime divide-by-zero panic is raised.
func NewUint16(d uint16) Uint16 {
	return Uint16{
		d: uint32(d),
		m: ^uint32(0)/uint32(d) + 1,
	}
}

// Div calculates n / d using the pre-computed inverse.
// Note must have d > 1.
func (d Uint16) Div(n uint16) uint16 {
	div, _ := bits.Mul32(d.m, uint32(n))
	return uint16(div)
}

// Mod calculates n % d using the pre-computed inverse.
func (d Uint16) Mod(n uint16) uint16 {
	fraction := d.m * uint32(n)
	mod, _ := bits.Mul32(fraction, d.d)
	return uint16(mod)
}

// DivMod calculates n / d and n % d using the pre-computed inverse.
// Note must have d > 1.
func (d Uint16) DivMod(n uint16) (uint16, uint16) {
	div, fraction := bits.Mul32(d.m, uint32(n))
	mod, _ := bits.Mul32(fraction, d.d)
	return uint16(div), uint16(mod)
}

// Divisible determines whether n is exactly divisible by d using the pre-computed inverse.
func (d Uint16) Divisible(n uint16) bool {
	return d.m*uint32(n) <= d.m-1
}
