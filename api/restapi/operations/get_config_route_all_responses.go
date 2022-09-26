// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// GetConfigRouteAllOKCode is the HTTP code returned for type GetConfigRouteAllOK
const GetConfigRouteAllOKCode int = 200

/*GetConfigRouteAllOK OK

swagger:response getConfigRouteAllOK
*/
type GetConfigRouteAllOK struct {

	/*
	  In: Body
	*/
	Payload *GetConfigRouteAllOKBody `json:"body,omitempty"`
}

// NewGetConfigRouteAllOK creates GetConfigRouteAllOK with default headers values
func NewGetConfigRouteAllOK() *GetConfigRouteAllOK {

	return &GetConfigRouteAllOK{}
}

// WithPayload adds the payload to the get config route all o k response
func (o *GetConfigRouteAllOK) WithPayload(payload *GetConfigRouteAllOKBody) *GetConfigRouteAllOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config route all o k response
func (o *GetConfigRouteAllOK) SetPayload(payload *GetConfigRouteAllOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigRouteAllOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigRouteAllBadRequestCode is the HTTP code returned for type GetConfigRouteAllBadRequest
const GetConfigRouteAllBadRequestCode int = 400

/*GetConfigRouteAllBadRequest Malformed arguments for API call

swagger:response getConfigRouteAllBadRequest
*/
type GetConfigRouteAllBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigRouteAllBadRequest creates GetConfigRouteAllBadRequest with default headers values
func NewGetConfigRouteAllBadRequest() *GetConfigRouteAllBadRequest {

	return &GetConfigRouteAllBadRequest{}
}

// WithPayload adds the payload to the get config route all bad request response
func (o *GetConfigRouteAllBadRequest) WithPayload(payload *models.Error) *GetConfigRouteAllBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config route all bad request response
func (o *GetConfigRouteAllBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigRouteAllBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigRouteAllUnauthorizedCode is the HTTP code returned for type GetConfigRouteAllUnauthorized
const GetConfigRouteAllUnauthorizedCode int = 401

/*GetConfigRouteAllUnauthorized Invalid authentication credentials

swagger:response getConfigRouteAllUnauthorized
*/
type GetConfigRouteAllUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigRouteAllUnauthorized creates GetConfigRouteAllUnauthorized with default headers values
func NewGetConfigRouteAllUnauthorized() *GetConfigRouteAllUnauthorized {

	return &GetConfigRouteAllUnauthorized{}
}

// WithPayload adds the payload to the get config route all unauthorized response
func (o *GetConfigRouteAllUnauthorized) WithPayload(payload *models.Error) *GetConfigRouteAllUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config route all unauthorized response
func (o *GetConfigRouteAllUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigRouteAllUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigRouteAllForbiddenCode is the HTTP code returned for type GetConfigRouteAllForbidden
const GetConfigRouteAllForbiddenCode int = 403

/*GetConfigRouteAllForbidden Capacity insufficient

swagger:response getConfigRouteAllForbidden
*/
type GetConfigRouteAllForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigRouteAllForbidden creates GetConfigRouteAllForbidden with default headers values
func NewGetConfigRouteAllForbidden() *GetConfigRouteAllForbidden {

	return &GetConfigRouteAllForbidden{}
}

// WithPayload adds the payload to the get config route all forbidden response
func (o *GetConfigRouteAllForbidden) WithPayload(payload *models.Error) *GetConfigRouteAllForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config route all forbidden response
func (o *GetConfigRouteAllForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigRouteAllForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigRouteAllNotFoundCode is the HTTP code returned for type GetConfigRouteAllNotFound
const GetConfigRouteAllNotFoundCode int = 404

/*GetConfigRouteAllNotFound Resource not found

swagger:response getConfigRouteAllNotFound
*/
type GetConfigRouteAllNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigRouteAllNotFound creates GetConfigRouteAllNotFound with default headers values
func NewGetConfigRouteAllNotFound() *GetConfigRouteAllNotFound {

	return &GetConfigRouteAllNotFound{}
}

// WithPayload adds the payload to the get config route all not found response
func (o *GetConfigRouteAllNotFound) WithPayload(payload *models.Error) *GetConfigRouteAllNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config route all not found response
func (o *GetConfigRouteAllNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigRouteAllNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigRouteAllConflictCode is the HTTP code returned for type GetConfigRouteAllConflict
const GetConfigRouteAllConflictCode int = 409

/*GetConfigRouteAllConflict Resource Conflict. VLAN already exists OR dependency VRF/VNET not found

swagger:response getConfigRouteAllConflict
*/
type GetConfigRouteAllConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigRouteAllConflict creates GetConfigRouteAllConflict with default headers values
func NewGetConfigRouteAllConflict() *GetConfigRouteAllConflict {

	return &GetConfigRouteAllConflict{}
}

// WithPayload adds the payload to the get config route all conflict response
func (o *GetConfigRouteAllConflict) WithPayload(payload *models.Error) *GetConfigRouteAllConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config route all conflict response
func (o *GetConfigRouteAllConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigRouteAllConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigRouteAllInternalServerErrorCode is the HTTP code returned for type GetConfigRouteAllInternalServerError
const GetConfigRouteAllInternalServerErrorCode int = 500

/*GetConfigRouteAllInternalServerError Internal service error

swagger:response getConfigRouteAllInternalServerError
*/
type GetConfigRouteAllInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigRouteAllInternalServerError creates GetConfigRouteAllInternalServerError with default headers values
func NewGetConfigRouteAllInternalServerError() *GetConfigRouteAllInternalServerError {

	return &GetConfigRouteAllInternalServerError{}
}

// WithPayload adds the payload to the get config route all internal server error response
func (o *GetConfigRouteAllInternalServerError) WithPayload(payload *models.Error) *GetConfigRouteAllInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config route all internal server error response
func (o *GetConfigRouteAllInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigRouteAllInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigRouteAllServiceUnavailableCode is the HTTP code returned for type GetConfigRouteAllServiceUnavailable
const GetConfigRouteAllServiceUnavailableCode int = 503

/*GetConfigRouteAllServiceUnavailable Maintanence mode

swagger:response getConfigRouteAllServiceUnavailable
*/
type GetConfigRouteAllServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigRouteAllServiceUnavailable creates GetConfigRouteAllServiceUnavailable with default headers values
func NewGetConfigRouteAllServiceUnavailable() *GetConfigRouteAllServiceUnavailable {

	return &GetConfigRouteAllServiceUnavailable{}
}

// WithPayload adds the payload to the get config route all service unavailable response
func (o *GetConfigRouteAllServiceUnavailable) WithPayload(payload *models.Error) *GetConfigRouteAllServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config route all service unavailable response
func (o *GetConfigRouteAllServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigRouteAllServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
