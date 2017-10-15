package models

import (
	"github.com/astaxie/beego/orm"
)

type Recordinfo struct {
	Id int
	userPhone string `orm:"size(255)"`
	recordTime string `orm:"size(255)"`
	recordTitle string `orm:"size(255)"`
	recordImag string `orm:"size(255)"`
}

func init() {
	orm.RegisterModel(new(Recordinfo))
}


