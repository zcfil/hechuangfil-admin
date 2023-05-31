package models

type SysProfitConfig struct {
	Id 		string `gorm:"column:id" json:"id"`
	Cvalue	string `gorm:"column:cvalue" json:"cvalue"`
	Cname	string `gorm:"column:cname" json:"cname"`
	Percent	float64 `gorm:"column:percent" json:"percent"`
	Userid	string 	`gorm:"column:userid" json:"userid"`
}
type SysProfitSalesman struct {
	Id 		string `gorm:"column:id" json:"id"`
	Amount	float64 `gorm:"column:amount" json:"amount"`
	Profits	float64 `gorm:"column:profits" json:"profits"`
	Userid	string 	`gorm:"column:userid" json:"userid"`
}