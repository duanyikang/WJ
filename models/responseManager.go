package models

type ResponseBean struct {
	Result int `json:"result"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}