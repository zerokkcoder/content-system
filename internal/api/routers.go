package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/content-system/internal/services"
)

const (
	rootPath   = "/api"
	noAuthPath = "/out/api"
)

func CmsRouters(r *gin.Engine) {
	cmsApp := services.NewCmsApp()
	session := &SessionAuth{}
	root := r.Group(rootPath).Use(session.Auth)
	{
		// /api/cms/ping
		root.GET("/cms/ping", cmsApp.Ping)
	}

	noAuth := r.Group(noAuthPath)
	{
		// /out/api/cms/register
		noAuth.POST("/cms/register", cmsApp.Register)
		// /out/api/cms/login
		noAuth.POST("/cms/login", cmsApp.Login)
	}
}
