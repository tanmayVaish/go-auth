package controllers

import (
	"go-auth/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_secret_key")

func Login(c *gin.Context) {
	var credential models.Credential

	if err := c.ShouldBindJSON(&credential); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var expectedPassword models.Credential

	models.DB.Where("username = ?", credential.Username).First(&expectedPassword)

	if expectedPassword.Password != credential.Password {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &models.Claims{
		Username: credential.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}

	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
}

func Signup(c *gin.Context) {

}
