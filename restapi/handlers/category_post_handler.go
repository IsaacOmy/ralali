package handlers

import (
	"log"
	"strconv"

	"github.com/isaacomy/ralali/db"

	"github.com/go-openapi/runtime/middleware"
	"github.com/isaacomy/ralali/models"
	"github.com/isaacomy/ralali/restapi/operations/category"
)

func CategoryPostHandler(params category.PostCategoryParams) middleware.Responder {

	if params.Body.Name == "" {
		log.Println("name cannot be empty")
		errMgs := models.ErrorMessage{Code: 100, Message: "name cannot be empty"}
		return category.NewPostCategoryBadRequest().WithPayload(&errMgs)
	}
	db := db.GetDB()
	var query string
	if params.Body.CategoryParentID == 0 {
		query = "INSERT INTO category (name, status) VALUES ('" + params.Body.Name + "', 1);"
	} else {
		categoryParentID := strconv.Itoa(int(params.Body.CategoryParentID))
		query = "INSERT INTO category (name, category_parent_id, status) VALUES ('" + params.Body.Name + "'," + categoryParentID + ", 1);"
	}

	err := db.Exec(query)
	if len(err.GetErrors()) != 0 {
		log.Println(err.GetErrors())
		
		errMsg := models.ErrorMessage{Code: 100, Message: "cannot insert to DB"}
		return category.NewPostCategoryBadRequest().WithPayload(&errMsg)
	}

	log.Println("create new category OK")

	return category.NewPostCategoryOK().WithPayload(&category.PostCategoryOKBody{Message: "Category created OK"})
}
