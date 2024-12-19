package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func IPLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		log.Printf("Request from IP: %s", ip)
		c.Next()
	}
}
