package permutation

import (
	"math/rand"
	"testing"
)

var examples = [][base]byte{
	[base]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
	[base]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 10},
	[base]byte{1, 0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
	[base]byte{0, 1, 9, 3, 5, 10, 6, 8, 7, 2, 11, 4},
}

const N = 1000

var randomized = [N][base]byte{}

func init() {
	for i := 0; i < N; i++ {
		var perm [base]byte
		for k, v := range rand.Perm(base) {
			perm[k] = byte(v)
		}
		randomized[i] = perm
	}
}

func TestCrossValidate(t *testing.T) {
	for _, example := range examples {
		check(t, example)
	}

}

func TestCrossValidateRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		var perm [base]byte
		for k, v := range rand.Perm(base) {
			perm[k] = byte(v)
		}
		check(t, perm)
	}
}

func check(t *testing.T, example [base]byte) {
	t.Helper()

	copying := CodeCopy(example)
	counting := CodeCount(example)
	table := CodeTable(example)
	table2 := CodeTable2(example)

	if copying != counting || copying != table || copying != table2 {
		t.Errorf("%v: copying %v, counting %v, table %v", example, copying, counting, table)
	}
}

func BenchmarkCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CodeCopy(randomized[i%N])
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CodeCount(randomized[i%N])
	}
}

func BenchmarkTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CodeTable(randomized[i%N])
	}
}

func BenchmarkTable2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = CodeTable2(randomized[i%N])
	}
}

func BenchmarkPackNybble(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PackNybble(randomized[i%N])
	}
}

func BenchmarkPackNybbleUnroll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PackNybbleUnroll(randomized[i%N])
	}
}

func BenchmarkPackNybbleUnroll2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PackNybbleUnroll2(randomized[i%N])
	}
}

func BenchmarkPackTight(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PackTight(randomized[i%N])
	}
}

func BenchmarkPackTightUnroll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = PackTightUnroll(randomized[i%N])
	}
}
