// create by d1y<chenhonzhou@gmail.com>
// 随机函数工具包

package randx

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// 随机最小长度
var randomMinLength = 1

// 随机最大长度
var randomMaxLength = 10

// CreateRandString 创建随机字符(指定长度)
//
// https://stackoverflow.com/a/31832326
func CreateRandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// CreateRandomNumber 创建随机数值(int)
//
// https://www.codenong.com/23577091
func CreateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	n := min + rand.Intn(max-min+1)
	return n
}

// CreateEasyRandomString 创建随机字符
func CreateEasyRandomString() string {
	var outputLength = CreateRandomNumber(randomMinLength, randomMaxLength)
	return CreateRandString(outputLength)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
