package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
	"github.com/kodah/blog/service"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer"

		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BearerSchema):]

		session := sessions.Default(c)
		raw := session.Get("token")
		if raw != nil && raw != "" {
			tokenString = raw.(string)
		}

		token, err := service.JWTAuthService().ValidateToken(tokenString)
		if token == nil || err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
