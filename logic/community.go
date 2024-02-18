package logic

import (
	"yanblue/dao/mysql"
	"yanblue/models"
)

// GetCommunityList return all community list
func GetCommunityList() ([]*models.Community, error) {
	// query database return all community list
	return mysql.GetCommunityList()
}
