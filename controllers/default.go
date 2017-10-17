package controllers

import (
	"github.com/astaxie/beego"
	"WJ/models"
	"github.com/astaxie/beego/orm"
	_"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"log"
)

type MainController struct {
	beego.Controller
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:dyk123@/wj?charset=utf8")
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"

}

func (c *MainController) Post() {

}

func (main *MainController) SelectAllUser() {

	main.TplName = "index.html"
}

/**
注册
 */
func (main *MainController) Register() {

	userphone := main.Input().Get("phone")
	userpasswd := main.Input().Get("passwd")
	username := main.Input().Get("name")
	usersex := main.Input().Get("sex")
	useravatar := main.Input().Get("avatar")
	usertitle := main.Input().Get("title")
	userfriend := main.Input().Get("friend")
	user, err := models.Register(userphone, userpasswd, username, usersex, useravatar, usertitle, userfriend)

	if err != nil {
		responsebean2 := models.ResponseBean{2, "注册失败", err.Error()}
		str, _ := json.Marshal(responsebean2)
		main.Ctx.WriteString(string(str))
	} else {
		responsebean1 := models.ResponseBean{1, "注册成功", user}
		str, _ := json.Marshal(responsebean1)
		main.Ctx.WriteString(string(str))
	}
}

/**
登陆接口
 */
func (main *MainController) Login() {
	userphone := main.Input().Get("phone")
	userpasswd := main.Input().Get("passwd")
	user, err := models.Login(userphone, userpasswd)

	if err != nil {
		responsebean2 := models.ResponseBean{2, "登陆失败", err.Error()}
		str, _ := json.Marshal(responsebean2)
		main.Ctx.WriteString(string(str))
	} else {
		responsebean1 := models.ResponseBean{1, "登陆成功", user}
		str, _ := json.Marshal(responsebean1)
		main.Ctx.WriteString(string(str))
	}
}

/**
接收上传头像
 */
func (this *MainController) UploadImag() {

	f, _, err := this.GetFile("img")
	userphone := this.Input().Get("phone")

	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()

	e := this.SaveToFile("img", "static/upload/"+userphone+".png")
	if e != nil {
		this.Ctx.WriteString(e.Error())
	} else {
		models.UploadUserAvatar(userphone, userphone+".png")
		this.Ctx.WriteString("OK")
	}

}

/**
下载APK
 */
func (this *MainController) DownLoadApk() {
	this.Ctx.Output.Download("static/apk/zhansha_v1.0.2.3.apk")
}
