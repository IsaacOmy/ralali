package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/isaacomy/ralali/db"
	"github.com/isaacomy/ralali/models"
	"github.com/isaacomy/ralali/restapi/operations/product"
)

func ProductGetByID(params product.GetProductsIDParams) middleware.Responder {

	db := db.GetDB()

	obj := models.Product{}

	if db.Where("status = ?", 1).First(&obj, params.ID).RecordNotFound() {
		log.Println("record not found")
		errMsg := models.ErrorMessage{Code: 100, Message: "record not found"}
		return product.NewGetProductsIDNotFound().WithPayload(&errMsg)
	}

	db.Preload("Categories", "status = ?", 1).Preload("Prices", "status = ?", 1).First(&obj)

	log.Println("Get Category by ID OK")

	return product.NewGetProductsIDOK().WithPayload(&obj)
}
