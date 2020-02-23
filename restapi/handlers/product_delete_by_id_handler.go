package handlers

import (
	"log"

	"github.com/isaacomy/ralali/db"

	"github.com/go-openapi/runtime/middleware"
	"github.com/isaacomy/ralali/models"
	"github.com/isaacomy/ralali/restapi/operations/product"
)

func ProductDeleteByIdHanlder(params product.DeleteProductsIDParams) middleware.Responder {

	obj := models.Product{}

	db := db.GetDB()

	if db.Where("status = ?", 1).First(&obj, params.ID).RecordNotFound() {
		log.Println("record not found")
		errMsg := models.ErrorMessage{Code: 100, Message: "record not found"}
		return product.NewDeleteProductsIDBadRequest().WithPayload(&errMsg)
	}

	obj.Status = 0

	err := db.Save(&obj)
	if err.GetErrors != nil {
		log.Println("Error Delete Categories")
		errMsg := models.ErrorMessage{Code: 100, Message: "Error Delete Product"}
		return product.NewDeleteProductsIDBadRequest().WithPayload(&errMsg)
	}

	log.Println("Delete product OK")

	return product.NewDeleteProductsIDOK().WithPayload(&product.DeleteProductsIDOKBody{Message: "Product deleted OK"})
}
