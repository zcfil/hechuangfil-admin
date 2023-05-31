package utils

import (
	"encoding/json"
	"strconv"
)

//Pagination 分页数据
type Pagination struct {
	PageIndex   int         `json:"pageIndex"`
	PageSize int         `json:"pageSize"`
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
}
type PageData struct {
	PageIndex   int         `json:"pageIndex"`
	PageSize int         `json:"pageSize"`
	List     interface{} `json:"list"`
	Total    int         `json:"total"`
	Code	 int		`json:"code"`
	Summation     interface{} `json:"summation"`
}
func NewPagination(pageNo, pageSize, total int, list interface{}) *Pagination {

	return &Pagination{
		PageIndex:   pageNo,
		PageSize: pageSize,
		List:     list,
		Total:    total,
	}
}

//Pagination
func NewPageData(param map[string]string, list interface{}) *PageData {
	no,_ := strconv.Atoi(param["pageIndex"])
	size,_ := strconv.Atoi(param["pageSize"])
	to,_ := strconv.Atoi(param["total"])
	return &PageData{
		PageIndex:   no,
		PageSize: size,
		List:     list,
		Total:    to,
		Code: 200,
	}
}
func NewPageDataTotal(param map[string]string, list interface{},total interface{}) *PageData {
	no,_ := strconv.Atoi(param["pageIndex"])
	size,_ := strconv.Atoi(param["pageSize"])
	to,_ := strconv.Atoi(param["total"])
	return &PageData{
		PageIndex:   no,
		PageSize: size,
		List:     list,
		Total:    to,
		Summation: total,
		Code: 200,
	}
}
func (p *Pagination)String() string {
	bytes, _ := json.Marshal(p)
	return string(bytes)
}

