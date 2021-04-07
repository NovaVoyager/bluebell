package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/miaogu-go/bluebell/models"
)

const (
	PasswordSalt = "20210407160200"
)

var cstZone = time.FixedZone("CST", 8*3600)

// CreateUser 创建用户
func CreateUser(user *models.User) error {
	sql := "INSERT INTO `user` (`user_id`,`username`,`password`,`create_time`,`update_time`) VALUES (?,?,?,?,?)"
	user.Password = encryptPassword(user.Password, PasswordSalt)
	date := time.Now().In(cstZone)
	_, err := db.Exec(sql, user.UserId, user.Username, user.Password, date, date)
	if err != nil {
		return err
	}

	return nil
}

func QueryUserByUsername() {

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

// encryptPassword 加密密码
func encryptPassword(str, salt string) string {
	h := md5.New()
	h.Write([]byte(str + salt))
	return hex.EncodeToString(h.Sum(nil))
}
