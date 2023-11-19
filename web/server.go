package web

import (
	"SongUser/mongo"
	"github.com/gin-gonic/gin"
	"log"
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

	repo, err := mongo.NewUserRepository("user", "userinfo")
	if err != nil {
		log.Printf("Error creating new user repository: %+v", err)
		return
	}

	err = mongo.Login(cred.Id, cred.Pw, repo)
	if err != nil {
		if err.Error() == "user not found" || err.Error() == "password mismatch" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
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

	mongo.Register(cred.Id, cred.Pw, cred.Name)

	c.JSON(http.StatusOK, gin.H{"message": "Register success"})
}
