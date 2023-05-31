package apis

import (
	"github.com/gin-gonic/gin"
	"hechuangfil-admin/models"
	"hechuangfil-admin/utils"
	"net/http"
)

func GetSettlementList(c *gin.Context) {
	var u models.SettleLog
	param := make(map[string]string)
	param["pageSize"] = c.DefaultQuery("pageSize","10")
	param["pageIndex"] = c.DefaultQuery("pageIndex","1")
	param["keyword"] = c.Request.FormValue("keyword")

	var re models.Response
	dataList , err := u.SettleLogList(param)
	if err != nil {
		re.Msg = err.Error()
		c.JSON(http.StatusOK, re.ReturnError(400))
		c.Abort()
		return
	}
	res := utils.NewPageData(param,dataList)

	c.JSON(http.StatusOK,res)
}