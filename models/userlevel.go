package models

import (
	"strings"
	orm "hechuangfil-admin/database"
)

type UserLevel struct {
	Percentreality	float64 	 `gorm:"column:percentreality" json:"percentreality"`
	Id	string `gorm:"column:id" json:"id"`
	Accumulative	float64 `gorm:"column:accumulative" json:"accumulative"`
	Levelvalue	float64 `gorm:"column:levelvalue" json:"levelvalue"`
	Percent	float64 	`gorm:"column:percent" json:"percent"`
	Levelname	string `gorm:"column:levelname" json:"levelname"`
}

func (u *UserLevel)GetUserLevelList(ids string)([]UserLevel,error){
	con := `and user_id in (`+ids+`)`
	sql := `select u.user_id,u.accumulative,max(l.levelvalue) levelvalue,sum(l.percent) percent from sys_user u
					left join (SELECT * from user_level)l on u.accumulative >= l.levelvalue
					where 1=1 `+con+`
					GROUP BY user_id order by levelvalue`
	var ul []UserLevel
	err := orm.Eloquent.Raw(sql).Scan(&ul).Error
	for i:=0;i<len(ul);i++{
		if i < len(ul)-1{
			if ul[i].Levelvalue == ul[i+1].Levelvalue{

			}
		}
	}
	return ul,err
}
//排除一样等级的
func (u *UserLevel)GetUserLevel(ids string,levelvalue float64)([]UserLevel,error){
	con := `and user_id in (`+ids+`)`
	sql := `select u.user_id,u.accumulative, b.levelvalue, b.percent from sys_user u
					left join (SELECT * from user_level)l on u.accumulative >= l.levelvalue
					left join (
							select count(1) count,levelvalue,sum(percent) percent from (
							select u.user_id,u.accumulative,max(l.levelvalue) levelvalue,sum(l.percent) percent from sys_user u
							left join (SELECT * from user_level)l on u.accumulative >= l.levelvalue
							where 1=1 `+con+` and l.levelvalue > ?
							GROUP BY user_id order by levelvalue
							)a GROUP BY levelvalue 
					)b on b.levelvalue = l.levelvalue
					where b.count = 1 `+con+` and l.levelvalue > ?
					order by levelvalue`
	var ul []UserLevel
	err := orm.Eloquent.Raw(sql,levelvalue,levelvalue).Scan(&ul).Error
	return ul,err
}
func (u *UserLevel)GetReferrerLevel(ids string)([]UserLevel,error){
	sql := `select u.user_id,u.accumulative,max(l.levelvalue) levelvalue,sum(l.percent) percent from sys_user u
					left join (SELECT * from user_level)l on u.accumulative >= l.levelvalue
					where user_id in (`+ids+`) and is_del = 0 and status = 0
					GROUP BY user_id order by levelvalue`
	var ul []UserLevel
	err := orm.Eloquent.Raw(sql).Scan(&ul).Error
	//mp := make(map[string]UserLevel)
	//for i:=0;i<len(ul);i++ {
	//	mp[ul[i].UserId] = ul[i]
	//}
	//同级排序，直接上级在前面
	refs := strings.Split(ids,",")
	mp := make(map[string]int)
	for i := len(refs)-1;i>=0;i--{
		mp[refs[i]] = i
	}
	var res []UserLevel
	//一级只取一个
	for i:=0;i<len(ul);i++{
		pre := ul[i]
		for j:=i+1;j<len(ul);j++{
			if pre.Levelvalue < ul[j].Levelvalue{
				//res = append(res, ul[i])
				break
			}else{
				i++
			}
			if mp[pre.Id]<mp[ul[j].Id]{
				pre = ul[j]
			}
		}
		res = append(res, pre)
	}
	return res,err
}

//获取设置等级
func (u *UserLevel)GetSetUserLevel(levelvalue float64)([]UserLevel,error){
	sql := `select * from user_level where levelvalue>=? order by levelvalue `
	var ul []UserLevel
	err := orm.Eloquent.Raw(sql,levelvalue).Scan(&ul).Error
	return ul,err
}

//获取设置等级
func (u *UserLevel)GetSetUserByUserid(userid string)(UserLevel,error){
	sql := `select u.user_id,u.accumulative,max(l.levelvalue) levelvalue,sum(l.percent) percent, max(l.percentreality) percentreality
			from sys_user u
			left join (select * from user_level ) l on  u.accumulative >= l.levelvalue
			where u.user_id = `+userid+`
			GROUP BY user_id 
			order by levelvalue `
	var ul UserLevel
	err := orm.Eloquent.Raw(sql).Scan(&ul).Error
	return ul,err
}

