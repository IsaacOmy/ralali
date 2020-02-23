// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"
	"strconv"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
)

// PostProductsHandlerFunc turns a function with the right signature into a post products handler
type PostProductsHandlerFunc func(PostProductsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostProductsHandlerFunc) Handle(params PostProductsParams) middleware.Responder {
	return fn(params)
}

// PostProductsHandler interface for that can handle valid post products params
type PostProductsHandler interface {
	Handle(PostProductsParams) middleware.Responder
}

// NewPostProducts creates a new http.Handler for the post products operation
func NewPostProducts(ctx *middleware.Context, handler PostProductsHandler) *PostProducts {
	return &PostProducts{Context: ctx, Handler: handler}
}

/*PostProducts swagger:route POST /products product postProducts

Create New Product

Add New Product

*/
type PostProducts struct {
	Context *middleware.Context
	Handler PostProductsHandler
}

func (o *PostProducts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostProductsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostProductsBody post products body
// swagger:model PostProductsBody
type PostProductsBody struct {

	// categories id
	Categories []int64 `json:"categories"`

	// name
	Name string `json:"name,omitempty"`

	// prices
	Prices []*PricesItems0 `json:"prices"`

	// series
	Series string `json:"series,omitempty"`
}

// Validate validates this post products body
func (o *PostProductsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePrices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostProductsBody) validatePrices(formats strfmt.Registry) error {

	if swag.IsZero(o.Prices) { // not required
		return nil
	}

	for i := 0; i < len(o.Prices); i++ {
		if swag.IsZero(o.Prices[i]) { // not required
			continue
		}

		if o.Prices[i] != nil {
			if err := o.Prices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "prices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostProductsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostProductsBody) UnmarshalBinary(b []byte) error {
	var res PostProductsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostProductsOKBody post products o k body
// swagger:model PostProductsOKBody
type PostProductsOKBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this post products o k body
func (o *PostProductsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostProductsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostProductsOKBody) UnmarshalBinary(b []byte) error {
	var res PostProductsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PricesItems0 prices items0
// swagger:model PricesItems0
type PricesItems0 struct {

	// min quantity
	MinQuantity int64 `json:"min_quantity,omitempty"`

	// price
	Price float64 `json:"price,omitempty"`
}

// Validate validates this prices items0
func (o *PricesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PricesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PricesItems0) UnmarshalBinary(b []byte) error {
	var res PricesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
