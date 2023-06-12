package ginaction

type Option func(do *Action)

// UseMidType 使用中间件类型
func UseMidType(t MidType) Option {
	return func(do *Action) {
		do.midType = t
	}
}

// UseLastPath 定义路由最后一节
func UseLastPath(path string) Option {
	return func(do *Action) {
		do.lastPath = path
	}
}

// UseMidSep 路由大小写分隔线
func UseMidSep(sep byte) Option {
	return func(do *Action) {
		do.midSep = sep
	}
}
