package models

import (
	orm "hechuangfil-admin/database"
	"strconv"
	"time"
)

type Order struct {
	OrderID   		int64			`gorm:"column:order_id" json:"order_id"`		// 订单id
	Amount    		float64			`gorm:"column:amount" json:"amount"`			// 下单金额
	CreateTime   	time.Time		`gorm:"column:create_time" json:"create_time"`	// 创建时间
	Phone		  	string			`gorm:"column:phone" json:"phone"`				// 用户手机号
	Hashrate 		int64			`gorm:"column:hashrate" json:"hashrate"`		// 算力
	SalesdepProfit  float64			`gorm:"column:salesdep_profit" json:"salesdep_profit"` // 业务部分润
}


func NewOrder() *Order {
	order := &Order{}
	return order
}

func (this *Order) GetOrders(param map[string]string) (res interface{}, err error) {
	sql1 := `select * from orders`
	count := GetTotalCount(sql1)

	param["total"] = count
	pageSize, err1 := strconv.ParseInt(param["pageSize"], 10, 64)
	if err1 != nil {
		err = err1
		return
	}
	pageIndex, err1 := strconv.ParseInt(param["pageIndex"], 10, 64)
	if err1 != nil {
		err = err1
		return
	}

	// 关联客户表取出电话号码
	sql2 := `select o.order_id, o.amount, o.create_time, o.hashrate, o.salesdep_profit, c.phone 
				from orders as o left join customer as c on c.customer_id = o.customer_id order by o.create_time desc`

	start := (pageIndex-1)*pageSize
	sql2 += ` limit ` + strconv.FormatInt(start, 10) + `, ` + param["pageSize"]

	finds := make([]*Order, 0)
	if err = orm.Eloquent.Raw(sql2).Scan(&finds).Error; err != nil {
		return
	}

	res = finds
	return
}