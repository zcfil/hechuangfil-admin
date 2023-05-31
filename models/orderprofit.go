package models

import (
	orm "hechuangfil-admin/database"
	"hechuangfil-admin/utils"
	"time"
)
type OrderProfits struct {
	UserId int64	`json:"user_id"`
	Amount float64	`json:"amount"`
	Profits float64	`json:"profits"`
	OrderId	int64	`json:"order_id"`
	Phone	string	`json:"phone"`
	Hashrate int64	`json:"hashrate"`
	CreateTime time.Time `json:"create_time"`
}
//订单分润报表
func (e *OrderProfits) StatementOrder(param map[string]string) (interface{}, error) {
	//concat(LAST_DAY('2021-08-01'),' 23:59:59')
	sql := `select ip.user_id,round(i.amount,2)amount,round(ip.profits,2)profits,ip.order_id,if(ip.user_id='0','业务部',c.phone)phone,i.create_time,i.hashrate from orders i
      left join ordersprofit ip on i.order_id = ip.order_id
      left join customer c on ip.user_id = c.customer_id
      where i.create_time <= :end and :start <= i.create_time `
	sql = utils.SqlReplaceParames(sql,param)
	param["total"] = GetTotalCount(sql)
	param["sort"] = "i.create_time"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)
	var su []OrderProfits
	if err := orm.Eloquent.Raw(sql).Scan(&su).Error;err!=nil{
		return nil, err
	}

	return su,nil
}
