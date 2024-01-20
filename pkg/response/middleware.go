package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddErrorHandlingMiddleware is used to add error handling middleware.
func AddErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if c.Errors.Last() == nil {
				return
			}

			err := c.Errors.Last()

			switch e := err.Err.(type) {
			default:
				c.AbortWithStatusJSON(http.StatusInternalServerError, Err.WrapError(e))
			}
		}()

		c.Next()
	}
}
