package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palashsinha14/go-rest-api/utils"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized."})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized."})
		return
	}
	c.Set("userId", userId)
	c.Next()
}
