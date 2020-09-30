package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/spf13/viper"
)

//密码加密
func Md5(str string) string {
	h := md5.New()
	//密码加盐
	h.Write([]byte(viper.GetString("app.secret")))
	//加密
	return hex.EncodeToString(h.Sum([]byte(str)))
}

//普通的MD5 不加盐
func Md5Normal(str string) string {
	h := md5.New()
	//加密
	return hex.EncodeToString(h.Sum([]byte(str)))
}
