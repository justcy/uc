package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(NotificatorService) NotificatorService

type loggingMiddleware struct {
	logger log.Logger
	next   NotificatorService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a NotificatorService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next NotificatorService) NotificatorService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) SendSms(ctx context.Context, mobile string, type1 int) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "SendSms", "mobile", mobile, "type1", type1, "rs", rs, "err", err)
	}()
	return l.next.SendSms(ctx, mobile, type1)
}
func (l loggingMiddleware) CheckSms(ctx context.Context, mobile string, code string) (rs bool, err error) {
	defer func() {
		l.logger.Log("method", "CheckSms", "mobile", mobile, "code", code, "rs", rs, "err", err)
	}()
	return l.next.CheckSms(ctx, mobile, code)
}
func (l loggingMiddleware) SendEmail(ctx context.Context, title string, content string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "SendEmail", "title", title, "content", content, "rs", rs, "err", err)
	}()
	return l.next.SendEmail(ctx, title, content)
}
