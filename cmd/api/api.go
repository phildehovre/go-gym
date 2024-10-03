package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phildehovre/go-gym/services/admin"
	"github.com/phildehovre/go-gym/services/location"
	"github.com/phildehovre/go-gym/services/membership"
	"github.com/phildehovre/go-gym/services/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, store *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: store}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	subrouter := router.PathPrefix("/api/v1").Subrouter()

	store := user.NewStore(s.db)
	userHandler := user.NewHandler(store)
	userHandler.RegisterRoutes(subrouter)

	locationStore := location.NewStore(s.db)
	locationHandler := location.NewHandler(locationStore)
	locationHandler.RegisterRoutes(subrouter)

	membershipStore := membership.NewStore(s.db)
	membershipHandler := membership.NewHandler(membershipStore, store)
	membershipHandler.RegisterRoutes(subrouter)

	adminStore := admin.NewStore(s.db)
	adminHandler := admin.NewHandler(adminStore, store)
	adminHandler.RegisterRoutes(subrouter)

	err := http.ListenAndServe(s.addr, router)

	if err != nil {
		return err
	}

	fmt.Printf("Server running on port %v", s.addr)
	return nil
}
