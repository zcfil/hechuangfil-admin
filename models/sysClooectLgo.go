package models

import (
	"errors"
	orm "hechuangfil-admin/database"
)

/**
 * @Author: king bo
 * @Author: kingbo@163.com
 * @Date: 2020/5/20 8:09 下午
 * @Desc: 钱包归集记录
 */

type CollectLog struct {
	ID         int64     `json:"id"`
	Symbol     string    `json:"symbol"`      //币种简称
	Quantity   float64   `json:"quantity"`    //归集数量
	Addres    string    `json:"addres"`     //归集地址
	Txid       string    `json:"txid"`        //交易hash
	CreateTime int64 `json:"create_time"` //归集时间
	Remark     string    `json:"remark"`      //备注
}

//添加新币种
func (c *CollectLog) Insert() (id int64, err error) {
	// 检查币种是否存在
	var count int
	orm.Eloquent.Table("sys_collect_log").Where("txid = ?", c.Txid).Count(&count)
	if count > 0 {
		err = errors.New("交易已存在！")
		return
	}
	//添加数据
	if err = orm.Eloquent.Table("sys_collect_log").Create(&c).Error; err != nil {
		return
	}
	id = c.ID
	return
}
