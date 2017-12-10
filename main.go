package main

import (
	_ "WJ/routers"
	"github.com/astaxie/beego"
	"WJ/models"
)



func main() {
	beego.SetStaticPath("/pic","static/upload")
	beego.Run()
	models.StartLonglink()
}
