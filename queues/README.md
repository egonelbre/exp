Go Queues Test
==

Performance test of different concurrent queue implementations

--

Warning: not for production use!

-- 

Available Structure:

* CFifo : Channel based fifo
* LcLifo : list.List based lifo using chan for locking
* LcFifo : list.List based fifo using chan for locking
* LmLifo : list.List based lifo using mutex for locking
* LmFifo : list.List based fifo using mutex for locking
* ZLifo : lockfree lifo implementation (broken, ABA problem)
* ZFifo : lockfree fifo implementation (probably broken)
* ZcFifo : lockfree fifo implementation using chan based freelist (crashes)
* ZrFifo : lockfree fifo implementation using ring.Ring based freelist (probably broken)
* RmLifo : ring.Ring based lifo using mutex for locking
* RmFifo : ring.Ring based fifo using mutex for locking
* SmLifo : slice based lifo using mutex for locking
* SmFifo : slice based fifo using mutex for locking

Tests:

1. Single threaded
2. Add N, Remove N
3. N times (Add 1, Remove 1)
4. N/2 times (Add 2, Remove 1), Remove N/2
5. Add N/2, N/2 times (Add 1, Remove 2)

ZcFifo was excluded due to crash.

General recommendation use chan if it suits otherwise use slice + "chan as a lock".