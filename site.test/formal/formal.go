package formal

import (
	"context"

	"site.test/user"
)

type Users interface {
	Create(ctx context.Context, u *user.User) (created *user.User, err error)
	FindAll(ctx context.Context) (loaded []*user.User, err error)
}

type Service struct {
	users Users
}

func NewService(users Users) *Service {
	return &Service{users: users}
}

func (service *Service) Create(ctx context.Context, u *user.User) (*user.User, error) {
	u, err := service.users.Create(ctx, u)
	if err != nil {
		return nil, err
	}

	u.Name = "Mr. " + u.Name
	return u, nil
}

func (service *Service) Get(ctx context.Context) ([]*user.User, error) {
	users, err := service.users.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, u := range users {
		u.Name = "Mr. " + u.Name
	}
	return users, nil
}
