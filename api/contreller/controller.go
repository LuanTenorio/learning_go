package contreller

import (
	"encoding/json"
	"net/http"

	"github.com/LuanTenorio/learning_go/api/model"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Home"))
}

func FindAllPersonalities(w http.ResponseWriter, r *http.Request) {
	personalities := model.FindPersonalities()

	json.NewEncoder(w).Encode(personalities)
}
