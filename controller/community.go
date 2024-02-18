package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"yanblue/logic"
)

func CommunityHandler(c *gin.Context) {
	// query all community (community_id, community_name)
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
