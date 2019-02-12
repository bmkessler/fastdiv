package fastdiv

import (
	"math"
	"testing"
	"testing/quick"
)

var (
	sinkUint32 uint32 = math.MaxUint32
	varUint32  uint32 = 123456789
)

const constUint32 uint32 = 123456789

func TestUint32Div(t *testing.T) {
	checkUint32Div := func(x, y uint32) bool {
		if y == 0 || y == 1 {
			return true
		}
		d := NewUint32(y)
		return (x / y) == d.Div(x)
	}

	if err := quick.Check(checkUint32Div, nil); err != nil {
		t.Error(err)
	}
}

func TestUint32Mod(t *testing.T) {
	checkUint32Mod := func(x, y uint32) bool {
		if y == 0 {
			return true
		}
		d := NewUint32(y)
		return (x % y) == d.Mod(x)
	}

	if err := quick.Check(checkUint32Mod, nil); err != nil {
		t.Error(err)
	}
}

func TestUint32DivMod(t *testing.T) {
	checkUint32ModDiv := func(x, y uint32) bool {
		if y == 0 {
			return true
		}
		d := NewUint32(y)
		q, r := d.DivMod(x)
		return (x/y) == q && (x%y) == r
	}

	if err := quick.Check(checkUint32ModDiv, nil); err != nil {
		t.Error(err)
	}
}

func TestUint32Divisible(t *testing.T) {
	checkUint32Divisible := func(x, y uint32) bool {
		if x < y {
			x, y = y, x
		}
		if y == 0 {
			return true
		}
		d := NewUint32(y)
		if ((x % y) == 0) != d.Divisible(x) {
			return false
		}
		x, y = x&math.MaxUint16, y&math.MaxUint16
		if x == 0 || y == 0 {
			return true
		}
		d = NewUint32(y)
		return d.Divisible(x * y)
	}

	if err := quick.Check(checkUint32Divisible, nil); err != nil {
		t.Error(err)
	}
}

func TestUint32DivSpeed(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping speed test in short mode")
	}
	const samples = 100
	results := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < samples; j++ {
				sinkUint32 = sinkUint32 / varUint32
			}
		}
	})
	baseline := float64(results.NsPerOp()) / float64(samples)
	t.Logf("baseline: %2.2f ns/op", baseline)

	benchmarkUint32Div := func(b *testing.B, n int) {
		for i := 0; i < b.N; i++ {
			d := NewUint32(varUint32)
			for j := 0; j < n; j++ {
				sinkUint32 = d.Div(sinkUint32)
			}
		}
	}

	sizes := []int{1, 2, 3, 5, 8, 10}
	for _, s := range sizes {
		results = testing.Benchmark(func(b *testing.B) {
			benchmarkUint32Div(b, s)
		})
		nsPerOp := float64(results.NsPerOp()) / float64(s)
		t.Logf(" %2d divs: %2.2f ns/op, faster: %v", s, nsPerOp, nsPerOp < baseline)
	}
}

func TestUint32ModSpeed(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping speed test in short mode")
	}
	const samples = 100
	results := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < samples; j++ {
				sinkUint32 = sinkUint32 % varUint32
			}
		}
	})
	baseline := float64(results.NsPerOp()) / float64(samples)
	t.Logf("baseline: %2.2f ns/op", baseline)

	benchmarkUint32Mod := func(b *testing.B, n int) {
		for i := 0; i < b.N; i++ {
			d := NewUint32(varUint32)
			for j := 0; j < n; j++ {
				sinkUint32 = d.Mod(sinkUint32)
			}
		}
	}

	sizes := []int{1, 2, 3, 5, 8, 10}
	for _, s := range sizes {
		results = testing.Benchmark(func(b *testing.B) {
			benchmarkUint32Mod(b, s)
		})
		nsPerOp := float64(results.NsPerOp()) / float64(s)
		t.Logf(" %2d mods: %2.2f ns/op, faster: %v", s, nsPerOp, nsPerOp < baseline)
	}
}

func BenchmarkUint32DivVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkUint32 = sinkUint32 / varUint32
	}
}

func BenchmarkUint32DivConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkUint32 = sinkUint32 / constUint32
	}
}

func BenchmarkUint32Div(b *testing.B) {
	d := NewUint32(varUint32)
	for i := 0; i < b.N; i++ {
		sinkUint32 = d.Div(sinkUint32)
	}
}

func BenchmarkUint32ModVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkUint32 = sinkUint32 % varUint32
	}
}

func BenchmarkUint32ModConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkUint32 = sinkUint32 % constUint32
	}
}

func BenchmarkUint32Mod(b *testing.B) {
	d := NewUint32(varUint32)
	for i := 0; i < b.N; i++ {
		sinkUint32 = d.Mod(sinkUint32)
	}
}

func BenchmarkUint32DivisibleVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if (uint32(i) % varUint32) == 0 {
			sinkUint32 = 2.0
		} else {
			sinkUint32 = 1.0
		}
	}
}

func BenchmarkUint32DivisibleConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if (uint32(i) % constUint32) == 0 {
			sinkUint32 = 2.0
		} else {
			sinkUint32 = 1.0
		}
	}
}

func BenchmarkUint32Divisible(b *testing.B) {
	d := NewUint32(varUint32)
	for i := 0; i < b.N; i++ {
		if d.Divisible(uint32(i)) {
			sinkUint32 = 2.0
		} else {
			sinkUint32 = 1.0
		}
	}
}
