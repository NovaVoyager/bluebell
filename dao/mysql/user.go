package mysql

import (
	"database/sql"
	"time"

	"github.com/miaogu-go/bluebell/pkg/tools"

	"github.com/miaogu-go/bluebell/models"
)

const (
	PasswordSalt = "20210407160200"
)

var cstZone = time.FixedZone("CST", 8*3600)

type User struct {
	Id         uint64 `db:"id"`
	UserId     int64  `db:"user_id"`
	Username   string `db:"username"`
	Password   string `db:"password"`
	Email      string `db:"email"`
	CreateTime string `db:"create_time"`
	UpdateTime string `db:"update_time"`
	Gender     int    `db:"gender"`
}

// CreateUser 创建用户
func CreateUser(user *models.User) error {
	sql := "INSERT INTO `user` (`user_id`,`username`,`password`,`create_time`,`update_time`) VALUES (?,?,?,?,?)"
	user.Password = tools.EncryptPassword(user.Password, PasswordSalt)
	date := time.Now().In(cstZone)
	_, err := db.Exec(sql, user.UserId, user.Username, user.Password, date, date)
	if err != nil {
		return err
	}

	return nil
}

func QueryUserByUsername(username string) (*User, error) {
	sqlStr := "SELECT * FROM `user` WHERE `username`=?"
	user := new(User)
	err := db.Get(user, sqlStr, username)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CheckUserExist 根据用户名检查用户是否存在
func CheckUserExist(username string) (bool, error) {
	var count int

	sql := "SELECT COUNT(user_id) FROM `user` WHERE username=?"
	if err := db.Get(&count, sql, username); err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetUserByUserId 根据userId获取用户信息
func GetUserByUserId(userId int64) (*User, error) {
	sqlStr := "SELECT * FROM `user` WHERE user_id=?"
	user := new(User)
	err := db.Get(user, sqlStr, userId)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return user, nil
}
