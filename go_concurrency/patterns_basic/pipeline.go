package main

import "fmt"

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// input
	nums := []int{2, 3, 4, 7, 1}
	// stage 1
	dataChannel := sliceToChannel(nums)
	//stage 2
	finalChannel := sq(dataChannel)
	//stage 3
	for n := range finalChannel { // range over unbuffered channel until its closed
		fmt.Println(n)
	}
}
