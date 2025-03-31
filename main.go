package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalf("failed to load file into memory: %s\n", err)
		return
	}

	defer file.Close()

	buffer := make([]byte, 8)

	for {
		n, err := file.Read(buffer)
		if err == io.EOF || n == 0 {
			break
		} else if err != nil {
			log.Fatalf("Failed to read 8 bytes: %v\n", err)
		}

		fmt.Printf("read: %s\n", string(buffer[:n]))
	}

}
