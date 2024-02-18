package logic

import (
	"yanblue/dao/mysql"
	"yanblue/models"
	"yanblue/pkg/jwt"
	"yanblue/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// check if user exist
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// generate uId
	userId := snowflake.GenID()
	// struct a user entity
	user := &models.User{
		UserID:   userId,
		Username: p.Username,
		Password: p.Password,
	}
	// insert user
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// generate token
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
