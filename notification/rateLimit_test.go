package notification

import (
	"gotest.tools/v3/assert"
	"testing"
	"time"
)

func TestRateLimit_validateRateLimit(t *testing.T) {
	t.Parallel()
	t.Run("should pass rate limit", func(t *testing.T) {
		rateLimit := NewRateLimiter()
		for i := 0; i < 10; i++ {
			isValid := rateLimit.ValidateRateLimit("user1", "Test")
			assert.Check(t, isValid)
			time.Sleep((time.Second * 2) / 3)
		}
	})
	t.Run("should not pass rate limit", func(t *testing.T) {
		rateLimit := NewRateLimiter()
		for i := 0; i < 3; i++ {
			isValid := rateLimit.ValidateRateLimit("user1", "Test")
			assert.Check(t, isValid)
		}
		isValid := rateLimit.ValidateRateLimit("user1", "Test")
		assert.Check(t, !isValid)
	})
	t.Run("should reset rate limit", func(t *testing.T) {
		rateLimit := NewRateLimiter()
		for i := 0; i < 10; i++ {
			isValid := rateLimit.ValidateRateLimit("user1", "Test")
			assert.Check(t, isValid)
			isValid = rateLimit.ValidateRateLimit("user1", "Test")
			assert.Check(t, isValid)
			isValid = rateLimit.ValidateRateLimit("user1", "Test")
			assert.Check(t, isValid)
			time.Sleep(time.Second * 2)
		}
	})
	t.Run("should manage different limit rates per user", func(t *testing.T) {
		rateLimit := NewRateLimiter()
		for i := 0; i < 10; i++ {
			isValid := rateLimit.ValidateRateLimit("user1", "Test")
			assert.Check(t, isValid)
			isValid = rateLimit.ValidateRateLimit("user1", "Test")
			assert.Check(t, isValid)
			isValid = rateLimit.ValidateRateLimit("user1", "Test")
			assert.Check(t, isValid)
			isValid = rateLimit.ValidateRateLimit("user2", "Test")
			assert.Check(t, isValid)
			isValid = rateLimit.ValidateRateLimit("user2", "Test")
			assert.Check(t, isValid)
			isValid = rateLimit.ValidateRateLimit("user2", "Test")
			assert.Check(t, isValid)
			time.Sleep(time.Second * 2)
		}
	})
}
