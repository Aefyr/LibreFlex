package responses

import "net/http"

type InternalServerErrorResponse struct {
}

func (r *InternalServerErrorResponse) code() int {
	return http.StatusInternalServerError
}

func (r *InternalServerErrorResponse) body() []byte {
	return []byte("internal server error")
}

func NewInternalServerErrorResponse() *InternalServerErrorResponse {
	return &InternalServerErrorResponse{}
}
