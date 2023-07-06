package repository

import (
	"context"
	"database/sql"
	"dating-apps/infras"
	"dating-apps/internal/domains/dating/model"
	"dating-apps/shared/failure"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type DatingRepository interface {
	SwipeProfile(ctx context.Context, swipe *model.NewSwipe) (err error)
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

func (repo *DatingRepositoryPostgres) exec(ctx context.Context, command string, args []interface{}) (sql.Result, error) {
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
