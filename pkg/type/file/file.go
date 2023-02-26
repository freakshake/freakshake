package file

import (
	"context"
	"io"
	"time"

	"github.com/freakshake/pkg/type/optional"
	"github.com/google/uuid"
)

type ID string

func NewFileID() ID {
	return ID(uuid.NewString())
}

type Info struct {
	ID        ID                           `json:"id"`
	Name      string                       `json:"name"`
	Type      string                       `json:"type"`
	Size      ByteSize                     `json:"size"`
	CreatedAt time.Time                    `json:"created_at"`
	DeletedAt optional.Optional[time.Time] `json:"deleted_at"`
}

type Service interface {
	Upload(context.Context, Info, io.Reader) (ID, error)
	Get(context.Context, ID) (io.ReadCloser, error)
	GetFileInfo(context.Context, ID) (Info, error)
	Delete(context.Context, ID) error
}

type Storage interface {
	Store(context.Context, Info) (ID, error)
	Find(context.Context, ID) (Info, error)
	Delete(context.Context, ID) error
}
