package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hechuangfil-admin/models"
	"hechuangfil-admin/pkg"
	"hechuangfil-admin/utils"
)

func GetcustomerList(c *gin.Context) {
	var u models.Customer
	var err error
	param := make(map[string]string)

	param["keyword"] = c.Request.FormValue("keyword")
	param["pageSize"] = c.DefaultQuery("pageSize","10")
	param["pageIndex"] = c.DefaultQuery("pageIndex","1")

	result, err := u.CustomerList(param)
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageData(param,result)

	c.JSON(http.StatusOK, res)
}

func GetCustomerByid(c *gin.Context) {
	var data models.Customer
	//err := c.BindWith(&data, binding.JSON)
	//pkg.AssertErr(err, "", 500)
	id := c.Request.FormValue("customerid")
	var res models.Response
	re,err := data.NewCustomer(id)
	if err!=nil{
		res.Code = 0
		res.Msg = err.Error()
	}
	res.Data = re
	c.JSON(http.StatusOK, res.ReturnOK())
}