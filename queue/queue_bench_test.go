package queue

import (
	"strconv"
	"sync"
	"testing"
)

const BenchSize = 8192

//var BenchWork = []int{0, 100}
var BenchWork = []int{0}

func localwork(amount int) {
	foo := 1
	for i := 0; i < amount; i++ {
		foo *= 2
		foo /= 2
	}
}

func bench(b *testing.B, ctor Ctor) {
	benchCommon(b, ctor)

	// blocking interface tests
	b.Run("b", func(b *testing.B) {
		if ctor := ctor.SPSC(); ctor != nil {
			b.Run("SPSC", func(b *testing.B) { benchSPSC(b, ctor) })
		}
		if ctor := ctor.SPMC(); ctor != nil {
			b.Run("SPMC", func(b *testing.B) { benchSPMC(b, ctor) })
		}
		if ctor := ctor.MPSC(); ctor != nil {
			b.Run("MPSC", func(b *testing.B) { benchMPSC(b, ctor) })
		}
		if ctor := ctor.MPMC(); ctor != nil {
			b.Run("MPMC", func(b *testing.B) { benchMPMC(b, ctor) })
		}
	})

	// non-blocking interface tests
	b.Run("n", func(b *testing.B) {
		if ctor := ctor.NonblockingSPSC(); ctor != nil {
			b.Run("SPSC", func(b *testing.B) { benchNonblockingSPSC(b, ctor) })
		}
		// if ctor := ctor.NonblockingSPMC(); ctor != nil {
		// 	b.Run("SPMC", func(b *testing.B) { benchNonblockingSPMC(b, ctor) })
		// }
		// if ctor := ctor.NonblockingMPSC(); ctor != nil {
		// 	b.Run("MPSC", func(b *testing.B) { benchNonblockingMPSC(b, ctor) })
		// }
		if ctor := ctor.NonblockingMPMC(); ctor != nil {
			b.Run("MPMC", func(b *testing.B) { benchNonblockingMPMC(b, ctor) })
		}
	})
}

func benchCommon(b *testing.B, ctor func(int) Queue) {
	b.Run("Create", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = ctor(TestSizes[i%len(TestSizes)])
		}
	})
}

func benchSPSC(b *testing.B, ctor func(int) SPSC) {
	b.Run("Single", func(b *testing.B) {
		q := ctor(BenchSize)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var v Value
			q.Send(v)
			q.Recv(&v)
		}
	})
	b.Run("Uncontended/x100", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			q := ctor(BenchSize)
			for pb.Next() {
				var v Value
				for i := 0; i < 100; i++ {
					q.Send(v)
					q.Recv(&v)
				}
			}
		})
	})

	for _, work := range BenchWork {
		suffix := ""
		if work > 0 {
			suffix = "/w" + strconv.Itoa(work)
		}
		b.Run("ProducerConsumer"+suffix, func(b *testing.B) {
			q := ctor(BenchSize)
			b.ResetTimer()
			var wg sync.WaitGroup
			wg.Add(2)
			go func() {
				for i := 0; i < b.N; i++ {
					var v Value
					q.Send(v)
					localwork(work)
				}
				wg.Done()
			}()
			go func() {
				for i := 0; i < b.N; i++ {
					var v Value
					q.Recv(&v)
					localwork(work)
				}
				wg.Done()
			}()
			wg.Wait()
		})
	}
}
func benchMPSC(b *testing.B, ctor func(int) MPSC) {
	for _, work := range BenchWork {
		suffix := ""
		if work > 0 {
			suffix = "/w" + strconv.Itoa(work)
		}
		b.Run("ProducerConsumer/x100"+suffix, func(b *testing.B) {
			q := ctor(BenchSize)
			b.ResetTimer()
			var wg sync.WaitGroup
			wg.Add(2)

			go func() {
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						for i := 0; i < 100; i++ {
							q.Send(0)
							localwork(work)
						}
					}
				})
				wg.Done()
			}()

			go func() {
				for i := 0; i < b.N; i++ {
					for i := 0; i < 100; i++ {
						var v Value
						q.Recv(&v)
						localwork(work)
					}
				}
				wg.Done()
			}()
			wg.Wait()
		})
	}
}
func benchSPMC(b *testing.B, ctor func(int) SPMC) {
	for _, work := range BenchWork {
		suffix := ""
		if work > 0 {
			suffix = "/w" + strconv.Itoa(work)
		}
		b.Run("ProducerConsumer/x100"+suffix, func(b *testing.B) {
			q := ctor(BenchSize)
			b.ResetTimer()
			var wg sync.WaitGroup
			wg.Add(2)

			go func() {
				for i := 0; i < b.N; i++ {
					for i := 0; i < 100; i++ {
						q.Send(0)
						localwork(work)
					}
				}
				wg.Done()
			}()

			go func() {
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						for i := 0; i < 100; i++ {
							var v Value
							q.Recv(&v)
						}
					}
				})
				wg.Done()
			}()
			wg.Wait()
		})
	}
}

func benchMPMC(b *testing.B, ctor func(int) MPMC) {
	b.Run("Contended/x100", func(b *testing.B) {
		q := ctor(BenchSize)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				var v Value
				for i := 0; i < 100; i++ {
					q.Send(v)
					q.Recv(&v)
				}
			}
		})
	})

	for _, work := range BenchWork {
		suffix := ""
		if work > 0 {
			suffix = "/w" + strconv.Itoa(work)
		}
		b.Run("ProducerConsumer/x100"+suffix, func(b *testing.B) {
			q := ctor(BenchSize)
			b.ResetTimer()
			var wg sync.WaitGroup
			wg.Add(2)

			go func() {
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						for i := 0; i < 100; i++ {
							q.Send(0)
							localwork(work)
						}
					}
				})
				wg.Done()
			}()

			go func() {
				b.RunParallel(func(pb *testing.PB) {
					for pb.Next() {
						for i := 0; i < 100; i++ {
							var v Value
							q.Recv(&v)
							localwork(work)
						}
					}
				})
				wg.Done()
			}()
			wg.Wait()
		})
	}
}
func benchNonblockingSPSC(b *testing.B, ctor func(int) NonblockingSPSC) {
	b.Run("Single", func(b *testing.B) {
		q := ctor(BenchSize)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			var v Value
			q.TrySend(v)
			q.TryRecv(&v)
		}
	})
	b.Run("Uncontended/x100", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			q := ctor(BenchSize)
			for pb.Next() {
				var v Value
				for i := 0; i < 100; i++ {
					q.TrySend(v)
					q.TryRecv(&v)
				}
			}
		})
	})
}

func benchNonblockingMPSC(b *testing.B, ctor func(int) NonblockingMPSC) { b.Skip("todo") }
func benchNonblockingSPMC(b *testing.B, ctor func(int) NonblockingSPMC) { b.Skip("todo") }

func benchNonblockingMPMC(b *testing.B, ctor func(int) NonblockingMPMC) {
	b.Run("Contended/x100", func(b *testing.B) {
		q := ctor(BenchSize)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				var v Value
				for i := 0; i < 100; i++ {
					q.TrySend(v)
					q.TryRecv(&v)
				}
			}
		})
	})
}
