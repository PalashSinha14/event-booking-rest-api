package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/palashsinha14/go-rest-api/models"
	"github.com/palashsinha14/go-rest-api/utils"
)

// signup godoc
// @Summary      Create a new account
// @Description  Register a new user with an email and password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      models.User  true  "Email and password"
// @Success      201          {object}  map[string]string
// @Failure      400          {object}  map[string]string
// @Failure      500          {object}  map[string]string
// @Router       /signup [post]
func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successsfully!"})
}

// login godoc
// @Summary      Log in
// @Description  Authenticate with email and password. On success, sets an httpOnly JWT cookie and redirects to /dashboard - this endpoint is built for the browser login form, not a JSON API client.
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body  models.User  true  "Email and password"
// @Success      303  "Redirects to /dashboard, JWT set as an httpOnly cookie"
// @Failure      400  {object}  map[string]string
// @Failure      401  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /login [post]
func login(c *gin.Context) {
	var user models.User

	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not login"})
		return
	}

	// Set cookie
	c.SetCookie("token", token, 3600, "/", "", true, true)

	// Redirect to dashboard
	c.Redirect(http.StatusSeeOther, "/dashboard")
}

/*func login(c *gin.Context) {
	var user models.User

	err := c.ShouldBindWith(&user, binding.Form) // works with form-data now
	if err != nil {
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": "Invalid input",
		})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "login.html", gin.H{
			"error": "Login failed",
		})
		return
	}

	// ✅ Set cookie
	c.SetCookie(
		"token",
		token,
		3600*2,
		"/",
		"",
		false,
		true,
	)

	// ✅ Redirect to dashboard
	c.Redirect(http.StatusSeeOther, "/dashboard")
}*/

/*
func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	//Setting jwt as cookie intead of pop up and manually copying jwt
	c.SetCookie(
		"token", // cookie name
		token,   // value
		3600*2,  // max age (2 hours)
		"/",     // path
		"",      // domain (empty = current)
		false,   // secure (true in HTTPS)
		true,    // httpOnly (important for security)
	)

	//c.JSON(http.StatusOK, gin.H{
	//	"message": "Login successful!",
	//})

	c.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})
}
*/
