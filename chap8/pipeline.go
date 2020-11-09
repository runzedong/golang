package main

import (
	"fmt"
	"time"
)

func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
		time.Sleep(1 * time.Second)
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for r := range in {
		fmt.Println(r)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(naturals, squares)
	printer(squares)
}
