package ginaction

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestServer(t *testing.T) {
	Server := gin.Default()

	AutoRegister(Server.Group("api"), TestDemo{})

	if err := Server.Run(":8080"); err != nil {
		panic(err)
	}
}

type TestDemo struct {
}

// ChooseMid 可以选择的服务中间件
func (a TestDemo) ChooseMid(router *gin.RouterGroup, t MidType) gin.IRoutes {
	switch t {
	case 1:
		return router.Use(Trace2())
	default:
		return router.Use(Trace())
	}
}

// Group 接口组标识
func (a TestDemo) Group() string {
	return "test"
}

// Register 接口注册
func (a TestDemo) Register() []Action {
	return []Action{
		NewAction("GET", a.test1),
		NewAction("GET", a.testPath2, UseMidType(1), UseLastPath("ttt")),
		NewAction("POST", a.testPath3, UseMidType(1), UseMidSep(0)),
		NewAction("POST", a.testPath4, UseMidSep('_')),
	}
}

func (a TestDemo) test1(c *gin.Context) {

}

func (a TestDemo) testPath2(c *gin.Context) {

}

func (a TestDemo) testPath3(c *gin.Context) {

}

func (a TestDemo) testPath4(c *gin.Context) {

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
