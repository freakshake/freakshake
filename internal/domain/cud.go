package domain

import (
	"time"

	"github.com/mehdieidi/storm/pkg/type/optional"
)

// CrUpDe represents the modification date of a record. i.e. created_at, updated_at, and deleted_at.
type CrUpDe struct {
	CreatedAt time.Time                    `json:"created_at"`
	UpdatedAt optional.Optional[time.Time] `json:"updated_at"`
	DeletedAt optional.Optional[time.Time] `json:"deleted_at"`
}
