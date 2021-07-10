package Model

import "gorm.io/gorm"

//定义数据结构，之后利用这个结构体创建数据库表
type DemoOrder struct{
	gorm.Model   //内嵌4个字段
	//Id int `gorm:"column:id"`
	OrderNo string `gorm:"column:order_no;not null;type:varchar(120)"`//订单号
	UserName string `gorm:"column:user_name;not null;type:varchar(120)"`//用户名
	Amount float64 `gorm:"column:amount;not null;default:0;type:float"`//金额
	Status string `gorm:"column:status;not null;type:varchar(120)"`//状态
	FileUrl string `gorm:"column:file_url"index:"addr"default:"null"type:"varchar(120)"`//文件地址
}

