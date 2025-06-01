package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Because stream is unbuffered, the goroutine sends to the channel and is blocked until the receiving goroutine has
// recieved it, so using the repeatFunc won't do extra iterations while take is iterating up to n
// synchronous channel
func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	// fn is a function that returns type T
	stream := make(chan T)

	// runs goroutine seperate of func, func returns stream while goroutine still runs
	go func() {
		defer close(stream) //schedules when function returns
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}

func take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)
	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream: // read from stream <-, write into taken <-
			}
		}
	}()

	return taken
}

func primeFinder(done <-chan int, randIntStream <-chan int) <-chan int {
	isPrime := func(randomInt int) bool {
		for i := randomInt - 1; i > 1; i-- {
			if randomInt%i == 0 {
				return false
			}
		}
		return true
	}

	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			for {
				select {
				case <-done:
					return
				case randomInt := <-randIntStream:
					if isPrime(randomInt) {
						primes <- randomInt
					}
				}
			}
		}
	}()

	return primes
}

func fanIn[T any](done <-chan int, channels []<-chan T) <-chan T {
	var wg sync.WaitGroup
	fannedInStream := make(chan T)

	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fannedInStream <- i:
			}
		}
	}

	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}

	go func() {
		wg.Wait()
		close(fannedInStream) //closes fanstream when wg is done
	}()

	return fannedInStream
}

func main() {
	start := time.Now()
	done := make(chan int)
	defer close(done)

	randomNumFetcher := func() int { return rand.Intn(500000000) }
	randIntStream := repeatFunc(done, randomNumFetcher)

	//primeStream := primeFinder(done, randIntStream)

	// naive pipeline (synchronous)
	// for rando := range take(done, primeStream, 10) { //repeatFunc(done, randomNumFetcher)
	// 	fmt.Println(rando)
	// }

	// fan out
	CPUCount := runtime.NumCPU()
	primeFinderChannels := make([]<-chan int, CPUCount)
	for i := 0; i < CPUCount; i++ {
		primeFinderChannels[i] = primeFinder(done, randIntStream)
	}

	// fan in
	fannedinStream := fanIn(done, primeFinderChannels)

	for rando := range take(done, fannedinStream, 10) { //repeatFunc(done, randomNumFetcher)
		fmt.Println(rando)
	}

	fmt.Println(time.Since(start))
	return
}
