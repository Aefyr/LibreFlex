package flexservice

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	resps "libreflex/responses"
	"net/http"
	"strings"
)

func (s *LibreFlexApiService) ApiCallMethod(w http.ResponseWriter, r *http.Request) {
	switch strings.ToLower(r.FormValue("method")) {
	case "getinfo":
		s.ApiGetInfo(w, r)
	case "getaudio":
		s.ApiGetAudio(w, r)
	default:
		s.ApiUnknownMethod(w, r)
	}
}

func (s *LibreFlexApiService) ApiGetInfo(w http.ResponseWriter, r *http.Request) {
	s.logApiRequest("GetInfo", r)

	resp := resps.NewGetInfoResponse(s.cfg.DonateUrl, s.cfg.AdsApiUrl)
	resps.WriteResponse(resp, w)
}

func (s *LibreFlexApiService) ApiGetAudio(w http.ResponseWriter, r *http.Request) {
	s.logApiRequest("GetAudio", r)

	audioIds := strings.Split(r.FormValue("audios"), ",")

	vkResp, err := s.cfg.VkMusicRepo.GetAudio(audioIds)
	if err != nil {
		s.logApiError("GetAudio", "VkMusicRepo.GetAudio", err)
		resps.WriteResponse(resps.NewInternalServerErrorResponse(), w)
		return
	}

	ad, err := s.cfg.AdsRepo.GetRandomAd()
	if err != nil {
		s.logApiError("GetAudio", "AdsRepo.GetRandomAd", err)
		resps.WriteResponse(resps.NewInternalServerErrorResponse(), w)
		return
	}

	vkResp["flexad"] = ad
	respData, err := json.Marshal(vkResp)
	if err != nil {
		s.logApiError("GetAudio", "Marshall response", err)
		resps.WriteResponse(resps.NewInternalServerErrorResponse(), w)
		return
	}

	resp := make([]byte, base64.StdEncoding.EncodedLen(len(respData)))
	base64.StdEncoding.Encode(resp, respData)

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (s *LibreFlexApiService) ApiUnknownMethod(w http.ResponseWriter, r *http.Request) {
	s.logApiRequest("UnknownMethod", r)

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "Uknown method passed")
}
