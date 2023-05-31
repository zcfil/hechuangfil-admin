package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strconv"
	"time"
	"hechuangfil-admin/models"
	"hechuangfil-admin/pkg"
	"hechuangfil-admin/pkg/googleAuthenticator"
	"hechuangfil-admin/utils"
)

// @Summary 列表数据
// @Description 获取JSON
// @Tags 用户
// @Param username query string false "username"
// @Success 200 {string} string "{"code": 200, "data": [...]}"
// @Success 200 {string} string "{"code": -1, "msg": "抱歉未找到相关信息"}"
// @Router /api/v1/sysUserList [get]
// @Security Bearer
func GetSysUserList(c *gin.Context) {
	var data models.SysUser
	var err error
	var pageSize = 10
	var pageIndex = 1

	size := c.Request.FormValue("pageSize")
	if size != "" {
		pageSize = pkg.StrToInt(err, size)
	}

	index := c.Request.FormValue("pageIndex")
	if index != "" {
		pageIndex = pkg.StrToInt(err, index)
	}

	data.Username = c.Request.FormValue("username")
	data.Phone = c.Request.FormValue("phone")

	postId := c.Request.FormValue("postId")
	data.PostId, _ = strconv.ParseInt(postId, 10, 64)

	deptId := c.Request.FormValue("deptId")
	data.DeptId, _ = strconv.ParseInt(deptId, 10, 64)

	data.DataScope = utils.GetUserIdStr(c)
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

// @Summary 获取用户
// @Description 获取JSON
// @Tags 用户
// @Param userId path int true "用户编码"
// @Success 200 {object} models.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sysUser/{userId} [get]
// @Security
func GetSysUser(c *gin.Context) {
	var SysUser models.SysUser
	SysUser.Id, _ = strconv.ParseInt(c.Param("userId"), 10, 64)
	result, err := SysUser.Get()
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)
	var SysRole models.SysRole
	var Post models.Post
	SysRole.RoleKey = utils.GetRolekey(c)
	roles, err := SysRole.GetList()
	posts, err := Post.GetList()

	postIds := make([]int64, 0)
	postIds = append(postIds, result.PostId)

	roleIds := make([]int64, 0)
	roleIds = append(roleIds, result.RoleId)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    result,
		"postIds": postIds,
		"roleIds": roleIds,
		"roles":   roles,
		"posts":   posts,
	})
}

// @Summary 获取当前登录用户
// @Description 获取JSON
// @Tags 个人中心
// @Success 200 {object} models.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/user/profile [get]
// @Security
func GetSysUserProfile(c *gin.Context) {
	var SysUser models.SysUser
	userId := utils.GetUserIdStr(c)
	SysUser.Id, _ = strconv.ParseInt(userId, 10, 64)
	result, err := SysUser.Get()
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)
	var SysRole models.SysRole
	var Post models.Post
	var Dept models.Dept
	SysRole.RoleKey = utils.GetRolekey(c)
	//获取角色列表
	roles, err := SysRole.GetList()
	//获取职位列表
	posts, err := Post.GetList()
	//获取部门列表
	Dept.Deptid = result.DeptId
	dept, err := Dept.Get()

	postIds := make([]int64, 0)
	postIds = append(postIds, result.PostId)

	roleIds := make([]int64, 0)
	roleIds = append(roleIds, result.RoleId)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    result,
		"postIds": postIds,
		"roleIds": roleIds,
		"roles":   roles,
		"posts":   posts,
		"dept":    dept,
	})
}

// @Summary 获取用户角色和职位
// @Description 获取JSON
// @Tags 用户
// @Success 200 {object} models.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/sysUser [get]
// @Security
func GetSysUserInit(c *gin.Context) {
	var SysRole models.SysRole
	var Post models.Post
	SysRole.RoleKey = utils.GetRolekey(c)
	roles, err := SysRole.GetList()
	posts, err := Post.GetList()
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)
	mp := make(map[string]interface{}, 2)
	mp["roles"] = roles
	mp["posts"] = posts
	var res models.Response
	res.Data = mp
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 创建用户
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body models.SysUser true "用户数据"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/sysUser [post]
func InsertSysUser(c *gin.Context) {
	var sysuser models.SysUser
	err := c.ShouldBind(&sysuser)
	fmt.Println(sysuser)
	pkg.AssertErr(err, "非法数据格式", 500)

	sysuser.CreateTime = time.Now()
	sysuser.UpdateTime = time.Now()
	sysuser.CreateBy = utils.GetUserIdStr(c)
	id, err := sysuser.Insert()
	var res models.Response
	if err != nil {
		res.Msg = err.Error()
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	res.Data = id
	res.Msg = "添加成功"
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 修改用户数据
// @Description 获取JSON
// @Tags 用户
// @Accept  application/json
// @Product application/json
// @Param data body models.SysUser true "body"
// @Success 200 {string} string	"{"code": 200, "message": "修改成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "修改失败"}"
// @Router /api/v1/sysuser/{userId} [put]
func UpdateSysUser(c *gin.Context) {
	var data models.SysUser
	err := c.ShouldBind(&data)
	pkg.AssertErr(err, "数据解析失败", -1)
	data.UpdateBy = utils.GetUserIdStr(c)
	result, err := data.Update(data.Id)
	var res models.Response
	if err != nil || result.Id == 0 {
		res.Msg = "修改失败"
		c.JSON(http.StatusOK, res.ReturnError(400))
		return
	}
	res.Msg = "修改成功"
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 删除用户数据
// @Description 删除数据
// @Tags 用户
// @Param userId path int true "userId"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/sysuser/{userId} [delete]
func DeleteSysUser(c *gin.Context) {
	var data models.SysUser
	data.UpdateBy = utils.GetUserIdStr(c)
	//IDS := utils.IdsStrToIdsInt64Group("userId", c)
	IDS := c.Param("userId")
	result, err := data.BatchDelete(IDS)
	if err != nil || !result {
		var res models.Response
		res.Msg = "删除失败"
		c.JSON(http.StatusOK, res.ReturnError(501))
		return
	}
	var res models.Response
	res.Msg = "删除成功"
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 修改头像
// @Description 获取JSON
// @Tags 用户
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/user/profileAvatar [post]
func InsetSysUserAvatar(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]
	guid := uuid.New().String()
	filPath := "static/uploadfile/" + guid + ".jpg"
	for _, file := range files {
		log.Println(file.Filename)
		// 上传文件至指定目录
		_ = c.SaveUploadedFile(file, filPath)
	}
	sysuser := models.SysUser{}
	sysuser.Id = utils.GetUserId(c)
	sysuser.Avatar = "/" + filPath
	sysuser.UpdateBy = utils.GetUserIdStr(c)
	sysuser.Update(sysuser.Id)

	c.JSON(http.StatusCreated, gin.H{
		"code": 200,
		"data": filPath,
	})

}

// @Summary 获取谷歌验证码密钥
// @Description 获取JSON
// @Tags 用户
// @Accept multipart/form-data
// @Success 200 {string} string	"{"code": 200, "message": "获取成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /user/getSecret [get]
func CreateSecret(c *gin.Context) {

	//判断用户是否已经绑定谷歌验证吗
	idstr := utils.GetUserIdStr(c)
	id, err := strconv.Atoi(idstr)
	pkg.AssertErr(err, "数据解析失败", -1)
	var sysu models.SysUser
	sysu.Id = int64(id)
	//sv, err := sysu.Get()
	//
	//var res models.Response
	//var goo models.GoogleAuth

	//if len(sv.Verification) <= 0 {
	//	ga := googleAuthenticator.NewGAuth()
	//	secret, err := ga.CreateSecret(16)
	//	pkg.AssertErr(err, "数据异常！", -1)
	//	//"googleSecret":"Z6QZ2QZ5CZE5I5WE",
	//	//"googleAccount":"云构:18902538879",
	//	//"url":"otpauth://totp/云构:18902538879?secret=Z6QZ2QZ5CZE5I5WE"
	//	goo.GoogleSecret = secret
	//	goo.GoogleAccount = "BET:" + utils.GetUserName(c)
	//	goo.URL = "otpauth://totp/" + goo.GoogleAccount + "?secret=" + secret
	//	goo.IsBind = false
	//	res.Data = goo
	//	res.Msg = "获取成功！"
	//} else {
	//	goo.IsBind = true
	//	res.Data = goo
	//	res.Msg = "已经绑定！"
	//}
	//c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 绑定谷歌验证码
// @Description 获取JSON
// @Tags 用户
// @Accept multipart/form-data
// @Success 200 {string} string	"{"code": 200, "message": "获取成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /user/bindVerificationCode [post]
func BindCode(c *gin.Context) {
	var data models.GoogleAuth
	err := c.BindWith(&data, binding.JSON)
	log.Println("BindCode DATA Secret:", data.GoogleSecret, " code :", data.Code)
	pkg.AssertErr(err, "数据解析失败", -1)

	//ga := googleAuthenticator.NewGAuth()
	//isOK, err := ga.VerifyCode(data.GoogleSecret, data.Code, 1)
	//log.Println("BindCode VerifyCode:", isOK)
	var res models.Response
	//if isOK {
	//	idstr := utils.GetUserIdStr(c)
	//	id, err := strconv.Atoi(idstr)
	//	pkg.AssertErr(err, "数据解析失败", -1)
	//	var sysu models.SysUser
	//	sysu.Id = int64(id)
	//	sysu.Verification = data.GoogleSecret
	//	result, err := sysu.Update(sysu.Id)
	//
	//	if err != nil || result.Id == 0 {
	//		res.Msg = "绑定失败！"
	//		c.JSON(http.StatusOK, res.ReturnError(200))
	//		return
	//	}
	//	res.Msg = "绑定成功！"
	//
	//} else {
	//	res.Msg = "验证码无效！"
	//}
	c.JSON(http.StatusOK, res.ReturnOK())
}

// @Summary 验证谷歌验证码
// @Description 获取JSON
// @Tags 用户
// @Accept multipart/form-data
// @Success 200 {string} string	"{"code": 200, "message": "获取成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /user/VerifyCode [post]
func VerifyCode(c *gin.Context) {

	ga := googleAuthenticator.NewGAuth()
	secret, err := ga.CreateSecret(16)
	pkg.AssertErr(err, "数据异常！", -1)

	var res models.Response
	res.Data = secret
	res.Msg = "获取成功！"
	c.JSON(http.StatusOK, res.ReturnOK())
}

func SysUserUpdatePwd(c *gin.Context) {
	var pwd models.SysAdminPwd
	err := c.Bind(&pwd)
	pkg.AssertErr(err, "数据解析失败", 500)
	sysuser := models.SysUser{}
	sysuser.Id = utils.GetUserId(c)
	sysuser.SetPwd(pwd)
	c.JSON(http.StatusCreated, gin.H{
		"code": 200,
		"data": "密码修改成功",
	})

}


func GetUserList(c *gin.Context) {
	var data models.SysUser
	//var res models.Response

	param := make(map[string]string)
	param["keyword"] = c.Request.FormValue("keyword")
	param["pageSize"] = c.DefaultQuery("pageSize","10")
	param["pageIndex"] = c.DefaultQuery("pageIndex","1")

	result, err := data.GetSysList(param)
	//if err !=nil{
	//	res.Msg = err.Error()
	//	c.JSON(http.StatusOK, res.ReturnError(0))
	//}
	//res.Data = result
	pkg.AssertErr(err, "抱歉未找到相关信息", -1)

	res := utils.NewPageData(param,result)

	c.JSON(http.StatusOK, res )
}

// GetUserPerformance 获取用户的业绩
func GetUserPerformance(c *gin.Context) {
	userId := utils.GetUserId(c)
	per := models.NewUserPerformance(userId)

	err := per.GetTotal()
	pkg.AssertErr(err, "获取总业绩失败", -1)
	err = per.GetToday()
	pkg.AssertErr(err, "获取当天业绩失败", -1)
	err = per.GetVipLevel()
	pkg.AssertErr(err, "获取vip等级失败", -1)

	var res models.Response
	res.Data = per
	c.JSON(http.StatusOK, res.ReturnOK())
}

// GetPassingSysUserList 获取审核中的用户列表
func GetPassingSysUserList(c *gin.Context) {
	pageSizeStr := c.Request.FormValue("pageSize")
	pageIndexStr := c.Request.FormValue("pageIndex")
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 64)
	pkg.AssertErr(err, "pageSize错误", -1)
	pageIndex, err := strconv.ParseInt(pageIndexStr, 10, 64)
	pkg.AssertErr(err, "pageIndex错误", -1)

	var user models.InReviewUser

	roleKey := utils.GetRolekey(c)
	curRoleID := utils.GetUserIdStr(c)
	list, total, err := user.GetUserPassingList(pageIndex, pageSize, roleKey, curRoleID)
	pkg.AssertErr(err, "获取通过列表错误", -1)
	var mp = make(map[string]interface{}, 3)
	mp["pageIndex"] = pageIndexStr
	mp["pageIndex"] = pageSizeStr
	mp["list"] = list
	mp["totalCount"] = total

	var res models.Response
	res.Data = mp
	c.JSON(http.StatusOK, res.ReturnOK())
}

// SubmitNewUser 提交新的用户
func SubmitNewUser(c *gin.Context) {
	var sysUser models.SysUser
	err := c.ShouldBind(&sysUser)
	pkg.AssertErr(err, "非法数据格式", 500)
	submitterID := utils.GetUserIdStr(c)
	err = sysUser.SubmitNewUser(submitterID)
	pkg.AssertErr(err, "提交新玩家数据失败", 501)

	var res models.Response
	c.JSON(http.StatusOK, res.ReturnOK())
}

// AllowNewUserPass 允许新用户通过 is_allow 1代表同意
func AllowNewUserPass(c *gin.Context) {
	userID := c.Request.FormValue("user_id")
	isAllow := c.Request.FormValue("is_allow")

	if isAllow == "1" {
		allowUser(c, userID)
		return
	}
	notAllowUser(c, userID)
}

func allowUser(c *gin.Context, userID string) {
	var sysUser models.SysUser
	err := sysUser.AllowNewPass(userID)
	pkg.AssertErr(err, "允许添加下级失败", 500)
	var res models.Response
	c.JSON(http.StatusOK, res.ReturnOK())
}

func notAllowUser(c *gin.Context, userID string) {
	var sysUser models.SysUser
	err := sysUser.NotAllowUserPass(userID)
	pkg.AssertErr(err, "拒绝添加下级失败", 500)
	var res models.Response
	c.JSON(http.StatusOK, res.ReturnOK())
}