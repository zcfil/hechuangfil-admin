package apis

import (
	"hechuangfil-admin/models"
	"hechuangfil-admin/pkg"
	"hechuangfil-admin/pkg/export"
	"hechuangfil-admin/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

// @Summary 配置列表数据
// @Description 获取JSON
// @Tags 配置
// @Param configKey query string false "configKey"
// @Param configName query string false "configName"
// @Param configType query string false "configType"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} models.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/configList [get]
// @Security
func GetConfigList(c *gin.Context) {
	var data models.Config
	var err error
	var pageSize = 10
	var pageIndex = 1

	if size := c.Request.FormValue("pageSize"); size != "" {
		pageSize = pkg.StrToInt(err, size)
	}

	if index := c.Request.FormValue("pageIndex"); index != "" {
		pageIndex = pkg.StrToInt(err, index)
	}

	data.Name = c.Request.FormValue("configName")
	data.Title = c.Request.FormValue("configTitle")
	data.Group = c.Request.FormValue("configGroup")
	//data.DataScope = utils.GetUserIdStr(c)
	result, count, err := data.GetPage(pageSize, pageIndex)
	pkg.AssertErr(err, "", -1)

	var mp = make(map[string]interface{}, 3)
	mp["list"] = result
	mp["count"] = count
	mp["pageIndex"] = pageIndex
	mp["pageIndex"] = pageSize

	var res models.Response
	res.Data = mp

	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 获取配置
// @Description 获取JSON
// @Tags 配置
// @Param configId path int true "配置编码"
// @Success 200 {object} models.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/config/{id} [get]
// @Security
func GetConfig(c *gin.Context) {
	var Config models.Config
	Config.ID, _ = utils.StringToInt64(c.Param("id"))
	result, err := Config.Get()
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	var res models.Response
	res.Data = result

	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 获取配置
// @Description 获取JSON
// @Tags 配置
// @Param configKey path int true "configKey"
// @Success 200 {object} models.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/configKey/{configKey} [get]
// @Security
func GetConfigByConfigKey(c *gin.Context) {
	var Config models.Config
	Config.Name = c.Param("name")
	result, err := Config.Get()
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	var res models.Response
	res.Data = result
	res.Msg = result.Value
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 添加配置
// @Description 获取JSON
// @Tags 配置
// @Accept  application/json
// @Product application/json
// @Param data body models.Config true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dict/data [post]
// @Security Bearer
func InsertConfig(c *gin.Context) {
	var data models.Config
	err := c.BindWith(&data, binding.JSON)
	data.CreateBy = utils.GetUserIdStr(c)
	pkg.AssertErr(err, "", 500)
	result, err := data.Create()
	pkg.AssertErr(err, "", -1)

	var res models.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 修改配置
// @Description 获取JSON
// @Tags 配置
// @Accept  application/json
// @Product application/json
// @Param data body models.Config true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/config [put]
// @Security Bearer
func UpdateConfig(c *gin.Context) {
	var data models.Config
	err := c.BindWith(&data, binding.JSON)
	pkg.AssertErr(err, "数据解析失败", -1)
	data.UpdateBy = utils.GetUserIdStr(c)
	result, err := data.Update(data.ID)
	pkg.AssertErr(err, "", -1)

	var res models.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}
func UpdateConfig1(c *gin.Context) {
	var data models.Config
	id := c.Request.FormValue("id")
	value := c.Request.FormValue("value")
	err := data.UpdateConfig(value,id)
	pkg.AssertErr(err, "", -1)

	var res models.Response
	c.JSON(http.StatusOK, res.ReturnOK())
}
func GetConfigKey(c *gin.Context) {
	var data models.Config
	configKey := c.Request.FormValue("configKey")
	var res models.Response
	var err error
	res.Data,err = data.GetConfig(configKey)
	if err!=nil{
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}

	c.JSON(http.StatusOK, res.ReturnOK())

}

// @Summary 删除配置
// @Description 删除数据
// @Tags 配置
// @Param configId path int true "configId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/config/{configId} [delete]
func DeleteConfig(c *gin.Context) {
	var data models.Config
	id, err := utils.StringToInt64(c.Param("id"))
	data.UpdateBy = utils.GetUserIdStr(c)
	_, err = data.Delete(id)
	pkg.AssertErr(err, "修改失败", 500)

	var res models.Response
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 导出配置
// @Description 导出系统配置信息
// @Tags 配置
// @Param configId path int true "configId"
// @Success 200 {string} string	"{"code": 200, "message": "导出成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/config/export [post]
func ExportConfig(c *gin.Context) {
	var data models.Config

	filename, err := data.Export()
	pkg.AssertErr(err, "导出失败", 500)

	var res models.Response
	res.Data = export.GetExcelFullUrl(filename)
	res.Msg = "导出成功"
	c.JSON(http.StatusOK, res.ReturnOK())
}
