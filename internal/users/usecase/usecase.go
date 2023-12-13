package usecase

import (
	"context"

	"github.com/martinyonatann/go-unit-test/internal/users"
	"github.com/martinyonatann/go-unit-test/internal/users/dtos"
	"github.com/martinyonatann/go-unit-test/internal/users/entities"
)

type usecase struct {
	repo users.Repositories
}

func NewUseCase(r users.Repositories) users.UseCases {
	return &usecase{repo: r}
}

func (uc *usecase) CreateUser(ctx context.Context, request dtos.Users) error {
	err := uc.repo.Create(ctx, entities.NewUsers(request))
	if err != nil {
		return err
	}

	return nil
}

func (uc *usecase) GetUserDetail(ctx context.Context, userUUID string) (userDetail dtos.Users, err error) {
	userData, err := uc.repo.Detail(ctx, userUUID)
	if err != nil {
		return userDetail, err
	}

	return entities.NewUsersDTO(userData), nil
}
