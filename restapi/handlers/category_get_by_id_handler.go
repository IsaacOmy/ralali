package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/isaacomy/ralali/db"
	"github.com/isaacomy/ralali/models"
	"github.com/isaacomy/ralali/restapi/operations/category"
)

func CategoryGetByIdHandler(params category.GetCategoryIDParams) middleware.Responder {
	obj := models.Category{}
	db := db.GetDB()

	if db.Where(&models.Category{ID: params.ID, Status: 1}).First(&obj).RecordNotFound() {
		log.Println("record not found")
		errMsg := models.ErrorMessage{Code: 100, Message: "record not found"}
		return category.NewGetCategoryIDNotFound().WithPayload(&errMsg)
	}

	db.Preload("CategoryParent").Find(&obj)

	log.Println("Get Category by ID OK")

	return category.NewGetCategoryIDOK().WithPayload(&obj)
}
