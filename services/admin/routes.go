package admin

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phildehovre/go-gym/services/auth"
	"github.com/phildehovre/go-gym/types"
	"github.com/phildehovre/go-gym/utils"
)

type Handler struct {
	store     Store
	userStore types.UserStore
}

func (h *Handler) NewHandler(store Store, userStore types.UserStore) *Handler {
	return &Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/admin/users", auth.AdminMiddleware(auth.WithJWTAuth(h.handleGetAllUsers, h.userStore)))
	router.HandleFunc("/admin/memberships", auth.AdminMiddleware(auth.WithJWTAuth(h.handleGetAllMemberships, h.userStore)))
}

func (h *Handler) handleGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.store.GetAllUsers()

	if err != nil {
		utils.WriteError(w, http.StatusNotFound, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, users)
}

func (h *Handler) handleGetAllMemberships(w http.ResponseWriter, r *http.Request) {
	memberships, err := h.store.GetAllMemberships()
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, err)
	}
	utils.WriteJSON(w, http.StatusOK, memberships)
}
