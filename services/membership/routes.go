package membership

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/phildehovre/go-gym/services/auth"
	"github.com/phildehovre/go-gym/types"
	"github.com/phildehovre/go-gym/utils"
)

type Handler struct {
	store     types.MembershipStore
	userStore types.UserStore
}

const (
	active   = "Active"
	expired  = "Expired"
	canceled = "Canceled"
)

func NewHandler(store types.MembershipStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/membership", auth.WithJWTAuth(h.handleCreateMembership, h.userStore)).Methods(http.MethodPost)
	router.HandleFunc("/membership", auth.WithJWTAuth(h.handleGetMembership, h.userStore)).Methods(http.MethodGet)
	router.HandleFunc("/membership", auth.WithJWTAuth(h.handleUpdateMembership, h.userStore)).Methods(http.MethodPatch)
	router.HandleFunc("/membership", auth.WithJWTAuth(h.handleDeactivateMembership, h.userStore)).Methods(http.MethodDelete)
	router.HandleFunc("/membership/locations", auth.WithJWTAuth(h.handleGetMembershipLocations, h.userStore)).Methods(http.MethodGet)
	router.HandleFunc("/membership/renew", auth.WithJWTAuth(h.handleRenewMembership, h.userStore)).Methods(http.MethodPatch)
}

func (h *Handler) handleCreateMembership(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateMembershipPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	userId := auth.GetUserIDFromContext(r.Context())

	_, err := h.store.CreateMembership(types.Membership{
		UserID:         userId,
		MembershipType: payload.MembershipType,
		Status:         payload.Status,
		StartDate:      payload.StartDate,
		EndDate:        payload.EndDate,
	}, payload.LocationIDS)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

func (h *Handler) handleGetMembership(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIDFromContext(r.Context())
	membership, err := h.store.GetMembership(userId)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
	}
	utils.WriteJSON(w, http.StatusOK, membership)
}

func (h *Handler) handleUpdateMembership(w http.ResponseWriter, r *http.Request) {
	var membershipUpdate *types.Membership

	err := utils.ParseJSON(r, &membershipUpdate)

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	err = h.store.UpdateMembership(membershipUpdate)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, "update route")
}

func (h *Handler) handleGetMembershipLocations(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIDFromContext(r.Context())
	fmt.Println("******Fetching membership data... ******")
	membership, err := h.store.GetMembership(userId)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}
	fmt.Println("******Fetching locations... ******")
	membershipLocations, err := h.store.GetMembershipLocations(membership.ID)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, membershipLocations)
}

func (h *Handler) handleDeactivateMembership(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIDFromContext(r.Context())

	membership, err := h.store.GetMembership(userId)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	membership.EndDate = time.Now()
	membership.Status = canceled

	err = h.store.UpdateMembership(membership)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, fmt.Sprintln("membership sucessfully deactivated"))
}

func (h *Handler) handleRenewMembership(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIDFromContext(r.Context())
	membership, err := h.store.GetMembership(userId)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}
	var updatedMembership types.RenewMembershipPayload
	err = utils.ParseJSON(r, &updatedMembership)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	membership.Status = updatedMembership.Status
	membership.StartDate = updatedMembership.StartDate
	membership.EndDate = updatedMembership.EndDate

	err = h.store.RenewMembership(membership)
	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, nil)
}
