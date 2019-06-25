package model

type FlexAd struct {
	ID        int64  `json:"-"`
	BannerURL string `json:"image_link"`
	TargetURL string `json:"target_url"`
}

func NewFlexAd(bannerURL, targetURL string) *FlexAd {
	return &FlexAd{BannerURL: bannerURL, TargetURL: targetURL}
}
