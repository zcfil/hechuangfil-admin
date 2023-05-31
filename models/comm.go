package models

import (
	"encoding/json"
	"errors"
	"fmt"
	orm "hechuangfil-admin/database"
	//"github.com/astaxie/beego/orm"
	"github.com/tealeg/xlsx"
	"log"
	"strconv"
	"strings"
	"hechuangfil-admin/pkg/export"
	"hechuangfil-admin/pkg/file"
	"hechuangfil-admin/utils"
)

type Func func(param map[string]string) (interface{}, error)     //定义函数结构
type FuncTotal func(param map[string]string) (interface{},interface{}, error)     //定义函数结构
//type FuncNew func(param map[string]interface{}) (in, int32, error) //定义函数结构

//获取总数
//func GetTotalCount(dbtype *gorm.DB, sql string, value ...string) (count string) {
//	sql = ` select count(1) totalcount from ( ` + sql + `)a`
//	sqltotal := ""
//	if strings.Contains(sql, "?") {
//		str := strings.Split(sql, "?")
//		for i := 0; i < len(str); i++ {
//			if i < len(str)-1 {
//				sqltotal += str[i] + `'` + value[i] + `'`
//			} else {
//				sqltotal += str[i]
//			}
//		}
//	} else {
//		sqltotal = sql
//	}
//	//总数
//	type Total struct {
//		Totalcount int
//	}
//	var total Total
//	dbtype.Raw(sql).Scan(&total)
//
//	return strconv.Itoa(total.Totalcount)
//}
func GetTotalCount(sql string, value ...string) (count string) {
	sql = ` select count(1) totalcount from ( ` + sql + `)a`
	sqltotal := ""
	if strings.Contains(sql, "?") {
		str := strings.Split(sql, "?")
		for i := 0; i < len(str); i++ {
			if i < len(str)-1 {
				sqltotal += str[i] + `'` + value[i] + `'`
			} else {
				sqltotal += str[i]
			}
		}
	} else {
		sqltotal = sql
	}
	//总数
	type Total struct {
		Totalcount int
	}
	var total Total
	orm.Eloquent.Raw(sql).Scan(&total)

	return strconv.Itoa(total.Totalcount)
}
func GetTotalCount1(sql string, value ...string) (count int) {
	sql = ` select count(1) totalcount from ( ` + sql + `)a`
	sqltotal := ""
	if strings.Contains(sql, "?") {
		str := strings.Split(sql, "?")
		for i := 0; i < len(str); i++ {
			if i < len(str)-1 {
				sqltotal += str[i] + `'` + value[i] + `'`
			} else {
				sqltotal += str[i]
			}
		}
	} else {
		sqltotal = sql
	}
	//总数
	type Total struct {
		Totalcount int
	}
	var total Total
	orm.Eloquent.Raw(sql).Scan(&total)

	return total.Totalcount
}
func GetSummation(sql string, value ...string) (count string) {
	sql = ` select count(1) totalcount from ( ` + sql + `)a`
	sqltotal := ""
	if strings.Contains(sql, "?") {
		str := strings.Split(sql, "?")
		for i := 0; i < len(str); i++ {
			if i < len(str)-1 {
				sqltotal += str[i] + `'` + value[i] + `'`
			} else {
				sqltotal += str[i]
			}
		}
	} else {
		sqltotal = sql
	}
	//总数
	type Total struct {
		Totalcount int
	}
	var total Total
	orm.Eloquent.Raw(sql).Scan(&total)

	return strconv.Itoa(total.Totalcount)
}
//导出Excel
func GetExcelURL(fun Func, param map[string]string) (string, error) {
	//获取参数
	sheetstr := param["sheet"]
	filefield := param["filefield"]
	filename := param["filename"]
	title := param["title"]

	//获取数据
	res,  _ := fun(param)
	//数据转为[]map
	fmt.Println(res)
	b, _ := json.Marshal(res)
	resmap := make([]map[string]interface{}, 0)
	if err := json.Unmarshal(b, &resmap); err != nil {
		fmt.Println(err)
	}
	//生成导出文件
	xlsFile := xlsx.NewFile()
	sheet, err := xlsFile.AddSheet(sheetstr)
	if err != nil {
		//return "", err
	}

	row := sheet.AddRow()
	var cell *xlsx.Cell
	names := strings.Split(filename, ",")
	//设置列名
	for _, name := range names {
		cell = row.AddCell()
		cell.Value = name
	}
	//查询字段表，字典类型需要做转换
	type Dict struct {
		Dicttype string
	}
	//var dic []Dict
	//sql := `  select dicttype from sys_dict_type `
	//orm.Eloquent.Raw(sql).Scan(&dic)
	//dictstr := ""
	//for i := 0; i < len(dic); i++ {
	//	dictstr += dic[i].Dicttype + ","
	//}

	//导出值
	fields := strings.Split(filefield, ",")
	for i := 0; i < len(resmap); i++ {
		row = sheet.AddRow()
		for _, field := range fields {
			//先转为字符串
			value := utils.Strval(resmap[i][field])
			//判断是否时间戳
			//times := strings.Split(strings.ToLower(field),"time")
			//dates := strings.Split(strings.ToLower(field),"date")
			//if len(field) > 3 {
			//	if (strings.ToLower(field[len(field)-4:]) == "time" || strings.ToLower(field[len(field)-4:]) == "date") && len(value) == 10 {
			//		//时间戳转换为字符串
			//		value = time.Unix(int64(resmap[i][field].(float64)), 0).Format("2006/01/02 15:04:05")
			//	}
			//}
			//字典数据转换
			//if strings.Contains(dictstr, field) {
			//	str := strings.Replace(field, "_","*",1)
			//	f := strings.Split(str, "*")
			//	if len(f) > 1 {
			//		value = SatausToChese(field, resmap[i][f[1]].(float64))
			//	}
			//}
			////字典数据转换
			//cn := strings.Split(names[j], "或")
			//if len(cn) > 1 {
			//	if value == "true" || value == "1" {
			//		value = cn[0]
			//	} else {
			//		value = cn[1]
			//	}
			//}
			//if strings.Contains(names[j], "是否") {
			//	if value == "true" || value == "1" {
			//		value = "是"
			//	} else {
			//		value = "否"
			//	}
			//}
			cell = row.AddCell()
			cell.Value = value
		}
	}

	//time := strconv.Itoa(int(time.Now().Unix()))
	Name := title + export.EXT

	dirFullPath := export.GetExcelFullPath()
	err = file.IsNotExistMkDir(dirFullPath)
	log.Print("DIR_FULL_PATH:", dirFullPath, " ERROR:", err)
	if err != nil {
		//return "", err
	}
	err = xlsFile.Save(dirFullPath + Name)

	log.Print("SAVE_FILE_NAME:", Name, " ERROR:", err)
	if err != nil {
		//return "", err
	}
	return "export/" + Name, err
}
func GetExcelTotal(fun FuncTotal, param map[string]string) (string, error) {
	//获取参数
	sheetstr := param["sheet"]
	filefield := param["filefield"]
	filename := param["filename"]
	title := param["title"]

	//获取数据
	res, total, _ := fun(param)
	//数据转为[]map
	b, _ := json.Marshal(res)
	resmap := make([]map[string]interface{}, 0)
	if err := json.Unmarshal(b, &resmap); err != nil {
		fmt.Println(err)
	}
	//生成导出文件
	xlsFile := xlsx.NewFile()
	sheet, err := xlsFile.AddSheet(sheetstr)
	if err != nil {
		//return "", err
	}

	row := sheet.AddRow()
	var cell *xlsx.Cell
	names := strings.Split(filename, ",")
	//设置列名
	for _, name := range names {
		cell = row.AddCell()
		cell.Value = name
	}

	//导出值
	fields := strings.Split(filefield, ",")
	for i := 0; i < len(resmap); i++ {
		row = sheet.AddRow()
		for _, field := range fields {
			//先转为字符串
			value := utils.Strval(resmap[i][field])
			cell = row.AddCell()
			cell.Value = value
		}
	}
	//合计
	t, _ := json.Marshal(total)
	totalmap := make(map[string]interface{}, 0)
	if err := json.Unmarshal(t, &totalmap); err != nil {
		fmt.Println(err)
	}
	rowt := sheet.AddRow()
	cellt := rowt.AddCell()
	cellt.Value = "合计"
	for i,v := range fields{
		if i>0{
			cellt = rowt.AddCell()
		}
		if val,ok := totalmap[v];ok {
			cellt.Value = utils.Strval(val)
		}
	}

	Name := title + export.EXT

	dirFullPath := export.GetExcelFullPath()
	err = file.IsNotExistMkDir(dirFullPath)
	log.Print("DIR_FULL_PATH:", dirFullPath, " ERROR:", err)
	if err != nil {
		//return "", err
	}
	err = xlsFile.Save(dirFullPath + Name)

	log.Print("SAVE_FILE_NAME:", Name, " ERROR:", err)
	if err != nil {
		//return "", err
	}
	return "export/" + Name, err
}
func GetExcelTotal1(fun FuncTotal, param map[string]string) (string, error) {
	//获取参数
	sheetstr := param["sheet"]
	filefield := param["filefield"]
	filename := param["filename"]
	title := param["title"]

	//获取数据
	res, total, _ := fun(param)
	//数据转为[]map
	b, _ := json.Marshal(res)
	resmap := make([]map[string]interface{}, 0)
	if err := json.Unmarshal(b, &resmap); err != nil {
		fmt.Println(err)
	}
	//生成导出文件
	xlsFile := xlsx.NewFile()
	sheet, err := xlsFile.AddSheet(sheetstr)
	if err != nil {
		//return "", err
	}

	row := sheet.AddRow()
	var cell *xlsx.Cell
	names := strings.Split(filename, ",")
	//设置列名
	for _, name := range names {
		cell = row.AddCell()
		cell.Value = name
	}

	//导出值
	fields := strings.Split(filefield, ",")
	for i := 0; i < len(resmap); i++ {
		row = sheet.AddRow()
		for _, field := range fields {
			//先转为字符串
			value := utils.Strval(resmap[i][field])
			cell := row.AddCell()
			cell.Value = value
			if row.Cells[0].Value == "合计"{
				st := cell.GetStyle()
				st.Font.Bold = true
			}
		}
		if row.Cells[0].Value == "合计"{
			sheet.AddRow()
		}
	}
	//合计
	t, _ := json.Marshal(total)
	totalmap := make(map[string]interface{}, 0)
	if err := json.Unmarshal(t, &totalmap); err != nil {
		fmt.Println(err)
	}
	rowt := sheet.AddRow()
	cellt := rowt.AddCell()
	cellt.Value = "总合计"
	st := cellt.GetStyle()
	st.Font.Bold = true
	for i,v := range fields{
		if i>0{
			cellt = rowt.AddCell()
		}
		st := cellt.GetStyle()
		st.Font.Bold = true
		if val,ok := totalmap[v];ok {
			cellt.Value = utils.Strval(val)
		}
	}

	Name := title + export.EXT

	dirFullPath := export.GetExcelFullPath()
	err = file.IsNotExistMkDir(dirFullPath)
	log.Print("DIR_FULL_PATH:", dirFullPath, " ERROR:", err)
	if err != nil {
		//return "", err
	}
	err = xlsFile.Save(dirFullPath + Name)

	log.Print("SAVE_FILE_NAME:", Name, " ERROR:", err)
	if err != nil {
		//return "", err
	}
	return "export/" + Name, err
}
func GetExcelSummary(fun FuncTotal, param map[string]string) (string, error) {
	//获取参数
	sheetstr := param["sheet"]
	filefield := ""
	filename := ""
	title := param["title"]

	//获取数据
	res,total,  _ := fun(param)
	re := res.(map[string]interface{})
	filefield = "customername,nick_name,amount,"+re["words"].(string)+",company,salesman,remark,create_time"
	filename = "客户姓名,业务员姓名,投资,"+re["keys"].(string)+",业务部门,业务员,备注,投资时间"
	//数据转为[]map
	b, _ := json.Marshal(re["profit"])
	resmap := make([]map[string]interface{}, 0)
	if err := json.Unmarshal(b, &resmap); err != nil {
		fmt.Println(err)
	}
	//生成导出文件
	xlsFile := xlsx.NewFile()
	sheet, err := xlsFile.AddSheet(sheetstr)
	if err != nil {
		//return "", err
	}

	row := sheet.AddRow()
	var cell *xlsx.Cell
	names := strings.Split(filename, ",")
	//设置列名
	for _, name := range names {
		cell = row.AddCell()
		cell.Value = name
	}
	//导出值
	fields := strings.Split(filefield, ",")
	for i := 0; i < len(resmap); i++ {
		row = sheet.AddRow()
		for _, field := range fields {
			value := utils.Strval(resmap[i][field])
			cell = row.AddCell()
			cell.Value = value
		}
	}
	//合计
	t, _ := json.Marshal(total)
	totalmap := make(map[string]interface{}, 0)
	if err := json.Unmarshal(t, &totalmap); err != nil {
		fmt.Println(err)
	}
	rowt := sheet.AddRow()
	cellt := rowt.AddCell()
	cellt.Value = "合计"
	for i,v := range fields{
		if i>0{
			cellt = rowt.AddCell()
		}
		if val,ok := totalmap[v];ok {
			cellt.Value = utils.Strval(val)
		}
	}


	Name := title + export.EXT

	dirFullPath := export.GetExcelFullPath()
	err = file.IsNotExistMkDir(dirFullPath)
	log.Print("DIR_FULL_PATH:", dirFullPath, " ERROR:", err)
	if err != nil {
		//return "", err
	}
	err = xlsFile.Save(dirFullPath + Name)

	log.Print("SAVE_FILE_NAME:", Name, " ERROR:", err)
	if err != nil {
		//return "", err
	}
	return "export/" + Name, err
}
//判断数据库的状态，转化为中文
func SatausToChese(dictType string, dictValue float64) (statstr string) {
	sql := " select dictlabel from sys_dict_data where dictType = ? and dictValue = ? "

	type Dic struct {
		Dictlabel string
	}
	var dic Dic
	orm.Eloquent.Raw(sql,dictType,dictValue).Scan(&dic)
	statstr = dic.Dictlabel
	return
}
//获取字典类型
func GetDictData(dictType string) (result map[string]string) {
	sql := " select dictvalue,dictlabel from sys_dict_data where dictType = ? "

	type Dic struct {
		Dictlabel string
		Dictvalue string
	}
	var dic []Dic
	orm.Eloquent.Raw(sql,dictType).Scan(&dic)
	result = make(map[string]string)
	for _,v := range dic{
		result[v.Dictlabel] = v.Dictvalue
	}
	return
}
//获取配置
func GetConfigData(name string) (result string) {
	sql := " select value from config where name = ? "

	type Conf struct {
		Value string
	}
	var conf Conf
	orm.Eloquent.Raw(sql,name).Scan(&conf)
	result = conf.Value
	return
}
func IsCheckByUid(uids string)(b bool,err error){
	ids := strings.Split(uids,",")
	sql := ` select id from user where id in(?) `
	type Uid struct {
		Id int64
	}
	var uid []Uid
	err = orm.Eloquent.Raw(sql,uids).Scan(&uid).Error
	if err!=nil{
		log.Println(err)
		return
	}
	//如果查出数量小于传入数量，要么重复，要么有用户不存在
	if len(uid)<len(ids){
		check :=make(map[string]int)
		for i:=0;i<len(uid);i++{
			check[strconv.FormatInt(uid[i].Id,10)] = i+1
		}
		//用户不存在
		flg := make(map[string]int)
		for j:=0;j<len(ids);j++{
			if _,ok := check[ids[j]];!ok{
				return false,errors.New("用户账号"+ids[j]+"不存在！")
			}
			//判断重复用户
			if _,ok := flg[ids[j]];ok{
				return false,errors.New("用户账号"+ids[j]+"重复！")
			}
			flg[ids[j]] = j+1
		}
	}
	return true,nil
}
