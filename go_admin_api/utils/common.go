package utils

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AdminResponse 管理后台响应结构
type AdminResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Count   int64       `json:"count"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// AdminSuccess 管理后台成功返回
func AdminSuccess(c *gin.Context, data interface{}, count int64) {
	c.JSON(200, AdminResponse{
		Code:    200,
		Message: "success",
		Data:    data,
		Count:   count,
	})
}

func OpensResponse(c *gin.Context, msg string) {
	c.JSON(200, Response{
		Code:    200,
		Message: msg,
	})
}

// AdminFailed 管理后台失败返回
func AdminFailed(c *gin.Context, message string) {
	c.JSON(500, AdminResponse{
		Code:    500,
		Message: message,
	})
}

func SuccessData(c *gin.Context, data interface{}) {
	c.JSON(200, AdminResponse{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
