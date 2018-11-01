package main

import (
	_ "cloudWorld/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"cloudWorld/models"
)

func main() {
	models.RegisterDB()
	// 开启 ORM 调试模式
	orm.Debug = true
	// 自动建表
	//orm.RunSyncdb("web", false, true)
	beego.Run()
	models.Read()
}

