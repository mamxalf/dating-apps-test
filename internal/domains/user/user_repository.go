package user

import (
	"context"
	"database/sql"
	"dating-apps/infras"
	"dating-apps/shared/failure"
	"dating-apps/shared/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"strings"
)

var userQueries = struct {
	registerNewUser   string
	getUser           string
	createUserSession string
}{
	registerNewUser:   "INSERT INTO users %s VALUES %s",
	getUser:           "SELECT * FROM users %s",
	createUserSession: "INSERT INTO user_sessions %s VALUES %s",
}

type UserRepository interface {
	RegisterNewUser(ctx context.Context, register *Register) (err error)
	GetUserByEmail(email string) (user User, err error)
	CreateUserSession(ctx context.Context, userSession *UserSession) (err error)
}

type UserRepositoryPostgres struct {
	DB *infras.PostgresConn
}

func ProvideUserRepositoryPostgres(db *infras.PostgresConn) *UserRepositoryPostgres {
	s := new(UserRepositoryPostgres)
	s.DB = db
	return s
}

func (repo *UserRepositoryPostgres) CreateUserSession(ctx context.Context, userSession *UserSession) (err error) {
	fieldsStr, valueListStr, args := composeInsertFieldAndParamsUserSession(*userSession)
	commandQuery := fmt.Sprintf(userQueries.createUserSession, fieldsStr, strings.Join(valueListStr, ","))
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateUserSession - Repository] failed exec query")
		err = failure.InternalError(err)
		return err
	}
	return
}

func composeInsertFieldAndParamsUserSession(userSession ...UserSession) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string = []string{
			"user_id",
			"access_token",
			"refresh_token",
			"is_active",
		}

		index = 0
	)
	fieldStr = fmt.Sprintf("(%s)", strings.Join(fields, ","))
	for _, params := range userSession {
		var values []string
		args = append(args,
			params.UserID,
			params.AccessToken,
			params.RefreshToken,
			params.IsActive,
		)
		for i := 1; i <= len(fields); i++ {
			values = append(values, fmt.Sprintf("$%d", index+i))
		}
		index += len(fields)

		valueListStr = append(valueListStr, fmt.Sprintf("(%s)", strings.Join(values, ",")))
	}
	return
}

func (repo *UserRepositoryPostgres) GetUserByEmail(email string) (user User, err error) {
	whereClauses := " WHERE email = $1 LIMIT 1"
	query := fmt.Sprintf(userQueries.getUser, whereClauses)
	err = repo.DB.Read.Get(&user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = failure.NotFound("User not found!")
			return
		}
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}

	return
}

func (repo *UserRepositoryPostgres) RegisterNewUser(ctx context.Context, register *Register) (err error) {
	fieldsStr, valueListStr, args := composeInsertFieldAndParamsUser(*register)
	commandQuery := fmt.Sprintf(userQueries.registerNewUser, fieldsStr, strings.Join(valueListStr, ","))
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[RegisterNewUser - Repository] failed exec query")
		err = failure.InternalError(err)
		return err
	}
	return
}

func composeInsertFieldAndParamsUser(register ...Register) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string = []string{
			"username",
			"email",
			"password",
		}

		index = 0
	)
	fieldStr = fmt.Sprintf("(%s)", strings.Join(fields, ","))
	for _, params := range register {
		var values []string
		args = append(args,
			params.Username,
			params.Email,
			params.Password,
		)
		for i := 1; i <= len(fields); i++ {
			values = append(values, fmt.Sprintf("$%d", index+i))
		}
		index += len(fields)

		valueListStr = append(valueListStr, fmt.Sprintf("(%s)", strings.Join(values, ",")))
	}
	return
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
