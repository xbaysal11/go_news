package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "hllo world!")
}
func collectHandler(c *gin.Context) {
	c.String(http.StatusOK, "hllo world!")
}
func resultHandler(c *gin.Context) {
	c.String(http.StatusOK, "hllo world!")
}
