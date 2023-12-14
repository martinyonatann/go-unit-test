//go:build unit
// +build unit

package usecase_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/martinyonatann/go-unit-test/internal/generated/mocks"
	"github.com/martinyonatann/go-unit-test/internal/users/dtos"
	"github.com/martinyonatann/go-unit-test/internal/users/entities"
	"github.com/martinyonatann/go-unit-test/internal/users/usecase"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type eqUsersMatcher struct {
	users entities.Users
}

func createUsersMatcher(u entities.Users) gomock.Matcher {
	return eqUsersMatcher{
		users: u,
	}
}

func (e eqUsersMatcher) Matches(x interface{}) bool {
	arg, ok := x.(entities.Users)
	if !ok {
		return false
	}

	return arg.Name == e.users.Name && arg.Password == e.users.Password
}

func (e eqUsersMatcher) String() string {
	return fmt.Sprintf("%v", e.users.Name)
}

func Test_usecase_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		repositories = mocks.NewMockRepositories(ctrl)
		usecase      = usecase.NewUseCase(repositories)
	)

	var request = dtos.Users{
		Name:     faker.Name(),
		Password: faker.Password(),
	}

	t.Run("positive_createUser", func(t *testing.T) {
		repositories.EXPECT().Create(context.Background(), createUsersMatcher(entities.NewUsers(request))).Return(nil)
		err := usecase.CreateUser(context.Background(), request)
		require.NoError(t, err)
	})

	t.Run("negative_createUser_failed_insert", func(t *testing.T) {
		repositories.EXPECT().Create(context.Background(), createUsersMatcher(entities.NewUsers(request))).Return(sql.ErrConnDone)
		err := usecase.CreateUser(context.Background(), request)
		require.EqualError(t, err, sql.ErrConnDone.Error())
	})

}

func Test_usecase_GetUserDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		repositories = mocks.NewMockRepositories(ctrl)
		usecase      = usecase.NewUseCase(repositories)
		id           = uuid.NewString()
	)

	t.Run("positive_GetUserDetail", func(t *testing.T) {
		var users = entities.Users{
			ID:        id,
			Name:      faker.Name(),
			Password:  faker.Password(),
			CreatedAt: time.Now(),
		}

		repositories.EXPECT().Detail(context.Background(), id).Return(users, nil)

		detail, err := usecase.GetUserDetail(context.Background(), id)
		require.NoError(t, err)
		require.Equal(t, entities.NewUsersDTO(users), detail)
	})

	t.Run("negative_GetUserDetail_not_found", func(t *testing.T) {
		repositories.EXPECT().Detail(context.Background(), id).Return(entities.Users{}, sql.ErrNoRows)

		detail, err := usecase.GetUserDetail(context.Background(), id)
		require.EqualError(t, sql.ErrNoRows, err.Error())
		require.Equal(t, detail, dtos.Users{})
	})
}
