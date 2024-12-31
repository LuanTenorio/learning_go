package route

import (
	"net/http"

	"github.com/LuanTenorio/learning_go/web/internal/controller"
)

func LoadRoutes() {
	http.HandleFunc("/", controller.Index)
}
