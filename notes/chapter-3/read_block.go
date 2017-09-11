package main

import (
	"fmt"
	"time"
)

func main() {
	var intStream <-chan int
	go func() {
		integer, ok := <-intStream
		fmt.Printf("(%v): %v", ok, integer)
	}()

	time.Sleep(time.Second)
}
