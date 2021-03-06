// Code generated by go-swagger; DO NOT EDIT.

package price

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/isaacomy/ralali/models"
)

// PostPriceOKCode is the HTTP code returned for type PostPriceOK
const PostPriceOKCode int = 200

/*PostPriceOK Insert New Price

swagger:response postPriceOK
*/
type PostPriceOK struct {

	/*
	  In: Body
	*/
	Payload *PostPriceOKBody `json:"body,omitempty"`
}

// NewPostPriceOK creates PostPriceOK with default headers values
func NewPostPriceOK() *PostPriceOK {

	return &PostPriceOK{}
}

// WithPayload adds the payload to the post price o k response
func (o *PostPriceOK) WithPayload(payload *PostPriceOKBody) *PostPriceOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post price o k response
func (o *PostPriceOK) SetPayload(payload *PostPriceOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostPriceOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostPriceBadRequestCode is the HTTP code returned for type PostPriceBadRequest
const PostPriceBadRequestCode int = 400

/*PostPriceBadRequest Bad Request / Validation exception

swagger:response postPriceBadRequest
*/
type PostPriceBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewPostPriceBadRequest creates PostPriceBadRequest with default headers values
func NewPostPriceBadRequest() *PostPriceBadRequest {

	return &PostPriceBadRequest{}
}

// WithPayload adds the payload to the post price bad request response
func (o *PostPriceBadRequest) WithPayload(payload *models.ErrorMessage) *PostPriceBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post price bad request response
func (o *PostPriceBadRequest) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostPriceBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostPriceDefault Error

swagger:response postPriceDefault
*/
type PostPriceDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewPostPriceDefault creates PostPriceDefault with default headers values
func NewPostPriceDefault(code int) *PostPriceDefault {
	if code <= 0 {
		code = 500
	}

	return &PostPriceDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post price default response
func (o *PostPriceDefault) WithStatusCode(code int) *PostPriceDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post price default response
func (o *PostPriceDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post price default response
func (o *PostPriceDefault) WithPayload(payload *models.ErrorMessage) *PostPriceDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post price default response
func (o *PostPriceDefault) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostPriceDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
