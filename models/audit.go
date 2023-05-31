package models

import (
	"context"
	"errors"
	"fmt"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/lotus/chain/types"
	orm "hechuangfil-admin/database"
	"hechuangfil-admin/pkg/lotus"
	"hechuangfil-admin/utils"
	"sync"
	"time"
)

type Withdraw struct {
	WithdrawId           string     `json:"withdraw_id"`
	ToAddres    	 string    		`json:"to_addres"`               //提现id
	Amount        string    		`json:"Amount"`            //金额
	CreateTime	time.Time			`json:"create_time"`	//创建时间
	UpdateTime	time.Time			`json:"update_time"`	//审核时间
	Status		int				`json:"status"`	//状态 0待审核，1已通过，2已拒绝，3已到账
	Cid    	 string    			`json:"cid"`	//消息ID
	Height		string			`json:"height"`	//高度
	Name    	 string    			`json:"name"`	//用户姓名
	Phone		string			`json:"phone"`	//手机号
	CustomerId string 			`json:"customer_id"`
}

func (e *Withdraw) AuditList(param map[string]string) (result interface{}, err error) {
	//拼凑筛选条件sql
	sql := ` select w.*,c.name,c.phone from withdraw w 
				left join customer c on w.customer_id = c.customer_id
				where w.status = 0 `

	keyword := param["keyword"]
	if keyword != "" {
		sql += ` and (c.phone like '%` + keyword + `%' ) `
	}
	//总数
	param["total"] = GetTotalCount(sql)
	//分页 and 排序
	param["sort"] = "w.create_time"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)

	user := make([]Withdraw, 0)
	orm.Eloquent.Raw(sql).Scan(&user)

	result = user

	return
}
func (e *Withdraw) AuditLogList(param map[string]string) (result interface{}, err error) {
	//拼凑筛选条件sql
	sql := ` select w.*,c.name,c.phone from withdraw w 
				left join customer c on w.customer_id = c.customer_id
				where w.status <> 0 `

	keyword := param["keyword"]
	if keyword != "" {
		sql += ` and (c.phone like '%` + keyword + `%' or c.name like '%` + keyword + `%') `
	}
	//总数
	param["total"] = GetTotalCount(sql)
	//分页 and 排序
	param["sort"] = "w.create_time"
	param["order"] = "desc"
	sql += utils.LimitAndOrderBy(param)

	user := make([]Withdraw, 0)
	orm.Eloquent.Raw(sql).Scan(&user)

	result = user
	return
}

func (e *Withdraw) WithdrawById(id string,status string) (result Withdraw, err error) {
	//拼凑筛选条件sql
	sql := ` select w.*,c.name,c.phone from withdraw w 
				left join customer c on w.customer_id = c.customer_id
				where w.withdraw_id = '`+id+`'`
	if status !=""{
		sql += ` and w.status = `+status
	}
	var w Withdraw
	err = orm.Eloquent.Raw(sql).Scan(&w).Error

	result = w
	return
}

var audit sync.Mutex
func (e *Withdraw) CheckAudit(param map[string]string) (err error) {
	orm1 := orm.Eloquent.Begin()
	defer func() {
		audit.Unlock()
		if err!=nil{
			orm1.Rollback()
			return
		}
		orm1.Commit()
	}()
	audit.Lock()
	w,err := e.WithdrawById(param["withdraw_id"],VERIFICATION_STR)
	if err!=nil{
		return err
	}

	if w == (Withdraw{}){
		return fmt.Errorf("该记录不存在或已经审核")
	}
	param["amount"] = w.Amount
	param["customer_id"] = w.CustomerId
	//拒绝
	if param["status"] == NO_PASS_VERIFICATION_STR{
		sql1 := ` update withdraw set status = :status ,update_time = now()  where withdraw_id = :withdraw_id `
		sql1 = utils.SqlReplaceParames(sql1,param)
		err = orm1.Exec(sql1).Error

		sql2 := ` update customer set frozen_capital=frozen_capital-:amount,balance=balance+:amount  where customer_id = :customer_id `
		sql2 = utils.SqlReplaceParames(sql2,param)
		err = orm1.Exec(sql2).Error
		return
	}


	addrTo, err := address.NewFromString(w.ToAddres)
	if err!=nil{
		return errors.New("钱包格式异常："+err.Error())
	}
	//val, err := types.ParseFIL(strconv.FormatFloat(fil,'f',-1,64))
	//if err != nil {
	//	return nil,fmt.Errorf("failed to parse amount: %w", err)
	//}
	var cf Config
	cf ,err = cf.GetConfig(FROM_ADDRESS)
	if err!=nil{
		return errors.New("未配置转出钱包!")
	}
	from, err := address.NewFromString(cf.Value)
	if err!=nil{
		return errors.New("配置钱包格式不正确:"+err.Error())
	}
	//fmt.Println(w.Amount,utils.NanoOrAttoToFILstr(w.Amount,utils.AttoFIL))
	val, err := types.ParseFIL(w.Amount)
	if err != nil {
		return fmt.Errorf("failed to parse amount: %w", err)
	}
	sm, err := lotus.FullAPI.MpoolPushMessage(context.Background(), &types.Message{
		To:         addrTo,
		From:       from,
		Value:      types.BigInt(val),
	},nil)
	if err != nil {
		return err
	}
	CidWithdraw[sm.Cid().String()] = true
	param["cid"] = sm.Cid().String()
	sql := ` update withdraw set status = 1 ,update_time = now(),cid = :cid where withdraw_id = :withdraw_id `
	sql = utils.SqlReplaceParames(sql,param)
	if err = orm1.Exec(sql).Error;err!=nil{
		return err
	}

	sql2 := ` update customer set frozen_capital=frozen_capital-:amount  where customer_id = :withdraw_id `
	sql2 = utils.SqlReplaceParames(sql2,param)
	if err = orm1.Exec(sql2).Error;err!=nil{
		return err
	}

	return
}

func (e *Withdraw) UpdateByCid(cid string,height string) (err error) {
	//拼凑筛选条件sql
	sql := ` update withdraw set status = 3,height = ? where cid = ? `

	err = orm.Eloquent.Exec(sql,height,cid).Error
	return
}