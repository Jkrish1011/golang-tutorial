package mutex

/*

The standard library also exposes a sync.RWMutex
It has Lock() and Unlock() methods

Also it has RLock() and RUnlock()
The sync.RWMutex can help with performance if we have a read-intensive process.
Many goroutines can safely read from the map at the same time(multiple RLock()) calls can happen simultaneously.
However, only 1 goroutine can hold a Lock() and all RLock()'s will also be excluded.

*/

// Assignment
/*
Let's update our same code from the last assignment, but this time we can speed it up by allowing readers to read from
 the map concurrently.

Run the new test suite. You'll notice that it hangs forever and you'll need to cancel it.

Update the val method to only lock the mutex for reading.

*/

import (
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mu     *sync.RWMutex
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mu.RLock()
	defer sc.mu.RUnlock()
	return sc.counts[key]
}

// don't touch below this line

func (sc safeCounter) slowIncrement(key string) {
	tempCounter := sc.counts[key]
	time.Sleep(time.Microsecond)
	tempCounter++
	sc.counts[key] = tempCounter
}
