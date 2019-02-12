package fastdiv

import (
	"math"
	"testing"
	"testing/quick"
)

var (
	sinkInt32 int32 = math.MaxInt32
	varInt32  int32 = 123456789
)

const constInt32 int32 = 123456789

func TestInt32Div(t *testing.T) {
	checkInt32Div := func(x, y int32) bool {
		if y == 0 || y == 1 || y == -1 || y == math.MinInt32 {
			return true
		}
		d := NewInt32(y)
		return (x / y) == d.Div(x)
	}

	if err := quick.Check(checkInt32Div, nil); err != nil {
		t.Error(err)
	}
}

func TestInt32Mod(t *testing.T) {
	checkInt32Mod := func(x, y int32) bool {
		if y == 0 || y == -1 || y == math.MinInt32 {
			return true
		}
		d := NewInt32(y)
		return (x % y) == d.Mod(x)
	}

	if err := quick.Check(checkInt32Mod, nil); err != nil {
		t.Error(err)
	}
}

func TestInt32DivMod(t *testing.T) {
	checkInt32DivMod := func(x, y int32) bool {
		if y == 0 || y == 1 || y == -1 || y == math.MinInt32 {
			return true
		}
		d := NewInt32(y)
		q, r := d.DivMod(x)
		return (x/y) == q && (x%y) == r
	}

	if err := quick.Check(checkInt32DivMod, nil); err != nil {
		t.Error(err)
	}
}

func TestInt32DivSpeed(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping speed test in short mode")
	}
	const samples = 100
	results := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < samples; j++ {
				sinkInt32 = sinkInt32 / varInt32
			}
		}
	})
	baseline := float64(results.NsPerOp()) / float64(samples)
	t.Logf("baseline: %2.2f ns/op", baseline)

	benchmarkInt32Div := func(b *testing.B, n int) {
		for i := 0; i < b.N; i++ {
			d := NewInt32(varInt32)
			for j := 0; j < n; j++ {
				sinkInt32 = d.Div(sinkInt32)
			}
		}
	}

	sizes := []int{1, 2, 3, 5, 8, 10}
	for _, s := range sizes {
		results = testing.Benchmark(func(b *testing.B) {
			benchmarkInt32Div(b, s)
		})
		nsPerOp := float64(results.NsPerOp()) / float64(s)
		t.Logf(" %2d divs: %2.2f ns/op, faster: %v", s, nsPerOp, nsPerOp < baseline)
	}
}

func TestInt32ModSpeed(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping speed test in short mode")
	}
	const samples = 100
	results := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < samples; j++ {
				sinkInt32 = sinkInt32 % varInt32
			}
		}
	})
	baseline := float64(results.NsPerOp()) / float64(samples)
	t.Logf("baseline: %2.2f ns/op", baseline)

	benchmarkInt32Mod := func(b *testing.B, n int) {
		for i := 0; i < b.N; i++ {
			d := NewInt32(varInt32)
			for j := 0; j < n; j++ {
				sinkInt32 = d.Mod(sinkInt32)
			}
		}
	}

	sizes := []int{1, 2, 3, 5, 8, 10}
	for _, s := range sizes {
		results = testing.Benchmark(func(b *testing.B) {
			benchmarkInt32Mod(b, s)
		})
		nsPerOp := float64(results.NsPerOp()) / float64(s)
		t.Logf(" %2d mods: %2.2f ns/op, faster: %v", s, nsPerOp, nsPerOp < baseline)
	}
}

func BenchmarkInt32DivVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkInt32 = sinkInt32 / varInt32
	}
}

func BenchmarkInt32DivConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkInt32 = sinkInt32 / constInt32
	}
}

func BenchmarkInt32Div(b *testing.B) {
	d := NewInt32(varInt32)
	for i := 0; i < b.N; i++ {
		sinkInt32 = d.Div(sinkInt32)
	}
}

func BenchmarkInt32ModVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkInt32 = sinkInt32 % varInt32
	}
}

func BenchmarkInt32ModConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkInt32 = sinkInt32 % constInt32
	}
}

func BenchmarkInt32Mod(b *testing.B) {
	d := NewInt32(varInt32)
	for i := 0; i < b.N; i++ {
		sinkInt32 = d.Mod(sinkInt32)
	}
}
