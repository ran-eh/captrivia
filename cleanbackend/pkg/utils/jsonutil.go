// pkg/utils/jsonutil.go
package utils

import (
	"encoding/json"
	"io"
)

// Marshal is a utility function that marshals an object into JSON.
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// Unmarshal is a utility function that unmarshals JSON into an object.
func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// DecodeJSONBody is a utility function to decode a JSON request body.
func DecodeJSONBody(r io.Reader, dest interface{}) error {
	return json.NewDecoder(r).Decode(dest)
}

// EncodeJSONResponse is a utility function to encode an object as a JSON response.
func EncodeJSONResponse(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}