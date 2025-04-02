package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

const port string = ":42069"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to load file into memory: %s\n", err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Could not establish connection:", err)
		}
		fmt.Println("A new connection has been accepted")
		c := getLinesChannel(conn)
		for line := range c {
			fmt.Println(line)
		}
		fmt.Println("Connection has been closed")

	}

}

func getLinesChannel(file io.ReadCloser) <-chan string {
	c := make(chan string)

	go func() {
		var curr string
		defer file.Close()
		defer close(c)
		for {
			buffer := make([]byte, 8)
			n, err := file.Read(buffer)
			if errors.Is(err, io.EOF) {
				break
			} else if err != nil {
				log.Fatalf("Failed to read 8 bytes: %v\n", err)
			}
			parts := strings.Split(string(buffer[:n]), "\n")
			for i, part := range parts {
				curr += part
				if i != len(parts)-1 {
					c <- curr
					curr = ""
				}
			}
		}
	}()
	return c

}
