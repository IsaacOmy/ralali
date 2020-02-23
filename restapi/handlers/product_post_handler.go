package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/isaacomy/ralali/db"
	"github.com/isaacomy/ralali/models"
	"github.com/isaacomy/ralali/restapi/operations/product"
)

func ProductPostHandler(params product.PostProductsParams) middleware.Responder {

	db := db.GetDB()

	categories := []*models.Category{}

	for _, c := range params.Body.Categories {
		category := models.Category{}
		if db.Where("status = ?", 1).First(&category, c).RecordNotFound() {
			log.Println("category not found")
			errMsg := models.ErrorMessage{Code: 100, Message: "cannot insert to DB"}
			return product.NewPutProductsIDNotFound().WithPayload(&errMsg)
		}
		categories = append(categories, &category)
	}

	// res2B, _ := json.Marshal(categories)
	// fmt.Println(string(res2B))

	prices := []*models.Price{}

	for _, p := range params.Body.Prices {
		price := models.Price{}
		price.Price = p.Price
		price.MinQuantity = p.MinQuantity
		price.Status = 1

		prices = append(prices, &price)
	}

	obj := models.Product{}
	obj.Name = params.Body.Name
	obj.Series = params.Body.Series
	obj.Status = 1
	obj.Prices = prices

	tx := db.Begin()

	err := tx.Create(&obj)
	if len(err.GetErrors()) != 0 {
		log.Println(err.GetErrors())
		errMsg := models.ErrorMessage{Code: 100, Message: "cannot insert to DB"}
		return product.NewPostProductsBadRequest().WithPayload(&errMsg)
	}

	phc := models.ProductHasCategory{}
	for _, c := range categories {
		phc = models.ProductHasCategory{CategoryID: c.ID, ProductID: obj.ID}

		err := tx.Create(&phc)
		if len(err.GetErrors()) != 0 {
			tx.Rollback()
			log.Println(err.GetErrors())
			errMsg := models.ErrorMessage{Code: 100, Message: "cannot insert to DB"}
			return product.NewPostProductsBadRequest().WithPayload(&errMsg)
		}
	}

	tx.Commit()

	log.Println("Create new product OK")

	return product.NewPostProductsOK().WithPayload(&product.PostProductsOKBody{Message: "Product created OK"})
}
