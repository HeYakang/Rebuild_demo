package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewOrm() *gorm.DB{
	//配置MySQL连接参数
	username := "root"  //账号
	password := "123"   //密码
	host := "127.0.0.1" //数据库地址，可以是Ip或者域名
	port := 3306        //数据库端口
	Dbname := "test_db" //数据库名

	//通过前面的数据库参数，拼接MYSQL DSN， 其实就是数据库连接串（数据源名称）
	//MYSQL dsn格式： {username}:{password}@tcp({host}:{port})/{Dbname}?charset=utf8&parseTime=True&loc=Local
	//类似{username}使用花括号包着的名字都是需要替换的参数
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	//db,err := gorm.Open("mysql","root:123@(127.0.0.1)/test?charset=utf8mb4&loc=Local")
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	//禁用复数结构
	db.SingularTable(true)

	//获取数据库
	d:= db.DB()

	if err := d.Ping(); err != nil {
		panic(err)
	}
	////如果表不存在则创建

	return db
}

//db.AutoMigrate(&DemoOrder{})