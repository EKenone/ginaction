package example

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tdeken/ginaction"
)

type Api struct {
}

// ChooseMid 可以选择的服务中间件
func (a Api) ChooseMid(router *gin.RouterGroup, t ginaction.MidType) gin.IRoutes {
	switch t {
	case 1:
		return router.Use(Trace2())
	default:
		return router.Use(Trace())
	}
}

// Trace 链路追踪
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(111)
		c.Next()
	}
}

// Trace 链路追踪
func Trace2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(222)
		c.Next()
	}
}
