package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/runzedong/golang/memo"
)

func getHTTPBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	incomingURLs := []string{
		"https://golang.org",
		"https://play.golang.org",
		"https://golang.org",
		"http://gopl.io",
		"http://gopl.io",
		"https://play.golang.org",
	}

	m := memo.New(getHTTPBody)

	var wg sync.WaitGroup
	for _, url := range incomingURLs {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	wg.Wait()
	fmt.Println("All works done")
}
