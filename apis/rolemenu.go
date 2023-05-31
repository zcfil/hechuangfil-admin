package apis

import (
	"hechuangfil-admin/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary RoleMenu列表数据
// @Description 获取JSON
// @Tags 角色菜单
// @Param RoleId query string false "RoleId"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/rolemenu [get]
// @Security Bearer
func GetRoleMenu(c *gin.Context) {
	var Rm models.RoleMenu
	err := c.ShouldBind(&Rm)
	result, err := Rm.Get()
	var res models.Response
	if err != nil {
		res.Msg = "抱歉未找到相关信息"
		c.JSON(http.StatusOK, res.ReturnError(200))
		return
	}
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

type RoleMenuPost struct {
	RoleId   string
	RoleMenu []models.RoleMenu
}

// @Summary 创建角色菜单
// @Description 获取JSON
// @Tags 角色菜单
// @Accept  application/x-www-form-urlencoded
// @Product application/x-www-form-urlencoded
// @Param role_id formData string true "role_id"
// @Param menu_id formData string true "menu_id"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/rolemenu [post]
func InsertRoleMenu(c *gin.Context) {
	//roleId := c.Request.FormValue("role_id")
	//menuId := c.Request.FormValue("menu_id")
	//menus := strings.Split(menuId, ",")
	//fmt.Println(menus)
	//var t models.RoleMenu
	//_,err := t.DeleteRoleMenu(roleId)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code":    -1,
	//		"message": "添加失败1",
	//	})
	//	return
	//}
	//_, err2 := t.Insert(roleId, menus)
	//if err2 != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code":    -1,
	//		"message": "添加失败2",
	//	})
	//	return
	//}
	var res models.Response
	res.Msg = "添加成功"
	c.JSON(http.StatusOK, res.ReturnOK())
	return

}

// @Summary 删除用户菜单数据
// @Description 删除数据
// @Tags 角色菜单
// @Param id path string true "id"
// @Param menu_id query string false "menu_id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/rolemenu/{id} [delete]
func DeleteRoleMenu(c *gin.Context) {
	fmt.Println("!!!!!!!!!!!!!!")
	var t models.RoleMenu
	id := c.Param("id")
	menuId := c.Request.FormValue("menu_id")
	fmt.Println(menuId)
	_, err := t.Delete(id, menuId)
	if err != nil {
		var res models.Response
		res.Msg = "删除失败"
		c.JSON(http.StatusOK, res.ReturnError(200))
		return
	}
	var res models.Response
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res.ReturnOK())
	return
}
