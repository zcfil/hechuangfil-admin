package models

import (
	orm "hechuangfil-admin/database"
	"fmt"
	"strconv"
)

//角色菜单
type RoleMenu struct {
	RoleId   int64
	MenuId   int64
	RoleName string
	CreateBy string
	UpdateBy string
}

type MenuPath struct {
	Path string `json:"path"`
}

func (rm *RoleMenu) Get() ([]RoleMenu, error) {
	var r []RoleMenu
	table := orm.Eloquent.Table("sys_role_menu")
	if rm.RoleId != 0 {
		table = table.Where("role_id = ?", rm.RoleId)

	}
	if err := table.Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

func (rm *RoleMenu) GetPermis() ([]string, error) {
	var r []Menu
	table := orm.Eloquent.Select("sys_menu.permission").Table("sys_role_menu").Joins("left join sys_menu on sys_menu.menuId = sys_role_menu.menu_id")

	table = table.Where("role_id = ?", rm.RoleId)

	table = table.Where("sys_menu.menuType in('F','C')")
	if err := table.Find(&r).Error; err != nil {
		return nil, err
	}
	var list []string
	for i := 0; i < len(r); i++ {
		list = append(list, r[i].Permission)
	}
	return list, nil
}

func (rm *RoleMenu) GetIDS() ([]MenuPath, error) {
	var r []MenuPath
	table := orm.Eloquent.Select("sys_menu.path").Table("sys_role_menu")
	table = table.Joins("left join sys_role on sys_role.id=sys_role_menu.role_id")
	table = table.Joins("left join sys_menu on sys_menu.id=sys_role_menu.menu_id")
	table = table.Where("sys_role.name = ? and sys_menu.type=1", rm.RoleName)
	if err := table.Find(&r).Error; err != nil {
		return nil, err
	}
	return r, nil
}

//删除权限菜单
func (rm *RoleMenu) DeleteRoleMenu(roleId int64) (s bool, err error) {
	var role SysRole
	if err := orm.Eloquent.Table("sys_role").Where("id = ?", roleId).First(&role).Error; err != nil {
		return false, err
	}
	if role.RoleKey == "admin" {

		return false, err
	}

	if err := orm.Eloquent.Table("sys_role_dept").Where("role_id = ?", roleId).Delete(&rm).Error; err != nil {
		return false, err
	}
	if err := orm.Eloquent.Table("sys_role_menu").Where("role_id = ?", roleId).Delete(&rm).Error; err != nil {
		return false, err
	}

	sql3 := "delete from casbin_rule where v0= '" + role.RoleKey + "';"
	orm.Eloquent.Exec(sql3)

	return true, nil

}

//批量删除角色菜单
func (rm *RoleMenu) BatchDeleteRoleMenu(roleIds []int64) (bool, error) {

	var role []SysRole
	if err := orm.Eloquent.Table("sys_role").Where("id in (?)", roleIds).Find(&role).Error; err != nil {
		return false, err
	}

	if err := orm.Eloquent.Table("sys_role_menu").Where("role_name !='admin' and role_id in (?)", roleIds).Delete(&rm).Error; err != nil {
		return false, err
	}

	sql := ""
	for i := 0; i < len(role); i++ {
		if role[i].Name != "admin" {
			sql += "delete from casbin_rule where v0= '" + role[i].Name + "';"
		}
	}
	orm.Eloquent.Exec(sql)
	return true, nil

}

func (rm *RoleMenu) Insert(roleId int64, menuId []int64) (false bool,err error) {
	var role SysRole
	if err := orm.Eloquent.Table("sys_role").Where("id = ?", roleId).First(&role).Error; err != nil {
		return false, err
	}
	var menu []Menu
	if err := orm.Eloquent.Table("sys_menu").Where("menuId in (?)", menuId).Find(&menu).Error; err != nil {
		return false, err
	}
	orm1 := orm.Eloquent.Begin()
	defer func() {
		if  err!=nil{
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()
	//先删除原有权限
	sqldel := ` delete from sys_role_menu where role_id = ? `

	if err = orm.Eloquent.Exec(sqldel,roleId).Error;err!=nil{
		return
	}
	//ORM不支持批量插入所以需要拼接 sql 串
	sql := ""

	sql2 := ""
	for i := 0; i < len(menu); i++ {
		if len(menu)-1 == i {
			//最后一条数据 以分号结尾
			sql += fmt.Sprintf("(%d,%d,'%s');", role.Id, menu[i].MenuId, role.RoleKey)
			if menu[i].MenuType == "A" {
				sql2 += fmt.Sprintf("('p','%s','%s','%s');", role.RoleKey, menu[i].Path, menu[i].Action)
			}
		} else {
			sql += fmt.Sprintf("(%d,%d,'%s'),", role.Id, menu[i].MenuId, role.RoleKey)
			if menu[i].MenuType == "A" {
				sql2 += fmt.Sprintf("('p','%s','%s','%s'),", role.RoleKey, menu[i].Path, menu[i].Action)
			}
		}
	}
	if sql!=""{
		sql = "INSERT INTO `sys_role_menu` (`role_id`,`menu_id`,`role_name`) VALUES "+sql
		if err = orm1.Exec(sql).Error;err!=nil{
			return
		}
	}
	if sql2 != ""{
		sql2 = "INSERT INTO casbin_rule  (`p_type`,`v0`,`v1`,`v2`) VALUES "+sql2
		sql2 = sql2[0:len(sql2)-1] + ";"
		if err = orm1.Exec(sql2).Error;err!=nil{
			return
		}
	}

	//修改sys_role表role权限

	//sqlr := " update `sys_role` set role = ? ,update_time = now() where id = ? "
	//strmid:= ""
	//for i:=0;i<len(menuId);i++{
	//	strmid+= strconv.Itoa(int(menuId[i]))
	//	if i<len(menuId)-1{
	//		strmid += ","
	//	}
	//}
	//ERR = orm.Eloquent.Exec(sqlr,strmid,roleId).Error
	return true, err
}

func (rm *RoleMenu) Delete(RoleId string, MenuID string) (bool, error) {
	rm.RoleId, _ = strconv.ParseInt(RoleId, 10, 64)
	table := orm.Eloquent.Table("sys_role_menu").Where("role_id = ?", RoleId)
	if MenuID != "" {
		table = table.Where("menu_id = ?", MenuID)
	}
	if err := table.Delete(&rm).Error; err != nil {
		return false, err
	}
	return true, nil

}
