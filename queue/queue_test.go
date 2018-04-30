package queue

import (
	"flag"
	"fmt"
	"runtime"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"github.com/egonelbre/async"
	"github.com/egonelbre/exp/sync2/spin"
)

const TestProcs = 2

var shake = flag.Int("shake", 1, "run test multiple times in a row")

var BatchSizes = [...]int{1, 8, 32, 64}
var TestSizes = [...]int{1, 2, 3, 7, 8, 9, 15, 127, 128, 129}

func flushsend(q Queue) {
	if qf, ok := q.(Flusher); ok {
		qf.FlushSend()
	}
}

func flushrecv(q Queue) {
	if qf, ok := q.(Flusher); ok {
		qf.FlushRecv()
	}
}

func tryclose(q Queue) {
	if qf, ok := q.(Closer); ok {
		qf.Close()
	}
}

func mustSend(q NonblockingSPSC, v Value) bool {
	var s spin.Second
	for s.Spin() {
		if q.TrySend(v) {
			return true
		}
	}
	return false
}

func mustRecv(q NonblockingSPSC, v *Value) bool {
	var s spin.Second
	for s.Spin() {
		if q.TryRecv(v) {
			return true
		}
	}
	return false
}

func test(t *testing.T, ctor Ctor) {
	for k := 0; k < *shake; k++ {
		t.Run("b", func(t *testing.T) {
			if ctor := ctor.SPSC(); ctor != nil {
				t.Run("SPSC", func(t *testing.T) { testSPSC(t, ctor) })
			}
			if ctor := ctor.SPMC(); ctor != nil {
				t.Run("SPMC", func(t *testing.T) { testSPMC(t, ctor) })
			}
			if ctor := ctor.MPSC(); ctor != nil {
				t.Run("MPSC", func(t *testing.T) { testMPSC(t, ctor) })
			}
			if ctor := ctor.MPMC(); ctor != nil {
				t.Run("MPMC", func(t *testing.T) { testMPMC(t, ctor) })
			}
		})

		t.Run("n", func(t *testing.T) {
			if ctor := ctor.NonblockingSPSC(); ctor != nil {
				t.Run("SPSC", func(t *testing.T) { testNonblockingSPSC(t, ctor) })
			}
			if ctor := ctor.NonblockingSPMC(); ctor != nil {
				t.Run("SPMC", func(t *testing.T) { testNonblockingSPMC(t, ctor) })
			}
			if ctor := ctor.NonblockingMPSC(); ctor != nil {
				t.Run("MPSC", func(t *testing.T) { testNonblockingMPSC(t, ctor) })
			}
			if ctor := ctor.NonblockingMPMC(); ctor != nil {
				t.Run("MPMC", func(t *testing.T) { testNonblockingMPMC(t, ctor) })
			}
		})
	}
}

func testSPSC(t *testing.T, ctor func(int) SPSC) {
	t.Run("Basic", func(t *testing.T) {
		for _, size := range TestSizes {
			q := ctor(size)
			count := Value(size/2) + 1

			result := async.All(func() error {
				for i := Value(0); i < count; i++ {
					if !q.Send(i) {
						return fmt.Errorf("%v: failed to send %v", size, i)
					}
				}
				flushsend(q)
				return nil
			}, func() error {
				for exp := Value(0); exp < Value(exp); exp++ {
					var got Value
					if !q.Recv(&got) || got != exp {
						return fmt.Errorf("%v: invalid value got %v, expected %v", size, got, exp)
					}
				}
				return nil
			})

			if errs := result.Wait(); errs != nil {
				t.Fatal(errs)
			}
		}
	})

	if _, isBatchReceiver := ctor(1).(BatchReceiver); isBatchReceiver {
		t.Run("Batch", func(t *testing.T) {
			for _, size := range TestSizes {
				q := ctor(size)
				count := Value(size/2) + 1

				result := async.All(func() error {
					for i := Value(0); i < count; i++ {
						if !q.Send(i) {
							return fmt.Errorf("%v: failed to send %v", size, i)
						}
					}
					flushsend(q)
					return nil
				}, func() error {
					br := q.(BatchReceiver)
					var err error
					exp := Value(0)
					for exp < count {
						br.BatchRecv(func(got Value) {
							if got != exp {
								if err == nil {
									err = fmt.Errorf("%v: invalid value got %v, expected %v", size, got, exp)
								}
								tryclose(br)
								return
							}
							exp++
							if exp >= count {
								tryclose(br)
							}
						})
						if err != nil {
							return err
						}
					}
					return nil
				})
				if errs := result.Wait(); errs != nil {
					t.Fatal(errs)
				}
			}
		})
	}

	if _, isBounded := ctor(1).(Bounded); isBounded {
		t.Run("BlockOnFull", func(t *testing.T) {
			for _, size := range TestSizes {
				q := ctor(size)
				capacity := q.(Bounded).Cap()
				for i := 0; i < capacity; i++ {
					if !q.Send(0) {
						t.Fatal("failed to send")
					}
				}
				flushsend(q)
				sent := uint32(0)
				go func() {
					if !q.Send(0) {
						t.Fatal("failed to send")
					}
					flushsend(q)
					atomic.StoreUint32(&sent, 1)
				}()
				runtime.Gosched()
				time.Sleep(time.Millisecond)
				if atomic.LoadUint32(&sent) != 0 {
					t.Fatalf("send to full chan")
				}

				var v Value
				if !q.Recv(&v) {
					t.Fatal("failed to recv")
				}
				flushrecv(q)
				runtime.Gosched()
				if atomic.LoadUint32(&sent) != 1 {
					runtime.Gosched()
					time.Sleep(time.Millisecond)
					if atomic.LoadUint32(&sent) != 1 {
						t.Fatalf("did not unblock")
					}
				}
			}
		})
	}
}
func testMPSC(t *testing.T, ctor func(int) MPSC) {
	t.Run("Basic", func(t *testing.T) {
		for _, size := range TestSizes {
			q := ctor(size)
			count := Value(size/2) + 1

			result := async.All(func() error {
				return async.SpawnWithResult(TestProcs, func(id int) error {
					for i := Value(0); i < count; i++ {
						if !q.Send(Value(id)<<32 | i) {
							return fmt.Errorf("%v: failed to send %v", size, i)
						}
					}
					flushsend(q)
					return nil
				}).WaitError()
			}, func() error {
				var exps [TestProcs]Value
				for i := Value(0); i < count*TestProcs; i++ {
					var val Value
					if !q.Recv(&val) {
						return fmt.Errorf("%v: failed to get", size)
					}
					id, got := val>>32, val&0xFFFFFFFF
					exp := exps[id]
					exps[id]++
					if exp != got {
						return fmt.Errorf("%v: invalid order got %v, expected %v", size, got, exp)
					}
				}
				return nil
			})

			if errs := result.Wait(); errs != nil {
				t.Fatal(errs)
			}
		}
	})

	if _, isBatchReceiver := ctor(1).(BatchReceiver); isBatchReceiver {
		t.Run("Batch", func(t *testing.T) {
			for _, size := range TestSizes {
				q := ctor(size)
				count := Value(size/2) + 1

				result := async.All(func() error {
					return async.SpawnWithResult(TestProcs, func(id int) error {
						for i := Value(0); i < count; i++ {
							if !q.Send(Value(id)<<32 | i) {
								return fmt.Errorf("%v: failed to send %v", size, i)
							}
						}
						flushsend(q)
						return nil
					}).WaitError()
				}, func() error {
					br := q.(BatchReceiver)
					var err error
					var exps [TestProcs]Value
					total := count * TestProcs
					for total > 0 {
						br.BatchRecv(func(val Value) {
							id, got := val>>32, val&0xFFFFFFFF
							exp := exps[id]
							exps[id]++

							if got != exp {
								if err == nil {
									err = fmt.Errorf("%v: invalid value got %v, expected %v", size, got, exp)
								}
								tryclose(br)
								return
							}

							total--
							if total == 0 {
								tryclose(br)
							} else if total < 0 {
								if err == nil {
									err = fmt.Errorf("%v: invalid value got %v, expected %v", size, got, exp)
								}
								tryclose(br)
							}
						})
						if err != nil {
							return err
						}
					}
					return nil
				})
				if errs := result.Wait(); errs != nil {
					t.Fatal(errs)
				}
			}
		})
	}
}
func testSPMC(t *testing.T, ctor func(int) SPMC) {
	t.Run("Basic", func(t *testing.T) {
		for _, size := range TestSizes {
			q := ctor(size)
			count := Value(size/2) + 1

			result := async.All(func() error {
				for i := Value(0); i < count*TestProcs; i++ {
					if !q.Send(i) {
						return fmt.Errorf("%v: failed to send %v", size, i)
					}
				}
				flushsend(q)
				return nil
			}, func() error {
				return async.SpawnWithResult(TestProcs, func(int) error {
					var lastexp Value = -1
					for i := Value(0); i < count; i++ {
						var got Value
						if !q.Recv(&got) {
							return fmt.Errorf("%v: failed to get", size)
						}
						exp := lastexp
						lastexp = got
						if got <= exp {
							return fmt.Errorf("%v: invalid order got %v, expected at least %v", size, got, exp)
						}
					}
					return nil
				}).WaitError()
			})

			if errs := result.Wait(); errs != nil {
				t.Fatal(errs)
			}
		}
	})
	// TODO: Batch
}
func testMPMC(t *testing.T, ctor func(int) MPMC) {
	t.Run("Basic", func(t *testing.T) {
		for _, size := range TestSizes {
			t.Run(strconv.Itoa(size), func(t *testing.T) {
				q := ctor(size)
				count := Value(size/2) + 1

				result := async.All(func() error {
					return async.SpawnWithResult(TestProcs, func(id int) error {
						for i := Value(0); i < count; i++ {
							if !q.Send(Value(id)<<32 | i) {
								return fmt.Errorf("%v: failed to send %v", size, i)
							}
						}
						flushsend(q)
						return nil
					}).WaitError()
				}, func() error {
					return async.SpawnWithResult(TestProcs, func(int) error {
						var exps [TestProcs]Value
						for i := range exps {
							exps[i] = -1
						}
						for i := Value(0); i < count; i++ {
							var val Value
							if !q.Recv(&val) {
								return fmt.Errorf("%v: failed to get", size)
							}
							id, got := val>>32, val&0xFFFFFFFF
							exp := exps[id]
							exps[id] = got
							if got <= exp {
								return fmt.Errorf("%v: invalid order got %v, expected %v", size, got, exp)
							}
						}
						return nil
					}).WaitError()
				})

				if errs := result.Wait(); errs != nil {
					t.Fatal(errs)
				}
			})
		}
	})

	if _, isBatchReceiver := ctor(1).(BatchReceiver); isBatchReceiver {
		t.Run("Batch", func(t *testing.T) {
			for _, size := range TestSizes {
				q := ctor(size)
				count := Value(size/2) + 1

				result := async.All(func() error {
					return async.SpawnWithResult(TestProcs, func(id int) error {
						for i := Value(0); i < count; i++ {
							if !q.Send(Value(id)<<32 | i) {
								return fmt.Errorf("%v: failed to send %v", size, i)
							}
						}
						flushsend(q)
						return nil
					}).WaitError()
				}, func() error {
					br := q.(BatchReceiver)
					var err error
					var exps [TestProcs]Value
					total := count * TestProcs
					for total > 0 {
						br.BatchRecv(func(val Value) {
							id, got := val>>32, val&0xFFFFFFFF
							exp := exps[id]
							exps[id] = got
							if got <= exp {
								if err == nil {
									err = fmt.Errorf("%v: invalid value got %v, expected %v", size, got, exp)
								}
								tryclose(br)
								return
							}

							total--
							if total == 0 {
								tryclose(br)
							} else if total < 0 {
								if err == nil {
									err = fmt.Errorf("%v: invalid value got %v, expected %v", size, got, exp)
								}
								tryclose(br)
							}
						})
						if err != nil {
							return err
						}
					}
					return nil
				})
				if errs := result.Wait(); errs != nil {
					t.Fatal(errs)
				}
			}
		})
	}
}

func testNonblockingSPSC(t *testing.T, ctor func(int) NonblockingSPSC) {
	t.Run("Basic", func(t *testing.T) {
		for _, size := range TestSizes {
			q := ctor(size)
			count := Value(size/2) + 1

			result := async.All(func() error {
				for i := Value(0); i < count; i++ {
					if !mustSend(q, i) {
						return fmt.Errorf("%v: failed to send %v", size, i)
					}
				}
				flushsend(q)
				return nil
			}, func() error {
				for exp := Value(0); exp < Value(exp); exp++ {
					var got Value
					if !mustRecv(q, &got) || got != exp {
						return fmt.Errorf("%v: invalid value got %v, expected %v", size, got, exp)
					}
				}
				return nil
			})

			if errs := result.Wait(); errs != nil {
				t.Fatal(errs)
			}
		}
	})

	if _, isBounded := ctor(1).(Bounded); isBounded {
		t.Run("NonblockOnFull", func(t *testing.T) {
			for _, size := range TestSizes {
				q := ctor(size)
				capacity := q.(Bounded).Cap()
				for i := 0; i < capacity; i++ {
					if !q.TrySend(0) {
						t.Fatal("failed to send")
					}
				}
				flushsend(q)
				if q.TrySend(0) {
					t.Fatal("send succeeded")
				}
				flushsend(q)
			}
		})
	}
}
func testNonblockingMPSC(t *testing.T, ctor func(int) NonblockingMPSC) {
	t.Run("Basic", func(t *testing.T) {
		for _, size := range TestSizes {
			q := ctor(size)
			count := Value(size/2) + 1

			result := async.All(func() error {
				return async.SpawnWithResult(TestProcs, func(id int) error {
					for i := Value(0); i < count; i++ {
						if !mustSend(q, Value(id)<<32|i) {
							return fmt.Errorf("%v: failed to send %v", size, i)
						}
					}
					flushsend(q)
					return nil
				}).WaitError()
			}, func() error {
				var exps [TestProcs]Value
				for i := Value(0); i < count*TestProcs; i++ {
					var val Value
					if !mustRecv(q, &val) {
						return fmt.Errorf("%v: failed to get", size)
					}
					id, got := val>>32, val&0xFFFFFFFF
					exp := exps[id]
					exps[id]++
					if got != exp {
						return fmt.Errorf("%v: invalid order got %v, expected %v", size, got, exp)
					}
				}
				return nil
			})

			if errs := result.Wait(); errs != nil {
				t.Fatal(errs)
			}
		}
	})
}
func testNonblockingSPMC(t *testing.T, ctor func(int) NonblockingSPMC) {
	t.Run("Basic", func(t *testing.T) {
		for _, size := range TestSizes {
			q := ctor(size)
			count := Value(size/2) + 1

			result := async.All(func() error {
				for i := Value(0); i < count*TestProcs; i++ {
					if !mustSend(q, i) {
						return fmt.Errorf("%v: failed to send %v", size, i)
					}
				}
				flushsend(q)
				return nil
			}, func() error {
				return async.SpawnWithResult(TestProcs, func(int) error {
					var lastexp Value = -1
					for i := Value(0); i < count; i++ {
						var got Value
						if !mustRecv(q, &got) {
							return fmt.Errorf("%v: failed to get", size)
						}
						exp := lastexp
						lastexp = got
						if got <= exp {
							return fmt.Errorf("%v: invalid order got %v, expected %v", size, got, exp)
						}
					}
					return nil
				}).WaitError()
			})

			if errs := result.Wait(); errs != nil {
				t.Fatal(errs)
			}
		}
	})
}
func testNonblockingMPMC(t *testing.T, ctor func(int) NonblockingMPMC) {
	t.Run("Basic", func(t *testing.T) {
		for _, size := range TestSizes {
			q := ctor(size)
			count := Value(size/2) + 1

			result := async.All(func() error {
				return async.SpawnWithResult(TestProcs, func(id int) error {
					for i := Value(0); i < count; i++ {
						if !mustSend(q, Value(id)<<32|i) {
							return fmt.Errorf("%v: failed to send %v", size, i)
						}
					}
					flushsend(q)
					return nil
				}).WaitError()
			}, func() error {
				return async.SpawnWithResult(TestProcs, func(int) error {
					var exps [TestProcs]Value
					for i := range exps {
						exps[i] = -1
					}
					for i := Value(0); i < count; i++ {
						var val Value
						if !mustRecv(q, &val) {
							return fmt.Errorf("%v: failed to get", size)
						}
						id, got := val>>32, val&0xFFFFFFFF
						exp := exps[id]
						exps[id] = got
						if got <= exp {
							return fmt.Errorf("%v: invalid order got %v, expected %v", size, got, exp)
						}
					}
					return nil
				}).WaitError()
			})

			if errs := result.Wait(); errs != nil {
				t.Fatal(errs)
			}
		}
	})
}
