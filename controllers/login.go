package controllers

import (
	"github.com/astaxie/beego"
	"cloudWorld/models"
	"encoding/json"
)

type LoginController struct {
	beego.Controller
}



func (c *LoginController) Post() {
	var request models.User//这是一个model，struct类型
	body := c.Ctx.Input.RequestBody//这是获取到的json二进制数据
	json.Unmarshal(body, &request)//解析二进制json，把结果放进ob中
	//user := &User{Id: ob.UserName, Mobile: ob.Mobile}
	user := models.CheckLogin(request.Username,request.Password)
	if user["code"]==200{
		c.SetSession("token","useraflsgllsglslgksldg")
	}
	c.Data["json"] = user
	c.ServeJSON()
}