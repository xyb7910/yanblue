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

// GetCommunityDetailByID get community detail by id
func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	community = &models.CommunityDetail{}
	sqlStr := `select community_id, community_name, introduction, create_time 
		from community 
		where community_id = ?`

	if err := db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return community, err
}
