package fastdiv

import (
	"math"
	"testing"
	"testing/quick"
)

var (
	sinkUint64 uint64 = math.MaxUint64
	varUint64  uint64 = 123456789
)

const constUint64 uint64 = 123456789

func TestUint64Div(t *testing.T) {
	checkUint64Div := func(x, y uint64) bool {
		if y == 0 || y == 1 {
			return true
		}
		d := NewUint64(y)
		return (x / y) == d.Div(x)
	}

	if err := quick.Check(checkUint64Div, nil); err != nil {
		t.Error(err)
	}
}

func TestUint64Mod(t *testing.T) {
	checkUint64Mod := func(x, y uint64) bool {
		if y == 0 {
			return true
		}
		d := NewUint64(y)
		return (x % y) == d.Mod(x)
	}

	if err := quick.Check(checkUint64Mod, nil); err != nil {
		t.Error(err)
	}
}

func TestUint64Divisible(t *testing.T) {
	checkUint64Divisible := func(x, y uint64) bool {
		if x < y {
			x, y = y, x
		}
		if y == 0 {
			return true
		}
		d := NewUint64(y)
		if ((x % y) == 0) != d.Divisible(x) {
			return false
		}
		x, y = x&math.MaxUint32, y&math.MaxUint32
		if x == 0 || y == 0 {
			return true
		}
		d = NewUint64(y)
		return d.Divisible(x * y)
	}

	if err := quick.Check(checkUint64Divisible, nil); err != nil {
		t.Error(err)
	}
}
func TestUint64DivSpeed(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping speed test in short mode")
	}
	const samples = 100
	results := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < samples; j++ {
				sinkUint64 = sinkUint64 / varUint64
			}
		}
	})
	baseline := float64(results.NsPerOp()) / float64(samples)
	t.Logf("baseline: %2.2f ns/op", baseline)

	benchmarkUint64Div := func(b *testing.B, n int) {
		for i := 0; i < b.N; i++ {
			d := NewUint64(varUint64)
			for j := 0; j < n; j++ {
				sinkUint64 = d.Div(sinkUint64)
			}
		}
	}

	sizes := []int{1, 2, 3, 5, 8, 10}
	for _, s := range sizes {
		results = testing.Benchmark(func(b *testing.B) {
			benchmarkUint64Div(b, s)
		})
		nsPerOp := float64(results.NsPerOp()) / float64(s)
		t.Logf(" %2d divs: %2.2f ns/op, faster: %v", s, nsPerOp, nsPerOp < baseline)
	}
}

func TestUint64ModSpeed(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping speed test in short mode")
	}
	const samples = 100
	results := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < samples; j++ {
				sinkUint64 = sinkUint64 % varUint64
			}
		}
	})
	baseline := float64(results.NsPerOp()) / float64(samples)
	t.Logf("baseline: %2.2f ns/op", baseline)

	benchmarkUint64Mod := func(b *testing.B, n int) {
		for i := 0; i < b.N; i++ {
			d := NewUint64(varUint64)
			for j := 0; j < n; j++ {
				sinkUint64 = d.Mod(sinkUint64)
			}
		}
	}

	sizes := []int{1, 2, 3, 5, 8, 10}
	for _, s := range sizes {
		results = testing.Benchmark(func(b *testing.B) {
			benchmarkUint64Mod(b, s)
		})
		nsPerOp := float64(results.NsPerOp()) / float64(s)
		t.Logf(" %2d mods: %2.2f ns/op, faster: %v", s, nsPerOp, nsPerOp < baseline)
	}
}

func BenchmarkUint64DivVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkUint64 = sinkUint64 / varUint64
	}
}

func BenchmarkUint64DivConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkUint64 = sinkUint64 / constUint64
	}
}

func BenchmarkUint64Div(b *testing.B) {
	d := NewUint64(varUint64)
	for i := 0; i < b.N; i++ {
		sinkUint64 = d.Div(sinkUint64)
	}
}

func BenchmarkUint64ModVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkUint64 = sinkUint64 % varUint64
	}
}

func BenchmarkUint64ModConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkUint64 = sinkUint64 % constUint64
	}
}

func BenchmarkUint64Mod(b *testing.B) {
	d := NewUint64(varUint64)
	for i := 0; i < b.N; i++ {
		sinkUint64 = d.Mod(sinkUint64)
	}
}

func BenchmarkUint64DivisibleVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if (uint64(i) % varUint64) == 0 {
			sinkUint64 = 2.0
		} else {
			sinkUint64 = 1.0
		}
	}
}

func BenchmarkUint64DivisibleConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if (uint64(i) % constUint64) == 0 {
			sinkUint64 = 2.0
		} else {
			sinkUint64 = 1.0
		}
	}
}

func BenchmarkUint64Divisible(b *testing.B) {
	d := NewUint64(varUint64)
	for i := 0; i < b.N; i++ {
		if d.Divisible(uint64(i)) {
			sinkUint64 = 2.0
		} else {
			sinkUint64 = 1.0
		}
	}
}
