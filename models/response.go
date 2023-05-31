package models

import (
	"net/http"
)

//请求响应
type Response struct {
	// 代码
	Code int `json:"code" example:"200"`
	// 数据集
	Data interface{} `json:"data"`
	// 错误消息
	Msg string `json:"msg"`
}

func (res *Response) ReturnOK() *Response {
	res.Code = 200
	return res
}

func (res *Response) ReturnError(code int) *Response {
	res.Code = code
	return res
}

//客户端请求错误
func BadRequestResponse(tips string) Response {
	return Response{Code:http.StatusBadRequest,Msg:"请求错误:"+tips}
}
//服务器内部错误
func ServerErrorResponse(tips string) Response {
	return Response{Code:http.StatusInternalServerError,Msg:"服务器内部错误:"+tips}
}
//成功响应
func SuccessResponse(data interface{}) Response {
	return Response{Code:http.StatusOK,Data:data,Msg:"SUCCESS"}
}
