package adsservice

import (
	"fmt"
	"log"
	"net/http"
)

func (s *LibreFlexAdsService) registerEndpoints() {
	mux := http.NewServeMux()
	s.mux = mux

	mux.HandleFunc("/FlexAds", s.ApiCallMethod)
}

func (s *LibreFlexAdsService) startApi() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.ApiServerPort), s.mux)
}

func (s *LibreFlexAdsService) logApiRequest(method string, req *http.Request) {
	log.Println(fmt.Sprintf("Got FlexAdsAPI request with method %s. URL: %s", method, req.URL.String()))
}

func (s *LibreFlexAdsService) logApiError(method, step string, err error) {
	log.Println(fmt.Sprintf("FlexAdsAPI error while handling method %s at step \"%s\": %s", method, step, err.Error()))
}
