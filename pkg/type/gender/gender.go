package gender

import (
	"strconv"
)

type Gender uint8

const (
	Unset Gender = iota
	Male
	Female
)

func (g Gender) String() string {
	switch g {
	case Unset:
		return "unset"
	case Male:
		return "male"
	case Female:
		return "female"
	default:
		return ""
	}
}

func (g Gender) MarshalText() ([]byte, error) {
	if s := g.String(); s != "" {
		return []byte(s), nil
	}
	return nil, ErrUnknownGender
}

func (g *Gender) UnmarshalText(p []byte) error {
	switch string(p) {
	case Unset.String():
		*g = Unset
	case Male.String():
		*g = Male
	case Female.String():
		*g = Female
	default:
		return ErrUnknownGender
	}
	return nil
}

func JoinGenders(genders []Gender, sep string) (joined string) {
	for i := 0; i < len(genders)-1; i++ {
		joined += strconv.Itoa(int(genders[i])) + sep
	}
	joined += strconv.Itoa(int(genders[len(genders)-1]))

	return
}
