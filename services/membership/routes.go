package membership

import (
	"fmt"
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

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/membership", auth.WithJWTAuth(h.handleCreateMembership, h.userStore)).Methods("POST")
}

func (h *Handler) handleCreateMembership(w http.ResponseWriter, r *http.Request) {
	var payload types.CreateMembershipPayload

	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	userId := auth.GetUserIDFromContext(r.Context())
	fmt.Println(userId)

	if err := h.store.CreateMembership(types.Membership{
		UserID:         userId,
		MembershipType: payload.MembershipType,
		Status:         payload.Status,
		StartDate:      payload.StartDate,
		EndDate:        payload.EndDate,
	}); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
