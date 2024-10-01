package membership

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
func (s *Store) CreateMembership(membership types.Membership, locationIDS []int) (int, error) {
	// Start a transaction
	tx, err := s.db.Begin()
	if err != nil {
		return -1, err
	}

	// Create membership
	res, err := tx.Exec(`INSERT INTO memberships (
		user_id, 
		membership_type,
		status,
		start_date,
		end_date
	)
	VALUES (?,?,?,?,?)`,
		membership.UserID,
		membership.MembershipType,
		membership.Status,
		membership.StartDate,
		membership.EndDate,
	)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	membershipId, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	for _, loc := range locationIDS {
		err = s.CreateMembershipLocation(&types.MembershipLocation{LocationID: loc, MembershipID: int(membershipId)}, tx)
		if err != nil {
			tx.Rollback()
			return -1, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return int(membershipId), nil
}

func (s *Store) CreateMembershipLocation(ml *types.MembershipLocation, tx *sql.Tx) error {
	_, err := tx.Exec(`
	INSERT INTO membershipLocations (
		membership_id,
		location_id
	) VALUES (?,?)
	`, ml.MembershipID, ml.LocationID)

	return err
}

func (s *Store) GetMembership(userId int) (*types.Membership, error) {
	rows, err := s.db.Query(`SELECT * FROM memberships WHERE user_id = ?`, userId)
	if err != nil {
		return nil, err
	}

	membership := new(types.Membership)
	for rows.Next() {
		m, err := ScanRowsIntoMembership(rows)
		if err != nil {
			return nil, err
		}
		membership = m
	}

	return membership, nil
}

func (s *Store) GetMembershipLocations(mId int) ([]*types.Location, error) {
	var membershipLocations []*types.Location
	query := `SELECT ml.membership_id, l.*
		FROM membershipLocations ml
		JOIN locations l ON ml.location_id = l.id
		WHERE ml.membership_id = ?
		;`
	rows, err := s.db.Query(query, mId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		loc, err := scanRowsIntoMembershipLocation(rows)

		if err != nil {
			return nil, err
		}
		membershipLocations = append(membershipLocations, loc)
	}

	return membershipLocations, nil
}

func (s *Store) UpdateMembership(m *types.Membership) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	var exists int
	err = tx.QueryRow("SELECT COUNT(*) FROM memberships WHERE id = ?", m.ID).Scan(&exists)
	if err != nil || exists == 0 {
		tx.Rollback()
		return fmt.Errorf("membership with ID %d does not exist", m.ID)
	}
	fmt.Println("MEMBERSHIP: ", m)
	_, err = tx.Exec(`UPDATE memberships
	SET user_id = ?,
	membership_type = ?,
	status = ?,
	start_date = ?,
	end_date = ?
	WHERE id = ?`,
		m.UserID, m.MembershipType, m.Status, m.StartDate, m.EndDate, m.ID)

	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func ScanRowsIntoMembership(rows *sql.Rows) (*types.Membership, error) {
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

func scanRowsIntoMembershipLocation(rows *sql.Rows) (*types.Location, error) {
	ml := new(types.MembershipLocation)
	loc := new(types.Location)

	err := rows.Scan(
		&ml.MembershipID,
		&loc.ID,
		&loc.Name,
		&loc.Address,
		&loc.City,
		&loc.State,
		&loc.PostalCode,
		&loc.Country,
		&loc.PhoneNumber,
		&loc.Email,
		&loc.Capacity,
		&loc.OperatingHours,
		&loc.IsActive,
		&loc.CreatedAt,
		&loc.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return loc, nil
}

func (s *Store) DeleteMembership(id int) error {
	_, err := s.db.Exec(`DELETE FROM memberships WHERE id = ? `, id)
	return err
}
