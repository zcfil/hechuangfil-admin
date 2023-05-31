package models

import (
	"errors"
	"github.com/tealeg/xlsx"
	"mime/multipart"
	"strings"
	orm "hechuangfil-admin/database"
	"hechuangfil-admin/pkg/eth"
)
type Addr struct {
	Id int64 `json:"id"`
	Addr string `json:"addr"`
	State int `json:"state"`
}
//导入数据
func (a *Addr)UploadAddress(file multipart.File,Size int64)error{
	buf := make([]byte,Size)
	n,_ := file.Read(buf)

	xf ,_ := xlsx.OpenBinary(buf[:n])
	sql := `insert into addr(addr)values`
	for _,sheet := range xf.Sheets{
		for _,row := range sheet.Rows{
			for i,cell := range row.Cells{
				if i==0{
					addr := cell.String()
					if eth.ValidAddress(addr) {
						return errors.New("地址"+addr+"不合法！")
					}
					sql += `('`+addr+`'),`
				}
			}
		}
	}
	if err := orm.Eloquent.Exec(sql[:len(sql)-1]).Error;err!=nil{
		str := strings.Split(err.Error(),"'")
		return errors.New("已存在地址:"+str[1])
	}
	return nil
}

