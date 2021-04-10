package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	TokenExpireDuration = 2 * time.Hour
)

var jwtSercet = []byte("jwt")

type MyClaims struct {
	jwt.StandardClaims
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
}

// GenerateToken 生成token
func GenerateToken(username string, userId int64) (string, error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "bluebell",
		},
		UserId:   userId,
		Username: username,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSercet)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*MyClaims, error) {
	myClaims := new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, myClaims, func(token *jwt.Token) (interface{}, error) {
		return jwtSercet, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return myClaims, nil
}
