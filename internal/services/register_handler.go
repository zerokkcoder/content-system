package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	// 密码加密
	// $2a$10$GiXgbsUv.3TS5Innkg6FgO0AfzDKfsdzpZiZjUZBd2/XaDCN/tI.y
	// $2a$10$lCim7/WvwQN7hCjRdVAVyeDkWiIVd//LvZJ3N69Zy6z39ULxsnGhK
	hashedPassword, err := encryptPassword(req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
	}
	// TODO: 账号校验
	// TODO: 账号信息持久化
	fmt.Printf("register req = %+v, hashedPassword = [%s]\n", req, hashedPassword)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": RegisterRsp{Message: "注册成功"},
	})
}

func encryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("bcrypt generate from password error = %v\n", err)
		return "", err
	}
	return string(hashedPassword), nil
}
