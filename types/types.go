package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateUser(user User) error
	GetUsers() ([]*User, error)
	GetUserById(id int) (*User, error)
}

type LocationStore interface {
	CreateLocation(Location) error
	GetLocations() ([]*Location, error)
	GetLocationByName(string) (*Location, error)
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Password  string    `json:"password"`
	Email     string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type Location struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Address        string    `json:"address"`
	City           string    `json:"city"`
	State          string    `json:"state"`
	PostalCode     string    `json:"postal_code"`
	Country        string    `json:"country"`
	PhoneNumber    string    `json:"phone_number,omitempty"`
	Email          string    `json:"email,omitempty"`
	Capacity       int       `json:"capacity,omitempty"`
	OperatingHours string    `json:"operating_hours,omitempty"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type CreateLocationPayload struct {
	Name           string `json:"name"`
	Address        string `json:"address"`
	City           string `json:"city"`
	State          string `json:"state"`
	PostalCode     string `json:"postal_code"`
	Country        string `json:"country"`
	PhoneNumber    string `json:"phone_number,omitempty"`
	Email          string `json:"email,omitempty"`
	Capacity       int    `json:"capacity,omitempty"`
	OperatingHours string `json:"operating_hours,omitempty"`
	IsActive       bool   `json:"is_active"`
}
