package models

import (
	orm "hechuangfil-admin/database"
	"hechuangfil-admin/utils"
	"time"
)

type Recharge struct {
	RechargeId string  `json:"recharge_id"`
	Cid string			`json:"cid"`
	ToAddress string	`json:"to_address"`
	Amount float64		`json:"amount"`
	FromAddress string	`json:"from_address"`
	CustomerId string	`json:"customer_id"`
	CreateTime time.Time	`json:"create_time"`
	Status int			`json:"status"`
}

func (r *Recharge)Insert(param map[string]string)(err error){
	orm1 := orm.Eloquent.Begin()
	defer func() {
		if err!=nil{
			orm1.Rollback()
			return
		}
		orm1.Commit()
		return
	}()
	sql := `insert into recharge(cid,to_address,amount,from_address,customer_id,height)values(:cid,:to_address,:amount,:from_address,:customer_id,:height) `
	sql = utils.SqlReplaceParames(sql,param)
	if err := orm1.Exec(sql).Error;err!=nil {
		return err
	}
	sql2 := `update customer set balance = balance + :amount where customer_id =:customer_id  `
	sql2 = utils.SqlReplaceParames(sql2,param)
	if err = orm1.Exec(sql2).Error;err!=nil {
		return err
	}

	return nil
}

//func (r *Recharge)Update(param map[string]string)(err error) {
//	orm1 := orm.Eloquent.Begin()
//	defer func() {
//		if err!=nil{
//			orm1.Rollback()
//			return
//		}
//		orm1.Commit()
//		return
//	}()
//	sql := `select * from recharge where to_address =:to_address and status = 0 `
//	sql = utils.SqlReplaceParames(sql,param)
//	var re []Recharge
//	if err = orm1.Scan(sql).Scan(&re).Error;err!=nil {
//		return err
//	}
//	if len(re) == 0 {
//		//无订单，新增订单
//		if err = r.Insert(param);err != nil{
//			return err
//		}
//		return nil
//	}
//	amount,_ := strconv.ParseInt(param["amount"],10,64)
//	min := utils.Abs(amount - re[0].Amount)
//	param["recharge_id"] = re[0].RechargeId
//	param["customer_id"] = re[0].CustomerId
//	for i:=1;i<len(re);i++{
//		if min < utils.Abs(amount-re[0].Amount){
//			param["recharge_id"] = re[i].Cid
//			param["customer_id"] = re[i].CustomerId
//		}
//	}
//	sql1 := `update recharge set cid =:cid,from_address=:from_address ,amount=:amount,status = 1,update_time=now(),height=:height where recharge_id =:recharge_id `
//	sql1 = utils.SqlReplaceParames(sql1,param)
//	if err = orm1.Exec(sql1).Error;err!=nil {
//		return err
//	}
//	sql2 := `update customer set balance = balance + :amount where customer_id =:customer_id  `
//	sql2 = utils.SqlReplaceParames(sql2,param)
//	if err = orm1.Exec(sql2).Error;err!=nil {
//		return err
//	}
//	return nil
//}

func (y *Recharge) RechargeList(param map[string]string) ( interface{}, error) {

	sql := `select r.*,c.name,c.phone from recharge r 
				left join customer c on c.customer_id = r.customer_id
				where 1=1 `
	if param["keyword"]!=""{
		sql += ` and (c.phone like '%` + param["keyword"] + `%' or c.name '%` + param["keyword"] + `%') `
	}
	param["total"] = GetTotalCount(sql)

	param["sort"] = "create_time"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)
	var res []Recharge
	if err := orm.Eloquent.Raw(sql).Scan(&res).Error;err != nil {
		return nil,err
	}
	return res,nil
}