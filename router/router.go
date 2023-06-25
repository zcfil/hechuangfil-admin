package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"hechuangfil-admin/apis"
	_ "hechuangfil-admin/docs"
	"hechuangfil-admin/handler"
	"hechuangfil-admin/handler/sd"
	_ "hechuangfil-admin/pkg/jwtauth"
	"hechuangfil-admin/router/middleware"
	"log"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//r.Use(middleware.LoggerToFile())
	r.Use(middleware.CustomError)
	r.Use(middleware.NoCache)
	r.Use(middleware.Options)
	r.Use(middleware.Secure)
	r.Use(middleware.RequestId())
	r.Use(middleware.DemoEvn())
	r.Use(gin.Recovery())
	r.Static("/static", "./static")
	//导出
	r.Static("/export", "./runtime/export")
	r.GET("/info", Ping)
	r.GET("/heath", Heath)

	// 监控信息
	svcd := r.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
		svcd.GET("/os", sd.OSCheck)
	}

	// the jwt middleware
	authMiddleware, err := middleware.AuthInit()

	if err != nil {
		_ = fmt.Errorf("JWT Error", err.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler) //登录

	// Refresh time can be longer than token timeout
	r.GET("/refresh_token", authMiddleware.RefreshHandler)
	//r.GET("/dashboard", apis.Dashboard)
	r.GET("/routes", Dashboard)

	//无权限限制接口
	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/getCaptcha", apis.GenerateCaptchaHandler) //获取图片验证码

		apiv1.GET("/rolemenu", apis.GetRoleMenu)
		apiv1.POST("/rolemenu", apis.InsertRoleMenu)
		apiv1.DELETE("/rolemenu/:id", apis.DeleteRoleMenu)
		apiv1.GET("/dict/databytype/:dictType", apis.GetDictDataByDictType)

	}

	//权限限制接口
	auth := r.Group("/api/v1")

	auth.Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		auth.Static("/export", "./runtime/export")
		auth.POST("/logout", handler.LogOut)                 //退出系统
		apiv1.GET("/menuTreeselect", apis.GetMenuTreeSelect) //获取菜单树

		auth.GET("/deptList", apis.GetDeptList)   //获取部门列表
		auth.GET("/deptTree", apis.GetDeptTree)   //获取部门树形图
		auth.GET("/dept/:deptId", apis.GetDept)   //根据部门ID 获取部门
		auth.POST("/dept", apis.InsertDept)       //添加部门
		auth.PUT("/dept", apis.UpdateDept)        //修改部门
		auth.DELETE("/dept/:id", apis.DeleteDept) //删除部门

		auth.GET("/dict/datalist", apis.GetDictDataList)         //获取字典列表
		auth.GET("/dict/data/:dictCode", apis.GetDictData)       //根据字典代码获取字典
		auth.POST("/dict/data", apis.InsertDictData)             //添加字典
		auth.PUT("/dict/data/", apis.UpdateDictData)             // 修改字典
		auth.DELETE("/dict/data/:dictCode", apis.DeleteDictData) //删除字典数据

		auth.GET("/dict/typelist", apis.GetDictTypeList)       //字典类型列表数据
		auth.GET("/dict/type/:dictId", apis.GetDictType)       //通过字典id获取字典类型
		auth.POST("/dict/type", apis.InsertDictType)           //添加字典类型
		auth.PUT("/dict/type", apis.UpdateDictType)            //修改字典类型
		auth.DELETE("/dict/type/:dictId", apis.DeleteDictType) //删除字典类型

		auth.GET("/sysUserList", apis.GetSysUserList)               // 系统用户列表
		auth.GET("/sysUser/:userId", apis.GetSysUser)               //根据用户ID 获取用户
		auth.GET("/sysUser/", apis.GetSysUserInit)                  //获取用户角色和职位
		auth.POST("/sysUser", apis.InsertSysUser)                   //添加管理员
		auth.PUT("/sysUser", apis.UpdateSysUser)                    //修改用户数据
		auth.DELETE("/sysUser/:userId", apis.DeleteSysUser)         //删除用户数据
		auth.GET("/passingSysUserList", apis.GetPassingSysUserList) // 获取审核中的列表
		auth.POST("/submitNewUser", apis.SubmitNewUser)             // 提交新的用户
		auth.POST("/allowNewUserPass", apis.AllowNewUserPass)       // 允许新用户通过

		auth.GET("/rolelist", apis.GetRoleList)              //角色列表
		auth.GET("/role/:roleId", apis.GetRole)              //角色
		auth.PUT("/roleAdd", apis.InsertRole)                //添加角色
		auth.PUT("/role", apis.UpdateRole)                   //修改角色
		auth.PUT("/roledatascope", apis.UpdateRoleDataScope) //修改角色数据
		auth.DELETE("/role/:roleId", apis.DeleteRole)        //删除角色
		auth.PUT("/roleDel", apis.DeleteRole)                //删除角色

		//参数设置
		auth.GET("/configList", apis.GetConfigList)    //配置列表
		auth.GET("/config/:id", apis.GetConfig)        //获取配置
		auth.POST("/config", apis.InsertConfig)        //添加配置
		auth.PUT("/config", apis.UpdateConfig)         //修改配置
		auth.DELETE("/config/:id", apis.DeleteConfig)  //删除配置
		auth.POST("/config/export", apis.ExportConfig) //导出配置

		auth.GET("/roleMenuTreeselect/:roleId", apis.GetMenuTreeRoleselect) //获取角色菜单树
		auth.GET("/roleDeptTreeselect/:roleId", apis.GetDeptTreeRoleselect) //获取部门菜单树

		auth.GET("/getinfo", apis.GetInfo)                        //获取管理员信息
		auth.GET("/user/profile", apis.GetSysUserProfile)         //获取当前登录用户
		auth.POST("/user/profileAvatar", apis.InsetSysUserAvatar) //修改用户头像
		auth.PUT("/user/pwd", apis.SysUserUpdatePwd)              //修改管理员密码
		auth.GET("/user/getVerificationCode", apis.CreateSecret)  //创建谷歌验证码
		auth.POST("/user/bindVerificationCode", apis.BindCode)    //绑定谷歌验证

		auth.GET("/postlist", apis.GetPostList)       //职位列表
		auth.GET("/post/:postId", apis.GetPost)       //职位列表数据
		auth.POST("/post", apis.InsertPost)           //添加职位
		auth.PUT("/post", apis.UpdatePost)            //修改职位
		auth.DELETE("/post/:postId", apis.DeletePost) //删除职位

		auth.GET("/menulist", apis.GetMenuList)   //菜单列表
		auth.GET("/menu/:id", apis.GetMenu)       //菜单数据
		auth.POST("/menu", apis.InsertMenu)       //添加菜单
		auth.PUT("/menu", apis.UpdateMenu)        //修改菜单
		auth.DELETE("/menu/:id", apis.DeleteMenu) //删除菜单
		auth.GET("/menurole", apis.GetMenuRole)   //获取角色权限
		auth.GET("/menuids", apis.GetMenuIDS)     //获取角色对应的菜单id数组

		auth.GET("/loginloglist", apis.GetLoginLogList)       //登录日志
		auth.GET("/loginlog/:infoId", apis.GetLoginLog)       //通过编码获取登录日志
		auth.POST("/loginlog", apis.InsertLoginLog)           //添加日志
		auth.PUT("/loginlog", apis.UpdateLoginLog)            //修改日志
		auth.DELETE("/loginlog/:infoId", apis.DeleteLoginLog) //删除日志

		auth.GET("/operloglist", apis.GetOperLogList)       //操作日志
		auth.GET("/operlog/:operId", apis.GetOperLog)       //获取日志
		auth.DELETE("/operlog/:operId", apis.DeleteOperLog) //删除操作日志

		auth.GET("/configKey/:configKey", apis.GetConfigByConfigKey) //通过配置Key获取配置信息

		//会员管理
		//auth.GET("/memberList", apis.GetUsersList)                   //会员管理列表
		//auth.POST("/member/memberDelete", apis.UsersDelete)          //删除会员
		//auth.POST("/member/memberEdit", apis.UsersEdit)              //编辑会员
		//
		//
		//auth.GET("/memberExport", apis.GetUsersListExport)              //会员列表 导出

		//矿多多
		auth.POST("/updateConfig", apis.UpdateConfig1)
		auth.GET("/getConfigKey", apis.GetConfigKey)

		auth.GET("/individualPerformance", apis.IndividualPerformance) //个人业绩

		auth.GET("/profitconfigList", apis.ProfitconfigList) //分润配置列表
		//auth.POST("/profitconfigOnce", apis.ProfitconfigOnce)             //添加分润配置（一次性分配）
		//auth.POST("/updateProfitConfigOnce", apis.UpdateProfitconfigOnce) //修改分润配置（一次性分配）
		auth.POST("/profitconfigEdit", apis.ProfitconfigEdit) //业务员分润配置(提交修改)
		//auth.POST("/delProfitconfigOnce", apis.DelProfitconfigOnce)       //添加分润配置（一次性分配）

		auth.GET("/userList", apis.GetUserList)

		auth.GET("/getUserVipLevel", apis.GetUserVipLevel)    // 获取用户的vip等级
		auth.POST("/editUserVipLevel", apis.EditUserVipLevel) // 编辑玩家的vip等级

		auth.GET("/getUserPerformance", apis.GetUserPerformance) // 获取主页显示内容

		//合创需求
		auth.GET("/customerList", apis.GetcustomerList)    //客户列表
		auth.GET("/getCustomerByid", apis.GetCustomerByid) //获取客户信息

		auth.GET("/auditList", apis.AuditList)       //审核列表
		auth.GET("/auditLogList", apis.AuditLogList) //审核记录
		auth.GET("/auditById", apis.AuditByID)       //审核记录
		auth.POST("/audit", apis.CheckAudit)         //审核
		auth.POST("/audits", apis.CheckAudits)       //批量审核

		auth.GET("/rechargeList", apis.GetRechargeList)     //充值列表
		auth.GET("/settlementList", apis.GetSettlementList) //结算列表

		auth.GET("/statementSummary", apis.StatementSummary)         //汇总报表
		auth.GET("/statementOrderProfit", apis.StatementOrderProfit) //订单一次性分配报表
		auth.GET("/getOrders", apis.GetOrders)                       // 获取订单列表

		//财务配置
		auth.GET("/financeConfigList", apis.FinanceConfigList) //配置列表
		auth.POST("/updateConfigById", apis.UpdateConfigById)  //修改配置

		// 新闻资讯
		auth.GET("/getAllNews", apis.GetAllNews)             // 获取所有新闻列表
		auth.POST("/addNews", apis.AddNews)                  // 添加新闻
		auth.POST("/updateNews", apis.UpdateNews)            // 修改新闻
		auth.POST("/deleteNews", apis.DeleteNews)            // 删除新闻
		auth.POST("/updateNewsStatus", apis.UpDateNewStatus) // 更新新闻状态

	}
	lanGroup := r.Group("/LAN")
	//lanGroup.Use(middleware.CheckLANIP())
	lanGroup.Use()
	{
		lanGroup.GET("/walletNew", apis.WalletNew)
	}
	//r.POST("/audit", apis.CheckAudit)              //审核
	r.GET("/rechargeList", apis.GetRechargeList) //充值列表

	r.GET("/settlementList", apis.GetSettlementList) //结算列表
	//批量导入地址
	r.POST("/api/v1/uploadAddress", apis.UploadAddress) //批量导入账号地址

	return r
}

func Ping(c *gin.Context) {
	log.Println(viper.GetString(`settings.database.type`))
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func Heath(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "UP",
	})
}

func Dashboard(c *gin.Context) {

	var user = make(map[string]interface{})
	user["login_name"] = "admin"
	user["user_id"] = 1
	user["user_name"] = "管理员"
	user["dept_id"] = 1

	var cmenuList = make(map[string]interface{})
	cmenuList["children"] = nil
	cmenuList["parent_id"] = 1
	cmenuList["title"] = "用户管理"
	cmenuList["name"] = "Sysuser"
	cmenuList["icon"] = "user"
	cmenuList["order_num"] = 1
	cmenuList["id"] = 4
	cmenuList["path"] = "sysuser"
	cmenuList["component"] = "sysuser/index"

	var lista = make([]interface{}, 1)
	lista[0] = cmenuList

	var menuList = make(map[string]interface{})
	menuList["children"] = lista
	menuList["parent_id"] = 1
	menuList["name"] = "Upms"
	menuList["title"] = "权限管理"
	menuList["icon"] = "example"
	menuList["order_num"] = 1
	menuList["id"] = 4
	menuList["path"] = "/upms"
	menuList["component"] = "Layout"

	var list = make([]interface{}, 1)
	list[0] = menuList
	var data = make(map[string]interface{})
	data["user"] = user
	data["menuList"] = list

	var r = make(map[string]interface{})
	r["code"] = 200
	r["data"] = data

	c.JSON(200, r)
}
