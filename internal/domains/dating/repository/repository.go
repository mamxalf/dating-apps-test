package repository

import (
	"context"
	"database/sql"
	"dating-apps/infras"
	"dating-apps/internal/domains/dating/model"
	"dating-apps/shared/failure"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type DatingRepository interface {
	SwipeProfile(ctx context.Context, swipe *model.NewSwipe, dbTx *sqlx.Tx) (err error)
	GetProfile(ctx context.Context, exceptID []string, gender string, page int, size int) (res []model.Profile, err error)

	// SwipeIncr SwipeCacheListID SwipeCacheExpiry for cache data
	SwipeIncr(userID uuid.UUID) (amount int64, err error)
	SwipeCacheListID(userID uuid.UUID, profileID uuid.UUID) (err error)
	SwipeCacheExpiry(userID uuid.UUID) (err error)
	GetSwipeCacheListID(userID uuid.UUID) (data []string, err error)

	// BeginTx RollbackTx CommitTx will be used in service layer
	BeginTx() (*sqlx.Tx, error)
	RollbackTx(tx *sqlx.Tx) error
	CommitTx(tx *sqlx.Tx) error
}

type DatingRepositoryPostgres struct {
	DB    *infras.PostgresConn
	Cache *infras.RedisConn
}

func ProvideDatingRepositoryPostgres(db *infras.PostgresConn, cache *infras.RedisConn) *DatingRepositoryPostgres {
	s := new(DatingRepositoryPostgres)
	s.DB = db
	s.Cache = cache
	return s
}

func (repo *DatingRepositoryPostgres) exec(ctx context.Context, command string, args []any) (sql.Result, error) { //nolint:unused
	var (
		stmt *sqlx.Stmt
		err  error
	)
	stmt, err = repo.DB.Write.PreparexContext(ctx, command)
	if err != nil {
		log.Error().Err(err).Msg("[exec] failed prepare query")
		return nil, failure.InternalError(err)
	}

	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		log.Error().Err(err).Msg("[exec] failed exec query")
		return nil, failure.InternalError(err)
	}

	return result, nil
}

// BeginTx will initiate the database transaction and will be used in use case layer later
func (repo *DatingRepositoryPostgres) BeginTx() (tx *sqlx.Tx, err error) {
	tx, err = repo.DB.Write.Beginx()
	if err != nil {
		log.Error().Err(err)
		err = failure.InternalError(err)
	}

	return
}

// RollbackTx will rollback the database transaction
func (repo *DatingRepositoryPostgres) RollbackTx(tx *sqlx.Tx) (err error) {
	err = tx.Rollback()
	if err != nil {
		log.Error().Err(err)
		err = failure.InternalError(err)
	}

	return
}

// CommitTx will commit the database transaction
func (repo *DatingRepositoryPostgres) CommitTx(tx *sqlx.Tx) (err error) {
	err = tx.Commit()
	if err != nil {
		log.Error().Err(err)
		err = failure.InternalError(err)
	}

	return
}
