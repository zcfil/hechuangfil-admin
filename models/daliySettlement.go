package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/robfig/cron"
	"hechuangfil-admin/common"
	orm "hechuangfil-admin/database"
	"hechuangfil-admin/define"
	"strconv"
	"time"

	log "hechuangfil-admin/logrus"
)

// AverageFilConfig 利润计算
type AverageFilConfig struct {
	ID         int64     `gorm:"column:id"`
	Height     int64     `gorm:"column:height"`
	CreateTime time.Time `gorm:"column:create_time"`
	AverageFil float64   `gorm:"column:average_fil"`
}

// LockBalanceRecord 锁定的历史记录
type LockBalanceRecord struct {
	Arr   []float64				`json:"arr"`
}

type DailySettlement struct {

	// 结算时使用数据
	mapRatio map[int8]int32
	ratioSum int32

	mapCustomerRatio map[int8]int32
	customerRatioSum int32

	aver float64
}

func NewSettlement() *DailySettlement {
	settle := new(DailySettlement)
	settle.startSettleLoop()
	return settle
}

func (this *DailySettlement) startSettleLoop() {
	spec := "0, 01, 00, *, *, *" // 每天0000 点
	c := cron.New()
	if err := c.AddFunc(spec, this.SettleFunc); err != nil {
		log.Error("Add settle func error:", err.Error())
		return
	}
	c.Start()
}

func (this *DailySettlement) SettleFunc() {
	//if !this.checkSettleDate() {
	//	log.Info("当天已经结算过了")
	//	return
	//}
	if err := this.settle(); err != nil {
		log.Error("结算错误, err:", err.Error())
	}
}

// 验证一下结算时间 一天只能结算一次
func (this *DailySettlement) checkSettleDate() bool {
	type findDate struct {
		Date time.Time `gorm:"column:date"`
	}
	sql := `select date from settle_date where id = (select max(id) from settle_date)`

	find := &findDate{}
	if err := orm.Eloquent.Raw(sql).Scan(find).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return true
		}
		log.Error("查找最近结算时间错误 err:", err.Error())
		return false
	}

	lastDay := common.TimeToDay(find.Date)
	curDay := common.TimeToDay(time.Now())
	if lastDay == curDay {
		return false
	}

	return true
}

func (this *DailySettlement) settle() error {
	// 先取出结算比例
	mapRatio, ratioSum, err := this.GetRatio()
	if err != nil {
		return err
	}
	this.mapRatio = mapRatio
	this.ratioSum = ratioSum
	log.Info("结算比例信息:", mapRatio)

	// 获取用户释放到余额和锁仓占比
	mapCustomerRatio, customerRatioSum, err := this.getCustomerBalanceLockRatio()
	if err != nil {
		return err
	}
	this.mapCustomerRatio = mapCustomerRatio
	this.customerRatioSum = customerRatioSum
	log.Info("用户结算余额锁仓占比:", mapCustomerRatio)

	// 获取fil获润比例
	average, err := this.getAverageFil()
	if err != nil {
		return err
	}
	log.Info("获润比例信息:", *average)

	this.aver = average.AverageFil

	// 插入总的结算记录 (用来判断每天结算一次)
	sql := `insert into settle_date (average_fil, filearnings_id, date) values (` +
		strconv.FormatFloat(average.AverageFil, 'f', 10, 64) + `,` + strconv.FormatInt(average.ID, 10) + `, now())`
	var id []int64
	if err := orm.Eloquent.Exec(sql).Raw("select LAST_INSERT_ID() as id").Pluck("id", &id).Error; err != nil {
		log.Error("插入结算记录数据错误， err:", err.Error())
		return err
	}

	lastID := id[0]
	log.Info("last id:", lastID)

	// 取出每个用户的订单总算力
	mapCustomerHashrate, err := this.selectUserOrderHashrate()
	if err != nil {
		return err
	}

	// 分别计算客户、分公司、公司收益	FIL nanoFIL 10^9 attoFIL 10^18
	for customerID, customerHashrate := range mapCustomerHashrate {
		if err := this.settleCustomer(customerID, customerHashrate, lastID); err != nil {
			log.Error("结算错误 顾客ID:", customerID, "  lastID:", lastID, "  errInfo:", err.Error())
			continue
		}
	}
	return nil
}

func (this *DailySettlement) settleCustomer(customerID int64, customerHashrate float64, lastID int64) (err error) {
	customerHashrate *= define.MININT_RATIO
	// 总收益
	totalIncome := float64(customerHashrate) * this.aver

	// 客户收益
	customerIncome := (float64(this.mapRatio[define.PROFIT_CUSTOMER_RATIO]) / float64(this.ratioSum)) * totalIncome
	// 分公司收益
	filialeIncome := (float64(this.mapRatio[define.PROFIT_FILIATE_RATIO])) / float64(this.ratioSum) * totalIncome
	// 公司收益
	company := (float64(this.mapRatio[define.PROFIT_COMPANY_RATIO])) / float64(this.ratioSum) * totalIncome

	// 客户释放到余额
	customerBalance := (float64(this.mapCustomerRatio[define.PROFIT_CUSTOMER_BALANCE_RATIO]) / float64(this.customerRatioSum)) * customerIncome
	// 客户分180天释放部分
	customerLock := (float64(this.mapCustomerRatio[define.PROFIT_CUSTOMER_LOCK_RATIO]) / float64(this.customerRatioSum)) * customerIncome

	// 保存到数据库
	session := orm.Eloquent.Begin()
	defer func() {
		if err != nil {
			session.Rollback()
			return
		}
		session.Commit()
	}()

	customerIDStr := strconv.FormatInt(customerID, 10)

	// 获取旧的用户锁定金额信息
	oldLockBalance, records, isExist, err1 := this.getLockBalanceMsg(customerIDStr, session)
	if err1 != nil {
		err = err1
		log.Error("获取锁仓数据错误， err:", err.Error())
		return
	}

	addLock := oldLockBalance / define.LOCK_BALANCE_DAY_COUNT
	balanceAdd := addLock + customerBalance

	newLockBalance := oldLockBalance + customerLock
	// 先减掉到期的
	if len(records) >= define.LOCK_BALANCE_DAY_COUNT {
		newLockBalance -= records[0]
		records = records[1:]
	}
	// 添加新的锁定部分
	customerLock,_ = strconv.ParseFloat(fmt.Sprintf("%.6f", customerLock), 64)   // 保留6位小数
	records = append(records, customerLock)
	if err = this.saveNewLockBalanceMsg(customerIDStr, session, newLockBalance, records, isExist); err != nil {
		log.Error("保存新的锁仓数据错误, err：", err.Error())
		return
	}

	// 添加用户数据
	sql1 := `update customer set balance = ifnull(balance,0)+` +
		strconv.FormatFloat(balanceAdd, 'f', 10, 64) +
		` where customer_id = ` + customerIDStr
	if err = session.Exec(sql1).Error; err != nil {
		log.Error("设置用户的余额和锁定余额错误, err:", err, "顾客ID:", customerID)
		return
	}

	// 添加公司收益数据
	companyID := define.COMPANY_INCOME_ID
	log.Info("用户ID:", customerID, " 增加公司收益:", company, "  增加分公司收益:", filialeIncome)
	sql2 := `update company_income set income = ifnull(income, 0) + ` + strconv.FormatFloat(company, 'f', 10, 64) +
		`, filiale_income = ifnull(filiale_income, 0) + ` + strconv.FormatFloat(filialeIncome, 'f', 10, 64) +
		` where id = ` + strconv.FormatInt(int64(companyID), 10)
	if err = session.Exec(sql2).Error; err != nil {
		log.Error("更新公司收益错误, err:", err, "顾客ID:", customerID)
		return
	}

	// 插入新的结算数据
	sql3 := `insert into settle_log (
						customer_id, 
						total_income, 
						customer_income, 
						company_income, 
						filiale_income, 
						to_customer_balance, 
						to_customer_lock, 
						customer_lock_release, 
						settle_date_id, time) values (
						%s, %s, %s, %s, %s, %s, %s, %s, %d, now())`
	sql3 = fmt.Sprintf(sql3,
		customerIDStr,
		strconv.FormatFloat(totalIncome, 'f', 10, 64),
		strconv.FormatFloat(customerIncome, 'f', 10, 64),
		strconv.FormatFloat(company, 'f', 10, 64),
		strconv.FormatFloat(filialeIncome, 'f', 10, 64),
		strconv.FormatFloat(customerBalance, 'f', 10, 64),
		strconv.FormatFloat(customerLock, 'f', 10, 64),
		strconv.FormatFloat(addLock, 'f', 10, 64), lastID)

	if err = session.Exec(sql3).Error; err != nil {
		log.Error("插入结算数据错误, err:", err.Error(), "顾客ID:", customerID)
		return
	}
	return
}

func (this *DailySettlement) saveNewLockBalanceMsg(customerID string, session *gorm.DB, lockBalance float64, record []float64, isExist bool) (err error) {
	lr := &LockBalanceRecord{Arr: record}
	bin, err1 := json.Marshal(lr)
	if err1 != nil {
		err = err1
		return
	}

	recordStr := string(bin)
	lockBalanceStr := strconv.FormatFloat(lockBalance, 'f', 10, 64)
	if isExist {
		sql := `update lock_balance set lock_balance = %s, record = '%s' where customer_id = %s`
		sql = fmt.Sprintf(sql, lockBalanceStr, recordStr, customerID)
		err = session.Exec(sql).Error
		if err != nil {
			return
		}
		return
	}

	sql := `insert into lock_balance (customer_id, lock_balance, record) values (%s, %s, '%s')`
	sql = fmt.Sprintf(sql, customerID, lockBalanceStr,recordStr)
	err = session.Exec(sql).Error
	if err != nil {
		return
	}
	return
}

// getLockBalanceMsg 取出分批释放的数据
func (this *DailySettlement) getLockBalanceMsg(customerID string, session *gorm.DB) (lockBalance float64, record []float64, isExist bool, err error) {
	isExist = true
	type findLockBalance struct {
		LockBalance float64 `gorm:"column:lock_balance"`
		Record 		string  `gorm:"column:record"`
	}
	flb := &findLockBalance{}
	record = make([]float64, 0)
	sql := `select * from lock_balance where customer_id = ` + customerID
	err = session.Raw(sql).Scan(flb).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
		isExist = false
		return
	}
	if err != nil {
		log.Error("查询数据错误:", err.Error(), "顾客ID:", customerID)
		return
	}

	lockBalance = flb.LockBalance
	if len(flb.Record) <= 0 {
		return
	}
	var lockRecord LockBalanceRecord
	err = json.Unmarshal([]byte(flb.Record), &lockRecord)
	if err != nil {
		return
	}
	record = lockRecord.Arr
	return
}

// selectUserOrderHashrate 选出每个用户订单总算力
func (this *DailySettlement) selectUserOrderHashrate() (map[int64]float64, error) {
	type CustomerAmount struct {
		CustomerID    int64   `gorm:"column:customer_id"`
		CustomerTotal float64 `gorm:"column:hashrate_total"`
	}
	sql := `select customer_id, sum(hashrate) as hashrate_total from orders where date_add(create_time, interval 1 day) < now() group by customer_id`  // 筛选订单要结算时间之前一天的订单
	//sql := `select customer_id, sum(hashrate) as hashrate_total from orders group by customer_id`			// 测试专用（结算订单没有做日期筛选）
	findList := make([]CustomerAmount, 0)
	if err := orm.Eloquent.Raw(sql).Scan(&findList).Error; err != nil {
		return nil, err
	}

	retData := make(map[int64]float64)
	for _, data := range findList {
		retData[data.CustomerID] = data.CustomerTotal
	}
	return retData, nil
}

// GetRatio 获取结算比例
func (this *DailySettlement) GetRatio() (mapRet map[int8]int32, ratioSum int32, err error) {
	type ratio struct {
		ID    int8  `gorm:"column:id"`
		Ratio int32 `gorm:"column:ratio"`
	}

	findList := make([]ratio, 0)
	sql := `select id, ratio from profit_config where id in (%d, %d, %d)`
	sql = fmt.Sprintf(sql, define.PROFIT_CUSTOMER_RATIO, define.PROFIT_FILIATE_RATIO, define.PROFIT_COMPANY_RATIO)
	if err = orm.Eloquent.Raw(sql).Scan(&findList).Error; err != nil {
		return
	}

	if len(findList) != 3 {
		err = errors.New("获取到的分润配置长度错误")
		return
	}

	ratioSum = 0
	mapRet = make(map[int8]int32)
	for _, data := range findList {
		mapRet[data.ID] = data.Ratio
		ratioSum += data.Ratio
	}
	return
}

// 获取收益比， fixme 暂时获取最新的数据
func (this *DailySettlement) getAverageFil() (*AverageFilConfig, error) {
	find := &AverageFilConfig{}
	sql := `select * from filearnings where id=(select max(id) from filearnings)`
	//sql := `select * from filearnings where id=562`

	if err := orm.Eloquent.Raw(sql).Scan(find).Error; err != nil {
		log.Error("查找分润比例错误, err:", err.Error())
		return nil, err
	}
	return find, nil
}

// 获取客户余额锁仓占比
func (this *DailySettlement) getCustomerBalanceLockRatio() (mapRet map[int8]int32, de int32, err error) {
	type findRatio struct {
		ID    int8  `gorm:"column:id"`
		Ratio int32 `gorm:"column:ratio"`
	}
	sql := `select id, ratio from profit_config where id in (%d, %d)`
	sql = fmt.Sprintf(sql, define.PROFIT_CUSTOMER_BALANCE_RATIO, define.PROFIT_CUSTOMER_LOCK_RATIO)

	findList := make([]findRatio, 0)
	err = orm.Eloquent.Raw(sql).Scan(&findList).Error
	if err != nil {
		log.Error("获取结算记录失败", err.Error())
		return
	}

	if len(findList) != 2 {
		err = errors.New("获取到的分润配置长度错误")
		return
	}

	de = 0
	mapRet = make(map[int8]int32)
	for _, msg := range findList {
		mapRet[msg.ID] = msg.Ratio
		de += msg.Ratio
	}
	return
}
