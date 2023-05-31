package handler

import (
	"context"
	"github.com/filecoin-project/go-jsonrpc"
	"hechuangfil-admin/pkg/lotus"

	//"github.com/filecoin-project/lotus/api/apistruct"
	"hechuangfil-admin/config"
	"log"
	"net/http"
)


const (
	SEND = 0
	PC2 = 6
	C2 = 7
)
func init() {
	if err := NewLotusRpc(config.LotusConfig);err!=nil{
		log.Fatal("lotus连接失败！",err.Error())
	}
}

//云构LOTUS api
func NewLotusRpc(l *config.Lotus) (error) {

	headers := http.Header{"Authorization": []string{"Bearer " + l.Token}}

	_, err := jsonrpc.NewMergeClient(context.Background(), "ws://"+l.Host+"/rpc/v0", "Filecoin", []interface{}{&lotus.FullAPI.Internal, &lotus.FullAPI.CommonStruct.Internal}, headers)
	if err != nil {
		log.Println("connecting with lotus failed: %s", err)
		return err
	}

	return nil
}