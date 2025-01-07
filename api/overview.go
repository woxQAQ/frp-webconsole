package api

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("frp-webconsole", func() {

	Title("FRP Webconsole")
	Description("FRP Webconsole")
	HTTP(func() {
		Path("/api/v1")
	})
	Error("InternalError", func() {
		Fault()
	})
	Error("Timeout", func() {
		Timeout()
	})
	Error("Unauthorized")
	License(func() {
		URL("https://github.com/woxQAQ/frp-webconsole/blob/main/LICENSE")
	})
	Version("1.0.0")
	Server("frp-webconsole", func() {
		Services("frpc")
		Host("localhost", func() {
			URI("http://localhost:8080")
			URI("grpc://localhost:8080")
		})
	})
})
