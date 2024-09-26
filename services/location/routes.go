package location

import (
	"fmt"
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
	router.HandleFunc("/location", h.handleGetLocations).Methods("GET")
	router.HandleFunc("/location/:name", h.handleGetLocationByName).Methods("GET")
}

func (h *Handler) handleCreateLocation(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateLocationPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	//todo
	// check if gym by the same name already exists
	// _, err := h.store.GetLocationByName(payload.Name)
	// if err == nil {
	// 	utils.WriteError(w, http.StatusBadRequest, err)
	// 	return
	// }

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
	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not a valid param %v", name))
	}

	location, err := h.store.GetLocationByName(name)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}
	utils.WriteJSON(w, http.StatusFound, location)

}
