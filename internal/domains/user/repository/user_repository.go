package repository

import (
	"context"
	"database/sql"
	"dating-apps/internal/domains/user/model"
	"dating-apps/shared/failure"
	"dating-apps/shared/logger"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
)

var userQueries = struct {
	registerNewUser string
	getUser         string
	getUserProfile  string
}{
	registerNewUser: "INSERT INTO users %s VALUES %s",
	getUser:         "SELECT * FROM users %s",
	getUserProfile: `
		SELECT
			u.id, u.username, u.email, u.is_verified,
			up.full_name, up.age, up.gender
		FROM users u
			INNER JOIN user_profiles up on u.id = up.user_id
		 %s`,
}

func (repo *UserRepositoryPostgres) GetUserByEmail(email string) (user model.User, err error) {
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

func (repo *UserRepositoryPostgres) GetUserProfileByEmail(email string) (user model.FullUserProfile, err error) {
	whereClauses := " WHERE u.email = $1 LIMIT 1"
	query := fmt.Sprintf(userQueries.getUserProfile, whereClauses)
	err = repo.DB.Read.Get(&user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = failure.NotFound("User Profile not found, Please fill User Profile first!")
			return
		}
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}

	return
}

func (repo *UserRepositoryPostgres) RegisterNewUser(ctx context.Context, register *model.UserRegister) (err error) {
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

func composeInsertFieldAndParamsUser(register ...model.UserRegister) (fieldStr string, valueListStr []string, args []any) {
	var (
		fields = []string{
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
