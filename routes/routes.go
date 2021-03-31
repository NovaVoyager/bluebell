package routes

import (
	"fmt"
	"net/http"
	"web_app/logger"
	"web_app/settings"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("%s", settings.Conf.AppConf.Name))
	})

	return r
}
