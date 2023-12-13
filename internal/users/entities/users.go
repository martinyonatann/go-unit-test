package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/martinyonatann/go-unit-test/internal/users/dtos"
)

type Users struct {
	ID        string     `db:"id"`
	Name      string     `db:"name"`
	Password  string     `db:"password"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

func NewUsers(request dtos.Users) Users {
	return Users{
		ID:        uuid.NewString(),
		Name:      request.Name,
		Password:  request.Password,
		CreatedAt: time.Now(),
	}
}

func NewUsersDTO(request Users) dtos.Users {
	return dtos.Users{
		ID:        request.ID,
		Name:      request.Name,
		Password:  request.Password,
		CreatedAt: request.CreatedAt,
		UpdatedAt: request.UpdatedAt,
	}
}
