// test10 project main.go
package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case c <- 1:
		case c <- 2:
		}
	}

	ch := make(chan bool)
	select {
	case v := <-ch:
		fmt.Println(v)
	case v := <-time.After(3 * time.Second):
		fmt.Println("timeout", v)
	}

}
