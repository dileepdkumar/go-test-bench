// Code generated by go-swagger; DO NOT EDIT.

package xss

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// XSSGetResponseSinkHandlerFunc turns a function with the right signature into a XSS get response sink handler
type XSSGetResponseSinkHandlerFunc func(XSSGetResponseSinkParams) middleware.Responder

// Handle executing the request and returning a response
func (fn XSSGetResponseSinkHandlerFunc) Handle(params XSSGetResponseSinkParams) middleware.Responder {
	return fn(params)
}

// XSSGetResponseSinkHandler interface for that can handle valid XSS get response sink params
type XSSGetResponseSinkHandler interface {
	Handle(XSSGetResponseSinkParams) middleware.Responder
}

// NewXSSGetResponseSink creates a new http.Handler for the XSS get response sink operation
func NewXSSGetResponseSink(ctx *middleware.Context, handler XSSGetResponseSinkHandler) *XSSGetResponseSink {
	return &XSSGetResponseSink{Context: ctx, Handler: handler}
}

/* XSSGetResponseSink swagger:route GET /xss/Sink/response/{safety} xss xssGetResponseSink

demonstrates Reflected XSS via response, with vulnerable function _

*/
type XSSGetResponseSink struct {
	Context *middleware.Context
	Handler XSSGetResponseSinkHandler
}

func (o *XSSGetResponseSink) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewXSSGetResponseSinkParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
