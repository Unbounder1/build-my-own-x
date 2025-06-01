package main

import "fmt"

func main() {

	myChannel := make(chan string) // Channels are FIFO

	go func() {
		myChannel <- "data"
	}()

	msg := <-myChannel // blocking

	fmt.Println(msg)
}
