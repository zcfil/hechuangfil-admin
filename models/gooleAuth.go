package models

/**
 *@Project     hechuangfil-admin
 *@Author      king
 *@CreateTime  2020/4/13 11:47 下午
 *@ClassName   gooleAuth
 *@Description TODO  谷歌验证器
 */

type GoogleAuth struct {
	GoogleSecret  string `json:"googleSecret,omitempty"`
	GoogleAccount string `json:"googleAccount,omitempty"`
	URL           string `json:"url,omitempty"`
	Code          string `json:"code,omitempty"`
	IsBind        bool   `json:"isBind,omitempty"`
}

//  "GoogleSecret":"Z6QZ2QZ5CZE5I5WE",
//	"googleAccount":"纽币:18902538879",
//	"url":"otpauth://totp/纽币:18902538879?secret=Z6QZ2QZ5CZE5I5WE"
