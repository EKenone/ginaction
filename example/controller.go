package example

import (
	"github.com/gin-gonic/gin"
	"github.com/tdeken/ginaction"
)

var Server *gin.Engine

func Register() {
	r := Server.Group("api")

	ginaction.AutoRegister(r, Demo{})
}
