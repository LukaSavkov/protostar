package utils

import (
	"encoding/json"
	"health-check/domain"
	"net/http"
	"time"
)

func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		WriteErrorResp(err.Error(), http.StatusInternalServerError, "Internal Server Error", w)
	}
}

func WriteErrorResp(err string, status int, path string, w http.ResponseWriter) {

	baseErrorResp := domain.BaseErrorHttpResponse{
		Error:  err,
		Path:   path,
		Status: status,
		Time:   time.Now().String(),
	}
	writeJSONResponse(w, status, baseErrorResp)
}

func WriteResp(resp any, statusCode int, w http.ResponseWriter) {
	if resp == nil {
		return
	}
	domainResponse := domain.BaseHttpResponse{
		Status: statusCode,
		Data:   resp,
	}
	writeJSONResponse(w, statusCode, domainResponse)
}

func DecodeJSONFromRequest(r *http.Request, rw http.ResponseWriter, v interface{}) bool {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&v); err != nil {
		WriteErrorResp("Internal server error", 500, r.URL.Path, rw)
		return false
	}
	return true
}
