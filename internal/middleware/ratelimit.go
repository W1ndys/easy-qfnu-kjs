package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// rateLimitEntry 记录单个客户端的最后请求时间
type rateLimitEntry struct {
	lastAccess time.Time
}

// RateLimiter 基于 IP + User-Agent 的速率限制器
type RateLimiter struct {
	mu       sync.Mutex
	entries  map[string]*rateLimitEntry
	interval time.Duration // 两次请求之间的最小间隔
}

// NewRateLimiter 创建一个速率限制器
// interval 为两次请求之间的最小时间间隔
func NewRateLimiter(interval time.Duration) *RateLimiter {
	rl := &RateLimiter{
		entries:  make(map[string]*rateLimitEntry),
		interval: interval,
	}
	// 启动后台清理协程，每 interval*2 清理一次过期条目
	go rl.cleanup(interval * 2)
	return rl
}

// buildKey 根据客户端 IP 和 User-Agent 生成唯一标识
// 对 User-Agent 做哈希处理，避免超长字符串作为 map key
func buildKey(clientIP, userAgent string) string {
	h := sha256.Sum256([]byte(userAgent))
	uaHash := hex.EncodeToString(h[:8]) // 取前 8 字节（16 个十六进制字符）足够区分
	return clientIP + "|" + uaHash
}

// Middleware 返回一个 Gin 中间件，对请求进行速率限制
func (rl *RateLimiter) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		userAgent := c.GetHeader("User-Agent")
		key := buildKey(clientIP, userAgent)

		now := time.Now()

		rl.mu.Lock()
		entry, exists := rl.entries[key]
		if exists && now.Sub(entry.lastAccess) < rl.interval {
			remaining := rl.interval - now.Sub(entry.lastAccess)
			rl.mu.Unlock()
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":       "请求过于频繁，请稍后再试",
				"retry_after": remaining.Seconds(),
			})
			c.Abort()
			return
		}

		// 记录/更新本次访问时间
		if !exists {
			rl.entries[key] = &rateLimitEntry{lastAccess: now}
		} else {
			entry.lastAccess = now
		}
		rl.mu.Unlock()

		c.Next()
	}
}

// cleanup 定期清理过期的限流条目，防止内存泄漏
func (rl *RateLimiter) cleanup(tick time.Duration) {
	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()
		rl.mu.Lock()
		for key, entry := range rl.entries {
			if now.Sub(entry.lastAccess) > rl.interval*2 {
				delete(rl.entries, key)
			}
		}
		rl.mu.Unlock()
	}
}
