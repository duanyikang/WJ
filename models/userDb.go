package models

import "github.com/astaxie/beego/orm"

type Userinfo struct {
	Id         int    `pk:"auto"`
	UserName   string `orm:"size(255)"`
	UserSex    string `orm:"size(20)"`
	UserPasswd string `orm:size(50)`
	UserPhone  string `orm:"size(20)"`
	UserTitle  string `orm:size(255)`
	UserFriend int `orm:size(20)`
}

func init() {
	orm.RegisterModel(new(Userinfo))
}
