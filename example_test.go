package fastdiv_test

import (
	"fmt"

	"github.com/bmkessler/fastdiv"
)

func Example() {
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
}
