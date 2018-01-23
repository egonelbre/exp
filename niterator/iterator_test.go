package niterator

import (
	"testing"

	"github.com/egonelbre/exp/niterator/basic"
	"github.com/egonelbre/exp/niterator/instruct"
	"github.com/egonelbre/exp/niterator/onearr"
	"github.com/egonelbre/exp/niterator/onearrpremul"
	"github.com/egonelbre/exp/niterator/onearrrev"
	"github.com/egonelbre/exp/niterator/onearrrevspecialize"
	"github.com/egonelbre/exp/niterator/ordone"
	"github.com/egonelbre/exp/niterator/premul"
	"github.com/egonelbre/exp/niterator/shape"
	"github.com/egonelbre/exp/niterator/specialize"
	"github.com/egonelbre/exp/niterator/unroll"
	"github.com/egonelbre/exp/niterator/unrollinreverse"
	"github.com/egonelbre/exp/niterator/unrollinreversebool"
	"github.com/egonelbre/exp/niterator/unrollinreversehardcode"
	"github.com/egonelbre/exp/niterator/unrollinreverseswitch"
	"github.com/egonelbre/exp/niterator/unrollinverse"
	"github.com/egonelbre/exp/niterator/unrollreverse"
)

var (
	example = shape.New(24, 20, 40, 30)
	target  = make([]int, example.TotalSize())
)

func init() {
	it := basic.NewIterator(example)
	for index, err := it.Next(); err == nil; index, err = it.Next() {
		target[index] = index + 1
	}
	for _, v := range target {
		if v == 0 {
			panic("invalid iterator")
		}
	}
}

func testIterator(t *testing.T, it Iterator) {
	t.Helper()
	count := 0
	for index, err := it.Next(); err == nil; index, err = it.Next() {
		count++
		if target[index] != index+1 {
			t.Fatalf("invalid at %d", index)
		}
	}
	if count != example.TotalSize() {
		t.Fatalf("invalid count: got %d expected %d", count, example.TotalSize())
	}
}

func TestBasic(t *testing.T) {
	testIterator(t, basic.NewIterator(example))
}

func BenchmarkBasic(b *testing.B) {
	b.SetBytes(int64(example.TotalSize()))
	for i := 0; i < b.N; i++ {
		it := basic.NewIterator(example)
		total := 0
		for index, err := it.Next(); err == nil; index, err = it.Next() {
			total += index
		}
		_ = total
	}
}

func TestInstruct(t *testing.T) {
	testIterator(t, instruct.NewIterator(example))
}

func BenchmarkInstruct(b *testing.B) {
	b.SetBytes(int64(example.TotalSize()))
	for i := 0; i < b.N; i++ {
		it := instruct.NewIterator(example)
		total := 0
		for index, err := it.Next(); err == nil; index, err = it.Next() {
			total += index
		}
		_ = total
	}
}

func TestOrdone(t *testing.T) {
	testIterator(t, ordone.NewIterator(example))
}

func BenchmarkOrdone(b *testing.B) {
	b.SetBytes(int64(example.TotalSize()))
	for i := 0; i < b.N; i++ {
		it := ordone.NewIterator(example)
		total := 0
		for index, err := it.Next(); err == nil; index, err = it.Next() {
			total += index
		}
		_ = total
	}
}

func TestOnearr(t *testing.T) {
	testIterator(t, onearr.NewIterator(example))
}

func BenchmarkOnearr(b *testing.B) {
	b.SetBytes(int64(example.TotalSize()))
	for i := 0; i < b.N; i++ {
		it := onearr.NewIterator(example)
		total := 0
		for index, err := it.Next(); err == nil; index, err = it.Next() {
			total += index
		}
		_ = total
	}
}

func TestOnearrRev(t *testing.T) {
	testIterator(t, onearrrev.NewIterator(example))
}

func BenchmarkOnearrRev(b *testing.B) {
	b.SetBytes(int64(example.TotalSize()))
	for i := 0; i < b.N; i++ {
		it := onearrrev.NewIterator(example)
		total := 0
		for index, err := it.Next(); err == nil; index, err = it.Next() {
			total += index
		}
		_ = total
	}
}

func TestOnearrRevSpecialize(t *testing.T) {
	testIterator(t, onearrrevspecialize.NewIterator(example))
}

func BenchmarkOnearrRevSpecialize(b *testing.B) {
	b.SetBytes(int64(example.TotalSize()))
	for i := 0; i < b.N; i++ {
		it := onearrrevspecialize.NewIterator(example)
		total := 0
		for index, err := it.Next(); err == nil; index, err = it.Next() {
			total += index
		}
		_ = total
	}
}

func TestPremul(t *testing.T) {
	testIterator(t, premul.NewIterator(example))
}

func BenchmarkPremul(b *testing.B) {
	b.SetBytes(int64(example.TotalSize()))
	for i := 0; i < b.N; i++ {
		it := premul.NewIterator(example)
		total := 0
		for index, err := it.Next(); err == nil; index, err = it.Next() {
			total += index
		}
		_ = total
	}
}

func TestOnearrPremul(t *testing.T) {
	testIterator(t, onearrpremul.NewIterator(example))
}

func BenchmarkOnearrPremul(b *testing.B) {
	b.SetBytes(int64(example.TotalSize()))
	for i := 0; i < b.N; i++ {
		it := onearrpremul.NewIterator(example)
		total := 0
		for index, err := it.Next(); err == nil; index, err = it.Next() {
			total += index
		}
		_ = total
	}
}

func TestSpecialize(t *testing.T) {
	testIterator(t, specialize.NewIterator(example))
}

func BenchmarkSpecialize(b *testing.B) {
	b.SetBytes(int64(example.TotalSize()))
	for i := 0; i < b.N; i++ {
		it := specialize.NewIterator(example)
		total := 0
		for index, err := it.Next(); err == nil; index, err = it.Next() {
			total += index
		}
		_ = total
	}
}

func TestUnroll(t *testing.T) {
	testIterator(t, unroll.NewIterator(example))
}

func BenchmarkUnroll(b *testing.B) {
	b.SetBytes(int64(example.TotalSize()))
	for i := 0; i < b.N; i++ {
		it := unroll.NewIterator(example)
		total := 0
		for index, err := it.Next(); err == nil; index, err = it.Next() {
			total += index
		}
		_ = total
	}
}

func TestUnrollReverse(t *testing.T) {
	testIterator(t, unrollreverse.NewIterator(example))
}

func BenchmarkUnrollReverse(b *testing.B) {
	b.SetBytes(int64(example.TotalSize()))
	for i := 0; i < b.N; i++ {
		it := unrollreverse.NewIterator(example)
		total := 0
		for index, err := it.Next(); err == nil; index, err = it.Next() {
			total += index
		}
		_ = total
	}
}

func TestUnrollInverse(t *testing.T) {
	testIterator(t, unrollinverse.NewIterator(example))
}

func BenchmarkUnrollInverse(b *testing.B) {
	b.SetBytes(int64(example.TotalSize()))
	for i := 0; i < b.N; i++ {
		it := unrollinverse.NewIterator(example)
		total := 0
		for index, err := it.Next(); err == nil; index, err = it.Next() {
			total += index
		}
		_ = total
	}
}

func TestUnrollInReverse(t *testing.T) {
	testIterator(t, unrollinreverse.NewIterator(example))
}

func BenchmarkUnrollInReverse(b *testing.B) {
	b.SetBytes(int64(example.TotalSize()))
	for i := 0; i < b.N; i++ {
		it := unrollinreverse.NewIterator(example)
		total := 0
		for index, err := it.Next(); err == nil; index, err = it.Next() {
			total += index
		}
		_ = total
	}
}

func TestUnrollInReverseHardcode(t *testing.T) {
	testIterator(t, unrollinreversehardcode.NewIterator(example))
}

func BenchmarkUnrollInReverseHardcode(b *testing.B) {
	b.SetBytes(int64(example.TotalSize()))
	for i := 0; i < b.N; i++ {
		it := unrollinreversehardcode.NewIterator(example)
		total := 0
		for index, err := it.Next(); err == nil; index, err = it.Next() {
			total += index
		}
		_ = total
	}
}

func TestUnrollInReverseSwitch(t *testing.T) {
	testIterator(t, unrollinreverseswitch.NewIterator(example))
}

func BenchmarkUnrollInReverseSwitch(b *testing.B) {
	b.SetBytes(int64(example.TotalSize()))
	for i := 0; i < b.N; i++ {
		it := unrollinreverseswitch.NewIterator(example)
		total := 0
		for index, err := it.Next(); err == nil; index, err = it.Next() {
			total += index
		}
		_ = total
	}
}

func BenchmarkUnrollInReverseBool(b *testing.B) {
	b.SetBytes(int64(example.TotalSize()))
	for i := 0; i < b.N; i++ {
		it := unrollinreversebool.NewIterator(example)
		total := 0
		for index, done := it.Next(); !done; index, done = it.Next() {
			total += index
		}
		_ = total
	}
}
