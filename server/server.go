package server

import (
	"Rebuild_demo/dao"
	"Rebuild_demo/dao/db"
	"Rebuild_demo/model"
	"Rebuild_demo/service"
)

type Server struct {
	service *service.OrderService
}

// 需要问一下... interface
//这边创建Serve并且绑定三个层之间的关系
func NewServer(det ...interface{}) *Server {
	//创建orm对象
  orm := db.NewOrm()
   orm.AutoMigrate(&model.DemoOrder{})
  //if err != nil{
  //	//需要了解一下panic的功能
  //	fmt.Print(err.Error)
  //	panic(err)
  //}

  //绑定ORM与DAO层的关系
  orderDao := dao.NewOrderDAO(orm)

  //绑定DAO层与Service层的关系

  orderService := service.NewOrderService(orderDao)

	return &Server{
		service : orderService,
	}
}
