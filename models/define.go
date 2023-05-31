package models

const (
	PASS_VERIFICATION = 1  			// 通过审核
	NO_PASS_VERIFICATION = 0  		// 没通过审核
)

var AUDIT = map[string]bool{
	"admin":true,
	"finance":true,
	"boss":true,
}

//configKey 配置key
const (
	FROM_ADDRESS = "from_address"	//转出钱包配置
	COLLECTION_ADDRESS = "collection_address"	//归集钱包
)
const (
	VERIFICATION_STR = "0" 			// 正在审核
	PASS_VERIFICATION_STR = "1"  			// 通过审核
	NO_PASS_VERIFICATION_STR = "2" 		// 没通过审核
)


//var FilAddress = make(map[string]string)
//var RechargeRecord = make(map[string]bool)