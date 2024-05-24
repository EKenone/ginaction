package ginaction

import "github.com/gin-gonic/gin"

// AutoRegister 自动注册路由
func AutoRegister(router *gin.RouterGroup, cs ...Controller) {
	for _, c := range cs {
		g := router.Group(c.Group())
		for _, r := range c.Register() {
			do := c.ChooseMid(r.midType)
			do = append(do, r.do)
			g.Handle(r.method, r.createLastPath(), do...)
		}
	}
}
