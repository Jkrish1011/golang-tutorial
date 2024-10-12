package mutex

/*
Mutex in Go;
------------
Mutexes allow us to lock access to data. This ensures that we can control which goroutines can access certain data at which time!

Go's standard library provides a built-in implementation of a mutex with the sync.Mutex type and two methods:
1/ Lock()
2/ Unlock()

We can protect a block of code by surrounding it with a call to Lock and Unlock as shown below:
Note: It's good practice to structure the protected code within a function so that defer can be used to ensure to unlock the mutex.

```
func protected() {
	mux.Lock()
	defer mux.Unlock()
	// Rest of the function is protected
	// Any other calls to the `mux.Lock()` will block
}

Note;
Mutex are used when we want to use a dangerous resource across many goroutines.

Important;

Maps are not thread-safe;
-------------------------
Maps are not safe for concurrent use. If we have multiple goroutines accessing the same map, and at least one of them
is writing to the map, it should be done that we have to lock the map with a mutex.
```

*/

// Assignment - 1

/*
Our safeCounter struct is unsafe! Update the inc() and val() methods so that they utilize the safeCounter's mutex
to ensure that the map is not accessed by multiple goroutines at the same time.


*/

// import (
// 	"sync"
// 	"time"
// )

// type safeCounter struct {
// 	counts map[string]int
// 	mu     *sync.Mutex
// }

// func (sc safeCounter) inc(key string) {
// 	sc.mu.Lock()
// 	defer sc.mu.Unlock()
// 	sc.slowIncrement(key)
// }

// func (sc safeCounter) val(key string) int {
// 	sc.mu.Lock()
// 	defer sc.mu.Unlock()
// 	return sc.slowVal(key)
// }

// // don't touch below this line

// func (sc safeCounter) slowIncrement(key string) {
// 	tempCounter := sc.counts[key]
// 	time.Sleep(time.Microsecond)
// 	tempCounter++
// 	sc.counts[key] = tempCounter
// }

// func (sc safeCounter) slowVal(key string) int {
// 	time.Sleep(time.Microsecond)
// 	return sc.counts[key]
// }
