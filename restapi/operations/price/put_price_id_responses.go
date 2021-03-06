// Code generated by go-swagger; DO NOT EDIT.

package price

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/isaacomy/ralali/models"
)

// PutPriceIDOKCode is the HTTP code returned for type PutPriceIDOK
const PutPriceIDOKCode int = 200

/*PutPriceIDOK Updates Price data

swagger:response putPriceIdOK
*/
type PutPriceIDOK struct {

	/*
	  In: Body
	*/
	Payload *PutPriceIDOKBody `json:"body,omitempty"`
}

// NewPutPriceIDOK creates PutPriceIDOK with default headers values
func NewPutPriceIDOK() *PutPriceIDOK {

	return &PutPriceIDOK{}
}

// WithPayload adds the payload to the put price Id o k response
func (o *PutPriceIDOK) WithPayload(payload *PutPriceIDOKBody) *PutPriceIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put price Id o k response
func (o *PutPriceIDOK) SetPayload(payload *PutPriceIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutPriceIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutPriceIDBadRequestCode is the HTTP code returned for type PutPriceIDBadRequest
const PutPriceIDBadRequestCode int = 400

/*PutPriceIDBadRequest Bad Request / Validation exception

swagger:response putPriceIdBadRequest
*/
type PutPriceIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewPutPriceIDBadRequest creates PutPriceIDBadRequest with default headers values
func NewPutPriceIDBadRequest() *PutPriceIDBadRequest {

	return &PutPriceIDBadRequest{}
}

// WithPayload adds the payload to the put price Id bad request response
func (o *PutPriceIDBadRequest) WithPayload(payload *models.ErrorMessage) *PutPriceIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put price Id bad request response
func (o *PutPriceIDBadRequest) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutPriceIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutPriceIDNotFoundCode is the HTTP code returned for type PutPriceIDNotFound
const PutPriceIDNotFoundCode int = 404

/*PutPriceIDNotFound Not found

swagger:response putPriceIdNotFound
*/
type PutPriceIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewPutPriceIDNotFound creates PutPriceIDNotFound with default headers values
func NewPutPriceIDNotFound() *PutPriceIDNotFound {

	return &PutPriceIDNotFound{}
}

// WithPayload adds the payload to the put price Id not found response
func (o *PutPriceIDNotFound) WithPayload(payload *models.ErrorMessage) *PutPriceIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put price Id not found response
func (o *PutPriceIDNotFound) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutPriceIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutPriceIDUnprocessableEntityCode is the HTTP code returned for type PutPriceIDUnprocessableEntity
const PutPriceIDUnprocessableEntityCode int = 422

/*PutPriceIDUnprocessableEntity Unprocessable Entity

swagger:response putPriceIdUnprocessableEntity
*/
type PutPriceIDUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewPutPriceIDUnprocessableEntity creates PutPriceIDUnprocessableEntity with default headers values
func NewPutPriceIDUnprocessableEntity() *PutPriceIDUnprocessableEntity {

	return &PutPriceIDUnprocessableEntity{}
}

// WithPayload adds the payload to the put price Id unprocessable entity response
func (o *PutPriceIDUnprocessableEntity) WithPayload(payload *models.ErrorMessage) *PutPriceIDUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put price Id unprocessable entity response
func (o *PutPriceIDUnprocessableEntity) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutPriceIDUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PutPriceIDDefault Error

swagger:response putPriceIdDefault
*/
type PutPriceIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewPutPriceIDDefault creates PutPriceIDDefault with default headers values
func NewPutPriceIDDefault(code int) *PutPriceIDDefault {
	if code <= 0 {
		code = 500
	}

	return &PutPriceIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put price ID default response
func (o *PutPriceIDDefault) WithStatusCode(code int) *PutPriceIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put price ID default response
func (o *PutPriceIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put price ID default response
func (o *PutPriceIDDefault) WithPayload(payload *models.ErrorMessage) *PutPriceIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put price ID default response
func (o *PutPriceIDDefault) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutPriceIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
