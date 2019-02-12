# fastdiv
[![GoDoc](https://godoc.org/github.com/bmkessler/fastdiv?status.svg)](https://godoc.org/github.com/bmkessler/fastdiv)

Fast division, modulus and divisibility checks for divisors known only at runtime via the method of:

"Faster Remainder by Direct Computation: Applications to Compilers and Software Libraries"
Daniel Lemire, Owen Kaser, Nathan Kurz
[arXiv:1902.01961 ](https://arxiv.org/abs/1902.01961)

Usage:
```
var divisor uint32 = 3

// intialize a divisor at runtime
d := fastdiv.NewUint32(divisor)

// use it repeatedly
var total uint32
for i := uint32(1); i < 10; i++ {
	total += d.Div(i)
	if d.Divisible(i) {
		fmt.Printf("%d is divisible by %d\n", i, divisor)
	}
}
fmt.Printf("Sum of quotients = %d", total)

// Output:
// 3 is divisible by 3
// 6 is divisible by 3
// 9 is divisible by 3
// Sum of quotients = 12
```

The method works by pre-computing an approximate inverse of the divisor such that the quotient is given by the high part of the multiplication and the remainder can be calculated by multiplying the fraction contained in the low part by the original divisor.
In general, the required accuracy for the approximate inverse is twice the width of the original divisor.
For divisors that are half the width of a register or less, this means that the quotient can be calculated with one high-multiplication (top word of a full-width multiplication), the remainder can be calculated with one low-multiplication followed by a high-multiplication and both can be calculated with one full-width multiplication and one high-multiplication.

On amd64 architecture for divisors that are 32-bits or less, this method can be faster than the traditional Granlund-Montgomery-Warren approach used to optimize constant divisions in the compiler.
The requirement that the approximate inverse be twice the divisor width means that extended arithmetic is required for 64-bit divisors.
The extended arithmetic makes this method is somewhat slower than the Granlund-Montgomery-Warren approach for these larger divisors, but still
faster than 64-bit division instructions.

The per operation speed up over a division instruction is ~2-3x and the overhead of pre-computing the inverse can be amortized after 1-6 repeated divisions with the same divisor.

| op  | size   | var     | const   | fastdiv | var / fastdiv | # to breakeven |
| --- | ------ | ------- | ------- | ------- | ------------- | -------------- |
| div | uint16 | 15.5 ns | 5.46 ns | 5.32 ns | 2.9x          | 1              |
| mod | uint16 | 16.2 ns | 7.68 ns | 7.68 ns | 2.1x          | 1              |
| div | int16  | 16.7 ns | 5.91 ns | 6.55 ns | 2.5x          | 1              |
| mod | int16  | 17.5 ns | 8.86 ns | 8.86 ns | 2.0x          | 1              |
| div | uint32 | 15.5 ns | 7.28 ns | 5.53 ns | 2.8x          | 2              |
| mod | uint32 | 15.7 ns | 9.63 ns | 7.09 ns | 2.2x          | 2              |
| div | int32  | 16.0 ns | 5.91 ns | 5.91 ns | 2.7x          | 2              |
| mod | int32  | 16.1 ns | 8.86 ns | 8.27 ns | 1.9x          | 2              |
| div | uint64 | 21.4 ns | 5.91 ns | 6.89 ns | 3.1x          | 5              |
| mod | uint64 | 20.3 ns | 8.30 ns | 8.87 ns | 2.3x          | 6              |
| div | int64  | 26.2 ns | 7.26 ns | 8.51 ns | 3.0x          | 5              |
| mod | int64  | 25.8 ns | 9.57 ns | 16.8 ns | 1.5x          | 6              |
