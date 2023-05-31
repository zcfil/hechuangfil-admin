package models

import (
	orm "hechuangfil-admin/database"
	"hechuangfil-admin/utils"
)

type Finance struct {
	// 代码
	ID    int64  `json:"id" gorm:"column:configId;primary_key"`     //编码
	ConfigKey  string `json:"name" gorm:"column:configKey;"` //参数名称 //参数键名ConfigKey string `json:"configKey" gorm:"column:configKey"`
	ConfigName string `json:"title" gorm:"column:configName"`           //变量标题  //参数名称ConfigName string `json:"Name" gorm:"column:name;primary_key"`
	ConfigValue      string `json:"value" gorm:"column:configValue"`           //参数变量值 	//参数键值 //ConfigValue string `json:"configValue" gorm:"column:configValue"`
	Balance      float64 `json:"balance"`
}

func (f *Finance)FinanceConfigList()([]Finance,error){
	sql := `select * from sys_config where is_del = 0`
	var fi []Finance
	err := orm.Eloquent.Raw(sql).Scan(&fi).Error
	return fi,err
}

func (f *Finance)UpdateConfigById(param map[string]string)([]Finance,error){
	sql := `update sys_config set configName=:title,configvalue=:value where configId = :id`
	sql = utils.SqlReplaceParames(sql,param)
	var fi []Finance
	err := orm.Eloquent.Raw(sql).Scan(&fi).Error
	return fi,err
}