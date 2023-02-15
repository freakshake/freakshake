package repository

import (
	"context"
	"database/sql"

	"github.com/mehdieidi/storm/internal/domain"
	"github.com/mehdieidi/storm/pkg/type/id"
	"github.com/mehdieidi/storm/pkg/type/offlim"
	"github.com/mehdieidi/storm/pkg/xsql"
)

const insertQuery = `
	INSERT INTO users (
		avatar,
		first_name,
		last_name,
		email,
		mobile_number,
		hashed_password
	) VALUES
`

const selectQuery = `
	SELECT
		id,
		avatar,
		first_name,
		last_name,
		email,
		mobile_number,
		hashed_password,
		created_at,
		updated_at,
		deleted_at
	FROM users
`

const updateQuery = `
	UPDATE users
	SET
		avatar = $2,
		first_name = $3,
		last_name = $4,
		email = $5,
		mobile_number = $6,
		hashed_password = $7,
		updated_at = NOW()
`

var scanInsert = func(s xsql.Scanner) (id id.ID[domain.User], err error) {
	if err := s.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

var scanRetrieve = func(s xsql.Scanner) (u domain.User, err error) {
	if err := s.Scan(
		&u.ID,
		&u.Avatar,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.MobileNumber,
		&u.HashedPassword,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.DeletedAt,
	); err != nil {
		return domain.User{}, err
	}
	return u, nil
}

type userPostgres struct {
	db *sql.DB
}

func NewUserPostgresStorage(db *sql.DB) domain.UserStorage {
	return &userPostgres{
		db: db,
	}
}

func (s *userPostgres) Store(ctx context.Context, u domain.User) (id.ID[domain.User], error) {
	query := insertQuery + `($1, $2, $3, $4, $5, $6) RETURNING id`

	id, err := xsql.QueryOne(ctx, s.db, scanInsert, query,
		u.Avatar,
		u.FirstName,
		u.LastName,
		u.Email,
		u.MobileNumber,
		u.HashedPassword,
	)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *userPostgres) Find(ctx context.Context, uID id.ID[domain.User]) (domain.User, error) {
	query := selectQuery + `WHERE id = $1 AND deleted_at IS NULL LIMIT 1`

	u, err := xsql.QueryOne(ctx, s.db, scanRetrieve, query, uID)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (s *userPostgres) FindAll(ctx context.Context, o offlim.Offset, l offlim.Limit) ([]domain.User, error) {
	query := selectQuery + `WHERE deleted_at IS NULL ORDER BY id OFFSET $1 LIMIT $2`

	u, err := xsql.QueryMany(ctx, s.db, scanRetrieve, query, o, l)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (s *userPostgres) Update(ctx context.Context, uID id.ID[domain.User], u domain.User) error {
	query := updateQuery + `WHERE id = $1 AND deleted_at IS NULL`

	_, err := s.db.ExecContext(ctx, query, uID,
		u.Avatar,
		u.FirstName,
		u.LastName,
		u.Email,
		u.MobileNumber,
		u.HashedPassword,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *userPostgres) Delete(ctx context.Context, uID id.ID[domain.User]) error {
	const query = `UPDATE users SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`

	_, err := s.db.ExecContext(ctx, query, uID)
	if err != nil {
		return err
	}

	return nil
}
