package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"hechuangfil-admin/auth"
	"hechuangfil-admin/config"
	orm "hechuangfil-admin/database"
	log "hechuangfil-admin/logrus"
	"hechuangfil-admin/models"
	"hechuangfil-admin/router"
)

var Cron = cron.New()

func init() {
	//0 */3 * * * ?    0 15,30,45 * * * ?
	//Cron.AddFunc("0 15,30,45 * * * ?", auth.CornTask)
	//Cron.AddFunc("@daily", auth.CornTaskGKC)
	//
	//Cron.Start()
}

// @title tco-admin API
// @version 0.0.1
// @description GKC管理系统的接口文档

// @description
// @license.name MIT
// @license.url
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	gin.SetMode(gin.DebugMode)

	if config.ApplicationConfig.IsInit {
		if err := models.InitDb(); err != nil {
			log.Fatal("数据库初始化失败！")
		} else {
			config.SetApplicationIsInit()
		}
	}

	log.NewLogger(config.ApplicationConfig.LogPath)

	// 获取链上数据
	data := models.NewDataCrawler(config.ApplicationConfig.DataCrawlerPeriod)
	data.Start()

	// 结算
	_ = models.NewSettlement()

	r := router.InitRouter()
	ctx := context.Background()
	auth.GetBlock()
	auth.CashSweep(ctx)
	defer orm.Eloquent.Close()
	if err := r.Run(config.ApplicationConfig.Host + ":" + config.ApplicationConfig.Port); err != nil {
		log.Fatal(err)
	}
}

