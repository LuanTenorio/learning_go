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

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var personality model.Personality
	json.NewDecoder(r.Body).Decode(&personality)

	err := model.CreatePersonality(&personality)

	if err != nil {
		logger.Println(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(personality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		logger.Println(err)
		return
	}

	err = model.DeletePersonality(id)

	if err != nil {
		logger.Println(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
}
