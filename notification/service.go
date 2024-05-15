package notification

import (
	"errors"
	"fmt"
	"modak/gateway"
)

type Notification interface {
	Send(typeMessage, userID, message string) error
}

func NewService(limiter RateLimiter) Notification {
	return Service{
		rateLimit: limiter,
		gateway:   gateway.Gateway{},
	}
}

type Service struct {
	rateLimit RateLimiter
	gateway   gateway.Gateway
}

func (s Service) Send(typeMessage, userID, message string) error {
	if !s.rateLimit.ValidateRateLimit(User(userID), TypeMessage(typeMessage)) {
		return errors.New(fmt.Sprintf("max rate limit achieve user: %s type:%s", userID, typeMessage))
	}
	return s.gateway.Send(userID, message)
}
