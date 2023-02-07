package nulltime

import (
	"database/sql/driver"
	"time"
	_ "unsafe"
)

type NullTime struct {
	Time  time.Time
	Valid bool
}

func (n *NullTime) Scan(value any) error {
	if value == nil {
		n.Time, n.Valid = time.Time{}, false
		return nil
	}

	n.Valid = true

	return convertAssign(&n.Time, value)
}

func (n NullTime) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}

func (n NullTime) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte("null"), nil
	}
	return n.Time.MarshalJSON()
}

func (n *NullTime) UnmarshalJSON(text []byte) error {
	if string(text) == "null" {
		n.Valid = false
		return nil
	}

	n.Valid = true

	return n.Time.UnmarshalJSON(text)
}

//go:linkname convertAssign database/sql.convertAssign
func convertAssign(dest, src any) error
