package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	TokenExpireDuration = 2 * time.Hour
)

var (
	jwtSercet                 = []byte("jwt")
	jwtSercetFunc jwt.Keyfunc = func(token *jwt.Token) (interface{}, error) {
		return jwtSercet, nil
	}
)

type MyClaims struct {
	jwt.StandardClaims
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
}

// GenerateToken 生成token
func GenerateToken(username string, userId int64) (aToken, rToken string, err error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "bluebell",
		},
		UserId:   userId,
		Username: username,
	}
	aToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSercet)
	rToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		Issuer:    "bluebell",
	}).SignedString(jwtSercet)

	return
}

// ParseToken 解析token
func ParseToken(tokenString string) (*MyClaims, error) {
	myClaims := new(MyClaims)
	token, err := jwt.ParseWithClaims(tokenString, myClaims, jwtSercetFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return myClaims, nil
}

// RefreshToken 刷新token
func RefreshToken(aToken, rToken string) (newAToken, newRToken string, err error) {
	//解析验证refresh token 是否有效
	_, err = jwt.Parse(rToken, jwtSercetFunc)
	if err != nil {
		return
	}

	var myClaims MyClaims
	//从access token中解析出用户数据
	_, err = jwt.ParseWithClaims(aToken, &myClaims, jwtSercetFunc)
	if err == nil { //access token 未过期，返回原本数据
		newAToken = aToken
		newRToken = rToken
		return
	}
	v, _ := err.(*jwt.ValidationError)
	//access token为过期错误
	if v.Errors == jwt.ValidationErrorExpired {
		return GenerateToken(myClaims.Username, myClaims.UserId)
	}
	return
}
