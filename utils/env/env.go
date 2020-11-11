package env

import (
	"os"
	"strconv"
)

// Mode 模式
//
// 参考: https://github.com/gin-gonic/gin/blob/master/mode.go
type Mode int

// AppMode 开发模式
var AppMode Mode

const envKey = "APP_ENV"

const (
	// DebugMode 开发
	DebugMode = "debug"
	// ReleaseMode 线上
	ReleaseMode = "release"
)

const (
	// DebugCode 开发
	DebugCode = iota
	// ReleaseCode 线上
	ReleaseCode
)

// GetMode 获取模式
func GetMode() Mode {
	var mode = os.Getenv(envKey)
	switch mode {
	case DebugMode:
		return DebugCode
	case ReleaseMode:
		return ReleaseCode
	default:
		return DebugCode
	}
}

// GetPort 获取端口
func GetPort() int {
	var port = os.Getenv("PORT")
	if port == "" {
		return 2333
	}
	var p, _ = strconv.Atoi(port)
	return p
}

func init() {
	AppMode = GetMode()
}
