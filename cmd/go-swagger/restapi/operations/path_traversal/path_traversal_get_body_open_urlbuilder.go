// Code generated by go-swagger; DO NOT EDIT.

package path_traversal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// PathTraversalGetBodyOpenURL generates an URL for the path traversal get body open operation
type PathTraversalGetBodyOpenURL struct {
	Safety string

	Input string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PathTraversalGetBodyOpenURL) WithBasePath(bp string) *PathTraversalGetBodyOpenURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *PathTraversalGetBodyOpenURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *PathTraversalGetBodyOpenURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/pathTraversal/Open/body/{safety}"

	safety := o.Safety
	if safety != "" {
		_path = strings.Replace(_path, "{safety}", safety, -1)
	} else {
		return nil, errors.New("safety is required on PathTraversalGetBodyOpenURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	inputQ := o.Input
	if inputQ != "" {
		qs.Set("input", inputQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *PathTraversalGetBodyOpenURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *PathTraversalGetBodyOpenURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *PathTraversalGetBodyOpenURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on PathTraversalGetBodyOpenURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on PathTraversalGetBodyOpenURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *PathTraversalGetBodyOpenURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
