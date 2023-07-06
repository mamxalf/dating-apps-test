package repository

import (
	"context"
	"database/sql"
	"dating-apps/infras"
	"dating-apps/internal/domains/user/model"
	"dating-apps/shared/failure"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type UserRepository interface {
	RegisterNewUser(ctx context.Context, register *model.UserRegister) (err error)
	GetUserByEmail(email string) (user model.User, err error)
	CreateUserSession(ctx context.Context, userSession *model.UserSession) (err error)
}

type UserRepositoryPostgres struct {
	DB *infras.PostgresConn
}

func ProvideUserRepositoryPostgres(db *infras.PostgresConn) *UserRepositoryPostgres {
	s := new(UserRepositoryPostgres)
	s.DB = db
	return s
}

func (repo *UserRepositoryPostgres) exec(ctx context.Context, command string, args []interface{}) (sql.Result, error) {
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
