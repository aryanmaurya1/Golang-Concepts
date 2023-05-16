package main

import (
	"fmt"
)

// There are three types on channels direction wise
// 1. receive-only channels (<-chan T)
// 2. send-only channels  (chan<- T)
// 3. bidirectional channels (chan T)

// func sendToReceiveOnlyChannel(c <-chan int) {
// 	c <- 100
// }

// OUTPUT
// channels/directional_channels.go:8:2: invalid operation: cannot send to receive-only channel c (variable of type <-chan int)

// -----------------------------------------------------

// func receiveFromSendOnlyChannel(c chan<- int) {
// 	fmt.Println(<-c)
// }

// OUTPUT
// channels/directional_channels.go:15:16: invalid operation: cannot receive from send-only channel c (variable of type chan<- int)

// -----------------------------------------------------


func sendAndReceiveBothFromChannel(c chan int){
	c <- 10
	fmt.Println(<-c)
}
// OUTPUT
// 10


func DirectionalChannels() {
	var c chan int = make(chan int, 1)
	sendAndReceiveBothFromChannel(c)
}