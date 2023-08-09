package Middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func RoleMiddleware(role string) gin.HandlerFunc {
	// return the HandlerFunc
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		onlyToken := tokenString[len("Bearer "):]
		token, _ := jwt.Parse(onlyToken, func(token *jwt.Token) (interface{}, error) {
			return privateKey, nil
		})

		var claims = token.Claims.(jwt.MapClaims)

		// validate the token
		if token.Valid == true {
			// check if the role matches
			if claims["role"] == "User" {
				// continue with the request
				c.Next()
			} else {
				// return a forbidden status
				c.AbortWithStatus(http.StatusForbidden)

			}
		} else {
			// return a forbidden status
			c.AbortWithStatus(http.StatusForbidden)
		}

	}

}

func UserMiddleware() gin.HandlerFunc {
	return RoleMiddleware("User")
}
