package request

import (
	"context"

	"github.com/adnanahmady/go-url-shortner/pkg/applog"
)

var (
	loggerKey    = &struct{ uint8 }{}
	requestIDKey = &struct{ uint8 }{}
)

func GetLogger(ctx context.Context) applog.Logger {
	return ctx.Value(loggerKey).(applog.Logger)
}

func SetLogger(ctx context.Context, logger applog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func GetRequestID(ctx context.Context) string {
	return ctx.Value(requestIDKey).(string)
}

func SetRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, requestIDKey, requestID)
}
