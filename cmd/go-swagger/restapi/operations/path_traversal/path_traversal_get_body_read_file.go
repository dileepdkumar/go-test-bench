// Code generated by go-swagger; DO NOT EDIT.

package path_traversal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PathTraversalGetBodyReadFileHandlerFunc turns a function with the right signature into a path traversal get body read file handler
type PathTraversalGetBodyReadFileHandlerFunc func(PathTraversalGetBodyReadFileParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PathTraversalGetBodyReadFileHandlerFunc) Handle(params PathTraversalGetBodyReadFileParams) middleware.Responder {
	return fn(params)
}

// PathTraversalGetBodyReadFileHandler interface for that can handle valid path traversal get body read file params
type PathTraversalGetBodyReadFileHandler interface {
	Handle(PathTraversalGetBodyReadFileParams) middleware.Responder
}

// NewPathTraversalGetBodyReadFile creates a new http.Handler for the path traversal get body read file operation
func NewPathTraversalGetBodyReadFile(ctx *middleware.Context, handler PathTraversalGetBodyReadFileHandler) *PathTraversalGetBodyReadFile {
	return &PathTraversalGetBodyReadFile{Context: ctx, Handler: handler}
}

/* PathTraversalGetBodyReadFile swagger:route GET /pathTraversal/ReadFile/body/{safety} path_traversal pathTraversalGetBodyReadFile

demonstrates Path Traversal via body, with vulnerable function os.ReadFile

*/
type PathTraversalGetBodyReadFile struct {
	Context *middleware.Context
	Handler PathTraversalGetBodyReadFileHandler
}

func (o *PathTraversalGetBodyReadFile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPathTraversalGetBodyReadFileParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
