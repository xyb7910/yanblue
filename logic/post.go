package logic

import (
	"go.uber.org/zap"
	"yanblue/dao/mysql"
	"yanblue/models"
	"yanblue/pkg/snowflake"
)

// CreatePost create a post
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

// GetPostDetailByID get post detail by id
func GetPostDetailByID(pid int64) (data *models.ApiPostDetail, err error) {
	// get post detail from db
	post, err := mysql.GetPostByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailByID(pid) failed",
			zap.Int64("pid", pid),
			zap.Error(err))
		return
	}
	// query author message by author id
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID(post.AuthorID) failed",
			zap.Int64("post.AuthorID", post.AuthorID),
			zap.Error(err))
		return
	}
	// query community message by community id
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed",
			zap.Int64("post.CommunityID", post.CommunityID),
			zap.Error(err))
		return
	}
	// pick up the author and community message
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}
