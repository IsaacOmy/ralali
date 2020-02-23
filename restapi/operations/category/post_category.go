// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
)

// PostCategoryHandlerFunc turns a function with the right signature into a post category handler
type PostCategoryHandlerFunc func(PostCategoryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostCategoryHandlerFunc) Handle(params PostCategoryParams) middleware.Responder {
	return fn(params)
}

// PostCategoryHandler interface for that can handle valid post category params
type PostCategoryHandler interface {
	Handle(PostCategoryParams) middleware.Responder
}

// NewPostCategory creates a new http.Handler for the post category operation
func NewPostCategory(ctx *middleware.Context, handler PostCategoryHandler) *PostCategory {
	return &PostCategory{Context: ctx, Handler: handler}
}

/*PostCategory swagger:route POST /category category postCategory

Create New Category

Add New Category

*/
type PostCategory struct {
	Context *middleware.Context
	Handler PostCategoryHandler
}

func (o *PostCategory) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostCategoryParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostCategoryBody post category body
// swagger:model PostCategoryBody
type PostCategoryBody struct {

	// category parent id
	CategoryParentID int64 `json:"category_parent_id,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this post category body
func (o *PostCategoryBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCategoryBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCategoryBody) UnmarshalBinary(b []byte) error {
	var res PostCategoryBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostCategoryOKBody post category o k body
// swagger:model PostCategoryOKBody
type PostCategoryOKBody struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this post category o k body
func (o *PostCategoryOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostCategoryOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostCategoryOKBody) UnmarshalBinary(b []byte) error {
	var res PostCategoryOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}