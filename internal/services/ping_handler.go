package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ca *CmdApp) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
