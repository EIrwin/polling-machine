package auth

import (
	"encoding/json"
	"net/http"
)

const (
	// APIBase is the base path for API access
	APIBase = "/api/v1/"

	LoginPath = APIBase + "login"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json"token"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	service := NewService()

	token, err := service.Login(req.Email, req.Password)
	if err != nil {
		//w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(token); err != nil {
		panic(err)
	}
}
