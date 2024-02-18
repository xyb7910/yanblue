package logic

import (
	"yanblue/dao/mysql"
	"yanblue/models"
	"yanblue/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	// generate post id
	p.ID = snowflake.GenID()
	// save post to db
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	// save post to redis
	//err = redis.CreatePost(p.ID, p.CommunityID)
	return
}
