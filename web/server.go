package web

import (
	"SongUser/mongo"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// NewRouter `Gin Engine`을 초기화하고 반환한다
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", homeHandler)
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

// loginHandler POST("/login") 엔드포인트
func loginHandler(c *gin.Context) {
	var cred LoginCredentials
	if err := c.ShouldBind(&cred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	repo, err := mongo.GetUserInfoRepository()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't access to user information database"})
		return
	}

	err = mongo.Login(cred.Id, cred.Pw, repo)
	if err != nil {
		var userNotFoundError *mongo.UserNotFoundError
		if errors.As(err, &userNotFoundError) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		var passwordMismatchError *mongo.PasswordMismatchError
		if errors.As(err, &passwordMismatchError) {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Password mismatch"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
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

	repo, err := mongo.GetUserInfoRepository()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't access to user information database"})
		return
	}

	err = mongo.Register(cred.Id, cred.Pw, cred.Name, repo)
	var userAlreadyExistsError *mongo.UserAlreadyExistsError
	if errors.As(err, &userAlreadyExistsError) {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Can't access to user information database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Register success"})
}
