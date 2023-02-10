package bucket

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// 令牌桶结构体
type TokenBucket struct {
	rate      int           // 每秒生成令牌数
	bucket    chan struct{} // 令牌桶
	lastCheck time.Time     // 上次检查时间
}

// 创建令牌桶
func NewTokenBucket(rate int) *TokenBucket {
	bucket := make(chan struct{}, rate)
	for i := 0; i < rate; i++ {
		bucket <- struct{}{}
	}
	return &TokenBucket{
		rate:      rate,
		bucket:    bucket,
		lastCheck: time.Now(),
	}
}

// 从令牌桶中取出令牌
func (b *TokenBucket) Take() bool {
	select {
	case <-b.bucket:
		return true
	default:
		return false
	}
}

// 检查令牌桶是否有剩余令牌，并加入新令牌
func (b *TokenBucket) Check() {
	now := time.Now()
	diff := now.Sub(b.lastCheck).Seconds()
	num := int(diff * float64(b.rate))
	for i := 0; i < num; i++ {
		select {
		case b.bucket <- struct{}{}:
		default:
		}
	}
	b.lastCheck = now
}

// IP 令牌桶中间件
func IpTokenBucketMiddleware(rate int) gin.HandlerFunc {
	buckets := make(map[string]*TokenBucket)
	return func(c *gin.Context) {
		ip := strings.Split(c.ClientIP(), ":")[0]
		bucket, ok := buckets[ip]
		if !ok {
			buckets[ip] = NewTokenBucket(rate)
		}
		bucket = buckets[ip]
		if bucket.Take() {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": http.StatusText(http.StatusTooManyRequests),
			})
		}
	}
}
