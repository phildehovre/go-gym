package membership

import (
	"database/sql"

	"github.com/phildehovre/go-gym/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}
func (s *Store) CreateMembership(membership types.Membership) error {
	_, err := s.db.Exec(`INSERT INTO memberships (
	user_id, 
	membership_type,
	status,
	start_date,
	end_date
)
	VALUES (?,?,?,?,?)
	`,
		membership.UserID,
		membership.MembershipType,
		membership.Status,
		membership.StartDate,
		membership.EndDate,
	)
	return err
}

func (s *Store) GetMembership(userId int) (*types.Membership, error) {
	rows, err := s.db.Query(`SELECT * FROM memberships WHERE user_id = ?`, userId)
	if err != nil {
		return nil, err
	}

	membership := new(types.Membership)
	for rows.Next() {
		m, err := scanRowsIntoMembership(rows)
		if err != nil {
			return nil, err
		}
		membership = m
	}

	return membership, nil

}

func scanRowsIntoMembership(rows *sql.Rows) (*types.Membership, error) {
	membership := new(types.Membership)

	err := rows.Scan(
		&membership.ID,
		&membership.UserID,
		&membership.MembershipType,
		&membership.Status,
		&membership.StartDate,
		&membership.EndDate,
		&membership.CreatedAt,
		&membership.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return membership, nil
}
