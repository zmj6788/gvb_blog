package random

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成随机验证码
func Code (length int) string{
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%4v",rand.Intn(10000))
}
