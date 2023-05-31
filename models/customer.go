package models

import (
	"errors"
	orm "hechuangfil-admin/database"
	"hechuangfil-admin/utils"
	"strconv"
	"time"
)

//客户表
type Customer struct {
	CustomerID   	int64    	`json:"customer_id" xorm:"customer_id"`
	Name 			string		`json:"name" xorm:"name"`
	Password 		string		`json:"password" xorm:"password"`
	Phone 			string		`json:"phone" xorm:"phone"`
	UserID          int			`json:"user_id" xorm:"user_id"`
	IsDel           int8		`json:"is_del" xorm:"is_del"`
	Status          int8		`json:"status"`
	CreateTime 		time.Time	`json:"create_time" xorm:"create_time"`
	UpdateTime 		time.Time	`json:"update_time" xorm:"update_time"`
	Sex 			int8		`json:"sex"`
	Identity 		string		`json:"identity"`
	Wallet 			string		`json:"wallet"`
	Balance 		float64		`json:"balance"`
}
func (e *Customer)NewCustomer(customerid string)(Customer,error){
	sql2 := `select * from customer where id = '`+customerid+"'"
	var c Customer
	if err := orm.Eloquent.Raw(sql2).Scan(&c).Error;err!=nil{
		return c,err
	}
	return c,nil
}

// 客户管理页面数据
func (e *Customer) CustomerList(param map[string]string) (result interface{}, err error) {
	//拼凑筛选条件sql
	sql := ` select * from customer c
				where c.is_del=0`
	//搜索框
	keyword := param["keyword"]
	if keyword != "" {
		sql += ` and (c.phone like '%` + keyword + `%' or c.name like '%` + keyword + `%') `
	}

	//if param["userid"] != "" {
	//	userID := param["userid"]
	//	filterUsers, err := e.getFilterUsers(userID)
	//	if err != nil {
	//		sql += ` and c.userid = ` + userID
	//	} else {
	//		sql += ` and c.userid in (` + filterUsers + `)`
	//	}
	//}
	//总数
	param["total"] = GetTotalCount(sql)
	//分页 and 排序
	param["sort"] = "id"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)

	user := make([]Customer, 0)
	orm.Eloquent.Raw(sql).Scan(&user)

	result = user

	return
}

func (e *Customer) getFilterUsers(userID string) (userIDs string, err error) {
	// 先判断部门，部门没有就取推荐人玩家列表
	userIDs, err = e.getDeptUsers(userID)
	if err == nil {
		return
	}
	return e.getUserReferralsIDs(userID)
}

func (e *Customer) getDeptUsers(userID string) (userIDs string, err error) {
	// 是否leader
	sql := `select deptId from sys_dept where leader_id=` + userID
	type findDept struct {
		Deptid int64 `gorm:"column:deptId;primary_key"`
	}
	var depts []findDept
	if err = orm.Eloquent.Raw(sql).Scan(&depts).Error; err != nil {
		return
	}

	if len(depts) <= 0 {
		err = errors.New("no records")
		return
	}

	deptStr := ""
	l := len(depts)
	for i, de := range depts {
		deptStr += strconv.FormatInt(de.Deptid, 10)
		if i != l-1 {
			deptStr += ","
		}
	}
	sql2 := `select user_id from sys_user where dept_id in (` + deptStr + `)`
	var findUsers []SysUserId
	if err = orm.Eloquent.Raw(sql2).Scan(&findUsers).Error; err != nil {
		return
	}

	userIDs = ""
	lu := len(findUsers)
	for i, u := range findUsers {
		userIDs += strconv.FormatInt(u.Id, 10)
		if i != lu-1 {
			userIDs += ","
		}
	}
	return
}

func (e *Customer) getUserReferralsIDs(userID string) (ret string, err error) {
	sql := `select referrals from referrer where userid=` + userID

	var referrers []Referrer
	if err = orm.Eloquent.Raw(sql).Scan(&referrers).Error; err != nil {
		return
	}
	if len(referrers) != 1 {
		err = errors.New("record error")
		return
	}
	ret = referrers[0].Referrals
	if len(ret) <= 0 {
		err = errors.New("no record")
		return
	}
	if len(ret) > 0 && ret[0] == ',' {
		ret = ret[1:]
	}
	ret += "," + userID  // 追加上自己的ID
	return
}
//	for key, val := range param {
//		if val != "" && key != "id" {
//			con +=  key + "='" + val.(string) + "',"
//		}
//	}
