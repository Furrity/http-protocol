package main

import (
	"fmt"
	"os"
)


func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Errorf("failed to load file into memory: %s", err)
		return
	}
	data = make([]byte, 8)

	for file.Read()

}
