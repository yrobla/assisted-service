// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// EnableHostOKCode is the HTTP code returned for type EnableHostOK
const EnableHostOKCode int = 200

/*EnableHostOK Success.

swagger:response enableHostOK
*/
type EnableHostOK struct {

	/*
	  In: Body
	*/
	Payload *models.Host `json:"body,omitempty"`
}

// NewEnableHostOK creates EnableHostOK with default headers values
func NewEnableHostOK() *EnableHostOK {

	return &EnableHostOK{}
}

// WithPayload adds the payload to the enable host o k response
func (o *EnableHostOK) WithPayload(payload *models.Host) *EnableHostOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enable host o k response
func (o *EnableHostOK) SetPayload(payload *models.Host) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnableHostOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnableHostUnauthorizedCode is the HTTP code returned for type EnableHostUnauthorized
const EnableHostUnauthorizedCode int = 401

/*EnableHostUnauthorized Unauthorized.

swagger:response enableHostUnauthorized
*/
type EnableHostUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewEnableHostUnauthorized creates EnableHostUnauthorized with default headers values
func NewEnableHostUnauthorized() *EnableHostUnauthorized {

	return &EnableHostUnauthorized{}
}

// WithPayload adds the payload to the enable host unauthorized response
func (o *EnableHostUnauthorized) WithPayload(payload *models.Error) *EnableHostUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enable host unauthorized response
func (o *EnableHostUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnableHostUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnableHostForbiddenCode is the HTTP code returned for type EnableHostForbidden
const EnableHostForbiddenCode int = 403

/*EnableHostForbidden Forbidden.

swagger:response enableHostForbidden
*/
type EnableHostForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewEnableHostForbidden creates EnableHostForbidden with default headers values
func NewEnableHostForbidden() *EnableHostForbidden {

	return &EnableHostForbidden{}
}

// WithPayload adds the payload to the enable host forbidden response
func (o *EnableHostForbidden) WithPayload(payload *models.Error) *EnableHostForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enable host forbidden response
func (o *EnableHostForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnableHostForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnableHostNotFoundCode is the HTTP code returned for type EnableHostNotFound
const EnableHostNotFoundCode int = 404

/*EnableHostNotFound Error.

swagger:response enableHostNotFound
*/
type EnableHostNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewEnableHostNotFound creates EnableHostNotFound with default headers values
func NewEnableHostNotFound() *EnableHostNotFound {

	return &EnableHostNotFound{}
}

// WithPayload adds the payload to the enable host not found response
func (o *EnableHostNotFound) WithPayload(payload *models.Error) *EnableHostNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enable host not found response
func (o *EnableHostNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnableHostNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnableHostMethodNotAllowedCode is the HTTP code returned for type EnableHostMethodNotAllowed
const EnableHostMethodNotAllowedCode int = 405

/*EnableHostMethodNotAllowed Method Not Allowed.

swagger:response enableHostMethodNotAllowed
*/
type EnableHostMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewEnableHostMethodNotAllowed creates EnableHostMethodNotAllowed with default headers values
func NewEnableHostMethodNotAllowed() *EnableHostMethodNotAllowed {

	return &EnableHostMethodNotAllowed{}
}

// WithPayload adds the payload to the enable host method not allowed response
func (o *EnableHostMethodNotAllowed) WithPayload(payload *models.Error) *EnableHostMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enable host method not allowed response
func (o *EnableHostMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnableHostMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnableHostConflictCode is the HTTP code returned for type EnableHostConflict
const EnableHostConflictCode int = 409

/*EnableHostConflict Error.

swagger:response enableHostConflict
*/
type EnableHostConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewEnableHostConflict creates EnableHostConflict with default headers values
func NewEnableHostConflict() *EnableHostConflict {

	return &EnableHostConflict{}
}

// WithPayload adds the payload to the enable host conflict response
func (o *EnableHostConflict) WithPayload(payload *models.Error) *EnableHostConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enable host conflict response
func (o *EnableHostConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnableHostConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// EnableHostInternalServerErrorCode is the HTTP code returned for type EnableHostInternalServerError
const EnableHostInternalServerErrorCode int = 500

/*EnableHostInternalServerError Error.

swagger:response enableHostInternalServerError
*/
type EnableHostInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewEnableHostInternalServerError creates EnableHostInternalServerError with default headers values
func NewEnableHostInternalServerError() *EnableHostInternalServerError {

	return &EnableHostInternalServerError{}
}

// WithPayload adds the payload to the enable host internal server error response
func (o *EnableHostInternalServerError) WithPayload(payload *models.Error) *EnableHostInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the enable host internal server error response
func (o *EnableHostInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *EnableHostInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
