package apis

/*// @Param sort query string false "id||其他字段"
// @Param order query string false "desc"
// @Param offset query string false "第几条数据"
// @Param limit query string false "每页显示数据条数"*/

// @Summary 会员管理列表数据
// @Description 获取JSON
// @Tags 会员管理
// @Param cid query string false "cid"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} models.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/memberList [get]
// @Security Bearer
//func GetUsersList(c *gin.Context) {
//	var u models.User
//	var err error
//	param := make(map[string]interface{})
//	var pageSize = 10
//	var pageIndex = 1
//
//	if size := c.Request.FormValue("pageSize"); size != "" {
//		pageSize = pkg.StrToInt(err, size)
//	}
//
//	if index := c.Request.FormValue("pageIndex"); index != "" {
//		pageIndex = pkg.StrToInt(err, index)
//	}
//
//	param["keyword"] = c.Request.FormValue("keyword")
//	param["pageSize"] = pageSize
//	param["pageIndex"] = pageIndex
//	param["status"] = c.Request.FormValue("status")
//	param["level"] = c.Request.FormValue("level")
//	result, count, err := u.UserList(param)
//	//u.GetPage(pageSize, pageIndex,param)
//	pkg.AssertErr(err, "抱歉未找到相关信息", -1)
//	var mp = make(map[string]interface{}, 3)
//	mp["list"] = result
//	mp["count"] = count
//	mp["pageIndex"] = pageIndex
//	mp["pageSize"] = pageSize
//	var res models.Response
//	res.Data = mp
//	res.Msg = "查询成功！"
//	c.JSON(http.StatusOK, res.ReturnOK())
//}
//
////会员列表导出
//func GetUsersListExport(c *gin.Context) {
//	var u models.User
//	param := make(map[string]interface{})
//	param["status"] = c.Request.FormValue("status")
//	param["keyword"] = c.Request.FormValue("keyword")
//	result, _ := u.ExportExcelUrl(param)
//	//c.Header("Content-Type", "application/octet-stream")
//	//c.Header("Content-Disposition", "attachment; filename="+"Workbook.xlsx")
//	//c.Header("Content-Transfer-Encoding", "binary")
//	//result.Write(c.Writer)
//
//	var res models.Response
//	res.Msg = result
//	c.JSON(http.StatusOK, res.ReturnOK())
//}
//
//// @Summary 删除会员
//// @Description 删除数据
//// @Tags 会员管理
//// @Param ids path int true "1,2,3"
//// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
//// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
//// @Router /user/memberDelete [post]
////删除会员
//func UsersDelete(c *gin.Context) {
//	var u models.User
//	param := make(map[string]interface{})
//
//	param["ids"] = c.Request.FormValue("ids")
//	_, err := u.UserDelete(param)
//	pkg.AssertErr(err, "删除失败", -1)
//
//	var res models.Response
//	res.Msg = "删除成功"
//	c.JSON(http.StatusOK, res.ReturnOK())
//}
//
////编辑会员
//func UsersEdit(c *gin.Context) {
//	var u models.User
//	param := make(map[string]interface{})
//	param["id"] = c.Request.FormValue("id")
//	param["email"] = c.Request.FormValue("email")
//	param["address"] = c.Request.FormValue("address")
//	param["money_address"] = c.Request.FormValue("money_address")
//	param["status"] = c.Request.FormValue("status")
//	param["level"] = c.Request.FormValue("level")
//	pwd := c.Request.FormValue("pwd")
//	param["pwd"] = utils.GenLowMD5([]byte(pwd))
//	_, err := u.UserEdit(param)
//	pkg.AssertErr(err, "修改失败", -1)
//
//	var res models.Response
//	res.Msg = "修改成功"
//	c.JSON(http.StatusOK, res.ReturnOK())
//}

