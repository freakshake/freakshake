package domain

import (
	"context"

	"github.com/mehdieidi/storm/pkg/type/file"
	"github.com/mehdieidi/storm/pkg/type/offlim"
)

type UserID uint

type User struct {
	Credential
	CrUpDe

	ID     UserID       `json:"id"`
	Avatar *file.FileID `json:"avatar"`
}

type UserStorage interface {
	Store(context.Context, User) (UserID, error)
	Find(context.Context, UserID) (User, error)
	FindAll(context.Context, offlim.Offset, offlim.Limit) ([]User, error)
	Update(context.Context, UserID, User) error
	Delete(context.Context, UserID) error
}

type UserService interface {
	Create(context.Context, User) (User, error)
	Get(context.Context, UserID) (User, error)
	List(context.Context, offlim.Offset, offlim.Limit) ([]User, error)
	Update(context.Context, UserID, User) error
	Delete(context.Context, UserID) error
}
