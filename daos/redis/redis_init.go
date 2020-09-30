package redis

//初始化redis
type RedisDao struct {
}

func NewRedisDao(con string) *RedisDao {
	return &RedisDao{}
}

//初始化mysql链接
func (this *RedisDao) Init() error {
	return nil
}
