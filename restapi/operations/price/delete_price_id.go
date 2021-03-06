// Code generated by go-swagger; DO NOT EDIT.

package price

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
)

// DeletePriceIDHandlerFunc turns a function with the right signature into a delete price ID handler
type DeletePriceIDHandlerFunc func(DeletePriceIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeletePriceIDHandlerFunc) Handle(params DeletePriceIDParams) middleware.Responder {
	return fn(params)
}

// DeletePriceIDHandler interface for that can handle valid delete price ID params
type DeletePriceIDHandler interface {
	Handle(DeletePriceIDParams) middleware.Responder
}

// NewDeletePriceID creates a new http.Handler for the delete price ID operation
func NewDeletePriceID(ctx *middleware.Context, handler DeletePriceIDHandler) *DeletePriceID {
	return &DeletePriceID{Context: ctx, Handler: handler}
}

/*DeletePriceID swagger:route DELETE /price/{id} price deletePriceId

Delete Price by ID

delete Price Data by ID

*/
type DeletePriceID struct {
	Context *middleware.Context
	Handler DeletePriceIDHandler
}

func (o *DeletePriceID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeletePriceIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeletePriceIDOKBody delete price ID o k body
// swagger:model DeletePriceIDOKBody
type DeletePriceIDOKBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delete price ID o k body
func (o *DeletePriceIDOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeletePriceIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeletePriceIDOKBody) UnmarshalBinary(b []byte) error {
	var res DeletePriceIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
