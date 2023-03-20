package main

import (
	"fmt"
	"time"
)

func main() {
  ch := make(chan string)
  for i:=0; i<5; i++ {
    go zuhörer(ch)
    time.Sleep(time.Second)
  }
  ch <- "eine Nachricht"
  close(ch)
  time.Sleep(time.Second)
  fmt.Println("Ende!")
}

func zuhörer(ch chan string) {
  fmt.Println("warte...")
  fmt.Println(<-ch)
  fmt.Println("Goroutine Ende!")
}
