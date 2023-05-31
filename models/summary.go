package models

import (
	orm "hechuangfil-admin/database"
	"hechuangfil-admin/utils"
)
type Summary struct {
	CustomerId		string `gorm:"column:customer_id" json:"customer_id"`
	CustomerIncome	float64 `gorm:"column:customer_income" json:"customer_income"`
	CompanyIncome	float64 `gorm:"column:company_income" json:"company_income"`
	FilialeIncome	float64`gorm:"column:filiale_income" json:"filiale_income"`
	ToCustomerBalance	float64 	`gorm:"column:to_customer_balance" json:"to_customer_balance"`
	ToCustomerLock	float64 	`gorm:"column:to_customer_lock" json:"to_customer_lock"`
	ReferrerProfit	float64 `gorm:"column:referrer_profit" json:"referrer_profit"`
	SalesProfit	float64 `gorm:"column:sales_profit" json:"sales_profit"`
	Phone 		string	`json:"phone"`
}
type SummaryTotal struct {
	Company	float64 `gorm:"column:company" json:"company"`
	Salesman	float64 `gorm:"column:salesman" json:"salesman"`
	Customerprofit	float64 `gorm:"column:customerprofit" json:"customerprofit"`
	Amount	float64 `gorm:"column:amount" json:"amount"`
	Total	string `gorm:"column:total" json:"total"`
}

//汇总报表
func (e *Summary) StatementSummary(param map[string]string) (interface{}, error) {
	//concat(LAST_DAY('2021-08-01'),' 23:59:59')
	sql := `select a.*
			,b.profits referrer_profit
			,b.salesdep_profit sales_profit
			,b.amount from (	
			select s.customer_id,c.phone,sum(customer_income)customer_income,sum(company_income)company_income,sum(filiale_income)filiale_income
			,sum(to_customer_balance+customer_lock_release)to_customer_balance,sum(to_customer_lock)to_customer_lock
		
			from settle_log s
			left join customer c on c.customer_id = s.customer_id
			GROUP BY s.customer_id
			) a 
			left join (select o.customer_id,sum(profits)profits,sum(os.amount)amount,max(os.salesdep_profit)salesdep_profit from ordersprofit o
											left join orders os on o.order_id = os.order_id
											GROUP BY o.customer_id
			)b on a.customer_id = b.customer_id
			`
	sql = utils.SqlReplaceParames(sql,param)
	param["total"] = GetTotalCount(sql)
	param["sort"] = "a.customer_id"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)
	var su []Summary
	if err := orm.Eloquent.Raw(sql).Scan(&su).Error;err!=nil{
		return nil, err
	}

	return su,nil
}
