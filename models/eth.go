package models

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"
	"hechuangfil-admin/config"
	"hechuangfil-admin/pkg/eth"
)

//合约地址转帐
func TokenTraction(fromAddress, toAddress, token string, decimal int64, value float64) (string, error) {

	//收款地址截取去掉0x
	toaddr := strings.Split(toAddress, "0x")[1]

	//转账代币数量 十六进制的value值去掉0x并由0补够64位数
	//v:=strconv.FormatFloat(value,0.00000,18,64)
	//vl := IntToHex(int(value * 1e6)) //"000000000000000000000000"  USDT
	//log.Println("十六进制：",IntToHex(int(value * 1e6)))
	//IntToHex(int(value*1e18)) //"000000000000000000000000"

	//这是处理位数的代码段
	amount := big.NewFloat(float64(value))
	tenDecimal := big.NewFloat(math.Pow(10, float64(decimal)))
	convertAmount, _ := new(big.Float).Mul(tenDecimal, amount).Int(&big.Int{})

	//log.Println("转账数量：", convertAmount)

	has := fmt.Sprintf("%x", convertAmount) //格式化数据
	//log.Println("Int十六进制：",has)

	//vs := strings.Split(has, "0x")[1]
	//log.Println("BIG十六进制：",vs)

	//data拼接： “0x”+"23b872dd"+"from地址去掉0x并由0补够64位数"+"to地址去掉0x并由0补够64位数"+"十六进制的value值去掉0x并由0补够64位数"
	//data:="0x70a08231"+faddr+taddr+vstr //data拼接 "0xa9059cbb"
	//0xa9059cbb

	data := eth.TRANSFER_CODE + addPreZero(toaddr) + addPreZero(has) //data拼接
	//gaspric,_:= GasPrice()
	//fmt.Println("Data拼装：", data)
	t := eth.T{
		From:     fromAddress,
		Gas:      600000,                 //600000
		GasPrice: big.NewInt(6000000000), //big.NewInt(4500000000) 最快到账 60000000000  2500000000  普通：20000000000
		To:       token,                  //合约地址
		//Value:
		Data: data,
	}
	//hash,err:= client.EthPerSendTransaction(t,PWD)
	//log.Println("GasPrice",gaspric)
	hash, err := config.ETHConn.EthSendTransaction(t)

	if err != nil {
		log.Println("转账错误:", err.Error())
	}

	return hash, err

}

// 补齐64位，不够前面用0补齐
func addPreZero(num string) string {
	t := len(num)
	s := ""
	for i := 0; i < 64-t; i++ {
		s += "0"
	}
	return s + num
}

//获取合約余额
func GetContractBalance(address string, token string, decimals int) (float64,error) {

	addrSplit := strings.Split(address, "0x")[1] //地址去掉0x
	//data数据格式：最前边的“0x70a08231000000000000000000000000”是固定的，后边的是钱包地址（不带“0x”前缀）
	data := "0x70a08231000000000000000000000000" + addrSplit //data拼接

	t := eth.T{
		From: address, //查詢地址
		To:   token,   //合约地址
		Data: data,    //data
	}

	//获取代币的余额，要通过rpc接口得到接口为：eth_call
	balance, err := config.ETHConn.EthCall(t, "latest")
	if err != nil {
		log.Println("错误信息:", err.Error())
	}
	//單位計算
	ethc, _ := eth.ParseBigInt(balance) //
	intwei, _ := strconv.ParseFloat(ethc.String(), decimals)
	inteth := intwei / math.Pow10(decimals)

	return inteth,err

}

//发送手续费
func SendGas(address string) (string, error) {
	var gas float64
	gas = 0.001
	GAS := int64(gas * 1e18)

	gasFrom := config.EthWalletConfig.ManageAddr
	pwd := config.EthWalletConfig.ManagePawd
	v := eth.T{
		From:  gasFrom,
		To:    address,
		Value: big.NewInt(GAS),
	}

	//l, err := UnLock(GASADDRES, "HJBJK5810929")
	//解锁钱包
	islock, rpcerr := config.ETHConn.EthPerUnLockAccount(gasFrom, pwd)
	if rpcerr!=nil || !islock {
		log.Println("解锁钱包异常ERR：", rpcerr.Error())
		return "", rpcerr
	}

	hash, err := config.ETHConn.EthSendTransaction(v)
	if err != nil {
		log.Println("ERR:", err.Error())
	}

	return hash, err
}
//获取ETH余额
func GetBalance(address string) (float64,error) {

	blance, err := config.ETHConn.EthGetBalance(address, "latest")

	//wei TO  ETH
	ethc, _ := eth.ParseBigInt(blance.String())
	intwei, _ := strconv.ParseInt(ethc.String(), 0, 64)
	inteth := float64(intwei) / 1e18

	return inteth,err
}
//通过txid获取交易信息
func GetScanningBlock(txid string)(tran *eth.Transaction,err error) {
	tran, err = config.ETHConn.Client.EthGetTransactionByHash(txid)
	//if tran!=nil{
	//	log.Println("tran:",tran)
	//}
	return
}
