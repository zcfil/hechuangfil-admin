package mycasbin

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/go-kit/kit/endpoint"
	_ "github.com/go-sql-driver/mysql"
	"hechuangfil-admin/config"
	"hechuangfil-admin/utils"
)

var Em endpoint.Middleware

func init() {
	//Apter ,_ := gormadapter.NewAdapter("mysql", "testdb:Qq123456@tcp(rm-uf60100537o86401eao.mysql.rds.aliyuncs.com:3306)/testdb",true) // Your driver and data source.
	//e,_ := casbin.NewEnforcer("config/rbac_model.conf", "config/policy.csv")
	//
	//err := Apter.SavePolicy(e.GetModel())
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Load the policy from DB.
	//e.LoadPolicy()

}

func Casbin() (*casbin.Enforcer, error) {
	conn := config.AdminDatabaseConfig.Username + ":" + config.AdminDatabaseConfig.Password + "@tcp(" + config.AdminDatabaseConfig.Host + ":" + utils.IntToString(config.AdminDatabaseConfig.Port) + ")/" + config.AdminDatabaseConfig.Database
	Apter, _ := gormadapter.NewAdapter(config.AdminDatabaseConfig.Dbtype, conn, true)
	e, _ := casbin.NewEnforcer("config/rbac_model.conf", Apter)

	if err := e.LoadPolicy(); err == nil {
		return e, err
	} else {
		fmt.Print("casbin rbac_model or policy init error, message: %v", err)
		return nil, err
	}
}
