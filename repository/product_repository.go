package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRespository struct {
	connection *sql.DB
}

func NewProductRespository(connection *sql.DB) ProductRespository {
	return ProductRespository{
		connection: connection,
	}
}

func (pr *ProductRespository) GetProducts() ([]model.Product, error) {

	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil
}
