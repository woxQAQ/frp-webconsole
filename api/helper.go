package api

import . "goa.design/goa/v3/dsl"

// ErrorResponse defines a response with an empty body.
func ErrorResponse(name string, code int, fn ...func()) {
	Response(name, code, func() {
		Body(Empty)
		if len(fn) > 0 {
			fn[0]()
		}
	})
}

func TypeTimeFormat(name string, description string) {
	Attribute(name, String, func() {
		Format("date-time")
		Description(description)
	})
}
