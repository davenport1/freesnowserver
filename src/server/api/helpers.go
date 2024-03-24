package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type envelope map[string]any

// writeJSON serializes the data to json and writes to the response writer the result
// Parameters: http.ResponseWriter, int, any
// Returns nil if writer writes the json object to response writer, otherwise error
func writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	jsonObj, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	jsonObj = append(jsonObj, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if _, err := w.Write(jsonObj); err != nil {
		// log.
	}

	return nil
}

// readJSON deserializes the data from json
// returns nil if JSON is successfully deserialized
func readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// disallow unknown fields in requests
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		// Custom error handling
		return err
	}

	if err := dec.Decode(&struct{}{}); err != io.EOF {
		return errors.New("body must only contain a single JSON object")
	}

	return nil
}
