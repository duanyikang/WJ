package routers

import (
	"WJ/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:SelectAllUser")
	beego.Router("/login.dgg", &controllers.MainController{}, "post:Login")
	beego.Router("/register.dgg", &controllers.MainController{}, "post:Register")
	beego.Router("/upload.dgg", &controllers.MainController{}, "post:UploadImag")
	beego.Router("/uploadrecord.dgg", &controllers.MainController{}, "post:UploadRecord")
	beego.Router("/allrecord.dgg", &controllers.MainController{}, "post:AllRecord")
	beego.Router("/download.dgg", &controllers.MainController{}, "get:DownLoadApk")
}
