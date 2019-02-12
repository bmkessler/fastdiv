package fastdiv

import (
	"math"
	"testing"
	"testing/quick"
)

var (
	sinkInt16 int16 = math.MaxInt16
	varInt16  int16 = 12345
)

const constInt16 int16 = 12345

func TestInt16Div(t *testing.T) {
	checkInt16Div := func(x, y int16) bool {
		if y == 0 || y == 1 || y == -1 || y == math.MinInt16 {
			return true
		}
		d := NewInt16(y)
		return (x / y) == d.Div(x)
	}

	if err := quick.Check(checkInt16Div, nil); err != nil {
		t.Error(err)
	}
}

func TestInt16Mod(t *testing.T) {
	checkInt16Mod := func(x, y int16) bool {
		if y == 0 || y == -1 || y == math.MinInt16 {
			return true
		}
		d := NewInt16(y)
		return (x % y) == d.Mod(x)
	}

	if err := quick.Check(checkInt16Mod, nil); err != nil {
		t.Error(err)
	}
}

func TestInt16DivMod(t *testing.T) {
	checkInt16DivMod := func(x, y int16) bool {
		if y == 0 || y == 1 || y == -1 || y == math.MinInt16 {
			return true
		}
		d := NewInt16(y)
		q, r := d.DivMod(x)
		return (x/y) == q && (x%y) == r
	}

	if err := quick.Check(checkInt16DivMod, nil); err != nil {
		t.Error(err)
	}
}

func TestInt16DivSpeed(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping speed test in short mode")
	}
	const samples = 100
	results := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < samples; j++ {
				sinkInt16 = sinkInt16 / varInt16
			}
		}
	})
	baseline := float64(results.NsPerOp()) / float64(samples)
	t.Logf("baseline: %2.2f ns/op", baseline)

	benchmarkInt16Div := func(b *testing.B, n int) {
		for i := 0; i < b.N; i++ {
			d := NewInt16(varInt16)
			for j := 0; j < n; j++ {
				sinkInt16 = d.Div(sinkInt16)
			}
		}
	}

	sizes := []int{1, 2, 3, 5, 8, 10}
	for _, s := range sizes {
		results = testing.Benchmark(func(b *testing.B) {
			benchmarkInt16Div(b, s)
		})
		nsPerOp := float64(results.NsPerOp()) / float64(s)
		t.Logf(" %2d divs: %2.2f ns/op, faster: %v", s, nsPerOp, nsPerOp < baseline)
	}
}

func TestInt16ModSpeed(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping speed test in short mode")
	}
	const samples = 100
	results := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < samples; j++ {
				sinkInt16 = sinkInt16 % varInt16
			}
		}
	})
	baseline := float64(results.NsPerOp()) / float64(samples)
	t.Logf("baseline: %2.2f ns/op", baseline)

	benchmarkInt16Mod := func(b *testing.B, n int) {
		for i := 0; i < b.N; i++ {
			d := NewInt16(varInt16)
			for j := 0; j < n; j++ {
				sinkInt16 = d.Mod(sinkInt16)
			}
		}
	}

	sizes := []int{1, 2, 3, 5, 8, 10}
	for _, s := range sizes {
		results = testing.Benchmark(func(b *testing.B) {
			benchmarkInt16Mod(b, s)
		})
		nsPerOp := float64(results.NsPerOp()) / float64(s)
		t.Logf(" %2d mods: %2.2f ns/op, faster: %v", s, nsPerOp, nsPerOp < baseline)
	}
}

func BenchmarkInt16DivVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkInt16 = sinkInt16 / varInt16
	}
}

func BenchmarkInt16DivConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkInt16 = sinkInt16 / constInt16
	}
}

func BenchmarkInt16Div(b *testing.B) {
	d := NewInt16(varInt16)
	for i := 0; i < b.N; i++ {
		sinkInt16 = d.Div(sinkInt16)
	}
}

func BenchmarkInt16ModVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkInt16 = sinkInt16 % varInt16
	}
}

func BenchmarkInt16ModConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkInt16 = sinkInt16 % constInt16
	}
}

func BenchmarkInt16Mod(b *testing.B) {
	d := NewInt16(varInt16)
	for i := 0; i < b.N; i++ {
		sinkInt16 = d.Mod(sinkInt16)
	}
}
