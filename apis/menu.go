package apis

import (
	"hechuangfil-admin/models"
	"hechuangfil-admin/pkg"
	"hechuangfil-admin/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
)

// @Summary Menu列表数据
// @Description 获取JSON
// @Tags 菜单
// @Param menuName query string false "menuName"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/menulist [get]
// @Security Bearer
func GetMenuList(c *gin.Context) {
	var Menu models.Menu
	Menu.MenuName = c.Request.FormValue("menuName")
	Menu.Visible = c.Request.FormValue("visible")
	//Menu.DataScope = utils.GetUserIdStr(c)
	result, err := Menu.SetMenu()
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	var res models.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary Menu列表数据
// @Description 获取JSON
// @Tags 菜单
// @Param menuName query string false "menuName"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/menu [get]
// @Security Bearer
func GetMenu(c *gin.Context) {
	var data models.Menu
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	data.MenuId = id
	result, err := data.GetByMenuId()
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)
	var res models.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

func GetMenuTreeRoleselect(c *gin.Context) {
	var Menu models.Menu
	var SysRole models.SysRole
	id, err := strconv.ParseInt(c.Param("roleId"), 10, 64)
	SysRole.Id = id
	result, err := Menu.SetMenuLable()
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)
	menuIds := make([]int64, 0)
	if id != 0 {
		menuIds, err = SysRole.GetRoleMeunId()
		pkg.AssertErr(err, "抱歉未找到相关信息", -1)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":        200,
		"menus":       result,
		"checkedKeys": menuIds,
	})
}

// @Summary 获取菜单树
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/x-www-form-urlencoded
// @Product application/x-www-form-urlencoded
// @Success 200 {string} string	"{"code": 200, "msg": "获取成功"}"
// @Success 200 {string} string	"{"code": -1, "msg": "获取失败"}"
// @Router /api/v1/menuTreeSelect [get]
// @Security Bearer
func GetMenuTreeSelect(c *gin.Context) {
	var data models.Menu
	result, err := data.SetMenuLable()
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)
	var res models.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 创建菜单
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/x-www-form-urlencoded
// @Product application/x-www-form-urlencoded
// @Param menuName formData string true "menuName"
// @Param Path formData string false "Path"
// @Param Action formData string true "Action"
// @Param Permission formData string true "Permission"
// @Param ParentId formData string true "ParentId"
// @Param IsDel formData string true "IsDel"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/menu [post]
// @Security Bearer
func InsertMenu(c *gin.Context) {
	var data models.Menu
	err := c.BindWith(&data, binding.JSON)
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)
	data.CreateBy = utils.GetUserIdStr(c)
	result, err := data.Create()
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	var res models.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 修改菜单
// @Description 获取JSON
// @Tags 菜单
// @Accept  application/x-www-form-urlencoded
// @Product application/x-www-form-urlencoded
// @Param id path int true "id"
// @Param data body models.Menu true "body"
// @Success 200 {string} string	"{"code": 200, "message": "修改成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "修改失败"}"
// @Router /api/v1/menu/{id} [put]
// @Security Bearer
func UpdateMenu(c *gin.Context) {
	var data models.Menu
	err2 := c.BindWith(&data, binding.JSON)
	data.UpdateBy = utils.GetUserIdStr(c)
	pkg.AssertErr(err2, "修改失败", -1)
	_, err := data.Update(data.MenuId)
	pkg.AssertErr(err, "", 501)

	var res models.Response
	res.Msg = "修改成功"
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 删除菜单
// @Description 删除数据
// @Tags 菜单
// @Param id path int true "id"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/menu/{id} [delete]
func DeleteMenu(c *gin.Context) {
	var data models.Menu
	id, err := utils.StringToInt64(c.Param("id"))
	data.UpdateBy = utils.GetUserIdStr(c)
	_, err = data.Delete(id)
	pkg.AssertErr(err, "删除失败", 500)

	var res models.Response
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 根据角色名称获取菜单列表数据（左菜单使用）
// @Description 获取JSON
// @Tags 菜单
// @Param id path int true "id"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/menurole [get]
// @Security Bearer
func GetMenuRole(c *gin.Context) {
	var Menu models.Menu
	result, err := Menu.SetMenuRole(utils.GetRoleName(c))
	pkg.AssertErr(err, "获取失败", 500)
	var res models.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 获取角色对应的菜单id数组
// @Description 获取JSON
// @Tags 菜单
// @Param id path int true "id"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "message": "抱歉未找到相关信息"}"
// @Router /api/v1/menuids/{id} [get]
// @Security Bearer
func GetMenuIDS(c *gin.Context) {
	var data models.RoleMenu
	data.RoleName = c.GetString("role")
	data.UpdateBy = utils.GetUserIdStr(c)
	result, err := data.GetIDS()
	pkg.AssertErr(err, "获取失败", 500)
	var res models.Response
	res.Data = result
	c.JSON(http.StatusOK, res.ReturnOK())
}
