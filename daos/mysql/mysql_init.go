package mysql

import (
	"fmt"
	"github.com/spf13/viper"
)

var db string

type MysqlDao struct {
}

func NewMysqlDao(con string) *MysqlDao {
	return &MysqlDao{}
}

//初始化mysql链接
func (this *MysqlDao) Init() error {

	mysqlMap := viper.Get("mysql").(map[string]interface{})
	fmt.Println("db - ", mysqlMap["db"])

	db = "aaa"
	return nil
}
