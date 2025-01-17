// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// PostConfigVlanVlanIDMemberIfNameNoContentCode is the HTTP code returned for type PostConfigVlanVlanIDMemberIfNameNoContent
const PostConfigVlanVlanIDMemberIfNameNoContentCode int = 204

/*PostConfigVlanVlanIDMemberIfNameNoContent OK

swagger:response postConfigVlanVlanIdMemberIfNameNoContent
*/
type PostConfigVlanVlanIDMemberIfNameNoContent struct {
}

// NewPostConfigVlanVlanIDMemberIfNameNoContent creates PostConfigVlanVlanIDMemberIfNameNoContent with default headers values
func NewPostConfigVlanVlanIDMemberIfNameNoContent() *PostConfigVlanVlanIDMemberIfNameNoContent {

	return &PostConfigVlanVlanIDMemberIfNameNoContent{}
}

// WriteResponse to the client
func (o *PostConfigVlanVlanIDMemberIfNameNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PostConfigVlanVlanIDMemberIfNameBadRequestCode is the HTTP code returned for type PostConfigVlanVlanIDMemberIfNameBadRequest
const PostConfigVlanVlanIDMemberIfNameBadRequestCode int = 400

/*PostConfigVlanVlanIDMemberIfNameBadRequest Malformed arguments for API call

swagger:response postConfigVlanVlanIdMemberIfNameBadRequest
*/
type PostConfigVlanVlanIDMemberIfNameBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanVlanIDMemberIfNameBadRequest creates PostConfigVlanVlanIDMemberIfNameBadRequest with default headers values
func NewPostConfigVlanVlanIDMemberIfNameBadRequest() *PostConfigVlanVlanIDMemberIfNameBadRequest {

	return &PostConfigVlanVlanIDMemberIfNameBadRequest{}
}

// WithPayload adds the payload to the post config vlan vlan Id member if name bad request response
func (o *PostConfigVlanVlanIDMemberIfNameBadRequest) WithPayload(payload *models.Error) *PostConfigVlanVlanIDMemberIfNameBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan vlan Id member if name bad request response
func (o *PostConfigVlanVlanIDMemberIfNameBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanVlanIDMemberIfNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigVlanVlanIDMemberIfNameUnauthorizedCode is the HTTP code returned for type PostConfigVlanVlanIDMemberIfNameUnauthorized
const PostConfigVlanVlanIDMemberIfNameUnauthorizedCode int = 401

/*PostConfigVlanVlanIDMemberIfNameUnauthorized Invalid authentication credentials

swagger:response postConfigVlanVlanIdMemberIfNameUnauthorized
*/
type PostConfigVlanVlanIDMemberIfNameUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanVlanIDMemberIfNameUnauthorized creates PostConfigVlanVlanIDMemberIfNameUnauthorized with default headers values
func NewPostConfigVlanVlanIDMemberIfNameUnauthorized() *PostConfigVlanVlanIDMemberIfNameUnauthorized {

	return &PostConfigVlanVlanIDMemberIfNameUnauthorized{}
}

// WithPayload adds the payload to the post config vlan vlan Id member if name unauthorized response
func (o *PostConfigVlanVlanIDMemberIfNameUnauthorized) WithPayload(payload *models.Error) *PostConfigVlanVlanIDMemberIfNameUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan vlan Id member if name unauthorized response
func (o *PostConfigVlanVlanIDMemberIfNameUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanVlanIDMemberIfNameUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigVlanVlanIDMemberIfNameForbiddenCode is the HTTP code returned for type PostConfigVlanVlanIDMemberIfNameForbidden
const PostConfigVlanVlanIDMemberIfNameForbiddenCode int = 403

/*PostConfigVlanVlanIDMemberIfNameForbidden Capacity insufficient

swagger:response postConfigVlanVlanIdMemberIfNameForbidden
*/
type PostConfigVlanVlanIDMemberIfNameForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanVlanIDMemberIfNameForbidden creates PostConfigVlanVlanIDMemberIfNameForbidden with default headers values
func NewPostConfigVlanVlanIDMemberIfNameForbidden() *PostConfigVlanVlanIDMemberIfNameForbidden {

	return &PostConfigVlanVlanIDMemberIfNameForbidden{}
}

// WithPayload adds the payload to the post config vlan vlan Id member if name forbidden response
func (o *PostConfigVlanVlanIDMemberIfNameForbidden) WithPayload(payload *models.Error) *PostConfigVlanVlanIDMemberIfNameForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan vlan Id member if name forbidden response
func (o *PostConfigVlanVlanIDMemberIfNameForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanVlanIDMemberIfNameForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigVlanVlanIDMemberIfNameNotFoundCode is the HTTP code returned for type PostConfigVlanVlanIDMemberIfNameNotFound
const PostConfigVlanVlanIDMemberIfNameNotFoundCode int = 404

/*PostConfigVlanVlanIDMemberIfNameNotFound Vlan interface is not defined

swagger:response postConfigVlanVlanIdMemberIfNameNotFound
*/
type PostConfigVlanVlanIDMemberIfNameNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanVlanIDMemberIfNameNotFound creates PostConfigVlanVlanIDMemberIfNameNotFound with default headers values
func NewPostConfigVlanVlanIDMemberIfNameNotFound() *PostConfigVlanVlanIDMemberIfNameNotFound {

	return &PostConfigVlanVlanIDMemberIfNameNotFound{}
}

// WithPayload adds the payload to the post config vlan vlan Id member if name not found response
func (o *PostConfigVlanVlanIDMemberIfNameNotFound) WithPayload(payload *models.Error) *PostConfigVlanVlanIDMemberIfNameNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan vlan Id member if name not found response
func (o *PostConfigVlanVlanIDMemberIfNameNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanVlanIDMemberIfNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigVlanVlanIDMemberIfNameConflictCode is the HTTP code returned for type PostConfigVlanVlanIDMemberIfNameConflict
const PostConfigVlanVlanIDMemberIfNameConflictCode int = 409

/*PostConfigVlanVlanIDMemberIfNameConflict Resource Conflict. VLAN member already exists on this VLAN interface OR Vlan member is being added to 2nd Vlan inteface as an untagged member.

swagger:response postConfigVlanVlanIdMemberIfNameConflict
*/
type PostConfigVlanVlanIDMemberIfNameConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanVlanIDMemberIfNameConflict creates PostConfigVlanVlanIDMemberIfNameConflict with default headers values
func NewPostConfigVlanVlanIDMemberIfNameConflict() *PostConfigVlanVlanIDMemberIfNameConflict {

	return &PostConfigVlanVlanIDMemberIfNameConflict{}
}

// WithPayload adds the payload to the post config vlan vlan Id member if name conflict response
func (o *PostConfigVlanVlanIDMemberIfNameConflict) WithPayload(payload *models.Error) *PostConfigVlanVlanIDMemberIfNameConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan vlan Id member if name conflict response
func (o *PostConfigVlanVlanIDMemberIfNameConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanVlanIDMemberIfNameConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigVlanVlanIDMemberIfNameInternalServerErrorCode is the HTTP code returned for type PostConfigVlanVlanIDMemberIfNameInternalServerError
const PostConfigVlanVlanIDMemberIfNameInternalServerErrorCode int = 500

/*PostConfigVlanVlanIDMemberIfNameInternalServerError Internal service error

swagger:response postConfigVlanVlanIdMemberIfNameInternalServerError
*/
type PostConfigVlanVlanIDMemberIfNameInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanVlanIDMemberIfNameInternalServerError creates PostConfigVlanVlanIDMemberIfNameInternalServerError with default headers values
func NewPostConfigVlanVlanIDMemberIfNameInternalServerError() *PostConfigVlanVlanIDMemberIfNameInternalServerError {

	return &PostConfigVlanVlanIDMemberIfNameInternalServerError{}
}

// WithPayload adds the payload to the post config vlan vlan Id member if name internal server error response
func (o *PostConfigVlanVlanIDMemberIfNameInternalServerError) WithPayload(payload *models.Error) *PostConfigVlanVlanIDMemberIfNameInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan vlan Id member if name internal server error response
func (o *PostConfigVlanVlanIDMemberIfNameInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanVlanIDMemberIfNameInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigVlanVlanIDMemberIfNameServiceUnavailableCode is the HTTP code returned for type PostConfigVlanVlanIDMemberIfNameServiceUnavailable
const PostConfigVlanVlanIDMemberIfNameServiceUnavailableCode int = 503

/*PostConfigVlanVlanIDMemberIfNameServiceUnavailable Maintanence mode

swagger:response postConfigVlanVlanIdMemberIfNameServiceUnavailable
*/
type PostConfigVlanVlanIDMemberIfNameServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanVlanIDMemberIfNameServiceUnavailable creates PostConfigVlanVlanIDMemberIfNameServiceUnavailable with default headers values
func NewPostConfigVlanVlanIDMemberIfNameServiceUnavailable() *PostConfigVlanVlanIDMemberIfNameServiceUnavailable {

	return &PostConfigVlanVlanIDMemberIfNameServiceUnavailable{}
}

// WithPayload adds the payload to the post config vlan vlan Id member if name service unavailable response
func (o *PostConfigVlanVlanIDMemberIfNameServiceUnavailable) WithPayload(payload *models.Error) *PostConfigVlanVlanIDMemberIfNameServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan vlan Id member if name service unavailable response
func (o *PostConfigVlanVlanIDMemberIfNameServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanVlanIDMemberIfNameServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
