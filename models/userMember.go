package models

//会员表
//type User struct {
//	ID           int64     `orm:"pk;auto;column(id)"json:"id"`
//	NickName     string    `orm:"size(32);column(nick_name)"json:"nick_name"`               //昵称
//	Gender       int8      `orm:"default(0);column(gender)"json:"gender"`                   //性别(0保密，1男，2女)
//	Level        int8      `orm:"default(1);column(level)"json:"level"`                     //允许为空，会员等级,1-9级,默认是1级
//	Score        int       `orm:"default(0);column(score)"json:"score"`                     //积分
//	Status       int8      `orm:"default(0);column(status)"json:"status"`                   //状态 0 表示注册未激活 1表示正常 2表示禁用  3表示删除
//	InvitePeople int64     `orm:"default(0);column(invite_people)"json:"invite_people,omitempty"`     //邀请人
//	InviteCode   string    `orm:"type(char);size(8);column(invite_code)"json:"invite_code"` //邀请码8位
//	Pwd          string    `orm:"type(char);size(32);column(pwd)"json:"pwd"`                //密码
//	Email        string    `orm:"size(64);column(email)"json:"email"`                       //电子邮箱
//	IP           string    `orm:"size(18);column(ip)"json:"ip"`                             //登录的ip
//	Phone        string    `orm:"type(char);size(11);column(phone)"json:"phone"`            //手机号
//	Avatar       string    `orm:"size(128);column(avatar)"json:"avatar"`                    //头像
//	AuthTime     time.Time `orm:"column(auth_time);null"json:"auth_time"`                   //授权时间
//	LoginTime    time.Time `orm:"auto_now_add;column(login_time)"json:"login_time"`         //登录时间
//	CreatedTime  time.Time `orm:"auto_now_add;column(created_time)"json:"created_time"`     //注册时间
//	AuthAddr		string   `json:"auth_addr"`   //授权地址
//	Total		string   `json:"total"`   //下注总数
//}

// 会员管理页面数据
//func (e *User) UserList(param map[string]interface{}) (result interface{}, count int32, err error) {
//	//拼凑筛选条件sql
//	con := ""
//	con1:=""
//	//搜索框
//	if param["keyword"] != nil {
//		keyword := param["keyword"].(string)
//		if keyword != "" {
//			con += ` and (u.id like '%` + keyword + `%' or invite_people like '%` + keyword + `%' or email like '%` + keyword + `%' or auth_addr like '%` + keyword + `%') `
//			con1 += ` and (user_id like '%` + keyword + `%') `
//		}
//	}
//	//状态
//	if param["status"] != nil {
//		status := param["status"].(string)
//		if param["status"] != "" {
//			con += ` and u.Status = ` + status
//		}
//	}
//	if param["level"] != nil {
//		level := param["level"].(string)
//		if param["level"] != "" {
//			con += ` and u.level = ` + level
//		}
//	}
//	//分页 and 排序
//	param["sort"] = "id"
//	param["order"] = "desc"
//	limit := utils.LimitAndOrderBy(param)
//	//查询数据
//	sql := ` select *,
//				(select sum(ifnull(count,0))total from dice_record d where d.user_id = u.id and symbol ='GKC' group by user_id,symbol ) total
//					from user u
//					where 1=1 `+con
//	user := make([]User, 0)
//	orm.Eloquent.Raw(sql+limit).Scan(&user)
//
//	result = user
//	//总数
//	count = GetTotalCount(orm.Eloquent, sql)
//	return
//}
//
//// 会员管理页面数据
//func (e *User) GetPage(pageSize int, pageIndex int, param map[string]interface{}) ([]User, int32, error) {
//	var members []User
//	table := orm.Eloquent.Table("user")
//
//	var count int32
//
//	if err := table.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&members).Error; err != nil {
//		return nil, 0, err
//	}
//	table.Count(&count)
//
//	return members, count, nil
//}
//
//// 删除会员
//func (e *User) UserDelete(param map[string]interface{}) (user []User, err error) {
//	ids := param["ids"].(string)
//	if ids == "" {
//		return
//	}
//	id := strings.Replace(ids, ",", `','`, -1)
//	sql := ` UPDATE user SET is_del=1 where id in('` + id + `') `
//
//	orm.Eloquent.Exec(sql)
//	return
//}
//
//// 编辑会员
//func (e *User) UserEdit(param map[string]interface{}) (user []User, err error) {
//	id := param["id"].(string)
//	if id == "" {
//		return
//	}
//	con := ``
//	for key, val := range param {
//		if val != "" && key != "id" {
//			con +=  key + "='" + val.(string) + "',"
//		}
//	}
//	sql := ` UPDATE user SET ` + con[:len(con)-1] + ` where id =? `
//	orm.Eloquent.Exec(sql, id)
//	return
//}
//
////会员列表导出
//func (e *User) ExportExcelUrl(param map[string]interface{}) (URL string, err error) {
//	param["isexp"] = "1"
//	param["sheet"] = "sheet1"
//	param["filefield"] = "id,email,auth_addr,auth_time,invite_code,invite_people,user_level,user_status,total"
//	param["filename"] = "用户账号,邮箱,授权地址,授权时间,邀请码,邀请人,会员等级,状态,投注总额"
//	param["title"] = "会员列表"
//	URL, err = GetExcelURL(e.UserList, param)
//	return
//}
//
////会员总数and今日注册人数
//func (e *User) GetUserCount(param map[string]interface{}) (result interface{}, err error) {
//	sql := ` select count(1) usercount,(
//				select count(1) from user where FROM_UNIXTIME(created_time,'%Y-%m-%d') = date(now())
//			) todaycount from user  `
//	type UserCount struct {
//		Usercount  string
//		Todaycount string
//	}
//	var u UserCount
//	orm.Eloquent.Raw(sql).Scan(&u)
//	result = u
//
//	return
//}
////获取用户邀请所有人
//func (e *User) GetUserInvitePeople(param map[string]interface{}) (result interface{},count int32, err error) {
//	con := ""
//	if param["id"]!=nil{
//		id := param["id"].(string)
//		if id!=""{
//			con += ` and u.invite_people = `+id
//		}
//	}
//	//分页
//	limit := utils.LimitAndOrderBy(param)
//	sql := ` select *,
//				(select sum(ifnull(count,0))total from dice_record d where d.user_id = u.id and symbol ='GKC' group by user_id,symbol ) total
//					from user u
//					where 1=1 `+con
//	var u []User
//	orm.Eloquent.Raw(sql+limit).Scan(&u)
//	result = u
//	count = GetTotalCount(orm.Eloquent,sql)
//	return
//}
////授权地址
//func (e *User) GetAuthAddr(param map[string]interface{}) (result string, err error) {
//	if param["cid"]!=nil{
//		id := param["cid"].(string)
//		if id==""{
//			return "",errors.New("cid不能为空！")
//		}
//	}
//	sql := ` select * from user where id=? `
//	var u User
//	orm.Eloquent.Raw(sql,param["cid"]).Scan(&u)
//	result = u.AuthAddr
//	return
//}
