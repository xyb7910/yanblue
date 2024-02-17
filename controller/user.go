package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"yanblue/dao/mysql"
	"yanblue/logic"
	"yanblue/models"
)

// SignUpHandler SinUpHandler is a function to handle sign up request
func SignUpHandler(c *gin.Context) {
	// 1. get the user info from the request body
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// query param error
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// check err is validator.ValidationErrors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//2.handle the sign-up logic
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp with error", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3. return the response
	ResponseSuccess(c, nil)
}
