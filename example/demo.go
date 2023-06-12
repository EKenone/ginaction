package example

import (
	"ginaction"
	"github.com/gin-gonic/gin"
)

type Demo struct {
	Api
}

// Group 接口组标识
func (a Demo) Group() string {
	return "test"
}

// Register 接口注册
func (a Demo) Register() []ginaction.Action {
	return []ginaction.Action{
		ginaction.NewAction("GET", a.test1),
		ginaction.NewAction("GET", a.testPath2, ginaction.UseMidType(1), ginaction.UseLastPath("ttt")),
		ginaction.NewAction("POST", a.testPath3, ginaction.UseMidType(1), ginaction.UseMidSep(0)),
		ginaction.NewAction("POST", a.testPath4, ginaction.UseMidSep('_')),
	}
}

func (a Demo) test1(c *gin.Context) {

}

func (a Demo) testPath2(c *gin.Context) {

}

func (a Demo) testPath3(c *gin.Context) {

}

func (a Demo) testPath4(c *gin.Context) {

}
