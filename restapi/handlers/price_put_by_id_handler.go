package handlers

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/isaacomy/ralali/db"
	"github.com/isaacomy/ralali/models"
	"github.com/isaacomy/ralali/restapi/operations/price"
)

func PricePutByIDHandler(params price.PutPriceIDParams) middleware.Responder {

	db := db.GetDB()

	obj := models.Price{}

	if db.Where("status = ?", 1).First(&obj, params.ID).RecordNotFound() {
		log.Println("record not found")
		errMsg := models.ErrorMessage{Code: 100, Message: "record not found"}
		return price.NewPutPriceIDBadRequest().WithPayload(&errMsg)
	}

	if params.Body.Price != 0 {
		obj.Price = params.Body.Price
	}

	if params.Body.MinQuantity != 0 {
		obj.MinQuantity = params.Body.MinQuantity
	}
	res2B, _ := json.Marshal(obj)
	fmt.Println(string(res2B))
	err := db.Save(&obj)
	if len(err.GetErrors()) != 0 {
		log.Println("Error Update Price")
		errMsg := models.ErrorMessage{Code: 100, Message: "Error Update Product"}
		return price.NewPutPriceIDBadRequest().WithPayload(&errMsg)
	}

	log.Println("Put Price by ID OK")

	return price.NewPutPriceIDOK().WithPayload(&price.PutPriceIDOKBody{Message: "Price Updated OK"})
}
