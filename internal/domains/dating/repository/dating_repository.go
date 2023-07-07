package repository

import (
	"context"
	"database/sql"
	"dating-apps/internal/domains/dating/model"
	"dating-apps/shared/failure"
	"dating-apps/shared/logger"
	"fmt"
)

var datingQueries = struct {
	getDatingProfile string
}{
	getDatingProfile: "SELECT * FROM user_profiles %s",
}

func (repo *DatingRepositoryPostgres) GetProfile(_ context.Context, exceptID []string, limit int, offset int) (res []model.Profile, err error) {
	whereClauses := " WHERE profile_id NOT IN ($1) LIMIT $2 OFFSET $3"
	query := fmt.Sprintf(datingQueries.getDatingProfile, whereClauses)
	err = repo.DB.Read.Get(&res, query, exceptID, limit, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			err = failure.NotFound("Dating profile not found!")
			return
		}
		logger.ErrorWithStack(err)
		err = failure.InternalError(err)
		return
	}

	return
}
