package main

import (
	"fmt"
	"log"
	"os"
)

var token = make(chan struct{}, 20)

func main() {
	workLists := make(chan []string)
	seen := make(map[string]bool)
	n := 0

	// Start with the command-line arguments.
	n++
	go func() { worklist <- os.Args[1:] }()

	// start crawl the work list
	for ; n > 0; n-- {
		list <- workLists
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					workLists <- crawl(link)
				}(link)
			}
		}
	}
}

func crawl(link string) []string {
	token <- struct{}{}
	fmt.Println(link)
	list, err := links.Extract(link)
	if err != nil {
		log.Fatal(err)
	}
	<-token
	return list
}
