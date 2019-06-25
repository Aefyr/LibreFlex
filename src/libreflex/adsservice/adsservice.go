package adsservice

import (
	"libreflex/repos"
	"net/http"
)

type LibreFlexAdsServiceConfig struct {
	ApiServerPort int
	AdsRepo       repos.AdsRepo
}

type LibreFlexAdsService struct {
	cfg *LibreFlexAdsServiceConfig
	mux *http.ServeMux
}

func NewLibreFlexAdsService(config *LibreFlexAdsServiceConfig) *LibreFlexAdsService {
	return &LibreFlexAdsService{cfg: config}
}

func (s *LibreFlexAdsService) Start() error {
	s.registerEndpoints()

	err := s.startApi()
	if err != nil {
		return err
	}

	return nil
}
