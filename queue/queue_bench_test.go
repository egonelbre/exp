package queue

import (
	"sync"
	"testing"
)

const (
	BenchSize = 1024
	BenchWork = 100
)

func localwork() {
	foo := 1
	for i := 0; i < BenchWork; i++ {
		foo *= 2
		foo /= 2
	}
}

func bench(b *testing.B, ctor Ctor) {
	benchCommon(b, ctor)

	b.Run("b", func(b *testing.B) {
		if ctor := ctor.BlockingSPSC(); ctor != nil {
			b.Run("SPSC", func(b *testing.B) { benchBlockingSPSC(b, ctor) })
		}
		if ctor := ctor.BlockingSPMC(); ctor != nil {
			b.Run("SPMC", func(b *testing.B) { benchBlockingSPMC(b, ctor) })
		}
		if ctor := ctor.BlockingMPSC(); ctor != nil {
			b.Run("MPSC", func(b *testing.B) { benchBlockingMPSC(b, ctor) })
		}
		if ctor := ctor.BlockingMPMC(); ctor != nil {
			b.Run("MPMC", func(b *testing.B) { benchBlockingMPMC(b, ctor) })
		}
	})

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

func benchBlockingSPSC(b *testing.B, ctor func(int) BlockingSPSC) {
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

	for _, work := range []bool{false, true} {
		suffix := ""
		if work {
			suffix = "/work"
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
					if work {
						localwork()
					}
				}
				wg.Done()
			}()
			go func() {
				for i := 0; i < b.N; i++ {
					var v Value
					q.Recv(&v)
					if work {
						localwork()
					}
				}
				wg.Done()
			}()
			wg.Wait()
		})
	}
}
func benchBlockingMPSC(b *testing.B, ctor func(int) BlockingMPSC) {
	for _, work := range []bool{false, true} {
		suffix := ""
		if work {
			suffix = "/work"
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
							if work {
								localwork()
							}
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
						if work {
							localwork()
						}
					}
				}
				wg.Done()
			}()
			wg.Wait()
		})
	}
}
func benchBlockingSPMC(b *testing.B, ctor func(int) BlockingSPMC) {
	for _, work := range []bool{false, true} {
		suffix := ""
		if work {
			suffix = "/work"
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
						if work {
							localwork()
						}
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

func benchBlockingMPMC(b *testing.B, ctor func(int) BlockingMPMC) {
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

	for _, work := range []bool{false, true} {
		suffix := ""
		if work {
			suffix = "/work"
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
							if work {
								localwork()
							}
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
							if work {
								localwork()
							}
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

// func benchNonblockingMPSC(b *testing.B, ctor func(int) NonblockingMPSC) { b.Skip("todo") }
// func benchNonblockingSPMC(b *testing.B, ctor func(int) NonblockingSPMC) { b.Skip("todo") }

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
