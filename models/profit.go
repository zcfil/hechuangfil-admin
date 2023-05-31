package models

import (
	"errors"
	orm "hechuangfil-admin/database"
	"hechuangfil-admin/utils"
	"strings"
	"time"
)

type Profit struct {
	ID           string     `gorm:"column:id" json:"id"`
	Percent    	 float64    `gorm:"column:percent" json:"percent"`               //
	Userid     	string    `gorm:"column:userid" json:"userid"`               //
	Profittype        int    `gorm:"column:profittype" json:"profittype"`            //
	Profitlevel		float64    `gorm:"column:profitlevel" json:"profitlevel"`
	//A			[]Aa			`gorm:"column:a" json:"a"`
	//IsDel		int			`gorm:"column:is_del" json:"is_del"`	//是否删除
	NickName        string    `gorm:"column:nick_name" json:"nick_name"`            //
	Username		string    `gorm:"column:username" json:"username"`
	CreateTime	time.Time			`gorm:"column:create_time" json:"create_time"`	//创建时间
	UpdateTime	time.Time			`gorm:"column:update_time" json:"update_time"`	//创建时间
}
type Aa struct {
	Name string `json:"name"`
	Sex string `json:"sex"`
}
func (e *Profit) ProfitconfigList(param map[string]string) (result interface{}, err error) {
	//状态
	sql := `select * from user_level `

	user := make([]Profit, 0)
	orm.Eloquent.Raw(sql).Scan(&user)

	result = user

	return
}

func (e *Profit) ProfitconfigOnce(param map[string]string) (err error) {

	//param["id"] = strconv.FormatInt(utils.Node().Generate().Int64(),10)
	var count int
	orm.Eloquent.Table("profitconfig").Where("userid = ? and is_del =0 ", param["userid"]).Count(&count)
	if count > 0 {
		err = errors.New("账户已存在！")
		return
	}
	sql := ` insert into profitconfig(percent,profittype,userid )value(:percent,0,:userid)`
	sql = utils.SqlReplaceParames(sql,param)
	if err = orm.Eloquent.Exec(sql).Error;err!=nil{
		return err
	}

	return
	//return orm.Eloquent.Table("customer").Create(&e).Error
}
func (e *Profit) DelProfitconfigOnce(param map[string]string) (err error) {

	sql := ` delete from profitconfig where id =:id and profittype =0 `
	sql = utils.SqlReplaceParames(sql,param)
	if err = orm.Eloquent.Exec(sql).Error;err!=nil{
		return err
	}

	return
	//return orm.Eloquent.Table("customer").Create(&e).Error
}
func (e *Profit) UpdateProfitconfigOnce(param map[string]string) (err error) {

	//param["id"] = strconv.FormatInt(utils.Node().Generate().Int64(),10)

	sql := ` update profitconfig set percent=:percent,userid=:userid where id=:id `
	sql = utils.SqlReplaceParames(sql,param)
	if err = orm.Eloquent.Exec(sql).Error;err!=nil{
		return err
	}

	return
	//return orm.Eloquent.Table("customer").Create(&e).Error
}
type ProfitEdit struct {
	Ids string `json:"ids"`
	//Userid string `json:"userid"`
	Profit []ProfitConfig `json:"profit"`
}
type ProfitConfig struct {
	//ID           string     `gorm:"column:id" json:"id"`
	Percent    	 float64    `gorm:"column:percent" json:"percent"`               //
	//Userid     	string    `gorm:"column:userid" json:"userid"`               //
	//Profittype        int    `gorm:"column:profittype" json:"profittype"`            //
	Levelvalue        float64    `gorm:"column:levelvalue" json:"levelvalue"`
	//A			[]Aa			`gorm:"column:a" json:"a"`
	//IsDel		int			`gorm:"column:is_del" json:"is_del"`	//是否删除
	//CreateTime	time.Time			`gorm:"column:create_time" json:"create_time"`	//创建时间
}
func (e *ProfitEdit) ProfitEdit() (err error) {
	orm1 := orm.Eloquent.Begin()

	//增加新配置
	flag := false
	sql1 := `insert into user_level(levelvalue, percent, percentreality)values `

	percentReality := 0.0
	for k,v := range e.Profit{
		flag = true

		percentReality += v.Percent
		sql1 +=`(`+utils.Float64ToString(v.Levelvalue) + "," +utils.Float64ToString(v.Percent)+ "," + utils.Float64ToString(percentReality) + `)`
		if k < len(e.Profit)-1{
			sql1 += ","
		}
	}
	//删除原有配置
	sql := ` delete from user_level where id in ('`+strings.Replace(e.Ids,",","','",-1)+`')`
	if err = orm1.Exec(sql).Error;err!=nil{
		orm1.Rollback()
		return err
	}
	if flag{
		if err = orm1.Exec(sql1).Error;err!=nil{
			orm1.Rollback()
			return err
		}
	}
	orm1.Commit()
	return
}
