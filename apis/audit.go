package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hechuangfil-admin/models"
	"hechuangfil-admin/utils"
	"strings"
)

func AuditList(c *gin.Context) {
	var u models.Withdraw
	param := make(map[string]string)
	var re models.Response
	param["keyword"] = c.Request.FormValue("keyword")
	param["pageSize"] = c.DefaultQuery("pageSize","10")
	param["pageIndex"] = c.DefaultQuery("pageIndex","1")

	result, err:= u.AuditList(param)
	if err!=nil{
		c.JSON(http.StatusOK, re.ReturnError(400))
		return
	}

	res := utils.NewPageData(param,result)

	c.JSON(http.StatusOK, res)
}

func AuditLogList(c *gin.Context) {
	var u models.Withdraw
	param := make(map[string]string)
	var re models.Response
	param["keyword"] = c.Request.FormValue("keyword")
	param["pageSize"] = c.DefaultQuery("pageSize","10")
	param["pageIndex"] = c.DefaultQuery("pageIndex","1")

	result, err:= u.AuditLogList(param)
	if err!=nil{
		c.JSON(http.StatusOK, re.ReturnError(400))
		return
	}
	//pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageData(param,result)

	c.JSON(http.StatusOK, res)
}
func AuditByID(c *gin.Context) {
	var u models.Withdraw
	var err error
	var res models.Response
	withdraw_id := c.Request.FormValue("withdraw_id")

	res.Data , err = u.WithdrawById(withdraw_id,"")
	if err!=nil{
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}
func CheckAudit(c *gin.Context) {
	var data models.Withdraw
	param := make(map[string]string)
	param["withdraw_id"] = c.Request.FormValue("withdraw_id")
	param["status"] = c.Request.FormValue("status")
	var res models.Response

	err := data.CheckAudit(param)
	if err!=nil{
		res.Code = 0
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	res.Msg = "审核成功！"
	c.JSON(http.StatusOK, res.ReturnOK())
}

func CheckAudits(c *gin.Context) {
	var data models.Withdraw
	param := make(map[string]string)
	param["ids"] = c.Request.FormValue("ids")
	param["status"] = c.Request.FormValue("status")
	var res models.Response
	ids := strings.Split(param["ids"],",")
	for _,v := range ids{
		param["withdraw_id"] = v
		err := data.CheckAudit(param)
		if err!=nil{
			res.Code = 0
			res.Msg = err.Error()
			c.JSON(http.StatusOK, res.ReturnError(400))
			return
		}
	}

	res.Msg = "审核成功！"
	c.JSON(http.StatusOK, res.ReturnOK())
}