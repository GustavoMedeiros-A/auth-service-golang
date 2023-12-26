package users

import (
	"auth-service/initializers"
	"auth-service/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// Get the email/password of req body
	var body struct {
		Email    string
		password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Respond
	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context) {
	// get the email and password of req body

	var body struct {
		Email    string
		password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Look up requested user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Invalid user email or password",
		})

		return
	}

	// Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.password))

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Invalid user email or password",
		})

		return
	}
	// Generate a jwt token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	fmt.Println("env token secret", os.Getenv("SECRET"))
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "fail to create token",
		})
		return
	}

	// send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"user":  user.Email,
		"token": tokenString,
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in my friend",
		"user":    user,
	})
}
