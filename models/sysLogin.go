package models

import (
	"errors"
	"log"
	"strings"
	orm "hechuangfil-admin/database"
	"hechuangfil-admin/pkg"
)

//登录
type Login struct {
	Username string `form:"UserName" json:"username" binding:"required"`
	Password string `form:"Password" json:"password" binding:"required"`
	Code     string `form:"Code" json:"code" binding:"required"`
	UUID     string `form:"UUID" json:"uuid" binding:"required"`
}

//管理员登录获取信息
func (u *Login) GetUser() (user SysUser, role SysRole, e error) {

	e = orm.Eloquent.Table("sys_user").Where("UserName = ? and is_del =0  and status = 0 ", u.Username).Find(&user).Error
	if e != nil {
		if strings.Contains(e.Error(), "record not found") {
			e = errors.New("账号不存在或已经停用")
		}
		log.Print(e)
		return
	}
	if user == (SysUser{}){
		e = errors.New("账号不存在或已经停用")
		return
	}
	_, e = pkg.CompareHashAndPassword(user.Password, u.Password)
	if e != nil {
		if strings.Contains(e.Error(), "hashedPassword is not the hash of the given password") {
			e = errors.New("账号或密码错误")
		}
		log.Print(e)
		return
	}
	e = orm.Eloquent.Table("sys_role").Where("id = ? ", user.RoleId).First(&role).Error
	if e != nil {
		log.Print(e.Error())
		return
	}
	return
}
