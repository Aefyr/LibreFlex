package flexservice

import (
	"libreflex/repos"
	"net/http"
)

//Dependecy injection LULW
type LibreFlexConfig struct {
	ApiServerPort int
	AdsApiUrl     string
	DonateUrl     string
	VkMusicRepo   repos.VkMusicRepo
	AdsRepo       repos.AdsRepo
}

//LibreFlex FlexAPI service
type LibreFlexApiService struct {
	cfg *LibreFlexConfig
	mux *http.ServeMux
}

func NewLibreFlexApiService(config *LibreFlexConfig) *LibreFlexApiService {
	return &LibreFlexApiService{cfg: config}
}

func (s *LibreFlexApiService) Start() error {
	s.registerEndpoints()

	err := s.startApi()
	if err != nil {
		return err
	}

	return nil
}
