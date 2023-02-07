package minute

import "strconv"

type Minute uint8

func (m Minute) String() string {
	return strconv.Itoa(int(m))
}

func (m Minute) MarshalJSON() ([]byte, error) {
	return []byte(m.String()), nil
}

func (m *Minute) UnmarshalJSON(data []byte) error {
	d, err := strconv.Atoi(string(data))
	if err != nil {
		return ErrInvalidMinute
	}

	*m = Minute(d)

	return nil
}
