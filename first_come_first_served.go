//go:build ignore

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := fanln(generator("Hello"), generator("Bye"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func fanln(ch1, ch2 <-chan string) <-chan string {
	new_ch := make(chan string)
	go func() {
		for {
			select {
			case s := <-ch1:
				new_ch <- s
			case s := <-ch2:
				new_ch <- s
			default:
				update_progressbar()
			}
		}
	}()
	return new_ch
}

func generator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}
