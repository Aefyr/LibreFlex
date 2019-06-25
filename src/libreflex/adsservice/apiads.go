package adsservice

import (
	"encoding/base64"
	"fmt"
	resps "libreflex/responses"
	"net/http"
)

func (s *LibreFlexAdsService) ApiCallMethod(w http.ResponseWriter, r *http.Request) {
	switch r.FormValue("action") {
	case base64enc("GenerateAd"):
		s.ApiGenerateAd(w, r)
	default:
		s.ApiUnknownMethod(w, r)
	}
}

func (s *LibreFlexAdsService) ApiGenerateAd(w http.ResponseWriter, r *http.Request) {
	s.logApiRequest("GenerateAd", r)

	ad, err := s.cfg.AdsRepo.GetRandomAd()
	if err != nil {
		s.logApiError("GenerateAd", "AdsRepo.GetRandomAd", err)
		resps.WriteResponse(resps.NewGenerateAdResponseWithError(), w)
		return
	}

	resps.WriteResponse(resps.NewGenerateAdResponse(ad), w)
}

func (s *LibreFlexAdsService) ApiUnknownMethod(w http.ResponseWriter, r *http.Request) {
	s.logApiRequest("UnknownMethod", r)

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Uknown method passed")
}

func base64enc(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
