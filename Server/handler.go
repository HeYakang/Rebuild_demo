package Server

import (
	"Rebuild_demo/Model"
	serverModel "Rebuild_demo/Server/server_model"
	"Rebuild_demo/Util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"net/http"
	"os"
	"time"
)

//按照用户订单号码进行查询
func (s *Server) QueryOrder() gin.HandlerFunc{
	return func(c *gin.Context) {
		OrderNo := c.Query("order_no")
        if len(OrderNo) == 0{
        	s.httpErr(c,http.StatusBadRequest,"Not Found OrderNo")
			return
		}

		o,err := s.service.QueryByNo(OrderNo)
		if err != nil{
			s.httpErr(c,http.StatusBadGateway,err.Error())
			return
		}

		//查询成功返回结果
		s.httpSuccess(c,&serverModel.Order{
			ID: o.ID,
			OrderNo: o.OrderNo,
			UserName: o.UserName,
			Amount: o.Amount,
			FileUrl: o.FileUrl,
		})
	}
}

//按名称模糊查询
func (s *Server) QueryByName() gin.HandlerFunc{
	return func(c *gin.Context) {
		name := c.Query("name")
		if len(name) == 0{
			s.httpErr(c,http.StatusBadRequest,"Not Found name")
			return
		}

		o,err := s.service.QueryListByName(name,"")
		if err != nil{
			s.httpErr(c,http.StatusBadGateway,err.Error())
			return
		}
		//解析结果
		////var result string
		//var result map[string]string
		//result = make(map[string]string)
		////格式化输出
		//for k, v := range o {
		//	//result +=fmt.Sprint("%d:%s",k,v)
		//	result[fmt.Sprintf("%d", k)] = fmt.Sprintf("%s", v)
		//
		//}
		//查询成功返回结果
		s.httpSuccess(c,o)
	}
}

//添加记录
func (s *Server) AddOrder() gin.HandlerFunc{
	return func(c *gin.Context) {
		var req *serverModel.AddOrderReq
		if err := c.ShouldBind(&req); err !=nil{
			s.httpErr(c,http.StatusBadRequest,err.Error())
			return
		}

		if err := req.IsValid(); err != nil{
			s.httpErr(c,http.StatusBadRequest,err.Error())
			return
		}

		o := Model.DemoOrder{
			OrderNo:  fmt.Sprintf("%s%d", time.Now().Format(Util.FormatDateTime), time.Now().UnixNano()),
			UserName: req.UserName,
			Amount:   req.Amount,
			Status:   "true",
			FileUrl:  req.FileUrl,
		}

		if err := s.service.Create(&o); err != nil{
			s.httpErr(c,http.StatusBadGateway,err.Error())
			return
		}

		//添加成功返回结果
		s.httpSuccess(c,&serverModel.Order{
			ID:       o.ID,
			OrderNo:  o.OrderNo,
			UserName: o.UserName,
			Amount:   o.Amount,
			FileUrl:  o.FileUrl,
		})
	}
}


//修改
func (s *Server) UpdateOrder() gin.HandlerFunc{
	return func(c *gin.Context) {
		var req *serverModel.Order
		if err := c.ShouldBind(&req); err != nil{
			s.httpErr(c,http.StatusBadRequest,err.Error())
			return
		}

		//上传数据不合法
		if err := req.IsValid(); err != nil{
			s.httpErr(c, http.StatusBadRequest,err.Error())
			return
		}

		//更新金额
		if err := s.service.UpdateByNo(req.OrderNo, map[string]interface{}{
			"amount":req.Amount,
		});err!=nil{
			s.httpErr(c,http.StatusBadGateway,err.Error())
			return
		}

		//更新成功
		s.httpSuccess(c,nil)
	}
}
//删除
func (s *Server) Delete() gin.HandlerFunc{
	return func(c *gin.Context) {
		var req *serverModel.Order
		if err := c.ShouldBind(&req); err != nil{
			s.httpErr(c,http.StatusBadRequest,err.Error())
			return
		}

		//上传数据不合法
		if err := req.IsDeleteValid(); err != nil{
			s.httpErr(c, http.StatusBadRequest,err.Error())
			return
		}

		//删除订单
		if _,err := s.service.DeleteByNo(req.OrderNo); err!=nil{
			s.httpErr(c,http.StatusBadGateway,err.Error())
			return
		}

		//删除成功
		s.httpSuccess(c,nil)
	}
}


//上传文件
func (s *Server) Upload() gin.HandlerFunc{
	return func(c *gin.Context) {
		orderNo := c.Query("order_no")
		if len(orderNo) == 0 {
			s.httpErr(c, http.StatusBadRequest, "Not found orderNo.")
			return
		}

		//获取文件
		file, err := c.FormFile("file")
		if err != nil {
			s.httpErr(c, http.StatusBadRequest, err.Error())
			return
		}

		//获得文件地址
		path := Util.FilePath(file.Filename)
		//保存文件
		if err = c.SaveUploadedFile(file, path); err != nil {
			s.httpErr(c, http.StatusBadGateway, err.Error())
			return
		}

		//上传数据库
		if err = s.service.UpdateByNo(orderNo, map[string]interface{}{
			"file_url":path,
		}); err != nil {
			s.httpErr(c, http.StatusBadGateway, err.Error())
			return
		}

		//返回成功，返回文件地址
		s.httpSuccess(c, path)
	}
}

//下载文件
func (s *Server) Download() gin.HandlerFunc{
	return func(c *gin.Context) {
		//获取订单号 验证是否符合
		orderNo := c.Query("order_no")
		if len(orderNo) == 0 {
			s.httpErr(c, http.StatusBadRequest, "Not found orderNo.")
			return
		}

		//获取URL地址
		order,err := s.service.QueryByNo(orderNo);
		if  err != nil {
			s.httpErr(c, http.StatusBadGateway, err.Error())
			return
		}

		//获取URL检验文件是否存在
		if _, err := os.Stat(order.FileUrl); os.IsNotExist(err) {
			s.httpErr(c, http.StatusBadGateway, err.Error())
			return
		}

		//下载文件
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", order.FileUrl)) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
		c.Writer.Header().Add("Content-Type", "application/octet-stream")
		c.File(order.FileUrl)


	}
}



//下载表单
func (s *Server) DownloadTableList() gin.HandlerFunc{
	return func(c *gin.Context) {
		//表头
		titleList := []string{"id", "created_at", "updated_at", "deleted_at", "order_no", "user_name", "amount", "status", "file_url"}
		//获取所有数据
		orders,err := s.service.QueryTable()
		if  err != nil {
			s.httpErr(c, http.StatusBadGateway, err.Error())
			return
		}

		//生成一个新文件
		file := xlsx.NewFile()
		//添加sheet页
		sheet, _ := file.AddSheet("demo_order")
		//插入表头
		titleRow := sheet.AddRow()
		for _, v := range titleList {
			cell := titleRow.AddCell()
			cell.Value = v
			//表头字体颜色
			cell.GetStyle().Font.Color = "00FF0000"
			//居中显示
			cell.GetStyle().Alignment.Horizontal = "center"
			cell.GetStyle().Alignment.Vertical = "center"
		}
		// 插入内容
		for _, v := range orders {
			row := sheet.AddRow()

			cell := row.AddCell()
			cell.Value = fmt.Sprintf("%d", v.ID)
			cell = row.AddCell()
			cell.Value = v.CreatedAt.String()
			cell = row.AddCell()
			cell.Value = v.UpdatedAt.String()
			cell = row.AddCell()
			cell.Value = v.DeletedAt.Time.String()
			row.WriteStruct(&v, -1)
		}

		c.Writer.Header().Set("Content-Type", "application/octet-stream")
		disposition := fmt.Sprintf("attachment; filename=\"%s-%d.xlsx\"", "odmo_order", time.Now().UnixNano())
		c.Writer.Header().Set("Content-Disposition", disposition)
		_ = file.Write(c.Writer)

	}
}





func (s *Server) httpErr(c *gin.Context,code int, err string){
	c.JSON(code, Util.NewHttpErrResp(err))
}

func (s *Server) httpSuccess(c *gin.Context,data interface{}){
	c.JSON(http.StatusOK,Util.NewHttpSuccessResp(data))
}
