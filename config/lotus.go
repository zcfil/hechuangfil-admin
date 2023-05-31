package config

import (
	"github.com/spf13/viper"
)
type Lotus struct {
	Host string
	Token string
}
//type YungoRpc struct {
//	Rpc apistruct.FullNodeStruct
//	Closer  jsonrpc.ClientCloser
//}

func InitLotus(cfg *viper.Viper) *Lotus {
	return &Lotus{
		Host:     cfg.GetString("host"),
		Token: 	cfg.GetString("token"),

	}
}
var LotusConfig = new(Lotus)
