package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")

// cancelled listens to done channel close operation
var done = make(chan struct{})

// dirents returns the entries of directory dir.
// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

func main() {
	// Cancel traversal when input is detected.
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()

	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	// Launch the dir walker in parallel
	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	go func() {
		for _, r := range roots {
			wg.Add(1)
			go walkDir(r, fileSizes, &wg)
		}
	}()

	// Wait & close channel
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	// Print the results.
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			for range fileSizes {
				// do nothing but drain existing goroutine to finish
			}
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		}
	}
	printDiskUsage(nfiles, nbytes)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, fileSizes chan<- int64, wg *sync.WaitGroup) {
	defer wg.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subDir := filepath.Join(dir, entry.Name())
			go walkDir(subDir, fileSizes, wg)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	select {
	case <-done:
		return nil // cancelled
	case sema <- struct{}{}:
		// acquire lock
	}
	defer func() { <-sema }() // release lock
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f KB\n", nfiles, float64(nbytes)/1e3)
}
