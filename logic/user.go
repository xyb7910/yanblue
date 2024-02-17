package logic

import (
	"yanblue/dao/mysql"
	"yanblue/models"
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
