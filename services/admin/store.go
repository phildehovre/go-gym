package admin

import (
	"database/sql"

	"github.com/phildehovre/go-gym/services/membership"
	"github.com/phildehovre/go-gym/services/user"
	"github.com/phildehovre/go-gym/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetAllMemberships() ([]*types.Membership, error) {
	query := `SELECT * FROM memberships`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	var memberships []*types.Membership
	for rows.Next() {
		m, err := membership.ScanRowsIntoMembership(rows)
		if err != nil {
			return nil, err
		}
		memberships = append(memberships, m)
	}
	return memberships, nil
}

func (s *Store) GetAllUsers() ([]*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	var users []*types.User

	for rows.Next() {
		u, err := user.ScanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (s *Store) UpdateUserRole(uid int) error {
	// todo
	return nil
}
