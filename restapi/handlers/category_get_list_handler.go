package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/isaacomy/ralali/db"
	"github.com/isaacomy/ralali/models"
	"github.com/isaacomy/ralali/restapi/operations/category"
)

func CategoryGetListHanler(params category.GetCategoryListParams) middleware.Responder {

	categories := []*models.Category{}

	db := db.GetDB()

	if params.Name != nil {
		db = db.Where("name like ?", "%"+*params.Name+"%")
	}
	db = db.Where("status = ?", 1)

	db = db.Preload("CategoryParent").Find(&categories)

	var count float64
	db.Find(&categories).Count(&count)

	if params.PerPage == nil {
		log.Println("PerPage cannot be nil")
		errMsg := models.ErrorMessage{Code: 100, Message: "PerPage cannot be nil"}
		return category.NewGetCategoryListBadRequest().WithPayload(&errMsg)
	}
	length := *params.PerPage

	if params.Page == nil {
		log.Println("Page cannot be nil")
		errMsg := models.ErrorMessage{Code: 100, Message: "Page cannot be nil"}
		return category.NewGetCategoryListBadRequest().WithPayload(&errMsg)
	}
	start := ((*params.Page - 1) * length)

	if err := db.Limit(length).Offset(start).Find(&categories).Error; err != nil {
		log.Println(err)
		errMsg := models.ErrorMessage{Code: 100, Message: "cannot query category list"}
		return category.NewGetCategoryListNotFound().WithPayload(&errMsg)
	}

	log.Println("Get Category List OK")
	result := &category.GetCategoryListOKBody{TotalItems: count, Items: categories}

	return category.NewGetCategoryListOK().WithPayload(result)
}
