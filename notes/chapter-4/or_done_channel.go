package main

import "fmt"

func main() {
	var (
		done chan interface{}
	)
	done = make(chan interface{})
	defer close(done)

	intStreamFn := func(done <-chan interface{}, num int) <-chan interface{} {
		ch := make(chan interface{})
		go func() {
			for i := 0; i < num; i++ {
				select {
				case <-done:
					break
				case ch <- i:
				}

			}
			close(ch)
		}()

		return ch
	}

	intStream := intStreamFn(done, 10)

loop:
	for {
		select {
		case <-done:
			break loop
		case val, ok := <-intStream:
			if ok == false {
				fmt.Println("not ok, may be channel close")
				break loop
			}

			fmt.Println(val)
		}
	}
}
