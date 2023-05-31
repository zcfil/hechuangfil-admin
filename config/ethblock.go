package config

import "github.com/spf13/viper"

/**
 * @Author: king bo
 * @Author: kingbo@163.com
 * @Date: 2020/5/20 3:40 下午
 * @Desc: eth 节点链接
 */

//ETH钱包rpc客户端服务器配置
type EthWallet struct {
	ServerHost      string
	ServerPort      int
	UseSSL          bool
	UnlockWalletPwd string
	ManageAddr      string
	ManagePawd      string
}

var EthWalletConfig = new(EthWallet)

func InitEthBlock(cfg *viper.Viper) *EthWallet {
	return &EthWallet{
		ServerHost:      cfg.GetString("serverHost"),
		ServerPort:      cfg.GetInt("serverPort"),
		UseSSL:          cfg.GetBool("useSSL"),
		UnlockWalletPwd: cfg.GetString("unlockManagePwd"),
		ManageAddr:      cfg.GetString("manageAddr"),
		ManagePawd:      cfg.GetString("managePwd"),
	}
}
