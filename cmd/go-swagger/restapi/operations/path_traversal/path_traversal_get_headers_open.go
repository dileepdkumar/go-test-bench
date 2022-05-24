// Code generated by go-swagger; DO NOT EDIT.

package path_traversal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PathTraversalGetHeadersOpenHandlerFunc turns a function with the right signature into a path traversal get headers open handler
type PathTraversalGetHeadersOpenHandlerFunc func(PathTraversalGetHeadersOpenParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PathTraversalGetHeadersOpenHandlerFunc) Handle(params PathTraversalGetHeadersOpenParams) middleware.Responder {
	return fn(params)
}

// PathTraversalGetHeadersOpenHandler interface for that can handle valid path traversal get headers open params
type PathTraversalGetHeadersOpenHandler interface {
	Handle(PathTraversalGetHeadersOpenParams) middleware.Responder
}

// NewPathTraversalGetHeadersOpen creates a new http.Handler for the path traversal get headers open operation
func NewPathTraversalGetHeadersOpen(ctx *middleware.Context, handler PathTraversalGetHeadersOpenHandler) *PathTraversalGetHeadersOpen {
	return &PathTraversalGetHeadersOpen{Context: ctx, Handler: handler}
}

/* PathTraversalGetHeadersOpen swagger:route GET /pathTraversal/Open/headers/{safety} path_traversal pathTraversalGetHeadersOpen

demonstrates Path Traversal via headers, with vulnerable function os.Open

*/
type PathTraversalGetHeadersOpen struct {
	Context *middleware.Context
	Handler PathTraversalGetHeadersOpenHandler
}

func (o *PathTraversalGetHeadersOpen) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPathTraversalGetHeadersOpenParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
