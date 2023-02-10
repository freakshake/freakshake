package xsql

import (
	"context"
	"database/sql"

	"github.com/mehdieidi/storm/pkg/xerror"
)

// Scanner is an interface used by scan function.
type Scanner interface {
	Scan(dest ...any) error
}

// GetOne is used to retrieve a single row from a database using the provided query and arguments.
//
// Example:
//
//	// Create a new context with a timeout of 10 seconds
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	// Create a new sql DB connection
//	db, err := sql.Open("mysql", "user@tcp(127.0.0.1:3306)/mydb")
//	if err != nil {
//		panic(err)
//	}
//	// Define our scan function
//	scanFunc := func(s Scanner) (string, error) {
//		var name string
//		if err := s.Scan(&name); err != nil {
//			return "", err
//		}
//		return name, nil
//	}
//	// Execute our query with our scan function and args
//	result, err := GetOne(ctx, db, scanFunc, "SELECT name FROM users WHERE id = ?", 1)
//	if err != nil {
//		panic(err)
//	}
func GetOne[T any](
	ctx context.Context,
	db *sql.DB,
	scan func(Scanner) (T, error),
	query string,
	args ...interface{},
) (_ T, err error) {
	defer xerror.Wrap(&err, "GetOne(ctx, db, scan, %q, %q)", query, args)

	row := db.QueryRowContext(ctx, query, args...)
	return scan(row)
}

// GetMany is used to retrieve multiple rows from a database using a query and arguments.
//
// Example:
//
//	// Create a new context with a timeout of 10 seconds
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	// Create a new sql DB connection
//	db, err := sql.Open("mysql", "user@tcp(127.0.0.1:3306)/mydb")
//	if err != nil {
//		panic(err)
//	}
//	// Define our scan function
//	scanFunc := func(s Scanner) (string, error) {
//		var name string
//		if err := s.Scan(&name); err != nil {
//			return "", err
//		}
//		return name, nil
//	}
//	// Execute our query with our scan function and args
//	results, err := GetMany(ctx, db, scanFunc, "SELECT name FROM users WHERE age = ?", 34)
//	if err != nil {
//		panic(err)
//	}
func GetMany[T any](
	ctx context.Context,
	db *sql.DB,
	scan func(Scanner) (_ T, err error),
	query string,
	args ...interface{},
) (_ []T, err error) {
	defer xerror.Wrap(&err, "GetMany(ctx, db, scan, %q, %q)", query, args)

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := rows.Close()
		if cerr != nil {
			xerror.Wrap(&err, "rows.Close()")
		}
	}()

	results := make([]T, 0, 20)

	for rows.Next() {
		res, err := scan(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, res)
	}

	return results, nil
}
