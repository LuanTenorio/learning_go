package main

import (
	"fmt"
	"net/http"

	"github.com/LuanTenorio/learning_go/api/database"
	"github.com/LuanTenorio/learning_go/api/routes"
)

func main() {
	database.DatabaseConnect()
	routes.CreateServer()

	fmt.Println("server is running...")
	http.ListenAndServe(":8000", nil)
}
