package middleware

import (
	"api-golang/service"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema = "Bearer "
		header := ctx.GetHeader("Authorization")
		if header == "" || !strings.HasPrefix(header, BearerSchema) {
			fmt.Println("Authorization header missing or invalid")
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Authorization header missing or invalid"})
			return
		}

		token := header[len(BearerSchema):]
		fmt.Println("Token extracted:", token)

		if !jwtService.Validate(token) {
			fmt.Println("Invalid token")
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}

		fmt.Println("Token valid, proceeding to next handler")
		ctx.Next()
	}
}
