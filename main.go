package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	sh := shell.NewShell("localhost:5001")
	var wg sync.WaitGroup
	concurrency := 100 // Number of concurrent requests

	for i := 0; i < 10000; i += concurrency {
		wg.Add(concurrency)

		for j := 0; j < concurrency; j++ {
			index := i + j
			if index > 10000 {
				break // Ensure we don't exceed the total count
			}

			go func(i int) {
				defer wg.Done()

				tokenURI := fmt.Sprintf("QmZcH4YvBVVRJtdn4RdbaqgspFU8gH6P9vomDpBVpAL3u4/%d", i)
				stream, err := sh.Cat(tokenURI)

				if err != nil {
					log.Printf("error: %s", err)
					return
				}
				defer stream.Close()

				data, err := io.ReadAll(stream)
				if err != nil {
					log.Printf("error: %s", err)
					return
				}

				fileName := fmt.Sprintf("results/token-%d.json", i)
				file, err := os.Create(fileName)
				if err != nil {
					log.Printf("error: %s", err)
					return
				}
				defer file.Close()

				_, err = file.Write(data)
				if err != nil {
					log.Printf("error: %s", err)
					return
				}
			}(index)
		}

		wg.Wait() // Wait for the current batch of goroutines to complete
	}
}
