package membership

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phildehovre/go-gym/services/auth"
	"github.com/phildehovre/go-gym/types"
	"github.com/phildehovre/go-gym/utils"
)

type Handler struct {
	store     types.MembershipStore
	userStore types.UserStore
}

func NewHandler(store types.MembershipStore, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/membership", auth.WithJWTAuth(h.handleCreateMembership, h.userStore)).Methods("POST")
	router.HandleFunc("/membership", auth.WithJWTAuth(h.handleGetMembership, h.userStore)).Methods("GET")
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
