package activeCount

import (
	"KpChatGpt/cache"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// DailyActiveCount 日活统计中间件
func DailyActiveCount(ctx *gin.Context) {
	// Get the current date
	now := time.Now()
	date := fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day())
	// Get the client's IP address
	ip := ctx.ClientIP()
	// Concatenate the date and IP to get the unique key
	key := date + ":" + ip
	// Check if the key exists in Redis
	count, err := cache.Rdb.Exists(ctx, key).Result()
	if err != nil {
		// Handle error
		fmt.Println("error checking key in Redis:", err)
	}
	if count == 0 {
		// If the key does not exist, increment the counter in Redis
		_, err := cache.Rdb.Incr(ctx, date).Result()
		if err != nil {
			// Handle error
			fmt.Println("error incrementing counter in Redis:", err)
		}
		// Set the key with an expiration time of one day
		_, err = cache.Rdb.Set(ctx, key, 1, time.Hour*24).Result()
		if err != nil {
			// Handle error
			fmt.Println("error setting key in Redis:", err)
		}
	}
	ctx.Next()
}
