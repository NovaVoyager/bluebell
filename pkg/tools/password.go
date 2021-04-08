package tools

import (
	"crypto/md5"
	"encoding/hex"
)

// EncryptPassword 加密密码
func EncryptPassword(str, salt string) string {
	h := md5.New()
	h.Write([]byte(str + salt))
	return hex.EncodeToString(h.Sum(nil))
}
