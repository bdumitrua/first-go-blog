package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Request struct {
	w http.ResponseWriter
	r *http.Request
}

func MakeRequest(w http.ResponseWriter, r *http.Request) *Request {
	return &Request{w: w, r: r}
}

func (req *Request) Validate() (map[string]interface{}, error) {
	body, err := io.ReadAll(req.r.Body)
	if err != nil {
		http.Error(req.w, "Unable to read request body", http.StatusBadRequest)
		return map[string]interface{}{}, errors.New("unable to read request body")
	}

	if !json.Valid(body) {
		http.Error(req.w, "Invalid JSON format", http.StatusBadRequest)
		return map[string]interface{}{}, errors.New("invalid JSON format")
	}

	var rawData map[string]interface{}
	err = json.Unmarshal(body, &rawData)
	if err != nil {
		http.Error(req.w, "Unable to parse JSON", http.StatusBadRequest)
		return map[string]interface{}{}, errors.New("unable to parse JSON")
	}

	return rawData, nil
}

func (req *Request) Writer() http.ResponseWriter {
	return req.w
}

func (req *Request) GetRequest() *http.Request {
	return req.r
}
