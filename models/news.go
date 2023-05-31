package models

import (
	"errors"
	"fmt"
	orm "hechuangfil-admin/database"
	"strconv"
	"time"
)

type News struct {
	ID         int32     `gorm:"id" json:"id"`
	Title      string    `gorm:"title" json:"title"`
	Content    string    `gorm:"content" json:"content"`
	Status     string    `gorm:"status" json:"status"` // 1 可见  2不可见
	UpdateTime time.Time `gorm:"update_time" json:"update_time"`
	CreateTime time.Time `gorm:"create_time" json:"create_time"`
}

func NewNews() *News {
	news := new(News)
	return news
}

func (this *News) GetAll(param map[string]string) (ret interface{}, err error) {
	sql := `select * from news`

	pageIndex, err1 := strconv.ParseInt(param["pageIndex"], 10, 64)
	if err1 != nil {
		err = err1
		return
	}
	pageSize, err1 := strconv.ParseInt(param["pageSize"], 10, 64)
	if err1 != nil {
		err = err1
		return
	}
	param["total"] = GetTotalCount(sql)
	start := (pageIndex - 1) * pageSize
	sql += ` limit ` + strconv.FormatInt(start, 10) + `, ` + param["pageSize"]

	findList := make([]News, 0)
	err = orm.Eloquent.Raw(sql).Scan(&findList).Error
	if err != nil {
		return
	}

	ret = findList
	return
}

func (this *News) AddNews(title, content string, status string) (err error) {
	sql := `select id from news where title = "%s"`
	sql = fmt.Sprintf(sql, title)
	if GetTotalCount1(sql) > 0 {
		return errors.New("已有相同的标题")
	}

	sql1 := `insert into news (title, content, status, update_time, create_time) values ("` + title + `", "` + content + `",` +
		status + `, now(), now())`
	return orm.Eloquent.Exec(sql1).Error
}

func (this *News) Update(id string, title, content string, status string) (err error) {
	sql := `update news set title = "%s", content = "%s", status = %s where id = %s`
	sql = fmt.Sprintf(sql, title, content, status, id)
	return orm.Eloquent.Exec(sql).Error
}

func (this *News) Del(id string) (err error) {
	sql := `delete from news where id = ` + id
	return orm.Eloquent.Exec(sql).Error
}

func (this *News) UpdateStatus(id, status string) (err error) {
	sql := `update news set status = %s where id = %s`
	sql = fmt.Sprintf(sql, status, id)
	return orm.Eloquent.Exec(sql).Error
}
