package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/isaacomy/ralali/db"
	"github.com/isaacomy/ralali/models"

	"github.com/isaacomy/ralali/restapi/operations/product"
)

func ProductPostCompareHandler(params product.PostProductsCompareParams) middleware.Responder {

	db := db.GetDB()

	products := []*models.Product{}

	for _, p := range params.Body.ProductsID {
		prod := models.Product{}
		if db.Where("status = ?", 1).First(&prod, p).RecordNotFound() {
			log.Println("product not found")
			errMsg := models.ErrorMessage{Code: 100, Message: "product not found"}
			return product.NewPostProductsCompareBadRequest().WithPayload(&errMsg)
		}
		db.Preload("Categories", "status = ?", 1).Preload("Prices", "status = ?", 1).First(&prod)
		products = append(products, &prod)
	}

	totalItem := int64(len(products))

	log.Println("List Compare Product OK")

	return product.NewPostProductsCompareOK().WithPayload(&product.PostProductsCompareOKBody{Items: products, TotalItems: totalItem})
}
