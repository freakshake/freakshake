package xsql

// Scanner is an interface used by scan function.
type Scanner interface {
	Scan(dest ...any) error
}
