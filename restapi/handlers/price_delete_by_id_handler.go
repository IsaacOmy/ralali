package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/isaacomy/ralali/db"
	"github.com/isaacomy/ralali/models"
	"github.com/isaacomy/ralali/restapi/operations/price"
)

func PriceDeleteByIDHandler(params price.DeletePriceIDParams) middleware.Responder {
	db := db.GetDB()

	obj := models.Price{}

	if db.Where("status = ?", 1).First(&obj, params.ID).RecordNotFound() {
		log.Println("record not found")
		errMsg := models.ErrorMessage{Code: 100, Message: "record not found"}
		return price.NewDeletePriceIDNotFound().WithPayload(&errMsg)
	}

	obj.Status = 0

	err := db.Save(&obj)
	if err.GetErrors != nil {
		log.Println("Error Delete Price")
		errMsg := models.ErrorMessage{Code: 100, Message: "Error Delete Product"}
		return price.NewDeletePriceIDBadRequest().WithPayload(&errMsg)
	}

	log.Println("Delete Price by ID OK")

	return price.NewDeletePriceIDOK().WithPayload(&price.DeletePriceIDOKBody{Message: "Price deleted OK"})
}
