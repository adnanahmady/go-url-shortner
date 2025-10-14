package request

import (
	"net/http"

	"github.com/adnanahmady/go-url-shortner/pkg/applog"
	"github.com/google/uuid"
)

type LoggingMiddleware struct {
	logger applog.Logger
}

func NewLoggingMiddleware(logger applog.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{logger: logger}
}

func (l *LoggingMiddleware) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rLogger := l.logger.With(applog.Arg{
			Key:   "request_id",
			Value: uuid.New().String(),
		})
		rLogger.Info("Logging middleware")

		ctx := SetLogger(r.Context(), rLogger)
		ctx = SetRequestID(ctx, uuid.New().String())
		r = r.WithContext(ctx)

		rLogger.Infof("Request Method (%v)", r.Method)
		rLogger.Infof("Request URL (%v)", r.URL)
		rLogger.Infof("Request Host (%v)", r.Host)
		rLogger.Infof("Request User Agent (%v)", r.UserAgent())
		rLogger.Infof("Request Referer (%v)", r.Referer())
		rLogger.Infof("Request Remote Addr (%v)", r.RemoteAddr)
		rLogger.Infof("Request Proto (%v)", r.Proto)

		writer := writerWrapper{
			ResponseWriter: w,
			logger:         rLogger,
		}
		next(&writer, r)

		rLogger.Infof("Request completed")
	}
}

// writerWrapper is a wrapper around the http.ResponseWriter that logs the response status and body
type writerWrapper struct {
	http.ResponseWriter
	logger     applog.Logger
}

func (w *writerWrapper) WriteHeader(statusCode int) {
	w.logger.Infof("Response Status (%v)", statusCode)
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *writerWrapper) Write(b []byte) (int, error) {
	w.logger.Infof("Response Body (%v)", string(b))
	return w.ResponseWriter.Write(b)
}