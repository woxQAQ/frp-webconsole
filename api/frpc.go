package api

import (
	. "goa.design/goa/v3/dsl"
)

var FrpRelease = Type("FrpRelease", func() {
	Description("A release of frp")
	Attribute("tag_name", String, "Tag name of release", func() {
		Example("v0.55.0")
	})
	Attribute("assets", ArrayOf(FrpAsset), "Assets of release")
	TypeTimeFormat("created_at", "Created at")
})

var FrpAsset = Type("FrpAsset", func() {
	Description("A asset of frp release")
	Attribute("name", String, "Name of asset", func() {
		Example("frp_0.55.0_linux_amd64.tar.gz", "frp_0.55.0_linux_arm64.tar.gz")
	})
	Attribute("download_url", String, "Download URL of asset", func() {
		Example("https://github.com/fatedier/frp/releases/download/v0.55.0/frp_0.55.0_linux_amd64.tar.gz", "https://github.com/fatedier/frp/releases/download/v0.55.0/frp_0.55.0_linux_arm64.tar.gz")
	})
	Attribute("size", Int, "Size of asset", func() {
		Example(1024)
	})
	Attribute("downloads", Int, "Downloads of asset", func() {
		Example(100)
	})
})

var _ = Service("frpc", func() {
	Description("FRP Webconsole")

	Method("ListFrpRelease", func() {
		Payload(func() {
			Attribute("os", String)
			Attribute("arch", String)
		})
		Result(ArrayOf(FrpRelease))
		HTTP(func() {
			Params(func() {
				Param("os", String, "OS")
				Param("arch", String, "Architecture")
			})
			GET("/frp/release")
		})
	})
})
