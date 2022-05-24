// Code generated by go-swagger; DO NOT EDIT.

package cmd_injection

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewCmdInjectionGetQueryCommandParams creates a new CmdInjectionGetQueryCommandParams object
//
// There are no default values defined in the spec.
func NewCmdInjectionGetQueryCommandParams() CmdInjectionGetQueryCommandParams {

	return CmdInjectionGetQueryCommandParams{}
}

// CmdInjectionGetQueryCommandParams contains all the bound params for the cmd injection get query command operation
// typically these are obtained from a http.Request
//
// swagger:parameters CmdInjectionGetQueryCommand
type CmdInjectionGetQueryCommandParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*the user provided input for the query vulnerability
	  Required: true
	  In: query
	*/
	Input string
	/*safety qualifier
	  Required: true
	  In: path
	*/
	Safety string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCmdInjectionGetQueryCommandParams() beforehand.
func (o *CmdInjectionGetQueryCommandParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qInput, qhkInput, _ := qs.GetOK("input")
	if err := o.bindInput(qInput, qhkInput, route.Formats); err != nil {
		res = append(res, err)
	}

	rSafety, rhkSafety, _ := route.Params.GetOK("safety")
	if err := o.bindSafety(rSafety, rhkSafety, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindInput binds and validates parameter Input from query.
func (o *CmdInjectionGetQueryCommandParams) bindInput(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("input", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("input", "query", raw); err != nil {
		return err
	}
	o.Input = raw

	return nil
}

// bindSafety binds and validates parameter Safety from path.
func (o *CmdInjectionGetQueryCommandParams) bindSafety(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Safety = raw

	if err := o.validateSafety(formats); err != nil {
		return err
	}

	return nil
}

// validateSafety carries on validations for parameter Safety
func (o *CmdInjectionGetQueryCommandParams) validateSafety(formats strfmt.Registry) error {

	if err := validate.EnumCase("safety", "path", o.Safety, []interface{}{"safe", "unsafe", "noop"}, true); err != nil {
		return err
	}

	return nil
}
