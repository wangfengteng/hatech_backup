package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NoRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": "request not found"})
}
