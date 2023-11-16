package web

import (
	"SongUser/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewRouter `Gin Engine`을 초기화하고 반환한다
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", homeHandler)
	r.GET("/hashPassword", hashPasswordHandler)
	r.POST("/login", loginHandler)
	r.POST("/register", registerHandler)

	return r
}

// homeHandler GET("/") 엔드포인트
func homeHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the Chat Server",
	})
}

// hashPasswordHandler GET("/hashPassword") 엔드포인트
func hashPasswordHandler(c *gin.Context) {
	password := c.Query("password")
	hashPassword, err := auth.HashPassword(password)
	if err != nil {
		c.JSON(500, gin.H{
			"error":   "Error hashing password",
			"details": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":      "Password hashed successfully",
		"hashPassword": hashPassword,
	})
}

// loginHandler POST("/login") 엔드포인트
func loginHandler(c *gin.Context) {
	var cred LoginCredentials
	if err := c.ShouldBind(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !(cred.Id == "admin" && cred.Pw == "admin") {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

// registerHandler POST("/register") 엔드포인트
func registerHandler(c *gin.Context) {
	var cred RegisterCredentials
	if err := c.ShouldBind(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Register success"})
}
