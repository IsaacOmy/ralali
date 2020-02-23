package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/isaacomy/ralali/db"
	"github.com/isaacomy/ralali/models"
	"github.com/isaacomy/ralali/restapi/operations/price"
)

func PricePostHandler(params price.PostPriceParams) middleware.Responder {

	db := db.GetDB()

	obj := models.Price{}

	if db.Where("status = ?", 1).First(&models.Product{}, params.Body.ProductID).RecordNotFound() {
		log.Println("record not found")
		errMsg := models.ErrorMessage{Code: 100, Message: "record not found"}
		return price.NewPostPriceBadRequest().WithPayload(&errMsg)
	}

	obj.Price = params.Body.Price
	obj.ProductID = params.Body.ProductID
	obj.MinQuantity = params.Body.MinQuantity
	obj.Status = 1

	tx := db.Begin()

	if err := tx.Create(&obj).Error; err != nil {
		tx.Rollback()
		log.Println("record not found")
		errMsg := models.ErrorMessage{Code: 100, Message: "cannot insert to DB"}
		return price.NewPostPriceBadRequest().WithPayload(&errMsg)
	}

	log.Println("create new category OK")

	return price.NewPostPriceOK().WithPayload(&price.PostPriceOKBody{Message: "Price created OK"})
}
