package derror

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func EncodeHTTPError(w http.ResponseWriter, err error) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	status, ok := errHTTPStatusMap[err]
	if !ok {
		status = http.StatusInternalServerError
	}
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(errorResponse{Error: err.Error()})

	return nil
}
