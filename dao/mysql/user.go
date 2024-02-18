package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"yanblue/models"
)

const SecrectKey = "yanblue.com"

// CheckUserExist check user exist
func CheckUserExist(username string) (err error) {
	sqlStr := "select count(user_id) from user where username = ?"
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}

	if count > 0 {
		return ErrorUserExist
	}
	return
}

// InsertUser insert user
func InsertUser(user *models.User) (err error) {
	// encrypt password
	user.Password = encryptPassword(user.Password)
	// commit sql
	sqlStr := "insert into user(user_id,username, password) values(?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// EncryptPassword encrypt password
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(SecrectKey + oPassword))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	// user login password
	oPassword := user.Password
	sqlStr := "select user_id, username, password from user where username = ?"
	err = db.Get(user, sqlStr, user.Username)
	if err != sql.ErrNoRows {
		return ErrorUserExist
	}
	if err != nil {
		// user not exist, query sql error
		return err
	}
	// check password correct
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}

func GetUserByID(uid int64) (user *models.User, err error) {
	user = &models.User{}
	sqlStr := "select user_id, username from user where user_id = ?"
	err = db.Get(user, sqlStr, uid)
	return
}
