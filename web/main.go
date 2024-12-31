package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type postgresConfig struct {
	url string
}

type Product struct {
	Name        string
	Description string
	Price       float64
	Amount      int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	loadEnv()

	postgresUrl := os.Getenv("PG_URL")
	db := connectDB(postgresUrl)
	defer db.Close()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"T-shirt", "Very pretty", 29, 10},
		{" Notebook", "Very fast", 1999, 5},
	}

	temp.ExecuteTemplate(w, "Index", products)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env: ", err)
		os.Exit(1)
	}
}

func connectDB(url string) *sql.DB {
	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err.Error())
	}

	return db
}
