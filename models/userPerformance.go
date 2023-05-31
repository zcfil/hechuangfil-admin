package models

import (
	"errors"
	"strconv"
	"time"
	orm "hechuangfil-admin/database"
)

type UserPerformance struct {
	userID int64
	NickName string					`json:"nick_name"`			 // 用户昵称
	TotalPerformance  float64		`json:"total_performance"`	 // 总业绩
	TodayPerformance  float64		`json:"today_performance"`	 // 当日业绩
	VipLevel 		  string		`json:"vip_level"`			 // vip等级
	VipTitle   		  string 		`json:"vip_title"`			 // vip称号
}


func NewUserPerformance(userID int64) *UserPerformance {
	return &UserPerformance{
		userID: userID,
	}
}

func(this *UserPerformance) GetTotal() (err error) {
	sql := `select nick_name, ifnull(accumulative, 0)+ ifnull(diff_value, 0) total_performance from sys_user where user_id=` + strconv.FormatInt(this.userID, 10)
	type totalPerformance struct {
		TotalPerformance  float64		`gorm:"column:total_performance"`
		NickName 		  string `gorm:"column:nick_name"`
	}

	findList := make([]totalPerformance, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&findList).Error; err != nil {
		return
	}
	if len(findList) != 1 {
		err = errors.New("符合查找条件的数量不对")
		return
	}
	this.TotalPerformance = findList[0].TotalPerformance
	this.NickName = findList[0].NickName
	return
}

// GetVipLevel 先获取总业绩再获取vip等级
func (this *UserPerformance) GetVipLevel() (err error) {
	sql := `select * from user_level`
	list := make([]UserLevelConfig, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&list).Error; err != nil {
		return
	}

	if len(list) <=0  {
		err = errors.New("符合查找条件的数量不对")
		return
	}

	// 根据业务员的业绩获取业务员的vip等级
	tempVipLevel := int64(0)
	allLen := len(list)
	for i, m := range list {
		vl := int64(i+1)
		if i+1 == allLen {
			tempVipLevel = vl
			break
		}
		next := list[i+1]
		if this.TotalPerformance >= m.LevelValue && this.TotalPerformance < next.LevelValue {
			tempVipLevel = vl
			break
		}
	}
	this.VipLevel = "V"+strconv.FormatInt(tempVipLevel, 10)
	this.setVipTitle(tempVipLevel)
	return
}

func (this *UserPerformance) GetToday() (err error) {
	referrals, err1 := this.getUserReferrals()
	if err1 != nil {
		err = err1
		return
	}

	l := len(referrals)
	sql := `select sum(i.amount) as total from investment as i where ( userid in (` + strconv.FormatInt(this.userID, 10)

	if l > 0 {
		sql += `,`
		sql += referrals
	}
	sql += `)`
	sql += `) and create_time >= "`
	today := time.Now().Format("2006-01-02") + ` 00:00:00"`
	sql += today

	type totalMsg struct {
		Total  float64   `gorm:"column:total" json:"total"`
	}
	finds := make([]totalMsg, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&finds).Error; err != nil {
		return
	}

	if len(finds) != 1 {
		err = errors.New("统计今日业绩数量错误")
		return
	}
	this.TodayPerformance = finds[0].Total
	return
}

// 获取推荐人列表
func (this *UserPerformance) getUserReferrals() (referrals string, err error) {
	// 先获取下级推荐人列表
	sql := `select referrals from referrer where userid = ` + strconv.FormatInt(this.userID, 10)
	type referral struct {
		Referrals  string   `gorm:"column:referrals" json:"referrals"`
	}
	find := make([]referral, 0)
	if err = orm.Eloquent.Raw(sql).Scan(&find).Error; err != nil {
		return
	}

	l := len(find)
	if l <= 0 {
		return
	}
	if l != 1 {
		err = errors.New("查找推荐列表数量错误")
		return
	}

	referrals = find[0].Referrals
	if len(referrals) == 0 {
		return
	}
	if referrals[0] == ',' {
		referrals = referrals[1:]
	}

	return
}

// setVipTitle 设置vip称号
func (this *UserPerformance) setVipTitle(vipLevel int64) {
	switch vipLevel {
	case 1:
		this.VipTitle = "业务员"
		return
	case 2:
		this.VipTitle = "铜冠"
		return
	case 3:
		this.VipTitle = "银冠"
		return
	case 4:
		this.VipTitle = "金冠"
		return
	default:
		this.VipTitle = "金冠"
		return
	}
}