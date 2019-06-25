package responses

import "net/http"

type response interface {
	code() int
	body() []byte
}

func WriteResponse(response response, w http.ResponseWriter) error {
	w.WriteHeader(response.code())
	_, err := w.Write(response.body())
	return err
}
