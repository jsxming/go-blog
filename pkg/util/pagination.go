package util

import (
	"blog/global"
	"github.com/gin-gonic/gin"
)

type Pager struct {
	Page   int
	Size   int
	Offset int
}

func NewPager(c *gin.Context) Pager {
	page := GetPage(c)
	size := GetPageSize(c)
	offset := GetPageOffset(page, size)
	return Pager{
		Page:   page,
		Size:   size,
		Offset: offset,
	}
}

func GetPage(c *gin.Context) int {
	page := Int(c.Query("page"))
	if page <= 0 {
		page = 1
	}
	return page
}

func GetPageSize(c *gin.Context) int {
	size := Int(c.Query("size"))
	if size <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if size > global.AppSetting.MaxPageSize {
		return global.AppSetting.MaxPageSize
	}
	return size

}

func GetPageOffset(page, size int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * size
	}
	return result
}
