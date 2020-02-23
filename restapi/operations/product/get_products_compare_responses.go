// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/isaacomy/ralali/models"
)

// GetProductsCompareOKCode is the HTTP code returned for type GetProductsCompareOK
const GetProductsCompareOKCode int = 200

/*GetProductsCompareOK get product list

swagger:response getProductsCompareOK
*/
type GetProductsCompareOK struct {

	/*
	  In: Body
	*/
	Payload *GetProductsCompareOKBody `json:"body,omitempty"`
}

// NewGetProductsCompareOK creates GetProductsCompareOK with default headers values
func NewGetProductsCompareOK() *GetProductsCompareOK {

	return &GetProductsCompareOK{}
}

// WithPayload adds the payload to the get products compare o k response
func (o *GetProductsCompareOK) WithPayload(payload *GetProductsCompareOKBody) *GetProductsCompareOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get products compare o k response
func (o *GetProductsCompareOK) SetPayload(payload *GetProductsCompareOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProductsCompareOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProductsCompareBadRequestCode is the HTTP code returned for type GetProductsCompareBadRequest
const GetProductsCompareBadRequestCode int = 400

/*GetProductsCompareBadRequest Bad Request / Validation exception

swagger:response getProductsCompareBadRequest
*/
type GetProductsCompareBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetProductsCompareBadRequest creates GetProductsCompareBadRequest with default headers values
func NewGetProductsCompareBadRequest() *GetProductsCompareBadRequest {

	return &GetProductsCompareBadRequest{}
}

// WithPayload adds the payload to the get products compare bad request response
func (o *GetProductsCompareBadRequest) WithPayload(payload *models.ErrorMessage) *GetProductsCompareBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get products compare bad request response
func (o *GetProductsCompareBadRequest) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProductsCompareBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetProductsCompareNotFoundCode is the HTTP code returned for type GetProductsCompareNotFound
const GetProductsCompareNotFoundCode int = 404

/*GetProductsCompareNotFound Not found

swagger:response getProductsCompareNotFound
*/
type GetProductsCompareNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetProductsCompareNotFound creates GetProductsCompareNotFound with default headers values
func NewGetProductsCompareNotFound() *GetProductsCompareNotFound {

	return &GetProductsCompareNotFound{}
}

// WithPayload adds the payload to the get products compare not found response
func (o *GetProductsCompareNotFound) WithPayload(payload *models.ErrorMessage) *GetProductsCompareNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get products compare not found response
func (o *GetProductsCompareNotFound) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProductsCompareNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetProductsCompareDefault Error

swagger:response getProductsCompareDefault
*/
type GetProductsCompareDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetProductsCompareDefault creates GetProductsCompareDefault with default headers values
func NewGetProductsCompareDefault(code int) *GetProductsCompareDefault {
	if code <= 0 {
		code = 500
	}

	return &GetProductsCompareDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get products compare default response
func (o *GetProductsCompareDefault) WithStatusCode(code int) *GetProductsCompareDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get products compare default response
func (o *GetProductsCompareDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get products compare default response
func (o *GetProductsCompareDefault) WithPayload(payload *models.ErrorMessage) *GetProductsCompareDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get products compare default response
func (o *GetProductsCompareDefault) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetProductsCompareDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
