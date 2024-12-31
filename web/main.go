package main

import (
	"fmt"
	"net/http"

	"github.com/LuanTenorio/learning_go/web/config"
	"github.com/LuanTenorio/learning_go/web/internal/route"
	_ "github.com/lib/pq"
)

func main() {
	config.LoadConfig()
	route.LoadRoutes()

	fmt.Println("server is running...")
	http.ListenAndServe(":8000", nil)
}
