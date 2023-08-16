package main

import (
	"fmt"
	"io"
	"log"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {
	sh := shell.NewShell("localhost:5001")

	tokenURI := "QmZcH4YvBVVRJtdn4RdbaqgspFU8gH6P9vomDpBVpAL3u4/20"

	stream, err := sh.Cat(tokenURI)
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer stream.Close()

	data, err := io.ReadAll(stream)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("data: %s\n", data)
}
