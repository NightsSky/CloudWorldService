package routers

import (
	"cloudWorld/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user",&controllers.SqlController{})
}
