package models

import (
	orm "hechuangfil-admin/database"
	"log"
)

var CidWithdraw = make(map[string]bool)
func RunInit(){
	//var conf Config
	//c ,err:= conf.GetConfig("customerratio")
	//if err!=nil||(c == Config{}){
	//	sql := `insert into sys_config(configKey,configValue)value('customerratio',?),('salesmanratio',?)`
	//	if err = orm.Eloquent.Exec(sql,config.ApplicationConfig.CustomerRatio,config.ApplicationConfig.SalesmanRatio).Error;err!=nil{
	//		log.Fatal(err)
	//	}
	//}
	sql := `select cid from withdraw where status = 1 and create_time >= DATE_SUB(now(),INTERVAL 7 day)`
	type Cids struct {
		Cid string
	}
	var c []Cids
	if err := orm.Eloquent.Raw(sql).Scan(&c).Error;err!=nil{
		log.Fatal("启动失败：",err)
	}
	for _,v := range c {
		CidWithdraw[v.Cid] = true
	}

}