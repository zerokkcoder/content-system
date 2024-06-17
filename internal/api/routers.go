package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/content-system/internal/services"
)

const (
	rootPath = "/api"
)

func CmsRouters(r *gin.Engine) {
	root := r.Group(rootPath)
	cmsApp := services.NewCmdApp()
	{
		// /api/cms/ping
		root.GET("/cms/ping", cmsApp.Ping)
	}
}
