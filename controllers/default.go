package controllers

import (
	"github.com/astaxie/beego"
	"wj/models"
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
	c.TplName = "default/hello.tpl"

}

func (c *MainController) Post() {

}

/**
检索所有用户
 */
func (main *MainController) SelectAllUser() {
	dbObj := orm.NewOrm()
	dbObj.Using("default")
	var users []models.Userinfo
	orm.NewOrm().QueryTable("userinfo").All(&users)
	b, _ := json.Marshal(users)
	main.Data["Data"] = string(b)
	main.TplName = "default/hello.tpl"
}

/**
注册
 */
func (main *MainController) Register() {
	dbObj := orm.NewOrm()
	dbObj.Using("default")

	user := new(models.Userinfo)
	user.UserName = main.GetString(":name")
	user.UserSex = main.GetString(":sex")
	user.UserPasswd = main.GetString(":passwd")
	user.UserPhone = main.GetString(":phone")
	user.UserFriend = -1

	_, err := dbObj.Insert(user)
	if err != nil {
		main.Data["Data"] = "输入的信息有🈚️问题啊"
	} else {
		main.Data["Data"] = user.Id
	}
	main.TplName = "default/register.tpl"
}

/**
登陆接口
 */
func (main *MainController) Login() {
	dbObj := orm.NewOrm()
	dbObj.Using("default")
	userId, _ := main.GetInt(":catId")

	user := new(models.Userinfo)
	user.Id = userId

	err := dbObj.Read(user)
	if err != nil {
		main.Data["Data"] = "没有这个用户，笨蛋"
	} else {
		b, _ := json.Marshal(user)
		main.Data["Data"] = string(b)
	}

	main.TplName = "default/login.tpl"
}

/**
上传头像
 */
func (main *MainController) UploadImag() {
	f, h, err := main.GetFile("img")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()
	main.SaveToFile("uploadname", "static/upload/"+h.Filename)
}
