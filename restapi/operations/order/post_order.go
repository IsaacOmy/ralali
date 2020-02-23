// Code generated by go-swagger; DO NOT EDIT.

package order

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
)

// PostOrderHandlerFunc turns a function with the right signature into a post order handler
type PostOrderHandlerFunc func(PostOrderParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostOrderHandlerFunc) Handle(params PostOrderParams) middleware.Responder {
	return fn(params)
}

// PostOrderHandler interface for that can handle valid post order params
type PostOrderHandler interface {
	Handle(PostOrderParams) middleware.Responder
}

// NewPostOrder creates a new http.Handler for the post order operation
func NewPostOrder(ctx *middleware.Context, handler PostOrderHandler) *PostOrder {
	return &PostOrder{Context: ctx, Handler: handler}
}

/*PostOrder swagger:route POST /order order postOrder

Create New Order

Add New Order

*/
type PostOrder struct {
	Context *middleware.Context
	Handler PostOrderHandler
}

func (o *PostOrder) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostOrderParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostOrderOKBody post order o k body
// swagger:model PostOrderOKBody
type PostOrderOKBody struct {

	// message
	Message string `json:"message,omitempty"`

	// order id
	OrderID int64 `json:"order_id,omitempty"`
}

// Validate validates this post order o k body
func (o *PostOrderOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostOrderOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOrderOKBody) UnmarshalBinary(b []byte) error {
	var res PostOrderOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostOrderParamsBodyItems0 post order params body items0
// swagger:model PostOrderParamsBodyItems0
type PostOrderParamsBodyItems0 struct {

	// product id
	ProductID int64 `json:"product_id,omitempty"`

	// quantity
	Quantity int64 `json:"quantity,omitempty"`
}

// Validate validates this post order params body items0
func (o *PostOrderParamsBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostOrderParamsBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostOrderParamsBodyItems0) UnmarshalBinary(b []byte) error {
	var res PostOrderParamsBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}