package main

import (
	"html/template"
	"net/http"
)

type Product struct {
	Name        string
	Description string
	Price       float64
	Amount      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"T-shirt", "Very pretty", 29, 10},
		{" Notebook", "Very fast", 1999, 5},
	}

	temp.ExecuteTemplate(w, "Index", products)
}
