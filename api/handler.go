package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"zt_test/entity"
)

type UserService interface {
	CreateNewTable() error
	AddNewUser(user entity.User) error
	GetUser() (entity.User, error)
	ComputeHmac(message, secret string) (string, error)
	Increase(ctx context.Context, key string, value int64) (int64, error)
}

type Handler struct {
	service UserService
}

func NewHandler(s UserService) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) Incr(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var req entity.KeyValue
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sendJsonError(w, err, http.StatusBadRequest)
		return
	}
	res, err := h.service.Increase(ctx, req.Key, req.Value)
	if err != nil {
		sendJsonError(w, err, http.StatusInternalServerError)
		return
	}

	sendJson(w, map[string]int64{"value": res})
}

func (h *Handler) GenerateHash(w http.ResponseWriter, r *http.Request) {
	var req entity.Key
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sendJsonError(w, err, http.StatusBadRequest)
		return
	}

	hash, err := h.service.ComputeHmac(req.Text, req.Key)
	if err != nil {
		sendJsonError(w, err, http.StatusInternalServerError)
		return
	}
	sendJson(w, hash)
}

func (h *Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	var req entity.User
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		sendJsonError(w, err, http.StatusBadRequest)
		return
	}

	err = h.service.CreateNewTable()
	if err != nil {
		sendJsonError(w, err, http.StatusInternalServerError)
		return
	}

	err = h.service.AddNewUser(req)
	if err != nil {
		sendJsonError(w, err, http.StatusInternalServerError)
		return
	}

	user, err := h.service.GetUser()
	if err != nil {
		sendJsonError(w, err, http.StatusInternalServerError)
		return
	}

	sendJson(w, map[string]int64{"id": user.ID})
}

type jsonError struct {
	Error string `json:"error"`
}

func sendJson(w http.ResponseWriter, data any, code ...int) {
	w.Header().Set("Content-Type", "application/json")

	if len(code) > 0 {
		w.WriteHeader(code[0])
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		sendJsonError(w, err, http.StatusInternalServerError)
	}
}
func sendJsonError(w http.ResponseWriter, err error, code int) {
	log.Println(err)
	sendJson(w, jsonError{Error: err.Error()}, code)
}
