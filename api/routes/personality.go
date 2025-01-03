package routes

import (
	"github.com/LuanTenorio/learning_go/api/contreller"
	"github.com/gorilla/mux"
)

func personalityHandleRequest(r *mux.Router) {
	r.HandleFunc("/personality/{id}", contreller.FindOnePersonality)
	r.HandleFunc("/personalities", contreller.FindAllPersonalities)
}
