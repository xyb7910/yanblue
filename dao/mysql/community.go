package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"yanblue/models"
)

// GetCommunityList get all community list
func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select community_id, community_name from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("no community found", zap.Error(err))
			err = nil
		}
	}
	return
}
