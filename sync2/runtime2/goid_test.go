package runtime2

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/egonelbre/async"
)

func TestGOID(t *testing.T) {
	r := async.SpawnWithResult(1000, func(id int) error {
		start := GOID()
		for i := 0; i < 100; i++ {
			runtime.Gosched()
			id := GOID()
			idslow := goidslow()
			if start != id || start != idslow {
				return fmt.Errorf("unstable gid %v / %v / %v", start, id, idslow)
			}
		}
		return nil
	})

	if errs := r.Wait(); errs != nil {
		n := len(errs)
		if n > 10 {
			n = 10
		}
		for _, err := range errs[:n] {
			t.Error(err)
		}
	}
}
