package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		r, _ := c.Get("role")
		if r != role {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}
