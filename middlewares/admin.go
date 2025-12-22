package middlewares

import(
	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		if role != "admin" {
			c.AbortWithStatusJSON(403, gin.H{
				"message": "admin only",
			})
			return
		}
		c.Next()
	}
}
