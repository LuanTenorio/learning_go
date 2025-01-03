package routes

import (
	"net/http"

	"github.com/LuanTenorio/learning_go/api/contreller"
	"github.com/gorilla/mux"
)

func CreateServer() {
	r := mux.NewRouter()
	handleRequest(r)
	http.Handle("/", r)
}

func handleRequest(r *mux.Router) {
	r.HandleFunc("/", contreller.Home)
	r.HandleFunc("/personalities", contreller.FindAllPersonalities)
}
