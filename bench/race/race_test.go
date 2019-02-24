package race

import (
	"runtime"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func BenchmarkGenerateFromPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hash, _ := bcrypt.GenerateFromPassword([]byte{1, 2, 3, 4, 5}, bcrypt.DefaultCost)
		runtime.KeepAlive(hash)
	}
}

func BenchmarkCompareHashAndPassw(b *testing.B) {
	hash, _ := bcrypt.GenerateFromPassword([]byte{1, 2, 3, 4, 5}, bcrypt.DefaultCost)
	for i := 0; i < b.N; i++ {
		err := bcrypt.CompareHashAndPassword(hash, []byte{1, 2, 3, 4, 5})
		runtime.KeepAlive(err)
	}
}
