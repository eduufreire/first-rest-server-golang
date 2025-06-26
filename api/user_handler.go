package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/eduufreire/rest-api-users/model"
	repository "github.com/eduufreire/rest-api-users/repository/user"
)

type Handler struct {
	repository *repository.UserRepository
}

func Init() *Handler {
	handler := Handler{}
	handler.repository = repository.New()
	return &handler
}

type RequestBody struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	body := RequestBody{}
	_ = json.NewDecoder(r.Body).Decode(&body)

	createUserEntity, err := model.ToUserEntity(body.Name, body.Birthday)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.repository.CreateUser(createUserEntity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	responseUserEntity := h.repository.GetUserById(id)
	userParsed, err := json.Marshal(responseUserEntity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userParsed)
}

func (h *Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	allUsers := h.repository.GetAllUsers()

	usersParsed, err := json.Marshal(allUsers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(usersParsed)
}

func (h *Handler) GetUserById(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	responseUserEntity := h.repository.GetUserById(id)
	if responseUserEntity.ID == 0 {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	userParsed, err := json.Marshal(responseUserEntity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(userParsed)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	responseUserEntity := h.repository.GetUserById(id)
	if responseUserEntity.ID == 0 {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	err = h.repository.DeleteUserById(id)
	if err != nil {
		http.Error(w, "it wasn't possible to delete the user", http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(204)
}

func (h *Handler) EditUser(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user := h.repository.GetUserById(id)
	if user.ID == 0 {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	body := RequestBody{}
	_ = json.NewDecoder(r.Body).Decode(&body)

	newData, err := model.ToUserEntity(body.Name, body.Birthday)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user.Age = newData.Age
	user.Name = newData.Name
	user.Birthday = newData.Birthday

	err = h.repository.UpdateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	w.WriteHeader(204)
}
