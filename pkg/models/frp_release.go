package models

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/go-github/v68/github"
)

type FrpRelease struct {
	TagName   string    `json:"tag_name"`
	Size      int64     `json:"size"` // in bytes
	Assets    FrpAsset  `json:"assets"`
	CreatedAt time.Time `json:"created_at"`
}

// FrpAsset 表示一个匹配当前系统的发布文件
type FrpAsset struct {
	Name        string `json:"name"`
	DownloadURL string `json:"browser_download_url"`
	Size        int64  `json:"size"`
	Downloads   int64  `json:"download_count"`
}

func FilterAssetByOS(assets []*github.ReleaseAsset, system System) (FrpAsset, error) {
	pattern := fmt.Sprintf(".*%s_%s", system.OS, system.Arch)
	for _, asset := range assets {
		if matched, _ := regexp.MatchString(pattern, asset.GetName()); matched {
			return FrpAsset{
				Name:        asset.GetName(),
				DownloadURL: asset.GetBrowserDownloadURL(),
				Size:        int64(asset.GetSize()),
				Downloads:   int64(asset.GetDownloadCount()),
			}, nil
		}
	}
	return FrpAsset{}, fmt.Errorf("no matching asset found for %s_%s", system.OS, system.Arch)
}

func NewFrpRelease(release *github.RepositoryRelease, system System) (*FrpRelease, error) {
	assets, err := FilterAssetByOS(release.Assets, system)
	if err != nil {
		return nil, err
	}
	return &FrpRelease{
		TagName: release.GetTagName(),
		Assets:  assets,
	}, nil
}
