package domain

import (
	"context"
	"time"

	"github.com/mehdieidi/storm/pkg/type/email"
	"github.com/mehdieidi/storm/pkg/type/file"
	"github.com/mehdieidi/storm/pkg/type/id"
	"github.com/mehdieidi/storm/pkg/type/mobile"
	"github.com/mehdieidi/storm/pkg/type/offlim"
	"github.com/mehdieidi/storm/pkg/type/optional"
	"github.com/mehdieidi/storm/pkg/type/password"
)

type User struct {
	ID             id.ID[User]                    `json:"id" swaggertype:"integer"`
	FirstName      string                         `json:"first_name"`
	LastName       string                         `json:"last_name"`
	Email          email.Email                    `json:"email"`
	MobileNumber   mobile.MobileNumber            `json:"mobile_number"`
	Password       password.Password              `json:"-"`
	HashedPassword string                         `json:"-"`
	Avatar         optional.Optional[file.FileID] `json:"avatar" swaggertype:"string"`
	CreatedAt      time.Time                      `json:"created_at"`
	UpdatedAt      optional.Optional[time.Time]   `json:"updated_at" swaggertype:"string"`
	DeletedAt      optional.Optional[time.Time]   `json:"deleted_at" swaggertype:"string"`
}

type UserStorage interface {
	Store(context.Context, User) (id.ID[User], error)
	Find(context.Context, id.ID[User]) (User, error)
	FindAll(context.Context, offlim.Offset, offlim.Limit) ([]User, error)
	Update(context.Context, id.ID[User], User) error
	Delete(context.Context, id.ID[User]) error
}

type UserService interface {
	Create(context.Context, User) (User, error)
	Get(context.Context, id.ID[User]) (User, error)
	List(context.Context, offlim.Offset, offlim.Limit) ([]User, error)
	Update(context.Context, id.ID[User], User) error
	Delete(context.Context, id.ID[User]) error
}
