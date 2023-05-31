package models

import (
	"time"
	orm "hechuangfil-admin/database"
)

type Post struct {
	//岗位编号
	PostId int64 `gorm:"column:postId;primary_key" json:"postId" example:"1" extensions:"x-description=标示"`

	//岗位名称
	PostName string `gorm:"column:postName" json:"postName"`

	//岗位代码
	PostCode string `gorm:"column:postCode" json:"postCode"`

	//岗位排序
	Sort int `gorm:"column:sort" json:"sort"`

	//状态
	Status string `gorm:"column:status" json:"status"`

	//描述
	Remark string `gorm:"column:remark" json:"remark"`

	//创建时间
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`

	//最后修改时间
	UpdateTime time.Time  `gorm:"column:update_time" json:"updateTime"`

	//是否删除
	IsDel int64 `gorm:"column:is_del" json:"isDel"`

	CreateBy string `gorm:"column:create_by" json:"createBy"`

	UpdateBy string `gorm:"column:update_by" json:"updateBy"`

	//DataScope string `gorm:"column:dataScope" json:"dataScope"`

	Params string `gorm:"column:params" json:"params"`
}

func (e *Post) Create() (Post, error) {
	var doc Post
	e.CreateTime = time.Now()
	e.UpdateTime = time.Now()
	result := orm.Eloquent.Table("sys_post").Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

func (e *Post) Get() (Post, error) {
	var doc Post

	table := orm.Eloquent.Table("sys_post")
	if e.PostId != 0 {
		table = table.Where("postId = ?", e.PostId)
	}
	if e.PostName != "" {
		table = table.Where("postName = ?", e.PostName)
	}
	if e.PostCode != "" {
		table = table.Where("postCode = ?", e.PostCode)
	}

	if err := table.Where("is_del = 0").First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (e *Post) GetList() ([]Post, error) {
	var doc []Post

	table := orm.Eloquent.Table("sys_post")
	if e.PostId != 0 {
		table = table.Where("postId = ?", e.PostId)
	}
	if e.PostName != "" {
		table = table.Where("postName = ?", e.PostName)
	}
	if e.PostCode != "" {
		table = table.Where("postCode = ?", e.PostCode)
	}

	if err := table.Where("is_del = 0").Find(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (e *Post) GetPage(pageSize int, pageIndex int) ([]Post, int32, error) {
	var doc []Post

	table := orm.Eloquent.Select("*").Table("sys_post")
	if e.PostId != 0 {
		table = table.Where("postId = ?", e.PostId)
	}
	if e.PostName != "" {
		table = table.Where("postName = ?", e.PostName)
	}

	// 数据权限控制
	//dataPermission := new(DataPermission)
	//dataPermission.UserId, _ = utils.StringToInt64(e.DataScope)
	//table = dataPermission.GetDataScope("sys_post", table)

	var count int32

	if err := table.Where("is_del = 0").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("is_del = 0").Count(&count)
	return doc, count, nil
}

func (e *Post) Update(id int64) (update Post, err error) {
	e.UpdateTime = time.Now()
	if err = orm.Eloquent.Table("sys_post").First(&update, id).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table("sys_post").Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}

func (e *Post) Delete(id int64) (success bool, err error) {
	var mp = map[string]string{}
	mp["is_del"] = "1"
	mp["update_time"] = time.Now().Format("2006/01/02 15:04:05")
	mp["update_by"] = e.UpdateBy
	if err = orm.Eloquent.Table("sys_post").Where("postId = ?", id).Update(mp).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}
