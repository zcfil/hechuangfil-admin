package apis

import (
	"github.com/gin-gonic/gin"
	"hechuangfil-admin/models"
	"hechuangfil-admin/utils"
	"net/http"
)

func StatementSummary(c *gin.Context) {
	var u models.Summary
	param := make(map[string]string)
	var re models.Response
	//param["keyword"] = c.Request.FormValue("keyword")
	param["pageSize"] = c.DefaultQuery("pageSize","10")
	param["pageIndex"] = c.DefaultQuery("pageIndex","1")
	param["start"] = c.Request.FormValue("start")
	param["end"] = c.Request.FormValue("end")
	//if param["end"]==""||param["start"]==""{
	//	c.JSON(http.StatusOK, re.ReturnError(400))
	//	return
	//}
	param["end"] += " 23:59:59"
	var result interface{}
	var err error
	result, err = u.StatementSummary(param)
	if err!=nil{
		c.JSON(http.StatusOK, re.ReturnError(400))
		return
	}
	//pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageData(param,result)

	c.JSON(http.StatusOK, res)
}
