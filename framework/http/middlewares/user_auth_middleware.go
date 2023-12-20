package middlewares

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CreateAuthMiddleware(c *gin.Context) {
	// get bearer token from header
	token := strings.Split(c.GetHeader("Authorization"), "Bearer ")

	if len(token) != 2 {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	parsedToken, err := jwt.Parse(token[1], func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized",
		})
		fmt.Println(err)
		return

	}

	parsedTokenClaims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok || !parsedToken.Valid {
		c.AbortWithStatusJSON(401, gin.H{
			"message": "Unauthorized",
		})
		return

	}

	c.Set("user_id", parsedTokenClaims["id"])

	c.Next()

}
