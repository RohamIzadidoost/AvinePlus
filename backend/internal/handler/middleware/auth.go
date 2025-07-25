package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Simple auth via header "X-User-ID" and "X-User-Role"
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.GetHeader("X-User-ID")
		role := c.GetHeader("X-User-Role")
		if idStr == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("userID", uint(id))
		c.Set("role", role)
		c.Next()
	}
}
