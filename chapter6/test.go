package main

import (
	"fmt"
)

func in(ch chan string) {
	ch <- "yes"
	//time.sleep(1)
	close(ch)
}

func main() {
	ch := make(chan string)

	go func() {
		ch <- "yes"
		//time.sleep(1)
		close(ch)
	}()

	for {
		out, ok := <-ch
		if !ok {
			fmt.Printf("not ok")
			return
		} else {
			fmt.Printf("out = %s\n", out)
		}
	}
}
