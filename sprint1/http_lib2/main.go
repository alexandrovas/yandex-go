package main

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

var cars = map[string]string{
	"id1": "Renault Logan",
	"id2": "Renault Duster",
	"id3": "BMW X6",
	"id4": "BMW M5",
	"id5": "VW Passat",
	"id6": "VW Jetta",
	"id7": "Audi A4",
	"id8": "Audi Q7",
}

// carsListFunc — вспомогательная функция для вывода всех машин.
func carsListFunc() []string {
	var list []string
	for _, c := range cars {
		list = append(list, c)
	}
	return list
}

// carFunc — вспомогательная функция для вывода определённой машины.
func carFunc(id string) string {
	if c, ok := cars[id]; ok {
		return c
	}
	return "unknown identifier " + id
}

func carsByBrand(brand string) []string {
	var list []string
	for _, c := range cars {
		if strings.EqualFold(strings.Split(c, " ")[0], brand) {
			list = append(list, c)
		}
	}
	return list
}

func carsByBrandModel(brand, model string) []string {
	var list []string
	for _, c := range cars {
		if strings.EqualFold(c, brand+" "+model) {
			list = append(list, c)
		}
	}
	return list
}

func carsHandle(rw http.ResponseWriter, r *http.Request) {
	carsList := carsListFunc()
	io.WriteString(rw, strings.Join(carsList, ", "))
}

func carHandle(rw http.ResponseWriter, r *http.Request) {
	carID := chi.URLParam(r, "id")
	rw.Write([]byte(carFunc(carID)))
}

func carsBrandHandle(rw http.ResponseWriter, r *http.Request) {
	brand := chi.URLParam(r, "brand")
	carsList := carsByBrand(brand)
	io.WriteString(rw, strings.Join(carsList, ", "))
}

func carsBrandModelHandle(rw http.ResponseWriter, r *http.Request) {
	brand := chi.URLParam(r, "brand")
	model := chi.URLParam(r, "model")
	carsList := carsByBrandModel(brand, model)
	io.WriteString(rw, strings.Join(carsList, ", "))
}

func main() {
	r := chi.NewRouter()

	// хендлер, который выводит определённую машину
	r.Get("/car/{id}", carHandle)

	r.Route("/cars", func(r chi.Router) {
		// хендлер, который выводит все машины
		r.Get("/", carsHandle)
		// хендлер, который выводит все модели определенного бренда
		r.Get("/{brand}", carsBrandHandle)
		// хендлер, который выводит машины определенного бренда и модели
		r.Get("/{brand}/{model}", carsBrandModelHandle)

	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
