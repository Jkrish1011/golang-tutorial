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
