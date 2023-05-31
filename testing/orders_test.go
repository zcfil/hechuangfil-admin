package testing

import (
	"fmt"
	"hechuangfil-admin/config"
	log "hechuangfil-admin/logrus"
	"hechuangfil-admin/models"
	"testing"
)

func Test_orders(t *testing.T) {
	if config.ApplicationConfig.IsInit {
		if err := models.InitDb(); err != nil {
			log.Fatal("数据库初始化失败！")
		} else {
			config.SetApplicationIsInit()
		}
	}

	log.NewLogger(config.ApplicationConfig.LogPath)

	order := models.NewOrder()
	param := map[string]string {
		"pageIndex":"1",
		"pageSize":"10",
	}

	res, _ := order.GetOrders(param)
	fmt.Println("res:", res)
}