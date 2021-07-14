package dao

import (
	"Rebuild_demo/dao/db"
	"Rebuild_demo/model"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including a T() method which
// returns the current testing context
type OrderTestSuite struct {
	suite.Suite
	orders []*model.DemoOrder
	updateDate []*map[string]interface{}
	dao OrderDAO
}

func (s *OrderTestSuite) SetupSuite(){
	//输出日志
	s.T().Log("SetupSuite")
	//新建数据库使用
	s.dao =NewOrderDAO(db.NewOrm())

	s.orders = []*model.DemoOrder{
		{OrderNo:  "1111111111", UserName: "heyakang1", Amount:   1000, Status:   "false", FileUrl:  "FileUrl"},
		{OrderNo:  "2222222222", UserName: "heyakang12", Amount:   1000, Status:   "false", FileUrl:  "FileUrl"},
		{OrderNo:  "3333333333", UserName: "heyakang123", Amount:   1000, Status:   "false", FileUrl:  "FileUrl"},
		{OrderNo:  "4444444444", UserName: "heyakang1234", Amount:   1000, Status:   "false", FileUrl:  "FileUrl"},
	}

	s.updateDate = []*map[string]interface{}{
		{"amount":111111.1},
		{"amount":222222.2},
		{"amount":333333.3},
		{"amount":444444.4},
	}
}

func (s *OrderTestSuite) SetupTest() {
	s.T().Log("SetupTest")
}

func (s *OrderTestSuite) TearDownSuite() {
	s.T().Log("TearDownSuite")
}

//创建
func (s *OrderTestSuite) Test_orderDAO_Create(){
	s.T().Log("Test_orderDAO_Create")

	for _,v := range s.orders{
		require.NoError(s.T(), s.dao.Create(v))

	}

	//require.NoError(s.T(), s.dao.Create(s.order))
	//require.NoError(s.T(), err, "条件不成立时输出")
	//require.NotNil(s.T(), o, "条件不成立时输出")
}

//通过用户NO查找
func (s *OrderTestSuite) Test_orderDAO_QueryByNo(){
	s.T().Log("Test_orderDAO_QueryByNo")
	for _,v := range s.orders{
		o,err := s.dao.QueryByNo(v.OrderNo)
		s.T().Log(o)
		require.NoError(s.T(), err, "err")
		require.NotNil(s.T(), o, "条件不成立时输出")
	}
}

//通过用户名查找
func (s *OrderTestSuite) Test_orderDAO_QueryListByName(){
	s.T().Log("Test_orderDAO_QueryListByName")
	for _,v := range s.orders{
		o,err := s.dao.QueryListByName(v.UserName,"")
		for _,v :=range o{
			s.T().Log(v)
		}
		require.NoError(s.T(), err, "err")
		require.NotNil(s.T(), o, "条件不成立时输出")
	}
}

//通过用户NO更新数据
func (s *OrderTestSuite) Test_orderDAO_UpdateByNo(){
	s.T().Log("Test_orderDAO_UpdateByNo")
	for i,v := range s.orders{
		err := s.dao.UpdateByNo(v.OrderNo,*s.updateDate[i])
		require.NoError(s.T(), err,"err")
	}
}


//通过用户NO删除用户
func (s *OrderTestSuite) Test_orderDAO_DeleteByNo(){
	s.T().Log("Test_orderDAO_DeleteByNo")
	for _,v := range s.orders{
		err := s.dao.DeleteByNo(v.OrderNo)
		require.NoError(s.T(), err, "err")
	}
}


//查询表格
func (s *OrderTestSuite) Test_orderDAO_QueryTable(){
	s.T().Log("Test_orderDAO_QueryTable")
	o,_ := s.dao.QueryTable()
	for _,v :=range o{
		s.T().Log(v)
	}

	//require.NoError(s.T(), err, err.Error())
	require.NotNil(s.T(), o, "条件不成立时输出")
}



// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestOrderTestSuite(t *testing.T) {
	suite.Run(t, new(OrderTestSuite))
}