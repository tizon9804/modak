package notification

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestService_Send(t *testing.T) {
	t.Run("should valid rate limit", func(t *testing.T) {
		rateLimiter := NewMockRateLimiter(t)
		service := NewService(rateLimiter)
		rateLimiter.On("ValidateRateLimit", User("user1"), TypeMessage("test")).Return(true)
		err := service.Send("test", "user1", "send message")
		assert.NilError(t, err)
	})

	t.Run("should not valid rate limit", func(t *testing.T) {
		rateLimiter := NewMockRateLimiter(t)
		service := NewService(rateLimiter)
		rateLimiter.On("ValidateRateLimit", User("user1"), TypeMessage("test")).Return(false)
		err := service.Send("test", "user1", "send message")
		assert.ErrorContains(t, err, "max rate limit achieve user")
	})
}
