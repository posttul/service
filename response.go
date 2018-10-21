package service

import (
	"encoding/json"
	"io"
	"net/http"
)

// Response interface is use to handle a service response.
type Response interface {
	OK(w io.Writer)
	Error(w io.Writer)
	Deny(w io.Writer)
}

// R basic response for json
type R struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Err    error       `json:"error,omitempty"`
}

func (r R) json(w io.Writer) {
	bts, err := json.Marshal(r)
	if err != nil && log != nil {
		log.Error(err)
	}
	io.WriteString(w, string(bts))
}

// OK write the ok response
func (r R) OK(w io.Writer) {
	r.Status = "ok"
	if hw, ok := w.(http.ResponseWriter); ok {
		hw.WriteHeader(http.StatusOK)
		r.json(hw)
		return
	}
	r.json(w)
}

// Error write the error response
func (r R) Error(w io.Writer) {
	r.Status = "error"
	if hw, ok := w.(http.ResponseWriter); ok {
		hw.WriteHeader(http.StatusInternalServerError)
		r.json(hw)
		return
	}
	r.json(w)
}

// Deny write the deny response
func (r R) Deny(w io.Writer) {
	r.Status = "deny"
	if hw, ok := w.(http.ResponseWriter); ok {
		hw.WriteHeader(http.StatusUnauthorized)
		r.json(hw)
		return
	}
	r.json(w)
}
