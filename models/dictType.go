package models

import (
	"time"
	orm "hechuangfil-admin/database"
)

type DictType struct {
	DictId int64 `gorm:"column:dictId;primary_key" json:"dictId" example:"1"`
	//字典名称
	DictName string `gorm:"column:dictName" json:"dictName"`
	//字典类型
	DictType string `gorm:"column:dictType" json:"dictType"`
	//状态
	Status string `gorm:"column:status" json:"status"`

	//DataScope string `gorm:"column:dataScope" json:"dataScope"`

	Params string `gorm:"column:params" json:"params"`
	//创建者
	CreateBy string `gorm:"column:create_by" json:"createBy"`
	//创建时间
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	//更新者
	UpdateBy string `gorm:"column:update_by" json:"updateBy"`
	//更新时间
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	//备注
	Remark string `gorm:"column:remark" json:"remark"`
	IsDel  string `gorm:"column:is_del" json:"isDel"`
}

func (e *DictType) Create() (DictType, error) {
	var doc DictType
	e.CreateTime = time.Now()
	e.UpdateTime = time.Now()
	result := orm.Eloquent.Table("sys_dict_type").Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

func (e *DictType) Get() (DictType, error) {
	var doc DictType

	table := orm.Eloquent.Table("sys_dict_type")
	if e.DictId != 0 {
		table = table.Where("dictId = ?", e.DictId)
	}
	if e.DictName != "" {
		table = table.Where("dictName = ?", e.DictName)
	}
	if e.DictType != "" {
		table = table.Where("dictType = ?", e.DictType)
	}

	if err := table.Where("is_del = 0").First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (e *DictType) GetPage(pageSize int, pageIndex int) ([]DictType, int32, error) {
	var doc []DictType

	table := orm.Eloquent.Select("*").Table("sys_dict_type")
	if e.DictId != 0 {
		table = table.Where("dictId = ?", e.DictId)
	}
	if e.DictName != "" {
		table = table.Where("dictName = ?", e.DictName)
	}

	// 数据权限控制
	//dataPermission := new(DataPermission)
	//dataPermission.UserId, _ = utils.StringToInt64(e.DataScope)
	//table = dataPermission.GetDataScope("sys_dict_type", table)

	var count int32

	if err := table.Where("is_del = 0").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("is_del = 0").Count(&count)
	return doc, count, nil
}

func (e *DictType) Update(id int64) (update DictType, err error) {
	e.UpdateTime = time.Now()
	if err = orm.Eloquent.Table("sys_dict_type").First(&update, id).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	update.UpdateTime = time.Now()
	if err = orm.Eloquent.Table("sys_dict_type").Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

func (e *DictType) Delete(id int64) (success bool, err error) {
	var mp = map[string]string{}
	mp["is_del"] = "1"
	mp["update_time"] = time.Now().Format("2006/01/02 15:04:05")
	mp["update_by"] = e.UpdateBy
	if err = orm.Eloquent.Table("sys_dict_type").Where("dictId = ?", id).Update(mp).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}
