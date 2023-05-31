package config

import (
	"github.com/spf13/viper"
	"log"
	"hechuangfil-admin/pkg/eth"
)

var adminDatabase *viper.Viper
var cfgApplication *viper.Viper
var cfgRedisConn *viper.Viper
var lotusConn *viper.Viper


var ETHConn *eth.ETHconn

//初始化配置文件数据
func init() {
	viper.SetConfigName("settings") // name of config file (without extension)
	viper.AddConfigPath("./config") // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {
		log.Println(err) // Handle errors reading the config file
	}

	//admin管理系统数据库配置
	adminDatabase = viper.Sub("settings.admindb")
	if adminDatabase == nil {
		panic("config not found settings.database")
	}
	AdminDatabaseConfig = InitAdminDatabase(adminDatabase)


	//应用配置
	cfgApplication = viper.Sub("settings.application")
	if cfgApplication == nil {
		panic("config not found settings.application")
	}
	ApplicationConfig = InitApplication(cfgApplication)

	//Redis配置
	cfgRedisConn = viper.Sub("settings.redis")
	if cfgRedisConn == nil {
		panic("config not found settings.redis")
	}
	RedisConnConfig = InitRedisConn(cfgRedisConn)

	//Lotus
	lotusConn = viper.Sub("settings.lotus")
	if lotusConn == nil {
		panic("config not found settings.lotus")
	}
	LotusConfig = InitLotus(lotusConn)
}

func SetApplicationIsInit() {
	viper.AddConfigPath("./config")
	viper.Set("settings.application.isInit", false)
	viper.WriteConfig()
}
