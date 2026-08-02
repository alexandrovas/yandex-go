package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	ID      int    `json:"-"`
	Name    string `json:"name"`
	Company string `json:"company,omitempty"`
}

// func main() {
// 	foo := []Data{
// 		{
// 			ID:   10,
// 			Name: "Gopher",
// 		},
// 		{
// 			Name:    "Вася",
// 			Company: "Яндекс",
// 		},
// 	}
// 	out, err := json.Marshal(foo)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(string(out))
// }

func main() {
	var unknownJSON any
	err := json.Unmarshal([]byte(`[0, 10, 30]`), &unknownJSON)
	fmt.Printf("%T, %v, %v", unknownJSON, unknownJSON, err)
}
