package codecancel

import (
	"context"
	"testing"
)

func BenchmarkSelect(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	for i := 0; i < b.N; i++ {
		select {
		case <-ctx.Done():
		default:
		}
	}
}

func BenchmarkErr(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	for i := 0; i < b.N; i++ {
		if ctx.Err() != nil {
			//
		}
	}
}

func BenchmarkChan(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	done := ctx.Done()

	for i := 0; i < b.N; i++ {
		select {
		case <-done:
		default:
		}
	}
}

func BenchmarkErrShort(b *testing.B) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	geterr := ctx.Err

	for i := 0; i < b.N; i++ {
		if geterr() != nil {
			//
		}
	}
}
