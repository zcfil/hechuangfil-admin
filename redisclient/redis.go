package redisclient

import (
	"encoding/json"
	"hechuangfil-admin/handler"
	"sync"
)


const (
	LOTUS_HEIGHT = "lotus_height"	//记录爬下的lotus高度
	FIL_ADDRESS = "fil_address"		//钱包地址列表
	RECHARGE_RECORD = "recharge_record"		//已到账消息
)

func GetLotusHeight()(int64,error){
	return handler.RedisClient.Get(LOTUS_HEIGHT).Int64()
}
func SetLotusHeight(height int64)error{
	return handler.RedisClient.Set(LOTUS_HEIGHT,height,-1).Err()
}
func GetFilAddress()map[string]string{
	FilAddress := make(map[string]string)
	by,err := handler.RedisClient.Get(FIL_ADDRESS).Bytes()
	if err!=nil{
		return FilAddress
	}
	_ = json.Unmarshal(by,&FilAddress)
	return FilAddress
}
func GetRechargeRecord()map[string]string{
	RechargeRecord := make(map[string]string)
	by,err := handler.RedisClient.Get(RECHARGE_RECORD).Bytes()
	if err!=nil{
		return RechargeRecord
	}
	_ = json.Unmarshal(by,&RechargeRecord)
	return RechargeRecord
}
var Mux sync.Mutex
func UpdateRechargeRecord(record map[string]string)(map[string]string,error){
	Mux.Lock()
	defer func() {
		Mux.Unlock()
	}()
	by,_ := handler.RedisClient.Get(RECHARGE_RECORD).Bytes()
	RechargeRecord := make(map[string]string)
	if len(by)>0{
		_ = json.Unmarshal(by,&RechargeRecord)
	}
	for k,v := range record{
		RechargeRecord[k] = v
	}
	b,_ := json.Marshal(RechargeRecord)
	if err := handler.RedisClient.Set(RECHARGE_RECORD,string(b),-1).Err();err!=nil{
		return RechargeRecord,err
	}
	return RechargeRecord,nil
}
