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

	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if ok == false {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	intStream := intStreamFn(done, 10)

	for val := range orDone(done, intStream) {
		fmt.Println(val)
	}
}
