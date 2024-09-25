package location

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phildehovre/go-gym/types"
	"github.com/phildehovre/go-gym/utils"
)

type Handler struct {
	store types.LocationStore
}

func NewHandler(store types.LocationStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/location", h.handleCreateLocation).Methods("POST")
}

func (h *Handler) handleCreateLocation(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateLocationPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	err := h.store.CreateLocation(types.Location{
		Name:           payload.Name,
		Address:        payload.Address,
		City:           payload.City,
		State:          payload.State,
		PostalCode:     payload.PostalCode,
		Country:        payload.Country,
		PhoneNumber:    payload.PhoneNumber,
		Email:          payload.Email,
		Capacity:       payload.Capacity,
		OperatingHours: payload.OperatingHours,
		IsActive:       payload.IsActive,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, nil)
}
