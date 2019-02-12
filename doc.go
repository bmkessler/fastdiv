/*
Package fastdiv implements fast division, modulus and divisibility checks
for divisors known only at runtime via the method of:

"Faster Remainder by Direct Computation:
Applications to Compilers and Software Libraries"
Daniel Lemire, Owen Kaser, Nathan Kurz
https://arxiv.org/abs/1902.01961

The method works by pre-computing an approximate inverse of the divisor
such that the quotient is given by the high part of the multiplication
and the remainder can be calculated by multiplying the fraction contained
in the low part by the original divisor.  In general, the required accuracy
for the approximate inverse is twice the width of the original divisor.
For divisors that are half the width of a register or less, this means that
the quotient can be calculated with one high-multiplication (top word of a
full-width multiplication), the remainder can be calculated with one
low-multiplication followed by a high-multiplication and both can be
calculated with one full-width multiplication and one high-multiplication.

On amd64 architecture for divisors that are 32-bits or less, this method
can be faster than the traditional Granlund-Montgomery-Warren approach used
to optimize constant divisions in the compiler. The requirement that the
approximate inverse be twice the divisor width means that extended arithmetic
is required for 64-bit divisors.  The extended arithmetic makes this
method is somewhat slower than the Granlund-Montgomery-Warren approach for
these larger divisors, but still faster than 64-bit division instructions.

The per operation speed up over a division instruction is ~2-3x and the
overhead of pre-computing the inverse can be amortized after 1-6 repeated
divisions with the same divisor.
*/
package fastdiv
