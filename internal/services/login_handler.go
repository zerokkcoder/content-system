package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/content-system/internal/dao"
	"golang.org/x/crypto/bcrypt"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRsp struct {
	SessionID string `json:"session_id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
}

func (ca *CmsApp) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	accountDao := dao.NewAccountDao(ca.db)
	account, err := accountDao.FirstByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "账号不存在",
		})
		return
	}
	// 密码校验
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "密码错误",
		})
		return
	}
	// 登录成功
	// 生成session

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": LoginRsp{
			SessionID: genetateSessionID(),
			Username:  account.Username,
			Nickname:  account.Nickname,
		},
	})
}

func genetateSessionID() string {
	// session id 生成

	// session id 持久化

	return "session-id"
}
