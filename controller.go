package ginaction

import "github.com/gin-gonic/gin"

type Controller interface {
	Register() []Action                    //需要中间件的方法
	Group() string                         //控制器分组
	ChooseMid(t MidType) []gin.HandlerFunc //选择中间件
}
