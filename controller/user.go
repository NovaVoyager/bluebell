package controller

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/bluebell/logic"
	"github.com/miaogu-go/bluebell/models"
)

// @Summary 用户注册
// @Description 用户注册
// @Tags 用户注册
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /api/v1/signup [post]
// SignUpHandler 用户注册
func SignUpHandler(c *gin.Context) {
	param := new(models.SignupReq)
	if err := c.ShouldBindJSON(param); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
			return
		}
		ResponseError(c, CodeInvalidParam)
		return
	}
	err := logic.Signup(param)
	if err != nil {
		if errors.Is(err, logic.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		zap.L().Error("注册失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// LoginHandler 用户登录
func LoginHandler(c *gin.Context) {
	//验证参数
	param := new(models.LoginReq)
	if err := c.ShouldBind(param); err != nil {
		zap.L().Error("login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
			return
		}
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		ResponseError(c, CodeInvalidParam)
		return
	}
	token, err := logic.Login(param)
	if err != nil {
		zap.L().Error("LoginHandler failed", zap.String("username", param.User), zap.Error(err))
		if errors.Is(err, logic.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		if errors.Is(err, logic.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	ResponseSuccess(c, token)
}

// PingHandler
func PingHandler(c *gin.Context) {
	ResponseSuccess(c, "success")
}

// RefreshTokenHandler 刷新token
func RefreshTokenHandler(c *gin.Context) {
	accessToken := c.Request.Header.Get("Authorization")
	refreshToken := c.Request.Header.Get("Refresh-Token")
	if accessToken == "" || refreshToken == "" {
		ResponseError(c, CodeTokenInvalid)
		return
	}
	token, err := logic.RefreshToken(c, accessToken, refreshToken)
	if err != nil {
		zap.L().Error("RefreshTokenHandler failed", zap.String("accessToken", accessToken),
			zap.String("refreshToken", refreshToken), zap.Error(err))
		ResponseError(c, CodeRefreshTokenFail)
		return
	}
	ResponseSuccess(c, token)
}
