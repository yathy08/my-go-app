package handler

import (
    "encoding/json"
    "net/http"
    "github.com/yourusername/my-go-app/internal/service"
)

type AuthHandler struct {
    userService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
    return &AuthHandler{userService: userService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    var creds struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    token, err := h.userService.Login(creds.Username, creds.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    w.Header().Set("Authorization", "Bearer "+token)
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
