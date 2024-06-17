package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const SessionKey = "session_id"

type SessionAuth struct{}

func (s *SessionAuth) Auth(c *gin.Context) {
	sessionID := c.GetHeader(SessionKey)
	// TODO: imp auth
	if sessionID == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "session id is null")
	}
	fmt.Println("session id: ", sessionID)
	c.Next()
}
