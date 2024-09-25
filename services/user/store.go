package user

import (
	"database/sql"
	"fmt"

	"github.com/phildehovre/go-gym/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query(`SELECT * FROM users WHERE email = ?`, email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found %s", email)
	}

	return u, nil
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (firstName, lastName, email, password) VALUES (?,?,?,?)", user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetUsers() ([]*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	users := []*types.User{}
	for rows.Next() {
		u, err := scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
		if u.ID == 0 {
			return nil, fmt.Errorf("user not found with ID %v", u.ID)
		}
		users = append(users, u)
	}
	return users, nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	rows, err := s.db.Query(`SELECT * WHERE id==?`, id)
	if err != nil {
		return nil, err
	}

	var user *types.User
	for rows.Next() {
		u, err := scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
		if u.ID == 0 {
			return nil, fmt.Errorf("not user found with id %d", id)
		}
		user = u
	}

	return user, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
