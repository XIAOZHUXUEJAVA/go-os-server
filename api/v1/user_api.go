package v1

import (
	"net/http"
	"oschina/model"

	"github.com/gin-gonic/gin"
)

// LoginRequest 登录请求参数
type LoginRequest struct {
	Username string `json:"user_name" binding:"required"`
	Password string `json:"user_password" binding:"required"`
}

// LoginResponse 登录响应参数
type LoginResponse struct {
	Message string `json:"message"`
}

// LoginHandler 用户登录处理函数
func LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.GetUserByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil || !model.CheckUserPassword(user, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Login Failed"})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		Message: "Login Success",
	})
}

// RegisterRequest 注册请求参数
type RegisterRequest struct {
	Username string `json:"user_name" binding:"required"`
	Password string `json:"user_password" binding:"required"`
}

// RegisterHandler 用户注册处理函数
func RegisterHandler(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := model.GetUserByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Username already exists"})
		return
	}

	newUser := model.User{
		Username: req.Username,
		Password: req.Password,
	}

	err = model.CreateUser(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}
