package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"unsafe"

	"github.com/loov/hrtime"
)

/*
#include <stdlib.h>
#include <unistd.h>
#include <stdint.h>

typedef struct {
	int64_t closed;
	int64_t token;
	int64_t sequence;
	int64_t request;
	int64_t response;
} queue;

void service(volatile queue *q) {
	while(q->closed == 0){
		int64_t sequence;
		while(1){
			__atomic_load(&q->sequence, &sequence, __ATOMIC_ACQUIRE);
			if(sequence % 4 == 1) {
				break;
			}
		}

		int64_t result = q->request;
		q->response = result * 10;

		__atomic_add_fetch(&q->sequence, 1, __ATOMIC_RELEASE);
	}
}

*/
import "C"

type Queue struct {
	closed   int64
	token    int64
	sequence int64
	request  int64
	response int64
}

func main() {
	var queue Queue

	go func() {
		runtime.LockOSThread()
		C.service((*C.queue)(unsafe.Pointer(&queue)))
	}()

	const N = 1 << 20
	start := hrtime.TSC()
	for i := 0; i < N; i++ {
		request := int64(i)

		token := atomic.AddInt64(&queue.token, 4) - 4

		// wait for our turn
		for atomic.LoadInt64(&queue.sequence) != token {
		}

		queue.request = request

		atomic.StoreInt64(&queue.sequence, token+1)

		for atomic.LoadInt64(&queue.sequence) != token+2 {
		}

		response := queue.response
		atomic.StoreInt64(&queue.sequence, token+4)

		_ = response
	}
	stop := hrtime.TSC()
	fmt.Println((stop - start).ApproxDuration() / N)

	runtime.KeepAlive(queue)
}
