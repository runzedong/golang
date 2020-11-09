package thumbnail

import (
	"log"
	"os"
	"sync"
)

// ImageFile reads an image from infile and writes
// a thumbnail-size version of it in the same directory.
// It returns the generated file name, e.g., "foo.thumb.jpg".
func ImageFile(infile string) (string, error) {
	return "", nil
}

// MakeThumbnails makes thumbnails for each file received from the channel.
// It returns the number of bytes occupied by the files it creates.
func MakeThumbnails(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		// process
		go func(f string) {
			defer wg.Done()
			thumb, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb) // OK to ignore error
			sizes <- info.Size()
		}(f)
	}
	// wait to close the sizes
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for s := range sizes {
		total += s
	}
	return total
}
