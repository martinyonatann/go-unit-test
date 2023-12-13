package users

import (
	"context"

	"github.com/martinyonatann/go-unit-test/internal/users/dtos"
)

type UseCases interface {
	CreateUser(ctx context.Context, request dtos.Users) error
	GetUserDetail(ctx context.Context, userUUID string) (userDetail dtos.Users, err error)
}
