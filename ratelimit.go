package ratelimit

import (
	"math"
	"time"
)

var (
	lastReqTime  int64 = time.Now().Unix() // 上个请求的时间
	leftTokenNum int64 = 1                 // 上次请求时剩余的令牌
	tokenRate    int64 = 1                 // 每秒增加一个token
	bucketCap    int64 = 100               // 令牌总数量

)

// 每个请求都会调用take方法
func take() bool {
	now := time.Now().Unix() // 当前时间
	// 以一定速率放入令牌，超过容量就会溢出，所以要 math.min()
	nowTokenNum := int64(math.Min(float64((now-lastReqTime)*tokenRate+leftTokenNum), float64(bucketCap)))
	if nowTokenNum > 0 {
		lastReqTime = now
		leftTokenNum--
		return true
	}
	return false
}
