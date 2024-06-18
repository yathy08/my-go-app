package main

import (
    "log"
    "net/http"
    "github.com/yourusername/my-go-app/internal/handler"
    "github.com/yourusername/my-go-app/internal/repository"
    "github.com/yourusername/my-go-app/internal/service"
    "github.com/yourusername/my-go-app/internal/middleware"
)

func main() {
    userRepo := repository.NewInMemoryUserRepository()
    userService := service.NewUserService(userRepo)
    authHandler := handler.NewAuthHandler(userService)
    adminHandler := handler.NewAdminHandler()

    mux := http.NewServeMux()
    mux.HandleFunc("/login", authHandler.Login)
    mux.Handle("/admin", middleware.JWTAuthMiddleware(http.HandlerFunc(adminHandler.Dashboard)))

    log.Println("Server running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
