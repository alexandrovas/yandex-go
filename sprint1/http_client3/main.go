package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get(`http://localhost:8080`)
	if err != nil {
		panic(err)
	}

	for _, v := range resp.Cookies() {
		fmt.Printf("%s = %s\r\n", v.Name, v.Value)
	}

	resp.Body.Close()
}
