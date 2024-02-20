package logic

import (
	"go.uber.org/zap"
	"yanblue/dao/mysql"
	"yanblue/dao/redis"
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

func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		// query author message by author id
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(post.AuthorID) failed",
				zap.Int64("post.AuthorID", post.AuthorID),
				zap.Error(err))
			continue
		}
		// query community message by community id
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed",
				zap.Int64("post.CommunityID", post.CommunityID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	// get id from redis
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("ids", ids))

	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("posts", posts))

	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	// query author message and back data to add post detail
	for idx, post := range posts {
		// get author message by author id
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(post.AuthorID) failed",
				zap.Int64("post.AuthorID", post.AuthorID),
				zap.Error(err))
			continue
		}
		// get community message by community id
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.CommunityID) failed",
				zap.Int64("post.CommunityID", post.CommunityID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}
