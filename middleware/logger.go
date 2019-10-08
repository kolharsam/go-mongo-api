package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var logger = log.New()

func init() {
	logger.SetFormatter(&log.JSONFormatter{
		PrettyPrint: true,
	})
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	logger.SetReportCaller(true)
	logger.Out = os.Stdout
}

// RequestLogger := middleware logging info of all the requests
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.WithFields(log.Fields{
			"route":     r.RequestURI,
			"method":    r.Method,
			"body":      r.Body,
			"headers":   r.Header,
			"createdAt": time.Now(),
		}).Info(r.URL)
		next.ServeHTTP(w, r)
	})
}
