package config

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
	"log"
)
type OSSData struct {
	Endpoint   		string
	AccessKeyID     string
	AccessKeySecret string
	BucketName 		string
	Client			*oss.Client
}
var OSSConfig = new(OSSData)
func InitOSSDatabase(cfg *viper.Viper) *OSSData {
	client, err := oss.New(cfg.GetString("endpoint"), cfg.GetString("accessKeyID"), cfg.GetString("accessKeySecret"))
	if err!=nil{
		log.Println("OSS连接失败！",err)
		return nil
	}
	return &OSSData{
		Client:			 client,
		Endpoint:    	 cfg.GetString("endpoint"),
		AccessKeyID: 	 cfg.GetString("accessKeyID"),
		AccessKeySecret: cfg.GetString("accessKeySecret"),
		BucketName: 	 cfg.GetString("bucketName"),
	}
}
