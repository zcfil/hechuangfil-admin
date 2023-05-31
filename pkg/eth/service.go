/***************************************************
 ** @Desc :
 ** @Time : 2019/12/27
 ** @Author : Administrator
 ** @File : service.go
 ** @Last Modified by : Administrator
 ** @Last Modified time: 2019-12-27-19:36
 ** @Software: GoLand
****************************************************/
package eth

import (
	"time"
)

//var client = NewEthRPC("http://47.113.108.103:11002")

//解锁账户
//func UnLock(address string, password string) (bool, error) {
//
//	addr := strings.ToLower(address)
//	l, err := client.EthPerUnLockAccount(addr, password)
//	if err != nil {
//		log.Println("ERR:", err.Error())
//	}
//	return l, err
//}

//USDT 归集
func USDTCollect() {
	//list := ListAccount() //账户列表
	/*	for i := 1; i < len(list); i++ {
			if list[i].Address != GASADDRES {
				//fmt.Println(i," 地址：",list[i].Address,"	ETH 余额：",list[i].ETH, "  USDT 余额：",list[i].USDT)
				//手续费不足发送0.001手续到有USDT的账户里面
				if list[i].ETH < 0.0001 && list[i].USDT >= 1 {
					txhash, _ := sendGas(list[i].Address, 0.0005)
					log.Println("GAS交易Hash：", txhash)
					//等待1分钟
					//time.Sleep(1*60*time.Second)
					////判断手续费是否确认
					//tr,err:=client.EthGetTransactionReceipt(txhash)
					//if err !=nil {
					//	log.Println("ERR",err.Error())
					//	return
					//}
					//
					//s,_ :=strconv.Atoi(strings.Split(tr.Status,"0x")[1])
					//if s==1{
					//	//解锁
					//	l,err:=	UnLock(list[i].Address,PWD)
					//	if err !=nil || l ==false{
					//		log.Println("解锁失败，错误提示：",err.Error())
					//		return
					//	}
					//	//开始收集USDT
					//	thash,_:=tokenTraction(list[i].Address,list[i].USDT)
					//	fmt.Println("归集交易hash:",thash)
					//}
				}
			}
		}
	*/
	//等待1分钟
	time.Sleep(60 * time.Second)
	//开始归集
	//判断手续费是否确认
	//tr,err:=client.EthGetTransactionReceipt(txhash)
	//if err !=nil {
	//	log.Println("ERR",err.Error())
	//	return
	//}

	//s,_ :=strconv.Atoi(strings.Split(tr.Status,"0x")[1])
	//if s==1{
	//}

	/*	for i := 1; i < len(list); i++ {
		if list[i].Address != GASADDRES {
			if list[i].USDT >= 1 {
				//解锁
				l, err := UnLock(list[i].Address, PWD)
				fmt.Println("归集账户：", list[i].Address)
				if err != nil || l == false {
					log.Println("解锁失败，错误提示：", err.Error())
					return
				}
				//开始收集USDT
				thash, _ := Traction(list[i].Address, list[i].USDT)
				fmt.Println("归集交易hash:", thash)
			}
		}
	}*/

	//time.Sleep(1*60*time.Second)
	//newlist:=listAccount()  //账户列表
	//for i:=1;i< len(newlist); i++ {
	//	fmt.Println(i," 地址：",newlist[i].Address,"	ETH 余额：",newlist[i].ETH, "  USDT 余额：",newlist[i].USDT)
	//
	//}
}

//Token 转帐
//func Traction(toaddress string, value float64) (string, error) {
//
//	//taddr := strings.Split(toaddress, "0x")[1]
//	//
//	////代币数量 十六进制的value值去掉0x并由0补够64位数
//	////v:=strconv.FormatFloat(value,0.00000,6,64)
//	//vl := IntToHex(int(value * 1e6)) //"000000000000000000000000"
//	//vs := strings.Split(vl, "0x")[1]
//	//
//	////fmt.Println("value :",vstr ," len:",len(vstr))
//	////data拼接： “0x”+"23b872dd"+"from地址去掉0x并由0补够64位数"+"to地址去掉0x并由0补够64位数"+"十六进制的value值去掉0x并由0补够64位数"
//	////data:="0x70a08231"+faddr+taddr+vstr //data拼接
//	//data := "0xa9059cbb" + addPreZero(taddr) + addPreZero(vs) //data拼接 addPreZero(faddr)
//	//
//	//fmt.Println("Data拼装：", data)
//	//t := T{
//	//	From: Fromadd,
//	//	To:   "NBH", //合约地址
//	//	//Value:
//	//	Data: data,
//	//}
//	////hash,err:= client.EthPerSendTransaction(t,PWD)
//	//
//	//hash, err := client.EthSendTransaction(t)
//	//
//	//if err != nil {
//	//	log.Println("错误信息:", err.Error())
//	//}
//	//fmt.Println("交易Hash",hash)
//
//	return hash, err
//}

////创建钱包地址
//func NewAccount(password string) {
//
//	//for i := 0; i <= 1000; i++ {
//	//	account, _ := client.EthPerNewAccount("luce!1989@9922sdf")
//	//	fmt.Println(nu, " 数量：", i, "-", account)
//	//	time.Sleep(300)
//	//}
//	account, _ := client.EthPerNewAccount(password)
//	log.Println(account)
//}

////列出本地钱包
//func ListAccount() []string {
//
//	list, err := client.EthPerListAccounts()
//	if err != nil {
//		log.Println("列出地址失败：", err.Error())
//	}
//
//	/*account := Account{}
//	listAccounts := []Account{}
//	//fmt.Println("本地钱包数量：",len(list))
//	for i := 0; i < len(list); i++ {
//		usdt := getUSDTBalance(list[i]) //USDT 余额
//		blance := getBalance(list[i])   //获取余额
//		if blance > 0 || usdt > 0 {
//			account.Address = list[i]
//			account.USDT = usdt
//			account.ETH = blance
//			listAccounts = append(listAccounts, account)
//		}
//	}*/
//	return list
//}

//获取USDT余额
//func getUSDTBalance(address string) float64 {
//
//	//0x7260c1661793170694344bC813BE6857ED16e58c
//	addr2 := strings.Split(address, "0x")[1]
//	data := "0x70a08231000000000000000000000000" + addr2 //data拼接
//
//	t := T{
//		From: address,
//		To:   "USDT", //USDT合约地址
//		Data: data,
//	}
//
//	balance, err := client.EthCall(t, "latest")
//	if err != nil {
//		log.Println("错误信息:", err.Error())
//	}
//
//	//wei TO  ETH
//	ethc, _ := ParseBigInt(balance)
//	intwei, _ := strconv.ParseFloat(ethc.String(), 6)
//	inteth := float64(intwei) / 1000000
//
//	return inteth
//	//fmt.Println("USDT余额：",ethc.String(),"    ",inteth, "USDT")
//
//	/**
//	acounts,_:=client.EthAccounts()
//	for i:=0;i<len(acounts);i++  {
//		a:=strings.Split(acounts[i],"0x")[1]
//		data:="0x70a08231000000000000000000000000"+a
//		fmt.Println(  "NU:",i," DATA：",data)
//
//		t:=eth.T{
//			//From:acounts[i],
//			To:"0xdac17f958d2ee523a2206206994597c13d831ec7", //USDT合约地址
//			Data:data,
//		}
//
//		blance,_:=client.EthCall(t,"latest")
//		//wei TO  ETH
//		ethc,_ :=eth.ParseBigInt(blance)
//		intwei,_:=eth.ParseInt(ethc.String())
//		inteth:=float64(intwei)/1e6
//
//		if inteth>0 {
//			fmt.Println("地址：",acounts[i],"   USDT余额：",inteth,"ETH")
//		}
//
//	}*/
//}

//获取合约余额
//func TokeBalance(address string, token string) float64 {
//
//	//0x7260c1661793170694344bC813BE6857ED16e58c
//	addr2 := strings.Split(address, "0x")[1]
//	data := "0x70a08231000000000000000000000000" + addr2 //data拼接
//
//	t := T{
//		From: address,
//		To:   token, //合约地址
//		Data: data,
//	}
//
//	balance, err := client.EthCall(t, "latest")
//	if err != nil {
//		log.Println("错误信息:", err.Error())
//	}
//
//	//wei TO  ETH
//	ethc, _ := ParseBigInt(balance)
//	intwei, _ := strconv.ParseFloat(ethc.String(), 6)
//	inteth := float64(intwei) / 1000000
//
//	return inteth
//	//fmt.Println("USDT余额：",ethc.String(),"    ",inteth, "USDT")
//
//	/**
//	acounts,_:=client.EthAccounts()
//	for i:=0;i<len(acounts);i++  {
//		a:=strings.Split(acounts[i],"0x")[1]
//		data:="0x70a08231000000000000000000000000"+a
//		fmt.Println(  "NU:",i," DATA：",data)
//
//		t:=eth.T{
//			//From:acounts[i],
//			To:"0xdac17f958d2ee523a2206206994597c13d831ec7", //USDT合约地址
//			Data:data,
//		}
//
//		blance,_:=client.EthCall(t,"latest")
//		//wei TO  ETH
//		ethc,_ :=eth.ParseBigInt(blance)
//		intwei,_:=eth.ParseInt(ethc.String())
//		inteth:=float64(intwei)/1e6
//
//		if inteth>0 {
//			fmt.Println("地址：",acounts[i],"   USDT余额：",inteth,"ETH")
//		}
//
//	}*/
//}

//获取GasPrice
//func GasPrice() (big.Int, error) {
//	return client.EthGasPrice()
//}

//客户端的coinbase地址
//func CoinsBase() (string, error) {
//	return client.EthCoinbase()
//}

//获取合約信息
//func GetContractInfo(address string, token string, code string) string {
//	if address == "" {
//		log.Println("地址不合法")
//		return ""
//	}
//	addrSplit := strings.Split(address, "0x")[1] //地址去掉0x
//	//data数据格式：最前边的“0x70a08231000000000000000000000000”是固定的，后边的是钱包地址（不带“0x”前缀）
//	//data := "0x70a08231000000000000000000000000" + addrSplit //data拼接
//	data := code + "000000000000000000000000" + addrSplit
//	t := T{
//		From: address, //查詢地址
//		To:   token,   //合约地址
//		Data: data,    //data
//	}
//
//	var dataStr string
//	switch code {
//	case "0x06fdde03": //获取合约名称
//		name, _ := client.EthCall(t, "latest")
//		//單位計算
//		ethc, _ := ParseBigInt(name) //
//		//intwei, _ := strconv.ParseFloat(ethc.String(), 18)
//		n, _ := ethc.GobEncode()
//		//fmt.Printf("合约名称:%s\n",n)
//		dataStr = fmt.Sprintf("%s", n)
//
//	case "0x95d89b41": //合约简称
//		symbol, _ := client.EthCall(t, "latest")
//		//單位計算
//		ethc, _ := ParseBigInt(symbol) //
//		//intwei, _ := strconv.ParseFloat(ethc.String(), 18)
//		s, _ := ethc.GobEncode()
//		//fmt.Printf("合约简称:%s\n", s)
//		dataStr = fmt.Sprintf("%s", s)
//
//	case "0x70a08231": //查询余额
//		//获取代币的余额，要通过rpc接口得到接口为：eth_call
//		balance, _ := client.EthCall(t, "latest")
//		//單位計算
//		ethc, _ := ParseBigInt(balance)                    //
//		intwei, _ := strconv.ParseFloat(ethc.String(), 18) //18
//
//		dataStr, _ = conver.String(intwei) //string(inteth)
//
//	case "0x313ce567": //合约精度
//		decimals, err := client.EthCall(t, "latest")
//		if err != nil {
//			log.Println("错误信息:", err.Error())
//		}
//		bigs, _ := ParseBigInt(decimals)
//
//		pdecimals, _ := strconv.ParseFloat(bigs.String(), 64)
//
//		dataStr, _ = conver.String(pdecimals)
//		//fmt.Println("位數:", pdecimals)
//
//	case "0x18160ddd": //发行总量
//		totalSupply, err := client.EthCall(t, "latest")
//		if err != nil {
//			log.Println("错误信息: TotalSupply", err.Error())
//		}
//		//單位計算
//		ethc, _ := ParseBigInt(totalSupply) //
//		intwei, _ := strconv.ParseFloat(ethc.String(), 18)
//		inteth := intwei / math.Pow10(18)
//		//itotal,_:=ParseInt(totalSupply)
//		ptotalSupply, _ := strconv.ParseFloat(totalSupply, 18)
//		fmt.Println("发行总量:", ptotalSupply)
//		dataStr, _ = conver.String(inteth) //string(inteth)
//	}
//
//	return dataStr
//}

//获取最新区块高度
//func GetEthBlockNumber() (int, error) {
//	number, err := client.EthBlockNumber()
//	return number, err
//}

//根据区块ID获取区块信息
//func GetBlockByNumber(number int) (*Block, error) {
//	blokc, err := client.EthGetBlockByNumber(number, false)
//	return blokc, err
//}
