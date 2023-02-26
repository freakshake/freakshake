package domain

import (
	"context"
	"time"

	"github.com/freakshake/pkg/type/email"
	"github.com/freakshake/pkg/type/file"
	"github.com/freakshake/pkg/type/id"
	"github.com/freakshake/pkg/type/mobile"
	"github.com/freakshake/pkg/type/offlim"
	"github.com/freakshake/pkg/type/optional"
	"github.com/freakshake/pkg/type/password"
)

const UserDomain = "user"

type User struct {
	ID             id.ID[User]                  `json:"id" swaggertype:"integer"`
	FirstName      string                       `json:"first_name"`
	LastName       string                       `json:"last_name"`
	Email          email.Email                  `json:"email"`
	MobileNumber   mobile.Number                `json:"mobile_number"`
	Password       password.Password            `json:"-"`
	HashedPassword string                       `json:"-"`
	Avatar         optional.Optional[file.ID]   `json:"avatar" swaggertype:"string"`
	CreatedAt      time.Time                    `json:"created_at"`
	UpdatedAt      optional.Optional[time.Time] `json:"updated_at" swaggertype:"string"`
	DeletedAt      optional.Optional[time.Time] `json:"deleted_at" swaggertype:"string"`
}

type UserStorage interface {
	Store(context.Context, User) (id.ID[User], error)
	Find(context.Context, id.ID[User]) (User, error)
	FindByEmail(context.Context, email.Email) (User, error)
	FindAll(context.Context, offlim.Offset, offlim.Limit) ([]User, error)
	Update(context.Context, id.ID[User], User) error
	Delete(context.Context, id.ID[User]) error
}

type UserService interface {
	Create(context.Context, User) (User, error)
	Get(context.Context, id.ID[User]) (User, error)
	GetByEmail(context.Context, email.Email) (User, error)
	List(context.Context, offlim.Offset, offlim.Limit) ([]User, error)
	Update(context.Context, id.ID[User], User) error
	Delete(context.Context, id.ID[User]) error
}
