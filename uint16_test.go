package fastdiv

import (
	"math"
	"testing"
	"testing/quick"
)

var (
	sinkUint16 uint16 = math.MaxUint16
	varUint16  uint16 = 14 //12345
)

const constUint16 uint16 = 12345

func TestUint16Div(t *testing.T) {
	checkUint16Div := func(x, y uint16) bool {
		if y == 0 || y == 1 {
			return true
		}
		d := NewUint16(y)
		return (x / y) == d.Div(x)
	}

	if err := quick.Check(checkUint16Div, nil); err != nil {
		t.Error(err)
	}
}

func TestUint16Mod(t *testing.T) {
	checkUint16Mod := func(x, y uint16) bool {
		if y == 0 {
			return true
		}
		d := NewUint16(y)
		return (x % y) == d.Mod(x)
	}

	if err := quick.Check(checkUint16Mod, nil); err != nil {
		t.Error(err)
	}
}

func TestUint16DivMod(t *testing.T) {
	checkUint16ModDiv := func(x, y uint16) bool {
		if y == 0 {
			return true
		}
		d := NewUint16(y)
		q, r := d.DivMod(x)
		return (x/y) == q && (x%y) == r
	}

	if err := quick.Check(checkUint16ModDiv, nil); err != nil {
		t.Error(err)
	}
}

func TestUint16Divisible(t *testing.T) {
	checkUint16Divisible := func(x, y uint16) bool {
		if x < y {
			x, y = y, x
		}
		if y == 0 {
			return true
		}
		d := NewUint16(y)
		if ((x % y) == 0) != d.Divisible(x) {
			return false
		}
		x, y = x&math.MaxUint8, y&math.MaxUint8
		if x == 0 || y == 0 {
			return true
		}
		d = NewUint16(y)
		return d.Divisible(x * y)
	}

	if err := quick.Check(checkUint16Divisible, nil); err != nil {
		t.Error(err)
	}
}

func TestUint16DivSpeed(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping speed test in short mode")
	}
	const samples = 100
	results := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < samples; j++ {
				sinkUint16 = sinkUint16 / varUint16
			}
		}
	})
	baseline := float64(results.NsPerOp()) / float64(samples)
	t.Logf("baseline: %2.2f ns/op", baseline)

	benchmarkUint16Div := func(b *testing.B, n int) {
		for i := 0; i < b.N; i++ {
			d := NewUint16(varUint16)
			for j := 0; j < n; j++ {
				sinkUint16 = d.Div(sinkUint16)
			}
		}
	}

	sizes := []int{1, 2, 3, 5, 8, 10}
	for _, s := range sizes {
		results = testing.Benchmark(func(b *testing.B) {
			benchmarkUint16Div(b, s)
		})
		nsPerOp := float64(results.NsPerOp()) / float64(s)
		t.Logf(" %2d divs: %2.2f ns/op, faster: %v", s, nsPerOp, nsPerOp < baseline)
	}
}

func TestUint16ModSpeed(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping speed test in short mode")
	}
	const samples = 100
	results := testing.Benchmark(func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < samples; j++ {
				sinkUint16 = sinkUint16 % varUint16
			}
		}
	})
	baseline := float64(results.NsPerOp()) / float64(samples)
	t.Logf("baseline: %2.2f ns/op", baseline)

	benchmarkUint16Mod := func(b *testing.B, n int) {
		for i := 0; i < b.N; i++ {
			d := NewUint16(varUint16)
			for j := 0; j < n; j++ {
				sinkUint16 = d.Mod(sinkUint16)
			}
		}
	}

	sizes := []int{1, 2, 3, 5, 8, 10}
	for _, s := range sizes {
		results = testing.Benchmark(func(b *testing.B) {
			benchmarkUint16Mod(b, s)
		})
		nsPerOp := float64(results.NsPerOp()) / float64(s)
		t.Logf(" %2d mods: %2.2f ns/op, faster: %v", s, nsPerOp, nsPerOp < baseline)
	}
}

func BenchmarkUint16DivVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkUint16 = sinkUint16 / varUint16
	}
}

func BenchmarkUint16DivConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkUint16 = sinkUint16 / constUint16
	}
}

func BenchmarkUint16Div(b *testing.B) {
	d := NewUint16(varUint16)
	for i := 0; i < b.N; i++ {
		sinkUint16 = d.Div(sinkUint16)
	}
}

func BenchmarkUint16ModVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkUint16 = sinkUint16 % varUint16
	}
}

func BenchmarkUint16ModConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sinkUint16 = sinkUint16 % constUint16
	}
}

func BenchmarkUint16Mod(b *testing.B) {
	d := NewUint16(varUint16)
	for i := 0; i < b.N; i++ {
		sinkUint16 = d.Mod(sinkUint16)
	}
}

func BenchmarkUint16DivisibleVar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if (uint16(i) % varUint16) == 0 {
			sinkUint16 = 2.0
		} else {
			sinkUint16 = 1.0
		}
	}
}

func BenchmarkUint16DivisibleConst(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if (uint16(i) % constUint16) == 0 {
			sinkUint16 = 2.0
		} else {
			sinkUint16 = 1.0
		}
	}
}

func BenchmarkUint16Divisible(b *testing.B) {
	d := NewUint16(varUint16)
	for i := 0; i < b.N; i++ {
		if d.Divisible(uint16(i)) {
			sinkUint16 = 2.0
		} else {
			sinkUint16 = 1.0
		}
	}
}
