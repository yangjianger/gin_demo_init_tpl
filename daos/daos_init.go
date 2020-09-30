package daos

import (
	"blog_tpl/daos/es"
	"blog_tpl/daos/mysql"
	"blog_tpl/daos/redis"
	"fmt"
)

type initFunc func() error

//初始化dao层数据
func InitDaos() {

	initFunMap := make(map[string]initFunc)

	//mysql 初始化
	mysqldao := mysql.NewMysqlDao("aa")
	initFunMap["mysql"] = mysqldao.Init

	//redis 初始化
	redisdao := redis.NewRedisDao("aa")
	initFunMap["redis"] = redisdao.Init

	//初始化es
	esdao := es.NewEsDao("aa")
	initFunMap["es"] = esdao.Init
}

//执行初始化函数
func execInit(initFunMap map[string]initFunc) {
	var err error
	for key, iFunc := range initFunMap {
		if err = iFunc(); err != nil {
			panic(fmt.Sprintf("%s, 初始化失败，错误信息是：%s", key, err))
		}
	}
}
