package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	event := func(c *sync.Cond, fn func()) {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			wg.Done()
			for i := 0; i <= 5; i++ {
				c.L.Lock()
				c.Wait()
				c.L.Unlock()

				fn()
			}
		}()
		wg.Wait()
	}
	c := sync.NewCond(&sync.Mutex{})

	event(c, func() {
		fmt.Println("event1")
	})
	event(c, func() {
		fmt.Println("event2")
	})

	for i := 0; i <= 5; i++ {
		c.Broadcast()
		time.Sleep(1 * time.Second)
	}
}
