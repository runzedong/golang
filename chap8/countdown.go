package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// abort program
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	// count down ticker
	fmt.Println("Starting count down..")
	ticker := time.Tick(1 * time.Second)
	for count := 10; count > 0; count-- {
		select {
		case <-ticker:
			fmt.Println(count)
		case <-abort:
			fmt.Println("Abort launching...")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Launch...")
}
