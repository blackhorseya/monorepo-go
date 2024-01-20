package httpx

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Server is a http server.
type Server struct {
	httpserver *http.Server
	router     *gin.Engine
}
