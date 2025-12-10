package router

import (
	"github.com/Laanaa/my-app/internal/app/auth"
	"github.com/Laanaa/my-app/internal/app/auth/middleware"
	"github.com/Laanaa/my-app/internal/app/user"
	"github.com/gin-gonic/gin"
)


func Router() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.SessionMiddleware())

	authService := auth.AuthService{
		Repo: user.UserRepository{},
	}

	authHandler := auth.AuthHandler{
		Service: authService,
	}

	auth := r.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout", authHandler.Logout)
	}

	return r
}