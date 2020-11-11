// create by d1y<chenhonzhou@gmail.com>
// frist write date: 2020-10-20

package conf

import "github.com/d1y/note/utils/env"

// Version 版本号
const Version = "0.0.1"

// Author 作者
const Author = "d1y"

// Mail 邮箱
const Mail = "chenhonzhou@gmail.com"

// ExposePort 端口
var ExposePort = 2333

// ReservedKeywords 保留关键字
var ReservedKeywords = []string{
	"api",
	"admin",
}

// WebPrefix web前端路径
const WebPrefix = "note"

// DatabaseName 数据库名
const DatabaseName = "note"

func init() {
	ExposePort = env.GetPort()
}
