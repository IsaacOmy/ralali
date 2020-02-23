package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/isaacomy/ralali/db"
	"github.com/isaacomy/ralali/models"
	"github.com/isaacomy/ralali/restapi/operations/category"
)

func CategoryPutByIdHandler(params category.PutCategoryIDParams) middleware.Responder {

	db := db.GetDB()

	obj := models.Category{}

	if db.Where("status = ?", 1).First(&obj, params.ID).RecordNotFound() {
		log.Println("record not found")
		errMsg := models.ErrorMessage{Code: 100, Message: "record not found"}
		return category.NewPutCategoryIDNotFound().WithPayload(&errMsg)
	}

	if params.Body.Name != "" {
		obj.Name = params.Body.Name
	}

	if params.Body.CategoryParentID != 0 {
		obj.CategoryParentID = params.Body.CategoryParentID
	}

	err := db.Save(&obj)
	if len(err.GetErrors()) != 0 {
		log.Println("Error Update Categories")
		errMsg := models.ErrorMessage{Code: 100, Message: "Error Update Product"}
		return category.NewPutCategoryIDBadRequest().WithPayload(&errMsg)
	}

	log.Println("Put Category by ID OK")

	return category.NewPutCategoryIDOK().WithPayload(&obj)
}
