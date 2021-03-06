// Code generated by go-swagger; DO NOT EDIT.

package category

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/isaacomy/ralali/models"
)

// GetCategoryIDOKCode is the HTTP code returned for type GetCategoryIDOK
const GetCategoryIDOKCode int = 200

/*GetCategoryIDOK Sends Category data

swagger:response getCategoryIdOK
*/
type GetCategoryIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Category `json:"body,omitempty"`
}

// NewGetCategoryIDOK creates GetCategoryIDOK with default headers values
func NewGetCategoryIDOK() *GetCategoryIDOK {

	return &GetCategoryIDOK{}
}

// WithPayload adds the payload to the get category Id o k response
func (o *GetCategoryIDOK) WithPayload(payload *models.Category) *GetCategoryIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get category Id o k response
func (o *GetCategoryIDOK) SetPayload(payload *models.Category) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCategoryIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCategoryIDBadRequestCode is the HTTP code returned for type GetCategoryIDBadRequest
const GetCategoryIDBadRequestCode int = 400

/*GetCategoryIDBadRequest Bad Request / Validation exception

swagger:response getCategoryIdBadRequest
*/
type GetCategoryIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetCategoryIDBadRequest creates GetCategoryIDBadRequest with default headers values
func NewGetCategoryIDBadRequest() *GetCategoryIDBadRequest {

	return &GetCategoryIDBadRequest{}
}

// WithPayload adds the payload to the get category Id bad request response
func (o *GetCategoryIDBadRequest) WithPayload(payload *models.ErrorMessage) *GetCategoryIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get category Id bad request response
func (o *GetCategoryIDBadRequest) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCategoryIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCategoryIDNotFoundCode is the HTTP code returned for type GetCategoryIDNotFound
const GetCategoryIDNotFoundCode int = 404

/*GetCategoryIDNotFound Not found

swagger:response getCategoryIdNotFound
*/
type GetCategoryIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetCategoryIDNotFound creates GetCategoryIDNotFound with default headers values
func NewGetCategoryIDNotFound() *GetCategoryIDNotFound {

	return &GetCategoryIDNotFound{}
}

// WithPayload adds the payload to the get category Id not found response
func (o *GetCategoryIDNotFound) WithPayload(payload *models.ErrorMessage) *GetCategoryIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get category Id not found response
func (o *GetCategoryIDNotFound) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCategoryIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetCategoryIDUnprocessableEntityCode is the HTTP code returned for type GetCategoryIDUnprocessableEntity
const GetCategoryIDUnprocessableEntityCode int = 422

/*GetCategoryIDUnprocessableEntity Unprocessable Entity

swagger:response getCategoryIdUnprocessableEntity
*/
type GetCategoryIDUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetCategoryIDUnprocessableEntity creates GetCategoryIDUnprocessableEntity with default headers values
func NewGetCategoryIDUnprocessableEntity() *GetCategoryIDUnprocessableEntity {

	return &GetCategoryIDUnprocessableEntity{}
}

// WithPayload adds the payload to the get category Id unprocessable entity response
func (o *GetCategoryIDUnprocessableEntity) WithPayload(payload *models.ErrorMessage) *GetCategoryIDUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get category Id unprocessable entity response
func (o *GetCategoryIDUnprocessableEntity) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCategoryIDUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetCategoryIDDefault Error

swagger:response getCategoryIdDefault
*/
type GetCategoryIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetCategoryIDDefault creates GetCategoryIDDefault with default headers values
func NewGetCategoryIDDefault(code int) *GetCategoryIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetCategoryIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get category ID default response
func (o *GetCategoryIDDefault) WithStatusCode(code int) *GetCategoryIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get category ID default response
func (o *GetCategoryIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get category ID default response
func (o *GetCategoryIDDefault) WithPayload(payload *models.ErrorMessage) *GetCategoryIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get category ID default response
func (o *GetCategoryIDDefault) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetCategoryIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
