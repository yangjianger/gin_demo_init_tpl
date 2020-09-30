package es

//初始化es
type EsDao struct {
}

func NewEsDao(con string) *EsDao {
	return &EsDao{}
}

//初始化es链接
func (this *EsDao) Init() error {
	return nil
}
