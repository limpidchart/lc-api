package middleware

import (
	"encoding/json"
	"net/http"
)

const (
	contentTypeHeader = "Content-Type"
	jsonContentType   = "application/json"
)

// MarshalJSON marshals 'v' to JSON, automatically escaping HTML, setting the Content-Type as application/json
// and writing headers with the provided status code.
// It calls http.Error in case of failures.
func MarshalJSON(w http.ResponseWriter, statusCode int, v interface{}) {
	w.Header().Set(contentTypeHeader, jsonContentType)
	w.WriteHeader(statusCode)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)

	if err := enc.Encode(v); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
