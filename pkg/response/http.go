package response

import (
	"encoding/json"
	"net/http"
)

type httpResponse struct {
	Data any `json:"data"`
}

func EncodeHTTP(w http.ResponseWriter, res any) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(httpResponse{Data: res})
}
