package testing

import (
	_ "hechuangfil-admin/config"
	_ "hechuangfil-admin/database"
	"hechuangfil-admin/models"
	"testing"
)

func Test_getData(t *testing.T) {
	data := models.NewDataCrawler(5)
	data.Start()

	//for true {
	//
	//}
}