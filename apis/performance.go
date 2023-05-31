package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hechuangfil-admin/models"
	"hechuangfil-admin/pkg"
	"hechuangfil-admin/utils"
)

func IndividualPerformance(c *gin.Context) {
	var data models.Performance
	var re models.Response
	//err := c.BindWith(&data, binding.JSON)
	//pkg.AssertErr(err, "", 500)
	param := make(map[string]string)
	param["date"] = c.Request.FormValue("date")
	param["namenew"] = c.Request.FormValue("name")
	param["pageSize"] = c.DefaultQuery("pageSize","10")
	param["pageIndex"] = c.DefaultQuery("pageIndex","1")
	if param["date"]==""{
		c.JSON(http.StatusOK, re.ReturnError(400))
		return
	}
	param["date"] += "-01"
	//switch utils.GetRolekey(c) {
	//case "finance":
	//case "boss":
	//case "admin":
	//default:
	//	param["userid"] = utils.GetUserIdStr(c)
	//}
	param["userid"] = utils.GetUserIdStr(c)
	result, err := data.IndividualPerformance(param)
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageData(param,result)

	c.JSON(http.StatusOK, res)
}
