package contextx

import (
	"github.com/gin-gonic/gin"
)

// AddContextxMiddleware is used to add contextx middleware.
func AddContextxMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("ctx", Background())

		c.Next()
	}
}
