package service

import "context"

type User struct {
}

// UsersService describes the service.
type UsersService interface {
	// Add your methods here
	Login(ctx context.Context, username string, password string) (rs string, err error)
	Register(ctx context.Context, username string, password string) (user User, err error)
	Update(ctx context.Context, user User) (rs string, err error)
}

type basicUsersService struct{}

func (b *basicUsersService) Login(ctx context.Context, username string, password string) (rs string, err error) {
	// TODO implement the business logic of Login
	return rs, err
}
func (b *basicUsersService) Register(ctx context.Context, username string, password string) (user User, err error) {
	// TODO implement the business logic of Register
	return user, err
}
func (b *basicUsersService) Update(ctx context.Context, user User) (rs string, err error) {
	// TODO implement the business logic of Update
	return rs, err
}

// NewBasicUsersService returns a naive, stateless implementation of UsersService.
func NewBasicUsersService() UsersService {
	return &basicUsersService{}
}

// New returns a UsersService with all of the expected middleware wired in.
func New(middleware []Middleware) UsersService {
	var svc UsersService = NewBasicUsersService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
