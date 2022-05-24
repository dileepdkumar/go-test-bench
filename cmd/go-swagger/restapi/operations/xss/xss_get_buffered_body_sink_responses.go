// Code generated by go-swagger; DO NOT EDIT.

package xss

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// XSSGetBufferedBodySinkOKCode is the HTTP code returned for type XSSGetBufferedBodySinkOK
const XSSGetBufferedBodySinkOKCode int = 200

/*XSSGetBufferedBodySinkOK returns the rendered response as a string

swagger:response xssGetBufferedBodySinkOK
*/
type XSSGetBufferedBodySinkOK struct {

	/*The response when succesful query happens
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewXSSGetBufferedBodySinkOK creates XSSGetBufferedBodySinkOK with default headers values
func NewXSSGetBufferedBodySinkOK() *XSSGetBufferedBodySinkOK {

	return &XSSGetBufferedBodySinkOK{}
}

// WithPayload adds the payload to the xss get buffered body sink o k response
func (o *XSSGetBufferedBodySinkOK) WithPayload(payload string) *XSSGetBufferedBodySinkOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the xss get buffered body sink o k response
func (o *XSSGetBufferedBodySinkOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *XSSGetBufferedBodySinkOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*XSSGetBufferedBodySinkDefault Error occured

swagger:response xssGetBufferedBodySinkDefault
*/
type XSSGetBufferedBodySinkDefault struct {
	_statusCode int
}

// NewXSSGetBufferedBodySinkDefault creates XSSGetBufferedBodySinkDefault with default headers values
func NewXSSGetBufferedBodySinkDefault(code int) *XSSGetBufferedBodySinkDefault {
	if code <= 0 {
		code = 500
	}

	return &XSSGetBufferedBodySinkDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the XSS get buffered body sink default response
func (o *XSSGetBufferedBodySinkDefault) WithStatusCode(code int) *XSSGetBufferedBodySinkDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the XSS get buffered body sink default response
func (o *XSSGetBufferedBodySinkDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *XSSGetBufferedBodySinkDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
