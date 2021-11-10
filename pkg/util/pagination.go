package util

import (
	"ZhangLin/pkg/setting"

	"github.com/unknwon/com"

	"github.com/gin-gonic/gin"
)

// com是一个工具包
func GetPage(c *gin.Context) int {
	res := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		res = (page - 1) * setting.PageSize
	}
	return res
}
