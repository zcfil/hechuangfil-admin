package handler

import (
	"context"
	"github.com/filecoin-project/lotus/api/client"
	"hechuangfil-admin/config"
	"hechuangfil-admin/pkg/lotus"
	"log"
	"net/http"
)

const (
	SEND = 0
	PC2  = 6
	C2   = 7
)

func init() {
	//if err := NewLotusRpc(config.LotusConfig); err != nil {
	//	log.Fatal("lotus连接失败！", err.Error())
	//}
	var err error
	if lotus.FullAPI, _, err = NewLotusApi(config.LotusConfig); err != nil {
		log.Fatal("lotus连接失败！", err.Error())
	}
}

// 云构LOTUS api
func NewLotusRpc(l *config.Lotus) error {

	headers := http.Header{"Authorization": []string{"Bearer " + l.Token}}

	_, err := jsonrpc.NewMergeClient(context.Background(), "ws://"+l.Host+"/rpc/v0", "Filecoin", []interface{}{&lotus.FullAPI.Internal, &lotus.FullAPI.CommonStruct.Internal}, headers)
	if err != nil {
		log.Println("connecting with lotus failed: %s", err)
		return err
	}

	return nil
}

func NewLotusApi(l *config.Lotus) (v0api.FullNode, jsonrpc.ClientCloser, error) {
	headers := http.Header{"Authorization": []string{"Bearer " + l.Token}}
	//var err error
	lapi, cl, err := client.NewFullNodeRPCV0(context.Background(), "http://"+l.Host+"/rpc/v0", headers)
	if err != nil {
		log.Println("connecting with lotus failed: ", err.Error())
		return nil, nil, err
	}
	return lapi, cl, nil
}
