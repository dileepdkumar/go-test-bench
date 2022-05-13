package common

import (
	"errors"
	"fmt"
)

// HandlerFn is a framework-agnostic function to handle a vulnerable endpoint.
// `opaque` can be set to some framework-specific struct - for example, gin.Context.
type HandlerFn func(safety Safety, in string, opaque interface{}) string

// VulnerableFnWrapper
type VulnerableFnWrapper func(opaque interface{}, payload string) (data string, err error)

//unlike HandlerFn, isTmpl is omitted as it seems that'll never be used
func GenericHandler(s Sink, safety Safety, payload string, opaque interface{}) (data string) {
	if s.Sanitize == nil {
		return fmt.Sprintf("sink %#v: internal error - Sanitizer cannot be nil", s)
	}
	if s.VulnerableFnWrapper == nil {
		return fmt.Sprintf("sink %#v: internal error - VulnerableFnWrapper cannot be nil", s)
	}
	switch safety {
	case Unsafe:
		// nothing to do here
	case Safe:
		payload = s.Sanitize(payload)
	default:
		return "NOOP"
	}
	var err error
	data, err = s.VulnerableFnWrapper(opaque, payload)
	if err == NoDecoration {
		//the vulnerable function writes sufficient information - no need to decorate.
		return data
	}
	switch safety {
	case Unsafe:
		if err != nil {
			data = fmt.Sprintf("%q: unsafe action failed: payload=%q err=%s", s.Name, payload, err)
		} else if len(data) == 0 {
			data = fmt.Sprintf("%q: unsafe action reported no error. payload=%q", s.Name, payload)
		}
	case Safe:
		if err != nil {
			return fmt.Sprintf("%q: safe action returned data=%s err=%s with payload %s", s.Name, data, err, payload)
		} else if len(data) == 0 {
			return fmt.Sprintf("%q: safe action returned no data or error. payload=%s", s.Name, payload)
		}
	}
	return data
}

// special error indicating no additional data should be written out
var NoDecoration = errors.New("no decoration needed")
