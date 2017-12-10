package main

import (
	_ "WJ/routers"
	"github.com/astaxie/beego"
	"WJ/models"
)



func main() {
go 	 models.StartLonglink()
	beego.SetStaticPath("/pic","static/upload")
	beego.Run()

}
