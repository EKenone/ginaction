package ginaction

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"runtime"
	"strings"
)

const defaultMidSep = '-'

type MidType uint8

type Controller interface {
	Register() []Action                                       //需要中间件的方法
	Group() string                                            //控制器分组
	ChooseMid(router *gin.RouterGroup, t MidType) gin.IRoutes //选择中间件
}

type Action struct {
	method   string          //请求方法
	do       gin.HandlerFunc //执行函数
	midType  MidType         //中间件类型
	midSep   byte            //路由大写分隔
	lastPath string          //最后一届路由
}

// NewAction 实例化一个action
func NewAction(method string, do gin.HandlerFunc, opts ...Option) Action {
	var action = Action{
		method: method,
		do:     do,
		midSep: defaultMidSep,
	}

	for _, opt := range opts {
		opt(&action)
	}

	return action
}

type midTypeMap struct {
	group  *gin.RouterGroup
	routes gin.IRoutes
}

// AutoRegister 自动注册路由
func AutoRegister(router *gin.RouterGroup, cs ...Controller) {
	for _, c := range cs {
		var midMap = make(map[MidType]midTypeMap)
		for _, v := range c.Register() {
			midR, ok := midMap[v.midType]
			if !ok {
				r := router.Group(c.Group())
				midR = midTypeMap{
					group:  r,
					routes: c.ChooseMid(r, v.midType),
				}

				midMap[v.midType] = midR
			}
			midR.routes.Handle(v.method, v.createLastPath(), v.do)
		}
	}
}

// 默认的路由最后一部分
func (a *Action) createLastPath() string {
	if a.lastPath != "" {
		return a.lastPath
	}

	// 获取函数名称
	fn := runtime.FuncForPC(reflect.ValueOf(a.do).Pointer()).Name()

	// 用 seps 进行分割
	fields := strings.FieldsFunc(fn, func(sep rune) bool {
		return sep == '.'
	})

	var lastPath string
	if size := len(fields); size > 0 {
		lastPath = strings.TrimSuffix(fields[size-1], "-fm")
	}

	if a.midSep != 0 {
		return midString(lastPath, a.midSep)
	}

	return lastPath
}
