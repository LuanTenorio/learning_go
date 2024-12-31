package controller

import (
	"html/template"
	"net/http"

	"github.com/LuanTenorio/learning_go/web/internal/model"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := model.GetProducts()

	temp.ExecuteTemplate(w, "Index", products)
}
