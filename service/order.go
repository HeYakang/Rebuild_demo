package service

import (
	"Rebuild_demo/dao"
	"Rebuild_demo/model"
)
//服务层结构  （为啥没用*）
type OrderService struct {
	orderDao dao.OrderDAO
}


//返回Service
func NewOrderService(orderDao dao.OrderDAO) *OrderService{
	return &OrderService{orderDao: orderDao}
}

////新建目录数据库接口
func (o *OrderService) Create(req *model.DemoOrder) error{
	return o.orderDao.Create(req)
}

//根据orderNo 更新订单
func (o *OrderService) UpdateByNo(orderNo string, m map[string]interface{}) error{
	return o.orderDao.UpdateByNo(orderNo,m)

}
//根据orderNo查找订单
func (o *OrderService) QueryByNo(orderNo string)(order *model.DemoOrder,err error){
	return o.orderDao.QueryByNo(orderNo)
}

//根据orderNo删除订单
func (o *OrderService) DeleteByNo(orderNo string)error{
	return o.orderDao.DeleteByNo(orderNo)
}
//根据姓名查找列表
func (o *OrderService) QueryListByName(userName string,orderBy string)(orders []*model.DemoOrder,err error){
	//排序没实现
	return o.orderDao.QueryListByName(userName,orderBy)
}
//获取表单
func (o *OrderService) QueryTable()(orders []*model.DemoOrder,err error){
	return o.orderDao.QueryTable()
}
