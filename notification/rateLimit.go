package notification

import "time"

type UserLimit struct {
	attempts    int
	lastAttempt time.Time
}

type Limiter struct {
	maxAttempts   int
	timeThreshold time.Duration
}

//go:generate mockery --name RateLimiter --inpackage --filename rateLimiter_mock.go
type RateLimiter interface {
	ValidateRateLimit(userID User, typeMessage TypeMessage) bool
}

type RateLimit struct {
	users        map[User]map[TypeMessage]UserLimit
	typeLimiters map[TypeMessage]Limiter
}

type User string
type TypeMessage string

func NewRateLimiter() RateLimiter {
	return &RateLimit{
		users: map[User]map[TypeMessage]UserLimit{},
		typeLimiters: map[TypeMessage]Limiter{
			"Status": {
				maxAttempts:   2,
				timeThreshold: time.Minute,
			},
			"News": {
				maxAttempts:   1,
				timeThreshold: time.Hour * 24,
			},
			"Marketing": {
				maxAttempts:   3,
				timeThreshold: time.Hour,
			},
			"Test": {
				maxAttempts:   3,
				timeThreshold: time.Second * 2,
			},
		},
	}
}

func (r *RateLimit) ValidateRateLimit(userID User, typeMessage TypeMessage) bool {
	limiter := r.typeLimiters[typeMessage]
	user := r.users[userID]
	if user == nil {
		user = map[TypeMessage]UserLimit{
			typeMessage: {
				attempts:    0,
				lastAttempt: time.Now(),
			},
		}
	}
	userLimit := user[typeMessage]
	duration := time.Now().Sub(userLimit.lastAttempt)
	isValid := true
	if duration < limiter.timeThreshold {
		userLimit.attempts += 1
		isValid = userLimit.attempts <= limiter.maxAttempts
	} else {
		userLimit.attempts = 1
		userLimit.lastAttempt = time.Now()
	}
	user[typeMessage] = userLimit
	r.users[userID] = user
	return isValid
}
