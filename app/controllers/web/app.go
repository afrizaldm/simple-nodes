package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	// c.HTML(http.StatusOK, "index.html", nil)
	c.Redirect(http.StatusFound, "/app")
}

func App(c *gin.Context) {
	c.HTML(http.StatusOK, "app.html", nil)
}
