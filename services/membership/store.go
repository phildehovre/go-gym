package membership

import (
	"database/sql"

	"github.com/phildehovre/go-gym/types"
)

type Store struct {
	db *sql.DB
}

func NewMembershipStore(db *sql.DB) *Store {
	return &Store{db: db}
}
func (s *Store) CreateMembership(membership *types.CreateMembershipPayload) error {
	return nil
}
