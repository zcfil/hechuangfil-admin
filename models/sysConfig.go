package models

import (
	"errors"
	"github.com/tealeg/xlsx"
	"log"
	"strconv"
	time2 "time"
	orm "hechuangfil-admin/database"
	"hechuangfil-admin/pkg/export"
	"hechuangfil-admin/pkg/file"
)


type Config struct {
	ID    int64  `json:"id" gorm:"column:configId;primary_key"`     //编码
	Name  string `json:"name" gorm:"column:configKey;"` //参数名称 //参数键名ConfigKey string `json:"configKey" gorm:"column:configKey"`
	Title string `json:"title" gorm:"column:configName"`           //变量标题  //参数名称ConfigName string `json:"Name" gorm:"column:name;primary_key"`

	Remark        string `json:"remark" gorm:"column:remark"`               //变量描述 //Remark string `json:"remark" gorm:"column:remark"` //备注
	Group      string `json:"group" gorm:"column:configType"`           //变量分组
	Value      string `json:"value" gorm:"column:configValue"`           //参数变量值 	//参数键值 //ConfigValue string `json:"configValue" gorm:"column:configValue"`
	IsDel      string `json:"isDel" gorm:"column:is_del"`          //是否删除 0 正常使用，1 已删除
	CreateBy   string `json:"createBy" gorm:"column:create_by"`
	CreateTime time2.Time `json:"createTime" gorm:"column:create_time"`
	UpdateBy   string `json:"updateBy" gorm:"column:update_by"`
	UpdateTime time2.Time `json:"updateTime" gorm:"column:update_time"`
}

// Config 创建
func (e *Config) Create() (Config, error) {
	var doc Config
	doc.IsDel = "0"
	e.CreateTime = time2.Now()
	e.UpdateTime = time2.Now()
	result := orm.Eloquent.Table("sys_config").Create(&e)
	if result.Error != nil {
		err := result.Error
		return doc, err
	}
	doc = *e
	return doc, nil
}

// 获取 Config
func (e *Config) Get() (Config, error) {
	var doc Config

	table := orm.Eloquent.Table("sys_config")
	if e.ID != 0 {
		table = table.Where("configId = ?", e.ID)
	}

	if e.Name != "" {
		table = table.Where("configKey = ?", e.Name)
	}

	if err := table.Where("is_del = 0").First(&doc).Error; err != nil {
		return doc, err
	}
	return doc, nil
}

func (e *Config) GetPage(pageSize int, pageIndex int) ([]Config, int32, error) {
	var doc []Config

	table := orm.Eloquent.Select("*").Table("sys_config")

	if e.Name != "" {
		table = table.Where("configKey = ?", e.Name)
	}
	if e.Title != "" {
		table = table.Where("configName like ?", "%"+e.Title+"%")
	}
	if e.Group != "" {
		table = table.Where("sys_config.group = ?", e.Group)
	}

	// 数据权限控制
	//dataPermission := new(DataPermission)
	//dataPermission.UserId, _ = utils.StringToInt64(e.DataScope)
	//table = dataPermission.GetDataScope("sys_config", table)

	var count int32

	if err := table.Where("is_del = 0").Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&doc).Error; err != nil {
		return nil, 0, err
	}
	table.Where("is_del = 0").Count(&count)
	return doc, count, nil
}

func (e *Config) Update(id int64) (update Config, err error) {
	e.UpdateTime = time2.Now()
	if err = orm.Eloquent.Table("sys_config").Where("configId = ?", id).First(&update).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table("sys_config").Model(&update).Updates(&e).Error; err != nil {
		return
	}
	return
}
func (e *Config) GetConfig(key string) (update Config, err error) {
	sql := `select * from sys_config where configKey = ?`
	err = orm.Eloquent.Raw(sql,key).Scan(&update).Error
	if update==(Config{}){
		return update,errors.New("未配置转出钱包！")
	}
	return
}
func (e *Config) UpdateConfig(value,id string) ( err error) {
	sql := `update sys_config set configValue=? where configId = ?`
	err = orm.Eloquent.Exec(sql,value,id).Error
	return
}

func (e *Config) Delete(id int64) (success bool, err error) {
	var mp = map[string]string{}
	mp["is_del"] = "1"
	mp["update_time"] = time2.Now().Format("2006/01/02 15:04:05")
	mp["update_by"] = e.UpdateBy
	if err = orm.Eloquent.Table("sys_config").Where("configId = ?", id).Update(mp).Error; err != nil {
		success = false
		return
	}
	success = true
	return
}
func (e *Config) Export() (string, error) {
	//1:获取需要导出的数据
	var doc []Config
	table := orm.Eloquent.Select("*").Table("sys_config")
	table.Find(&doc)

	//2：生成导出文件
	xlsFile := xlsx.NewFile()
	sheet, err := xlsFile.AddSheet("系统配置")
	log.Print("XLSX:", sheet.Name, " ERROR:", err)
	if err != nil {
		return "", err
	}

	titles := []string{"ID", "配置名称", "值", "类型", "备注"}
	row := sheet.AddRow()

	var cell *xlsx.Cell
	//设置标题
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}

	//导出值
	for _, v := range doc {
		values := []string{
			strconv.Itoa(int(v.ID)),
			v.Title,
			v.Value,
			v.Group,
			v.Remark,
		}
		row = sheet.AddRow()
		for _, value := range values {
			cell = row.AddCell()
			cell.Value = value
		}
	}
	time := strconv.Itoa(int(time2.Now().Unix()))
	filename := "配置信息" + time + export.EXT

	dirFullPath := export.GetExcelFullPath()
	err = file.IsNotExistMkDir(dirFullPath)
	log.Print("DIR_FULL_PATH:", dirFullPath, " ERROR:", err)
	if err != nil {
		return "", err
	}
	err = xlsFile.Save(dirFullPath + filename)

	log.Print("SAVE_FILE_NAME:", filename, " ERROR:", err)
	if err != nil {
		return "", err
	}

	return filename, nil
}
