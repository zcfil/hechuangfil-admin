package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hechuangfil-admin/models"
	"hechuangfil-admin/pkg"
)

func GetUserVipLevel(c *gin.Context) {
	userID := c.Request.FormValue("userid")
	var conf models.UserLevelConfig
	resultData, err := conf.GetVipLevelList(userID)
	pkg.AssertErr(err, "未找到相关信息", -1)
	var rsp models.Response
	rsp.Data = resultData
	c.JSON(http.StatusOK, rsp.ReturnOK())
}

func EditUserVipLevel(c *gin.Context) {
	userID := c.Request.FormValue("userid")
	levelID := c.Request.FormValue("levelid")
	var conf models.UserLevelConfig
	err := conf.EditUserVipLevel(userID, levelID)
	pkg.AssertErr(err, "跟新vip等级失败", -1)
	var rsp models.Response
	c.JSON(http.StatusOK, rsp.ReturnOK())
}