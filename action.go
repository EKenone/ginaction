package ginaction

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"runtime"
	"strings"
)

const defaultMidSep = '-'

type MidType any

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
