package endpoint

import (
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
	service "github.com/justcy/uc/users/pkg/service"
)

// LoginRequest collects the request parameters for the Login method.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse collects the response parameters for the Login method.
type LoginResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeLoginEndpoint returns an endpoint that invokes Login on the service.
func MakeLoginEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		rs, err := s.Login(ctx, req.Username, req.Password)
		return LoginResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r LoginResponse) Failed() error {
	return r.Err
}

// RegisterRequest collects the request parameters for the Register method.
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterResponse collects the response parameters for the Register method.
type RegisterResponse struct {
	User service.User `json:"user"`
	Err  error        `json:"err"`
}

// MakeRegisterEndpoint returns an endpoint that invokes Register on the service.
func MakeRegisterEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RegisterRequest)
		user, err := s.Register(ctx, req.Username, req.Password)
		return RegisterResponse{
			Err:  err,
			User: user,
		}, nil
	}
}

// Failed implements Failer.
func (r RegisterResponse) Failed() error {
	return r.Err
}

// UpdateRequest collects the request parameters for the Update method.
type UpdateRequest struct {
	User service.User `json:"user"`
}

// UpdateResponse collects the response parameters for the Update method.
type UpdateResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeUpdateEndpoint returns an endpoint that invokes Update on the service.
func MakeUpdateEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateRequest)
		rs, err := s.Update(ctx, req.User)
		return UpdateResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r UpdateResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Login implements Service. Primarily useful in a client.
func (e Endpoints) Login(ctx context.Context, username string, password string) (rs string, err error) {
	request := LoginRequest{
		Password: password,
		Username: username,
	}
	response, err := e.LoginEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(LoginResponse).Rs, response.(LoginResponse).Err
}

// Register implements Service. Primarily useful in a client.
func (e Endpoints) Register(ctx context.Context, username string, password string) (user service.User, err error) {
	request := RegisterRequest{
		Password: password,
		Username: username,
	}
	response, err := e.RegisterEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(RegisterResponse).User, response.(RegisterResponse).Err
}

// Update implements Service. Primarily useful in a client.
func (e Endpoints) Update(ctx context.Context, user service.User) (rs string, err error) {
	request := UpdateRequest{User: user}
	response, err := e.UpdateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(UpdateResponse).Rs, response.(UpdateResponse).Err
}
