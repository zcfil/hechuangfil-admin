package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"hechuangfil-admin/models"
)

func ProfitconfigList(c *gin.Context) {
	var u models.UserLevel
	var err error
	var res models.Response
	result, err := u.GetSetUserLevel(0)
	//pkg.AssertErr(err, "抱歉未找到相关信息", -1)
	if err!=nil{
		res.Code = 0
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res)
	}
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

func ProfitconfigOnce(c *gin.Context) {
	var data models.Profit
	param := make(map[string]string)
	param["percent"] = c.Request.FormValue("percent")
	param["userid"] = c.Request.FormValue("userid")
	//param["userid"] = utils.GetUserIdStr(c)
	//pkg.AssertErr(err, "", 400)
	var res models.Response
	err := data.ProfitconfigOnce(param)
	if err!=nil{
		res.Code = 0
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}
func DelProfitconfigOnce(c *gin.Context) {
	var data models.Profit
	param := make(map[string]string)
	param["id"] = c.Request.FormValue("id")
	//param["userid"] = utils.GetUserIdStr(c)
	//pkg.AssertErr(err, "", 400)
	var res models.Response
	err := data.DelProfitconfigOnce(param)
	if err!=nil{
		res.Code = 0
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res)
		return
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}


func UpdateProfitconfigOnce(c *gin.Context) {
	var data models.Profit
	param := make(map[string]string)
	param["percent"] = c.Request.FormValue("percent")
	param["userid"] = c.Request.FormValue("userid")
	param["id"] = c.Request.FormValue("id")
	//param["userid"] = utils.GetUserIdStr(c)
	//pkg.AssertErr(err, "", 400)
	var res models.Response
	err := data.UpdateProfitconfigOnce(param)
	if err!=nil{
		res.Code = 0
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res)
	}

	c.JSON(http.StatusOK, res.ReturnOK())
}

func ProfitconfigEdit(c *gin.Context) {
	var data models.ProfitEdit

	//err := c.BindWith(&data, binding.JSON)
	//pkg.AssertErr(err, "", 500)
	//err := c.MustBindWith(&data, binding.JSON)
	err := c.ShouldBind(&data)
	fmt.Println(data)
	var res models.Response
	//err := data.ProfitconfigEdit(param)
	err = data.ProfitEdit()
	if err!=nil{
		res.Code = 0
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res)
	}
	res.Msg = "操作成功！"
	c.JSON(http.StatusOK, res.ReturnOK())
}
