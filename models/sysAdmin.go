package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strconv"
	"strings"
	"time"
	orm "hechuangfil-admin/database"
	"hechuangfil-admin/pkg"
	"hechuangfil-admin/utils"
)

// SysUser
type AdminB struct {
	// key
	IdentityKey string
	// 用户名
	UserName  string
	FirstName string
	LastName  string
	// 角色
	Role string
}
type Referrer struct {
	//Id int64
	Userid int64
	Referrers string
	Referrals string
}

type UserName struct {
	Username string `gorm:"column:username" json:"username"`
}

type PassWord struct {
	// 密码
	Password string `gorm:"column:password" json:"password"`
}

type LoginM struct {
	UserName
	PassWord
}

type SysUserId struct {
	// 编码
	Id int64 `gorm:"column:user_id;primary_key"  json:"userId"`
}

type Admin struct {
	// 昵称
	NickName string `gorm:"column:nick_name" json:"nickName"`
	// 手机号
	Phone string `gorm:"column:phone" json:"phone"`
	// 角色编码
	RoleId int64 `gorm:"column:role_id" json:"roleId"`
	//盐
	Salt string `gorm:"column:salt" json:"salt"`
	//头像
	Avatar string `gorm:"column:avatar" json:"avatar"`
	//性别
	Sex string `gorm:"column:sex" json:"sex"`
	//邮箱
	Email string `gorm:"column:email" json:"email"`
	// 创建时间
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	// 修改时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`

	//部门编码
	DeptId int64 `gorm:"column:dept_id" json:"deptId"`

	//职位编码
	PostId int64 `gorm:"column:post_id" json:"postId"`

	CreateBy string `gorm:"column:create_by" json:"createBy"`
	UpdateBy string `gorm:"column:update_by" json:"updateBy"`

	//备注
	Remark       string `gorm:"column:remark" json:"remark"`
	Params       string `gorm:"column:params" json:"params"`
	Status       string `gorm:"column:status" json:"status"`
	DataScope    string `gorm:"column:dataScope" json:"dataScope"`
	IsDel        int `gorm:"column:is_del" json:"isDel"`
	Referrer     string `gorm:"column:referrer" json:"referrer"`
	//Referrers     string `gorm:"column:referrers" json:"referrers"`
	//Verification string `gorm:"type:varchar(255)" json:"verification"` //google验证

	// 银行相关
	BankUserName	string   `gorm:"column:bank_user_name" json:"bank_user_name"`		// 银行户名
	BankNum 		string 	 `gorm:"column:bank_num" json:"bank_num"`					// 银行卡号
	OpenBank 	    string 	 `gorm:"column:open_bank" json:"open_bank"`					// 开户行

	IsPass          int      		`gorm:"column:is_pass"`									// 是否通过审核
	PassTime 	    time.Time 		`gorm:"column:pass_time"`									// 审核通过时间
}

type SysUser struct {
	SysUserId
	Admin
	LoginM
}

type SysAdminPwd struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type SysAdminPage struct {
	SysUserId
	Admin
	LoginM
	DeptName string `gorm:"-" json:"deptName"`
}

type SysAdminView struct {
	SysUserId
	Admin
	LoginM
	RoleName string `gorm:"column:role_name"  json:"role_name"`
	ReferrerName string `gorm:"column:referrer_name" json:"referrer_name"`
}
func GetSysUserById(id string) (a SysAdminView) {
	sql := `select * from sys_user where user_id = ?`
	orm.Eloquent.Raw(sql,id).Scan(&a)
	return
}
func (r *Referrer) GetReferrer(userid string) (a Referrer) {
	sql := `select * from referrer where userid = ?`
	orm.Eloquent.Raw(sql,userid).Scan(&a)
	return
}
// 获取用户数据
func (e *SysUser) Get() (SysUserView SysAdminView, err error) {

	//table := orm.Eloquent.Table("sys_user").Select([]string{"b.*", "sys_role.name"})
	//table = table.Joins("left join sys_role on sys_user.role_id=sys_role.id")
	sql := `select a.*,r.name,b.nick_name referrer_name from sys_user a
			left join sys_role r on a.role_id=r.id
			left join sys_user b on a.referrer = b.user_id
			where 1=1 `
	if e.Id != 0 {
		sql += ` and a.user_id= `+strconv.FormatInt(e.Id,10)
	}

	if e.Username != "" {
		sql += ` and a.username= `+e.Username
		//table = table.Where("username = ?", e.Username)
	}

	if e.Password != "" {
		sql += ` and a.password= `+e.Password
		//table = table.Where("password = ?", e.Password)
	}

	if e.RoleId != 0 {
		sql += ` and a.role_id= `+strconv.FormatInt(e.RoleId,10)
		//table = table.Where("role_id = ?", e.RoleId)
	}

	if e.DeptId != 0 {
		sql += ` and a.dept_id= `+strconv.FormatInt(e.DeptId,10)
		//table = table.Where("dept_id = ?", e.DeptId)
	}

	if e.PostId != 0 {
		sql += ` and a.post_id= `+strconv.FormatInt(e.PostId,10)
		///table = table.Where("post_id = ?", e.PostId)
	}

	//if err = table.First(&SysUserView).Error; err != nil {
	//	return
	//}
	err = orm.Eloquent.Raw(sql).Scan(&SysUserView).Error
	return
}

//获取系统用户分页数据
func (e *SysUser) GetPage(pageSize int, pageIndex int) ([]SysAdminPage, int32, error) {

	log.Println("获取系统用户分页数据 GetPage ***************")
	var doc []SysAdminPage

	table := orm.Eloquent.Select("sys_user.*,sys_dept.dept_name").Table("sys_user")
	table = table.Joins("left join sys_dept on sys_dept.deptId = sys_user.dept_id")

	if e.Username != "" {
		table = table.Where("username = ?", e.Username)
	}
	if e.Phone != "" {
		table = table.Where("sys_user.phone = ?", e.Phone)
	}


	if e.DeptId != 0 {
		table = table.Where("dept_id in (select deptId from sys_dept where dept_path like ? )", "%"+utils.Int64ToString(e.DeptId)+"%")
	}

	// 数据权限控制
	//dataPermission := new(DataPermission)
	//dataPermission.UserId, _ = utils.StringToInt64(e.DataScope)
	//table = dataPermission.GetDataScope("sys_user", table)

	var count int32

	if err := table.Where("sys_user.is_del = 0").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("sys_user.is_del = 0").Count(&count)
	return doc, count, nil
}

//加密
func (e *SysUser) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

//添加
func (e SysUser) Insert() (id int64, err error) {
	e.CreateTime = time.Now()
	e.UpdateTime = time.Now()
	e.PassTime = time.Now()
	e.IsPass = PASS_VERIFICATION
	if err = e.Encrypt(); err != nil {
		return
	}
	orm1 := orm.Eloquent.Begin()
	defer func() {
		if err!=nil{
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()
	// check 用户名
	var count int
	orm.Eloquent.Table("sys_user").Where("username = ?", e.Username).Count(&count)
	if count > 0 {
		err = errors.New("账户已存在！")
		return
	}
	var ref Referrer
	if e.Referrer!=""{
		sql := `select * from referrer where userid=?`
		if err =orm1.Raw(sql,e.Referrer).Scan(&ref).Error; err != nil {
			return
		}
		if len(ref.Referrers)>0{
			ref.Referrers += ","
		}
		ref.Referrers += e.Referrer

		if err = orm1.Table("sys_user").Create(&e).Error; err != nil {
			return
		}
		id = e.Id

		sql2 := `update referrer set referrals = CONCAT(ifnull(referrals,''),',`+strconv.FormatInt(id,10)+`') where userid in (`+ref.Referrers+`) `
		if err = orm1.Exec(sql2).Error;err!=nil{
			return
		}
		sql3 := `insert into referrer (userid,referrers)value(`+strconv.FormatInt(id,10)+`,'`+ref.Referrers+`') `
		//e.Referrers = ref.Referrers
		if err = orm1.Exec(sql3).Error;err!=nil{
			return
		}
		return
	}
	e.Referrer = "0"
	//添加数据
	if err = orm1.Table("sys_user").Create(&e).Error; err != nil {
		return
	}
	id = e.Id
	sql3 := `insert into referrer (userid)value(`+strconv.FormatInt(id,10)+`) `
	//e.Referrers = ref.Referrers
	if err = orm1.Exec(sql3).Error;err!=nil{
		return
	}
	// 添加客户数据 （业务员也要添加到客户）
	sex := "女"
	if e.Sex == "0" {
		sex = "男"
	}
	sql4 := ` insert into customer(name,sex,phone,bank,banknum,userid,status )values("`+
		e.Username + `","` +
		sex + `","` +
		e.Phone + `","` +
		e.OpenBank + `","` +
		e.BankNum + `",` +
		strconv.FormatInt(e.Id, 10) + `,` +
		`0`
	sql4 += `)`
	if err = orm1.Exec(sql4).Error;err!=nil{
		return
	}

	return
}

//修改
func (e *SysUser) Update(id int64) (update SysUser, err error) {
	e.UpdateTime = time.Now()
	if err = e.Encrypt(); err != nil {
		return
	}
	idstr := strconv.FormatInt(id,10)
	if err = orm.Eloquent.Table("sys_user").First(&update, id).Error; err != nil {
		return
	}
	if e.RoleId == 0 {
		e.RoleId = update.RoleId
	}
	orm1 := orm.Eloquent.Begin()
	defer func() {
		if err!=nil{
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()
	if update.Referrer != e.Referrer && e.Referrer!=""{
		var old Referrer
		sql := `select * from referrer where userid = '`+update.Referrer+`'`
		if err = orm1.Raw(sql).Scan(&old).Error;err!=nil{
			if err == gorm.ErrRecordNotFound{
				rid := update.Referrer
				if update.Referrer ==""{
					rid = e.Referrer
				}
				//处理老数据
				sql00 := `insert into referrer(userid,referrals)value(`+rid+`,`+idstr+`)`
				if err = orm1.Exec(sql00).Error;err!=nil{
					return
				}
				old.Userid,_ = strconv.ParseInt(update.Referrer,10,64)
				old.Referrals = idstr
			}else{
				return
			}
		}
		fmt.Println(old)
		//修改原来上级的下级
		var olds []Referrer
		if old.Referrers !=""{
			sql1 := `select * from referrer where userid in (`+old.Referrers+`)`
			if err = orm1.Raw(sql1).Scan(&olds).Error;err!=nil{
				return
			}
		}
		olds = append(olds, old)
		for _,v := range olds{
			oldref := strings.Split(v.Referrals,",")
			als := ""
			for i,val := range oldref{
				if val != idstr{
					als += val
				}
				if i < len(oldref)-1{
					als += ","
				}
			}
			sql0 := `update referrer set referrals =  '`+als+`',update_time=now() where userid=`+strconv.FormatInt(v.Userid,10)
			if err = orm1.Exec(sql0).Error;err!=nil{
				return
			}
		}

		var new1 Referrer
		sql3 := `select * from referrer where userid = `+e.Referrer
		if err = orm1.Raw(sql3).Scan(&new1).Error;err!=nil{
			if err == gorm.ErrRecordNotFound{

				//处理老数据
				sql00 := `insert into referrer(userid,referrals)value(`+e.Referrer+`,`+idstr+`)`
				if err = orm1.Exec(sql00).Error;err!=nil{
					return
				}
				old.Userid,_ = strconv.ParseInt(update.Referrer,10,64)
				old.Referrals = idstr
			}else{
				return
			}
		}
		var news []Referrer
		if new1.Referrers!=""{
			sql4 := `select * from referrer where userid in (`+new1.Referrers+`)`
			if err = orm1.Raw(sql4).Scan(&news).Error;err!=nil{
				return
			}
		}
		//修改新的上级的下级
		for _,v := range news{
			oldref := strings.Split(v.Referrals,",")
			als := ""
			for i,val := range oldref{
				if val != idstr{
					als += val
				}
				if i < len(oldref)-1{
					als += ","
				}
			}
			if len(als) > 0{
				als += ","
			}
			als += idstr
			sql0 := `update referrer set referrals =  '`+als+`',update_time=now() where userid=`+strconv.FormatInt(v.Userid,10)
			if err = orm1.Exec(sql0).Error;err!=nil{
				return
			}
		}
		//修改自己的上级
		ers := new1.Referrers
		if len(ers)>0{
			ers += ","
		}
		ers += e.Referrer
		sql0 := `update referrer set referrers ='`+ers +`',update_time=now() where userid = ` +idstr
		if err = orm1.Exec(sql0).Error;err!=nil{
			return
		}
		//修改自己的下级的上级
		var my Referrer
		sql5 := `select * from referrer where userid = (`+idstr+`)`
		if err = orm1.Raw(sql5).Scan(&my).Error;err!=nil{
			return
		}
		var mys []Referrer
		if my.Referrals!=""{
			sql6 := `select * from referrer where userid in (`+my.Referrals+`)`
			if err = orm1.Raw(sql6).Scan(&mys).Error;err!=nil{
				return
			}
		}

		for _,v := range mys{

			newers := strings.Replace(v.Referrers,old.Referrers,ers,-1)

			sql0 := `update referrer set referrals =  '`+newers+`',update_time=now() where userid=`+strconv.FormatInt(v.Userid,10)
			if err = orm1.Exec(sql0).Error;err!=nil{
				return
			}
		}

	}
	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm1.Table("sys_user").Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

func (e *SysUser) BatchDelete(id string) (Result bool, err error) {
	var mp = map[string]string{}
	mp["is_del"] = "1"
	mp["status"] = "1"
	mp["update_time"] = time.Now().Format("2006/01/02 15:04:05")
	mp["update_by"] = e.UpdateBy
	if err = orm.Eloquent.Table("sys_user").Where("is_del=0 and user_id in (?)", id).Update(mp).Error; err != nil {
		return
	}
	Result = true
	return
}

func (e *SysUser) SetPwd(pwd SysAdminPwd) (Result bool, err error) {
	user, _ := e.Get()
	_, err = pkg.CompareHashAndPassword(user.Password, pwd.OldPassword)
	if err != nil {
		if strings.Contains(err.Error(), "hashedPassword is not the hash of the given password") {
			pkg.AssertErr(err, "密码错误(代码202)", 500)
		}
		log.Print(err)
		//return
	}
	e.Password = pwd.NewPassword
	_, err = e.Update(e.Id)
	pkg.AssertErr(err, "更新密码失败(代码202)", 500)
	return
}

type User struct {
	UserId string `gorm:"column:user_id" json:"userid"`
	Username string `gorm:"column:username" json:"username"`
	NickName string `gorm:"column:nick_name" json:"nick_name"`
	Phone string `gorm:"column:phone" json:"phone"`
}
//获取系统用户信息
func (e *SysUser) GetSysList(param map[string]string) (interface{}, error) {
	sql := `select user_id,username,nick_name,phone from sys_user where 1=1 `
	keyword := param["keyword"]
	if keyword!="" {
		sql += ` and (phone like '%` + keyword + `%' or nick_name like '%` + keyword + `%'`+ ` or username like '%` + keyword + `%') `
	}
	param["total"] = GetTotalCount(sql)
	//分页 and 排序
	param["sort"] = "user_id"
	param["order"] = "desc"
	var u []User
	sql += utils.LimitAndOrderBy(param)
	err := orm.Eloquent.Raw(sql).Scan(&u).Error

	return u, err
}

func (e *SysUser) GetUserPassingList(pageIndex, pageSize int64) (list interface{}, total string,  err error)  {
	sql := `select * from sys_user where is_pass = 0`
	total = GetTotalCount(sql)
	start := (pageIndex-1) * pageSize

	sql += ` limit `
	sql += strconv.FormatInt(start, 10)
	sql += ` `
	sql += strconv.FormatInt(pageSize, 10)
	finds := make([]SysUser, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&finds).Error; err != nil {
		return
	}
	list = finds
	return
}

func (e *SysUser) SubmitNewUser(submitterID string) (err error) {
	timeNow := time.Now()
	e.CreateTime = timeNow
	e.UpdateTime = timeNow
	e.PassTime = timeNow
	e.IsPass = NO_PASS_VERIFICATION
	e.Referrer = submitterID

	if err = e.Encrypt(); err != nil {
		return
	}

	orm1 := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()

	// 先判断用户名
	sql := `select user_id from sys_user where username = "` + e.Username + `" limit 1`
	type find struct {
		UserID int64 `gorm:"column:user_id"`
	}
	findList := make([]find, 0)
	if err = orm1.Raw(sql).Scan(&findList).Error; err != nil {
		return
	}

	if len(findList) > 0 {
		err = errors.New("该用户名已存在")
		return
	}

	if err = orm1.Table("sys_user").Create(&e).Error; err != nil {
		return
	}


	// 只添加自己的上级，通过以后自己的上级再追加自己的ID到被推荐人
	sql2 := `select referrers from referrer where userid = ` + submitterID
	type findReferrer struct {
		Referrers string `gorm:"column:referrers"`
	}
	findReferrers := make([]findReferrer, 0)
	if err = orm1.Raw(sql2).Scan(&findReferrers).Error; err != nil {
		return
	}

	if len(findReferrers) != 1 {
		return errors.New("查找推荐人列表错误")
	}

	referrers := findReferrers[0].Referrers+"," + submitterID
	sql3 := `insert into referrer (userid, referrers, update_time) values (` + strconv.FormatInt(e.Id, 10) + `,"` + referrers + `", now())`
	if err = orm1.Exec(sql3).Error; err != nil {
		return
	}
	return
}

func (e *SysUser) AllowNewPass(userID string) (err error) {
	sql := `select * from sys_user where user_id = ` + userID
	if err = orm.Eloquent.Raw(sql).Scan(e).Error; err != nil {
		return
	}

	orm1 := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()
	sql1 := `update sys_user set is_pass = 1, pass_time = now() where user_id = ` + userID
	db := orm1.Exec(sql1)
	if db.Error != nil {
		err = db.Error
		return
	}

	if db.RowsAffected == 0 {
		return
	}

	// 为自己的上级添加新的下级用户
	type find struct {
		Referrer int64 	`gorm:"column:referrer"`
	}
	sql2 := `select referrer from sys_user where user_id = ` + userID
	var findMsg find
	if err = orm1.Raw(sql2).Scan(&findMsg).Error; err != nil {
		return
	}

	referrerStr := strconv.FormatInt(findMsg.Referrer, 10)
	var ref Referrer
	sql3 := `select * from referrer where userid=` + referrerStr
	if err = orm1.Raw(sql3).Scan(&ref).Error; err != nil {
		return
	}

	if len(ref.Referrers) > 0 {
		ref.Referrers += ","
	}
	ref.Referrers += referrerStr

	sql4 := `update referrer set referrals = CONCAT(ifnull(referrals,''),',`+userID+`') where userid in (`+ref.Referrers+`) `
	if err = orm1.Exec(sql4).Error;err!=nil{
		return
	}

	sql5 := `update referrer set referrers = "`+ref.Referrers+`" where userid = ` + userID
	if err = orm1.Exec(sql5).Error;err!=nil{
		return
	}

	// 添加客户数据 （业务员也要添加到客户）
	sex := "女"
	if e.Sex == "0" {
		sex = "男"
	}
	sql6 := ` insert into customer(name,sex,phone,bank,banknum,userid,status )values("`+
		e.Username + `","` +
		sex + `","` +
		e.Phone + `","` +
		e.OpenBank + `","` +
		e.BankNum + `",` +
		strconv.FormatInt(e.Id, 10) + `,` +
		`0`
	sql6 += `)`
	if err = orm1.Exec(sql6).Error;err!=nil{
		return
	}

	return
}

func (e *SysUser) NotAllowUserPass(userID string) (err error) {
	orm1 := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()
	// 拒绝删除用户数据
	sql1 := `update sys_user set is_del = 1 where user_id = ` + userID
	if err = orm1.Exec(sql1).Error; err != nil {
		return
	}

	// 删除上下级关系数据表的数据
	sql2 := `delete from referrer where userid = ` + userID
	if err = orm1.Exec(sql2).Error; err != nil {
		return
	}
	return
}

