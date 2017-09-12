package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand := func() interface{} { return rand.Intn(50000000) }

	repeatFn := func(
		done <-chan interface{},
		fn func() interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	toInt := func(
		done <-chan interface{},
		valueStream <-chan interface{},
	) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int):
				}
			}
		}()
		return intStream
	}

	primeFinder := func(
		done <-chan interface{},
		valueStream <-chan int,
	) <-chan interface{} {
		intStream := make(chan interface{})
		go func() {
			defer close(intStream)
			var isPrime bool
			for v := range valueStream {
				if v <= 0 {
					continue
				}

				isPrime = true
				for i := 2; i < v; i++ {
					if v%i == 0 {
						isPrime = false
						break
					}
				}

				if isPrime == false {
					continue
				}

				select {
				case <-done:
					return
				case intStream <- v:
				}
			}
		}()
		return intStream
	}

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := toInt(done, repeatFn(done, rand))
	fmt.Println("Primes:")
	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}
