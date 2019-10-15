package main

import "fmt"

func main() {
	ch := make(chan string)
	// printer goroutine
	go func() {
		for message := range ch {
			fmt.Println(message)
		}
	}()

	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("Hello #%d", i)
	}
	close(ch)
}
