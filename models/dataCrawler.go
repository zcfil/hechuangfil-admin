package models

import (
	"errors"
	"fmt"
	orm "hechuangfil-admin/database"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type DataCrawler struct {
	periodTime int64
}

func NewDataCrawler(periodTime int64) *DataCrawler {
	dc := new(DataCrawler)
	dc.periodTime =  periodTime
	return dc
}


func (this *DataCrawler) Start() {
	go func() {
		c := time.Tick(time.Duration(this.periodTime) * time.Second)
		for {
			this.writeData()
			<- c
		}
	}()
}

func (this *DataCrawler) writeData() {
	bodyStr, err := this.GetData()
	if err != nil {
		fmt.Println("Get data body error:", err)
		return
	}
	height, err := this.getHeight(bodyStr)
	if err != nil {
		fmt.Println("Get height error:", err)
		return
	}

	aFil, err := this.getAverageFil(bodyStr)
	if err != nil {
		fmt.Println("Get average fil error:", err)
		return
	}

	// 插入到数据库
	sql := `insert into filearnings (height, create_time, average_fil) values (` + height + `, now(), ` + aFil + `)`
	if err := orm.Eloquent.Exec(sql).Error; err != nil {
		fmt.Println("插入爬取的数据错误")
	}
}

func (this *DataCrawler)GetData() (bodyStr string, err error)  {
	resp, err1 := http.Get("https://filfox.info/zh")
	if err1 != nil {
		fmt.Println(err)
		err = err1
		return
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		err = err2
		return
	}

	bodyStr = string(body)
	return
}

// 区块高度
func (this *DataCrawler) getHeight(bodyStr string) (height string, err error) {
	strList := strings.Split(bodyStr, `区块高度 </span><!----></div><div class="text-left lg:text-center text-sm lg:text-2xl items-start lg:mx-auto"> `)
	l := len(strList)
	// 取第二个字符串
	fmt.Println("区块高度的字符串长度:", l)
	if l < 2 {
		err = errors.New("数据格式错误")
		return
	}
	// 取出高度的字符串数据
	hBin := make([]byte, 0)
	for _, b := range strList[1] {
		if b == ' ' {
			break
		}
		if b == ',' {
			continue
		}
		hBin = append(hBin, byte(b))
	}

	height = string(hBin)
	return
}

// 获取平均Fil
func (this *DataCrawler) getAverageFil(bodyStr string) (averageFil string, err error) {
	strList := strings.Split(bodyStr, ` FIL/TiB`)
	fmt.Println("list len:", len(strList))
	if len(strList) <= 0 {
		err = errors.New("数据格式错误")
		return
	}
	// 第一个数组里面是我们要的数值, 再次切割取出需要的数值
	// 从后向前取 然后反转字符列表
	filList := make([]byte, 0)
	l := len(strList[0])
	for i := l-1; i>= 0 ;i-- {
		if strList[0][i] == ' ' {
			break
		}
		filList = append(filList, strList[0][i])
	}
	for i, j := 0, len(filList)-1; i < j; i, j = i+1, j-1 {
		filList[i], filList[j] = filList[j], filList[i]
	}
	averageFil = string(filList)
	return
}

func (this *DataCrawler) test() {
	//reg1 := regexp.MustCompile(`<div class="text-left lg:text-center text-sm lg:text-2xl items-start lg:mx-auto"> (.*)`)
	//if reg1 == nil {
	//	fmt.Println("regexp err")
	//	return
	//}
	////根据规则提取关键信息
	//result1 := reg1.FindAllStringSubmatch(strList[0], -1)
	//for _, l := range result1[0] {
	//	fmt.Println("l = ", l)
	//}
	//fmt.Println("result1 = ", len(result1))
}