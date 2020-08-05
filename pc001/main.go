package main

import (
	"fmt"
	"time"
)

func main() {
	worksChan := make(chan int, 10)
	i := 0
	go func() {
		for i < 10 {
			worksChan <- i
			time.Sleep(1 * time.Second)
			i++
		}
	}()
	go func() {
		for g := range worksChan {
			fmt.Println(g)
		}
	}()
	time.Sleep(30 * time.Second)
}
