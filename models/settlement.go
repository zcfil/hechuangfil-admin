package models

import (
	orm "hechuangfil-admin/database"
	"hechuangfil-admin/utils"
	"time"
)

type SettleLog struct {
	Id 				int64 `json:"id"`
	CustomerId 		int64  `json:"customer_id"`
	TotalIncome 	float64			`json:"total_income"`
	CustomerIncome 	float64	`json:"customer_income"`
	CompanyIncome 	float64		`json:"company_income"`
	FilialeIncome 	float64	`json:"filiale_income"`
	ToCustomerBalance float64	`json:"to_customer_balance"`
	ToCustomerLock 	float64	`json:"to_customer_lock"`
	CustomerLockRelease float64			`json:"customer_lock_release"`
	SettleDateId 	int64			`json:"settle_date_id"`
	Time 			time.Time			`json:"time"`
}
type SettleLogView struct {
	SettleLog
	Name string `json:"name"`
	Phone string `json:"phone"`
}

func (y *SettleLog) SettleLogList(param map[string]string) ( interface{}, error) {

	sql := `select s.*,c.name,c.phone from settle_log s
		left join customer c on s.customer_id = c.customer_id 
		where 1=1 `
	if param["keyword"]!=""{
		sql += ` and (c.phone like '%` + param["keyword"] + `%' or c.name '%` + param["keyword"] + `%') `
	}
	param["total"] = GetTotalCount(sql)

	param["sort"] = "time"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)
	var res []SettleLogView
	if err := orm.Eloquent.Raw(sql).Scan(&res).Error;err != nil {
		return nil,err
	}
	return res,nil
}