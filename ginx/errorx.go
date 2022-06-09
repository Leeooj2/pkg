package ginx

import "github.com/hololee2cn/pkg/errorx"

func BombErr(code int, format string, p ...interface{}) {
	errorx.BombErr(code, format, p...)
}

func CustomErr(v interface{}, code ...int) {
	errorx.CustomErr(v, code...)
}
