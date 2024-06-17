package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/content-system/internal/services"
)

const (
	rootPath = "/api"
)

func CmsRouters(r *gin.Engine) {
	cmsApp := services.NewCmdApp()
	session := &SessionAuth{}
	root := r.Group(rootPath).Use(session.Auth)
	{
		// /api/cms/ping
		root.GET("/cms/ping", cmsApp.Ping)
	}
}
