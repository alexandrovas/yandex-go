package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("https://practicum.yandex.ru")
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(response.Body)
	response.Body.Close()

	if len(body) > 512 {
		body = body[:512]
	}
	fmt.Printf("First 512 bytes of body:\n%s", string(body))

	// fmt.Printf("Status Code: %d\r\n", response.StatusCode)
	// for k, v := range response.Header {
	// 	// заголовок может иметь несколько значений,
	// 	// но для простоты запросим только первое
	// 	fmt.Printf("%s: %v\r\n", k, v[0])
	// }

}
