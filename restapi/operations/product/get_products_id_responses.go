// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/isaacomy/ralali/models"
)

// GetProductsIDOKCode is the HTTP code returned for type GetProductsIDOK
const GetProductsIDOKCode int = 200

/*GetProductsIDOK Sends Product data

swagger:response getProductsIdOK
*/
type GetProductsIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Product `json:"body,omitempty"`
}

// NewGetProductsIDOK creates GetProductsIDOK with default headers values
func NewGetProductsIDOK() *GetProductsIDOK {

	return &GetProductsIDOK{}
}

// WithPayload adds the payload to the get products Id o k response
func (o *GetProductsIDOK) WithPayload(payload *models.Product) *GetProductsIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get products Id o k response
func (o *GetProductsIDOK) SetPayload(payload *models.Product) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProductsIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProductsIDBadRequestCode is the HTTP code returned for type GetProductsIDBadRequest
const GetProductsIDBadRequestCode int = 400

/*GetProductsIDBadRequest Bad Request / Validation exception

swagger:response getProductsIdBadRequest
*/
type GetProductsIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetProductsIDBadRequest creates GetProductsIDBadRequest with default headers values
func NewGetProductsIDBadRequest() *GetProductsIDBadRequest {

	return &GetProductsIDBadRequest{}
}

// WithPayload adds the payload to the get products Id bad request response
func (o *GetProductsIDBadRequest) WithPayload(payload *models.ErrorMessage) *GetProductsIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get products Id bad request response
func (o *GetProductsIDBadRequest) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProductsIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProductsIDNotFoundCode is the HTTP code returned for type GetProductsIDNotFound
const GetProductsIDNotFoundCode int = 404

/*GetProductsIDNotFound Not found

swagger:response getProductsIdNotFound
*/
type GetProductsIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetProductsIDNotFound creates GetProductsIDNotFound with default headers values
func NewGetProductsIDNotFound() *GetProductsIDNotFound {

	return &GetProductsIDNotFound{}
}

// WithPayload adds the payload to the get products Id not found response
func (o *GetProductsIDNotFound) WithPayload(payload *models.ErrorMessage) *GetProductsIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get products Id not found response
func (o *GetProductsIDNotFound) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProductsIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProductsIDUnprocessableEntityCode is the HTTP code returned for type GetProductsIDUnprocessableEntity
const GetProductsIDUnprocessableEntityCode int = 422

/*GetProductsIDUnprocessableEntity Unprocessable Entity

swagger:response getProductsIdUnprocessableEntity
*/
type GetProductsIDUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetProductsIDUnprocessableEntity creates GetProductsIDUnprocessableEntity with default headers values
func NewGetProductsIDUnprocessableEntity() *GetProductsIDUnprocessableEntity {

	return &GetProductsIDUnprocessableEntity{}
}

// WithPayload adds the payload to the get products Id unprocessable entity response
func (o *GetProductsIDUnprocessableEntity) WithPayload(payload *models.ErrorMessage) *GetProductsIDUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get products Id unprocessable entity response
func (o *GetProductsIDUnprocessableEntity) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProductsIDUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetProductsIDDefault Error

swagger:response getProductsIdDefault
*/
type GetProductsIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetProductsIDDefault creates GetProductsIDDefault with default headers values
func NewGetProductsIDDefault(code int) *GetProductsIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetProductsIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get products ID default response
func (o *GetProductsIDDefault) WithStatusCode(code int) *GetProductsIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get products ID default response
func (o *GetProductsIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get products ID default response
func (o *GetProductsIDDefault) WithPayload(payload *models.ErrorMessage) *GetProductsIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get products ID default response
func (o *GetProductsIDDefault) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProductsIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
