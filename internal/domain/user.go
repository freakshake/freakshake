package domain

import (
	"context"

	"github.com/mehdieidi/storm/pkg/type/file"
	"github.com/mehdieidi/storm/pkg/type/id"
	"github.com/mehdieidi/storm/pkg/type/offlim"
)

type User struct {
	Credential
	CrUpDe

	ID     id.ID[User]  `json:"id"`
	Avatar *file.FileID `json:"avatar"`
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
