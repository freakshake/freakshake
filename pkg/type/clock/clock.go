package clock

import (
	"database/sql/driver"
	"time"
)

type Clock time.Time

func (c Clock) Value() (driver.Value, error) {
	return time.Time(c), nil
}

func (c Clock) String() string {
	return time.Time(c).Format(time.Kitchen)
}

func (c Clock) MarshalJSON() ([]byte, error) {
	return []byte(`"` + c.String() + `"`), nil
}

func (c *Clock) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return ErrInvalidClock
	}

	t, err := time.Parse(`"`+time.Kitchen+`"`, string(data))
	if err != nil {
		return err
	}

	*c = Clock(t)

	return nil
}
