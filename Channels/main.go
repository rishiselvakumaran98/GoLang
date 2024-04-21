package main

import (
	"fmt"

)
// Channels are a way to communicate between goroutines
// Channels are typed and can only send and receive the type they are declared with
// Channels are unbuffered by default
// Channels are a reference type and can be passed around like pointers to other functions
// Chennels are thread safe, and listening to a channel will block the current goroutine until a value is received
func main(){
	// think of channels as underlying arrays
	var c = make(chan int)
	// c <- 1 // unbuffered array has enough room for 1 value
	// retrive value using similar syntax
	// var i = <- c // i = 1 
	// This can technically cause the program to hang if the channel is not closed
	// Lets see how we can use this more effectively
	go process(c)
	// for i:=0; i<5; i++ {
	// 	fmt.Println(<-c)
	// }
	// Same as stating the following
	for i := range c {
		fmt.Println(i)
	}
}

func process (c chan int){
	for i:=0; i<5; i++ {
		c <- i
	}
	close(c) // we need to close the channel to avoid a deadlock
}

// we can create a buffered channel that can hold more than one value
// The buffer channel can wait upto 5 values before it blocks
