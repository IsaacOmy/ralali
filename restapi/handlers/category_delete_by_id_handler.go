package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/isaacomy/ralali/db"
	"github.com/isaacomy/ralali/models"
	"github.com/isaacomy/ralali/restapi/operations/category"
)

func CategoryDeleteByIdHandler(params category.DeleteCategoryIDParams) middleware.Responder {
	db := db.GetDB()

	obj := models.Category{}

	if db.Where("status = ?", 1).First(&obj, params.ID).RecordNotFound() {
		log.Println("record not found")
		errMsg := models.ErrorMessage{Code: 100, Message: "record not found"}
		return category.NewDeleteCategoryIDNotFound().WithPayload(&errMsg)
	}

	err := db.Save(&obj)
	if err.GetErrors != nil {
		log.Println("Error Delete Categories")
		errMsg := models.ErrorMessage{Code: 100, Message: "Error Delete Product"}
		return category.NewDeleteCategoryIDBadRequest().WithPayload(&errMsg)
	}

	log.Println("Delete Category by ID OK")

	return category.NewDeleteCategoryIDOK().WithPayload(&category.DeleteCategoryIDOKBody{Message: "Category deleted OK"})
}
