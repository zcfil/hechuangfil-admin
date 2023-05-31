package testing

import (
	"hechuangfil-admin/config"
	"hechuangfil-admin/models"
	"testing"

	log "hechuangfil-admin/logrus"
)

func Test_settlement(t *testing.T) {
	if config.ApplicationConfig.IsInit {
		if err := models.InitDb(); err != nil {
			log.Fatal("数据库初始化失败！")
		} else {
			config.SetApplicationIsInit()
		}
	}

	log.NewLogger(config.ApplicationConfig.LogPath)

	settle := models.NewSettlement()
	settle.SettleFunc()
}
