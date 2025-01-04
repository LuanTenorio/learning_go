package routes

import (
	"github.com/LuanTenorio/learning_go/api/contreller"
	"github.com/gorilla/mux"
)

func personalityHandleRequest(r *mux.Router) {
	r.HandleFunc("/personalities/{id}", contreller.FindOnePersonality).Methods("Get")
	r.HandleFunc("/personalities", contreller.FindAllPersonalities).Methods("Get")
	r.HandleFunc("/personalities", contreller.CreatePersonality).Methods("Post")
}
