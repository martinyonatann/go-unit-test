package dtos

import (
	"time"

	"github.com/invopop/validation"
)

type Users struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Password  string     `json:"password"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type UserDetailRequest struct {
	ID string `param:"id"`
}

func (x *UserDetailRequest) Validate() error {
	return validation.ValidateStruct(&x,
		validation.Field(&x.ID, validation.Required, validation.NotNil),
	)
}
