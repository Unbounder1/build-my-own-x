package main

import (
	"fmt"
	"time"
)

// read only channel
func doWork(done <-chan bool) {
	// infinitely running goroutine
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("Doing something") // goroutine runs infinitely looping
		}
	}
}

func main() {
	charChannel := make(chan string, 3) // buffer channel (capacity)
	// unbuffered channel -> goroutine sends data to the unbuffered channel -> goroutine waits for unbuffered channel (blocked) -> recieving go routine awaits data
	// synchronous /\

	// buffered channel -> goroutine communication is asynchronous -> not blocked when adding to channel

	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		select {
		case charChannel <- s:
		}
	}

	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}

	//donechannel pattern
	done := make(chan bool)

	go doWork(done)

	time.Sleep((time.Second * 3))

	close(done)
}
