package middleware

import (
    "net/http"
    "strings"
    "github.com/yourusername/my-go-app/pkg/utils"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "missing token", http.StatusUnauthorized)
            return
        }

        token := strings.TrimPrefix(authHeader, "Bearer ")
        _, err := utils.ParseJWT(token)
        if err != nil {
            http.Error(w, "invalid token", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}
