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

/**
检索所有用户
 */
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
	usertime := main.Input().Get("passwd")
	user, err := models.Register(userphone, userpasswd, username, usersex, useravatar, usertitle, userfriend, usertime)
	if err != nil {
		main.Ctx.WriteString("注册失败:" + err.Error())
	} else {
		b, _ := json.Marshal(user)
		main.Ctx.WriteString("注册成功:" + string(b))
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
		main.Ctx.WriteString("登陆失败:" + err.Error())
	} else {
		b, _ := json.Marshal(user)
		main.Ctx.WriteString("登陆成功:" + string(b))
	}
}

/**
接收上传头像
 */
func (this *MainController) UploadImag() {

	f, h, err := this.GetFile("img")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()
	e := this.SaveToFile("img", "static/upload/"+h.Filename)
	if e != nil {
		this.Ctx.WriteString(e.Error())
	} else {
		this.Ctx.WriteString("OK")
	}

}
