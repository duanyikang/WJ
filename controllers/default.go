package controllers

import (
	"github.com/astaxie/beego"
	"WJ/models"
	"github.com/astaxie/beego/orm"
	_ "encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"log"
	"time"
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

/**
翻译
 */
func (main *MainController) Translation() {
	key := main.GetString("key")
	if (len(key) < 2) {
		main.TplName = "index.html"
		return
	}

	str := models.Search(key)

	main.Data["s"] = str
	main.TplName = "index.html"
}

func (main *MainController) Chat() {

	key := main.GetString("key")
	if (len(key) < 2) {
		go models.ClientStart()
		main.TplName = "chat.html"
		return
	}
	go models.ClientSendmsg(key)
	main.Data["s"] = "可以啦？"
	main.TplName = "chat.html"
}

func (main *MainController) ChatSend() {
	key := main.GetString("key")
	if (len(key) < 2) {
		main.TplName = "index.html"
		return
	}

	go models.ClientSendmsg(key)
	main.Data["s"] = key
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

	e := this.SaveToFile("img", "static/upload/"+userphone+"avatar.png")
	if e != nil {
		this.Ctx.WriteString(e.Error())
	} else {
		models.UploadUserAvatar(userphone, userphone+"avatar.png")
		this.Ctx.WriteString("OK")
	}

}

func (this *MainController) AllRecord() {
	userphone := this.Input().Get("phone")

	records, err := models.SelectRecord(userphone)
	if err != nil {

	} else {
		responsebean1 := models.ResponseBean{1, "上传成功", records}
		str, _ := json.Marshal(responsebean1)
		this.Ctx.WriteString(string(str))
	}
}

/**
上传日记
 */
func (this *MainController) UploadRecord() {

	f, _, err := this.GetFile("img")
	userphone := this.Input().Get("phone")
	title := this.Input().Get("title")

	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()

	t := time.Now().Unix()
	fileName := time.Unix(t, 0).Format("2006_01_02 15:04:05")

	e := this.SaveToFile("img", "static/upload/"+userphone+fileName+".png")
	if e != nil {
		this.Ctx.WriteString(e.Error())
	} else {
		recordBean, _ := models.InsertRecord(userphone, title, userphone+fileName+".png")
		responsebean1 := models.ResponseBean{1, "上传成功", recordBean}
		str, _ := json.Marshal(responsebean1)
		this.Ctx.WriteString(string(str))
	}

}

/**
下载APK
 */
func (this *MainController) DownLoadApk() {
	this.Ctx.Output.Download("static/apk/zhansha_v1.0.2.3.apk")
}
