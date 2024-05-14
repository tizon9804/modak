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

type RateLimit struct {
	users        map[User]map[TypeMessage]UserLimit
	typeLimiters map[TypeMessage]Limiter
}

type User string
type TypeMessage string

func NewRateLimit() RateLimit {
	return RateLimit{
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
		},
	}
}

func (r *RateLimit) validateRateLimit(userID User, typeMessage TypeMessage) bool {
	limiter := r.typeLimiters[typeMessage]
	user := r.users[userID]
	userLimit := user[typeMessage]
	duration := time.Now().Sub(userLimit.lastAttempt)
	isValid := true
	if duration < limiter.timeThreshold {
		userLimit.attempts += 1
		isValid = userLimit.attempts > limiter.maxAttempts
	} else {
		userLimit.attempts = 1
		userLimit.lastAttempt = time.Now()
	}

	user[typeMessage] = userLimit
	r.users[userID] = user
	return isValid
}
