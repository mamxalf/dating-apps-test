package repository

import (
	"context"
	"dating-apps/internal/domains/user/model"
	"dating-apps/shared/failure"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"
)

var userProfileQueries = struct {
	insertUserProfile string
}{
	insertUserProfile: "INSERT INTO user_profiles %s VALUES %s",
}

func (repo *UserRepositoryPostgres) InsertUserprofile(ctx context.Context, newModel *model.UserProfile) (err error) {
	fieldsStr, valueListStr, args := composeInsertFieldAndParamsUserProfile(*newModel)
	commandQuery := fmt.Sprintf(userProfileQueries.insertUserProfile, fieldsStr, strings.Join(valueListStr, ","))
	_, err = repo.exec(ctx, commandQuery, args)
	if err != nil {
		log.Error().Err(err).Msg("[InsertUserprofile - Repository] failed exec query")
		err = failure.InternalError(err)
		return err
	}

	return
}

func composeInsertFieldAndParamsUserProfile(userSession ...model.UserProfile) (fieldStr string, valueListStr []string, args []any) {
	var (
		fields = []string{
			"user_id",
			"full_name",
			"age",
			"gender",
		}

		index = 0
	)
	fieldStr = fmt.Sprintf("(%s)", strings.Join(fields, ","))
	for _, params := range userSession {
		var values []string
		args = append(args,
			params.UserID,
			params.FullName,
			params.Age,
			params.Gender,
		)
		for i := 1; i <= len(fields); i++ {
			values = append(values, fmt.Sprintf("$%d", index+i))
		}
		index += len(fields)

		valueListStr = append(valueListStr, fmt.Sprintf("(%s)", strings.Join(values, ",")))
	}
	return
}
