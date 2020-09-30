package models

type User struct {
	UserName string `json:"username"`
	Age      int    `json:"age"`
	Sex      string `json:"sex"`
	Addr     string `json:"addr"`
}
