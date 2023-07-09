package repository

import (
	"context"
	"dating-apps/internal/domains/user/model"
	"dating-apps/shared/failure"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
)

var userSessionQueries = struct {
	createUserSession string
}{
	createUserSession: "INSERT INTO user_sessions %s VALUES %s",
}

func (repo *UserRepositoryPostgres) CreateUserSession(ctx context.Context, userSession *model.UserSession) (err error) {
	fieldsStr, valueListStr, args := composeInsertFieldAndParamsUserSession(*userSession)
	commandQuery := fmt.Sprintf(userSessionQueries.createUserSession, fieldsStr, strings.Join(valueListStr, ","))
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[CreateUserSession - Repository] failed exec query")
		err = failure.InternalError(err)
		return err
	}
	return
}

func composeInsertFieldAndParamsUserSession(userSession ...model.UserSession) (fieldStr string, valueListStr []string, args []any) {
	var (
		fields = []string{
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
