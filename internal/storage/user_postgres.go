package storage

import (
	"context"
	"database/sql"
	"errors"

	"github.com/freakshake/internal/derror"
	"github.com/freakshake/internal/domain"
	"github.com/freakshake/pkg/logger"
	"github.com/freakshake/pkg/type/email"
	"github.com/freakshake/pkg/type/id"
	"github.com/freakshake/pkg/type/offlim"
	"github.com/freakshake/xsql"
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

const deleteQuery = `
	UPDATE users
	SET 
		deleted_at = NOW() 
	WHERE id = $1 AND deleted_at IS NULL
`

var scanUser = func(s xsql.Scanner) (u domain.User, err error) {
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
	db     *sql.DB
	logger logger.Logger
}

func NewUserPostgresStorage(db *sql.DB, l logger.Logger) domain.UserStorage {
	return &userPostgres{
		db:     db,
		logger: l,
	}
}

func (s *userPostgres) Store(ctx context.Context, u domain.User) (id.ID[domain.User], error) {
	query := insertQuery + `($1, $2, $3, $4, $5, $6) RETURNING id`

	id, err := xsql.QueryOne(ctx, s.db, xsql.ScanID[id.ID[domain.User]], query,
		u.Avatar,
		u.FirstName,
		u.LastName,
		u.Email,
		u.MobileNumber,
		u.HashedPassword,
	)
	if err != nil {
		s.logger.Error(domain.UserDomain, logger.StorageLayer, err, logger.Args{})
		return 0, err
	}

	return id, nil
}

func (s *userPostgres) Find(ctx context.Context, uID id.ID[domain.User]) (domain.User, error) {
	query := selectQuery + `WHERE id = $1 AND deleted_at IS NULL LIMIT 1`

	u, err := xsql.QueryOne(ctx, s.db, scanUser, query, uID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, derror.ErrUnknownUser
		}
		s.logger.Error(domain.UserDomain, logger.StorageLayer, err, logger.Args{})
		return domain.User{}, err
	}

	return u, nil
}

func (s *userPostgres) FindByEmail(ctx context.Context, e email.Email) (domain.User, error) {
	query := selectQuery + `WHERE email = $1 AND deleted_at IS NULL LIMIT 1`

	u, err := xsql.QueryOne(ctx, s.db, scanUser, query, e)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.User{}, derror.ErrUnknownUser
		}
		s.logger.Error(domain.UserDomain, logger.StorageLayer, err, logger.Args{})
		return domain.User{}, err
	}

	return u, nil
}

func (s *userPostgres) FindAll(ctx context.Context, o offlim.Offset, l offlim.Limit) ([]domain.User, error) {
	query := selectQuery + `WHERE deleted_at IS NULL ORDER BY id OFFSET $1 LIMIT $2`

	var limit *offlim.Limit
	if l > 0 {
		limit = &l
	}

	u, err := xsql.QueryMany(ctx, s.db, scanUser, query, o, limit)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, derror.ErrUnknownUser
		}
		s.logger.Error(domain.UserDomain, logger.StorageLayer, err, logger.Args{})
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
		s.logger.Error(domain.UserDomain, logger.StorageLayer, err, logger.Args{})
		return err
	}

	return nil
}

func (s *userPostgres) Delete(ctx context.Context, uID id.ID[domain.User]) error {
	_, err := s.db.ExecContext(ctx, deleteQuery, uID)
	if err != nil {
		s.logger.Error(domain.UserDomain, logger.StorageLayer, err, logger.Args{})
		return err
	}
	return nil
}
