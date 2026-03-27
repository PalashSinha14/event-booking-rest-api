package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palashsinha14/go-rest-api/utils"
)
/*
func Authenticate(c *gin.Context) {
	//token := c.Request.Header.Get("Authorization")
	//Reading token as cookie
	token := c.GetHeader("Authorization")

	if token == "" {
		cookieToken, err := c.Cookie("token")
		if err == nil {
			token = cookieToken
		}
	}

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
}*/
/*
func AuthMiddlewareHTML(c *gin.Context) {
	token, err := c.Cookie("token")

	if err != nil || token == "" {
		c.Redirect(http.StatusSeeOther, "/login-page")
		c.Abort()
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login-page")
		c.Abort()
		return
	}

	c.Set("userId", userId)
	c.Next()
}*/
func Authenticate(c *gin.Context) {

	token, err := c.Cookie("token")
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login-page")
		c.Abort()
		return
	}

	userId, email, err := utils.VerifyToken(token)
	if err != nil {
		c.Redirect(http.StatusSeeOther, "/login-page")
		c.Abort()
		return
	}

	c.Set("userId", userId)
	c.Set("email", email)
	c.Next()
}