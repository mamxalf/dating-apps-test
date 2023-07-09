package repository

import (
	"context"
	"dating-apps/shared/failure"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func (repo *UserRepositoryPostgres) SubscribeUserPremium(ctx context.Context, userID uuid.UUID) (err error) {
	setUpdate := "is_verified = true"
	whereClause := "id = $1"
	commandQuery := fmt.Sprintf(userQueries.updateUser, setUpdate, whereClause)
	_, err = repo.exec(ctx, commandQuery, []any{userID.String()})
	if err != nil {
		log.Error().Err(err).Msg("[RegisterNewUser - Repository] failed exec query")
		err = failure.InternalError(err)
		return err
	}
	return
}
