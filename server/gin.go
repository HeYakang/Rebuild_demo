package server

import (
	"Rebuild_demo/server/middleware"
	"Rebuild_demo/utils/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"go/build"
)

//参数是数据库Server
func NewEngine(h *Server) *gin.Engine {
	g := gin.Default()
	//最终版本模式
	gin.SetMode(gin.ReleaseMode)
	logger.InitLogger("",fmt.Sprintf("%s/",build.Default.GOPATH))
    g.Use(middleware.Logger())

	userGroup := g.Group("/userInfo")
	{
		//通过订单orderNo查找
		userGroup.GET("/userInfo",h.QueryOrder())
		//添加订单
		userGroup.POST("/userInfo",h.AddOrder())
		//修改订单信息
		userGroup.PUT("/userInfo",h.UpdateOrder())
		//删除
		userGroup.DELETE("/userInfo", h.Delete())
		//根据姓名模糊查找
		userGroup.GET("/userInfoList", h.QueryByName())
		//下载订单列表为EXCEL
		userGroup.GET("/excel",h.DownloadTableList())

	}
	serviceGroup := g.Group("/service")
	{
		//上传文件，更新URL
		serviceGroup.PUT("/file", h.Upload())
		//下载文件
		serviceGroup.GET("/file", h.Download())
	}

	return g
}
