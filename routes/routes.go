package routes

import (
	"github.com/miaogu-go/bluebell/middlewares"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/miaogu-go/bluebell/controller"
	"github.com/miaogu-go/bluebell/logger"
	"github.com/miaogu-go/bluebell/settings"

	"github.com/gin-gonic/gin"
	_ "github.com/miaogu-go/bluebell/docs"
)

func Setup() *gin.Engine {
	setRunMode()
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")
	//注册
	v1.POST("/signup", controller.SignUpHandler)
	//登录
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware())
	//验证token
	v1.POST("/ping", controller.PingHandler)
	{
		//获取社区列表
		v1.GET("/community", controller.CommunityHandler)
		//获取社区详情
		v1.GET("/community/:id", controller.GetCommunityDetail)
		//创建帖子
		v1.POST("/post/create", controller.CreatePostHandler)
		//获取帖子详情
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		//帖子列表
		v1.POST("/posts", controller.GetPostsHandler)
		v1.POST("/posts2", controller.GetPosts2Handler)
		//投票
		v1.POST("/vote", controller.VoteHandler)
	}
	//刷新token
	r.POST("/refresh", controller.RefreshTokenHandler)

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	return r
}

// setRunMode 设置运行模式
func setRunMode() {
	switch settings.Conf.AppConf.Mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}
