package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var _db *sql.DB
var err error

func init(){
	println("初始化")
	//配置MySQL连接参数
	username := "root"       //账号
	password := "123456"     //密码
	host := "192.168.10.188" //数据库地址，可以是Ip或者域名
	port := 3307             //数据库端口
	Dbname := "test"         //数据库名
	timeout := "10s"         //连接超时，10秒

	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	_db, err = sql.Open("mysql", dsn)
	if err != nil{
		log.Fatalln(err)
	}

	_db.SetMaxIdleConns(20)
	_db.SetMaxOpenConns(20)

	if err := _db.Ping(); err != nil{
		log.Fatalln(err)
	}

}

func GetDB() *sql.DB{
	return _db
}