package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Recordinfo struct {
	Id          int    `pk:"auto" json:"id"`
	UserPhone   string `orm:"size(255)" json:"userPhone"`
	RecordTime  string `orm:"size(255)" json:"recordTime"`
	RecordTitle string `orm:"size(255)" json:"recordTitle"`
	RecordImage  string `orm:"size(255)" json:"recordImage"`
}

func init() {
	orm.RegisterModel(new(Recordinfo))
}

/**
写入数据
 */
func InsertRecord(userPhone, recrodTitle, recordImge string) (recordinfo Recordinfo, err error) {
	dbObj := orm.NewOrm()
	dbObj.Using("default")

	recordinfo.UserPhone = userPhone
	recordinfo.RecordTitle = recrodTitle
	recordinfo.RecordImage = recordImge
	t := time.Now().Unix()
	recordinfo.RecordTime = time.Unix(t, 0).Format("2006_01_02 15:04:05")
	_, err = dbObj.Insert(&recordinfo)
	return recordinfo, err
}

/**
检索一个用户的所有数据，暂时先不考虑分页
 */
func SelectRecord(userPhone string) (recordinfos []Recordinfo, err error) {
	dbObj := orm.NewOrm()
	dbObj.Using("default")

	_,err = dbObj.QueryTable("recordinfo").Filter("user_phone", userPhone).All(&recordinfos)
	return recordinfos, err
}
