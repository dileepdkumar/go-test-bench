// Code generated by go-swagger; DO NOT EDIT.

package path_traversal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PathTraversalGetQueryReadFileHandlerFunc turns a function with the right signature into a path traversal get query read file handler
type PathTraversalGetQueryReadFileHandlerFunc func(PathTraversalGetQueryReadFileParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PathTraversalGetQueryReadFileHandlerFunc) Handle(params PathTraversalGetQueryReadFileParams) middleware.Responder {
	return fn(params)
}

// PathTraversalGetQueryReadFileHandler interface for that can handle valid path traversal get query read file params
type PathTraversalGetQueryReadFileHandler interface {
	Handle(PathTraversalGetQueryReadFileParams) middleware.Responder
}

// NewPathTraversalGetQueryReadFile creates a new http.Handler for the path traversal get query read file operation
func NewPathTraversalGetQueryReadFile(ctx *middleware.Context, handler PathTraversalGetQueryReadFileHandler) *PathTraversalGetQueryReadFile {
	return &PathTraversalGetQueryReadFile{Context: ctx, Handler: handler}
}

/* PathTraversalGetQueryReadFile swagger:route GET /pathTraversal/ReadFile/query/{safety} path_traversal pathTraversalGetQueryReadFile

demonstrates Path Traversal via query, with vulnerable function os.ReadFile

*/
type PathTraversalGetQueryReadFile struct {
	Context *middleware.Context
	Handler PathTraversalGetQueryReadFileHandler
}

func (o *PathTraversalGetQueryReadFile) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPathTraversalGetQueryReadFileParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
