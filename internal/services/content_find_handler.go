package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/content-system/internal/dao"
)

type ContentFindReq struct {
	ID       int64  `json:"id"`
	Author   string `json:"author"`
	Title    string `json:"title"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

type ContentFindRsp struct {
	Message  string    `json:"message"`
	Contents []Content `json:"contents"`
	Total    int64     `json:"total"`
}

type Content struct {
	ID             int64         `json:"id"`
	Title          string        `json:"title"`
	VideoURL       string        `json:"video_url"`
	Author         string        `json:"author"`
	Description    string        `json:"description"`
	Thumbnail      string        `json:"thumbnail"`
	Category       string        `json:"category"`
	Duration       time.Duration `json:"duration"`
	Resolution     string        `json:"resolution"`
	FileSize       int64         `json:"file_size"`
	Format         string        `json:"format"`
	Quality        int           `json:"quality"`
	ApprovalStatus int           `json:"approval_status"`
}

func (ca *CmsApp) ContentFind(c *gin.Context) {
	var req ContentFindReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	contentDao := dao.NewContentDao(ca.db)
	contentList, total, err := contentDao.Find(&dao.FindParams{
		ID:       req.ID,
		Author:   req.Author,
		Title:    req.Title,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	contents := make([]Content, 0, len(contentList))
	for _, content := range contentList {
		contents = append(contents, Content{
			ID:             content.ID,
			Title:          content.Title,
			VideoURL:       content.VideoURL,
			Author:         content.Author,
			Description:    content.Description,
			Thumbnail:      content.Thumbnail,
			Category:       content.Category,
			Duration:       content.Duration,
			Resolution:     content.Resolution,
			FileSize:       content.FileSize,
			Format:         content.Format,
			Quality:        content.Quality,
			ApprovalStatus: content.ApprovalStatus,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": ContentFindRsp{
			Message:  "ok",
			Contents: contents,
			Total:    total,
		},
	})
}
