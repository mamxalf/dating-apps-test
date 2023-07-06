package repository

import (
	"context"
	"dating-apps/internal/domains/dating/model"
	"dating-apps/shared/failure"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"strings"
)

var swipeQueries = struct {
	insertSwipe string
}{
	insertSwipe: "INSERT INTO swipes %s VALUES %s",
}

func (repo *DatingRepositoryPostgres) SwipeProfile(_ context.Context, swipe *model.NewSwipe, dbTx *sqlx.Tx) (err error) {
	fieldsStr, valueListStr, args := composeInsertFieldAndParamsSwipeProfile(*swipe)
	commandQuery := fmt.Sprintf(swipeQueries.insertSwipe, fieldsStr, strings.Join(valueListStr, ","))
	_, err = dbTx.Exec(commandQuery, args...)
	if err != nil {
		log.Error().Err(err).Msg("[SwipeProfile - Repository] failed exec query")
		err = failure.InternalError(err)
		return err
	}
	return
}

func composeInsertFieldAndParamsSwipeProfile(register ...model.NewSwipe) (fieldStr string, valueListStr []string, args []interface{}) {
	var (
		fields []string = []string{
			"user_id",
			"profile_id",
			"is_like",
		}

		index = 0
	)
	fieldStr = fmt.Sprintf("(%s)", strings.Join(fields, ","))
	for _, params := range register {
		var values []string
		args = append(args,
			params.UserID,
			params.ProfileID,
			params.IsLike,
		)
		for i := 1; i <= len(fields); i++ {
			values = append(values, fmt.Sprintf("$%d", index+i))
		}
		index += len(fields)

		valueListStr = append(valueListStr, fmt.Sprintf("(%s)", strings.Join(values, ",")))
	}
	return
}
