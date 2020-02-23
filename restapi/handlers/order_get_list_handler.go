package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/isaacomy/ralali/db"
	"github.com/isaacomy/ralali/models"
	"github.com/isaacomy/ralali/restapi/operations/order"
)

func OrderGetDetailHandler(params order.GetOrderIDParams) middleware.Responder {

	db := db.GetDB()
	obj := models.Order{}

	if db.First(&obj, params.ID).RecordNotFound() {
		log.Println("record not found")
		errMsg := models.ErrorMessage{Code: 100, Message: "record not found"}
		return order.NewGetOrderIDBadRequest().WithPayload(&errMsg)
	}

	db.Preload("ItemOrders").Preload("ItemOrders.Product").Find(&obj)

	log.Println("Get Order by ID OK")

	return order.NewGetOrderIDOK().WithPayload(&obj)
}
