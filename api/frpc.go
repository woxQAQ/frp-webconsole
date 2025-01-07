package api

import (
	. "goa.design/goa/v3/dsl"
)

var FrpRelease = Type("FrpRelease", func() {
	Description("A bottle of wine")
	Attribute("tag_name", String, "Tag name of release")
	Attribute("size", Int, "Size of release")
	Attribute("assets", FrpAsset, "Assets of release")
	TypeTimeFormat("created_at", "Created at")
})

var FrpAsset = Type("FrpAsset", func() {
	Description("A bottle of wine")
	Attribute("name", String, "Name of asset")
	Attribute("download_url", String, "Download URL of asset")
	Attribute("size", Int, "Size of asset")
	Attribute("downloads", Int, "Downloads of asset")
})

var _ = Service("frpc", func() {
	Description("FRP Webconsole")

	Method("ListFrpRelease", func() {
		Payload(func() {
			Attribute("os", String)
			Attribute("arch", String)
		})
		Result(FrpRelease)
		HTTP(func() {
			Params(func() {
				Param("os", String, "OS")
				Param("arch", String, "Architecture")
			})
			GET("/frp/release")
		})
	})
})
