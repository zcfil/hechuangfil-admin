package models

import (
	"sort"
	"strconv"
	orm "hechuangfil-admin/database"
)

type SettleCustomerHistory struct {
	ID  			int64  		`gorm:"column:id" json:"id"`
	Name 			string 		`gorm:"column:name" json:"name"`					// 名称
	Bank  			string 		`gorm:"column:bank" json:"bank"`					// 开户行
	BankNum  		string 		`gorm:"column:banknum" json:"banknum"`				// 卡号
	UserID 			int64 		`gorm:"column:user_id" json:"user_id"`				// 业务员ID
	NickName 		string 		`gorm:"column:nick_name" json:"nickname"`			// 业务员昵称
	SettleTime 		string 		`gorm:"column:settle_time" json:"settle_time"`		// 结算时间
	InvestID 		int64 		`gorm:"column:invest_id" json:"invest_id"`			// investment表ID
	CustomerID      int64		`gorm:"column:customer_id" json:"customer_id_id"`	// customer表ID
}


type sortHistory []SettleCustomerHistory
func (s sortHistory) Len() int {return len(s)}
func (s sortHistory) Less(i,j int) bool {return s[i].ID > s[j].ID}
func (s sortHistory) Swap(i, j int) {s[i], s[j] = s[j], s[i]}

func (this *SettleCustomerHistory) GetHistory(pageSize, pageIndex string) (dataList interface{}, total int, err error) {
	pSize, err1 := strconv.ParseInt(pageSize, 10, 64)
	if err1 != nil {
		err = err1
		return
	}

	pIndex, err1 := strconv.ParseInt(pageIndex, 10, 64)
	if err1 != nil {
		err = err1
		return
	}

	sql2 := `select id from customer_settle`
	total = GetTotalCount1(sql2)

	start := int64(total) - pSize * pIndex
	if start < 0 {
		pSize = pSize + start
		start = 0
	}

	if pSize < 0 {
		pSize = 0
	}

	sql := `select * from customer_settle limit ` + strconv.FormatInt(start, 10) + `,` + strconv.FormatInt(pSize, 10)
	retList := make([]SettleCustomerHistory, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&retList).Error; err != nil {
		return
	}

	sort.Sort(sortHistory(retList))
	dataList = retList
	return
}