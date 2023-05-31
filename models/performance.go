package models

import (
	"time"
	orm "hechuangfil-admin/database"
	"hechuangfil-admin/utils"
)

type Performance struct {
	ID           string     `gorm:"column:id" json:"id"`
	Userid           string     `gorm:"column:userid" json:"userid"`
	Amount    	 float64    `gorm:"column:amount" json:"amount"`               //金额
	Customerid     	string    `gorm:"column:customerid" json:"customerid"`               //客户id
	Remark        string    `gorm:"column:remark" json:"remark"`            //备注
	Status 		int 		`gorm:"column:status" json:"status"`
	Customername        string    `gorm:"column:customername" json:"customername"`
	NickName        string    `gorm:"column:nick_name" json:"nick_name"`
	Profits float64    `gorm:"column:profits" json:"profits"`
	Percent float64    `gorm:"column:percent" json:"percent"`
	CreateTime	time.Time	`gorm:"column:create_time" json:"create_time"`	//创建时间
}

func (p *Performance)IndividualPerformance(param map[string]string)(result interface{},err error){
	//sql := `select i.id,c.userid,amount,i.customerid,i.remark,i.status,c.name customername,u.nick_name,p.percent, round(p.percent * i.amount,2) profits,i.create_time
	//	from investment i
	//	left join customer c on i.customerid = c.id
	//	left join sys_user u on c.userid = u.user_id
	//	left join (		select sum(percent)percent,a.userid from (
	//				select c.amount,ifnull(p.userid,0)pid,c.userid from (select sum(i1.amount)amount,c1.userid from customer c1
	//						left join investment i1 on i1.customerid = c1.id
	//						where c1.is_del = 0 and i1.is_del = 0 and i1.create_time <= concat(LAST_DAY(:date),' 23:59:59') and
	//										date_add(:date,interval -day(:date)+1 day) <= i1.create_time
	//						group by c1.userid
	//						)c
	//						left join (select userid from profitconfig WHERE profittype=1 GROUP BY userid) p on p.userid = c.userid
	//		)a
	//		left join profitconfig p1 on a.pid = p1.userid
	//		where p1.profitlevel < a.amount  and p1.profittype = if(ifnull(p1.userid,0)=0,2,1)
	//		GROUP BY a.userid) p on p.userid = c.userid
	//	where i.create_time <= concat(LAST_DAY(:date),' 23:59:59') and
	//										date_add(:date,interval -day(:date)+1 day) <= i.create_time  `
	con := ""
	if param["userid"] !=""{
		con += " and a.userid=:userid"
	}
	sql := `select amount,a.userid,a.id,sum(p1.levelgain)+(amount-max(p1.profitlevel))*sum(percent) profits,customername,u.username,u.nick_name,a.create_time,a.remark,a.status from (
					select c.amount,ifnull(p.userid,0)pid,c.userid,c.id, customername,create_time,remark,status from (
							select i1.amount,c1.userid,i1.id,c1.name customername,i1.create_time,i1.remark,i1.status from customer c1
								left join investment i1 on i1.customerid = c1.id
								where c1.is_del = 0 and i1.is_del = 0 and i1.create_time <= concat(LAST_DAY(:date),' 23:59:59') and
											date_add(:date,interval -day(:date)+1 day) <= i1.create_time
							)c 
							left join (select userid from profitconfig WHERE profittype=1 GROUP BY userid) p on p.userid = c.userid

				)a 
				left join profitconfig p1 on a.pid = p1.userid
				left join sys_user u on a.userid = u.user_id
				where p1.profitlevel < a.amount and p1.profittype = if(ifnull(p1.userid,0)=0,2,1) `+con+`
				GROUP BY a.id`

	sql = utils.SqlReplaceParames(sql,param)
	//总数
	param["total"] = GetTotalCount(sql)
	//分页 and 排序
	param["sort"] = "a.create_time"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)
	var res []Performance
	err = orm.Eloquent.Raw(sql).Scan(&res).Error
	return res,err
}