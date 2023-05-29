package boot

type MysqlRepository struct {
}

func NewMysqlRepository() Repository {
	return MysqlRepository{}
}
