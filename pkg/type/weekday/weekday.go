package weekday

import (
	"time"
)

type Weekday int

const (
	Unset Weekday = iota
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func NewWeekday(weekday string) (Weekday, error) {
	if weekday == time.Friday.String() {
		return Friday, nil
	}
	if weekday == time.Saturday.String() {
		return Saturday, nil
	}
	if weekday == time.Sunday.String() {
		return Sunday, nil
	}
	if weekday == time.Monday.String() {
		return Monday, nil
	}
	if weekday == time.Tuesday.String() {
		return Tuesday, nil
	}
	if weekday == time.Wednesday.String() {
		return Wednesday, nil
	}
	if weekday == time.Thursday.String() {
		return Thursday, nil
	}

	return Sunday, ErrUnknownWeekday
}

func (w Weekday) String() string {
	switch w {
	case Unset:
		return "unset"
	case Sunday:
		return "Sun"
	case Monday:
		return "Mon"
	case Tuesday:
		return "Tue"
	case Wednesday:
		return "Wed"
	case Thursday:
		return "Thu"
	case Friday:
		return "Fri"
	case Saturday:
		return "Sat"
	default:
		return ""
	}
}

func (w Weekday) MarshalText() ([]byte, error) {
	if s := w.String(); s != "" {
		return []byte(s), nil
	}
	return nil, ErrUnknownWeekday
}

func (w *Weekday) UnmarshalText(data []byte) error {
	switch string(data) {
	case Unset.String():
		*w = Unset
	case Sunday.String():
		*w = Sunday
	case Monday.String():
		*w = Monday
	case Tuesday.String():
		*w = Tuesday
	case Wednesday.String():
		*w = Wednesday
	case Thursday.String():
		*w = Thursday
	case Friday.String():
		*w = Friday
	case Saturday.String():
		*w = Saturday
	default:
		return ErrUnknownWeekday
	}
	return nil
}
