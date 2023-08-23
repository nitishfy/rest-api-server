// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetAPIAPIVersionOKCode is the HTTP code returned for type GetAPIAPIVersionOK
const GetAPIAPIVersionOKCode int = 200

/*
GetAPIAPIVersionOK returns the api version

swagger:response getApiApiVersionOK
*/
type GetAPIAPIVersionOK struct {

	/*
	  In: Body
	*/
	Payload int64 `json:"body,omitempty"`
}

// NewGetAPIAPIVersionOK creates GetAPIAPIVersionOK with default headers values
func NewGetAPIAPIVersionOK() *GetAPIAPIVersionOK {

	return &GetAPIAPIVersionOK{}
}

// WithPayload adds the payload to the get Api Api version o k response
func (o *GetAPIAPIVersionOK) WithPayload(payload int64) *GetAPIAPIVersionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get Api Api version o k response
func (o *GetAPIAPIVersionOK) SetPayload(payload int64) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAPIAPIVersionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAPIAPIVersionBadRequestCode is the HTTP code returned for type GetAPIAPIVersionBadRequest
const GetAPIAPIVersionBadRequestCode int = 400

/*
GetAPIAPIVersionBadRequest Invaid character as api version provided

swagger:response getApiApiVersionBadRequest
*/
type GetAPIAPIVersionBadRequest struct {
}

// NewGetAPIAPIVersionBadRequest creates GetAPIAPIVersionBadRequest with default headers values
func NewGetAPIAPIVersionBadRequest() *GetAPIAPIVersionBadRequest {

	return &GetAPIAPIVersionBadRequest{}
}

// WriteResponse to the client
func (o *GetAPIAPIVersionBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}