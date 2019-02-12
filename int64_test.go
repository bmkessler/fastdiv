package fastdiv

import (
	"math"
	"testing"
	"testing/quick"
)

var (
	sinkInt64 int64 = math.MaxInt64
	varInt64  int64 = 123456789
)

const (
	constInt64 int64 = 123456789
)

func TestInt64Div(t *testing.T) {
	checkInt64Div := func(x, y int64) bool {
		if y == 0 || y == 1 || y == -1 || y == math.MinInt64 {
			return true
		}
		d := NewInt64(y)
		return (x / y) == d.Div(x)
	}

	if err := quick.Check(checkInt64Div, nil); err != nil {
		t.Error(err)
	}
}

func TestInt64Mod(t *testing.T) {
	checkInt64Mod := func(x, y int64) bool {
		if y == 0 || y == -1 || y == math.MinInt64 {
			return true
		}
		d := NewInt64(y)
		return (x % y) == d.Mod(x)
	}

	if err := quick.Check(checkInt64Mod, nil); err != nil {
		t.Error(err)
	}
}

func TestInt64DivMod(t *testing.T) {
	checkInt64DivMod := func(x, y int64) bool {
		if y == 0 || y == 1 || y == -1 || y == math.MinInt64 {
			return true
		}
		d := NewInt64(y)
		q, r := d.DivMod(x)
		return (x/y) == q && (x%y) == r
	}

	if err := quick.Check(checkInt64DivMod, nil); err != nil {
		t.Error(err)
	}
}

func TestInt64DivSpeed(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping speed test in short mode")
	}
	const samples = 100
	results := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < samples; j++ {
				sinkInt64 = sinkInt64 / varInt64
			}
		}
	})
	baseline := float64(results.NsPerOp()) / float64(samples)
	t.Logf("baseline: %2.2f ns/op", baseline)

	benchmarkInt64Div := func(b *testing.B, n int) {
		for i := 0; i < b.N; i++ {
			d := NewInt64(varInt64)
			for j := 0; j < n; j++ {
				sinkInt64 = d.Div(sinkInt64)
			}
		}
	}

	sizes := []int{1, 2, 3, 5, 8, 10}
	for _, s := range sizes {
		results = testing.Benchmark(func(b *testing.B) {
			benchmarkInt64Div(b, s)
		})
		nsPerOp := float64(results.NsPerOp()) / float64(s)
		t.Logf(" %2d divs: %2.2f ns/op, faster: %v", s, nsPerOp, nsPerOp < baseline)
	}
}

func TestInt64ModSpeed(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping speed test in short mode")
	}
	const samples = 100
	results := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < samples; j++ {
				sinkInt64 = sinkInt64 % varInt64
			}
		}
	})
	baseline := float64(results.NsPerOp()) / float64(samples)
	t.Logf("baseline: %2.2f ns/op", baseline)

	benchmarkInt64Mod := func(b *testing.B, n int) {
		for i := 0; i < b.N; i++ {
			d := NewInt64(varInt64)
			for j := 0; j < n; j++ {
				sinkInt64 = d.Mod(sinkInt64)
			}
		}
	}

	sizes := []int{1, 2, 3, 5, 8, 10}
	for _, s := range sizes {
		results = testing.Benchmark(func(b *testing.B) {
			benchmarkInt64Mod(b, s)
		})
		nsPerOp := float64(results.NsPerOp()) / float64(s)
		t.Logf(" %2d mods: %2.2f ns/op, faster: %v", s, nsPerOp, nsPerOp < baseline)
	}
}

func BenchmarkInt64DivVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkInt64 = sinkInt64 / varInt64
	}
}

func BenchmarkInt64DivConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkInt64 = sinkInt64 / constInt64
	}
}

func BenchmarkInt64Div(b *testing.B) {
	d := NewInt64(varInt64)
	for i := 0; i < b.N; i++ {
		sinkInt64 = d.Div(sinkInt64)
	}
}

func BenchmarkInt64ModVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkInt64 = sinkInt64 % varInt64
	}
}

func BenchmarkInt64ModConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkInt64 = sinkInt64 % constInt64
	}
}

func BenchmarkInt64Mod(b *testing.B) {
	d := NewInt64(varInt64)
	for i := 0; i < b.N; i++ {
		sinkInt64 = d.Mod(sinkInt64)
	}
}
