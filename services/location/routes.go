package location

import (
	"fmt"
	"net/http"
	"strconv"

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
	router.HandleFunc("/location", h.handleGetLocations).Methods("GET")
	router.HandleFunc("/location/{id}", h.handleGetLocationByID).Methods("GET")
}

func (h *Handler) handleCreateLocation(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateLocationPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	locations, _ := h.store.GetLocationsByKey("name", payload.Name)
	if len(locations) > 0 {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("a location with this name already exists: %s", payload.Name))
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

func (h *Handler) handleGetLocations(w http.ResponseWriter, r *http.Request) {
	locations, err := h.store.GetLocations()
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, locations)

}

func (h *Handler) handleGetLocationByName(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleGetLocationByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString, ok := vars["id"]
	id, _ := strconv.Atoi(idString)

	fmt.Printf("ID: %d", id)
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not a valid param %v", id))
	}

	location, err := h.store.GetLocationByID(id)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, location)

}
