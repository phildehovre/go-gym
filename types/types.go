package types

import "time"

type User struct {
	ID        int
	FirstName string
	LastName  string
	CreatedAt time.Time
}

type UserStore interface {
	GetUserByEmail(email string) *User
}

type LoginUserPayload struct {
	Email    string
	Password string
}
