package database

import (
	"bytes"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
	config2 "hechuangfil-admin/config"
)

var (
	Eloquent *gorm.DB
)

func init() {
	initAdmin()
}

//初始化管理系统数据库链接
func initAdmin() {

	var err error
	conn, dbType := Mysqlconn("admin")
	log.Println("管理系统数据库链接：" + conn)
	var db Database
	if dbType == "mysql" {
		db = new(Mysql)
	} else {
		panic("db type unknow")
	}

	Eloquent, err = db.Open(dbType, conn)
	Eloquent.LogMode(true)
	if err != nil {
		log.Fatalln("mysql admin connect error %v", err)
	} else {
		log.Println("mysql admin connect success!")
	}
	if Eloquent.Error != nil {
		log.Fatalln("database error %v", Eloquent.Error)
	}
	//config2.AdminBeegoOrmJoinMysql() //初始化beego 数据库链接

}


//数据库链接
func Mysqlconn(typesql string) (conns string, dbType string) {
	var host, database, username, password string
	var port int

	switch typesql {
	case "admin":
		dbType = config2.AdminDatabaseConfig.Dbtype
		host = config2.AdminDatabaseConfig.Host
		port = config2.AdminDatabaseConfig.Port
		database = config2.AdminDatabaseConfig.Database
		username = config2.AdminDatabaseConfig.Username
		password = config2.AdminDatabaseConfig.Password
	}

	if dbType != "mysql" {
		fmt.Println("db type unknow")
	}

	var conn bytes.Buffer
	conn.WriteString(username)
	conn.WriteString(":")
	conn.WriteString(password)
	conn.WriteString("@tcp(")
	conn.WriteString(host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(port))
	conn.WriteString(")")
	conn.WriteString("/")
	conn.WriteString(database)
	conn.WriteString("?charset=utf8&parseTime=true&loc=Local&timeout=5s")
	conns = conn.String()
	return
}

type Database interface {
	Open(dbType string, conn string) (db *gorm.DB, err error)
}

type Mysql struct {
}

func (*Mysql) Open(dbType string, conn string) (db *gorm.DB, err error) {
	eloquent, err := gorm.Open(dbType, conn)
	return eloquent, err
}

type SqlLite struct {
}

func (*SqlLite) Open(dbType string, conn string) (db *gorm.DB, err error) {
	eloquent, err := gorm.Open(dbType, conn)
	return eloquent, err
}
