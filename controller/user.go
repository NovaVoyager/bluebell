package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/miaogu-go/bluebell/logic"
)

func SignUpHandler(c *gin.Context) {
	logic.Signup()
}
