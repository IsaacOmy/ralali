package handlers

import (
	"log"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/isaacomy/ralali/db"
	"github.com/isaacomy/ralali/models"

	"github.com/isaacomy/ralali/restapi/operations/order"
	"github.com/isaacomy/ralali/restapi/operations/product"
)

func OrderPostHandler(params order.PostOrderParams) middleware.Responder {

	db := db.GetDB()
	tx := db.Begin()
	t := time.Now()
	strTime := t.Format("2006-01-02T15:04:05")
	orders := models.Order{}
	orders.CreatedDate = strTime
	tx.Create(&orders)

	itemOrders := []*models.ItemOrder{}

	for _, p := range params.Body {
		itemOrder := models.ItemOrder{}
		prod := models.Product{}
		price := models.Price{}
		log.Println(p.ProductID)
		if tx.Where("status = ?", 1).First(&prod, p.ProductID).RecordNotFound() {
			tx.Rollback()
			log.Println("product not found")
			errMsg := models.ErrorMessage{Code: 100, Message: "product not found"}
			return product.NewPostProductsCompareBadRequest().WithPayload(&errMsg)
		}
		tx.Preload("Categories", "status = ?", 1).Preload("Prices", "status = ?", 1).First(&prod)
		tx.Where("status = ? AND product_id = ? AND (NOT min_quantity > ?)", 1, prod.ID, p.Quantity).
			Order("min_quantity DESC").First(&price)
		itemOrder.OrderID = orders.ID
		itemOrder.Price = price.Price
		itemOrder.ProductID = prod.ID
		itemOrder.Quantity = p.Quantity
		itemOrders = append(itemOrders, &itemOrder)
	}

	for _, itemOrder := range itemOrders {
		err := tx.Create(&itemOrder)
		if len(err.GetErrors()) != 0 {
			tx.Rollback()
			log.Println(err.GetErrors())
			errMsg := models.ErrorMessage{Code: 100, Message: "cannot insert to DB"}
			return order.NewPostOrderBadRequest().WithPayload(&errMsg)
		}
	}

	tx.Commit()

	log.Println("Create new Order OK")

	return order.NewPostOrderOK().WithPayload(&order.PostOrderOKBody{Message: "Create new Order OK", OrderID: orders.ID})
}
