/***************************************************
 ** @Desc :  erc20 智能合约各方法对应的签名编码
			transfer(address,uint256)： 0xa9059cbb
			balanceOf(address)：0x70a08231
			decimals()：0x313ce567
			allowance(address,address)： 0xdd62ed3e
			symbol()：0x95d89b41
			totalSupply()：0x18160ddd
			name()：0x06fdde03
			approve(address,uint256)：0x095ea7b3
			transferFrom(address,address,uint256)： 0x23b872dd
 ** @Time : 2019/12/28
 ** @Author : Administrator
 ** @File : signcode.go
 ** @Last Modified by : Administrator
 ** @Last Modified time: 2019-12-28-16:42
 ** @Software: GoLand
****************************************************/
package eth

const (
	TRANSFER_CODE      = "0xa9059cbb" //转账签名编码
	BALANCE_OF_CODE    = "0x70a08231" //余额查询编码
	DECIMALS_CODE      = "0x313ce567" //位數
	ALLOWANCE_CODE     = "0xdd62ed3e" //合约拥有者
	SYMBOL_CODE        = "0x95d89b41" //合约简称
	TOTALSUPPLY_CODE   = "0x18160ddd" //发行总量
	NAME_CODE          = "0x06fdde03" //合约名称
	APPROVE_CODE       = "0x095ea7b3" //认证
	TRANSFER_FROM_CODE = "0x23b872dd" //交易
)
