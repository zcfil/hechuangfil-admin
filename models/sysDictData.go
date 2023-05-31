package models

import (
	"time"
	orm "hechuangfil-admin/database"
)

type DictData struct {
	//字典编码
	DictCode int64 `gorm:"column:dictCode;primary_key" json:"dictCode" example:"1"`

	//显示顺序
	DictSort int64 `gorm:"column:dictSort" json:"dictSort" example:"1"`

	//数据标签
	DictLabel string `gorm:"column:dictLabel" json:"dictLabel"`

	//数据键值
	DictValue string `gorm:"column:dictValue" json:"dictValue"`

	//字典类型
	DictType  string `gorm:"column:dictType" json:"dictType"`
	CssClass  string `gorm:"column:cssClass" json:"cssClass"`
	ListClass string `gorm:"column:listClass" json:"listClass"`
	IsDefault string `gorm:"column:isDefault" json:"isDefault"`

	//状态
	Status     string `gorm:"column:status" json:"status"`
	Default    string `gorm:"column:default" json:"default"`
	CreateBy   string `gorm:"column:create_by" json:"createBy"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UpdateBy   string `gorm:"column:update_by" json:"updateBy"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`

	//备注
	Remark    string `gorm:"column:remark" json:"remark"`
	Params    string `gorm:"column:params" json:"params"`
	//DataScope string `gorm:"column:dataScope" json:"dataScope"`
	IsDel     string `gorm:"column:is_del" json:"isDel"`
}

func (e *DictData) Create() (DictData, error) {
	var doc DictData
	e.CreateTime = time.Now()
	e.UpdateTime = time.Now()
	result := orm.Eloquent.Table("sys_dict_data").Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

func (e *DictData) GetByCode() (DictData, error) {
	var doc DictData

	table := orm.Eloquent.Table("sys_dict_data")
	if e.DictCode != 0 {
		table = table.Where("dictCode = ?", e.DictCode)
	}
	if e.DictLabel != "" {
		table = table.Where("dictLabel = ?", e.DictLabel)
	}
	if e.DictType != "" {
		table = table.Where("dictType = ?", e.DictType)
	}

	if err := table.Where("is_del = 0").First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (e *DictData) Get() ([]DictData, error) {
	var doc []DictData

	table := orm.Eloquent.Table("sys_dict_data")
	if e.DictCode != 0 {
		table = table.Where("dictCode = ?", e.DictCode)
	}
	if e.DictLabel != "" {
		table = table.Where("dictLabel = ?", e.DictLabel)
	}
	if e.DictType != "" {
		table = table.Where("dictType = ?", e.DictType)
	}

	if err := table.Where("is_del = 0").Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (e *DictData) GetPage(pageSize int, pageIndex int) ([]DictData, int32, error) {
	var doc []DictData

	table := orm.Eloquent.Select("*").Table("sys_dict_data")
	if e.DictCode != 0 {
		table = table.Where("dictCode = ?", e.DictCode)
	}
	if e.DictType != "" {
		table = table.Where("dictType = ?", e.DictType)
	}
	if e.DictLabel != "" {
		table = table.Where("dictLabel = ?", e.DictLabel)
	}
	if e.Status != "" {
		table = table.Where("status = ?", e.Status)
	}

	// 数据权限控制
	//dataPermission := new(DataPermission)
	//dataPermission.UserId, _ = utils.StringToInt64(e.DataScope)
	//table = dataPermission.GetDataScope("sys_dict_data", table)

	var count int32

	if err := table.Where("is_del = 0").Order("dictSort").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("is_del = 0").Count(&count)
	return doc, count, nil
}

func (e *DictData) Update(id int64) (update DictData, err error) {
	if err = orm.Eloquent.Table("sys_dict_data").Where("dictCode = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	update.UpdateTime = time.Now()
	e.UpdateTime = time.Now()
	if err = orm.Eloquent.Table("sys_dict_data").Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

func (e *DictData) Delete(id int64) (success bool, err error) {
	var mp = map[string]string{}
	mp["is_del"] = "1"
	mp["update_time"] = time.Now().Format("2006/01/02 15:04:05")
	mp["update_by"] = e.UpdateBy
	if err = orm.Eloquent.Table("sys_dict_data").Where("dictCode = ?", id).Update(mp).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}
