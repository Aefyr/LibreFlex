package responses

import (
	"libreflex/model"
	"net/http"
)

type GenerateAdResponse struct {
	Ok      bool          `json:"success"`
	Message string        `json:"message"`
	Ad      *model.FlexAd `json:"generated,omitempty"`
}

func (r *GenerateAdResponse) code() int {
	if r.Ok {
		return http.StatusOK
	}
	return http.StatusInternalServerError
}

func (r *GenerateAdResponse) body() []byte {
	return b64encode(jsonEncode(r))
}

func NewGenerateAdResponse(ad *model.FlexAd) *GenerateAdResponse {
	return &GenerateAdResponse{Ok: true, Message: "asdf", Ad: ad}
}

func NewGenerateAdResponseWithError() *GenerateAdResponse {
	return &GenerateAdResponse{Ok: false, Message: "internal server error"}
}
