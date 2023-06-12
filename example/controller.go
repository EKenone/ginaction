package example

import (
	"ginaction"
	"github.com/gin-gonic/gin"
)

var Server *gin.Engine

func Register() {
	r := Server.Group("api")

	ginaction.AutoRegister(r, Demo{})
}
