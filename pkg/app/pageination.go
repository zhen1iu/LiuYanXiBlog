package app

import (
	"LiuYanXiBlog/pkg/convert"

	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context)int{
	page:=convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 { 
		return 1
	 }
	 return page
}

func GetPageSize(c *gin.Context)int{
	pageSize:=convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 { 
		return global.
	 }
	 return pageSize
}

