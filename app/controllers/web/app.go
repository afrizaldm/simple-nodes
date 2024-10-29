package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func App(c *gin.Context) {
	c.String(http.StatusOK, "ping pong")
}
