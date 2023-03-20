
package main

import (
  "fmt"
  "sync"
)

func main() {
	wg := &sync.WaitGroup{}
	c := make(chan string)
	wg.Add(2)
	go sender(wg, c)
	go reciver(wg, c)
	wg.Wait()
}

func sender(wg *sync.WaitGroup ,c chan string) {
	c <- "Hallo Empfänger"
  wg.Done()
}

func reciver(wg *sync.WaitGroup ,c chan string) {
	msg := <-c
	fmt.Println("Empfangen!\nNachricht:", msg)
  wg.Done()
}

