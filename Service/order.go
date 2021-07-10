package Service

import (
	"Rebuild_demo/Dao"
	"Rebuild_demo/Model"
)
//服务层结构  （为啥没用*）
type OrderService struct {
	orderDao Dao.OrderDAO
}


//返回Service
func NewOrderService(orderDao Dao.OrderDAO) *OrderService{
	return &OrderService{orderDao: orderDao}
}

////新建目录数据库接口
func (o *OrderService) Create(req *Model.DemoOrder) error{
	return o.orderDao.Create(req)
}

//根据orderNo 更新订单
func (o *OrderService) UpdateByNo(orderNo string, m map[string]interface{}) error{
	return o.orderDao.UpdateByNo(orderNo,m)

}
//根据orderNo查找订单
func (o *OrderService) QueryByNo(orderNo string)(order *Model.DemoOrder,err error){
	return o.orderDao.QueryByNo(orderNo)
}

//根据orderNo删除订单
func (o *OrderService) DeleteByNo(orderNo string)(order *Model.DemoOrder,err error){
	return o.orderDao.DeleteByNo(orderNo)
}
//根据姓名查找列表
func (o *OrderService) QueryListByName(userName string,orderBy string)(orders []Model.DemoOrder,err error){
	//排序没实现
	return o.orderDao.QueryListByName(userName,orderBy)
}
//获取表单
func (o *OrderService) QueryTable()(orders []Model.DemoOrder,err error){
	return o.orderDao.QueryTable()
}
