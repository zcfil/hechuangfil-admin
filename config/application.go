package config

import "github.com/spf13/viper"

type Application struct {
	IsInit        bool
	ReadTimeout   int
	WriterTimeout int
	Host          string
	Port          string
	Name          string
	JwtSecret     string
	LogPath       string
	Env           string
	EnvMsg        string

	RuntimeRootPath string
	ExportSavePath  string

	SetHeight	int64

	DataCrawlerPeriod  int64		// 爬取数据高度的周期
	SalesmanRatio	float64			//业务员分润比例

}

func InitApplication(cfg *viper.Viper) *Application {
	return &Application{
		IsInit:          cfg.GetBool("isInit"),
		ReadTimeout:     cfg.GetInt("readTimeout"),
		WriterTimeout:   cfg.GetInt("writerTimeout"),
		Host:            cfg.GetString("host"),
		Port:            cfg.GetString("port"),
		Name:            cfg.GetString("name"),
		JwtSecret:       cfg.GetString("jwtSecret"),
		LogPath:         cfg.GetString("logPath"),
		Env:             cfg.GetString("env"),
		EnvMsg:          cfg.GetString("envMsg"),
		RuntimeRootPath: cfg.GetString("runTimeRootPath"),
		ExportSavePath:  cfg.GetString("exportSavePath"),
		DataCrawlerPeriod: cfg.GetInt64("dataCrawlerPeriod"),
		SetHeight: cfg.GetInt64("setheight"),
		SalesmanRatio:  cfg.GetFloat64("salesmanratio"),
	}
}

var ApplicationConfig = new(Application)
