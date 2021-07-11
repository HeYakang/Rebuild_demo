package Dao


import (
	"Rebuild_demo/Model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//创建接口
type OrderDAO interface {
	//新建目录数据库接口
	Create(req *Model.DemoOrder) error
	//根据orderNo 更新订单
	UpdateByNo(orderNo string, m map[string]interface{}) error
	//根据orderNo查找订单
	QueryByNo(orderNo string)(order *Model.DemoOrder,err error)
	//根据orderNo删除订单
	DeleteByNo(orderNo string)(order *Model.DemoOrder,err error)
	//根据姓名查找列表
	QueryListByName(userName string,orderBy string)(orders []Model.DemoOrder,err error)
	//获取表单
	QueryTable()(orders []Model.DemoOrder,err error)
}



type  orderDAO struct {
	db *gorm.DB
}

//返回DAO
func NewOrderDAO(db *gorm.DB) *orderDAO{
	return &orderDAO{db: db}
}

////新建目录数据库接口
func (o *orderDAO) Create(req *Model.DemoOrder) error{
	return o.db.Create(req).Error
}

//根据orderNo 更新订单
func (o *orderDAO) UpdateByNo(orderNo string, m map[string]interface{}) error{
	//这边如果不指定find（）First等，就算没找到也不会报错
	err := o.db.Model(&Model.DemoOrder{}).Where("order_no = ?",orderNo).First(&Model.DemoOrder{}).Update(m).Error

	return err

}
//根据orderNo查找订单
func (o *orderDAO) QueryByNo(orderNo string)(order *Model.DemoOrder,err error){
	result := Model.DemoOrder{}
	err = o.db.Model(&Model.DemoOrder{}).Where("order_no=?",orderNo).First(&result).Error
	order = &result
	return
}

//根据orderNo删除订单
func (o *orderDAO) DeleteByNo(orderNo string)(order *Model.DemoOrder,err error){
	result := Model.DemoOrder{}
	//这边如果不制定First等限定则会删除对应orderNo下所有的与参数相同的数据，所有需要小心使用，根据需求使用
	err = o.db.Model(&Model.DemoOrder{}).Where("order_no=?",orderNo).First(result).Unscoped().Delete(&result).Error
	order = &result
	return
}
//根据姓名查找列表
func (o *orderDAO) QueryListByName(userName string,orderBy string)(orders []Model.DemoOrder,err error){
	//排序没实现
	result := []Model.DemoOrder{}
	err = o.db.Model(&Model.DemoOrder{}).Where("user_name LIKE ?", "%"+userName+"%").Find(&result).Error
	orders = result
	return
}
//获取表单
func (o *orderDAO) QueryTable()(orders []Model.DemoOrder,err error){
	result := []Model.DemoOrder{}
	err = o.db.Model(&Model.DemoOrder{}).Find(&result).Error
	orders = result
	return
}