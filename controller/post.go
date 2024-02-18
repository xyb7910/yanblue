package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"yanblue/logic"
	"yanblue/models"
)

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
