package location

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

func (s *Store) CreateLocation(location types.Location) error {
	_, err := s.db.Exec(`INSERT INTO locations (
    name, 
    address, 
    city, 
    state, 
    postal_code, 
    country, 
    phone_number, 
    email, 
    capacity, 
    operating_hours, 
    is_active
) 
VALUES (?,?,?,?,?,?,?,?,?,?,?)`,
		location.Name,
		location.Address,
		location.City,
		location.State,
		location.PostalCode,
		location.Country,
		location.PhoneNumber,
		location.Email,
		location.Capacity,
		location.OperatingHours,
		location.IsActive,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetLocations() ([]*types.Location, error) {
	rows, err := s.db.Query("SELECT * FROM locations")
	if err != nil {
		return nil, err
	}

	var locations []*types.Location

	for rows.Next() {
		location, err := scanRowsIntoLocation(rows)
		if err != nil {
			return nil, err
		}
		locations = append(locations, location)
	}
	return locations, nil

}

func scanRowsIntoLocation(rows *sql.Rows) (*types.Location, error) {
	location := new(types.Location)

	err := rows.Scan(
		&location.ID,
		&location.Name,
		&location.Address,
		&location.City,
		&location.State,
		&location.PostalCode,
		&location.Country,
		&location.PhoneNumber,
		&location.Email,
		&location.Capacity,
		&location.OperatingHours,
		&location.IsActive,
		&location.CreatedAt,
		&location.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return location, nil
}

func (s *Store) GetLocationByName(name string) (*types.Location, error) {
	rows, err := s.db.Query(`SELECT * FROM locations WHERE name = ?`, name)
	if err != nil {
		return nil, err
	}

	var loc *types.Location

	for rows.Next() {
		location, err := scanRowsIntoLocation(rows)

		if err != nil {
			return nil, err
		}
		loc = location
	}
	return loc, nil
}

func (s *Store) GetLocationByID(id int) (*types.Location, error) {
	rows, err := s.db.Query(`SELECT * FROM locations WHERE id = ?`, id)
	if err != nil {
		return nil, err
	}

	var loc *types.Location

	for rows.Next() {
		location, err := scanRowsIntoLocation(rows)

		if err != nil {
			return nil, err
		}
		loc = location
	}
	return loc, nil
}

func (s *Store) GetLocationsByKey(key string, value string) ([]*types.Location, error) {
	var locations []*types.Location

	rows, err := s.db.Query(fmt.Sprintf(`SELECT * FROM locations WHERE %s = ?`, key), value)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		location, err := scanRowsIntoLocation(rows)
		if err != nil {
			return nil, err
		}
		locations = append(locations, location)
	}

	return locations, nil
}
