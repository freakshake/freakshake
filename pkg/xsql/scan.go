package xsql

// Scanner is a type which can scan data to destinations.
// e.g. sql.Rows implements Scanner.
// It is used by scan function.
type Scanner interface {
	Scan(dest ...any) error
}

func ScanID[T any](s Scanner) (id T, err error) {
	if err := s.Scan(&id); err != nil {
		return id, err
	}
	return id, nil
}
