package handlers

import (
	"log"

	"github.com/go-openapi/runtime/middleware"
	"github.com/isaacomy/ralali/db"
	"github.com/isaacomy/ralali/models"
	"github.com/isaacomy/ralali/restapi/operations/product"
)

func ProductPutHandler(params product.PutProductsIDParams) middleware.Responder {

	db := db.GetDB()

	obj := models.Product{}

	if db.Where("status = ?", 1).First(&obj, params.ID).RecordNotFound() {
		log.Println("record not found")
		errMsg := models.ErrorMessage{Code: 100, Message: "record not found"}
		return product.NewGetProductsIDNotFound().WithPayload(&errMsg)
	}

	if params.Body.Name != "" {
		obj.Name = params.Body.Name
	}

	if params.Body.Series != "" {
		obj.Series = params.Body.Series
	}

	tx := db.Begin()

	if len(params.Body.Catogories) != 0 {
		if err := tx.Where("product_id = ?", params.ID).Delete(models.ProductHasCategory{}).Error; err != nil {
			log.Println("Error Update Categories")
			errMsg := models.ErrorMessage{Code: 100, Message: "Error Update Categories"}
			return product.NewPutProductsIDBadRequest().WithPayload(&errMsg)
		}

		phc := models.ProductHasCategory{}
		for _, c := range params.Body.Catogories {
			phc = models.ProductHasCategory{CategoryID: c, ProductID: obj.ID}

			err := tx.Create(&phc)
			if len(err.GetErrors()) != 0 {
				log.Println(err.GetErrors())
				errMsg := models.ErrorMessage{Code: 100, Message: "cannot insert to DB"}
				return product.NewPutProductsIDBadRequest().WithPayload(&errMsg)
			}
		}
	}

	err := tx.Debug().Model(&obj).Where("id = ?", obj.ID).Updates(&obj).Error
	if err != nil {
		log.Println("Error Update Categories")
		errMsg := models.ErrorMessage{Code: 100, Message: "Error Update Product"}
		return product.NewPutProductsIDBadRequest().WithPayload(&errMsg)
	}

	tx.Commit()

	db.Preload("Categories").Preload("Prices").First(&obj, params.ID)

	log.Println("Put Product by ID OK")

	return product.NewPutProductsIDOK().WithPayload(&obj)
}
