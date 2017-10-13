package routers

import (
	"wj/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"get:SelectAllUser")
	beego.Router("/login.dgg/:catId",&controllers.MainController{},"get:Login")
	beego.Router("/register.dgg/:name/:sex/:passwd/:phone",&controllers.MainController{},"get:Register")
	beego.Router("/upload.dgg",&controllers.MainController{},"post:UploadImag")
}
