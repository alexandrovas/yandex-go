package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var buf bytes.Buffer
	logger := log.New(&buf, "mylog: ", log.Lmsgprefix)

	logger.Println("Hello, world!")
	logger.Println("Goodbye")

	fmt.Print(&buf)
}
