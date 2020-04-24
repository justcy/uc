package service

import "context"

// NotificatorService describes the service.
type NotificatorService interface {
	// Add your methods here
	// e.x: Foo(ctx context.Context,s string)(rs string, err error)
	SendSms(ctx context.Context, mobile string, type1 int) (rs string, err error)
	CheckSms(ctx context.Context, mobile string, code string) (rs bool, err error)
	SendEmail(ctx context.Context, title string, content string) (rs string, err error)
}

type basicNotificatorService struct{}

func (b *basicNotificatorService) SendSms(ctx context.Context, mobile string, type1 int) (rs string, err error) {
	// TODO implement the business logic of SendSms
	return rs, err
}
func (b *basicNotificatorService) CheckSms(ctx context.Context, mobile string, code string) (rs bool, err error) {
	// TODO implement the business logic of CheckSms
	return rs, err
}
func (b *basicNotificatorService) SendEmail(ctx context.Context, title string, content string) (rs string, err error) {
	// TODO implement the business logic of SendEmail
	return rs, err
}

// NewBasicNotificatorService returns a naive, stateless implementation of NotificatorService.
func NewBasicNotificatorService() NotificatorService {
	return &basicNotificatorService{}
}

// New returns a NotificatorService with all of the expected middleware wired in.
func New(middleware []Middleware) NotificatorService {
	var svc NotificatorService = NewBasicNotificatorService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
