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
		log.WithFields(logrus.Fields{
			"Method": r.Method,
			"Path":   r.URL,
			"Status": wl.statusCode,
			"Time":   time.Since(t),
		}).Info("Request")
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
