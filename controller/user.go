package controller

import (
	"errors"
	"fmt"
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

// LoginHandler LoginHandler is a function to handle login request
func LoginHandler(c *gin.Context) {
	// 1. get the user info from the request body
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		// query param error
		zap.L().Error("Login with invalid param", zap.Error(err))
		// check err is validator.ValidationErrors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 2. handle the login logic
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login with error", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 3. return the response
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID),
		"user_name": user.Username,
		"token":     user.Token,
	})
}
