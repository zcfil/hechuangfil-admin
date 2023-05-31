package utils

import (
	"github.com/gin-gonic/gin"
	"hechuangfil-admin/pkg"
)

func GetParam(c *gin.Context)(param map[string]interface{},err error){
	var pageSize = 10
	var pageIndex = 1
	param = make(map[string]interface{})
	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = pkg.StrToInt(err, size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = pkg.StrToInt(err, index)
	}
	param["pageIndex"] = pageIndex
	param["pageSize"] = pageSize
	return
}
