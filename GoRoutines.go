package main

import (
	"fmt"
	"sync"
)

var WG = sync.WaitGroup{}
var m = sync.RWMutex{} //read/write mutex
//RW mutex, many things can read this data, but only
//		one can write at a time
// If something makes a write request, it waits until all readers are done

var counter = 0

// A waitgroup synchronizes multiple go routines together.

func GoRoutineMain() {
	// runtime.GOMAXPROCS(100)

	// go sayHello()
	//green threads
	//schedule maps goroutines on operating system threads for a certain
	// period of time.
	// go routines can be reallocated quickly
	// var msg = "hello"
	// WG.Add(1)
	//Adding another go routine to the waitgroup
	// go func(msg string) {
	// 	fmt.Println(msg)
	// 	WG.Done()
	// 	//Letting the waitgroup know that this process is finished,
	// 	// to decrement the current count.
	// 	// From two go functions in the wait group, Main and sayHello
	// 	// to just Main
	// }(msg)
	// msg = "Goodbye"
	// WG.Wait()
	// Telling the waitgroup to wait until all go routines have completed?

	for i := 0; i < 10; i++ {
		WG.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()

		//locking them in the same context
		//asynchronously unlocking when operations/go routines are completed
	}
	WG.Wait()

}

func sayHello() {
	// m.RLock() //moved to main function after go routines called
	//obtain read lock on message
	//print message
	//then unlock
	fmt.Printf("hello, #%v\n", counter)
	m.RUnlock()
	// prints hello, #2
	// 		hello, #2
	// 		hello, #0
	// 		hello, #2
	// 		hello, #10
	// 		hello, #9
	//Go routines are competing against each other and
	// 	out of sync. So we will use a Mutex to sync them together.

	//A mutex is "a lock that the application is going to honor."

	WG.Done()
}

func increment() {
	// m.Lock() //moved to main function after go routines called
	//locking write
	//incrementing
	//unlocking write
	counter++
	m.Unlock()
	WG.Done()
}

//Best practices
// Don't create goroutines in libraries.
//		Let consumer control concurrency.
// When creating a goroutine, know how it will end.
// 		Avoid subtle memory leaks.
// Check for race conditions at compile time.
//		go run -race filename
