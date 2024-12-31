package model

import (
	"github.com/LuanTenorio/learning_go/web/config"
	"github.com/LuanTenorio/learning_go/web/internal/db"
)

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Amount            int
}

func GetProducts() []Product {
	db := db.ConnectDB(config.AppConfig.Pg.Url)

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
