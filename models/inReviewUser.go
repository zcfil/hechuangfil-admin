package models

import (
	"strconv"
	"time"
	orm "hechuangfil-admin/database"
)

// 审核中的用户信息
type InReviewUser struct {
	UserID   	 	int64 			`gorm:"column:user_id" json:"user_id"`
	UserName  		string   		`gorm:"column:username" json:"username"`
	NickName  		string			`gorm:"column:nick_name" json:"nick_name"`
	DeptID          int				`gorm:"column:dept_id" json:"dept_id"`						// 部门ID
	Phone 			string 			`gorm:"column:phone" json:"phone"`
	ReferrerName    string 			`gorm:"column:referrer_name" json:"referrer_name"`			// 推荐者名字
	CreateTime 		time.Time 		`gorm:"column:create_time" json:"create_time"`				// 创建时间
}


func (this *InReviewUser) GetUserPassingList(pageIndex, pageSize int64, roleKey, curRoleID string) (list interface{}, total string,  err error)  {
	sql := `select u.user_id, u.username, u.nick_name, u.dept_id, u.phone, u1.username as referrer_name, u.create_time from sys_user u
			left join sys_user u1 on u.referrer = u1.user_id
 			where u.is_pass = 0 and u.is_del != 1`

	if roleKey == "common" {
		sql += ` and u.referrer = `
		sql += curRoleID
	}

	total = GetTotalCount(sql)
	start := (pageIndex-1) * pageSize

	sql += ` limit `
	sql += strconv.FormatInt(start, 10)
	sql += `, `
	sql += strconv.FormatInt(pageSize, 10)
	finds := make([]InReviewUser, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&finds).Error; err != nil {
		return
	}
	list = finds
	return
}