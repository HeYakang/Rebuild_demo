package main

import (
	"Rebuild_demo/Model"
	"Rebuild_demo/Server"
	"fmt"
)

func main(){
	//启动 这是
	//启动 这就是创建好的router
	g := Server.NewEngine(Server.NewServer(
		&Model.DemoOrder{},
	))
	if err := g.Run(":9090"); err != nil {
		fmt.Print(err.Error())
		panic(err)
	}
}

