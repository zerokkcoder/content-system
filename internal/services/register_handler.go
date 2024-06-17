package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}

type RegisterRsp struct {
	Message string `json:"message"`
}

func (ca *CmdApp) Register(c *gin.Context) {
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// TODO：密码加密
	// TODO: 账号校验
	// TODO: 账号信息持久化  
	fmt.Printf("register req = %+v", req)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": RegisterRsp{Message: "注册成功"},
	})
}
