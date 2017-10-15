package models

import (
	"github.com/astaxie/beego/orm"
	"errors"
)

type Userinfo struct {
	Id         int    `pk:"auto"`
	UserPhone  string   `orm:"size(255)"`
	UserPasswd string `orm:"size(255)"`
	UserName   string `orm:"size(255)"`
	UserSex    string `orm:"size(20)"`
	UserAvatar string `orm:"size(255)"`
	UserTitle  string `orm:size(255)`
	UserFriend string `orm:size(20)`
	UserTime   string `orm:"size(255)"`
}

func init() {
	orm.RegisterModel(new(Userinfo))
}

/**
登陆方法
 */
func Login(userPhone, userPasswd string) (user1 Userinfo, err error) {

	dbObj := orm.NewOrm()
	dbObj.Using("default")

	err = dbObj.QueryTable("userinfo").Filter("user_phone", userPhone).Limit(1).One(&user1)

	if err != nil {
		err = errors.New("没有该用户")
		return user1, err
	}

	if (user1.UserPasswd == userPasswd) {
		return user1, nil
	} else {
		err = errors.New("您输入的密码不正确")
		return user1, err;
	}
}

/**
注册
 */
func Register(userPhone, userPassswd, userName, userSex, userAvatar, userTitle, userFriend, userTime string) (user Userinfo, err error) {
	dbObj := orm.NewOrm()
	dbObj.Using("default")

	user.UserPhone = userPhone
	user.UserPasswd = userPhone
	user.UserName = userName
	user.UserSex = userSex
	user.UserAvatar = userAvatar
	user.UserTitle = userTitle
	user.UserFriend = userFriend
	user.UserTime = userTime

	_, err = dbObj.Insert(&user)

	return user, err
}
