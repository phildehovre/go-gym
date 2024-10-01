package types

import (
	"database/sql"
	"time"
)

type MembershipStore interface {
	CreateMembership(Membership, []int) (int, error)
	GetMembership(int) (*Membership, error)
	CreateMembershipLocation(*MembershipLocation, *sql.Tx) error
	GetMembershipLocations(int) ([]*Location, error)
	UpdateMembership(*Membership) error

	// New methods
	DeleteMembership(int) error
	RenewMembership(int) error
}

type LocationStore interface {
	CreateLocation(Location) error
	GetLocations() ([]*Location, error)
	GetLocationByName(string) (*Location, error)
	GetLocationsByKey(string, string) ([]*Location, error)
	GetLocationByID(int) (*Location, error)

	// New methods
	// DeleteLocation(int) error
	// UpdateLocation(Location) error
	// GetLocationsByMembershipID(int) ([]*Location, error)
}

type UserStore interface {
	// Existing methods
	GetUserByEmail(email string) (*User, error)
	CreateUser(user User) error
	GetUsers() ([]*User, error)
	GetUserById(id int) (*User, error)

	// New methods
	// DeleteUser(int) error
	// UpdateUser(User) error
	// GetUsersByMembershipID(int) ([]*User, error)
}

type AdminStore interface {
	GetAllUsers() ([]*User, error)
	UpdateUserRole(userID int, roleID int) error
	GetAllMemberships() []*Membership
	CancelMembership(userID int) error
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	RoleID    int       `json:"role_id"`
	CreatedAt time.Time `json:"createdAt"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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

type Membership struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	MembershipType string    `json:"membership_type"`
	Status         string    `json:"status"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	LocationIDS    []int     `json:"location_ids"`
}

type MembershipLocation struct {
	MembershipID int `json:"membership_id"`
	LocationID   int `json:"location_id"`
}

type UserMembershipLocations struct {
	ID         int        `json:"id"`
	FirstName  string     `json:"firstName"`
	LastName   string     `json:"lastName"`
	Email      string     `json:"email"`
	CreatedAt  time.Time  `json:"createdAt"`
	Locations  []Location `json:"locations"`
	Membership Membership `json:"membership"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"`
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" validate:"required,email"`
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

type CreateMembershipPayload struct {
	UserID         int       `json:"user_id"`
	MembershipType string    `json:"membership_type"`
	Status         string    `json:"status"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	LocationIDS    []int     `json:"location_id"`
}
