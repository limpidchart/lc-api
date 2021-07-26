package testutils

import (
	"bytes"
	"encoding/json"
	"testing"
)

// EncodeToJSON encodes the provided `v` into JSON string.
func EncodeToJSON(t *testing.T, v interface{}) string {
	t.Helper()

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)

	if err := enc.Encode(v); err != nil {
		t.Fatalf("unable to endode JSON: %s", err)
	}

	return buf.String()
}
