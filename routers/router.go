package routers

import (
	"wj/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"get:SelectAllUser")
	beego.Router("/login.dgg",&controllers.MainController{},"post:Login")
	beego.Router("/register.dgg",&controllers.MainController{},"post:Register")
	beego.Router("/upload.dgg",&controllers.MainController{},"post:UploadImag")
}

