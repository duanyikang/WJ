package main

import (
	_ "WJ/routers"
	"github.com/astaxie/beego"
)



func main() {
	beego.SetStaticPath("/pic","static/upload")
	beego.Run()
}
