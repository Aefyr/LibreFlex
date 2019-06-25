package flexservice

import (
	"fmt"
	"log"
	"net/http"
)

func (s *LibreFlexApiService) registerEndpoints() {
	mux := http.NewServeMux()
	s.mux = mux

	mux.HandleFunc("/FlexMusic", s.ApiCallMethod)
}

func (s *LibreFlexApiService) startApi() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.ApiServerPort), s.mux)
}

func (s *LibreFlexApiService) logApiRequest(method string, req *http.Request) {
	log.Println(fmt.Sprintf("Got FlexAPI request with method %s. URL: %s", method, req.URL.String()))
}

func (s *LibreFlexApiService) logApiError(method, step string, err error) {
	log.Println(fmt.Sprintf("FlexAPI error while handling method %s at step \"%s\": %s", method, step, err.Error()))
}
