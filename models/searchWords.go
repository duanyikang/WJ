package models

import (
	"net/url"
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/hex"
	"crypto/md5"
)

func Search(words string) (value string) {
	data := make(url.Values)
	data["q"] = []string{words}
	data["from"] = []string{"auto"}
	data["to"] = []string{"auto"}
	data["appKey"] = []string{"008fa90b96e19eca"}
	data["salt"] = []string{"666"}

	var s string
	s = "008fa90b96e19eca" + words + "666" + "YLKADTcEjUaIxkRJUPyPunxXrIENnhd5"
	sin := MD5(s)
	data["sign"] = []string{sin}
	//把post表单发送给目标服务器
	res, err := http.PostForm("http://openapi.youdao.com/api", data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println("post send success")

	return string(body)

}

func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
