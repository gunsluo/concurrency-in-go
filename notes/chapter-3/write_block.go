package main

import (
	"time"
)

func main() {
	var intStream chan<- int
	go func() {
		intStream <- 100
	}()

	time.Sleep(time.Second)
}
