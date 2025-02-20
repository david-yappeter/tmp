package jwt

import (
	"time"
)

type Payload struct {
	UserId    string
	CreatedAt time.Time
	ExpiredAt time.Time
}
