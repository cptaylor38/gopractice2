package main

import (
	"fmt"
)

func ChannelsMain() {
	//Channels allow the passing of data between goroutines.
	ch := make(chan int, 50)
	//Adding a second paramater to the make function and provide
	//		an integer tells go to create a channel that has an internal data
	//			store that can store int # of integers.
	WG.Add(2)
	// go func() {
	// 	i := <-ch //pulling data into the channel by establishing
	// 	// 				variable i = <-ch
	// 	fmt.Println(i)
	// 	WG.Done()
	// }() // this routine is the receiving goroutine
	// go func() {
	// 	ch <- 42 //sending 42 into the channel <- 42
	// 	// By default channels are unbuffered, which means
	// 	// the channel flow? will halt until something receives the data.

	// 	WG.Done()
	// }() // this routine is the sending goroutine

	// go func(ch <-chan int) {
	// 	//channel that only receives
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	WG.Done()
	// }(ch)

	// go func(ch chan<- int) {
	// 	//send only channel
	// 	ch <- 42
	// 	WG.Done()
	// }(ch)

	go func(ch <-chan int) {
		for i := range ch {
			//instead of ranging over a collection
			//ranging over a channel
			//different syntax. when ranging over a channel
			// the value you pull is the value out of the channel
			fmt.Println(i)
		}
		WG.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch) // to close the channel and signal there is no more
		//   data being sent
		// Have to be careful when closing a channel
		// Sending message on closed channel causes a panic.
		WG.Done()
	}(ch)

	WG.Wait()

}

//A buffer channel is designed so that if the sender or receiver operate
//	at a different frequency than the other side. When sender or
// 	receiver need more time to process.

// Can gracefully close channel by:

// func main() {
// 	defer func(){
// 		close(logCh)
// 		//fake channel being closed with defer, to run JUST before
// 		//the main function closes down
// 	}()
// }

// Or you can use a select statement
// var doneCh = make(chan struct{}) - Signal only channel
// struct with no fields is unique in that it requires 0 memory allocation
// channel set up with empty struct can't receive any data through
// except that a message was sent

//select statement - entire statement blocks until a message is received
// on one of the channels it is listening for.

//if we get message from done channel, we break out of endless for loop?
//work like switch statements but only in context of channels.
// allow a go routine to monitor several channels at the same time
// if they block, all channels are blocked.
// if no messages on any channel, blocks by default.
// if it receives message, processes on correct channel
// if multiple channels receive values simultaneously, behavior is undefined.

//adding a default case changes it from a blocked select statement

// Example of a blocked select statement
// func logger() {
// 	for {
// 		select {
// 		case entry := <- logCh:
// 			fmt.Printf("%v - [%v]%v\n", entry.time.Format("2006-01"))
// 		case <- doneCh:
// 			break
// 		}
// 	}
// }
