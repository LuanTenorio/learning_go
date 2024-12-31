package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type PostgresConfig struct {
	Url string
}

type Env struct {
	Pg PostgresConfig
}

var AppConfig Env

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env: ", err)
		os.Exit(1)
	}

	pgConfig := PostgresConfig{os.Getenv("PG_URL")}

	AppConfig = Env{Pg: pgConfig}
}
