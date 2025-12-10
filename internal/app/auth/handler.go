package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service AuthService
}

func (h AuthHandler) Register(c *gin.Context) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	c.BindJSON(&req)
	err := h.Service.Register(req.Name, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (h AuthHandler) Login(c *gin.Context){
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	c.BindJSON(&req)

	user, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Save()
	c.JSON(http.StatusOK, gin.H{"message": "Login success"})
}


func (h AuthHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout Successful",
	})
}