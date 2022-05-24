// Code generated by go-swagger; DO NOT EDIT.

package cmd_injection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// CmdInjectionGetCookiesCommandContextOKCode is the HTTP code returned for type CmdInjectionGetCookiesCommandContextOK
const CmdInjectionGetCookiesCommandContextOKCode int = 200

/*CmdInjectionGetCookiesCommandContextOK returns the rendered response as a string

swagger:response cmdInjectionGetCookiesCommandContextOK
*/
type CmdInjectionGetCookiesCommandContextOK struct {

	/*The response when succesful query happens
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewCmdInjectionGetCookiesCommandContextOK creates CmdInjectionGetCookiesCommandContextOK with default headers values
func NewCmdInjectionGetCookiesCommandContextOK() *CmdInjectionGetCookiesCommandContextOK {

	return &CmdInjectionGetCookiesCommandContextOK{}
}

// WithPayload adds the payload to the cmd injection get cookies command context o k response
func (o *CmdInjectionGetCookiesCommandContextOK) WithPayload(payload string) *CmdInjectionGetCookiesCommandContextOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the cmd injection get cookies command context o k response
func (o *CmdInjectionGetCookiesCommandContextOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CmdInjectionGetCookiesCommandContextOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*CmdInjectionGetCookiesCommandContextDefault Error occured

swagger:response cmdInjectionGetCookiesCommandContextDefault
*/
type CmdInjectionGetCookiesCommandContextDefault struct {
	_statusCode int
}

// NewCmdInjectionGetCookiesCommandContextDefault creates CmdInjectionGetCookiesCommandContextDefault with default headers values
func NewCmdInjectionGetCookiesCommandContextDefault(code int) *CmdInjectionGetCookiesCommandContextDefault {
	if code <= 0 {
		code = 500
	}

	return &CmdInjectionGetCookiesCommandContextDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the cmd injection get cookies command context default response
func (o *CmdInjectionGetCookiesCommandContextDefault) WithStatusCode(code int) *CmdInjectionGetCookiesCommandContextDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the cmd injection get cookies command context default response
func (o *CmdInjectionGetCookiesCommandContextDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *CmdInjectionGetCookiesCommandContextDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
