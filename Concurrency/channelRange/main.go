package main

import "fmt"

func main() { 
  ch := make(chan int)
  go func(c chan int){
    c <- 1
    c <- 2
    c <- 3
    close(c)
  }(ch)
  for n := range ch {
    fmt.Println(n)
  }
}
  