package middleware

import (
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"

	"github.com/lingojack/proj_template/config"
	"golang.org/x/time/rate"
)

type rateLimiter struct {
	limiters sync.Map
	rate     rate.Limit
	burst    int
}

func (rl *rateLimiter) getLimiter(key string) *rate.Limiter {
	if limiter, ok := rl.limiters.Load(key); ok {
		return limiter.(*rate.Limiter)
	}
	limiter := rate.NewLimiter(rl.rate, rl.burst)
	actual, _ := rl.limiters.LoadOrStore(key, limiter)
	return actual.(*rate.Limiter)
}

func RateLimit(cfg *config.Config) echo.MiddlewareFunc {
	if !cfg.Middleware.RateLimit.Enabled {
		return Passthrough()
	}

	rl := &rateLimiter{
		rate:  rate.Limit(cfg.Middleware.RateLimit.RequestsPerSecond),
		burst: cfg.Middleware.RateLimit.Burst,
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			key := c.RealIP()
			limiter := rl.getLimiter(key)

			if !limiter.Allow() {
				return c.JSON(http.StatusTooManyRequests, map[string]string{
					"message": "too many requests",
				})
			}
			return next(c)
		}
	}
}
