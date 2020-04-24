package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/justcy/uc/notificator/pkg/service"
)

// SendSmsRequest collects the request parameters for the SendSms method.
type SendSmsRequest struct {
	Mobile string `json:"mobile"`
	Type1  int    `json:"type1"`
}

// SendSmsResponse collects the response parameters for the SendSms method.
type SendSmsResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeSendSmsEndpoint returns an endpoint that invokes SendSms on the service.
func MakeSendSmsEndpoint(s service.NotificatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendSmsRequest)
		rs, err := s.SendSms(ctx, req.Mobile, req.Type1)
		return SendSmsResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r SendSmsResponse) Failed() error {
	return r.Err
}

// CheckSmsRequest collects the request parameters for the CheckSms method.
type CheckSmsRequest struct {
	Mobile string `json:"mobile"`
	Code   string `json:"code"`
}

// CheckSmsResponse collects the response parameters for the CheckSms method.
type CheckSmsResponse struct {
	Rs  bool  `json:"rs"`
	Err error `json:"err"`
}

// MakeCheckSmsEndpoint returns an endpoint that invokes CheckSms on the service.
func MakeCheckSmsEndpoint(s service.NotificatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CheckSmsRequest)
		rs, err := s.CheckSms(ctx, req.Mobile, req.Code)
		return CheckSmsResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r CheckSmsResponse) Failed() error {
	return r.Err
}

// SendEmailRequest collects the request parameters for the SendEmail method.
type SendEmailRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// SendEmailResponse collects the response parameters for the SendEmail method.
type SendEmailResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeSendEmailEndpoint returns an endpoint that invokes SendEmail on the service.
func MakeSendEmailEndpoint(s service.NotificatorService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendEmailRequest)
		rs, err := s.SendEmail(ctx, req.Title, req.Content)
		return SendEmailResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r SendEmailResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// SendSms implements Service. Primarily useful in a client.
func (e Endpoints) SendSms(ctx context.Context, mobile string, type1 int) (rs string, err error) {
	request := SendSmsRequest{
		Mobile: mobile,
		Type1:  type1,
	}
	response, err := e.SendSmsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendSmsResponse).Rs, response.(SendSmsResponse).Err
}

// CheckSms implements Service. Primarily useful in a client.
func (e Endpoints) CheckSms(ctx context.Context, mobile string, code string) (rs bool, err error) {
	request := CheckSmsRequest{
		Code:   code,
		Mobile: mobile,
	}
	response, err := e.CheckSmsEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CheckSmsResponse).Rs, response.(CheckSmsResponse).Err
}

// SendEmail implements Service. Primarily useful in a client.
func (e Endpoints) SendEmail(ctx context.Context, title string, content string) (rs string, err error) {
	request := SendEmailRequest{
		Content: content,
		Title:   title,
	}
	response, err := e.SendEmailEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(SendEmailResponse).Rs, response.(SendEmailResponse).Err
}
