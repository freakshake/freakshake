package file

import (
	"context"
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/mehdieidi/storm/pkg/type/optional"
)

type FileID string

func NewFileID() FileID {
	return FileID(uuid.NewString())
}

type FileInfo struct {
	ID        FileID                       `json:"id"`
	Name      string                       `json:"name"`
	Type      string                       `json:"type"`
	Size      ByteSize                     `json:"size"`
	CreatedAt time.Time                    `json:"created_at"`
	DeletedAt optional.Optional[time.Time] `json:"deleted_at"`
}

type FileService interface {
	Upload(context.Context, FileInfo, io.Reader) (FileID, error)
	Get(context.Context, FileID) (io.ReadCloser, error)
	GetFileInfo(context.Context, FileID) (FileInfo, error)
	Delete(context.Context, FileID) error
}

type FileStorage interface {
	Store(context.Context, FileInfo) (FileID, error)
	Find(context.Context, FileID) (FileInfo, error)
	Delete(context.Context, FileID) error
}
