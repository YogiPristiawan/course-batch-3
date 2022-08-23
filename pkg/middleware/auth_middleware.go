package middleware

import (
	"course/pkg/tokenize"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		if header == "" || !strings.HasPrefix(header, "Bearer") {
			c.AbortWithStatusJSON(401, map[string]string{
				"message": "unauthorize",
			})
		}

		token := strings.Split(header, " ")[1]

		claims, err := tokenize.DecodeJwt(token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"messag": err.Error(),
			})
		}
		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}
