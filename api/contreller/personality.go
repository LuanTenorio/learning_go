package contreller

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/LuanTenorio/learning_go/api/model"
	"github.com/gorilla/mux"
)

var logger = log.New(os.Stdout, "PersonalityController: ", 2)

func FindAllPersonalities(w http.ResponseWriter, r *http.Request) {
	personalities := model.FindPersonalities()

	json.NewEncoder(w).Encode(personalities)
}

func FindOnePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		logger.Println(err)
	}

	personality, err := model.FindPersonalityById(id)

	if err != nil {
		logger.Println(err.Error())
		w.WriteHeader(404)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(personality)
}
