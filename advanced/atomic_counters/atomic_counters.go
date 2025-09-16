package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Atomic (indivisible and uninterruptible) operations
// Atomic operations are used in concurrent programming to ensure
// that a particular operation on shared data is completed without interruption.
// This is crucial in multi-threaded environments where multiple threads
// may attempt to read or write shared data simultaneously, leading to race conditions.

type AtomicCounter struct {
	count int64
}

func (ac *AtomicCounter) increment() {
	atomic.AddInt64(&ac.count, 1) // Atomically increments the counter by 1
}

func (ac *AtomicCounter) getCount() int64 {
	return atomic.LoadInt64(&ac.count) // Atomically loads the current counter value
}

func main() {
	// Why atomic operations?
	// Performance: Atomic operations are generally faster than using mutexes
	//     because they avoid the overhead of locking and unlocking.
	//     They are implemented at the hardware level, making them more efficient
	//     for simple operations like incrementing a counter.
	// Low latency: Atomic operations can reduce latency in high-concurrency scenarios
	//     since they do not require context switching or blocking threads.
	//     This is particularly beneficial in real-time applications where responsiveness is critical.
	// Reduced contention: In scenarios with high contention for shared resources,
	//     atomic operations can help reduce bottlenecks that arise from thread contention
	//     on locks, leading to better overall throughput.
	// Simplicity: For simple operations (like incrementing a counter),
	//     atomic operations can be easier to implement and reason about compared to mutexes,
	//     which require careful handling to avoid deadlocks and other concurrency issues.

	// Use cases:
	//     * Window duration
	//     * Request counting
	//     * Resource pooling
	//     * Sequence generation
	//     * Reference counting
	//     * Caching
	//     * Rate limiting
	//     * Metrics collection
	//     * Lock-free data structures
	//     * Concurrent algorithms

	var wg sync.WaitGroup
	numGoRoutines := 10
	counter := &AtomicCounter{}

	for range numGoRoutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter.getCount()) // Should print 10000
}
