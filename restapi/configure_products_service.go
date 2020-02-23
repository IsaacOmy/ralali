// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/isaacomy/ralali/restapi/handlers"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/isaacomy/ralali/db"
	"github.com/isaacomy/ralali/restapi/operations"
	"github.com/isaacomy/ralali/restapi/operations/category"
	"github.com/isaacomy/ralali/restapi/operations/order"
	"github.com/isaacomy/ralali/restapi/operations/price"
	"github.com/isaacomy/ralali/restapi/operations/product"
)

//go:generate swagger generate server --target ../../ralali --name ProductsService --spec ../apispec/products.yaml

func configureFlags(api *operations.ProductsServiceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ProductsServiceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Init DB
	db.ConfigureDB()

	api.CategoryDeleteCategoryIDHandler = category.DeleteCategoryIDHandlerFunc(handlers.CategoryDeleteByIdHandler)

	api.PriceDeletePriceIDHandler = price.DeletePriceIDHandlerFunc(handlers.PriceDeleteByIDHandler)

	api.ProductDeleteProductsIDHandler = product.DeleteProductsIDHandlerFunc(handlers.ProductDeleteByIdHanlder)

	api.CategoryGetCategoryIDHandler = category.GetCategoryIDHandlerFunc(handlers.CategoryGetByIdHandler)

	api.CategoryGetCategoryListHandler = category.GetCategoryListHandlerFunc(handlers.CategoryGetListHanler)

	api.OrderGetOrderIDHandler = order.GetOrderIDHandlerFunc(handlers.OrderGetDetailHandler)

	api.ProductGetProductsIDHandler = product.GetProductsIDHandlerFunc(handlers.ProductGetByID)

	api.ProductGetProductsListHandler = product.GetProductsListHandlerFunc(handlers.ProductGetListHandler)

	api.CategoryPostCategoryHandler = category.PostCategoryHandlerFunc(handlers.CategoryPostHandler)

	api.OrderPostOrderHandler = order.PostOrderHandlerFunc(handlers.OrderPostHandler)

	api.PricePostPriceHandler = price.PostPriceHandlerFunc(handlers.PricePostHandler)

	api.ProductPostProductsHandler = product.PostProductsHandlerFunc(handlers.ProductPostHandler)

	api.CategoryPutCategoryIDHandler = category.PutCategoryIDHandlerFunc(handlers.CategoryPutByIdHandler)

	api.PricePutPriceIDHandler = price.PutPriceIDHandlerFunc(handlers.PricePutByIDHandler)

	api.ProductPutProductsIDHandler = product.PutProductsIDHandlerFunc(handlers.ProductPutHandler)

	api.ProductPostProductsCompareHandler = product.PostProductsCompareHandlerFunc(handlers.ProductPostCompareHandler)

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
