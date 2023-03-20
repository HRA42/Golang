package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go send_only(c)
	fmt.Println(<-c)
  go rec_only(c)
}

func rec_only(ch <-chan int) {
	// ...
	rec := <-ch
	fmt.Println(rec)
	// ...
}

func send_only(ch chan<- int) {
	//...
	ch <- 1
	//...
}
