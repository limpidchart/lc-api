package middleware

import (
	"encoding/json"
	"net/http"
)

const (
	contentTypeHeader = "Content-Type"
	contentTypeJSON   = "application/json; charset=utf-8"
)

// MarshalJSON marshals 'v' to JSON, automatically escaping HTML and setting the Content-Type as application/json.
// It will call http.Error in case of failures.
func MarshalJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set(contentTypeHeader, contentTypeJSON)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(true)

	if err := enc.Encode(v); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
