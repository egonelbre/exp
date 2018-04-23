package queue_test

import (
	"runtime"
	"strconv"
	"sync/atomic"
	"testing"
	"time"

	"github.com/egonelbre/exp/queue"
)

func TestChan(t *testing.T) {
	defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(4))
	N := 200
	if testing.Short() {
		N = 20
	}
	for chanCap := 1; chanCap < N; chanCap++ {
		t.Run("Cap"+strconv.Itoa(chanCap), func(t *testing.T) {
			t.Run("RecvOnEmpty", func(t *testing.T) {
				// Ensure that receive from empty chan blocks.
				c := queue.NewMPMC(chanCap)
				recv1 := false
				go func() {
					_, _ = c.RecvValue()
					recv1 = true
				}()
				time.Sleep(time.Millisecond)
				if recv1 {
					t.Fatalf("receive from empty chan")
				}

				c.Send(0)
				c.Send(0)
			})

			t.Run("BlockOnSendFull", func(t *testing.T) {
				// Ensure that send to full chan blocks.
				c := queue.NewMPMC(chanCap)
				for i := 0; i < chanCap; i++ {
					c.Send(i)
				}
				sent := uint32(0)
				go func() {
					c.Send(0)
					atomic.StoreUint32(&sent, 1)
				}()
				time.Sleep(time.Millisecond)
				if atomic.LoadUint32(&sent) != 0 {
					t.Fatalf("send to full chan")
				}
				_, _ = c.RecvValue()
			})

			t.Run("Long", func(t *testing.T) {
				// Send 100 integers,
				// ensure that we receive them non-corrupted in FIFO order.
				c := queue.NewMPMC(chanCap)

				t.Run("Send100", func(t *testing.T) {
					go func() {
						for i := 0; i < 100; i++ {
							c.Send(i)
						}
					}()
					for i := 0; i < 100; i++ {
						v, ok := c.RecvValue()
						if !ok {
							t.Fatalf("receive failed, expected %v", i)
						}
						if v != i {
							t.Fatalf("received %v, expected %v", v, i)
						}
					}
				})

				t.Run("Send1000x4", func(t *testing.T) {
					// Send 1000 integers in 4 goroutines,
					// ensure that we receive what we send.
					const P = 4
					const L = 1000
					for p := 0; p < P; p++ {
						go func() {
							for i := 0; i < L; i++ {
								c.Send(i)
							}
						}()
					}

					done := make(chan map[int]int)
					for p := 0; p < P; p++ {
						go func() {
							recv := make(map[int]int)
							for i := 0; i < L; i++ {
								v, _ := c.RecvValue()
								recv[v] = recv[v] + 1
							}
							done <- recv
						}()
					}
					recv := make(map[int]int)
					for p := 0; p < P; p++ {
						for k, v := range <-done {
							recv[k] = recv[k] + v
						}
					}
					if len(recv) != L {
						t.Fatalf("received %v values, expected %v", len(recv), L)
					}
					for _, v := range recv {
						if v != P {
							t.Fatalf("received %v values, expected %v", v, P)
						}
					}
				})
			})
		})
	}
}
