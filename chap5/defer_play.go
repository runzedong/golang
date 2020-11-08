package main

import (
	"fmt"
	"log"
	"time"
)

func trace(msg string) func() {
	fmt.Println("Executing defer...")
	start := time.Now()
	log.Printf("enter: %s", msg)
	return func() { log.Printf("exit %s since time: %s", msg, time.Since(start)) }
}

// deferred func run after `return` statment have updated function's result variable
func double(x int) (result int) {
	defer func() {
		log.Printf("enter defer double\n")
		fmt.Printf("double(%d) = %d\n", x, result)
	}()
	fmt.Printf("Going to evaluate double result..\n")
	return x + x
}

// func main() {
// 	// here. this defer evaluate a function but not invoked.
// 	defer trace("main function")()
// 	fmt.Println("Program is running")
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("Program is finished")
// }
