package model

type LibreFlexConfig struct {
	FlexApiServerConfig *FlexApiServerConfig `json:"flexApiServerConfig"`
	FlexAdsServerConfig *FlexAdsServerConfig `json:"flexAdsServerConfig"`
}

type FlexApiServerConfig struct {
	Enabled    bool   `json:"enabled"`
	Port       int    `json:"port"`
	DonateUrl  string `json:"donateUrl"`
	AdsApiUrl  string `json:"adsApiUrl"`
	VkApiToken string `json:"vkApiToken"`

	AdsRepoProvider string                 `json:"adsRepoProvider"`
	AdsRepoConfig   map[string]interface{} `json:"adsRepoConfig"`
}

type FlexAdsServerConfig struct {
	Enabled bool `json:"enabled"`
	Port    int  `json:"port"`

	AdsRepoProvider string                 `json:"adsRepoProvider"`
	AdsRepoConfig   map[string]interface{} `json:"adsRepoConfig"`
}
