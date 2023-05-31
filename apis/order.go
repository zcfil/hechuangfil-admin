package apis

import (
	"github.com/gin-gonic/gin"
	"hechuangfil-admin/models"
	"hechuangfil-admin/result"
	"hechuangfil-admin/utils"
	"net/http"
)

// 充值列表
func GetRechargeList(c *gin.Context) {
	var u models.Recharge
	param := make(map[string]string)
	param["pageSize"] = c.DefaultQuery("pageSize","10")
	param["pageIndex"] = c.DefaultQuery("pageIndex","1")
	param["user_id"] = utils.GetUserIdStr(c)
	param["keyword"] = c.Request.FormValue("keyword")

	var re models.Response
	dataList , err := u.RechargeList(param)
	if err != nil {
		re.Msg = err.Error()
		c.JSON(http.StatusOK, re.ReturnError(400))
		c.Abort()
		return
	}
	res := utils.NewPageData(param,dataList)

	c.JSON(http.StatusOK,res)
}

//func GetWithdrawList(c *gin.Context) {
//	var u models.Recharge
//	param := make(map[string]string)
//	param["pageSize"] = c.DefaultQuery("pageSize","10")
//	param["pageIndex"] = c.DefaultQuery("pageIndex","1")
//	param["user_id"] = utils.GetUserIdStr(c)
//	var re models.Response
//	dataList , err := u.RechargeList(param)
//	if err != nil {
//		re.Msg = err.Error()
//		c.JSON(http.StatusOK, re.ReturnError(400))
//		c.Abort()
//		return
//	}
//	res := utils.NewPageData(param,dataList)
//
//	c.JSON(http.StatusOK,res)
//}

func GetOrders(c *gin.Context) {
	pageSize := c.Request.FormValue("pageSize")
	if pageSize == "" {
		c.JSON(http.StatusOK, result.Failstr("pageSize不能为空"))
		c.Abort()
		return
	}
	pageIndex := c.Request.FormValue("pageIndex")
	if pageIndex == "" {
		c.JSON(http.StatusOK, result.Failstr("pageIndex不能为空"))
		c.Abort()
		return
	}

	param := make(map[string]string)
	param["pageSize"] = pageSize
	param["pageIndex"] = pageIndex

	order := models.NewOrder()
	list, err := order.GetOrders(param)
	if err != nil {
		c.JSON(http.StatusOK, result.Fail(err))
		c.Abort()
		return
	}

	pageData := utils.NewPageData(param, list)
	c.JSON(http.StatusOK, pageData)
}