package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

//生成32位md5字串[小写]
func GenLowMD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

//生成32位md5字串[大写]
func GenUpperMD5(data []byte) string {
	h := md5.New()
	h.Write(data)
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

//返回一个16位md5加密后的字符串
func Gen16MD5(data []byte) string {
	return GenLowMD5(data)[8:24]
}
