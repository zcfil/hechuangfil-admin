package main

import (
	"hechuangfil-admin/pkg/googleAuthenticator"
	"flag"
	"fmt"
)

//生成秘钥
func createSecret(ga *googleAuthenticator.GAuth) string {
	secret, err := ga.CreateSecret(16)
	if err != nil {
		return ""
	}
	return secret
}

//获取code
func getCode(ga *googleAuthenticator.GAuth, secret string) string {
	code, err := ga.GetCode(secret)
	if err != nil {
		return "*"
	}
	return code
}

//校验code
func verifyCode(ga *googleAuthenticator.GAuth, secret, code string) bool {
	// 1:30sec
	ret, err := ga.VerifyCode(secret, code, 1)
	if err != nil {
		return false
	}
	return ret
}

func main() {
	flag.Parse()
	ga := googleAuthenticator.NewGAuth()
	//secret := flag.Arg(0)+
	secret := "6AXXTCAMWI2PLHQ7" //createSecret(ga) // "ZCSDXNLJ5GZX7JMB"
	//fmt.Println(secret)
	//fmt.Println(getCode(ga, secret))

	fmt.Println(verifyCode(ga, secret, "983471"))
}
