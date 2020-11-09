package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	workList := make(chan []string)
	unseenLinks := make(chan string)

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main gorotine de-duplicate links and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for work := range workList {
		for _, link := range work {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

// crawl is a helper function # NOT IMPLEMENTATED
func crawl(link string) []string {
	fmt.Println(link)
	list, err := links.Extract(link)
	if err != nil {
		log.Fatal(err)
	}
	return list, err
}
