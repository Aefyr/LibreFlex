package responses

import "net/http"

type GetInfoResponse struct {
	DonateUrl string `json:"donate_url"`
	AdsApiUrl string `json:"ads_url"`
}

func (r *GetInfoResponse) code() int {
	return http.StatusOK
}

func (r *GetInfoResponse) body() []byte {
	return jsonEncode(r)
}

func NewGetInfoResponse(donateUrl string, adsApiUrl string) *GetInfoResponse {
	return &GetInfoResponse{DonateUrl: donateUrl, AdsApiUrl: adsApiUrl}
}
