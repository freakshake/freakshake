package date

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Date time.Time

func (d Date) Value() (driver.Value, error) {
	t := time.Time(d)
	return fmt.Sprintf("%04d-%02d-%02d", t.Year(), t.Month(), t.Day()), nil
}

func (d *Date) Scan(v interface{}) error {
	switch s := v.(type) {
	case string:
		t, err := time.Parse("2006-01-02", s)
		if err != nil {
			return err
		}
		*d = Date(t)
	case time.Time:
		*d = Date(s)
	}

	return nil
}

func (b Date) MarshalText() ([]byte, error) {
	t := time.Time(b)
	return []byte(fmt.Sprintf("%04d-%02d-%02d", t.Year(), t.Month(), t.Day())), nil
}

func (d *Date) UnmarshalText(p []byte) error {
	s := string(p)
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}

	*d = Date(t)

	return nil
}
