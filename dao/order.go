package dao


import (
	"Rebuild_demo/model"
	_"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//创建接口
type OrderDAO interface {
	//新建目录数据库接口
	Create(req *model.DemoOrder) error
	//根据orderNo 更新订单
	UpdateByNo(orderNo string, m map[string]interface{}) error
	//根据orderNo查找订单
	QueryByNo(orderNo string)(order *model.DemoOrder,err error)
	//根据orderNo删除订单
	DeleteByNo(orderNo string)error
	//根据姓名查找列表
	QueryListByName(userName string,orderBy string)(orders []*model.DemoOrder,err error)
	//获取表单
	QueryTable()(orders []*model.DemoOrder,err error)
}



type  orderDAO struct {
	db *gorm.DB
}

//返回DAO
func NewOrderDAO(db *gorm.DB) *orderDAO{
	return &orderDAO{db: db}
}

////新建目录数据库接口
func (o *orderDAO) Create(req *model.DemoOrder) error{
	return o.db.Create(req).Error
}

//根据orderNo 更新订单
func (o *orderDAO) UpdateByNo(orderNo string, m map[string]interface{}) error{
	//这边如果不指定find（）First等，就算没找到也不会报错
	//err := o.db.Model(&model.DemoOrder{}).Where("order_no = ?",orderNo).First(&model.DemoOrder{}).Update(m).Error
	return o.db.Model(&model.DemoOrder{}).Where("order_no = ?",orderNo).First(&model.DemoOrder{}).Updates(m).Error

}
//根据orderNo查找订单
func (o *orderDAO) QueryByNo(orderNo string)(order *model.DemoOrder,err error){
	//result := model.DemoOrder{}
	err = o.db.Model(&model.DemoOrder{}).Where("order_no=?",orderNo).First(&order).Error
	//order = &result
	return
}

//根据orderNo删除订单
func (o *orderDAO) DeleteByNo(orderNo string) error{
	//result := model.DemoOrder{}
	////这边如果不制定First等限定则会删除对应orderNo下所有的与参数相同的数据，所有需要小心使用，根据需求使用
	//err = o.db.Model(&model.DemoOrder{}).Where("order_no=?",orderNo).First(order).Unscoped().Delete().Error
	//order = &result
	return o.db.Model(&model.DemoOrder{}).Where("order_no=?",orderNo).First(&model.DemoOrder{}).Unscoped().Delete(&model.DemoOrder{}).Error
}
//根据姓名查找列表
func (o *orderDAO) QueryListByName(userName string,orderBy string)(orders []*model.DemoOrder,err error){
	//排序没实现
	//result := []model.DemoOrder{}
	err = o.db.Model(&model.DemoOrder{}).Where("user_name LIKE ?", "%"+userName+"%").Find(&orders).Error
	//orders = result
	return
}
//获取表单
func (o *orderDAO) QueryTable()(orders []*model.DemoOrder,err error){
	//result := []model.DemoOrder{}
	err = o.db.Model(&model.DemoOrder{}).Find(&orders).Error
	//orders = result
	return
}