package service

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// SetLog on this package
func SetLog(extlog *logrus.Logger) {
	log = extlog
}

type logHandler struct {
	handler http.Handler
}

func (l *logHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	// Init the responseWriter to the statusOK
	wl := &responseWriterLog{w, http.StatusOK}
	l.handler.ServeHTTP(wl, r)
	if log != nil {
		entry := log.WithFields(logrus.Fields{
			"Method": r.Method,
			"Path":   r.URL,
			"Status": wl.statusCode,
			"Time":   time.Since(t),
		})
		switch wl.statusCode {
		case http.StatusOK:
			entry.Info("ok")
		case http.StatusInternalServerError:
			entry.Error("error")
		case http.StatusUnauthorized:
			entry.Error("unauthorized")
		case http.StatusForbidden:
			entry.Error("forbidden")
		default:
			entry.Info("request")
		}
	}
}

type responseWriterLog struct {
	http.ResponseWriter
	statusCode int
}

func (r *responseWriterLog) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}
