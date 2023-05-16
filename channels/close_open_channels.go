package main

import "fmt"

// Few points to keep in mind :
// - Values can be received from a closed channel but not sent.
// - Goroutine never waits in case of receiving from closed channel.
// - Sending value on closed channel will cause goroutine to panic.

// If we try to send data on a closed channel goroutine will panic.
func case1() {
	var c chan int = make(chan int, 10)
	close(c)
	c <- 10
}

// OUTPUT
// panic: send on closed channel

// goroutine 1 [running]:
// main.case1(...)
//         /mnt/c/Users/rynm/code/go/golang_concepts/channels/close_open_channels.go:6
// main.CloseOpenChannel(...)
//         /mnt/c/Users/rynm/code/go/golang_concepts/channels/close_open_channels.go:12
// main.main()
//         /mnt/c/Users/rynm/code/go/golang_concepts/channels/main.go:4 +0x45
// exit status 2

// -----------------------------------------------------

// If we try to receive a value from a close channel we will get zero value of
// the channel's datatype. But how we will determine that channel is closed
// and some other goroutine has not put actual zero value on channel, see case 3.
func case2() {
	var c chan int = make(chan int, 10)
	close(c)
	fmt.Println(<-c)
}

// OUTPUT
// 0

// -----------------------------------------------------

// Just like a map receiving from a channel will return two
// values if there are two variables on left hand side. The
// second variable will be a boolean, if value of the boolean
// is 'true' this means channel may be still open or some values
// are present in the channel's buffer and the value received
// is the actual value sent on channel even if it is zero value of that
// datatype.
// If channel is closed and no value is present in the buffer then
// the boolen will be false and received value will be zero value of
// that datatype.
func case3() {
	var c chan int = make(chan int, 10)
	c <- 10
	value, isChannelOpen := <-c
	fmt.Println(value, isChannelOpen)
	close(c)

	value, isChannelOpen = <-c
	fmt.Println(value, isChannelOpen)
}

// OUTPUT
// 10 true
// 0 false

func CloseOpenChannel() {
}
