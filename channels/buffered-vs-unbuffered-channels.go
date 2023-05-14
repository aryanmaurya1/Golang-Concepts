package main

import (
	"fmt"
	"math/rand"
	"time"
)

func putIntOnChannel(c chan<- int) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Int()
	c <- n
}

// This is an unbuffered channel, it cannot hold any value for later receiving.
// It means if we are sending any value in the channel there must a receiver goroutine
// present at the same time otherwise deadlock will occur. In this case receiver
// goroutine is 'main' which invoked this funtion.
func caseOne() {

	c := make(chan int) // len(c) --> 0

	go putIntOnChannel(c)
	n := <-c
	close(c)
	fmt.Println(n)
}

// OUTPUT
// 6946467344481185046

//  --------------------------------------------------------------------------------- //

// We are calling the 'putInChannel' which puts value on channel
// without setting any receiver goroutine beforehand, this will result in
// deadlock because channel is unbuffered and cannot store any value for
// later receiving.
func caseTwo() {
	c := make(chan int) // len(c) --> 0

	putIntOnChannel(c)
	n := <-c
	close(c)
	fmt.Println(n)
}

// OUTPUT
// goroutine 1 [chan send]:
// main.putIntOnChannel(0x4b3c98?)
//         /mnt/c/Users/rynm/code/go/Golang-Concepts/channels/main.go:12 +0xa5
// main.caseTwo()
//         /mnt/c/Users/rynm/code/go/Golang-Concepts/channels/main.go:33 +0x30
// main.main()
//         /mnt/c/Users/rynm/code/go/Golang-Concepts/channels/main.go:41 +0x1c
// exit status 2

//  --------------------------------------------------------------------------------- //

// We are calling the 'putInChannel' which puts value on channel
// without setting any receiver goroutine beforehand, but this time
// it won't cause any deadlock because 'c' is a buffered channel and
// it can store one value which can be received later.
func caseThree() {
	c := make(chan int, 1) // len(c) --> 1

	putIntOnChannel(c)
	n := <-c
	close(c)
	fmt.Println(n)
}

// OUTPUT
// 295296556638589988

//  --------------------------------------------------------------------------------- //

// In this case we are sending data concurrently through 3 goroutines.
// As channel is unbuffered we have to setup receiver goroutine before
// sending data on channel, in this case receiver goroutine is 'main'
// which only receives one value, prints it and terminates itself.
// This is what's going to happen, any one of three goroutines which
// are trying to send value will succeed and other two will block and
// wait for value to be received, at the same time 'main' goroutine will wait
// for any values comming on channel. As soon as it finds a value, it will
// print it and terminate itself. After termination of the 'main' goroutine
// all other blocked goroutines will be killed and program exits.
func caseFour() {
	c := make(chan int) // len(c) --> 0

	go putIntOnChannel(c)
	go putIntOnChannel(c)
	go putIntOnChannel(c)
	n := <-c
	close(c)
	fmt.Println(n)
}

// OUTPUT
// 2654207484686224392
