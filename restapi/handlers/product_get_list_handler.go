package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/isaacomy/ralali/db"
	"github.com/isaacomy/ralali/models"
	"github.com/isaacomy/ralali/restapi/operations/product"
)

func ProductGetListHandler(params product.GetProductsListParams) middleware.Responder {

	products := []*models.Product{}

	keyword := params.Keyword
	db := db.GetDB()

	log.Println(keyword)

	if keyword != nil {
		db = db.Joins("left join product_has_category on product_has_category.product_id = product.id")
		db = db.Joins("left join category on product_has_category.category_id = category.id")
		db = db.Where("product.name rlike '(^|[[:space:]])" + *keyword + "([[:space:]]|$)'").
			Or("series rlike '(^|[[:space:]])" + *keyword + "([[:space:]]|$)'").
			Or("category.name rlike '(^|[[:space:]])" + *keyword + "([[:space:]]|$)'")
		db = db.Group("product.id")
	}

	db = db.Preload("Categories")
	db = db.Preload("Prices")

	var count float64
	db.Find(&products).Count(&count)

	if params.PerPage == nil {
		log.Println("PerPage cannot be nil")
		errMsg := models.ErrorMessage{Code: 100, Message: "PerPage cannot be nil"}
		return product.NewGetProductsListNotFound().WithPayload(&errMsg)
	}
	length := *params.PerPage

	if params.Page == nil {
		log.Println("Page cannot be nil")
		errMsg := models.ErrorMessage{Code: 100, Message: "Page cannot be nil"}
		return product.NewGetProductsListNotFound().WithPayload(&errMsg)
	}
	start := ((*params.Page - 1) * length)

	if err := db.Limit(length).Offset(start).Find(&products).Error; err != nil {
		log.Println(err)
		errMsg := models.ErrorMessage{Code: 100, Message: "cannot query category list"}
		return product.NewGetProductsListNotFound().WithPayload(&errMsg)
	}
	log.Println("Get Product List OK")

	result := &product.GetProductsListOKBody{TotalItems: count, Items: products}

	return product.NewGetProductsListOK().WithPayload(result)
}
