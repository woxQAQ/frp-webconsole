package models

type FrpRelease struct {
	TagName string `json:"tag_name"`
	Assets  []string
}
