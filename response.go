package service

import (
	"encoding/json"
	"io"
	"net/http"
)

// Response interface is use to handle a service response.
type Response interface {
	SetStatus(status string)
}

// Writer is a func that can write a response to writer.
type Writer func(w io.Writer, r Response)

// R basic response for json.
type R struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Error  error       `json:"error,omitempty"`
}

// SetStatus implemented response interface.
func (r *R) SetStatus(status string) {
	r.Status = status
}

// JSON a response to a writer.
func JSON(w io.Writer, r Response) {
	bts, err := json.Marshal(r)
	if err != nil && log != nil {
		log.Error(err)
	}
	io.WriteString(w, string(bts))
}

// OK a response to a writer.
func OK(w io.Writer, r Response, wf Writer) {
	r.SetStatus("ok")
	if hw, ok := w.(http.ResponseWriter); ok {
		hw.WriteHeader(http.StatusOK)
		wf(hw, r)
		return
	}
	wf(w, r)
}

// Error set error status to response to a writer.
func Error(w io.Writer, r Response, wf Writer) {
	r.SetStatus("error")
	if hw, ok := w.(http.ResponseWriter); ok {
		hw.WriteHeader(http.StatusInternalServerError)
		wf(hw, r)
		return
	}
	wf(w, r)
}

// Deny write the deny response.
func Deny(w io.Writer, r Response, wf Writer) {
	r.SetStatus("unauthorized")
	if hw, ok := w.(http.ResponseWriter); ok {
		hw.WriteHeader(http.StatusUnauthorized)
		wf(hw, r)
		return
	}
	wf(w, r)
}

// Forbid write the forbidden response.
func Forbid(w io.Writer, r Response, wf Writer) {
	r.SetStatus("forbidden")
	if hw, ok := w.(http.ResponseWriter); ok {
		hw.WriteHeader(http.StatusForbidden)
		wf(hw, r)
		return
	}
	wf(w, r)
}
