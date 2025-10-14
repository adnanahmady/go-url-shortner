package request

import (
	"context"
	"fmt"

	"github.com/adnanahmady/go-url-shortner/pkg/applog"
)

var (
	loggerKey    = &struct{}{}
	requestIDKey = &struct{}{}
)

func GetLogger(ctx context.Context) applog.Logger {
	key := fmt.Sprintf("%v", &loggerKey)
	return ctx.Value(key).(applog.Logger)
}

func GetRequestID(ctx context.Context) string {
	key := fmt.Sprintf("%v", &requestIDKey)
	return ctx.Value(key).(string)
}

func SetLogger(ctx context.Context, logger applog.Logger) context.Context {
	key := fmt.Sprintf("%v", &loggerKey)
	return context.WithValue(ctx, key, logger)
}

func SetRequestID(ctx context.Context, requestID string) context.Context {
	key := fmt.Sprintf("%v", &requestIDKey)
	return context.WithValue(ctx, key, requestID)
}
