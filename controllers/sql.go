package controllers

import (
	"github.com/astaxie/beego"

	"cloudWorld/models"
)

type SqlController struct {
	beego.Controller
}

func (c *SqlController) Get() {
	users:=models.Read()
	c.Data["json"] = users

	c.ServeJSON()
	//for _,m:=range users{
	//	fmt.Println(m)
	//}
}


