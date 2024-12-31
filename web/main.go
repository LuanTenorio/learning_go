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
	Id                int
	Name, Description string
	Price             float64
	Amount            int
}

var postgresconfig postgresConfig
var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	loadEnv()

	postgresconfig = postgresConfig{os.Getenv("PG_URL")}

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := getProducts()

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

func getProducts() []Product {
	db := connectDB(postgresconfig.url)

	selectAllProducts, err := db.Query("select * from products")
	if err != nil {
		panic(err.Error())
	}

	product := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, amount int
		var name, description string
		var price float64

		err := selectAllProducts.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount

		products = append(products, product)
	}

	defer db.Close()
	return products
}
