package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// bucket is a hand-rolled token-bucket rate limiter for one client: it
// holds up to `burst` tokens, refilling at `ratePerSec` tokens per
// second. Every allowed request consumes one token.
type bucket struct {
	mu         sync.Mutex
	tokens     float64
	ratePerSec float64
	burst      float64
	lastCheck  time.Time
	lastSeen   time.Time
}

func newBucket(ratePerSec float64, burst int) *bucket {
	now := time.Now()
	return &bucket{
		tokens:     float64(burst),
		ratePerSec: ratePerSec,
		burst:      float64(burst),
		lastCheck:  now,
		lastSeen:   now,
	}
}

// allow reports whether one request may proceed right now, consuming a
// token if so. Tokens are refilled lazily based on how much time has
// passed since the last check, rather than on a background timer.
func (b *bucket) allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(b.lastCheck).Seconds()
	b.lastCheck = now
	b.lastSeen = now

	b.tokens += elapsed * b.ratePerSec
	if b.tokens > b.burst {
		b.tokens = b.burst
	}

	if b.tokens < 1 {
		return false
	}
	b.tokens--
	return true
}

var (
	buckets   = make(map[string]*bucket)
	bucketsMu sync.Mutex
)

// getBucket returns the token bucket for the given client IP, creating
// one (with limit ratePerSec and the given burst) the first time that IP
// is seen.
func getBucket(ip string, ratePerSec float64, burst int) *bucket {
	bucketsMu.Lock()
	defer bucketsMu.Unlock()

	b, exists := buckets[ip]
	if !exists {
		b = newBucket(ratePerSec, burst)
		buckets[ip] = b
	}
	return b
}

// StartRateLimiterCleanup periodically forgets IPs that haven't made a
// request recently, so the in-memory bucket map doesn't grow forever.
// Call this once at startup.
func StartRateLimiterCleanup() {
	go func() {
		for {
			time.Sleep(time.Minute)
			bucketsMu.Lock()
			for ip, b := range buckets {
				b.mu.Lock()
				idle := time.Since(b.lastSeen) > 3*time.Minute
				b.mu.Unlock()
				if idle {
					delete(buckets, ip)
				}
			}
			bucketsMu.Unlock()
		}
	}()
}

// RateLimit returns a Gin middleware that allows at most ratePerSec
// requests per second per client IP (with a burst of `burst`), responding
// 429 Too Many Requests beyond that. Intended for brute-force-sensitive
// endpoints like login and signup.
func RateLimit(ratePerSec float64, burst int) gin.HandlerFunc {
	return func(c *gin.Context) {
		b := getBucket(c.ClientIP(), ratePerSec, burst)
		if !b.allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"message": "Too many requests. Please try again later.",
			})
			return
		}
		c.Next()
	}
}

// AuthRateLimit is a pre-configured RateLimit for the login/signup
// endpoints: 1 request/sec sustained per IP, with a burst of 5 (so a
// normal user retrying a mistyped password isn't blocked, but a rapid
// brute-force loop is throttled hard).
func AuthRateLimit() gin.HandlerFunc {
	return RateLimit(1, 5)
}
