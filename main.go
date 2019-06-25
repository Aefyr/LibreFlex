package main

import (
	"encoding/json"
	"io/ioutil"
	ads "libreflex/adsservice"
	flex "libreflex/flexservice"
	"libreflex/model"
	"libreflex/repos"
	"log"
)

func main() {
	cfg := readConfig()

	serversStarted := 0

	if cfg.FlexApiServerConfig.Enabled {
		startFlexApiServer(cfg.FlexApiServerConfig)
		serversStarted++
		log.Println("FlexApi server started")
	}

	if cfg.FlexAdsServerConfig.Enabled {
		startFlexAdsServer(cfg.FlexAdsServerConfig)
		serversStarted++
		log.Println("FlexAds server started")
	}

	if serversStarted == 0 {
		log.Fatal("Nothing to start, check your config and make sure at least one server is enabled")
		return
	}

	select {}
}

func startFlexApiServer(config *model.FlexApiServerConfig) {
	adsRepo := getAdsRepo(config.AdsRepoProvider, config.AdsRepoConfig)

	flexConfig := &flex.LibreFlexConfig{
		ApiServerPort: config.Port,
		AdsApiUrl:     config.AdsApiUrl,
		DonateUrl:     config.DonateUrl,
		VkMusicRepo:   repos.NewDefaultVkMusicRepo(config.VkApiToken),
		AdsRepo:       adsRepo}

	go func() {
		err := flex.NewLibreFlexApiService(flexConfig).Start()
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func startFlexAdsServer(config *model.FlexAdsServerConfig) {
	adsRepo := getAdsRepo(config.AdsRepoProvider, config.AdsRepoConfig)

	adsConfig := &ads.LibreFlexAdsServiceConfig{
		ApiServerPort: config.Port,
		AdsRepo:       adsRepo}

	go func() {
		err := ads.NewLibreFlexAdsService(adsConfig).Start()
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func readConfig() *model.LibreFlexConfig {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
		return nil
	}

	config := &model.LibreFlexConfig{}
	err = json.Unmarshal(data, config)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return config
}

func getAdsRepo(adsRepoProvider string, adsRepoConfig map[string]interface{}) repos.AdsRepo {
	switch adsRepoProvider {
	case "dummy":
		return repos.NewDummyAdsRepo(adsRepoConfig)
	case "postgres":
		adsRepo, err := repos.NewPostgresAdsRepo(adsRepoConfig)
		if err != nil {
			log.Fatal(err)
		}
		return adsRepo
	default:
		log.Fatal("Unknown adsRepoProvider: " + adsRepoProvider)
	}

	return nil
}
