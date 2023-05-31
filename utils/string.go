package utils

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func StringToInt64(e string) (int64, error) {
	return strconv.ParseInt(e, 10, 64)
}

func IntToString(e int) string {
	return strconv.Itoa(e)
}

func Float64ToString(e float64) string {
	return strconv.FormatFloat(e, 'f', -1, 64)
}

func Int64ToString(e int64) string {
	return strconv.FormatInt(e, 10)
}

//获取URL中批量id并解析
func IdsStrToIdsInt64Group(key string, c *gin.Context) []int64 {
	IDS := make([]int64, 0)
	ids := strings.Split(c.Request.FormValue(key), ",")
	for i := 0; i < len(ids); i++ {
		ID, _ := strconv.ParseInt(ids[i], 10, 64)
		IDS = append(IDS, ID)
	}
	return IDS
}

func GetCurrntTime() string {
	return time.Now().Format("2006/01/02 15:04:05")
}

func GetLocation(ip string) string {
	if ip == "127.0.0.1" || ip == "localhost" {
		return "内部IP"
	}
	resp, err := http.Get("https://restapi.amap.com/v3/ip?ip=" + ip + "&key=3fabc36c20379fbb9300c79b19d5d05e")
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	fmt.Printf(string(s))

	m := make(map[string]string)

	err = json.Unmarshal(s, &m)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
	}
	if m["province"] == "" {
		return "未知位置"
	}
	return m["province"] + "-" + m["city"]
}

func StructToJsonStr(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

func GetBodyString(c *gin.Context) (string, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return string(body), nil
	} else {
		return "", err
	}
}

func JsonStrToMap(e string) (map[string]interface{}, error) {
	var dict map[string]interface{}
	if err := json.Unmarshal([]byte(e), &dict); err == nil {
		return dict, err
	} else {
		return nil, err
	}
}
//日期转化为时间戳
func DateToTimeStamp(date string,timeLayout string)(int64 ,error){
	if timeLayout=="" {
		timeLayout="2006-01-02 15:04:05"
	}
	loc, err := time.LoadLocation("Local")    //获取时区
	tmp, err := time.ParseInLocation(timeLayout, date, loc)
	return tmp.Unix(),err   //转化为时间戳 类型是int64
}
//分页 and 排序
//func LimitAndOrderBy(param map[string]interface{})string{
//	str := ""
//	//排序
//	if param["sort"]!=nil{
//		str += ` order by `+param["sort"].(string)
//		if param["order"]!=nil{
//			str +=" "+ param["order"].(string)
//		}
//	}
//	if param["isexp"]==nil{
//		param["isexp"]="0"
//	}
//	if param["isexp"].(string)!="1"{
//		//分页
//		switch param["pageIndex"].(type){
//		case int:
//			if param["pageIndex"]!=nil&&param["pageSize"]!=nil{
//				pageIndex := param["pageIndex"].(int)
//				pageSize := param["pageSize"].(int)
//				if pageIndex!=0 && pageSize!=0{
//					str += ` limit `+strconv.Itoa((pageIndex-1)*pageSize)+`,`+strconv.Itoa(pageSize)
//				}
//			}
//		case string:
//			if param["pageIndex"]!=nil&&param["pageSize"]!=nil{
//				pageIndex,_ := strconv.Atoi(param["pageIndex"].(string))
//				pageSize,_:= strconv.Atoi(param["pageSize"].(string))
//				if pageIndex!=0 && pageSize!=0{
//					str += ` limit `+strconv.Itoa((pageIndex-1)*pageSize)+`,`+param["pageSize"].(string)
//				}
//			}
//		}
//
//	}
//
//	return str
//}
func LimitAndOrderBy(param map[string]string)string{
	str := ""
	//排序
	if param["sort"]!=""{
		str += ` order by `+param["sort"]
		if param["order"]!=""{
			str +=" "+ param["order"]
		}
	}
	if param["isexp"]==""{
		param["isexp"]="0"
	}
	if param["isexp"]!="1"{
		//分页
		if param["pageIndex"]!=""&&param["pageSize"]!=""{
			pageNum,_ := strconv.Atoi(param["pageIndex"])
			pageSize,_:= strconv.Atoi(param["pageSize"])
			if pageNum!=0 && pageSize!=0{
				str += ` limit `+strconv.Itoa((pageNum-1)*pageSize)+`,`+param["pageSize"]
			}
		}

	}

	return str
}
//获取时间戳时间区间
func BeginToEndTimestampstr(param map[string]interface{})(con string){
	field := ""
	if param["field"]==nil{
		return
	}else{
		field = param["field"].(string)
	}

	if param["beginTime"]!=nil{
		beginTime := param["beginTime"].(string)
		if beginTime!=""{
			t1,_ := DateToTimeStamp(beginTime+" 00:00:00","2006-01-02 15:04:05")
			con +=" and "+ field+">= "+strconv.FormatInt(t1,10)
		}
	}
	if param["endTime"]!=nil{
		endTime := param["endTime"].(string)
		if endTime!=""{
			t2, _ := DateToTimeStamp(endTime+" 23:59:59","2006-01-02 15:04:05")
			con += " and "+ field+"<= "+strconv.FormatInt(t2,10)
		}
	}
	return con
}

//获取date时间区间
func BeginToEndDatestr(param map[string]interface{})(con string) {
	field := ""
	if param["field"] == nil {
		return
	} else {
		field = param["field"].(string)
	}

	if param["beginTime"] != nil {
		beginTime := param["beginTime"].(string)
		if beginTime != "" {
			con += " and " + field + ">= '" + beginTime+"'"
		}
	}
	if param["endTime"] != nil {
		endTime := param["endTime"].(string)
		if endTime != "" {
			con += " and " + field + "<= '" + endTime+" 23:59:59'"
		}
	}
	return con
}
//interface类型转string类型
func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
//获取一个随机的16进制字符串
func GetRandString(num int)string{
	s,_ := rand.Prime(rand.Reader,32)
	str := strconv.FormatInt(time.Now().UnixNano(),10)+s.String()
	sh := sha1.New()
	sh.Write([]byte(str))
	res := sh.Sum(nil)
	return  hex.EncodeToString(res)[0:num]
}
//将sql中的占位符'？'替换成map中的参数
//func SqlReplaceParames(sql string,param map[string]interface{})(string){
//	fa := false
//	start := 0
//	sqlstr := sql
//	for i:=0;i<len(sql);i++{
//		if sql[i]==':'{
//			start = i+1
//			fa = true
//		}
//		if (sql[i] == '\n'|| sql[i]=='\t'||sql[i]==' '||sql[i]==','||sql[i]==')'||sql[i]=='%')&&fa{
//			field := sql[start:i]
//			if param[field]!=nil&&start>3{
//				if sql[start-3]=='%'{
//					sqlstr = strings.Replace(sqlstr,"%%:"+field+"%%",`'%%`+param[field].(string)+`%%'`,-1)
//				}else if sql[start-2]=='%' {
//					sqlstr = strings.Replace(sqlstr,"%:"+field+"%",`'%`+param[field].(string)+`%'`,-1)
//				}else{
//					sqlstr = strings.Replace(sqlstr,":"+field,`'`+param[field].(string)+`'`,-1)
//				}
//				fa = false
//			}else{
//				sqlstr = field + "参数不存在!"
//				return sqlstr
//			}
//		}
//
//	}
//	return sqlstr
//}

//将sql中的占位符':'替换成map中的参数
func SqlReplaceParames(sql string,param map[string]string)(string){
	fa := false
	start := 0
	sqlstr := sql
	fl := true
	for i,v := range sql{
		if v==':'{
			start = i+1
			fa = true
		}
		if (v == '\n'|| v=='\t'||v==' '||v==','||v==')'||v=='%'||v=='"'||v=='='||len(sql)-1==i)&&fa&&fl{
			field := sql[start:i]
			//最后一个
			if len(sql)-1==i&&v!=' '&&v != '\n'&& v!='\t'&&v!=')'{
				field = sql[start:i+1]
			}
			if param[field]!=""{
				if sql[start-3]=='%'{
					sqlstr = strings.Replace(sqlstr,"%%:"+field+"%%",`'%%`+param[field]+`%%'`,1)
				}else if sql[start-2]=='%' {
					sqlstr = strings.Replace(sqlstr,"%:"+field+"%",`'%`+param[field]+`%'`,1)
				}else{
					flen := len(field)
					//避免包含于字段 如 :bank  :banknum
					if len(sql) > start+flen{
						v1 := sql[start+flen]
						//v2 := ','
						if v1 == '\n'|| v1=='\t'||v1==' '||v1==','||v1==')'||v1=='%'||v1=='"'||v1=='='{
							//fmt.Println(field,v1)
							sqlstr = strings.Replace(sqlstr,":"+field,`'`+param[field]+`'`,1)
							fa = false
						}
						continue
					}
					sqlstr = strings.Replace(sqlstr,":"+field,`'`+param[field]+`'`,-1)
				}
				fa = false
			}else{
				if _,ok := param[field];ok{
					sqlstr = strings.Replace(sqlstr,":"+field,`'`+param[field]+`'`,1)
					fa = false
					continue
				}
				if sql[i-1]=='\''||sql[i-1]=='"'{
					fa = false
					continue
				}
				sqlstr = field + " 参数不存在!"
				return sqlstr
			}
		}

	}
	return sqlstr
}