package repository

import (
	"context"
	"database/sql"
	"dating-apps/internal/domains/dating/model"
	"dating-apps/shared/failure"
	"dating-apps/shared/logger"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

var datingQueries = struct {
	getDatingProfile string
}{
	getDatingProfile: "SELECT *, COUNT(*) OVER() AS total_data FROM user_profiles %s",
}

func (repo *DatingRepositoryPostgres) GetProfile(
	_ context.Context,
	exceptID []string, gender string, page, size int,
) (res []model.Profile, err error) {
	whereClauses := " WHERE id NOT IN ($1) AND gender = $2 LIMIT $3 OFFSET $4"
	query := fmt.Sprintf(datingQueries.getDatingProfile, whereClauses)
	if len(exceptID) == 0 {
		exceptID = []string{uuid.Nil.String()}
	}
	err = repo.DB.Read.Select(&res, query, strings.Join(exceptID, ","), gender, size, (page-1)*size)
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
