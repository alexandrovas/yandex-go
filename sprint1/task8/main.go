package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir(".."))
	mux.Handle("/golang/", http.StripPrefix("/golang/", fs))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./main.go")
	})
	mux.Handle("/search/", http.RedirectHandler("https://yandex.ru/", http.StatusMovedPermanently))

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}
