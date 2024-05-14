package notification

import (
	"errors"
	"modak/gateway"
)

type Notification interface {
	Send(typeMessage, userID, message string) error
}

func NewService() Notification {
	return Service{
		rateLimit: NewRateLimit(),
		gateway:   gateway.Gateway{},
	}
}

type Service struct {
	rateLimit RateLimit
	gateway   gateway.Gateway
}

func (s Service) Send(typeMessage, userID, message string) error {
	if !s.rateLimit.validateRateLimit(User(userID), TypeMessage(typeMessage)) {
		return errors.New("max rate limit achieve")
	}
	return s.gateway.Send(userID, message)
}
