package main

import (
	"Rebuild_demo/model"
	"Rebuild_demo/server"
	"fmt"
)

func main(){
	//启动 这是
	//启动 这就是创建好的router
	g := server.NewEngine(server.NewServer(
		&model.DemoOrder{},
	))
	if err := g.Run(":9091"); err != nil {
		fmt.Print(err.Error())
		panic(err)
	}
}

