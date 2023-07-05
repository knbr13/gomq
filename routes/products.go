package routes

import "github.com/abdullah-alaadine/basic-rest-api/models"

type Product struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serialNumber"`
}

func CreateResponseProduct(product *models.Product) Product {
	return Product{
		ID:           product.ID,
		Name:         product.Name,
		SerialNumber: product.SerialNumber,
	}
}
