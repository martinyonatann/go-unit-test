package users

import (
	"context"

	"github.com/martinyonatann/go-unit-test/internal/users/entities"
)

type Repositories interface {
	Create(ctx context.Context, request entities.Users) error
	Detail(ctx context.Context, userUUID string) (response entities.Users, err error)
}
