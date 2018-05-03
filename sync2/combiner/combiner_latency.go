package combiner

import (
	"fmt"
	"sync"
	"testing"

	"github.com/loov/hrtime"
)

func RunLatency(t *testing.T, procs int, create func(Batcher) Combiner) {
	const N = 10000
	const K = 10
	var workSize = []int{0, 100}
	for _, startWork := range workSize {
		for _, includeWork := range workSize {
			for _, finishWork := range workSize {
				name := fmt.Sprintf("s%vi%vr%v", startWork, includeWork, finishWork)
				t.Run(name, func(t *testing.T) {
					hrs := make([]*hrtime.BenchmarkTSC, procs)
					for i := range hrs {
						hrs[i] = hrtime.NewBenchmarkTSC(N)
					}

					var wg sync.WaitGroup
					wg.Add(procs)

					var work Worker
					work.WorkStart = startWork
					work.WorkInclude = includeWork
					work.WorkFinish = finishWork

					combiner := create(&work)
					defer StartClose(combiner)()

					for i := 0; i < procs; i++ {
						go func(b *hrtime.BenchmarkTSC) {
							defer wg.Done()
							v := int64(0)
							for b.Next() {
								for k := 0; k < K; k++ {
									combiner.Do(v)
									v++
								}
							}
						}(hrs[i])
					}

					// TODO: save benchmark results to file
				})
			}
		}
	}
}
