package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func CoresMiddleware() gin.HandlerFunc {
	cor := cors.Default()

	return cor
}
