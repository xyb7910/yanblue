package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"yanblue/logic"
	"yanblue/models"
)

// CreatePostHandler create post handler
func CreatePostHandler(c *gin.Context) {
	// 1.get param and check param
	p := &models.Post{}
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p) err", zap.Error(err))
		zap.L().Error("create post failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// get user id from context
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 2.create post
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.return response
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler get post detail handler
func GetPostDetailHandler(c *gin.Context) {
	// 1.get param from url
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt(pidStr, 10, 64) failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	// 2.get post detail by id
	data, err := logic.GetPostDetailByID(pid)
	if err != nil {
		zap.L().Error("logic.GetPostDetailByID(pid) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.return response
	ResponseSuccess(c, data)
}

// GetPostListHandler get post list handler
func GetPostListHandler(c *gin.Context) {
	// 1.query page and limit param
	page, size := getPageAndSize(c)
	// 2.get post list by page and size
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList(page, size) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
