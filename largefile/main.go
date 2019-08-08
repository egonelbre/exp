package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"

	mmap "github.com/edsrzf/mmap-go"
)

func main() {
	workerCount := flag.Int("workers", runtime.NumCPU(), "how many workers to run concurrently")
	chunkSize := flag.Int("chunk-size", 1<<20, "how large a single chunk should be")
	flag.Parse()

	start := time.Now()
	defer func() { fmt.Println("time:", time.Since(start)) }()

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer logError(file.Close)

	mmap, err := mmap.Map(file, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer logError(mmap.Unmap)

	type Result struct {
		Lines int64
		Names map[string]int
		Dates map[int]int
	}

	newResult := func() *Result {
		return &Result{
			Lines: 0,
			Names: make(map[string]int, 1<<10),
			Dates: make(map[int]int, 1<<10),
		}
	}

	var workers sync.WaitGroup
	chunks := make(chan []byte, *workerCount*64)
	partials := make(chan *Result, *workerCount)

	go func() { // chunker
		defer close(chunks)

		start := 0
		for {
			tail := start + *chunkSize
			if tail > len(mmap) {
				chunks <- mmap[start:]
				break
			}

			offset := bytes.IndexByte(mmap[tail:], '\n')
			if offset < 0 {
				chunks <- mmap[start:]
				break
			}

			tail += offset
			chunks <- mmap[start:tail]
			start = tail
		}
	}()

	workers.Add(*workerCount)
	for i := 0; i < *workerCount; i++ {
		go func() {
			defer workers.Done()
			partial := newResult()

			for chunk := range chunks {
				for len(chunk) > 0 {
					lineEnd := bytes.IndexByte(chunk, '\n')
					if lineEnd < 0 {
						lineEnd = len(chunk)
					} else {
						lineEnd++
					}

					var line []byte
					line, chunk = chunk[:lineEnd], chunk[lineEnd:]
					if len(line) <= 1 {
						continue
					}

					partial.Lines++

					// C00384818|N|M2|P|201702039042412114|15|IND|DENTON, DAVID|WOONSOCKET|RI|028956146|CVS HEALTH|EVP & CFO, CVS HEALTH|01122017|208||2017020211435-885|1147467|||4020820171370030293
					split := bytes.SplitN(line, []byte("|"), 9)
					if len(split) <= 7 {
						fmt.Println(string(line), len(line), len(split))
						panic("adsf")
					}

					if len(split[7]) > 0 {
						partial.Names[string(split[7])]++
					}

					date, _ := strconv.Atoi(string(split[4][:6]))
					partial.Dates[date]++
				}
			}

			partials <- partial
		}()
	}

	// close partials, when all workers have finished
	go func() {
		workers.Wait()
		close(partials)
	}()

	result := newResult()
	for partial := range partials {
		result.Lines += partial.Lines
		for name, count := range partial.Names {
			result.Names[name] += count
		}
		for date, count := range partial.Dates {
			result.Dates[date] += count
		}
	}

	fmt.Println("Line count", result.Lines)
	fmt.Println("Names count", len(result.Names))
	fmt.Println("Dates count", len(result.Dates))
}

func logError(fn func() error) {
	if err := fn(); err != nil {
		log.Println(err)
	}
}
