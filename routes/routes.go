package routes

import (
	"fmt"
	"net/http"

	"github.com/miaogu-go/bluebell/controller"
	"github.com/miaogu-go/bluebell/logger"
	"github.com/miaogu-go/bluebell/settings"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	setRunMode()
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("%s", settings.Conf.AppConf.Name))
	})

	r.POST("/signup", controller.SignUpHandler)

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
